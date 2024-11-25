package huffman

import (
	"container/heap"
)

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

func buildHuffmanCodes(node *MinNode, currentCode string, codes map[rune]string) {
	if node == nil {
		return
	}
	// basicaly im checking if the current node is leaf if it is leaf node that means im interested in writing its code
	// if it is not it means that is just temp
	if node.left == nil && node.right == nil {
		codes[node.value] = currentCode
		return
	}

	buildHuffmanCodes(node.left, currentCode+"0", codes)
	buildHuffmanCodes(node.right, currentCode+"1", codes)
}

func Encode(input string) string {
	freqTable := calculateFrequancy(input)
	pq := &MinHeap{}

	for key, val := range freqTable {
		heap.Push(pq, &MinNode{
			value:    key,
			priority: val,
		})
	}
	heap.Init(pq)

	var left, right *MinNode
	for pq.Len() > 1 {
		left = heap.Pop(pq).(*MinNode)
		right = heap.Pop(pq).(*MinNode)

		top := &MinNode{
			value:    ' ', // blank char
			priority: left.priority + right.priority,
			left:     left,
			right:    right,
		}

		heap.Push(pq, top)
	}

	root := heap.Pop(pq).(*MinNode)

	codes := make(map[rune]string)
	buildHuffmanCodes(root, "", codes)
	var encoded string
	for _, char := range input {
		encoded += codes[char]
	}

	return encoded
}

func Decode(input string, codes map[rune]string) string {

	currentString := ""
	result := ""
	reversedCodes := map[string]rune{}

	//reversing the keys to values to easily find todo look up there is a better way
	for key, value := range codes {
		reversedCodes[value] = key
	}

	for _, value := range input {
		currentString += string(value)
		if t_value, ok := reversedCodes[currentString]; ok {
			result += string(t_value)
			currentString = ""
		}
	}

	return result
}
