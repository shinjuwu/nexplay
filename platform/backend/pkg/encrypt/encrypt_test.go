package encrypt_test

import (
	"backend/pkg/encrypt"
	"log"
	"testing"
)

func TestTimeFormat(t *testing.T) {

	// The key must be 16, 24, 32 bytes
	secret_key := "6b14314760bd9280695a95d38082478b"
	orgData := "123456789999"
	enData, err := encrypt.EncryptSaltToken(orgData, secret_key)
	if err != nil {
		t.Errorf("Encrypt error: %v", err)
	}
	log.Println(enData)

	deData, err := encrypt.DecryptSaltToken(enData, secret_key)
	if err != nil {
		t.Errorf("Encrypt error: %v", err)
	}
	log.Println(deData)
}
