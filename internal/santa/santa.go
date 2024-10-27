package santa

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"time"

	myslices "github.com/pypaut/slices"
)

var ErrBadShuffle = errors.New("bad shuffle")

type Person struct {
	Name    string `yaml:"name"`
	Clan    string `yaml:"clan"`
	Email   string `yaml:"email"`
	Gifted  []*Person
	NbGifts int
}

func SecretSanta(config string, nbGifts int) (persons []*Person) {
		persons, err := loadPersons(config)
		if err != nil {
			panic("AAAAAAAAH")
		}

		initChecks(persons, nbGifts)

		for {
			shadowErr := findSantas(persons, nbGifts)
			if shadowErr != nil {
				switch {
				case errors.Is(shadowErr, ErrBadShuffle):
					resetSantas(persons)
					continue
				default:
					panic(shadowErr)
				}
			}

			break
		}

		finalChecks(persons, nbGifts)
		return
}

func findSantas(persons []*Person, nbSantas int) error {
	for _, p := range persons {
		otherPersons := myslices.Filter(
			persons, func(per *Person) bool {
				return per.NbGifts < nbSantas && per.Clan != p.Clan
			},
		)

    for i:=0; i<nbSantas; i++ {
      // Get random other person
      indices, err := generateRandomIndices(1, len(otherPersons))
      if err != nil {
        return err
      }

      // Assign this person to our current santa
      p.Gifted = append(p.Gifted, otherPersons[indices[0]])
      otherPersons[indices[0]].NbGifts++
      clanToRm := otherPersons[indices[0]].Clan

      // Rm this person's clan from the list (make this an option later)
      otherPersons = myslices.Filter(
        otherPersons, func(per *Person) bool {
          return per.Clan != clanToRm
        },
      )
    }
	}

	return nil
}

func resetSantas(persons []*Person) {
	for _, p := range persons {
		p.Gifted = nil
		p.NbGifts = 0
	}
}

func finalChecks(persons []*Person, nbSantas int) {
	for _, p := range persons {
    // Number of santas
		if len(p.Gifted) != nbSantas {
			panic("wrong number of santas")
		}

    // Number of gifts
		if p.NbGifts != nbSantas {
			panic("wrong number of gifts")
		}

    clansGifted, err := myslices.Map(p.Gifted, func(per *Person) (string, error) {
      return per.Clan, nil
    })

    // Offers to different clans only
    if err != nil {
      panic(err)
    }
    if slices.Contains(clansGifted, p.Clan) {
      panic("gifts to its own clan")
    }

    // Doesn't offer several times to the same clan (make it an option later)
    if myslices.CheckDuplicates(clansGifted) {
      panic("multiple gifts to the same clan")
    }
	}
}

func initChecks(persons []*Person, nbSantas int) {
	if len(persons) < nbSantas {
		panic("number santas > number people")
	}

	maxLen := biggestClanLen(persons)
	if len(persons)-maxLen < nbSantas {
		panic("one clan is too big")
	}
}

func loadPersons(configFile string) (persons []*Person, err error) {
	file, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &persons)
	if err != nil {
		return nil, err
	}

	return
}

func (p *Person) String() string {
	gifted, err := myslices.Map(p.Gifted, func(per *Person) (string, error) { return per.Name, nil })
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s -> %v", p.Name, gifted)
}

func generateRandomIndices(nbIndices, maxIndex int) (indices []int, err error) {
	indices = make([]int, 0, nbIndices)

	if maxIndex < nbIndices {
		return nil, ErrBadShuffle
	} else if nbIndices == maxIndex {
		for i := 0; i < maxIndex; i++ {
			indices = append(indices, i)
		}

		return indices, nil
	}

	for i := 0; i < nbIndices; i++ {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		newIndice := rand.Intn(maxIndex)
		for slices.Contains(indices, newIndice) {
			rand.New(rand.NewSource(time.Now().UnixNano()))
			newIndice = rand.Intn(maxIndex)
		}

		indices = append(indices, newIndice)
	}

	return
}

func biggestClanLen(santas []*Person) int {
	clansLens := make(map[string]int)
	for _, s := range santas {
		clansLens[s.Clan]++
	}

	theMax := clansLens[santas[0].Name]
	for _, l := range clansLens {
		if l > theMax {
			theMax = l
		}
	}

	return theMax
}
