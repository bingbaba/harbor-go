package harbor

import "fmt"

var (
	NotFoundError       = fmt.Errorf("not found")
	UserNotLoginError   = fmt.Errorf("user need to login")
	CallApiError        = fmt.Errorf("call api failed")
	ServerInternalError = fmt.Errorf("server internal error")
	NotAllowError       = fmt.Errorf("not allow")
)
