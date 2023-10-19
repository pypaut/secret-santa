package santa

import (
	"encoding/json"
	"os"
)

type Santa struct {
	Name  string `yaml:"name"`
	Clan  string `yaml:"clan"`
	Email string `yaml:"email"`
}

func LoadSantas(configFile string) (santas []*Santa, err error) {
	file, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &santas)
	if err != nil {
		return nil, err
	}

	return
}
