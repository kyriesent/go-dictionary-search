package main

import "os"
import "fmt"
import "encoding/xml"
import "encoding/json"

import "trie"
import "websocketserver"

import "net/http"

var index *trie.Vertex

func main() {
    index = buildIndexFromXMLFile("/home/ubuntu/dev/go-dictionary-search/files/strongsgreek.xml")

    if index == nil {
        return
    }

    websocketserver.NewHandler("query", queryToResultJson)

    fmt.Println("Listening on port " + "8080")
    err := http.ListenAndServe(":" + "8080", nil)
    if err != nil {
        panic("ListenAndServe failed: " + err.Error())
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

func queryToResultJson (in []byte) []byte {
    text := string(in)

    results := index.GetWordsOfPrefix(text)

    jsonData, _ := json.Marshal(results)

    return jsonData
}
