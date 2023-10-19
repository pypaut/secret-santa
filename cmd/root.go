package cmd

import (
	"os"

	"github.com/pypaut/secret-santa/internal/santa"
	"github.com/spf13/cobra"

	"fmt"
)

var (
	configFile string
	nbSantas   int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "secret-santa",
	Short: "Find out your santas for this year!",
	Long: `Fill a santas.json file with all the santas. "Clan" means the
household: it is relevant in case you don't want to link santas that live
in the same household. Usage example:

./secret-santa --nb_santas 2 --file santas.json`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load santas
		santas, err := santa.LoadSantas(configFile)
		if err != nil {
			panic(err)
		}

		for _, s := range santas {
			fmt.Printf("%v\n", s)
		}

		// Select people for each santa

		// Send mail to each santa with their people to gift
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVar(&nbSantas, "nb_santas", 1, "Number of santas to attribute each santa")
	rootCmd.Flags().StringVar(&configFile, "config", "santas.json", "File containing the list of santas (default: santas.json)")
}
