package oslinfra

import (
	"errors"
)

var (
	ValidRegions = []string{
		"ap-southeast-1",
		"cn-north-1",
		"eu-west-1",
	}
	ValidEnvs = []string{
		"prod",
		"prod-cn",
		"stage",
		"stage-cn",
		"test",
		"dev",
	}
	ErrInvalidRegion = errors.New("Invalid region")
	ErrInvalidEnv    = errors.New("Invalid Env")
)

type Region struct {
	name string
}

type Env struct {
	name string
}

type Resource interface {
	Update(newer *Resource) error
}

func (r *Region) Update(newR *Region) error {
	for _, validRegion := range ValidRegions {
		if newR.name == validRegion {
			r.name = newR.name
			return nil
		}
	}
	return ErrInvalidRegion
}

func (env *Env) Update(newEnv *Env) error {
	for _, validEnv := range ValidEnvs {
		if newEnv.name == validEnv {
			env.name = newEnv.name
			return nil
		}
	}
	return ErrInvalidEnv
}
