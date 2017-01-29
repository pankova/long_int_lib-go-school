package longint

import (
	"testing"
)

func TestSumOfEqualLengthNumbers (t *testing.T) {
	// Max length oh sign int 64 is 19. Give example of 20 digits
	val1 := NewLongIntFromString("92233720368547758071")
	val2 := NewLongIntFromString("54535454324435454524")
	res := Sum(val1, val2)
	if res.String() != "146769174692983212595" {
		t.Errorf("Res must be 146769174692983212595, not %s", res)
	}
}

func TestSumOfDifferentLengthNumbers (t *testing.T) {
	val1 := NewLongIntFromString("92233720368547758071")
	val2 := NewLongIntFromString("42884999423")
	res := Sum(val1, val2)
	if res.String() != "92233720411432757494" {
		t.Errorf("Res must be 92233720411432757494, not %s", res)
	}
}
func TestSumOfnumberAndZero (t *testing.T) {
	val1 := NewLongIntFromString("12345")
	val2 := NewLongIntFromString("0")
	res := Sum(val1, val2)
	if res.String() != "12345" {
		t.Errorf("Res must be 12345, not %s", res)
	}
}
func TestSumOfZeroAndNumber(t *testing.T) {
	val1 := NewLongIntFromString("0")
	val2 := NewLongIntFromString("12345")
	res := Sum(val1, val2)
	if res.String() != "12345" {
		t.Errorf("Res must be 12345, not %s", res)
	}
}
func TestSumOfZeroAndZero(t *testing.T) {
	val1 := NewLongIntFromString("0")
	val2 := NewLongIntFromString("0")
	res := Sum(val1, val2)
	if res.String() != "0" {
		t.Errorf("Res must be 0, not %s", res)
	}
}
func TestSubOfEqualLengthNumbers(t *testing.T) {
	val1 := NewLongIntFromString("92233720368547758071")
	val2 := NewLongIntFromString("54535454324435454524")
	res := Sub(val1, val2)
	if res.String() != "37698266044112303547" {
		t.Errorf("Res must be 37698266044112303547, not %s", res)
	}
}
func TestSubFirstNumLongerThenSecond(t *testing.T) {
	val1 := NewLongIntFromString("92233720368547758071")
	val2 := NewLongIntFromString("42524524")
	res := Sub(val1, val2)
	if res.String() != "92233720368505233547" {
		t.Errorf("Res must be 92233720368505233547, not %s", res)
	}
}

func TestSubSecondNumLongerThenFirst(t *testing.T) {
	val1 := NewLongIntFromString("12345")
	val2 := NewLongIntFromString("92233720368547758071")
	res := Sub(val1, val2)
	if res.String() != "-92233720368547745726" {
		t.Errorf("Res must be -92233720368547745726, not %s", res)
	}
}

func TestSubNumberFromZero (t *testing.T) {
	val1 := NewLongIntFromString("0")
	val2 := NewLongIntFromString("12345")
	res := Sub(val1, val2)
	if res.String() != "-12345" {
		t.Errorf("Res must be -12345, not %s", res)
	}
}

func TestSubZeroFromNumber (t *testing.T) {
	val1 := NewLongIntFromString("12345")
	val2 := NewLongIntFromString("0")
	res := Sub(val1, val2)
	if res.String() != "12345" {
		t.Errorf("Res must be 12345, not %s", res)
	}
}

func TestSubZeroFromZero (t *testing.T) {
	val1 := NewLongIntFromString("0")
	val2 := NewLongIntFromString("0")
	res := Sub(val1, val2)
	if res.String() != "0" {
		t.Errorf("Res must be 0, not %s", res)
	}
}

func TestMulEqualLengthNumber (t *testing.T) {
	val1 := NewLongIntFromString("92233720368547758071")
	val2 := NewLongIntFromString("54535454324435454524")
	res := Mul(val1, val2)
	if res.String() != "5030007844331688297293457917566074463204" {
		t.Errorf("Res must be 5030007844331688297293457917566074463204, not %s", res)
	}
}

func TestMulDifferentLengthNumber (t *testing.T) {
	val1 := NewLongIntFromString("92233720368547758071")
	val2 := NewLongIntFromString("12345")
	res := Mul(val1, val2)
	if res.String() != "1138625277949722073386495" {
		t.Errorf("Res must be 1138625277949722073386495, not %s", res)
	}
}

func TestMulNumOnZero (t *testing.T) {
	val1 := NewLongIntFromString("12345")
	val2 := NewLongIntFromString("0")
	res := Mul(val1, val2)
	if res.String() != "0" {
		t.Errorf("Res must be 1556, not %s", res)
	}
}

func TestMulZeroOnNum (t *testing.T) {
	val1 := NewLongIntFromString("0")
	val2 := NewLongIntFromString("12345")
	res := Mul(val1, val2)
	if res.String() != "0" {
		t.Errorf("Res must be 0, not %s", res)
	}
}

func TestMulZeroOnZero (t *testing.T) {
	val1 := NewLongIntFromString("0")
	val2 := NewLongIntFromString("0")
	res := Mul(val1, val2)
	if res.String() != "0" {
		t.Errorf("Res must be 0, not %s", res)
	}
}


