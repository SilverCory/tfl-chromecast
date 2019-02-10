package env_test

import (
	"os"
	"testing"
	. "tfl-chromecast/env"
)

func TestGet(t *testing.T) {
	if err := os.Setenv("ENV", "TESTING"); err != nil {
		panic(err)
	}

	var e Config
	e = Get()

	if e.Env.Production() {
		t.Fatal("WE ARE NOT IN PRODUCTION (I hope)")
	}

	if !e.Env.Testing() {
		t.Fatal("WE ARE IN TESTING (I hope)")
	}
}
