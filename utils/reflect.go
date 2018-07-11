package utils

import (
	"reflect"
	"errors"
)

func setField(obj interface{}, fieldName string, value interface{}) error {
	refVal := reflect.ValueOf(obj)
	if refVal.Kind() != reflect.Ptr {
		return errors.New("object should be a pointer")
	}

	if refVal.IsNil() {
		return errors.New("object should not be nil pointer")
	}

	pData := refVal.Elem()
	if !pData.IsValid() {
		return errors.New("object should not be nil pointer")
	}

	if pData.Kind() != reflect.Struct {
		return errors.New("object should be a struct")
	}

	fieldVal := pData.FieldByName(fieldName)

	if !fieldVal.IsValid() {
		return errors.New("field is invalid")
	}

	if !fieldVal.CanSet() {
		return errors.New("field cannot set")
	}


	val := reflect.ValueOf(value)

	if val.Kind() != fieldVal.Kind() {
		return errors.New("value does not match field's type")
	}

	fieldVal.Set(val)
	return nil
}

func SetStructByJSON(obj interface{}, json map[string]interface{}) error {
	for k, v := range json {
		if err := setField(obj, k, v); err != nil {
			return err
		}
	}

	return nil
}