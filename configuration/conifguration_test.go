package configuration

import (
	"os"
	"testing"
)

func TestConfiguration(t *testing.T) {
	os.Setenv("PORT", "2000")

	InitEnv()

	if StaticConfiguration.Port != "2000" {
		t.Fatalf("expected the port value at 2000 and gets %s\n", StaticConfiguration.Port)
	}

	v := os.Getenv("PORT")
	if v != "" {
		t.Fatalf("expected empty env variable and gets %s\n", v)
	}
}
