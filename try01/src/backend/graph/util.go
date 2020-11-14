package graph

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

const cursorSeps = "#####"

func createCursor(modelName string, key int64) string {
	return base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%s%s%d", modelName, cursorSeps, key)))
}

func decodeCursor(cursor string) (string, int64, error) {
	byteArray, err := base64.RawURLEncoding.DecodeString(cursor)
	if err != nil {
		return "", 0, err
	}
	elements := strings.SplitN(string(byteArray), cursorSeps, 2)
	key, err := strconv.Atoi(elements[1])
	if err != nil {
		return "", 0, err
	}
	return elements[0], int64(key), nil
}
