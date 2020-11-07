package graph

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func CreateCursor(modelName string, key interface{}) string {
	return base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%s#####%v", modelName, key)))
}

func DecodeCursor(cursor string) (string, string, error) {
	byteArray, err := base64.RawURLEncoding.DecodeString(cursor)
	if err != nil {
		return "", "", err
	}
	elements := strings.SplitN(string(byteArray), "#####", 2)
	return elements[0], elements[1], nil
}
