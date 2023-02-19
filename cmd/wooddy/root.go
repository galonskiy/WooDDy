package wooddy

import (
	"os"
	"wooddy/internal/md2bash"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wooddy file [files]",
	Short: "A brief description of your application",
	Long: `
WooDDy is a simple command line utility for running scripts from md files`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			execute, _ := cmd.Flags().GetBool("execute")
			saveFilename, _ := cmd.Flags().GetString("save")

			md2bash.ReadMds(args, execute, saveFilename)
		} else {
			cmd.Help()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("save", "s", "", "Help message for save")
	rootCmd.Flags().BoolP("execute", "e", false, "Help message for execute")
}
