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
	"fmt"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"github.com/chronowave/ssql/go/gen"
)

type parser struct {
	gen.BaseSSQLVisitor
	errors []Error
	stmt   *Statement
}

var (
	lexers = &sync.Pool{
		New: func() interface{} {
			return gen.NewSSQLLexer(nil)
		},
	}

	parsers = &sync.Pool{
		New: func() interface{} {
			return gen.NewSSQLParser(nil)
		},
	}
)

func Parse(query string) (*Statement, []Error) {
	p := parser{stmt: &Statement{Aux: &Aux{}}}

	l := lexers.Get().(*gen.SSQLLexer)
	l.SetInputStream(antlr.NewInputStream(query))
	defer lexers.Put(l)

	s := parsers.Get().(*gen.SSQLParser)
	s.SetInputStream(antlr.NewCommonTokenStream(l, 0))
	defer parsers.Put(s)

	l.RemoveErrorListeners()
	s.RemoveErrorListeners()
	l.AddErrorListener(&p)
	s.AddErrorListener(&p)

	p.Visit(s.Start())

	if len(p.errors) == 0 && p.stmt != nil {
		var (
			aggregates []uint32
			groups     []uint32
		)
		for i, attr := range p.stmt.Find {
			if attr.Group {
				groups = append(groups, uint32(i))
			}

			if attr.Func != nil {
				aggregates = append(aggregates, uint32(i))
			}
		}

		if len(groups) > 0 && len(aggregates) == 0 {
			p.errors = append(p.errors, Error{
				Message: "missing aggregate functions from group by statement",
			})
		}

		p.stmt.Aux.AggregateAttributes = aggregates
		p.stmt.Aux.GroupAttributes = groups
	}

	return p.stmt, p.errors
}

// Visitor implementations.
func (p *parser) Visit(tree antlr.ParseTree) interface{} {
	switch tree.(type) {
	case *gen.StartContext:
		return p.VisitStart(tree.(*gen.StartContext))
	case *gen.SelectionContext:
		return p.VisitSelection(tree.(*gen.SelectionContext))
	case *gen.FromContext:
		return p.VisitFrom(tree.(*gen.FromContext))
	case *gen.ExpressionContext:
		return p.VisitExpression(tree.(*gen.ExpressionContext))
	case *gen.TupleContext:
		return p.VisitTuple(tree.(*gen.TupleContext))
	case *gen.VectorContext:
		return p.VisitVector(tree.(*gen.VectorContext))
	case *gen.OrContext:
		return p.VisitOr(tree.(*gen.OrContext))
	case *gen.AndContext:
		return p.VisitAnd(tree.(*gen.AndContext))
	case *gen.OrderByContext:
		return p.VisitOrderBy(tree.(*gen.OrderByContext))
	case *gen.LimitContext:
		return p.VisitLimit(tree.(*gen.LimitContext))
	case antlr.TerminalNode:
		return nil
	}

	// Report at least one error if the parser reaches an unknown parse element.
	// Typically, this happens if the parser has already encountered a syntax error elsewhere.
	if len(p.errors) > 0 {
		fmt.Printf("error : %v\n", p.errors)
	}
	panic("unknown parser state due to previous errors")
}

func (p *parser) VisitChildren(node antlr.RuleNode) interface{} {
	for _, n := range node.GetChildren() {
		p.Visit(n.(antlr.ParseTree))
	}
	return nil
}

func (p *parser) VisitStart(ctx *gen.StartContext) interface{} {
	return p.VisitChildren(ctx)
}
