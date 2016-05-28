package trie

import "testing"
import "fmt"

func TestTrie(t *testing.T) {
  words := []*Word{
    &Word{Value: "amy", Definition: "my wife"},
    &Word{Value: "ames", Definition: "also my wife"},
    &Word{Value: "ami", Definition: "my wife in French"},
    &Word{Value: "a", Definition: "a letter"},
    &Word{Value: "α", Definition: "a Greek letter"},
    &Word{Value: "καὶ", Definition: "a Greek word"},
  }
  index := BuildIndex(words)

  var results []*Word
  results = index.GetWordsOfPrefix("a")
  fmt.Printf("%v\n", results)

  if results[1].Definition != "also my wife" {
    t.Error("Expected 'also my wife', got ", results[1].Definition)
  }

  if len(results) != 4 {
    t.Error("Expected 4, got ", len(results))
  }

  results = index.GetWordsOfPrefix("κ")
  fmt.Printf("%v\n", results)
  if len(results) != 1 {
    t.Error("Expected 1, got ", len(results))
  }
}
