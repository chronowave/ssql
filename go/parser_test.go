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

//go:generate protoc --go_out=./ --proto_path=../ ssql.proto
//go:generate cp antlr4/SSQL.g4 gen/
//go:generate ./antlr4/gen -Dlanguage=Go -visitor -long-messages -package gen gen/SSQL.g4

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want *Statement
	}{
		{
			"complete",
			args{`find group-by(part($c,10),$b), avg($g)
                              where [$b /adf/adf eq(10)]
                                    [$g /dkf/adf]
                                    [/adf/dkf [/df between(15, 20)]
                                              [/adf eq(2.5)]
                                    ]`},
			&Statement{
				Find: []*Attribute{
					{
						Name:  "c",
						Group: true,
						Func: &Function{
							Name: Function_PARTITION,
							Parameter: []*Parameter{
								{
									Value: &Parameter_Int{
										Int: 10,
									},
								},
							},
						},
					},
					{
						Name:  "b",
						Group: true,
					},
					{
						Name: "g",
						Func: &Function{Name: Function_AVG},
					},
				},
				Where: []*Expr{
					{
						Field: &Expr_Tuple{
							&Tuple{
								Name: "b",
								Path: "/adf/adf",
								Predicate: &Tuple_Unary{
									Unary: &Unary{
										Op: Operator_EQ,
										Operand: &Operand{
											Content: "10",
										},
									},
								},
							},
						},
					},
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "g",
							Path: "/dkf/adf",
						}},
					},
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Path: "/adf/dkf",
							Predicate: &Tuple_Nested{
								Nested: &Nested{
									Expr: []*Expr{
										{
											Field: &Expr_Tuple{
												&Tuple{
													Path: "/df",
													Predicate: &Tuple_Binary{
														&Binary{
															First: &Operand{
																Content: "15",
															},
															Second: &Operand{
																Content: "20",
															},
														},
													},
												},
											},
										},
										{
											Field: &Expr_Tuple{
												&Tuple{
													Path: "/adf",
													Predicate: &Tuple_Unary{
														&Unary{
															Op: Operator_EQ,
															Operand: &Operand{
																Content: "2.5",
															},
														},
													},
												},
											},
										},
									}},
							},
						},
						},
					}},
				Aux: &Aux{
					AggregateAttributes: []uint32{0, 2},
					GroupAttributes:     []uint32{0, 1},
				},
			},
		},
		{
			"zero_int",
			args{`find $g where [$g /adf/adf eq(0)]`},
			&Statement{
				Find: []*Attribute{{Name: "g"}},
				Where: []*Expr{
					{
						Field: &Expr_Tuple{
							&Tuple{
								Name: "g",
								Path: "/adf/adf",
								Predicate: &Tuple_Unary{
									&Unary{
										Op: Operator_EQ,
										Operand: &Operand{
											Content: "0",
										},
									},
								},
							},
						},
					}},
				Aux: &Aux{},
			},
		},
		{
			"zeros_int",
			args{`find $g where [$g /adf/adf eq(000)]`},
			&Statement{
				Find: []*Attribute{{Name: "g"}},
				Where: []*Expr{
					{
						Field: &Expr_Tuple{
							&Tuple{
								Name: "g",
								Path: "/adf/adf",
								Predicate: &Tuple_Unary{
									&Unary{
										Op: Operator_EQ,
										Operand: &Operand{
											Content: "000",
										},
									},
								},
							},
						},
					}},
				Aux: &Aux{},
			},
		},
		{
			"zero_double",
			args{`find $g where [$g /adf/adf eq(0.0)]`},
			&Statement{
				Find: []*Attribute{{Name: "g"}},
				Where: []*Expr{
					{
						Field: &Expr_Tuple{
							&Tuple{
								Name: "g",
								Path: "/adf/adf",
								Predicate: &Tuple_Unary{
									&Unary{
										Op: Operator_EQ,
										Operand: &Operand{
											Content: "0.0",
										},
									},
								},
							},
						},
					}},
				Aux: &Aux{},
			},
		},
		{
			"nested",
			args{`find $g where [$r / [/adf/adf eq(0.0)]]`},
			&Statement{
				Find: []*Attribute{{Name: "g"}},
				Where: []*Expr{
					{
						Field: &Expr_Tuple{Tuple: &Tuple{
							Name: "r",
							Path: "/",
							Predicate: &Tuple_Nested{
								Nested: &Nested{
									Expr: []*Expr{{
										Field: &Expr_Tuple{
											&Tuple{
												Path: "/adf/adf",
												Predicate: &Tuple_Unary{
													&Unary{
														Op: Operator_EQ,
														Operand: &Operand{
															Content: "0.0",
														},
													},
												},
											},
										},
									}},
								},
							}}},
					},
				},
				Aux: &Aux{},
			},
		},
		{
			"or",
			args{`find $g where {[/adf/adf eq(0.0)] [/d lt(1)]}`},
			&Statement{
				Find: []*Attribute{{Name: "g"}},
				Where: []*Expr{{
					Field: &Expr_Or{Or: &OR{
						Expr: []*Expr{
							{
								Field: &Expr_Tuple{Tuple: &Tuple{
									Path: "/adf/adf",
									Predicate: &Tuple_Unary{
										&Unary{
											Op: Operator_EQ,
											Operand: &Operand{
												Content: "0.0",
											},
										},
									},
								}},
							},
							{
								Field: &Expr_Tuple{Tuple: &Tuple{
									Path: "/d",
									Predicate: &Tuple_Unary{
										&Unary{
											Op: Operator_LT,
											Operand: &Operand{
												Content: "1",
											},
										},
									},
								}},
							},
						}},
					},
				}},
				Aux: &Aux{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if len(errs) > 0 {
				t.Errorf("parsing errors %v", errs)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
