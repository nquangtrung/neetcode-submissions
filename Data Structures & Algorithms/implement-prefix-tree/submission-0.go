type PrefixTree struct {
	children map[byte]*PrefixTree
	fullWord bool
}

func Constructor() PrefixTree {
	return *childConstructor(false) 
}

func childConstructor(fullWord bool) *PrefixTree {
	return &PrefixTree{
		children: map[byte]*PrefixTree{},
		fullWord: fullWord,
	} 
}

func (this *PrefixTree) insertWithPosition(word string, position int) {
	if _, ok := this.children[word[position]]; !ok {
		// fmt.Printf("create child for %s:%d\n", word, position)
		this.children[word[position]] = childConstructor(position == len(word) - 1) 
	}
	if position == len(word) - 1 {
		this.children[word[position]].fullWord = true
		return
	}
	this.children[word[position]].insertWithPosition(word, position + 1)
}

func (this *PrefixTree) Insert(word string) {
	this.insertWithPosition(word, 0)
}

func (this *PrefixTree) searchWithPosition(word string, position int) bool {
	// fmt.Printf("Search %s at %d\n", word, position)
	child, ok := this.children[word[position]];
	if (!ok) {
		return false
	}

	if position == len(word) - 1 {
		return child.fullWord
	}

	return child.searchWithPosition(word, position + 1)
}

func (this *PrefixTree) Search(word string) bool {
	return this.searchWithPosition(word, 0)
}

func (this *PrefixTree) startsWithPosition(word string, position int) bool {
	child, ok := this.children[word[position]];
	if (!ok) {
		return false
	}

	if position == len(word) - 1 {
		return true
	}

	return child.startsWithPosition(word, position + 1)
}


func (this *PrefixTree) StartsWith(prefix string) bool {
	return this.startsWithPosition(prefix, 0)
}
