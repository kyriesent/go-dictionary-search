package main

import "bufio"
import "os"
import "fmt"
import "encoding/xml"

import "trie"


func main() {
    index := buildIndexFromXMLFile("/home/ubuntu/dev/go-dictionary-search/files/strongsgreek.xml")

    if index == nil {
        return
    }

    fmt.Print("Enter text: ")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    bytes := scanner.Bytes()
    text := string(bytes)

    fmt.Println("You entered: " + text)
    results := index.GetWordsOfPrefix(text)
    fmt.Print("Results: ")
    for _, w := range results {
        fmt.Printf("%v: ", w.Value)
        fmt.Printf("%v\n", w.Definition)
    }
}

func buildIndexFromXMLFile (filename string) *trie.Vertex {
    fmt.Println("Building index from XML file:", filename)

    file, _ := os.Open(filename)
    decoder := xml.NewDecoder(file)

    dictionary := StrongsDictionary{
        Entries: []StrongsEntry{},
    }

    err := decoder.Decode(&dictionary)
    if err != nil {
        fmt.Printf("error: %v", err)
        return nil
    }

    index := trie.NewTrie()

    for _, v := range dictionary.Entries {
        index.AddWord(&trie.Word{
            Value: v.Greek.Unicode,
            Definition: v.StrongsDef,
        })
    }
    return index
}

type StrongsDictionary struct {
    XMLName xml.Name `xml:"strongsdictionary"`
    Entries []StrongsEntry `xml:"entries>entry"`
}

type StrongsEntry struct {
    Strongs string `xml:"strongs"`
    Greek Greek `xml:"greek"`
    StrongsDef string `xml:"strongs_def"`
}

type Greek struct {
    Unicode string `xml:"unicode,attr"`
    Translit string `xml:"translit,attr"`
}