package utils

import (
	"strconv"
)

// AtoSecond convert string format to second. It will error if the format not meet specification.
// eg: 11s -> 11 second,
// 	12m -> 12 minute,
// 	13h- > 13 hour,
// 	14d -> 14 day,
// 	15w -> 15 week,
// 	16M -> 16 month,
// 	17y -> 17 year.
func AtoSecond(v string) (t int, err error) {
	if v == "" {
		return
	}

	length := len(v)
	val, unit := v[:length-1], v[length-1:]

	if t, err = strconv.Atoi(val); err != nil {
		return 0, err
	}

	switch unit {
	case "s":
	case "m":
		t = t * 60
	case "h":
		t = t * 60 * 60
	case "d":
		t = t * 60 * 60 * 24
	case "w":
		t = t * 60 * 60 * 24 * 14
	case "M":
		t = t * 60 * 60 * 24 * 30
	case "y":
		t = t * 60 * 60 * 24 * 365 // dot not check overfilow
	}

	return
}
