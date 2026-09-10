package cmd

import (
	"fmt"
	"os/exec"

	"github.com/STELLAR-PATH/stellar-scaffold/internal/templates"
	"github.com/STELLAR-PATH/stellar-scaffold/pkg/generator"
	"github.com/spf13/cobra"
)

var (
	templateName string
	initGit      bool
)

var newCmd = &cobra.Command{
	Use:   "new <project-name>",
	Short: "Scaffold a new Stellar & Soroban smart contract project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		
		if templateName != "basic" && templateName != "token" {
			return fmt.Errorf("invalid template '%s', options are 'basic' or 'token'", templateName)
		}

		gen := &generator.Generator{
			TemplatesFS: templates.FS,
			ProjectName: projectName,
			Template:    templateName,
			OutputDir:   projectName,
		}

		fmt.Printf("Scaffolding project '%s' using template '%s'...\n", projectName, templateName)
		
		if err := gen.Generate(); err != nil {
			return err
		}

		if initGit {
			fmt.Println("Initializing git repository...")
			gitCmd := exec.Command("git", "init")
			gitCmd.Dir = projectName
			if err := gitCmd.Run(); err != nil {
				return fmt.Errorf("failed to initialize git: %w", err)
			}
		}

		fmt.Printf("Successfully generated project %s!\n", projectName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVarP(&templateName, "template", "t", "basic", "Project template (options: 'basic', 'token')")
	newCmd.Flags().BoolVar(&initGit, "git", true, "Initialize a git repo inside the generated folder")
}
