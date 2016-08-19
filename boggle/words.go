package boggle

type WordList struct {
	Words map[string]bool
}

// type WordTree struct {
//	Letter byte
//	Children [26]*WordTree
// }

// func (*WordTree tree) AddWord(word string) error {

// }

func (list *WordList) AddWord(word string) {
	list.Words[word] = true
}

func (list *WordList) HasWord(word string) bool {
	return list.Words[word] == true
}
