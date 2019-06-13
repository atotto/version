package file_test

import (
	"testing"

	"github.com/atotto/version/internal/file"
)

func TestVersion(t *testing.T) {
	t.Run("text", func(t *testing.T) {
		v, err := file.Version("./testdata/version")
		if err != nil {
			t.Fatal(err)
		}

		if v != "1.2.5" {
			t.Fatalf("want 1.2.5, got %s", v)
		}
	})

	t.Run("json", func(t *testing.T) {
		v, err := file.Version("./testdata/deb.json")
		if err != nil {
			t.Fatal(err)
		}

		if v != "0.1.2" {
			t.Fatalf("want 0.1.2, got %s", v)
		}
	})
}
