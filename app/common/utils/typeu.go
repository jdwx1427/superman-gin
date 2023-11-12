package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func InterfaceToInt64(value interface{}) int64 {
	result := int64(0)
	if value == nil {
		return result
	}
	switch value.(type) {
	case int:
		i, _ := value.(int)
		result = int64(i)
	case uint:
		i, _ := value.(uint)
		result = int64(i)
	case int16:
		i, _ := value.(int16)
		result = int64(i)
	case uint16:
		i, _ := value.(uint16)
		result = int64(i)
	case int32:
		i, _ := value.(int32)
		result = int64(i)
	case uint32:
		i, _ := value.(uint32)
		result = int64(i)
	case int64:
		result, _ = value.(int64)
	case uint64:
		i, _ := value.(uint64)
		result = int64(i)
	case json.Number:
		i, _ := value.(json.Number).Int64()
		result = i
	case string:
		res, _ := value.(string)
		result = StringToInt64(res)
	}

	return result
}

func InterfaceToString(value interface{}) string {
	result := ""
	switch value.(type) {
	case string:
		result, _ = value.(string)
	}

	return result
}

func InterfaceToFloat64(value interface{}) float64 {
	result := float64(0)
	switch value.(type) {
	case float64:
		result, _ = value.(float64)
	}

	return result
}

func StringToInt64(str string) int64 {
	if str == "" {
		return int64(0)
	}

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return int64(0)
		}
		i = int64(math.Ceil(f))
	}

	return i
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int64ArrayToString(iList []int64) string {
	iStrArr := make([]string, 0)
	for _, v := range iList {
		iStrArr = append(iStrArr, Int64ToString(v))
	}
	iStr := strings.Join(iStrArr, ",")

	return iStr
}

func StringToInt64Array(s string) []int64 {
	sList := strings.Split(s, ",")

	iList := make([]int64, 0)
	for _, v := range sList {
		iList = append(iList, StringToInt64(v))
	}

	return iList
}

func StringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return float64(0)
	}

	return f
}

func Float64ToString(f float64, precision int) string {
	return strconv.FormatFloat(f, 'f', precision, 64)
}

func Float64ToStringFast(f float64) string {
	return Float64ToString(f, -1)
}

func Float64ToDecimal(f float64) float64 {
	v, _ := strconv.ParseFloat(Float64ToString(f, 2), 64)

	return v
}

func StringToDecimal(str string) float64 {
	return Float64ToDecimal(StringToFloat64(str))
}

func IsFloat64(value interface{}) bool {
	switch value.(type) {
	case float64:
		return true
	}

	return false
}

func Int64ToFloat64(i int64) float64 {
	return float64(i)
}

func Float64Retain3(f float64) float64 {
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", f), 64)

	return v
}

func BoolToInt64(b bool) int64 {
	i := int64(0)
	if b {
		i = 1
	}

	return i
}
