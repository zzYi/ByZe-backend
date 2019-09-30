package service

import (
	"ByZe/pkg/gredis"
	"math"
)

func ProcessFileDataService(fileid string) error {
	var x [1000]int
	for k := range x {
		x[k] = k - 500
	}

	var y [1000]float64
	for k, v := range x {
		y[k] = math.Pow(float64(v), 2)
	}

	var testData map[string]interface{}
	testData = make(map[string]interface{})
	testData["x"] = x
	testData["y"] = y

	err := gredis.Set(fileid, testData, 0)
	if err != nil {
		return err
	}

	return nil
}
