package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//func isIntegral(val float64) bool {
//	return val == float64(int(val))
//}
//
//func isFloatint(floatValue float64) bool {
//	return math.Mod(floatValue, 1.0) == 0
//}
//
//func isFloatInt(value string) interface{} {
//	val := strings.Split(value, ".")
//	fmt.Println("Length of val is ", len(val))
//	if len(val) > 1 {
//		fmt.Println("Value in float true and value is", value)
//		floatVal, err := strconv.ParseFloat(value, 64)
//		fmt.Println("Float val is ", floatVal, " and err is ", err)
//		return floatVal
//	}
//	intVal, _ := strconv.Atoi(value)
//	return intVal
//}
//
//func isNumeric1(value string) bool {
//	for _, c := range value {
//		if c < '0' || c > '9' {
//			return false
//		}
//	}
//	return true
//}

func isFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

func isNumeric(value string) (int, error) {
	return strconv.Atoi(value)
}

func isBoolean(value string) (bool, error) {
	return strconv.ParseBool(value)
}

func covertIntoDataType(listToConvert []string) interface{} {
	var arr1 []interface{}
	for _, value := range listToConvert {
		num, err := isNumeric(value)
		if err == nil {
			arr1 = append(arr1, num)
			continue
		}
		floatVal, err := isFloat(value)
		if err == nil {
			arr1 = append(arr1, floatVal)
			continue
		}
		boolVal, err := isBoolean(value)
		if err == nil {
			arr1 = append(arr1, boolVal)
			continue
		}
		arr1 = append(arr1, value)
	}
	return arr1
}

func main() {
	readers := []string{"4.0", "Alice", "False", "Charlie"}
	readers = append(readers, "42")
	readers = append(readers, "0.3")
	readers = append(readers, "True")
	fmt.Println("List values are: ", readers)
	arr1 := covertIntoDataType(readers)
	for _, value := range arr1.([]interface{}) {
		fmt.Println("Value: ", value, ", Data type: ", reflect.TypeOf(value))
	}
}
