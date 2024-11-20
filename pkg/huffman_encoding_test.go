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

func TestDecodeString(t *testing.T) {
	input := "00010111011100011011100101100101110111001010111110101100100010111011100010111010100111111001100100111010001100001001110111100011101111011001"
	table := map[rune]string{
		'a': "000",
		'v': "001",
		'c': "010",
		'x': "0110",
		'b': "0111",
		'n': "10000",
		'h': "10001",
		'g': "1001",
		's': "101",
		'd': "110",
		'f': "111",
	}
	expectedOutput := "asdfadfvxcfbvcfdsgasdfasdsvfdxcbcvngdfhdfsg"
	actualOutput := Decode(input, table)

	if expectedOutput != actualOutput {
		t.Fatalf(`Decode("%s",{a : 1 , s = 01 , d = 1}) Expected Output = %s Actual Output = %s`, input, expectedOutput, actualOutput)
	}
}
