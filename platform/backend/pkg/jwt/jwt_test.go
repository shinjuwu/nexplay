package jwt_test

import (
	"backend/pkg/config"
	"backend/pkg/jwt"
	"testing"
)

func TestJwt(t *testing.T) {

	// load default config
	c := config.NewConfig()

	jwtObj := jwt.NewJwtManager(c)

	customClaim := jwtObj.CreateClaims(jwt.BaseClaims{})

	t.Logf("CreateClaims uerClaim is: %v", customClaim.ExpiresAt)

	token, err := jwtObj.GenerateToken(customClaim)
	if err != nil {
		t.Errorf("GenerateToken has error: %v", err)
		t.Fail()
	}

	t.Logf("GenerateToken token is: %v", token)

	claims, err := jwtObj.ParseToken(token)
	if err != nil {
		t.Errorf("ParseToken has error: %v", err)
		t.Fail()
	}

	t.Logf("ParseToken claims is: %v", claims)
}
