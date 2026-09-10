package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"
)

func TestGenerator(t *testing.T) {
	mockFS := fstest.MapFS{
		"internal/templates/basic/Cargo.toml": {Data: []byte("workspace")},
		"internal/templates/basic/contracts/tmpl_project_name/Cargo.toml": {Data: []byte("name = \"{{.ProjectName}}\"")},
		"internal/templates/basic/README.md": {Data: []byte("# {{.ProjectName}}")},
	}

	tempDir := t.TempDir()

	gen := &Generator{
		TemplatesFS: mockFS,
		ProjectName: "my_token",
		Template:    "basic",
		OutputDir:   tempDir,
	}

	err := gen.Generate()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify file was created and templated
	content, err := os.ReadFile(filepath.Join(tempDir, "README.md"))
	if err != nil {
		t.Fatalf("Failed to read README.md: %v", err)
	}
	if !strings.Contains(string(content), "# my_token") {
		t.Errorf("README.md templating failed. Got: %s", string(content))
	}

	// Verify project name replacement in path
	_, err = os.Stat(filepath.Join(tempDir, "contracts/my_token/Cargo.toml"))
	if err != nil {
		t.Fatalf("Failed to create templated path: %v", err)
	}
	
	// Verify project name replacement in content
	cargoTomlContent, err := os.ReadFile(filepath.Join(tempDir, "contracts/my_token/Cargo.toml"))
	if err != nil {
		t.Fatalf("Failed to read Cargo.toml: %v", err)
	}
	if !strings.Contains(string(cargoTomlContent), "name = \"my_token\"") {
		t.Errorf("Cargo.toml templating failed. Got: %s", string(cargoTomlContent))
	}
}
