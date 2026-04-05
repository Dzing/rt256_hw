package clientgrpc

import (
	"fmt"
)

type (
	NoConnectionError struct {
		Addr string
	}
)

func (e *NoConnectionError) Error() string {
	return fmt.Sprintf("no connection to gRPC server: =%v", e.Addr)
}

var ErrNoConnection *NoConnectionError
