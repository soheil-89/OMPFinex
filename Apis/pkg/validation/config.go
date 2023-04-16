package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func JsonErrorResponse(val validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range val {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}
	return errs
}
