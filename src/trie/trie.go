package trie

// import "fmt"
import "sort"

type Word struct {
    Value string
    Definition string
}

type Vertex struct {
    wordCount uint8
    prefixCount uint8
    definition string
    // Edges [255]*Vertex
    Edges map[uint8]*Vertex
}

type ByWord []*Word

func (s ByWord) Len() int {
    return len(s)
}
func (s ByWord) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByWord) Less(i, j int) bool {
    return s[i].Value < s[j].Value
}

func NewTrie() *Vertex {
    index := &Vertex{ Edges: make(map[uint8]*Vertex) }
    return index
}

func BuildIndex(words []*Word) *Vertex {
    index := NewTrie()
    for _, word := range words {
        index.AddWord(word)
        // fmt.Println("")
    }

    return index
}

func (vertex *Vertex) AddWord(word *Word) {
    if word.Value == "" {
        vertex.wordCount++
        vertex.definition = word.Definition
    } else {
        vertex.prefixCount++
        firstByte := word.Value[0]
        if vertex.Edges[firstByte] == nil {
            vertex.Edges[firstByte] = NewTrie()
        }
        // fmt.Printf("%v ", firstByte)
        word.Value = word.Value[1:len(word.Value)]
        vertex.Edges[firstByte].AddWord(word)
    }
}

func (vertex *Vertex) GetWordsOfPrefix (prefix string) []*Word {
    // fmt.Println("Getting words of prefix: " + prefix)
    suffixes := getSuffixesOfPrefix(vertex, prefix)
    // words := []*Word{}
    for _, suffix := range suffixes {
        suffix.Value = prefix + suffix.Value
        // words = append(words, prefix + suffix)
    }
    sort.Sort(ByWord(suffixes))
    return suffixes
}

func getSuffixesOfPrefix(vertex *Vertex, prefix string) []*Word {
    // fmt.Println("Getting suffixes of prefix: " + prefix)
    if prefix == "" {
        // fmt.Println("");
        suffixes := getSuffixesOfVertex(vertex, "", []*Word{})
        return suffixes
    } else {
        firstByte := prefix[0]
        // fmt.Printf("%v ", firstByte);
        if vertex.Edges[firstByte] == nil {
            return []*Word{}
        }
        return getSuffixesOfPrefix(vertex.Edges[firstByte], prefix[1:len(prefix)])
    }
}

func getSuffixesOfVertex(vertex *Vertex, word string, suffixes []*Word) []*Word{
    // fmt.Println(vertex.wordCount)
    // fmt.Println(vertex.prefixCount)
    // fmt.Println(word)
    if vertex.wordCount > 0 {
        suffixes = append(suffixes, &Word{
            Value: word,
            Definition: vertex.definition,
        })
    }
    for letter, v := range vertex.Edges {
        if v != nil {
            suffixes = getSuffixesOfVertex(v, string(append([]byte(word), byte(letter))), suffixes)
        }
    }
    return suffixes
}