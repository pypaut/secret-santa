package santa

import "testing"

func TestResetSantas(t *testing.T) {
	persons, err := loadPersons("../../santas-test.json")
	if err != nil {
		t.Errorf("error loading persons: %v", err)
		return
	}

	persons[0].NbGifts = 3
	resetSantas(persons)

	if persons[0].NbGifts != 0 {
		t.Error("nb_gifts should be 0")
	}
}
