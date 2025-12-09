package android

import (
	"flumint/internal/config"
)

type Android struct {
	Config config.AndroidConfig
}

func (android *Android) ChangePackageName(newPackageName string) error {
	return nil
}
