package config_test

import (
	"os"
	"testing"

	"cheki-back/config"
)

func TestNew(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}

	want := os.Getenv("CHEKI_DB_USER")
	if got := cfg.DBUser; got != want {
		t.Errorf("Got %s, want %s", got, want)
	}

	want = os.Getenv("CHEKI_DB_PASSWORD")
	if got := cfg.DBPassword; got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}
