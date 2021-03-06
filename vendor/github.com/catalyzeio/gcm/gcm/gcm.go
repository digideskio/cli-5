package gcm

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"
)

const (
	chunkSize = 1024 * 1024 // 1 MB

	// AAD (Additional authenticated data) is to be used in the GCM algorithm
	AAD = "7f57c07ee9459ed704d5e403086f6503"
)

// EncryptFile encrypts the file at the specified path using GCM
func EncryptFile(inFilePath, outFilePath string, key, givenIV, aad []byte) error {
	if _, err := os.Stat(inFilePath); os.IsNotExist(err) {
		return fmt.Errorf("A file does not exist at %s", inFilePath)
	}
	// copy the IV since it will potentially be incremented
	iv := make([]byte, len(givenIV))
	copy(iv, givenIV)

	inFile, err := os.Open(inFilePath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.OpenFile(outFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer outFile.Close()

	aes, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCMWithNonceSize(aes, len(iv))
	if err != nil {
		return err
	}

	chunk := make([]byte, chunkSize)
	written := false
	for {
		read, err := inFile.Read(chunk)
		// ensure we have written at least one chunk before breaking
		if read > 0 || !written {
			encrChunk := gcm.Seal(nil, iv, chunk[:read], aad)
			if _, err := outFile.Write(encrChunk); err != nil {
				return err
			}
			written = true
		}
		if err == io.EOF {
			break
		}
		incrementIV(iv)
	}
	return nil
}

// DecryptFile decrypts the file at the specified path using GCM
func DecryptFile(inFilePath, outFilePath string, key, givenIV, aad []byte) error {
	if _, err := os.Stat(inFilePath); os.IsNotExist(err) {
		return fmt.Errorf("A file does not exist at %s", inFilePath)
	}
	// copy the IV since it will potentially be incremented
	iv := make([]byte, len(givenIV))
	copy(iv, givenIV)

	inFile, err := os.Open(inFilePath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.OpenFile(outFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer outFile.Close()

	aes, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCMWithNonceSize(aes, len(iv))
	if err != nil {
		return err
	}

	written := false
	chunk := make([]byte, chunkSize+gcm.Overhead())
	for {
		read, err := inFile.Read(chunk)
		// ensure we have written at least one chunk before breaking
		if read > 0 || !written {
			decrChunk, err := gcm.Open(nil, iv, chunk[:read], aad)
			if err != nil {
				return err
			}
			if _, err := outFile.Write(decrChunk); err != nil {
				return err
			}
			written = true
		}
		if err == io.EOF {
			break
		}
		incrementIV(iv)
	}
	return nil
}

func incrementIV(iv []byte) {
	for i := len(iv) - 1; i >= 0; i-- {
		iv[i]++
		if iv[i] != 0 {
			return
		}
	}
}
