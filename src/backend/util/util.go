package util

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func CreateUniqueID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func CreateCursor(modelName, uniqueKey string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", modelName, uniqueKey)))
}
