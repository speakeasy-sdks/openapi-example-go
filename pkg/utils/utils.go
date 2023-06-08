// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"regexp"
	"strings"
	"time"

	"openapi-example/pkg/types"
)

const (
	queryParamTagKey  = "queryParam"
	headerParamTagKey = "header"
	pathParamTagKey   = "pathParam"
)

var (
	paramRegex                       = regexp.MustCompile(`({.*?})`)
	SerializationMethodToContentType = map[string]string{
		"json":      "application/json",
		"form":      "application/x-www-form-urlencoded",
		"multipart": "multipart/form-data",
		"raw":       "application/octet-stream",
		"string":    "text/plain",
	}
)

func UnmarshalJsonFromResponseBody(body io.Reader, out interface{}) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	if err := json.Unmarshal(data, &out); err != nil {
		return fmt.Errorf("error unmarshalling json response body: %w", err)
	}

	return nil
}

func ReplaceParameters(stringWithParams string, params map[string]string) string {
	if len(params) == 0 {
		return stringWithParams
	}

	return paramRegex.ReplaceAllStringFunc(stringWithParams, func(match string) string {
		match = match[1 : len(match)-1]
		return params[match]
	})
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func parseStructTag(tagKey string, field reflect.StructField) map[string]string {
	tag := field.Tag.Get(tagKey)
	if tag == "" {
		return nil
	}

	values := map[string]string{}

	options := strings.Split(tag, ",")
	for _, optionConf := range options {
		parts := strings.Split(optionConf, "=")

		switch len(parts) {
		case 1:
			// flag option
			parts = append(parts, "true")
		case 2:
			// key=value option
			break
		default:
			// invalid option
			continue
		}

		values[parts[0]] = parts[1]
	}

	return values
}

func parseParamTag(tagKey string, field reflect.StructField, defaultStyle string, defaultExplode bool) *paramTag {
	// example `{tagKey}:"style=simple,explode=false,name=apiID"`
	values := parseStructTag(tagKey, field)
	if values == nil {
		return nil
	}

	tag := &paramTag{
		Style:     defaultStyle,
		Explode:   defaultExplode,
		ParamName: strings.ToLower(field.Name),
	}

	for k, v := range values {
		switch k {
		case "style":
			tag.Style = v
		case "explode":
			tag.Explode = v == "true"
		case "name":
			tag.ParamName = v
		case "serialization":
			tag.Serialization = v
		}
	}

	return tag
}

func valToString(val interface{}) string {
	switch v := val.(type) {
	case time.Time:
		return v.Format(time.RFC3339Nano)
	case types.BigInt:
		return v.String()
	case big.Int:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func populateFromGlobals(fieldType reflect.StructField, valType reflect.Value, paramType string, globals map[string]map[string]map[string]interface{}) reflect.Value {
	if globals != nil && fieldType.Type.Kind() == reflect.Ptr {
		parameters, ok := globals["parameters"]
		if ok {
			paramsOfType, ok := parameters[paramType]
			if ok {
				globalVal, ok := paramsOfType[fieldType.Name]
				if ok {
					if reflect.TypeOf(globalVal).Kind() == fieldType.Type.Elem().Kind() && valType.IsNil() {
						valType = reflect.ValueOf(&globalVal)
					}
				}
			}
		}
	}

	return valType
}
