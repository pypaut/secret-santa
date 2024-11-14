package show

import "fmt"

func SantasJsonTemplate() {
	fmt.Print(`[
  {
    "name": "John",
    "clan": "Cook",
    "email": "john@cook.io"
  },
  {
    "name": "James",
    "clan": "Potter",
    "email": "james@potter.uk"
  },
  {
    "name": "Ginny",
    "clan": "Potter",
    "email": "ginny@potter.uk"
  },
  {
    "name": "Jaina",
    "clan": "Proudmore",
    "email": "jaina@proudmore.com"
  },
  {
    "name": "Walther",
    "clan": "Heiseinberg",
    "email": "walther@white.us"
  },
  {
    "name": "Megumi",
    "clan": "Fushiguro",
    "email": "megumi@fushiguro.jp"
  },
  {
    "name": "Toji",
    "clan": "Fushiguro",
    "email": "toji@fushiguro.jp"
  }
]
`)
}

func MailJsonTemplate() {
	fmt.Printf(`{
  "smtp-address": "",
  "smtp-port": 587,
  "email-address": ""
}`)
}
