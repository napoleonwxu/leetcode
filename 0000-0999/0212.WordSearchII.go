func findWords(board [][]byte, words []string) []string {
    ans := []string{}
    root := buildTrie(words)
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            dfs(board, i, j, root, &ans)
        }
    }
    return ans
}

func dfs(board [][]byte, i, j int, node *TrieNode, ans *[]string) {
    if node.word != "" {
        *ans = append(*ans, node.word)
        node.word = ""
    }
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == '#' || node.child[board[i][j]-'a'] == nil {
        return
    }
    tmp := board[i][j]
    board[i][j] = '#'
    node = node.child[tmp-'a']
    dfs(board, i-1, j, node, ans)
    dfs(board, i+1, j, node, ans)
    dfs(board, i, j+1, node, ans)
    dfs(board, i, j-1, node, ans)
    board[i][j] = tmp
}

type TrieNode struct {
    child [26]*TrieNode
    word string
}

func buildTrie(words []string) *TrieNode {
    root := &TrieNode{}
    for _, w := range words {
        node := root
        for _, ch := range w {
            if node.child[ch-'a'] == nil {
                node.child[ch-'a'] = &TrieNode{}
            }
            node = node.child[ch-'a']
        }
        node.word = w
    }
    return root
}