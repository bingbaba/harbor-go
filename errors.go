package harbor

import (
	"errors"
	"fmt"
)

var (
	NotFoundError       = fmt.Errorf("not found")
	UserNotLoginError   = fmt.Errorf("user need to login")
	CallApiError        = fmt.Errorf("call api failed")
	ServerInternalError = fmt.Errorf("server internal error")
	NotAllowError       = fmt.Errorf("not allow")
)

var (
	ERROR_THE_GINSENG     = errors.New("Passing parameters cannot nil")
	ERROR_THE_FORMAT      = errors.New("Unsatisfied with constraints of the user creation.")
	ERROR_THE_PERMISSIONS = errors.New("User registration can only be used by admin role user when self-registration is off")
	ERROR_THE_TYPE        = errors.New("The Media Type of the request is not supported, it has to be 'application/json'")
	ERROR_THE_SERVER      = errors.New("Unexpected internal errors")
	ERROR_THE_PKG         = errors.New("Running errors")
	ERROR_THE_401         = errors.New("User need to log in first.")
	ERROR_THE_409         = errors.New("Maybe existing")
)
