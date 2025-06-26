package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _, str := range strings {
		value, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}

		floats = append(floats, value)
	}

	return floats, nil
}
