package generator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Generator handles scaffolding a project from an embedded template
type Generator struct {
	TemplatesFS fs.FS
	ProjectName string
	Template    string
	OutputDir   string
}

func (g *Generator) Generate() error {
	basePath := g.Template
	
	err := fs.WalkDir(g.TemplatesFS, basePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Calculate relative path from template base
		var relPath string
		if path == basePath {
			relPath = "."
		} else {
			relPath = strings.TrimPrefix(path, basePath+"/")
		}

		if relPath == "." {
			return nil
		}

		// Replace placeholders in path
		destPath := strings.ReplaceAll(relPath, "tmpl_project_name", g.ProjectName)
		destPath = filepath.Join(g.OutputDir, destPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		// Read file content
		content, err := fs.ReadFile(g.TemplatesFS, path)
		if err != nil {
			return err
		}

		// Replace placeholders in content
		contentStr := string(content)
		contentStr = strings.ReplaceAll(contentStr, "{{.ProjectName}}", g.ProjectName)

		// Create directories for the file if they don't exist
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return err
		}

		// Write to destination
		return os.WriteFile(destPath, []byte(contentStr), 0644)
	})

	if err != nil {
		return fmt.Errorf("failed to generate project: %w", err)
	}

	return nil
}
