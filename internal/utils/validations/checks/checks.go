package checks

import "github.com/google/uuid"

func IsUUIDValid(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}
