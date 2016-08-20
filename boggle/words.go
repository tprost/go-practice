package boggle

import "os"
import "bufio"

type WordList struct {
	Words map[string]bool
}

// type WordTree struct {
//	Letter byte
//	Children [26]*WordTree
// }

func NewWordList() *WordList {
	wordList := new(WordList)
	wordList.Words = make(map[string]bool)
	return wordList
}

func ReadWordList(file *os.File) (*WordList, error) {
	wordList := NewWordList()
	reader := bufio.NewReader(file)
	line, _, error := reader.ReadLine()
	for error == nil {
		wordList.AddWord(NewWord(string(line)))
		line, _, error = reader.ReadLine()
	}
	return wordList, nil
}

func (list *WordList) AddWord(word Word) {
	list.Words[word.String()] = true
}

func (list *WordList) HasWord(word Word) bool {
	return list.Words[word.String()] == true
}

func (list *WordList) Size() int {
	return len(list.Words)
}
