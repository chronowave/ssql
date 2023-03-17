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

	"github.com/chronowave/ssql/go/gen"
)

func (p *parser) VisitSelection(ctx *gen.SelectionContext) interface{} {
	for _, n := range ctx.AllAttribute() {
		p.stmt.Find = append(p.stmt.Find, p.VisitAttribute(n.(*gen.AttributeContext)).([]*Attribute)...)
	}

	for _, n := range ctx.AllGroupBy() {
		p.stmt.Find = append(p.stmt.Find, p.VisitGroupBy(n.(*gen.GroupByContext)).(*Attribute))
	}

	for _, n := range ctx.AllAggregate() {
		p.stmt.Find = append(p.stmt.Find, p.VisitAggregate(n.(*gen.AggregateContext)).(*Attribute))
	}

	return nil
}

func (p *parser) VisitAttribute(ctx *gen.AttributeContext) interface{} {
	if f := ctx.Aggregate(); f != nil {
		return p.VisitAggregate(f.(*gen.AggregateContext))
	}

	return []*Attribute{{
		Name: extractVariableName(ctx.IDENTIFIER().GetText()),
	}}
}

func (p *parser) VisitAggregate(ctx *gen.AggregateContext) interface{} {
	if pctl := ctx.Percentile(); pctl != nil {
		return p.VisitPercentile(pctl.(*gen.PercentileContext))
	}

	// AVG | MAX | MIN | SUM | COUNT
	var fn Function_Func
	if ctx.AVG() != nil {
		fn = Function_AVG
	} else if ctx.MAX() != nil {
		fn = Function_MAX
	} else if ctx.MIN() != nil {
		fn = Function_MIN
	} else if ctx.SUM() != nil {
		fn = Function_SUM
	} else if ctx.COUNT() != nil {
		fn = Function_COUNT
	} else {
		token := ctx.GetStart()
		p.errors = append(p.errors, Error{
			Line:    token.GetLine(),
			Column:  token.GetColumn(),
			Message: "missing aggregate function from group by statement",
		})
		return &Attribute{}
	}

	return &Attribute{
		Name: extractVariableName(ctx.IDENTIFIER().GetText()),
		Func: &Function{Name: fn},
	}
}

func (p *parser) VisitPercentile(ctx *gen.PercentileContext) interface{} {
	v, err := strconv.ParseFloat(ctx.REAL_NUMBER().GetText(), 64)
	if err != nil {
		// log err
		return nil
	}

	return &Attribute{
		Name: extractVariableName(ctx.IDENTIFIER().GetText()),
		Func: &Function{
			Name: Function_PERCENTILE,
			Parameter: []*Parameter{
				{
					Value: &Parameter_Double{
						Double: v,
					},
				},
			},
		},
	}
}

func (p *parser) VisitGroupBy(ctx *gen.GroupByContext) interface{} {
	if part := ctx.Partition(); part != nil {
		return p.VisitPartition(part.(*gen.PartitionContext))
	}

	var name string
	if ctx.IDENTIFIER() == nil {
		token := ctx.GetStart()
		p.errors = append(p.errors, Error{
			Line:    token.GetLine(),
			Column:  token.GetColumn(),
			Message: "missing group by attribute name",
		})
	} else {
		name = extractVariableName(ctx.IDENTIFIER().GetText())
	}

	return &Attribute{Name: name, Group: true}
}

func (p *parser) VisitPartition(ctx *gen.PartitionContext) interface{} {
	v, err := strconv.ParseInt(ctx.INTEGER().GetText(), 10, 64)
	if err != nil {
		token := ctx.GetStart()
		p.errors = append(p.errors, Error{
			Line:    token.GetLine(),
			Column:  token.GetColumn(),
			Message: "partition value must be an integer",
		})
	}

	return &Attribute{
		Name: extractVariableName(ctx.IDENTIFIER().GetText()),
		Func: &Function{
			Name: Function_PARTITION,
			Parameter: []*Parameter{
				&Parameter{
					Value: &Parameter_Int{
						Int: v,
					},
				},
			},
		},
		Group: true,
	}
}

func extractVariableName(name string) string {
	if name[0] == '$' {
		return name[1:]
	}

	return name
}
