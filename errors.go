package goutility

import (
	"fmt"
	"reflect"
)

const (
	CodeContext = "goutility"

	UnableToFindExternalIPAddressErrorCode ErrorCode = 1 << iota

	MarshalErrorCode
	UnmarshalErrorCode
)

// region UnableToFindExternalIPAddressError
type UnableToFindExternalIPAddressError struct {
	ErrorType
}

func MakeUnableToFindExternalIPAddressError() UnableToFindExternalIPAddressError {
	return UnableToFindExternalIPAddressError{
		ErrorType: MakeErrorWithCode("Unable to find external ip address", UnableToFindExternalIPAddressErrorCode, CodeContext),
	}
}

// endregion

// region MashalError
type MashalError struct {
	ErrorType

	object interface{}
	error  error
}

func MakeMashalError(object interface{}, error error) MashalError {
	message := fmt.Sprintf("Error while mashalling an object of type '%s': %s", reflect.TypeOf(object).String(), error.Error())
	return MashalError{
		ErrorType: MakeErrorWithCode(message, MarshalErrorCode, CodeContext),
		object:    object,
		error:     error,
	}
}

// endregion

// region UnmarshalError
type UnmarshalError struct {
	ErrorType

	object interface{}
	error  error
}

func MakeUnmashalError(object interface{}, error error) UnmarshalError {
	message := fmt.Sprintf("Error while unmarshalling an object of type '%s' from data: %s", reflect.TypeOf(object).String(), error.Error())
	return UnmarshalError{
		ErrorType: MakeErrorWithCode(message, UnmarshalErrorCode, CodeContext),
		object:    object,
		error:     error,
	}
}

// endregion
