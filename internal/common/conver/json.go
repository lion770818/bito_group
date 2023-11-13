package conver

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func Json2Map(data []byte) map[string]interface{} {
	var jsonMap map[string]interface{}
	err := json.Unmarshal(data, &jsonMap)
	if err != nil {
		log.Println(err)
		return nil
	}
	return jsonMap
}

func Map2Json(data map[string]interface{}) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return jsonData
}

func Interface2Int(data interface{}) int {
	return int(data.(float64))
}

func Interface2Float64(data interface{}) float64 {
	return data.(float64)
}

func InterfaceD2Int(data interface{}) int {
	return data.(int)
}

func Interface2Int64(data interface{}) int64 {
	return int64(data.(float64))
}

func Interface2Int64D(data interface{}) int64 {
	return data.(int64)
}

func Interface2Str(data interface{}) string {
	return data.(string)
}

func Interface2NullStr(data interface{}) string {
	if data == nil {
		return ""
	}
	return data.(string)
}

func Interface2InterfaceArray(data interface{}) []interface{} {
	return data.([]interface{})
}

func Interface2IntArray(data interface{}) []int {
	arr := Interface2InterfaceArray(data)

	arrList := make([]int, 0, len(arr))
	for i := range arr {
		arrList = append(arrList, Interface2Int(arr[i]))
	}
	return arrList
}

func InterfaceMap(data interface{}) map[string]interface{} {
	return data.(map[string]interface{})
}

func Interface2IntWithDefault(data interface{}) int {
	if data == nil {
		return 0
	}
	return data.(int)
}

func Interface2IntWithDefaultValue(data interface{}, defaultValue int) int {
	if data == nil {
		return defaultValue
	}
	return data.(int)
}

func Str2Int(data string) int {
	v, e := strconv.Atoi(data)
	if e != nil {
		return 0
	}
	return v
}

func Str2Int32(data string) int32 {
	v, e := strconv.Atoi(data)
	if e != nil {
		return 0
	}
	return int32(v)
}

func Str2Int64(data string) int64 {
	v, e := strconv.ParseInt(data, 10, 0)
	if e != nil {
		return 0
	}
	return v
}

func InterStr2Int32(data interface{}) int32 {
	return Str2Int32(Interface2Str(data))
}

func Inter2Str(inter interface{}) string {
	if inter == nil {
		return ""
	}

	switch inter.(type) {
	case string:
		return fmt.Sprintf("%s", inter.(string))
	case int:
		return fmt.Sprintf("%d", inter.(int))
	case float64:
		return fmt.Sprintf("%d", int64(inter.(float64)))
	case int64:
		return fmt.Sprintf("%d", inter.(int64))
	}
	return ""
}

func Inter2SNum(inter interface{}) int64 {
	valStr := Inter2Str(inter)
	if valStr == "" {
		return -1
	}
	val, _ := strconv.ParseInt(valStr, 10, 0)
	return val
}
