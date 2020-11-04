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
	return base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%s#####%s", modelName, uniqueKey)))
}

func DecodeCursor(cursor string) (string, string, error) {
	byteArray, err := base64.RawURLEncoding.DecodeString(cursor)
	if err != nil {
		return "", "", err
	}
	elements := strings.SplitN(string(byteArray), "#####", 2)
	return elements[0], elements[1], nil
}
