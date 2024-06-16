package main

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	SEPERATOR                                = "."
	keyNillableMap        map[string]keyMeta = make(map[string]keyMeta)
	RETURN_VAL_TYPE_MACRO                    = "##VAL##"
)

func initF(obj any, variableName string) string {
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
	returnValueType := fmt.Sprintf("func(%s %s, i int) interface{}", variableName, variableDataType)
	mapString = strings.ReplaceAll(mapString, RETURN_VAL_TYPE_MACRO, returnValueType)

	// fmt.Println(mapString)
	return mapString
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
	// var lastObjDataType reflect.Kind
	for _, obj := range objs {
		// lastObjDataType = keyNillableMap[obj].dataType
		if keyNillableMap[obj].isCollection {
			obj = fmt.Sprintf("%s[i]", obj)
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

	function := `func(%s %s, i int) %s {
		%s
	}`
	return fmt.Sprintf("\"%s\": %s,\n", key, fmt.Sprintf(function, objs[0], variableNameDataType, "interface{}", ifCondition))
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
