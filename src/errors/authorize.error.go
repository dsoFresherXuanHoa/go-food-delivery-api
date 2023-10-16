package exceptions

import "errors"

var (
	ErrorEmptyBearer = errors.New("empty bearer header: you must contain accessToken in bearer header to validate your permission")
)
