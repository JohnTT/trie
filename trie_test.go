package trie

import (
	"bufio"
	"flag"
	"log"
	"os"
	"testing"
)

var path string

func init() {
	flag.StringVar(&path, "path", "", "Path to file with words")
}

func TestTrieASCII(t *testing.T) {
	var got bool
	var want bool

	obj := Constructor()

	obj.Insert("apple")

	got = obj.Search("apple")
	want = true
	if got != want {
		t.Errorf("obj.Search(\"apple\") = %v, want %v", got, want)
	}

	got = obj.Search("app")
	want = false
	if got != want {
		t.Errorf("obj.Search(\"app\") = %v, want %v", got, want)
	}

	got = obj.StartsWith("app")
	want = true
	if got != want {
		t.Errorf("obj.StartsWith(\"app\") = %v, want %v", got, want)
	}

	obj.Insert("app")

	got = obj.Search("app")
	want = true
	if got != want {
		t.Errorf("obj.Search(\"app\") = %v, want %v", got, want)
	}
}

func TestTrieUnicode(t *testing.T) {
	var got bool
	var want bool

	obj := Constructor()

	obj.Insert("广东省2022")

	got = obj.Search("广东省2022")
	want = true
	if got != want {
		t.Errorf("obj.Search(\"广东省2022\") = %v, want %v", got, want)
	}

	got = obj.Search("广东省")
	want = false
	if got != want {
		t.Errorf("obj.Search(\"广东省\") = %v, want %v", got, want)
	}

	obj.Insert("广东省")

	got = obj.Search("广东省")
	want = true
	if got != want {
		t.Errorf("obj.Search(\"广东省\") = %v, want %v", got, want)
	}

	got = obj.Search("广州市")
	want = false
	if got != want {
		t.Errorf("obj.Search(\"广州市\") = %v, want %v", got, want)
	}

	got = obj.StartsWith("广")
	want = true
	if got != want {
		t.Errorf("obj.StartsWith(\"广\") = %v, want %v", got, want)
	}

	obj.Insert("پنجاب")

	got = obj.Search("پنجاب")
	want = true
	if got != want {
		t.Errorf("obj.Search(\"پنجاب\") = %v, want %v", got, want)
	}

	obj.Insert("बैंगलोर")

	got = obj.Search("बैंगलोर")
	want = true
	if got != want {
		t.Errorf("obj.Search(\"बैंगलोर\") = %v, want %v", got, want)
	}
}

func readWords() {
	flag.Parse()
	if path == "" {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	obj := Make()
	for scanner.Scan() {
		obj.Insert(scanner.Text())
		obj.Search(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func BenchmarkTrie(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readWords()
	}
}
