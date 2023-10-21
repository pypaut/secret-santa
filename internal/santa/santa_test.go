package santa

import "testing"

func TestSelectSantasMapLen(t *testing.T) {
	inSantas, err := LoadSantas("../../santas-sample.json")
	if err != nil {
		t.Log("error: could not load config file")
		t.Fail()
	}

	outSantas, err := SelectSantas(inSantas, 1)
	if err != nil {
		t.Log("error during SelectSantas")
		t.Fail()
	}

	if len(outSantas) != len(inSantas) {
		t.Logf("error: should have %d santas, but have %d\n", len(inSantas), len(outSantas))
		t.Fail()
	}
}

func TestSelectSantasElementsLen(t *testing.T) {
	inSantas, err := LoadSantas("../../santas-sample.json")
	if err != nil {
		t.Log("error: could not load config file")
		t.Fail()
	}

	outSantas, err := SelectSantas(inSantas, 1)
	if err != nil {
		t.Log("error during SelectSantas")
		t.Fail()
	}

	for name, _ := range outSantas {
		actualLen := len(outSantas[name])
		if actualLen != 1 {
			t.Logf("error: should have 1 santa for each santa, but have %d\n", actualLen)
			t.Fail()
		}
	}
}

func TestSelectSantasReceiversHaveSantasNames(t *testing.T) {
	inSantas, err := LoadSantas("../../santas-sample.json")
	if err != nil {
		t.Log("error: could not load config file")
		t.Fail()
	}

	outSantas, err := SelectSantas(inSantas, 2)
	if err != nil {
		t.Log("error during SelectSantas")
		t.Fail()
	}

	for _, santaReceivers := range outSantas {
		for _, receiverName := range santaReceivers {
			receiverIsSanta := false
			for _, giver := range inSantas {
				if giver.Name == receiverName {
					receiverIsSanta = true
					break
				}
			}

			if !receiverIsSanta {
				t.Logf("error: receiver santa is not a santa: \"%s\"\n", receiverName)
				t.Fail()
				return
			}
		}
	}
}

func TestSelectSantasDontGiveToYourself(t *testing.T) {
	inSantas, err := LoadSantas("../../santas-sample.json")
	if err != nil {
		t.Log("error: could not load config file")
		t.Fail()
	}

	outSantas, err := SelectSantas(inSantas, 2)
	if err != nil {
		t.Log("error during SelectSantas")
		t.Fail()
	}

	for santaGiver, santaReceivers := range outSantas {
		for _, receiverName := range santaReceivers {
			if receiverName == santaGiver {
				t.Logf("error: santa \"%s\" gives to himself", santaGiver)
				t.Fail()
				return
			}
		}
	}
}

// TODO : should have different gifted
// TODO : gifter and gifted should be from different clans
