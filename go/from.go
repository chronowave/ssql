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
	"strings"

	"github.com/chronowave/ssql/go/gen"
)

func (p *parser) VisitFrom(ctx *gen.FromContext) interface{} {
	p.stmt.Aux.LocalOnly = ctx.LOCAL() != nil

	if ctx.NAME() == nil {
		token := ctx.GetStart()
		p.errors = append(p.errors, Error{
			Line:    token.GetLine(),
			Column:  token.GetColumn(),
			Message: "missing name in query statement FROM clause",
		})
	} else {
		p.stmt.From = strings.TrimSpace(ctx.NAME().GetText())
	}

	return nil
}
