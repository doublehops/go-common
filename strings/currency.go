package strings

import (
	"fmt"
	"strconv"
)

func AddSeparator(value float64) string {
	i := int(value)
	d := fmt.Sprintf("%.2g", value-float64(i))[2:]

	val := separate(strconv.Itoa(i), ",")

	v := fmt.Sprintf("%s.%s", val, d)

	return v
}

func separate(str string, separator string) string {
	var newValue string
	for {
		if len(str) > 3 {
			newValue = separator + str[len(str)-3:] + newValue
			str = str[:len(str)-3]
		} else {
			return str + newValue
		}
	}

	return newValue
}
