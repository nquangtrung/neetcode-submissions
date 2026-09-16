type Node struct {
	children map[rune]Node
	isWord bool
}

type WordDictionary struct {
	children map[rune]Node

}

func Constructor() WordDictionary {
	return WordDictionary {
		children: map[rune]Node{},
	}
}

func (this *WordDictionary) AddWord(word string)  {
	current := this.children
    for i, ch := range word {
		_, ok := current[ch]
		if !ok {
			current[ch] = Node{
				children: map[rune]Node{},
				isWord: i == len(word) - 1,
			}
		}

		current = current[ch].children 
	}
}

func (this *WordDictionary) searchAtPosition(current map[rune]Node, runes []rune, pos int) bool {
	r := runes[pos]

	if r == '.' {
		for _, node := range current {
			if pos == len(runes) - 1 {
				// Last char of the word, so true if this node is a word
				if node.isWord {
					return true
				}
				continue
			}

			ok := this.searchAtPosition(node.children, runes, pos + 1)
			if (ok) {
				return true
			}
		}
		// No path match this word
		return false
	} else {
		_, ok := current[r]
		if (!ok) {
			return false
		}

		if pos == len(runes) - 1 {
			return current[r].isWord
		}

		return this.searchAtPosition(current[r].children, runes, pos + 1)
	}
}

func (this *WordDictionary) Search(word string) bool {
	runes := []rune(word)
   	current := this.children
	return this.searchAtPosition(current, runes, 0)
}
