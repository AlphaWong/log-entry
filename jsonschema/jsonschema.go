package jsonschema

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

var ReferenceSchema gojsonschema.JSONLoader

func Init() {
	ReferenceSchema = gojsonschema.NewStringLoader(RequestSchema)
}

func IsValidRequest(message string) error {
	var subject = gojsonschema.NewStringLoader(message)
	var result, err = gojsonschema.Validate(ReferenceSchema, subject)
	if err != nil {
		return err
	}
	if !result.Valid() {
		var s strings.Builder
		for _, desc := range result.Errors() {
			s.WriteString(fmt.Sprintf("- %s\n", desc))
		}
		return errors.New(s.String())
	}
	return nil
}
