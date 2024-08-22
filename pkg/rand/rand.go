package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/oklog/ulid/v2"
)

type Rand []byte

func MakeULID(prefix string) string {
	return fmt.Sprintf("%s%s", prefix, ulid.Make().String())
}

func MakeState() (state string, err error) {
	var r Rand
	err = r.Set(32)
	state = r.ToBase64()
	return state, err
}

func MakeUUID() (uuid string, err error) {
	var r Rand
	err = r.Set(16)
	r[8] = r[8]&^0xc0 | 0x80
	r[6] = r[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", r[0:4], r[4:6], r[6:8], r[8:10], r[10:]), err
}

func (r *Rand) Set(length int) error {
	*r = make([]byte, length)
	_, err := io.ReadFull(rand.Reader, *r)
	if err != nil {
		return err
	}
	return nil
}

func (r *Rand) ToBase64() string {
	return base64.StdEncoding.EncodeToString(*r)
}
