package main

import "fmt"

type T9 struct {
	root *Node
}

func (t *T9) add(word string) {
	cur_node := t.root
	fmt.Println("adding word: ", word)
	runes := []rune(word)
	var code int
	for i := 0; i < len(runes); i++ {
		char := runes[i]
		code = get_code_from_char(char)
		if cur_node.next[code] == nil {
			cur_node.next[code] = new_node()
		}
		cur_node = cur_node.next[code]
	}
	cur_node.word = true
}
func (t *T9) start_search(sequence string) []string {
	possible := t.root.search(sequence, "")
	return possible
}
func (n *Node) search(sequence string, path string) []string {
	length := len(sequence)
	possible := make([]string, 0)

	if length > 0 {
		index := index_from_key(rune(sequence[0]))
		plussed := false
		next_seq := sequence[1:length]
		if n.next[index] != nil {
			path += string(get_char_from_code(index))
			possible = append(possible, n.next[index].search(next_seq, path)...)
			plussed = true
		}
		if n.next[index+1] != nil {
			if plussed {
				path = path[0 : len(path)-1]
			}
			plussed = true
			path += string(get_char_from_code(index + 1))
			possible = append(possible, n.next[index+1].search(next_seq, path)...)
		}
		if n.next[index+2] != nil {
			if plussed {
				path = path[0 : len(path)-1]
			}
			path += string(get_char_from_code(index + 2))
			possible = append(possible, n.next[index+2].search(next_seq, path)...)
		}
	} else {
		if n.word {
			possible = append(possible, path)
			return possible
		} else {
			return nil
		}
	}
	return possible

}
