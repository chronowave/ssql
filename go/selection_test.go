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

func TestFind(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Statement
		wantErr bool
	}{
		{"group-by", args{"find group-by($b) where [$b /adf/adf]"},
			&Statement{
				Find: []*Attribute{{Name: "b", Group: true}},
				Where: []*Expr{{
					Field: &Expr_Tuple{Tuple: &Tuple{
						Name: "b",
						Path: "/adf/adf",
					}},
				}}}, true,
		},
		{"avg", args{"find group-by($b), avg($d) where [$b /adf/adf] [$d /df]"},
			&Statement{
				Find: []*Attribute{
					{Name: "b", Group: true},
					{Name: "d", Func: &Function{Name: Function_AVG}},
				},
				Where: []*Expr{
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "b",
							Path: "/adf/adf",
						}},
					},
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "d",
							Path: "/df",
						}},
					},
				}}, false,
		},
		{"max", args{"find group-by($b), max($d) where [$b /adf/adf] [$d /df]"},
			&Statement{
				Find: []*Attribute{
					{Name: "b", Group: true},
					{Name: "d", Func: &Function{Name: Function_MAX}},
				},
				Where: []*Expr{
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "b",
							Path: "/adf/adf",
						}},
					},
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "d",
							Path: "/df",
						}},
					},
				}}, false,
		},
		{"min", args{"find group-by($b), min($d) where [$b /adf/adf] [$d /df]"},
			&Statement{
				Find: []*Attribute{
					{Name: "b", Group: true},
					{Name: "d", Func: &Function{Name: Function_MIN}},
				},
				Where: []*Expr{
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "b",
							Path: "/adf/adf",
						}},
					},
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "d",
							Path: "/df",
						}},
					},
				}}, false,
		},
		{"count", args{"find group-by($b), count($d) where [$b /adf/adf] [$d /df]"},
			&Statement{
				Find: []*Attribute{
					{Name: "b", Group: true},
					{Name: "d", Func: &Function{Name: Function_COUNT}},
				},
				Where: []*Expr{
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "b",
							Path: "/adf/adf",
						}},
					},
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "d",
							Path: "/df",
						}},
					},
				}}, false,
		},
		{"sum,count", args{`find group-by($a), sum($c), count($c)
			from flight_test
			where [$a /field3/field1 gt(1)]
		[$c /field4/field2 between(0.3, 180.2)]`},
			&Statement{
				Find: []*Attribute{
					{Name: "a", Group: true},
					{Name: "c", Func: &Function{Name: Function_SUM}},
					{Name: "c", Func: &Function{Name: Function_COUNT}},
				},
				Where: []*Expr{
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "b",
							Path: "/adf/adf",
						}},
					},
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "d",
							Path: "/df",
						}},
					},
				}}, false,
		},
		{"pctl", args{"find group-by($b), pctl($d, 0.6) where [$b /adf/adf] [$d /df]"},
			&Statement{
				Find: []*Attribute{
					{Name: "b", Group: true},
					{Name: "d", Func: &Function{
						Name: Function_PERCENTILE,
						Parameter: []*Parameter{
							{
								Value: &Parameter_Double{
									Double: 0.6,
								},
							},
						},
					}},
				},
				Where: []*Expr{
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "b",
							Path: "/adf/adf",
						}},
					},
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "d",
							Path: "/df",
						}},
					},
				}}, false,
		},
		{"part", args{"find group-by(part($b, 20)), pctl($d, 0.6) where [$b /adf/adf] [$d /df]"},
			&Statement{
				Find: []*Attribute{
					{Name: "b", Group: true, Func: &Function{
						Name: Function_PARTITION,
						Parameter: []*Parameter{
							{
								Value: &Parameter_Int{
									Int: 20,
								},
							},
						},
					}},
					{Name: "d", Func: &Function{
						Name: Function_PERCENTILE,
						Parameter: []*Parameter{
							{
								Value: &Parameter_Double{
									Double: 0.6,
								},
							},
						},
					}},
				},
				Where: []*Expr{
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "b",
							Path: "/adf/adf",
						}},
					},
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "d",
							Path: "/df",
						}},
					},
				}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if tt.wantErr {
				if len(errs) == 0 {
					t.Error("parse() doesn't produce expected errors")
				}
			} else if len(errs) > 0 {
				t.Errorf("parse() error = %v", errs)
			}

			if !reflect.DeepEqual(got.Find, tt.want.Find) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
