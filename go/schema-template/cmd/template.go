/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"text/template"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// findSchemaFile looks for <name>.json then <name>.yaml in the given FS,
// returning the filename of the first match or an error if neither exists.
func findSchemaFile(templateFS fs.FS, name string) (string, error) {
	for _, ext := range []string{".json", ".yaml"} {
		candidate := name + ext
		if _, err := fs.Stat(templateFS, candidate); err == nil {
			return candidate, nil
		}
	}
	return "", fmt.Errorf("no schema file found for %q (looked for .json and .yaml)", name)
}

// loadSchema finds and compiles the JSON schema for name from the given FS.
func loadSchema(templateFS fs.FS, name string) (*jsonschema.Schema, error) {
	schemaFile, err := findSchemaFile(templateFS, name)
	if err != nil {
		return nil, err
	}

	schema, err := jsonschema.NewCompiler().Compile(schemaFile)
	if err != nil {
		return nil, fmt.Errorf("compiling schema %q: %w", schemaFile, err)
	}

	return schema, nil
}

// loadTemplate reads and parses the Go template for name from the given FS.
func loadTemplate(templateFS fs.FS, name string) (*template.Template, error) {
	tmplName := name + ".gotmpl"
	content, err := fs.ReadFile(templateFS, tmplName)
	if err != nil {
		return nil, fmt.Errorf("reading template file %q: %w", tmplName, err)
	}

	tmpl, err := template.New(name).Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("parsing template %q: %w", tmplName, err)
	}

	return tmpl, nil
}
