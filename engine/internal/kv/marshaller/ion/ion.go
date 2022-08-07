package ion

import (
	"github.com/amzn/ion-go/ion"
)

func Set[T any](input T) ([]byte, error) {
	bytes, err := ion.MarshalBinary(input)
	return bytes, err
}

func Get[T any](raw []byte) (T, error) {
	var resp T

	err := ion.Unmarshal(raw, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
