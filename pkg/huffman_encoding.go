package huffman

func calculateFrequancy(input string) map[rune]int {

	res := map[rune]int{}
	for _, value := range input {
		if _, ok := res[value]; ok {
			res[value] += 1
		} else {
			res[value] = 1
		}
	}

	return res
}

func Encode(input string) string {

	freqTable := calculateFrequancy(input)

	return "asdasd"

}

func Decode(input string) {}
