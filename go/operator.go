/*
 *  Copyright 2022 Chronowave Authors
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  Package parser declares an expression parser with support for macro
 *  expansion.
 */

package ssql

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"github.com/chronowave/ssql/go/gen"
)

func parseString(node antlr.TerminalNode, op Operator) interface{} {
	var text string
	if node != nil {
		text = stripQuote(node.GetText())
	}

	return &Tuple{
		Predicate: &Tuple_Unary{
			&Unary{
				Op:      op,
				Operand: &Operand{Content: text},
			},
		},
	}
}

func (p *parser) visitUnaryOperator(ctx *gen.ScalarContext, operator Operator) interface{} {
	return &Tuple{
		Predicate: &Tuple_Unary{
			&Unary{
				Op:      operator,
				Operand: p.VisitScalar(ctx).(*Operand),
			},
		},
	}
}

func (p *parser) VisitEq(ctx *gen.EqContext) interface{} {
	return p.visitUnaryOperator(ctx.Scalar().(*gen.ScalarContext), Operator_EQ)
}

func (p *parser) VisitNeq(ctx *gen.NeqContext) interface{} {
	return p.visitUnaryOperator(ctx.Scalar().(*gen.ScalarContext), Operator_NEQ)
}

func (p *parser) VisitGt(ctx *gen.GtContext) interface{} {
	return p.visitUnaryOperator(ctx.Scalar().(*gen.ScalarContext), Operator_GT)
}

func (p *parser) VisitGe(ctx *gen.GeContext) interface{} {
	return p.visitUnaryOperator(ctx.Scalar().(*gen.ScalarContext), Operator_GE)
}

func (p *parser) VisitLt(ctx *gen.LtContext) interface{} {
	return p.visitUnaryOperator(ctx.Scalar().(*gen.ScalarContext), Operator_LT)
}

func (p *parser) VisitLe(ctx *gen.LeContext) interface{} {
	return p.visitUnaryOperator(ctx.Scalar().(*gen.ScalarContext), Operator_LE)
}

func (p *parser) VisitIn(ctx *gen.InContext) interface{} {
	list := ctx.List()
	if list == nil {
		return &Tuple{
			Predicate: &Tuple_List{
				&List{
					Operands: []*Operand{},
				},
			},
		}
	}

	return &Tuple{
		Predicate: &Tuple_List{
			&List{
				Operands: p.VisitList(list.(*gen.ListContext)).([]*Operand),
			},
		},
	}
}

func (p *parser) VisitBetween(ctx *gen.BetweenContext) interface{} {
	operands := []*Operand{{}, {}}
	tuple := &Tuple{
		Predicate: &Tuple_Binary{
			Binary: &Binary{
				First:  operands[0],
				Second: operands[1],
			},
		},
	}

	signs := ctx.AllSIGN()
	nodes := ctx.AllINTEGER()
	if len(nodes) != 2 {
		nodes = ctx.AllREAL_NUMBER()
	}
	if len(nodes) != 2 {
		nodes = ctx.AllISO8601_DATE_TIME()
	}

	if len(nodes) != 2 {
		token := ctx.GetStart()
		p.errors = append(p.errors, Error{
			Line:    token.GetLine(),
			Column:  token.GetColumn(),
			Message: "wrong format, accepts number or iso8601 datetime",
		})
	} else {
		j := 0
		for i, n := range nodes {
			var sign string
			if len(signs) > 0 && n.GetSymbol().GetColumn() > signs[j].GetSymbol().GetColumn() {
				sign = signs[j].GetText()
				j++
			}
			operands[i].Content = sign + n.GetText()
		}
	}

	return tuple
}

// Visit a parse tree produced by SSQLParser#contain.
func (p *parser) VisitContain(ctx *gen.ContainContext) interface{} {
	return parseString(ctx.STRING(), Operator_CONTAIN)
}

