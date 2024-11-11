package src

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var (
	Version        = "dev"
	CommitHash     = "n/a"
	BuildTimestamp = "n/a"
)

func BuildVersion() string {
	return fmt.Sprintf("v%s-%s build time: %s", Version, CommitHash, BuildTimestamp)
}

/*

   __                            __
  / /_____  _________  ___  ____/ ____
 / __/ __ \/ ___/ __ \/ _ \/ __  / __ \
/ /_/ /_/ / /  / /_/ /  __/ /_/ / /_/ /
\__/\____/_/  / .___/\___/\__,_/\____/
             /_/


*/

var banner = `

       __________  ____  ____  __________  ____
  ____/_  __/ __ \/ __ \/ __ \/ ____/ __ \/ __ \_____
 /____// / / / / / /_/ / /_/ / __/ / / / / / / /____/
/____// / / /_/ / _, _/ ____/ /___/ /_/ / /_/ /____/
     /_/  \____/_/ |_/_/   /_____/_____/\____/

Torpedo is a clean code generator following hexagonal architecture approach
Complete documentation is available at https://darksubmarine.com/docs/torpedo
`

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Torpedo",
	Long:  `Print the version number of Torpedo`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(banner)
		fmt.Printf("Version: v%s-%s\n", Version, CommitHash)
		fmt.Printf("Build time: %s\n", BuildTimestamp)

	},
}
