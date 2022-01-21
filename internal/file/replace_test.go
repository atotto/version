package file_test

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/atotto/version/internal/file"
)

func TestReplace(t *testing.T) {
	dir := t.TempDir()
	b, err := ioutil.ReadFile("testdata/version2")
	if err != nil {
		t.Fatal(err)
	}
	f := filepath.Join(dir, "version2")
	if err := ioutil.WriteFile(f, b, 0644); err != nil {
		t.Fatal(err)
	}
	if err := file.Replace(f, "0.1.123", "1.0.0"); err != nil {
		t.Fatal(err)
	}

	b2, err := ioutil.ReadFile(f)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(b2, []byte("1.0.0\n")) {
		t.Fatalf("unexpected content %s", string(b2))
	}
}
