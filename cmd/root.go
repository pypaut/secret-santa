package cmd

import (
	"os"

	"github.com/pypaut/secret-santa/internal/santa"
	"github.com/spf13/cobra"

	"fmt"
)

var (
	configFile string
	nbGifts    int
)

var rootCmd = &cobra.Command{
	Use:   "secret-santa",
	Short: "Find out your santas for this year!",
	Long: `Fill a santas.json file with all the santas. "Clan" means the
  household: it is relevant in case you don't want to link santas that live
  in the same household. Usage example:

  ./secret-santa --nb_gifts 2 --file santas.json`,
	Run: func(cmd *cobra.Command, args []string) {
    persons := santa.SecretSanta(configFile, nbGifts)

		for _, p := range persons {
			fmt.Printf("%v\n", p)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVar(&nbGifts, "nb_gifts", 2, "How many persons should receive gifts from a single person (default: 2)")
	rootCmd.Flags().StringVar(&configFile, "config", "santas-sample.json", "File containing the list of santas (default: santas.json)")
}
