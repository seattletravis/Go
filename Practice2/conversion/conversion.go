package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	for stringIndex, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			fmt.Println("Could not parse data to float64.")
			fmt.Println(err)
			file.Close()
			return
		}

		prices[lineIndex] = floatPrice
	}
}