package service

import (
	"strings"

	"github.com/google/uuid"
)

// CreateUniqueID ...
func CreateUniqueID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
