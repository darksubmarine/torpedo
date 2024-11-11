package src

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "torpedo",
	Short: "Torpedo is a clean code generator following hexagonal architecture approach",
	Long: `
       __________  ____  ____  __________  ____      
  ____/_  __/ __ \/ __ \/ __ \/ ____/ __ \/ __ \_____
 /____// / / / / / /_/ / /_/ / __/ / / / / / / /____/
/____// / / /_/ / _, _/ ____/ /___/ /_/ / /_/ /____/ 
     /_/  \____/_/ |_/_/   /_____/_____/\____/       
            
Torpedo is a clean code generator following hexagonal architecture approach
Complete documentation is available at https://darksubmarine.com/docs/torpedo
	`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
