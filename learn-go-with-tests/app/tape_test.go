package app

import (
	"io"
	"io/ioutil"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	newText := "abc"

	tape := &tape{file}

	tape.Write([]byte(newText))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := newText

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
