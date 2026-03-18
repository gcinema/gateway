// Package httpreq
package httpreq

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gcinema/gateway/pkg/errconst"
)

type StructValidator interface {
	Struct(s any) error
}

func DecodeAndValidateBody[T any](r *http.Request, validator StructValidator, res *T) error {
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return fmt.Errorf("%w: %v", errconst.ErrInvalidArgument, err)
	}

	if err := validator.Struct(res); err != nil {
		return fmt.Errorf("%w: %v", errconst.ErrInvalidArgument, err)
	}

	return nil
}
