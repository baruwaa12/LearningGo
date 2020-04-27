package main

import (
	"gopl.io/ch7/eval"
	"time"
	"fmt"
)


type mydate interface {
	now() mydate
	seconds() int
	parse(dateString string) (bool, mydate)
	addSeconds(currentTime mydate, seconds int)

}

func parseAndCheck(s string) (eval.Expr, error) {
	if s == ""
		return nil,
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}