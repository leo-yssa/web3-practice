package util

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

func GenerateULID(prefix string) string {
	return fmt.Sprintf("%s%s", prefix, ulid.Make().String())
}
