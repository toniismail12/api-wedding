package services

import (
	"errors"
	"reflect"
)

func ValidationInputString(Input, Name string) error {

	if Input == "" || Input == " " {

		return errors.New(Name + " Wajib Diisi, Type String")
	}

	return nil
}

func ValidationInputInt(Input uint, Name string) error {
	if reflect.ValueOf(Input).Kind() != reflect.Uint {
		return errors.New(Name + " Type Input Wajib Angka")
	}

	return nil
}
