// Package request
package request

import (
	"encoding/json"
	"io"
	"net/http"
)

type StructValidator interface {
	Struct(s any) error
}

func DecodeAndValidateBody[T any](r *http.Request, validator StructValidator, res *T) error {
	if err := decodeBody(r.Body, res); err != nil {
		return err
	}

	err := validate(validator, res)
	if err != nil {
		return err
	}

	return nil
}

func validate[T any](validator StructValidator, payload *T) error {
	return validator.Struct(payload)
}

func decodeBody[T any](body io.ReadCloser, res *T) error {
	err := json.NewDecoder(body).Decode(res)
	if err != nil {
		return err
	}

	return nil
}
