package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func isFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

func isNumeric(value string) (int, error) {
	return strconv.Atoi(value)
}

func isBoolean(value string) (bool, error) {
	return strconv.ParseBool(value)
}
func isJsonObj(value string)(bool, map[string]interface{}){
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(value), &jsonMap)
	if jsonMap != nil{
		return true, jsonMap
	}
	return false, jsonMap
}


func isJsonArray(value string)(bool, [] interface{}){
	var jsonArr [] interface{}
	json.Unmarshal([]byte(value), &jsonArr)
	if jsonArr != nil{
		return true, jsonArr
	}
	return false, jsonArr
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
		var jsonMap map[string]interface{}
		checkJsonObj, jsonMap := isJsonObj(value)
		if checkJsonObj {
			arr1 = append(arr1, jsonMap)
			continue
		}
		checkJsonArr, jsonArr := isJsonArray(value)
		if checkJsonArr{
			arr1 = append(arr1, jsonArr)
			continue
		}
		arr1 = append(arr1, value)
	}
	return arr1
}

func main() {
	readers := []string{"4.0", "Alice", "False", "Charlie", "42", "0.3", "True", "{\"data\": \"key\"}",
		"[{\"data\":\"key\"}]", "[\"Hello world\", 10, false]"}
	fmt.Println("List values are: ", readers)
	arr1 := covertIntoDataType(readers)
	for _, value := range arr1.([]interface{}) {
		fmt.Println("Value: ", value, ", Data type: ", reflect.TypeOf(value))
	}
}
