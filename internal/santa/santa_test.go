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
			t.Logf("error: should have 1 santa, but have %d\n", actualLen)
			t.Fail()
		}
	}
}
