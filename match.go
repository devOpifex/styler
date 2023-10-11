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
	"t":    "font",
	"a":    "align",
	"j":    "justify",
	"i":    "items",
	"bk":   "background",
	"d":    "display",
	"pos":  "position",
	"rel":  "relative",
	"abs":  "absolute",
	"full": "100%",
	"ov":   "overflow",
	"sh":   "shadow",
}

func parseError(s string) error {
	return errors.New("could not map class:" + s)
}

func makeClass(prefix, class, attr string) (string, error) {
	if attr == "" {
		return "", parseError(class)
	}

	if !strings.Contains(attr, ":") {
		return "", parseError(class)
	}
	return "." + class + prefix + "{" + attr + ";}", nil
}

func match(str string) (string, error) {
	var style string

	prefixSplit := strings.Split(str, ":")

	var prefix = ""
	if len(prefixSplit) > 1 {
		prefix = ":" + prefixSplit[0]
	}

	str = strings.Replace(str, "hover:", "", -1)
	str = strings.Replace(str, "focus:", "", -1)

	str = preReplace(str)

	// split class into tokens
	splitClass := strings.Split(str, "-")

	var sep string
	var previousToken string
	for i, token := range splitClass {
		newtoken, ok := classesComponents[token]

		if ok {
			token = newtoken
		}

		if len(splitClass) > 1 && i == 0 {
			// special case for shadow
			// e.g.: sh-sm
			shadow, ok := makeShadow(token, splitClass[i+1])

			if ok {
				return makeClass(prefix, str, shadow)
			}
		}

		if len(splitClass)-1 > i && i > 0 {
			// catches use of x, y, b, t
			// e.g.: p-t-2 or p-x-2
			newtoken, valid := makeXYTB(token, splitClass[i+1], previousToken)

			if valid {
				return makeClass(prefix, str, newtoken)
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
			return makeClass(prefix, str, style)
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
			return makeClass(prefix, str, style)
		}

		// it's a "simple color"
		// e.g.: bk-white
		if errors.Is(err, errSimpleColor) {
			// text-white converts to font-white
			// we want color: white
			if token == "font" {
				style += color
				return makeClass(prefix, str, style)
			}
			// results in e.g.: background-color: white;
			style += token + "-" + color
			return makeClass(prefix, str, style)
		}

		// it's not a color
		if errors.Is(err, errNotColor) {
			style += sep + token
			continue
		}

		if err != nil {
			return "", err
		}

		return makeClass(prefix, str, style)
	}

	return makeClass(prefix, str, style)
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
