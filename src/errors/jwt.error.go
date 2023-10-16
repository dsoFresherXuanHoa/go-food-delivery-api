package exceptions

import "errors"

var (
	ErrInvalidAccessToken = errors.New("invalid accessToken: token was expired or damaged")
	ErrInvalidAuthClaims  = errors.New("invalid authClaim struct: don't try access this resource if you don't has permission")
)
