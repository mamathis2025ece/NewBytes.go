package main
import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)
type Options struct {
	DecimalPlaces      int
	FixedDecimals      bool
	ThousandsSeparator string
	Unit               string
	UnitSeparator      string
}
var unitMap = map[string]float64{
	"b":  1,
	"kb": 1024,
	"mb": 1024 * 1024,
	"gb": 1024 * 1024 * 1024,
	"tb": math.Pow(1024, 4),
	"pb": math.Pow(1024, 5),
}
var parseRegex = regexp.MustCompile(`^((-|\+)?(\d+(\.\d+)?)) *(kb|mb|gb|tb|pb)$`)
func Format(value float64, opt *Options) string {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return ""
	}
	options := Options{
		DecimalPlaces: 2,
	}
	if opt != nil {
		options = *opt
		if options.DecimalPlaces == 0 {
			options.DecimalPlaces = 2
		}
	}
	unit := strings.ToUpper(options.Unit)
	mag := math.Abs(value)
	if unit == "" {
		switch {
		case mag >= unitMap["pb"]:
			unit = "PB"
		case mag >= unitMap["tb"]:
			unit = "TB"
		case mag >= unitMap["gb"]:
			unit = "GB"
		case mag >= unitMap["mb"]:
			unit = "MB"
		case mag >= unitMap["kb"]:
			unit = "KB"
		default:
			unit = "B"
		}
	}
	val := value / unitMap[strings.ToLower(unit)]
	str := fmt.Sprintf("%.*f", options.DecimalPlaces, val)
	if !options.FixedDecimals {
		str = strings.TrimRight(str, "0")
		str = strings.TrimRight(str, ".")
	}
	return str + options.UnitSeparator + unit
}
func Parse(input string) (int64, error) {
	input = strings.TrimSpace(strings.ToLower(input))
	match := parseRegex.FindStringSubmatch(input)
	var value float64
	unit := "b"
	if match == nil {
		v, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return 0, err
		}
		value = v
	} else {
		v, err := strconv.ParseFloat(match[1], 64)
		if err != nil {
			return 0, err
		}
		value = v
		unit = match[5]
	}
	return int64(math.Floor(value * unitMap[unit])), nil
}

