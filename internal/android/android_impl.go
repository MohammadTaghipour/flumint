package android

import "flumint/internal/config"

type Android struct {
	Config config.AndroidConfig
}

func (this *Android) ChangePackageName(newPackageName string) error {
	return nil
}
