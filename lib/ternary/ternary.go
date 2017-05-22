package ternary

import (
	"errors"
	"fmt"
	"strings"
)

type Value int

const (
	FALSE Value = iota - 1
	UNKNOWN
	TRUE
)

var literals = map[Value]string{
	FALSE:   "FALSE",
	UNKNOWN: "UNKNOWN",
	TRUE:    "TRUE",
}

func (v Value) String() string {
	return literals[v]
}

func (v Value) EqualTo(v2 Value) Value {
	if v == v2 {
		return TRUE
	}
	return FALSE
}

func (v Value) BoolValue() bool {
	if v != TRUE {
		return false
	}
	return true
}

func Parse(s string) (Value, error) {
	switch strings.ToUpper(s) {
	case "FALSE", "-1":
		return FALSE, nil
	case "TRUE", "1":
		return TRUE, nil
	case "UNKNOWN", "0":
		return UNKNOWN, nil
	}
	return -1, errors.New(fmt.Sprintf("parsing %q: invalid syntax", s))
}

func ParseBool(b bool) Value {
	if b {
		return TRUE
	}
	return FALSE
}

func Not(v Value) Value {
	switch v {
	case FALSE:
		return TRUE
	case TRUE:
		return FALSE
	}
	return UNKNOWN
}

func And(v1 Value, v2 Value) Value {
	switch {
	case v1 == FALSE || v2 == FALSE:
		return FALSE
	case v1 == UNKNOWN || v2 == UNKNOWN:
		return UNKNOWN
	}
	return TRUE
}

func Or(v1 Value, v2 Value) Value {
	switch {
	case v1 == TRUE || v2 == TRUE:
		return TRUE
	case v1 == UNKNOWN || v2 == UNKNOWN:
		return UNKNOWN
	}
	return FALSE
}
