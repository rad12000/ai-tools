package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:   "execute {template_name}",
	Short: "executes the specified template, using the provided values from stdin as arguments",
	Long: `executes the specified template, using the provided values from stdin as arguments.
	For example:
	echo '{"key": "value"}' | schema-template execute foo`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		templateFS := getTemplateDir()
		name := args[0]

		schema, err := loadSchema(templateFS, name)
		if err != nil {
			return err
		}

		tmpl, err := loadTemplate(templateFS, name)
		if err != nil {
			return err
		}

		// Read and unmarshal JSON input from stdin.
		var input any
		err = json.NewDecoder(os.Stdin).Decode(&input)
		if err != nil {
			return fmt.Errorf("reading stdin JSON: %w", err)
		}

		// Validate input against the schema, then render the template.
		if err := schema.Validate(input); err != nil {
			fmt.Fprintf(os.Stderr, "Input validation failed: %v\nTry running `%s get %s` to inspect the schema for this template", err, os.Args[0], name)
			os.Exit(1)
		}

		return tmpl.Execute(os.Stdout, input)
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)
}
