package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var rPadding = regexp.MustCompile("^p(x|y|b|t)-|^p-")
var rMargin = regexp.MustCompile("^m(x|y|b|t)-|^m-")
var rWidth = regexp.MustCompile("^w-")
var rHeight = regexp.MustCompile("^h-")
var rRounded = regexp.MustCompile("^r-")
var rBorderWidth = regexp.MustCompile("^bw-")
var rBorderColor = regexp.MustCompile("^bc-")
var rBorderStyle = regexp.MustCompile("^bs-")

func parseError(s string) error {
	return errors.New("could not map:" + s)
}

func match(str string) (string, error) {
	var class string

	if rPadding.MatchString(str) {
		return makeAll(str, "padding", "p")
	}

	if rMargin.MatchString(str) {
		return makeAll(str, "margin", "m")
	}

	if rWidth.MatchString(str) {
		return makeSimple(str, "width", "w")
	}

	if rHeight.MatchString(str) {
		return makeSimple(str, "height", "h")
	}

	if rRounded.MatchString(str) {
		return makeSimple(str, "border-radius", "r")
	}

	if rBorderWidth.MatchString(str) {
		return makeSimple(str, "border-width", "bw")
	}

	if rBorderStyle.MatchString(str) {
		return makeSimple(str, "border-style", "bs")
	}

	return class, parseError(str)
}

func makeXYTB(s, t, m string) (string, error) {
	split := strings.Split(s, "-")

	if len(split) != 2 {
		return "", parseError(s)
	}

	var end string
	if isNumber(split[1]) {
		end = "rem"
	}

	if split[0] == m+`x` {
		return `.` + s + `{` + t + `-left:` + split[1] + end + ";" + t + `-right:` + split[1] + end + `;}`, nil
	}

	if split[0] == m+`y` {
		return `.` + s + `{` + t + `-top:` + split[1] + end + ";" + t + `-bottom:` + split[1] + end + `;}`, nil
	}

	if split[0] == m+`t` {
		return `.` + s + `{` + t + `-top:` + split[1] + end + `;}`, nil
	}

	if split[0] == m+`b` {
		return `.` + s + `{` + t + `-bottom:` + split[1] + end + `;}`, nil
	}

	return "", parseError(s)
}

func makeSimple(s, t, m string) (string, error) {
	split := strings.Split(s, "-")

	if len(split) != 2 {
		return "", parseError(s)
	}

	// for heights and widths
	if split[1] == "100" {
		split[1] = "100%"
	}

	// for heights and widths
	if split[1] == "50" {
		split[1] = "50%"
	}

	var end string
	if isNumber(split[1]) {
		end = "rem"
	}

	if split[0] == m {
		return `.` + s + `{` + t + `:` + split[1] + end + `;}`, nil
	}

	return "", parseError(s)
}

func makeAll(s, t, m string) (string, error) {
	cl, err := makeSimple(s, t, m)

	if err == nil {
		return cl, err
	}

	return makeXYTB(s, t, m)
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)

	return err == nil
}
