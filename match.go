package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var classesComponents = map[string]string{
	"b":    "border",
	"c":    "color",
	"s":    "size",
	"r":    "radius",
	"m":    "margin",
	"p":    "padding",
	"w":    "width",
	"f":    "flex",
	"gr":   "grow",
	"sh":   "shrink",
	"bk":   "background",
	"d":    "display",
	"pos":  "position",
	"rel":  "relative",
	"abs":  "absolute",
	"full": "100%",
	"ov":   "overflow",
}

func parseError(s string) error {
	return errors.New("could not map:" + s)
}

func makeClass(prefix, class, attr string) string {
	return "." + class + prefix + "{" + attr + ";}"
}

func match(str string) (string, error) {
	var style string

	prefixSplit := strings.Split(str, ":")

	var prefix = ""
	if len(prefixSplit) > 1 {
		prefix = ":" + prefixSplit[0]
	}

	str = strings.Replace(str, "hover:", "", -1)
	splitClass := strings.Split(str, "-")

	var sep string
	var previousToken string
	for i, token := range splitClass {
		newtoken, ok := classesComponents[token]

		if ok {
			token = newtoken
		}

		// catches use of x, y, b, t
		// e.g.: p-t-2 or p-x-2
		if len(splitClass)-1 > i && i > 0 {
			newtoken, valid := makeXYTB(token, splitClass[i+1], previousToken)

			if valid {
				return makeClass(prefix, str, newtoken), nil
			}
		}

		// last should always be the attribute value
		if !isColor(token) && len(splitClass) > 1 && i == len(splitClass)-1 {

			// numbers are treated as rem
			// except 50 and 100 as %
			// and shrink and grow as int, e.g.: flex-grow:1
			if isNumber(token) {
				token = converNumber(token, splitClass[i-1])
			}

			style += ":" + token
			return makeClass(prefix, str, style), nil
		}

		previousToken = token

		if sep != "-" && i > 0 {
			sep = "-"
		}

		// special handling of colors
		color, err := makeColor(splitClass[i:])

		// it's a color
		if err == nil {
			style += color
			return makeClass(prefix, str, style), nil
		}

		// it's a "simple color"
		// e.g.: bk-white
		if errors.Is(err, errSimpleColor) {
			// text-white converts to font-white
			// we want color: white
			if token == "font" {
				style += color
				return makeClass(prefix, str, style), nil
			}
			// results in e.g.: background-color: white;
			style += token + "-" + color
			return makeClass(prefix, str, style), nil
		}

		// it's not a color
		if errors.Is(err, errNotColor) {
			style += sep + token
			continue
		}

		if err != nil {
			return "", err
		}

		return makeClass(prefix, str, style), nil
	}

	return makeClass(prefix, str, style), nil
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)

	return err == nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func converNumber(s, previous string) string {
	if previous == "grow" || previous == "shrink" {
		return s
	}

	if s == "50" || s == "100" {
		return s + "%"
	}

	num, _ := strconv.Atoi(s)

	// we divide by 4
	numf := roundFloat(float64(num)/4, 2)

	return fmt.Sprintf("%vrem", numf)
}

func makeXYTB(token, next, previous string) (string, bool) {
	if token == "x" {
		return previous + "-left:" + next + ";" + previous + "-right:" + next, true
	}

	if token == "y" {
		return previous + "-top:" + next + ";" + previous + "-bottom:" + next, true
	}

	if token == "t" {
		return previous + "-top:" + next + ";", true
	}

	if token == "b" {
		return previous + "-bottom:" + next + ";", true
	}

	return "", false
}
