package uuid

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateUuid() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
