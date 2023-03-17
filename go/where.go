/*
 *  Copyright 2020 ChronoWave Authors
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
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"github.com/chronowave/ssql/go/gen"
)

func (p *parser) VisitExpression(ctx *gen.ExpressionContext) interface{} {
	for _, t := range ctx.AllTuple() {
		p.stmt.Where = append(p.stmt.Where, p.Visit(t.(antlr.ParseTree)).([]*Expr)...)
	}
	return nil
}

func (p *parser) VisitTuple(ctx *gen.TupleContext) interface{} {
	if ctx.Vector() != nil {
		return p.VisitVector(ctx.Vector().(*gen.VectorContext))
	}
	if ctx.Or() != nil {
		return p.VisitOr(ctx.Or().(*gen.OrContext))
	}
	if ctx.And() != nil {
		return p.VisitAnd(ctx.And().(*gen.AndContext))
	}
	return nil
}

func (p *parser) VisitVector(ctx *gen.VectorContext) interface{} {
	var tuple *Tuple
	if pred := ctx.Predicate(); pred != nil {
		// predicate
		tuple = p.VisitPredicate(pred.(*gen.PredicateContext)).(*Tuple)
	} else if nested := ctx.AllVector(); len(nested) > 0 {
		// nested
		var expr []*Expr
		for _, n := range nested {
			expr = append(expr, p.Visit(n.(antlr.ParseTree)).([]*Expr)...)
		}

		tuple = &Tuple{
			Predicate: &Tuple_Nested{
				Nested: &Nested{Expr: expr},
			},
		}
	} else {
		// selection only
		tuple = &Tuple{}
	}

	if id := ctx.IDENTIFIER(); id != nil {
		tuple.Name = extractVariableName(id.GetText())
	}

	if path := ctx.PATH(); path != nil {
		tuple.Path = path.GetText()
	}

	return []*Expr{{
		Field: &Expr_Tuple{
			Tuple: tuple,
		},
	}}
}

// Visit a parse tree produced by stmtLParser#OR.
func (p *parser) VisitOr(ctx *gen.OrContext) interface{} {
	var expr []*Expr
	for _, t := range ctx.AllTuple() {
		expr = append(expr, p.VisitTuple(t.(*gen.TupleContext)).([]*Expr)...)
	}

	return []*Expr{{
		Field: &Expr_Or{
			Or: &OR{
				Expr: expr,
			},
		},
	}}
}

// Visit a parse tree produced by stmtLParser#AND.
func (p *parser) VisitAnd(ctx *gen.AndContext) interface{} {
	var expr []*Expr
	for _, t := range ctx.AllTuple() {
		expr = append(expr, p.VisitTuple(t.(*gen.TupleContext)).([]*Expr)...)
	}

	return expr
}

func (p *parser) VisitPredicate(ctx *gen.PredicateContext) interface{} {
	n := ctx.GetChild(0)
	switch n.(type) {
	case *gen.EqContext:
		return p.VisitEq(n.(*gen.EqContext))
	case *gen.NeqContext:
		return p.VisitNeq(n.(*gen.NeqContext))
	case *gen.GtContext:
		return p.VisitGt(n.(*gen.GtContext))
	case *gen.GeContext:
		return p.VisitGe(n.(*gen.GeContext))
	case *gen.LtContext:
		return p.VisitLt(n.(*gen.LtContext))
	case *gen.LeContext:
		return p.VisitLe(n.(*gen.LeContext))
	case *gen.InContext:
		return p.VisitIn(n.(*gen.InContext))
	case *gen.BetweenContext:
		return p.VisitBetween(n.(*gen.BetweenContext))
	case *gen.ContainContext:
		return p.VisitContain(n.(*gen.ContainContext))
	case *gen.ExistContext:
		return p.VisitExist(n.(*gen.ExistContext))
	case *gen.IsnullContext:
		return p.VisitIsnull(n.(*gen.IsnullContext))
	case *gen.BooleanContext:
		return p.VisitBoolean(n.(*gen.BooleanContext))
	}

	return nil
}
