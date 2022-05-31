package err

import "fmt"

type BizError struct {
	StatusCode int
	BizCode    string
	Msg        string
}

func (be *BizError) Error() string {
	return fmt.Sprintf("code: %s  status: %d", be.BizCode, be.StatusCode)
}
