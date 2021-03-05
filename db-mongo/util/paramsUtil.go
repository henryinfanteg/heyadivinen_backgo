package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Filters
const (
	filterContains   = "[contains]"
	filterBetween    = "[between]"
	filterStartswith = "[startswith]"
	filterEndswith   = "[endswith]"
	filterNE         = "[ne]"
	filterGT         = "[gt]"
	filterLT         = "[lt]"
	filterGTE        = "[gte]"
	filterLTE        = "[lte]"

	dateFormat = "2006-01-02T15:04:05-07"
)

// Data types
const (
	TypeString   = "STRING"
	TypeInteger  = "INTEGER"
	TypeBoolean  = "BOOLEAN"
	TypeDate     = "DATE"
	TypeObjectID = "OBJECT_ID"
)

// GetParametrosFiltro obtiene los parametros para filtro
func GetParametrosFiltro(queryParams map[string][]string, typeFields map[string]string) map[string]interface{} {
	params := make(map[string]interface{})

	// Recorremos los parametros
	for key, value := range queryParams {
		// fmt.Printf("key: %s -> value: %s\n", key, value)
		field, values := getParametros(key, value, typeFields)

		// Validamos que sea un campo de filtro valido
		if len(field) > 0 && len(typeFields[field]) > 0 {
			params[field] = values
		}
	}
	return params
}

