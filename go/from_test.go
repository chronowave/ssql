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
	"reflect"
	"testing"
)

func TestFrom(t *testing.T) {
	stmt, errs := Parse(`find group-by(part($c,10),$b), avg($g)
            from table_8
			where [$b /adf/adf eq(10)]
				  [$g /dkf/adf]
				  [/adf/dkf [/df between(15, 20)]
							[/dd eq(2.5)]
							[/adf eq(2.5)]
				  ]`)
	if errs != nil {
		t.Errorf("unexpected error %v", errs)
	}

	if !reflect.DeepEqual(stmt.From, "table_8") {
		t.Errorf("statement from doesn't match expected=table_8, got=%s", stmt.From)
	}
}
