package test

import (
    "testing"
    "github.com/timurgulov/calc_go/internal/calculation"
)

func TestCalc(t *testing.T) {

    SuccessCases := []struct {
        Level          string
        Expression     string
        ExpectedResult float64
    }{
        {
            Level:          "first level",
            Expression:     "2+3",
            ExpectedResult: 5.0,
        },
        {
            Level:          "second level",
            Expression:     "(2+3)*4",
            ExpectedResult: 20.0,
        },
        {
            Level:          "third level",
            Expression:     "2*(3+4)/7",
            ExpectedResult: 2.0,
        },
        {
            Level:          "fourth level",
            Expression:     "2.5*4.2",
            ExpectedResult: 10.5,
        },
        {
            Level:          "fifth level",
            Expression:     "2+3*5-(2+1)/3",
            ExpectedResult: 16.0,
        },
    }

    for _, testcase := range SuccessCases {

        t.Run(testcase.Level, func(t *testing.T) {
            result, err := calculation.Calc(testcase.Expression)
            if err != nil {
                t.Fatalf("Expected result %f, but got error: %v", testcase.ExpectedResult, err)
            }
            if result != testcase.ExpectedResult {
                t.Fatalf("Expected result %f, but got %f", testcase.ExpectedResult, result)
            }
        })
    }

    BadCases := []struct {
        Level      string
        Expression string
    }{
        {
            Level:      "first level",
            Expression: "",
        },
        {
            Level:      "second level",
            Expression: "2+(3*5",
        },
        {
            Level:      "third level",
            Expression: "2+3&5",
        },
        {
            Level:      "fourth level",
            Expression: "5/0",
        },
        {
            Level:      "fifth level",
            Expression: "+",
        },
    }

    for _, testcase := range BadCases {
        t.Run(testcase.Level, func(t *testing.T) {
            result, err := calculation.Calc(testcase.Expression)
            if err == nil {
                t.Fatalf("Expected error for expression '%s', but got result %f", testcase.Expression, result)
            }
        })
    }
}