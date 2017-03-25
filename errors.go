package goutility

const (
	CodeContext = "goutility"

	UnableToFindExternalIPAddressErrorCode ErrorCode = 1 << iota
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
