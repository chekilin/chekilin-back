package config

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestNew(t *testing.T) {
	err := godotenv.Load("./env/.env.test")
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}

	if got.Port != wantPort {
		t.Errorf("want %d, but %d", wantPort, got.Port)
	}
}
