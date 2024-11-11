package src

import (
	"github.com/darksubmarine/torpedo/console"
	"github.com/spf13/cobra"
	"os"
	"path"
)

func init() {
	rootCmd.AddCommand(initCmd)

	if dir, err := os.Getwd(); err != nil {
		initCmd.Flags().StringP("dir", "d", "", "Absolute path to your application directory")
	} else {
		initCmd.Flags().StringP("dir", "d", dir, "Absolute path to your application directory")
	}
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes an empty application project",
	Long:  `Initializes an empty application project`,
	Run: func(cmd *cobra.Command, args []string) {
		if outputDir, err := cmd.Flags().GetString("dir"); err != nil {
			console.ExitWithError(err)
		} else {
			console.ExitIfError(os.MkdirAll(path.Join(outputDir, torpedoDir, torpedoEntitiesDir, torpedoDocsDir), os.ModePerm))
			console.ExitIfError(os.MkdirAll(path.Join(outputDir, torpedoDir, torpedoUseCasesDir, torpedoDocsDir), os.ModePerm))
		}

	}}
