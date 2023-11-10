package exceptions

import "errors"

var (
	ErrorFinishRejectedOrder                = errors.New("finish rejected order: can't finish rejected order")
	ErrorCompensatedFinishedOrRejectedOrder = errors.New("compensated rejected or finished order: can't compensated finish rejected or finished order")
	ErrorInvalidSecretCode                  = errors.New("invalid secretCode: can't create a booking because your secretCode is invalid, check your secretCode again")
	ErrReasonEmpty                          = errors.New("invalid reason: reason can not be empty")
)
