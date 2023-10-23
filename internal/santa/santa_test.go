package santa

import "testing"

func TestSelectGiftedMapLen(t *testing.T) {
	inSantas, err := LoadSantas("../../santas-sample.json")
	if err != nil {
		t.Log("error: could not load config file")
		t.Fail()
	}

	outSantas, err := SelectGifted(inSantas, 1)
	if err != nil {
		t.Log("error during SelectGifted")
		t.Fail()
	}

	if len(outSantas) != len(inSantas) {
		t.Logf("error: should have %d santas, but have %d\n", len(inSantas), len(outSantas))
		t.Fail()
	}
}

func TestSelectGiftedElementsLen(t *testing.T) {
	inSantas, err := LoadSantas("../../santas-sample.json")
	if err != nil {
		t.Log("error: could not load config file")
		t.Fail()
	}

	outSantas, err := SelectGifted(inSantas, 1)
	if err != nil {
		t.Log("error during SelectGifted")
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

func TestSelectGiftedReceiversHaveSantasNames(t *testing.T) {
	inSantas, err := LoadSantas("../../santas-sample.json")
	if err != nil {
		t.Log("error: could not load config file")
		t.Fail()
	}

	outSantas, err := SelectGifted(inSantas, 2)
	if err != nil {
		t.Log("error during SelectGifted")
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

func TestSelectGiftedDontGiveToYourself(t *testing.T) {
	inSantas, err := LoadSantas("../../santas-sample.json")
	if err != nil {
		t.Log("error: could not load config file")
		t.Fail()
	}

	outSantas, err := SelectGifted(inSantas, 2)
	if err != nil {
		t.Log("error during SelectGifted")
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

func TestSelectGiftedShouldHaveDifferentGifted(t *testing.T) {
	inSantas, err := LoadSantas("../../santas-sample.json")
	if err != nil {
		t.Log("error: could not load config file")
		t.Fail()
	}

	outSantas, err := SelectGifted(inSantas, 3)
	if err != nil {
		t.Log("error during SelectGifted")
		t.Fail()
	}

	for gifter, gifted := range outSantas {
		if gifted[0] == gifted[1] || gifted[1] == gifted[2] || gifted[0] == gifted[2] {
			t.Logf("error: several times the same gifted: %s -> %v\n", gifter, gifted)
			t.Fail()
			return
		}
	}
}

// TODO : gifter and gifted should be from different clans
