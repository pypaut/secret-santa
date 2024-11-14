package cmd

import (
	"os"

	"github.com/pypaut/secret-santa/internal/mail"
	"github.com/pypaut/secret-santa/internal/santa"
	"github.com/pypaut/secret-santa/internal/show"
	"github.com/spf13/cobra"

	"fmt"
)

var (
	configFile     string
	mailConfigFile string
	nbGifts        int
	showMailJson   bool
	showSantasJson bool
	withMail       bool
)

var rootCmd = &cobra.Command{
	Use:   "secret-santa",
	Short: "Find out your santas for this year!",
	Long: `Usage example:
./secret-santa --nb_gifts 2 --config my-custom-conf.json --with_mail

Config files:
- santas.json
- mail.json

Get templates of these configuration files by running respectively:
- ./secret-santa --show_santas_json
- ./secret-santa --show_mail_json`,
	Run: func(cmd *cobra.Command, args []string) {
		if showSantasJson {
			show.SantasJsonTemplate()
			return
		}
		if showMailJson {
			show.MailJsonTemplate()
			return
		}

		persons := santa.SecretSanta(configFile, nbGifts)
		for _, p := range persons {
			fmt.Printf("%v\n", p)
		}

		if withMail {
			mailConfig, err := mail.LoadConfig("mail-conf.json")
			if err != nil {
				fmt.Print("error while loading mail config\n")
				panic(err)
			}

			err = mail.SendMails(mailConfig, persons)
			if err != nil {
				fmt.Print("error while sending mail\n")
				panic(err)
			}
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
	rootCmd.Flags().BoolVar(&showMailJson, "show_mail_json", false, "Whether to show the mail.json template")
	rootCmd.Flags().BoolVar(&showSantasJson, "show_santas_json", false, "Whether to show the santas.json template")
	rootCmd.Flags().BoolVar(&withMail, "with_mail", false, "Whether to send mails to the persons or not")
	rootCmd.Flags().IntVar(&nbGifts, "nb_gifts", 2, "How many persons should receive gifts from a single person")
	rootCmd.Flags().StringVar(&configFile, "config", "santas.json", "File containing the list of santas")
	rootCmd.Flags().StringVar(&mailConfigFile, "mail_config", "mail.json", "File containing the SMTP configuration to send mails")
}
