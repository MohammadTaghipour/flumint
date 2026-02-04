package main

import (
	"flumint/internal/flutter"
	"fmt"

	"github.com/spf13/cobra"
)

// This will hold the value of the --config flag.
var configPath string

var rootCmd = &cobra.Command{
	Use:   "flumint",
	Short: "A CLI tool to automate Flutter build and packaging processes.",
	Long: `Flumint helps you create customized builds from a single
Flutter source code for different customers or environments.`,
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a default flumint.yaml configuration file in the current directory.",
	Run: func(cmd *cobra.Command, args []string) {
		// --- Your logic for the 'init' command goes here. ---
		fmt.Println("--> 'init' command called. Logic to create a template YAML file will go here.")
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds the Flutter project based on the specified configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		// --- Your logic for the 'build' command goes here. ---
		fmt.Println("--> 'build' command called. Logic to start the build process will go here.")

		// You can access the config path like this:
		fmt.Printf("Configuration file to use: %s\n", configPath)
	},
}

var listConfigsCmd = &cobra.Command{
	Use:   "list-configs",
	Short: "Lists all available build configurations (e.g., .yaml files in a configs/ directory).",
	Run: func(cmd *cobra.Command, args []string) {
		// --- Your logic for the 'list-configs' command goes here. ---
		fmt.Println("--> 'list-configs' command called. Logic to find and list YAML files will go here.")
	},
}

// The init() function is called by Go when the program starts.
// We use it to set up our commands and flags.
func init() {
	// Add a persistent flag for the config file path.
	// "Persistent" means it will be available to 'flumint' and all its subcommands.
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "flumint.yaml", "Path to the configuration file")

	// Add the subcommands to our root command.
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(listConfigsCmd)
}

// The main() function is the true entry point. Its only job is to execute the root command.
func main() {
	// init components
	f := flutter.Flutter{}
	ver, err := f.GetVersion()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ver.Version)
		fmt.Println(ver.Dart)
		fmt.Println(ver.Channel)
		fmt.Println(ver.DevTools)
	}
	doctor, err := f.RunDoctor()
	fmt.Println(doctor)

	//if err := rootCmd.Execute(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
}
