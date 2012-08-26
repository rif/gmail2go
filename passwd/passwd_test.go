package passwd

import (
	"bytes"
	"testing"
)

func TestEncryptDescrypt(t *testing.T) {
	var dst bytes.Buffer
	err := Encrypt(&dst, bytes.NewBufferString("mama are mere"), make([]byte, 16), make([]byte, 16))
	if err != nil {
		t.Error("Encryption failed: ", err)
	}
	res, err := Decrypt(&dst, make([]byte, 16), make([]byte, 16))
	if err != nil {
		t.Error("Descryption failed: ", err)
	}
	if res.String() != "mama are mere" {
		t.Error("Process failed: ", res.String())
	}
}
