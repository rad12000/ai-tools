package cmd

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"text/template"

	"github.com/google/jsonschema-go/jsonschema"
	"gopkg.in/yaml.v3"
)

func loadSchema(templateFS fs.FS, name string) (*jsonschema.Resolved, error) {
	for i, ext := range []string{".json", ".yaml"} {
		candidate := name + ext
		contents, err := fs.ReadFile(templateFS, candidate)
		if err != nil {
			continue
		}

		isYAML := i > 0

		if isYAML {
			var yamlContents any
			err := yaml.Unmarshal(contents, &yamlContents)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal yaml schema file for %q: %w", name, err)
			}

			contents, err = json.Marshal(yamlContents)
			if err != nil {
				return nil, fmt.Errorf("failed to convert yaml schema file into json for %q: %w", name, err)
			}
		}

		var schema jsonschema.Schema
		if err := json.Unmarshal(contents, &schema); err != nil {
			return nil, fmt.Errorf("failed to unmarshal json schema for %q: %w", name, err)
		}

		resolvedSchema, err := schema.Resolve(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve schema from %q: %w", name, err)
		}

		return resolvedSchema, nil
	}

	return nil, fmt.Errorf("no schema file found for %q (looked for .json and .yaml)", name)
}

var tmplFuncs = template.FuncMap{}

// loadTemplate reads and parses the Go template for name from the given FS.
func loadTemplate(templateFS fs.FS, name string) (*template.Template, error) {
	tmplName := name + ".gotmpl"
	content, err := fs.ReadFile(templateFS, tmplName)
	if err != nil {
		return nil, fmt.Errorf("reading template file %q: %w", tmplName, err)
	}

	tmpl := template.New(name).Funcs(tmplFuncs)
	tmpl, err = tmpl.Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("parsing template %q: %w", tmplName, err)
	}

	return tmpl, nil
}
