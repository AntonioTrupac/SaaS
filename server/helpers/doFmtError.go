package helpers

import (
	"fmt"
)

func DoFmtError(errorText string) error {
	errCode := 401
	err := fmt.Errorf(errorText, errCode)
	return err
}
