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

func TestEqual(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Statement
		wantErr bool
	}{
		{"eq_int", args{"find $b where [$b /adf/adf eq(2)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_EQ,
									Operand: &Operand{
										Content: "2",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"eq_double", args{"find $b where [$b /adf/adf eq(2.5)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
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
				}}}, false,
		},
		{"eq_string", args{"find $b where [$b /adf/adf eq('abc')]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_EQ,
									Operand: &Operand{
										Content: "abc",
									},
								},
							},
						},
					},
				}}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if tt.wantErr {
				if len(errs) == 0 {
					t.Error("parse() doesn't produce expected errors")
				} else {
					t.Logf("expected errors: %v", errs)
				}
			} else if len(errs) > 0 {
				t.Errorf("parse() error = %v", errs)
			} else if !reflect.DeepEqual(got.Where, tt.want.Where) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotEqual(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Statement
		wantErr bool
	}{
		{"neq_int", args{"find $b where [$b /adf/adf neq(2)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_NEQ,
									Operand: &Operand{
										Content: "2",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"neq_double", args{"find $b where [$b /adf/adf neq(2.5)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_NEQ,
									Operand: &Operand{
										Content: "2.5",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"neq_date", args{"find $b where [$b /adf/adf neq(2006-07-08)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_NEQ,
									Operand: &Operand{
										Content: "2006-07-08",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"neq_string", args{"find $b where [$b /adf/adf neq('abc')]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_NEQ,
									Operand: &Operand{
										Content: "abc",
									},
								},
							},
						},
					},
				}}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if tt.wantErr {
				if len(errs) == 0 {
					t.Error("parse() doesn't produce expected errors")
				} else {
					t.Logf("expected errors: %v", errs)
				}
			} else if len(errs) > 0 {
				t.Errorf("parse() error = %v", errs)
			} else if !reflect.DeepEqual(got.Where, tt.want.Where) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGT(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Statement
		wantErr bool
	}{
		{"gt_int", args{"find $b where [$b /adf/adf gt(2)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_GT,
									Operand: &Operand{
										Content: "2",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"gt_double", args{"find $b where [$b /adf/adf gt(2.5)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_GT,
									Operand: &Operand{
										Content: "2.5",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"gt_string", args{"find $b where [$b /adf/adf gt('abc')]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_GT,
									Operand: &Operand{
										Content: "abc",
									},
								},
							},
						},
					},
				}}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if tt.wantErr {
				if len(errs) == 0 {
					t.Error("parse() doesn't produce expected errors")
				} else {
					t.Logf("expected errors: %v", errs)
				}
			} else if len(errs) > 0 {
				t.Errorf("parse() error = %v", errs)
			} else if !reflect.DeepEqual(got.Where, tt.want.Where) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGE(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Statement
		wantErr bool
	}{
		{"ge_int", args{"find $b where [$b /adf/adf ge(2)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_GE,
									Operand: &Operand{
										Content: "2",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"ge_double", args{"find $b where [$b /adf/adf ge(2.5)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_GE,
									Operand: &Operand{
										Content: "2.5",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"ge_string", args{"find $b where [$b /adf/adf ge('abc')]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_GE,
									Operand: &Operand{
										Content: "abc",
									},
								},
							},
						},
					},
				}}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if tt.wantErr {
				if len(errs) == 0 {
					t.Error("parse() doesn't produce expected errors")
				} else {
					t.Logf("expected errors: %v", errs)
				}
			} else if len(errs) > 0 {
				t.Errorf("parse() error = %v", errs)
			} else if !reflect.DeepEqual(got.Where, tt.want.Where) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLT(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Statement
		wantErr bool
	}{
		{"lt_int", args{"find $b where [$b /adf/adf lt(2)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_LT,
									Operand: &Operand{
										Content: "2",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"lt_double", args{"find $b where [$b /adf/adf lt(2.5)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_LT,
									Operand: &Operand{
										Content: "2.5",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"lt_string", args{"find $b where [$b /adf/adf lt('abc')]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_LT,
									Operand: &Operand{
										Content: "abc",
									},
								},
							},
						},
					},
				}}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if tt.wantErr {
				if len(errs) == 0 {
					t.Error("parse() doesn't produce expected errors")
				} else {
					t.Logf("expected errors: %v", errs)
				}
			} else if len(errs) > 0 {
				t.Errorf("parse() error = %v", errs)
			} else if !reflect.DeepEqual(got.Where, tt.want.Where) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLE(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Statement
		wantErr bool
	}{
		{"le_int", args{"find $b where [$b /adf/adf le(2)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_LE,
									Operand: &Operand{
										Content: "2",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"le_double", args{"find $b where [$b /adf/adf le(2.5)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_LE,
									Operand: &Operand{
										Content: "2.5",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"le_string", args{"find $b where [$b /adf/adf le('abc')]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_LE,
									Operand: &Operand{
										Content: "abc",
									},
								},
							},
						},
					},
				}}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if tt.wantErr {
				if len(errs) == 0 {
					t.Error("parse() doesn't produce expected errors")
				} else {
					t.Logf("expected errors: %v", errs)
				}
			} else if len(errs) > 0 {
				t.Errorf("parse() error = %v", errs)
			} else if !reflect.DeepEqual(got.Where, tt.want.Where) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIN(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Statement
		wantErr bool
	}{
		{"in_int", args{"find $b where [$b /adf/adf in(2, 2)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_List{
								&List{
									Operands: []*Operand{
										{
											Content: "2",
										},
										{
											Content: "2",
										},
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"in_double", args{"find $b where [$b /adf/adf in(2.5, 2.0)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_List{
								&List{
									Operands: []*Operand{
										{
											Content: "2.5",
										},
										{
											Content: "2.0",
										},
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"in_string", args{"find $b where [$b /adf/adf in('abc', 'adf')]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_List{
								&List{
									Operands: []*Operand{
										{
											Content: "abc",
										},
										{
											Content: "adf",
										},
									},
								},
							},
						},
					},
				}}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if tt.wantErr {
				if len(errs) == 0 {
					t.Error("parse() doesn't produce expected errors")
				} else {
					t.Logf("expected errors: %v", errs)
				}
			} else if len(errs) > 0 {
				t.Errorf("parse() error = %v", errs)
			} else if !reflect.DeepEqual(got.Where, tt.want.Where) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBetween(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Statement
		wantErr bool
	}{
		{"between_int", args{"find $b where [$b /adf/adf between(2, 3)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Binary{
								&Binary{
									First: &Operand{
										Content: "2",
									},
									Second: &Operand{
										Content: "3",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"between_double", args{"find $b where [$b /adf/adf between(2.5, 5.0)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Binary{
								&Binary{
									First: &Operand{
										Content: "2.5",
									},
									Second: &Operand{
										Content: "5.0",
									},
								},
							},
						},
					},
				}}}, false,
		},
		{"between_string", args{"find $b where [$b /adf/adf between('abc', 'adf')]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Binary{
								&Binary{
									First: &Operand{
										Content: "abc",
									},
									Second: &Operand{
										Content: "adf",
									},
								},
							},
						},
					},
				}}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if tt.wantErr {
				if len(errs) == 0 {
					t.Error("parse() doesn't produce expected errors")
				} else {
					t.Logf("expected errors: %v", errs)
				}
			} else if len(errs) > 0 {
				t.Errorf("parse() error = %v", errs)
			} else if !reflect.DeepEqual(got.Where, tt.want.Where) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContain(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Statement
		wantErr bool
	}{
		{"contain_int", args{"find $b where [$b /adf/adf contain(2)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_CONTAIN,
									Operand: &Operand{
										Content: "2",
									},
								},
							},
						},
					},
				}}}, true,
		},
		{"contain_double", args{"find $b where [$b /adf/adf contain(5.0)]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_CONTAIN,
									Operand: &Operand{
										Content: "5.0",
									},
								},
							},
						},
					},
				}}}, true,
		},
		{"contain_string", args{"find $b where [$b /adf/adf contain('abc')]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_CONTAIN,
									Operand: &Operand{
										Content: "abc",
									},
								},
							},
						},
					},
				}}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if tt.wantErr {
				if len(errs) == 0 {
					t.Error("parse() doesn't produce expected errors")
				} else {
					t.Logf("expected errors: %v", errs)
				}
			} else if len(errs) > 0 {
				t.Errorf("parse() error = %v", errs)
			} else if !reflect.DeepEqual(got.Where, tt.want.Where) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExist(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Statement
		wantErr bool
	}{
		{"exist ()", args{"find $b where [$b /adf/adf exist()]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_EXIST,
								},
							},
						},
					},
				}}}, false,
		},
		{"exist", args{"find $b where [$b /adf/adf exist]"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{
						&Tuple{
							Name: "b",
							Path: "/adf/adf",
							Predicate: &Tuple_Unary{
								&Unary{
									Op: Operator_EXIST,
								},
							},
						},
					},
				}}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, errs := Parse(tt.args.query)
			if tt.wantErr {
				if len(errs) == 0 {
					t.Error("parse() doesn't produce expected errors")
				} else {
					t.Logf("expected errors: %v", errs)
				}
			} else if len(errs) > 0 {
				t.Errorf("parse() error = %v", errs)
			} else if !reflect.DeepEqual(got.Where, tt.want.Where) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
