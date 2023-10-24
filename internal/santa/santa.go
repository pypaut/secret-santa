package santa

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"time"
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
	for _, s := range inSantas {
		otherSantas := GetSantasNotInClan(s.Clan, inSantas)

		indices := GenerateRandomIndices(nbGifted, len(otherSantas))

		// Assign
		outSantas[s.Name] = make([]string, nbGifted, nbGifted)
		for i, index := range indices {
			outSantas[s.Name][i] = otherSantas[index].Name
		}
	}

	return outSantas, nil
}

func GenerateRandomIndices(nbGifted, maxIndex int) (indices []int) {
	indices = make([]int, nbGifted, nbGifted)

	for i := 0; i < nbGifted; i++ {
		rand.Seed(time.Now().UnixNano())
		newIndice := rand.Intn(maxIndex)
		for slices.Contains(indices, newIndice) {
			rand.Seed(time.Now().UnixNano())
			newIndice = rand.Intn(maxIndex)
		}

		indices[i] = newIndice
	}

	return
}

func GetClan(santaName string, santas []*Santa) (clan string, err error) {
	for _, s := range santas {
		if s.Name == santaName {
			return s.Clan, nil
		}
	}

	return "", errors.New(fmt.Sprintf("could not find %s in santas", santaName))
}

func IsOfSameClan(santaName1, santaName2 string, santas []*Santa) (isOfSameClan bool, err error) {
	clan1, err := GetClan(santaName1, santas)
	if err != nil {
		return true, err
	}

	clan2, err := GetClan(santaName2, santas)
	if err != nil {
		return true, err
	}

	return clan1 == clan2, nil
}

func GetSantasNotInClan(clan string, inSantas []*Santa) (outSantas []*Santa) {
	for _, s := range inSantas {
		if s.Clan != clan {
			outSantas = append(outSantas, s)
		}
	}

	return
}

func BiggestClanLen(santas []*Santa) int {
	clansLens := make(map[string]int)
	for _, s := range santas {
		clansLens[s.Clan]++
	}

	max := clansLens[santas[0].Name]
	for _, l := range clansLens {
		if l > max {
			max = l
		}
	}

	return max
}
