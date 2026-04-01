/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get {template_name}",
	Short: "Get the JSON Schema for the specified template, and write it to STDOUT.",
	Long:  "Get the JSON Schema for the specified template, and write it to STDOUT.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		templateFS := getTemplateDir()
		name := args[0]

		// Verify a .gotmpl file exists for this template name.
		if _, err := loadTemplate(templateFS, name); err != nil {
			return err
		}

		// Load and compile the schema, then print it to stdout.
		schema, err := loadSchema(templateFS, name)
		if err != nil {
			return err
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(schema)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
