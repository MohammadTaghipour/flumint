package service

import (
	"fmt"

	"github.com/MohammadTaghipour/flumint/internal/flutter"
	"github.com/MohammadTaghipour/flumint/internal/utils"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

func RunDoctor(cmd *cobra.Command) error {
	fmt.Println()

	s := spinner.New(spinner.CharSets[utils.SpinnerCharset], utils.SpinnerDuration)
	s.Suffix = " Running Flumint doctor..."
	s.Color(utils.SpinnerColor)
	s.Start()

	flutterV, err := flutter.GetVersion()
	if err != nil {
		s.Stop()
		fmt.Println("Flumint doctor failed ✖")
		return fmt.Errorf("failed to get flutter version: %w", err)
	}

	s.Stop()

	fmt.Println(utils.BrandWriter("Flumint Doctor Information"))
	fmt.Println("------------------------")
	fmt.Println(utils.SuccessWriter("Version    : " + flutterV.Version))
	fmt.Println(utils.SuccessWriter("Channel    : " + flutterV.Channel))
	fmt.Println(utils.SuccessWriter("Dart       : " + flutterV.Dart))
	fmt.Println(utils.SuccessWriter("DevTools   : " + flutterV.DevTools))

	fmt.Println()

	return nil
}
