package cmd

import (
	"fmt"
	"os"

	"github.com/santhosh-tekuri/jsonschema/v6"
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
		input, err := jsonschema.UnmarshalJSON(os.Stdin)
		if err != nil {
			return fmt.Errorf("reading stdin JSON: %w", err)
		}

		// Validate input against the schema, then render the template.
		if err := schema.Validate(input); err != nil {
			return fmt.Errorf("validation failed: %w", err)
		}

		return tmpl.Execute(os.Stdout, input)
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// executeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// executeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
