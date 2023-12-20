package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"reflect"
)

func StringifyArrayOfStructs(data interface{}) string {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Slice {
		return "Not a slice"
	}

	var str string
	for i := 0; i < val.Len(); i++ {
		structElement := val.Index(i)
		str += StringifyStruct(structElement.Interface()) + "\n"
	}

	return str
}

func StringifyStruct(data interface{}) string {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	if val.Kind() != reflect.Struct {
		return "Not a struct"
	}

	var str string
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		str += fmt.Sprintf("%s: %v\n", fieldType.Name, field.Interface())
	}

	return str
}

func GetHashOfData(str string) string {
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(str))
	hashBytes := sha1Hash.Sum(nil)
	sha1String := hex.EncodeToString(hashBytes)
	return sha1String
}
