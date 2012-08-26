package passwd

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"
)

func Encrypt(dst io.Writer, data *bytes.Buffer, key, iv []byte) (err error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	w := &cipher.StreamWriter{S: cipher.NewOFB(c, iv), W: dst}
	io.Copy(w, data)
	return
}

func Decrypt(src io.Reader, key, iv []byte) (data *bytes.Buffer, err error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	r := &cipher.StreamReader{S: cipher.NewOFB(c, iv), R: src}
	data = new(bytes.Buffer)
	io.Copy(data, r)
	return
}
