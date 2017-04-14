package goutility

import "encoding/json"

func MarshalToJSON(v interface{}) ([]byte, ErrorTypeInterface) {
	rawData, error := json.Marshal(v)
	if error != nil {
		return nil, MakeMashalError(v, error)
	}

	return rawData, nil
}

func MarshalIndentToJSON(v interface{}, prefix string, indent string) ([]byte, ErrorTypeInterface) {
	rawData, error := json.MarshalIndent(v, prefix, indent)
	if error != nil {
		return nil, MakeMashalError(v, error)
	}

	return rawData, nil
}

func UnmarshalFromJSON(data []byte, v interface{}) ErrorTypeInterface {
	error := json.Unmarshal(data, v)
	if error != nil {
		return MakeUnmashalError(v, error)
	}

	return nil
}
