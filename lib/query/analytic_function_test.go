package query

import (
	"reflect"
	"testing"

	"github.com/mithrandie/csvq/lib/parser"
)

type analyticFunctionTest struct {
	Name     string
	View     *View
	Function parser.AnalyticFunction
	Result   *View
	Error    string
}

func testAnalyticFunction(t *testing.T, f func(*View, parser.AnalyticFunction) error, tests []analyticFunctionTest) {
	for _, v := range tests {
		ViewCache.Clear()
		err := f(v.View, v.Function)
		if err != nil {
			if len(v.Error) < 1 {
				t.Errorf("%s: unexpected error %q", v.Name, err)
			} else if err.Error() != v.Error {
				t.Errorf("%s: error %q, want error %q", v.Name, err.Error(), v.Error)
			}
			continue
		}
		if 0 < len(v.Error) {
			t.Errorf("%s: no error, want error %q", v.Name, v.Error)
			continue
		}
		if !reflect.DeepEqual(v.View, v.Result) {
			t.Errorf("%s: result = %q, want %q", v.Name, v.View, v.Result)
		}
	}
}

var rowNumberTests = []analyticFunctionTest{
	{
		Name: "RowNumber",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(3),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(4),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(5),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "row_number",
			AnalyticClause: parser.AnalyticClause{
				OrderByClause: parser.OrderByClause{
					Items: []parser.Expression{
						parser.OrderItem{
							Value: parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
						},
					},
				},
			},
		},
		Result: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(3),
					parser.NewInteger(3),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(4),
					parser.NewInteger(4),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(5),
					parser.NewInteger(5),
				}),
			},
		},
	},
	{
		Name: "RowNumber with Partition",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(3),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(4),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(5),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "row_number",
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
				OrderByClause: parser.OrderByClause{
					Items: []parser.Expression{
						parser.OrderItem{
							Value: parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
						},
					},
				},
			},
		},
		Result: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(3),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(4),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(5),
					parser.NewInteger(3),
				}),
			},
		},
	},
	{
		Name: "RowNumber Arguments Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "row_number",
			Args: []parser.Expression{
				parser.NewInteger(1),
			},
		},
		Error: "[L:- C:-] function row_number takes no argument",
	},
	{
		Name: "RowNumber Partition Value Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(3),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(4),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(5),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "row_number",
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "notexist"}},
					},
				},
				OrderByClause: parser.OrderByClause{
					Items: []parser.Expression{
						parser.OrderItem{
							Value: parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
						},
					},
				},
			},
		},
		Error: "[L:- C:-] field notexist does not exist",
	},
}

func TestRowNumber(t *testing.T) {
	testAnalyticFunction(t, RowNumber, rowNumberTests)
}

var rankTests = []analyticFunctionTest{
	{
		Name: "Rank",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "rank",
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
				OrderByClause: parser.OrderByClause{
					Items: []parser.Expression{
						parser.OrderItem{
							Value: parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
						},
					},
				},
			},
		},
		Result: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(2),
					parser.NewInteger(3),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
					parser.NewInteger(2),
				}),
			},
		},
	},
	{
		Name: "Rank Arguments Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "rank",
			Args: []parser.Expression{
				parser.NewInteger(1),
			},
		},
		Error: "[L:- C:-] function rank takes no argument",
	},
	{
		Name: "Rank Partition Value Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "rank",
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "notexist"}},
					},
				},
				OrderByClause: parser.OrderByClause{
					Items: []parser.Expression{
						parser.OrderItem{
							Value: parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
						},
					},
				},
			},
		},
		Error: "[L:- C:-] field notexist does not exist",
	},
	{
		Name: "Rank Order Value Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "rank",
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
				OrderByClause: parser.OrderByClause{
					Items: []parser.Expression{
						parser.OrderItem{
							Value: parser.FieldReference{Column: parser.Identifier{Literal: "notexist"}},
						},
					},
				},
			},
		},
		Error: "[L:- C:-] field notexist does not exist",
	},
}

func TestRank(t *testing.T) {
	testAnalyticFunction(t, Rank, rankTests)
}

var denseRankTests = []analyticFunctionTest{
	{
		Name: "DenseRank",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "dense_rank",
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
				OrderByClause: parser.OrderByClause{
					Items: []parser.Expression{
						parser.OrderItem{
							Value: parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
						},
					},
				},
			},
		},
		Result: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(2),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
					parser.NewInteger(2),
				}),
			},
		},
	},
}

func TestDenseRank(t *testing.T) {
	testAnalyticFunction(t, DenseRank, denseRankTests)
}

