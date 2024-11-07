package mail

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"os"
	"syscall"

	"github.com/pypaut/secret-santa/internal/santa"
	myslices "github.com/pypaut/slices"
	"golang.org/x/term"
	"gopkg.in/gomail.v2"
)

type Config struct {
	SmtpAddress  string `json:"smtp-address"`
	SmtpPort     int    `json:"smtp-port"`
	EmailAddress string `json:"email-address"`
}

func LoadConfig(configFile string) (mailConfig *Config, err error) {
	file, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &mailConfig)
	if err != nil {
		return nil, err
	}

	return mailConfig, nil
}

func SendMails(config *Config, persons []*santa.Person) error {
	// Get password for mail
	fmt.Print("Mailbox password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}

	// Connect to SMTP
	d := gomail.NewDialer(
		config.SmtpAddress, config.SmtpPort, config.EmailAddress, string(bytePassword),
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Build messages
	var mails []*gomail.Message
	for _, person := range persons {
		m := gomail.NewMessage()
		m.SetHeader("From", config.EmailAddress)
		m.SetHeader("To", person.Email)
		m.SetHeader("Subject", "Secret santa !")
		strMessage := fmt.Sprintf("Bonjour %s, cette année tu seras le santa des personnes suivantes : ", person.Name)
		names, err2 := myslices.Map[*santa.Person, string](
			person.Gifted,
			func(p *santa.Person) (string, error) { return p.Name, nil },
		)
		if err2 != nil {
			return err
		}
		strMessage += fmt.Sprintf("<b>%v</b>", names)
		m.SetBody("text/html", strMessage)
		mails = append(mails, m)
	}

	// Send emails
	err = d.DialAndSend(mails...)
	if err != nil {
		return err
	}

	fmt.Print("Emails were sent\n")

	return nil
}