// Visit a parse tree produced by SSQLParser#exist.
func (p *parser) VisitExist(_ *gen.ExistContext) interface{} {
	return &Tuple{
		Predicate: &Tuple_Unary{
			&Unary{
				Op: Operator_EXIST,
			},
		},
	}
}

func (p *parser) VisitIsnull(_ *gen.IsnullContext) interface{} {
	return &Tuple{
		Predicate: &Tuple_Unary{
			Unary: &Unary{
				Op: Operator_ISNULL,
			},
		},
	}
}

func (p *parser) VisitBoolean(ctx *gen.BooleanContext) interface{} {
	return &Tuple{
		Predicate: &Tuple_Unary{
			Unary: &Unary{
				Op: Operator_EQ,
				Operand: &Operand{
					Value: &Operand_Bool{
						Bool: ctx.TRUE() != nil,
					},
				},
			},
		},
	}
}

func (p *parser) VisitScalar(ctx *gen.ScalarContext) interface{} {
	var content string
	if s := ctx.SIGN(); s != nil {
		content = s.GetText()
	}

	if r := ctx.REAL_NUMBER(); r != nil {
		content += r.GetText()
	} else if integers := ctx.AllINTEGER(); len(integers) == 1 {
		content += integers[0].GetText()
	} else if date := ctx.ISO8601_DATE_TIME(); date != nil {
		content = date.GetText()
	} else {
		token := ctx.GetStart()
		p.errors = append(p.errors, Error{
			Line:    token.GetLine(),
			Column:  token.GetColumn(),
			Message: "wrong format, accepts number or iso8601 datetime",
		})
	}

	return &Operand{Content: content}
}

func (p *parser) VisitList(ctx *gen.ListContext) interface{} {
	str := ctx.StringList()
	if str != nil {
		return p.VisitStringList(str.(*gen.StringListContext))
	}

	double := ctx.DoubleList()
	if double != nil {
		return p.VisitDoubleList(double.(*gen.DoubleListContext))
	}

	integer := ctx.IntList()
	if integer != nil {
		return p.VisitIntList(integer.(*gen.IntListContext))
	}

	token := ctx.GetStart()
	p.errors = append(p.errors, Error{
		Line:    token.GetLine(),
		Column:  token.GetColumn(),
		Message: "operator in must contain a list of string, double or integer",
	})

	return []*Operand{}
}

func (p *parser) VisitStringList(ctx *gen.StringListContext) interface{} {
	nodes := ctx.AllSTRING()
	operands := make([]*Operand, len(nodes))
	for i, n := range nodes {
		operands[i] = &Operand{Content: stripQuote(n.GetText())}
	}

	return operands
}

func (p *parser) VisitDoubleList(ctx *gen.DoubleListContext) interface{} {
	nodes := ctx.AllREAL_NUMBER()
	operands := make([]*Operand, len(nodes))
	for i, n := range nodes {
		_, err := strconv.ParseFloat(n.GetText(), 64)
		if err != nil {
			token := n.GetSymbol()
			p.errors = append(p.errors, Error{
				Line:    token.GetLine(),
				Column:  token.GetColumn(),
				Message: err.Error(),
			})
			break
		}
		operands[i] = &Operand{Content: n.GetText()}
	}

	return operands
}

func (p *parser) VisitIntList(ctx *gen.IntListContext) interface{} {
	nodes := ctx.AllINTEGER()
	operands := make([]*Operand, len(nodes))
	for i, n := range nodes {
		_, err := strconv.ParseInt(n.GetText(), 10, 64)
		if err != nil {
			token := n.GetSymbol()
			p.errors = append(p.errors, Error{
				Line:    token.GetLine(),
				Column:  token.GetColumn(),
				Message: err.Error(),
			})
			break
		}
		operands[i] = &Operand{Content: n.GetText()}
	}

	return operands
}

func stripQuote(text string) string {
	if len(text) < 2 {
		return text
	}
	return text[1 : len(text)-1]
}
