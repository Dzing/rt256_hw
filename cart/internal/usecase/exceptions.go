package usecase

import (
	"fmt"
)

type (
	CartIsEmptyError struct {
		User TUserId
	}
)

func (e *CartIsEmptyError) Error() string {
	return fmt.Sprintf("cart is empty userId = %v", e.User)
}

var ErrCartIsEmpty *CartIsEmptyError
