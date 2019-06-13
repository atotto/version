package file

import (
	"path/filepath"

	"github.com/atotto/version/internal/file/json"
	"github.com/atotto/version/internal/file/text"
)

func Version(fpath string) (string, error) {
	ext := filepath.Ext(fpath)
	switch ext {
	case ".json":
		return json.Version(fpath)
	default:
		return text.Version(fpath)
	}
}
