package rpc

import (
	"crypto/rand"
	"fmt"
)

// Note - NOT RFC4122 compliant
func pseudo_uuid() (uuid string) {
	b := make([]byte, 16)
	_, _ = rand.Read(b)

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}
