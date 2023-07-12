package jwts

import "testing"

func TestParesToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODE0NzMwMjksInRva2VuIjoiMTAwMyJ9.FP6WpwTpRi_QVtDtsE2ulvYpy0B3bZs6ahCnnPWBl0M"
	ParseToken(tokenString, "msproject")

}
