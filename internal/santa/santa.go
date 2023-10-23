package santa

import (
	"encoding/json"
	"errors"
	"fmt"
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

func SelectGifted(inSantas []*Santa, nbGifted int) (map[string][]string, error) {
	outSantas := make(map[string][]string)
	for iGifter, s := range inSantas {
		outSantas[s.Name] = make([]string, nbGifted, nbGifted)
		for iGifted := 0; iGifted < nbGifted; iGifted++ {
			nextGifterIndex := (iGifter + iGifted + 1) % len(inSantas)
			outSantas[s.Name][iGifted] = inSantas[nextGifterIndex].Name
		}
	}

	return outSantas, nil
}

func GetClan(santaName string, santas []*Santa) (clan string, err error) {
	for _, s := range santas {
		if s.Name == santaName {
			return s.Clan, nil
		}
	}

	return "", errors.New(fmt.Sprintf("could not find %s in santas", santaName))
}
