package main

import "regexp"

type replacor struct {
	expr *regexp.Regexp
	str  string
}

type replacors []replacor

var rBold = regexp.MustCompile(`-bold$`)
var rTextCenter = regexp.MustCompile(`^(text-center|t-center)$`)
var rJustify = regexp.MustCompile(`^(justify-|j-)`)

var replace = replacors{
	{expr: rBold, str: "-weight-bold"},
	{expr: rJustify, str: "text-center-"},
	{expr: rTextCenter, str: "text-align-center"},
}

func preReplace(str string) string {
	for _, r := range replace {
		str = r.expr.ReplaceAllString(str, r.str)
	}
	return str
}
