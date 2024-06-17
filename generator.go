package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

var (
	SEPERATOR                                       = "."
	keyNillableMap               map[string]keyMeta = make(map[string]keyMeta)
	RETURN_VAL_TYPE_MACRO                           = "##VAL##"
	INDEX_INFO_STRUCT_NAME                          = "IndexInfo"
	indexInfoFieldsWithDataTypes                    = make([]string, 0)
)

func initF(obj any, variableName, outputFileName string) string {
	fields := reflect.TypeOf(obj)
	variableDataType := fmt.Sprintf("%T", obj)
	setKeyMeta(obj, variableName)
	// values := reflect.ValueOf(obj)
	num := fields.NumField()
	mapString := fmt.Sprintf("m := map[string]%s{\n", RETURN_VAL_TYPE_MACRO)
	for i := 0; i < num; i++ {
		field := fields.Field(i)
		mapString = itr(field, variableName, mapString, variableDataType)
	}
	mapString += "}"
	// set return value type
	returnValueType := fmt.Sprintf("func(%s %s, indexInfo %s) interface{}", variableName, variableDataType, INDEX_INFO_STRUCT_NAME)
	mapString = strings.ReplaceAll(mapString, RETURN_VAL_TYPE_MACRO, returnValueType)

	// fmt.Println(mapString)
	indexInfoStruct := indexInfoStruct(indexInfoFieldsWithDataTypes...)
	output := indexInfoStruct + "\n" + mapString
	err := os.WriteFile(outputFileName, []byte(output), os.ModeDevice)
	if err != nil {
		fmt.Println(err.Error())
	}
	return output
}

func itr(field reflect.StructField, prefix string, mapString string, variableNameDataType string) string {

	// Get the type of the variable using reflection
	// rt := reflect.ValueOf(field)

	// Check if the type is an array
	switch field.Type.Kind() {
	case reflect.Slice:
		// mapString += mapKV(ognl(strings.ToLower(prefix), strings.ToLower(field.Name)), ognl(prefix, field.Name))
		fields := reflect.SliceOf(field.Type.Elem())
		setKeyMeta(field, field.Name)
		structFields := fields.Elem()
		if structFields.Kind() == reflect.Struct {
			num := structFields.NumField()
			prefix = ognl(prefix, field.Name)
			for i := 0; i < num; i++ {
				sfield := structFields.Field(i)
				mapString = itr(sfield, prefix, mapString, variableNameDataType)
			}
		}
		// fmt.Println(field, "is a slice with element type", field.Type.Elem())
	// case reflect.Array:
	// mapString += mapKV(strings.ToLower(field.Name), ognl(prefix, field.Name))
	// fmt.Println(field, "is an array with element type", field.Type.Elem())
	default:
		// m[strings.ToLower(field.Name)] = ognl(prefix, field.Name)
		setKeyMeta(field, field.Name)
		mapString += mapKV(ognlKey(prefix, field.Name), ognl(prefix, field.Name), variableNameDataType)
		// fmt.Println(field, "is something else entirely")
	}
	return mapString
}

func mapKV(key, val string, variableNameDataType string) string {
	objs := strings.Split(val, SEPERATOR)
	ifCondition := ""
	prevStmt := ""
	// collectionCounterVariables := []string{}
	// var lastObjDataType reflect.Kind
	for _, obj := range objs {
		// lastObjDataType = keyNillableMap[obj].dataType
		if keyNillableMap[obj].isCollection {
			counterVariable := fmt.Sprintf("%sIndex", strings.ToLower(obj))
			obj = fmt.Sprintf("%s[indexInfo.%s]", obj, counterVariable)
			// collectionCounterVariables = append(collectionCounterVariables, counterVariable)
			fieldPresent := false
			for _, field := range indexInfoFieldsWithDataTypes {
				if field == counterVariable {
					fieldPresent = true
					break
				}
			}
			if !fieldPresent {
				indexInfoFieldsWithDataTypes = append(indexInfoFieldsWithDataTypes, counterVariable)
			}
		}

		if prevStmt == "" {
			prevStmt = obj
		} else {
			prevStmt = ognl(prevStmt, obj)
		}

		if !keyNillableMap[obj].isNilable {
			continue
		}

		if ifCondition == "" {
			ifCondition = "if "
		} else {
			ifCondition += "|| "
		}
		ifCondition += fmt.Sprintf("%s == nil ", prevStmt)
	}
	if ifCondition == "" {
		ifCondition = "\n return " + prevStmt
	} else {
		ifCondition += "{\n return nil \n}\n return " + prevStmt
	}

	// %s
	// 1st - parent object ::  obj[0]
	// 2nd - parent object datatype  :: variableNameDataType
	// 3rd - counterVariables
	// 4th - return type :: interface{}
	// 5th - ifCondition :: either contains nil checks or simple return
	function := `func(%s %s %s) %s {
		%s
	}`
	// counters := ""
	// if len(collectionCounterVariables) > 0 {
	// 	// counters = strings.Join(collectionCounterVariables, " int,")
	// 	// counters += " int"
	// 	// counters = "," + counters
	// 	counters = fmt.Sprintf("indexInfo " + INDEX_INFO_STRUCT_NAME)
	// }
	valFun := fmt.Sprintf(function, objs[0], variableNameDataType, ",indexInfo "+INDEX_INFO_STRUCT_NAME, "interface{}", ifCondition)
	return fmt.Sprintf("\"%s\": %s,\n", key, valFun)
	// following returns actual datatype
	// return fmt.Sprintf("\"%s\": %s,\n", key, fmt.Sprintf(function, objs[0], variableNameDataType, lastObjDataType, ifCondition))
}
func ognlKey(prefix, variable string) string {
	return fmt.Sprintf("%s.%s", strings.ToLower(prefix), strings.ToLower(variable))
}
func ognl(prefix, variable string) string {
	return fmt.Sprintf("%s.%s", prefix, variable)
}

func isNillable(i any) bool {
	value := reflect.ValueOf(i)
	switch value.Kind() {
	case reflect.Pointer:
	case reflect.Map:
	case reflect.Slice:
	case reflect.Func:
	case reflect.Chan:
	case reflect.Interface:
		return true

	default:
		return false

	}
	return false
}

type keyMeta struct {
	isNilable    bool
	isCollection bool
	dataType     reflect.Kind
}

func setKeyMeta(key any, keyName string) {

	keyMeta := keyMeta{
		isNilable:    isNillable(key),
		isCollection: isCollection(key),
		// dataType:     reflect.TypeOf(key).Kind(),
	}

	if sField, ok := key.(reflect.StructField); ok {
		keyMeta.dataType = sField.Type.Kind()
	}
	keyNillableMap[keyName] = keyMeta
}

func isCollection(obj any) bool {
	oType := reflect.TypeOf(obj)
	if sObj, ok := obj.(reflect.StructField); ok {
		oType = sObj.Type
	}
	// fmt.Println(oType.Kind().String())
	// fields := reflect.SliceOf(oType.Elem())
	switch oType.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func indexInfoStruct(indices ...string) string {
	indexInfo := `type ` + INDEX_INFO_STRUCT_NAME + ` struct {
		%s
	}
	`
	fields := ""
	for _, field := range indices {
		fields += fmt.Sprintf("	%s int\n", field)
	}
	return fmt.Sprintf(indexInfo, fields)
}
