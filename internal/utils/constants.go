package utils

import (
	"time"

	"github.com/fatih/color"
)

const (
	SpinnerColor    string        = "fgHiBlue"
	SpinnerCharset  int           = 14
	SpinnerDuration time.Duration = 90 * time.Millisecond
)

var (
	ErrorWriter   = color.New(color.FgHiRed).SprintFunc()
	SuccessWriter = color.New(color.FgHiGreen).SprintFunc()
	InfoWriter    = color.New(color.FgHiYellow).SprintFunc()
	BrandWriter   = color.New(color.FgHiBlue).SprintFunc()
)
