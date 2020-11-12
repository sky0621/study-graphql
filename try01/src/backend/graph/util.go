package graph

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func createCursor(modelName string, key int64) string {
	return base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%s#####%d", modelName, key)))
}

func decodeCursor(cursor string) (string, string, error) {
	byteArray, err := base64.RawURLEncoding.DecodeString(cursor)
	if err != nil {
		return "", "", err
	}
	elements := strings.SplitN(string(byteArray), "#####", 2)
	return elements[0], elements[1], nil
}
