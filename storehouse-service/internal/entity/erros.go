package entity

import (
	er1 "errors"
	"fmt"
)

var (
	ErrorAlreadyExists    = fmt.Errorf("entity already exists")
	ErrorNotFound         = fmt.Errorf("entity not found")
	ErrorInvalidArguments = fmt.Errorf("invalid arguments")
	ErrInternal           = fmt.Errorf("internal error")
	ErrReqIsEmpty         = fmt.Errorf("req is empty")
	ErrUserNotRegistered  = fmt.Errorf("user not registered")
	ErrUserDeleted        = fmt.Errorf("user is deleted")
)

func ErrRegisteredOrDeleted(err error) error {
	if er1.Is(err, ErrUserDeleted) {
		return err
	}
	if er1.Is(err, ErrUserNotRegistered) {
		return err
	}
	if er1.Is(err, ErrorNotFound) {
		return err
	}
	if er1.Is(err, ErrorAlreadyExists) {
		return err
	}
	if er1.Is(err, ErrorInvalidArguments) {
		return err
	}

	return err
}
