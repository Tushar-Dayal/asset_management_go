package utils

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

func ParseJSONBody(r *http.Request, dst interface{}) error {
	decoder := jsoniter.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(dst)
	if err != nil {
		return err
	}
	return nil
}

type Role string

const (
	adminRole          Role = "admin"
	employeeMangerRole Role = "employee_manager"
	assetManagerRole   Role = "asset_manager"
)
