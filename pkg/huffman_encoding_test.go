package huffman

import (
	"reflect"
	"testing"
)

func TestEncodeString(t *testing.T) {

	input := "fsddsdaaaa"

	expectedOutput := "1101111010111100000"

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
