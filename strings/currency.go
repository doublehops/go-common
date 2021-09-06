package strings

import (
	"fmt"
	"strconv"
)

func AddSeparator(value float64, separator string) string {
	i := int(value)
	d := fmt.Sprintf("%.2g", value-float64(i))[2:]

	//val := separate(strconv.Itoa(i), ",")
	str := strconv.Itoa(i)

	var newValue string
	for {
		if len(str) > 3 {
			newValue = separator + str[len(str)-3:] + newValue
			str = str[:len(str)-3]
		} else {
			return str + newValue
		}
	}

	v := fmt.Sprintf("%s.%s", newValue, d)

	return v

	//v := fmt.Sprintf("%s.%s", val, d)
	//
	//return v
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
