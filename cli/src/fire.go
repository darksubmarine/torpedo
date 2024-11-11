package src

import (
	"errors"
	"github.com/darksubmarine/torpedo/console"
	"github.com/darksubmarine/torpedo/generator/stack/golang/goengine"
	"github.com/darksubmarine/torpedo/parserx"
	v1 "github.com/darksubmarine/torpedo/parserx/v1"
	"github.com/darksubmarine/torpedo/parserx/vx"
	"github.com/spf13/cobra"
	"os"
	"path"
)

// init command
func init() {

	rootCmd.AddCommand(fireCmd)

	if dir, err := os.Getwd(); err != nil {
		fireCmd.Flags().StringP("dir", "d", "", "Absolute path to your application directory")
	} else {
		fireCmd.Flags().StringP("dir", "d", dir, "Absolute path to your application directory")
	}
}

// Fire command
var fireCmd = &cobra.Command{
	Use:   "fire",
	Short: "Generates code from yaml files",
	Long:  `Generates application/entity code based on the defined spec at provided yaml files`,

	Run: func(cmd *cobra.Command, args []string) {
		var outputDir string
		if oDir, err := cmd.Flags().GetString("dir"); err != nil {
			console.ExitWithError(err)
		} else {
			outputDir = oDir
		}

		p := parserx.New()
		console.ExitIfErrors(p.ParseYaml(path.Join(outputDir, torpedoDir, torpedoAppYaml)))

		if p.Kind() != vx.KApp {
			console.ExitWithError(errors.New("fire only works with kind: app"))
		}

		switch p.Version() {
		case vx.V1:
			data, ok := p.Data().(v1.RootApp)
			if !ok {
				console.ExitWithError(errors.New("invalid yaml format"))
			}

			entityYamlFiles := make([]string, len(data.App.Domain.Entities))
			for i, f := range data.App.Domain.Entities {
				entityYamlFiles[i] = path.Join(outputDir, torpedoDir, torpedoEntitiesDir, f)
			}

			useCaseYamlFiles := make([]string, len(data.App.Domain.UseCases))
			for i, f := range data.App.Domain.UseCases {
				useCaseYamlFiles[i] = path.Join(outputDir, torpedoDir, torpedoUseCasesDir, f)
			}

			// GoLang is only supported so far.
			if data.App.Stack.Lang != v1.Go {
				console.ExitWithError(errors.New("only Go stack is supported"))
			} else {
				opts := goengine.DefaultOptionsForApp(outputDir, data.App.Stack.Package, entityYamlFiles, useCaseYamlFiles)
				engine := goengine.New(opts)
				console.ExitIfErrors(engine.Fire())
			}

			console.Println("Entities parsed:")
			for _, f := range entityYamlFiles {
				console.Println("\t-", f)
			}

			console.Println("Use cases parsed:")
			for _, f := range useCaseYamlFiles {
				console.Println("\t-", f)
			}

		default:
			console.ExitWithError(errors.New("fire only v1 is supported"))

		}
		console.ExitOk("\n  + Build has been successfully!")
	},
}
