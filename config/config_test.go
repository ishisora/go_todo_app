package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := New()
	if err != nil {
		// t.Fatalfは処理を中断する、テストは失敗
		t.Fatalf("cannot create config: %v", err)
	}
	if got.Port != wantPort {
		// t.Errorfは処理は続行するが、テストは失敗
		t.Errorf("want %d, but %d", wantPort, got.Port)
	}
	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("want %s, but %s", wantEnv, got.Env)
	}
}