var firstValueTests = []analyticFunctionTest{
	{
		Name: "FirstValue",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(4),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(5),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "first_value",
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
			},
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Result: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
					parser.NewNull(),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(4),
					parser.NewNull(),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(5),
					parser.NewNull(),
				}),
			},
		},
	},
	{
		Name: "FirstValue Ignore Nulls",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(4),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(5),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "first_value",
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
			},
			IgnoreNulls: true,
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Result: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
					parser.NewInteger(4),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(4),
					parser.NewInteger(4),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(5),
					parser.NewInteger(4),
				}),
			},
		},
	},
	{
		Name: "FirstValue Argument Length Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(2),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "first_value",
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Error: "[L:- C:-] function first_value takes 1 argument",
	},
	{
		Name: "FirstValue Partition Value Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(2),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "first_value",
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
			},
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "notexist"}},
					},
				},
			},
		},
		Error: "[L:- C:-] field notexist does not exist",
	},
	{
		Name: "FirstValue Argument Value Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(2),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "first_value",
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "notexist"}},
			},
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Error: "[L:- C:-] field notexist does not exist",
	}}

func TestFirstValue(t *testing.T) {
	testAnalyticFunction(t, FirstValue, firstValueTests)
}

var lastValueTests = []analyticFunctionTest{
	{
		Name: "LastValue",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(3),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(4),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "last_value",
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
			},
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Result: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(3),
					parser.NewNull(),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(4),
					parser.NewNull(),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
					parser.NewNull(),
				}),
			},
		},
	},
	{
		Name: "LastValue Ignore Nulls",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(3),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(4),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "last_value",
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
			},
			IgnoreNulls: true,
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Result: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(3),
					parser.NewInteger(4),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(4),
					parser.NewInteger(4),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
					parser.NewInteger(4),
				}),
			},
		},
	},
}

func TestLastValue(t *testing.T) {
	testAnalyticFunction(t, LastValue, lastValueTests)
}

var analyzeAggregateValueTests = []analyticFunctionTest{
	{
		Name: "AnalyzeAggregateValue",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "count",
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
			},
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Result: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewInteger(2),
				}),
			},
		},
	},
	{
		Name: "AnalyzeAggregateValue With Distinct",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name:     "count",
			Distinct: parser.Token{Token: parser.DISTINCT, Literal: "distinct"},
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
			},
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Result: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewInteger(1),
				}),
			},
		},
	},
	{
		Name: "AnalyzeAggregateValue With Wildcard",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "count",
			Args: []parser.Expression{
				parser.AllColumns{},
			},
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Result: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewInteger(3),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
					parser.NewInteger(3),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewInteger(3),
				}),
			},
		},
	},
	{
		Name: "AnalyzeAggregateValue Argument Length Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "count",
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Error: "[L:- C:-] function count takes 1 argument",
	},
	{
		Name: "AnalyzeAggregateValue Argument Value Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "count",
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "notexist"}},
			},
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Error: "[L:- C:-] field notexist does not exist",
	},
	{
		Name: "AnalyzeAggregateValue Partition Value Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "count",
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
			},
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "notexist"}},
					},
				},
			},
		},
		Error: "[L:- C:-] field notexist does not exist",
	},
}

func TestAnalyzeAggregateValue(t *testing.T) {
	testAnalyticFunction(t, AnalyzeAggregateValue, analyzeAggregateValueTests)
}

var analyzeListAggTests = []analyticFunctionTest{
	{
		Name: "AnalyzeListAgg",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "listagg",
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
				parser.NewString(","),
			},
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Result: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
					parser.NewString("1,2"),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
					parser.NewString("1,2"),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewString("1,1"),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
					parser.NewString("1,1"),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewString("1,1"),
				}),
			},
		},
	},
	{
		Name: "AnalyzeListAgg With Default Separator",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "listagg",
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
			},
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Result: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
					parser.NewString("12"),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
					parser.NewString("12"),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewString("11"),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewNull(),
					parser.NewString("11"),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("b"),
					parser.NewInteger(1),
					parser.NewString("11"),
				}),
			},
		},
	},
	{
		Name: "AnalyzeListAgg Argument Length Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "listagg",
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Error: "[L:- C:-] function listagg takes 1 or 2 arguments",
	},
	{
		Name: "AnalyzeListAgg Partition Value Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "listagg",
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "column2"}},
				parser.NewString(","),
			},
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "notexist"}},
					},
				},
			},
		},
		Error: "[L:- C:-] field notexist does not exist",
	},
	{
		Name: "AnalyzeListAgg Argument Value Error",
		View: &View{
			Header: NewHeaderWithoutId("table1", []string{"column1", "column2"}),
			Records: []Record{
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(1),
				}),
				NewRecordWithoutId([]parser.Primary{
					parser.NewString("a"),
					parser.NewInteger(2),
				}),
			},
		},
		Function: parser.AnalyticFunction{
			Name: "listagg",
			Args: []parser.Expression{
				parser.FieldReference{Column: parser.Identifier{Literal: "notexist"}},
				parser.NewString(","),
			},
			AnalyticClause: parser.AnalyticClause{
				Partition: parser.Partition{
					Values: []parser.Expression{
						parser.FieldReference{Column: parser.Identifier{Literal: "column1"}},
					},
				},
			},
		},
		Error: "[L:- C:-] field notexist does not exist",
	},
}

func TestAnalyzeListAgg(t *testing.T) {
	testAnalyticFunction(t, AnalyzeListAgg, analyzeListAggTests)
}