func getParametros(paramKey string, paramValue []string, typeFields map[string]string) (string, interface{}) {
	values := strings.Split(paramValue[0], ",")

	if strings.Contains(paramKey, filterNE) {
		options := strings.Split(paramKey, filterNE)
		field := options[0]

		// Validamos si es vector
		if len(values) > 1 {
			if typeFields[field] == TypeString {
				return field, bson.M{"$nin": values}
			} else if typeFields[field] == TypeInteger {
				return field, bson.M{"$nin": convertSliceStringToInteger(values)}
			}
		} else {
			if typeFields[field] == TypeString {
				if values[0] != "null" {
					return field, bson.M{"$ne": values[0]}
				} else {
					return field, bson.M{"$ne": nil}
				}
			} else if typeFields[field] == TypeInteger {
				value, _ := strconv.Atoi(values[0])
				return field, bson.M{"$ne": value}
			} else if typeFields[field] == TypeBoolean {
				value, _ := strconv.ParseBool(values[0])
				return field, bson.M{"$ne": value}
			} else if typeFields[field] == TypeDate {
				if values[0] != "null" {
					value, _ := time.Parse(dateFormat, values[0])
					return field, bson.M{"$ne": value}
				} else {
					return field, bson.M{"$ne": nil}
				}
			}
		}
	} else if strings.Contains(paramKey, filterContains) {
		options := strings.Split(paramKey, filterContains)
		field := options[0]

		// Validamos si es vector
		if len(values) > 1 {
			arraybson := make([]bson.RegEx, 0)
			for _, obj := range values {
				arraybson = append(arraybson, bson.RegEx{Pattern: obj, Options: "i"})
			}
			return field, bson.M{"$in": arraybson}
		} else {
			return field, bson.RegEx{Pattern: values[0], Options: "i"}
		}
	} else if strings.Contains(paramKey, filterStartswith) {
		options := strings.Split(paramKey, filterStartswith)
		field := options[0]

		// Validamos si es vector
		if len(values) > 1 {
			arraybson := make([]bson.RegEx, 0)
			for _, obj := range values {
				arraybson = append(arraybson, bson.RegEx{Pattern: fmt.Sprintf("^%v", obj), Options: "i"})
			}
			return field, bson.M{"$in": arraybson}
		} else {
			return field, bson.RegEx{Pattern: fmt.Sprintf("^%v", values[0]), Options: "i"}
		}
	} else if strings.Contains(paramKey, filterEndswith) {
		options := strings.Split(paramKey, filterEndswith)
		field := options[0]

		// Validamos si es vector
		if len(values) > 1 {
			arraybson := make([]bson.RegEx, 0)
			for _, obj := range values {
				arraybson = append(arraybson, bson.RegEx{Pattern: fmt.Sprintf("%v$", obj), Options: "i"})
			}
			return field, bson.M{"$in": arraybson}
		} else {
			return field, bson.RegEx{Pattern: fmt.Sprintf("%v$", values[0]), Options: "i"}
		}
	} else if strings.Contains(paramKey, filterGT) {
		options := strings.Split(paramKey, filterGT)
		field := options[0]

		if typeFields[field] == TypeString {
			return field, bson.M{"$gt": values[0]}
		} else if typeFields[field] == TypeInteger {
			value, _ := strconv.Atoi(values[0])
			return field, bson.M{"$gt": value}
		}

	} else if strings.Contains(paramKey, filterLT) {
		options := strings.Split(paramKey, filterLT)
		field := options[0]

		if typeFields[field] == TypeString {
			return field, bson.M{"$lt": values[0]}
		} else if typeFields[field] == TypeInteger {
			value, _ := strconv.Atoi(values[0])
			return field, bson.M{"$lt": value}
		}

	} else if strings.Contains(paramKey, filterGTE) {
		options := strings.Split(paramKey, filterGTE)
		field := options[0]

		if typeFields[field] == TypeString {
			return field, bson.M{"$gte": values[0]}
		} else if typeFields[field] == TypeInteger {
			value, _ := strconv.Atoi(values[0])
			return field, bson.M{"$gte": value}
		}

	} else if strings.Contains(paramKey, filterLTE) {
		options := strings.Split(paramKey, filterLTE)
		field := options[0]

		if typeFields[field] == TypeString {
			return field, bson.M{"$lte": values[0]}
		} else if typeFields[field] == TypeInteger {
			value, _ := strconv.Atoi(values[0])
			return field, bson.M{"$lte": value}
		}

	} else if strings.Contains(paramKey, filterBetween) {
		options := strings.Split(paramKey, filterBetween)
		field := options[0]

		if len(values) > 1 {
			if typeFields[field] == TypeString {
				return field, bson.M{"$gte": values[0], "$lte": values[1]}
			} else if typeFields[field] == TypeInteger {
				sliceInteger := convertSliceStringToInteger(values)
				return field, bson.M{"$gte": sliceInteger[0], "$lte": sliceInteger[1]}
			} else if typeFields[field] == TypeDate {
				sliceTime := convertSliceStringToTime(values)
				if len(sliceTime) == 2 {
					return field, bson.M{"$gte": sliceTime[0], "$lte": sliceTime[1]}
				}
			}
		}

	} else {
		if typeFields[paramKey] == TypeString {
			if len(values) > 1 {
				return paramKey, bson.M{"$in": values}
			} else {
				if values[0] != "null" {
					return paramKey, values[0]
				} else {
					return paramKey, nil
				}
			}
		} else if typeFields[paramKey] == TypeInteger {
			if len(values) > 1 {
				return paramKey, bson.M{"$in": convertSliceStringToInteger(values)}
			} else {
				value, _ := strconv.Atoi(values[0])
				return paramKey, value
			}
		} else if typeFields[paramKey] == TypeBoolean {
			value, _ := strconv.ParseBool(values[0])
			return paramKey, value
		} else if typeFields[paramKey] == TypeDate {
			if values[0] != "null" {
				value, _ := time.Parse(dateFormat, values[0])
				return paramKey, value
			} else {
				return paramKey, nil
			}
		} else if typeFields[paramKey] == TypeObjectID {
			if bson.IsObjectIdHex(values[0]) {
				return paramKey, bson.ObjectIdHex(values[0])
			} else {
				return paramKey, nil
			}
		}
	}
	return "", nil
}

func convertSliceStringToInteger(sliceString []string) []int {

	var sliceInteger = []int{}

	for _, obj := range sliceString {
		value, err := strconv.Atoi(obj)
		if err == nil {
			sliceInteger = append(sliceInteger, value)
		}
	}
	return sliceInteger
}

func convertSliceStringToTime(sliceString []string) []time.Time {

	var sliceTime = []time.Time{}

	for _, obj := range sliceString {
		value, err := time.Parse(dateFormat, obj) // 2019-05-18T10:03:00-05,2019-05-18T15:59:59-05
		if err == nil {
			sliceTime = append(sliceTime, value)
		} else {
			fmt.Println("err:", err, " ---- ", obj)
		}
	}
	return sliceTime
}
