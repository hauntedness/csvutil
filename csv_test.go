package csvutil

import (
	"os"
	"testing"
)

type Book struct {
	Name      string `json:"name,omitempty"`
	WordCount int    `json:"word_count,string,omitempty"`
	Author    string `json:"author,omitempty"`
}

func TestCsv(t *testing.T) {
	file, err := os.Open("./testdata/books.csv")
	if err != nil {
		t.Errorf("Want %v, Got %v", "no error", err)
		return
	}
	defer file.Close()
	books, err := ReadCsv[Book](file)
	if err != nil {
		t.Error("error read csv", err)
		return
	}
	if len(books) != 4 {
		t.Errorf("Want value %v, Got value %v", 5, len(books))
		return
	}
	for i := range books {
		if books[i].Author == "" {
			t.Errorf("Want %v, Got %v", "not empty", `""`)
		}
		if books[i].WordCount == 0 {
			t.Errorf("Want %v, Got %v", "not 0", `""`)
		}
		if books[i].Name == "" {
			t.Errorf("Want %v, Got %v", "not empty", `""`)
		}
	}
}
