package oslinfra

import (
	"testing"
)

func TestRegionUpdate(t *testing.T) {
	r := &Region{
		name: "ap-southeast-1",
	}
	newR := &Region{
		name: "cn-north-1",
	}
	invalidR := &Region{
		name: "us-west-2",
	}

	// Valid region
	err := r.Update(newR)
	if err != nil {
		t.Errorf("expected %s, actual %s", newR.name, r.name)
	}

	// Invalid region
	err = r.Update(invalidR)
	if err != ErrInvalidRegion {
		t.Errorf("expected %s, actual %s", ErrInvalidRegion, err)
	}
}

func TestEnvUpdate(t *testing.T) {
	env := &Env{
		name: "prod",
	}
	newEnv := &Env{
		name: "stage",
	}
	invalidEnv := &Env{
		name: "alien",
	}

	// Valid region
	err := env.Update(newEnv)
	if err != nil {
		t.Errorf("expected %s, actual %s", newEnv.name, env.name)
	}

	// Invalid region
	err = env.Update(invalidEnv)
	if err != ErrInvalidEnv {
		t.Errorf("expected %s, actual %s", ErrInvalidEnv, err)
	}
}
