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

func TestOrderBy(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Statement
		wantErr bool
	}{
		{"desc", args{"find $b where [$b /] order-by $b desc"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{Tuple: &Tuple{
						Name: "b",
						Path: "/",
					}},
				}},
				OrderBy: []*OrderBy{{
					Name:      "b",
					Direction: OrderBy_DESC,
				}},
			}, false,
		},
		{"asc", args{"find $b where [$b /] order-by $b"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{Tuple: &Tuple{
						Name: "b",
						Path: "/",
					}},
				}},
				OrderBy: []*OrderBy{{
					Name:      "b",
					Direction: OrderBy_ASC,
				}},
			}, false,
		},
		{"double", args{"find $b where [$b /] order-by $b desc, $b asc"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{Tuple: &Tuple{
						Name: "b",
						Path: "/",
					}},
				}},
				OrderBy: []*OrderBy{{
					Name:      "b",
					Direction: OrderBy_DESC,
				},
					{
						Name:      "b",
						Direction: OrderBy_ASC,
					},
				},
			}, false,
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

			if !reflect.DeepEqual(got.OrderBy, tt.want.OrderBy) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLimit(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Statement
		wantErr bool
	}{
		{"limit", args{"find $b where [$b /] order-by $b desc limit 30"},
			&Statement{
				Find: []*Attribute{{Name: "b"}},
				Where: []*Expr{{
					Field: &Expr_Tuple{Tuple: &Tuple{
						Name: "b",
						Path: "/",
					}},
				}},
				OrderBy: []*OrderBy{{
					Name:      "b",
					Direction: OrderBy_DESC,
				}},
				Limit: 30,
			}, false,
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

			if !reflect.DeepEqual(got.Limit, tt.want.Limit) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
