package huffman

import (
	"reflect"
	"testing"
)

func TestEncodeString(t *testing.T) {

	input := "ssssssssdadasddsfsad"

	expectedOutput := "000000001110111101011110100010111"

	actualOutput := Encode(input)

	if actualOutput != expectedOutput {
		t.Fatalf(`huffman.Encode("%s") Expected Output = %s Actual Output = %s`, input, expectedOutput, actualOutput)
	}
}

func TestDecodeString(t *testing.T) {}

func TestFrequencyCount(t *testing.T) {

	input := "aaadff"

	expectedOutput := map[rune]int{
		'a': 3,
		'd': 1,
		'f': 2,
	}

	actualOutput := calculateFrequancy(input)

	if !reflect.DeepEqual(expectedOutput, actualOutput) {
		t.Fatalf(`huffman.frequencyCount("%s") Expected Output = %q Actual Output = %q`, input, expectedOutput, actualOutput)
	}

}
