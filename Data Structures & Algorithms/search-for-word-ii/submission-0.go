type TrieNode struct {
    Children map[byte]*TrieNode
    Word bool
}

type Trie struct {
    Root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{
        Root: &TrieNode{
            Children: make(map[byte]*TrieNode),
        },
    }
}

func (t *Trie) Insert(word string) {
    if t.Root == nil {
        return
    }
    curr := t.Root

    for ix := range word {
        c := word[ix]
        if _, ok := curr.Children[c]; !ok {
            curr.Children[c] = &TrieNode{
                Children: make(map[byte]*TrieNode),
            }
        } 
        curr = curr.Children[c]
    }
    curr.Word = true
}

func (t *Trie) StartWith(prefix string) bool {
    curr := t.Root

    for ix := range prefix {
        c := prefix[ix]
        if _, ok := curr.Children[c]; !ok {
            return false
        }
    }
    return true
}

func findWords(board [][]byte, words []string) []string {
    trie := NewTrie()
    for _, word := range words {
        trie.Insert(word)
    }

    type pair struct {
        x int
        y int
    }
    rows, cols := len(board), len(board[0])
    result := make(map[string]bool)
    visit := make(map[pair]bool, 0)

    var dfs func(r, c int, node *TrieNode, word string)
    dfs = func(r, c int, node *TrieNode, word string) {
        if r < 0 || c < 0 || r == rows || c == cols ||
        visit[pair{r,c}] || node.Children[board[r][c]] == nil {
            return
        }
        visit[pair{r,c}] = true

        node = node.Children[board[r][c]]
        word += string(board[r][c])
        if node.Word {
            result[word] = true
        }

        dfs(r-1,c,node,word)
        dfs(r+1,c,node,word)
        dfs(r,c-1,node,word)
        dfs(r,c+1,node,word)

        delete(visit, pair{r,c})
    }

    for ix := range board {
        for iy := range board[ix] {
            dfs(ix,iy,trie.Root, "")
        }
    }
    response := []string{}
    for word := range result {
        response = append(response, word)
    }
    return response
}

