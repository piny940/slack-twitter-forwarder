package api

import (
	"slices"
	"strings"
	"testing"
)

func TestOauthNonce(t *testing.T) {
	nonces := make([]string, 0)
	for i := 0; i < 100; i++ {
		nonce := getOauthNonce()
		for _, letter := range oauthNgLetters {
			if strings.Contains(nonce, letter) {
				t.Errorf("Nonce must not contain %s, but nonce is %s", letter, nonce)
				return
			}
		}
		if slices.Contains(nonces, nonce) {
			t.Errorf("Nonce must be unique, but nonce is %s", nonce)
			return
		}
		nonces = append(nonces, nonce)
	}
}
