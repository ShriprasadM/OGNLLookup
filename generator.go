package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	mapString := fmt.Sprintf("var %sLookup = map[string]%s{\n", variableName, RETURN_VAL_TYPE_MACRO)
	for i := 0; i < num; i++ {
		field := fields.Field(i)
		mapString = itr(field, variableName, mapString, variableDataType)
	}
	mapString += "}"
	// set return value type
	returnValueType := fmt.Sprintf("func(%s %s, indexInfo %s) (interface{}, error)", variableName, variableDataType, INDEX_INFO_STRUCT_NAME)
	mapString = strings.ReplaceAll(mapString, RETURN_VAL_TYPE_MACRO, returnValueType)

	// fmt.Println(mapString)
	indexInfoStruct := indexInfoStruct(indexInfoFieldsWithDataTypes...)
	output := indexInfoStruct + "\n" + mapString
	path := filepath.Dir(outputFileName)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	// use rightmost path as package name
	paths := strings.Split(path, "/")
	packageName := paths[len(paths)-1]
	output = fmt.Sprintf("package %s\n%s", packageName, output)
	// add lookup function
	output += generateLookupFunc(variableName, variableDataType)
	err = os.WriteFile(outputFileName, []byte(output), os.ModePerm)
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
	conditionStmt := ""
	// collectionCounterVariables := []string{}
	// var lastObjDataType reflect.Kind
	for index, obj := range objs {
		// lastObjDataType = keyNillableMap[obj].dataType
		counterVariable := ""
		if keyNillableMap[obj].isCollection {
			counterVariable = fmt.Sprintf("%sIndex", strings.Title(strings.ToLower(obj)))
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
			if counterVariable != "" {
				conditionStmt = orCondition("", fmt.Sprintf("indexInfo.%s >= len(%s)", counterVariable, ognl(prevStmt, objs[index])))
			}
			prevStmt = ognl(prevStmt, obj)
		}

		if !keyNillableMap[objs[index]].isNilable {
			continue
		}

		if ifCondition == "" {
			ifCondition = "if "
		} else {
			ifCondition += "|| "
		}
		// ifCondition += fmt.Sprintf("%s == nil ", conditionStmt)
		ifCondition += conditionStmt
	}
	if ifCondition == "" {
		ifCondition = "\n return " + prevStmt + ", nil"
	} else {

		errorObject := fmt.Sprintf(`fmt.Errorf("invalid index present in indexInfo object '%s'", indexInfo)`, "%+v")
		ifCondition += "{\n return nil, " + errorObject + " \n}\n return " + prevStmt + ", nil"
	}

	// %s
	// 1st - parent object ::  obj[0]
	// 2nd - parent object datatype  :: variableNameDataType
	// 3rd - counterVariables
	// 4th - return type :: interface{}
	// 5th - ifCondition :: either contains nil checks or simple return
	function := `func(%s %s %s) (%s, error) {
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
	tpe := reflect.TypeOf(i)
	nillable := isKindNillable(tpe.Kind())

	if sField, ok := i.(reflect.StructField); ok {
		nillable = isKindNillable(sField.Type.Kind())
	}
	return nillable
}

func isKindNillable(kind reflect.Kind) bool {
	switch kind {
	case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan, reflect.Interface:
		return true
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

/*
Example

	func LookUpBidResponse[R any](key string, bidResponse openrtb.BidResponse, indexInfo IndexInfo, returnType R) (R, error) {
		fn := bidResponseLookup[key]
		if fn == nil {
			return returnType, fmt.Errorf("invalid OGNL key [%s]", key)
		}
		result := fn(bidResponse, indexInfo)
		if rTypeValue, ok := result.(R); ok {
			return rTypeValue, nil
		}
		return returnType, fmt.Errorf("data type mismatch.Expected [%T] but found [%T]", returnType, result)
	}
*/
func generateLookupFunc(variableName, variableDataType string) string {
	tmpl := `// LookUpBidResponse obtains value of type R using key from %s. It uses indexIndo if required
	// In case of failure returns error
	func LookUpBidResponse[R any](key string, %s %s, indexInfo IndexInfo, returnType R) (R, error) {
	fn := bidResponseLookup[key]
	if fn == nil {
		return returnType, fmt.Errorf("invalid OGNL key [%s]", key)
	}
	result, err := fn(bidResponse, indexInfo)
	if err != nil {
		// Note: returntype here is not actual value
		return returnType, err
	}
	if rTypeValue, ok := result.(R); ok {
		return rTypeValue, nil
	}
	return returnType, fmt.Errorf("data type mismatch.Expected [%s] but found [%s]", returnType, result)
}`
	return fmt.Sprintf("\n"+tmpl, variableName, variableName, variableDataType, "%s", "%T", "%T")
}

func orCondition(prevStmt, newStmt string) string {
	if prevStmt != "" {
		return prevStmt + " || " + newStmt
	}
	return newStmt
}
