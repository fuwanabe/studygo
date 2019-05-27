package main

import "fmt"

func main() {
	fmt.Println(RemoveChar("ABCDEFG"))
}

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
	//wordArray := strings.Split(word, "")
	//wordArray[0] = ""
	//wordArray[len(wordArray)-1] = ""
	//return strings.Join(wordArray, "")

}
