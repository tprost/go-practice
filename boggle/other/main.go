package main

import "os"
import "bufio"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, error := os.Open("words.txt")
	check(error)
	reader := bufio.NewReader(file)
	line, _, error := reader.ReadLine()
	check(error)
	boggleWords, error := os.Create("boggle-words.txt")
	check(error)
	writer := bufio.NewWriter(boggleWords)
	for error == nil && line != nil {
		word := string(line)
		if len(word) < 5 {
			_, error := writer.WriteString(word + "\n")
			check(error)
		}
		line, _, error = reader.ReadLine()
	}
	writer.Flush()
}
