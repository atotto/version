package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type VersionFile struct {
	Version string `json:"version"`
}

func Version(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	v := VersionFile{}
	if err := json.NewDecoder(f).Decode(&v); err != nil {
		return "", err
	}
	if v.Version == "" {
		return "", fmt.Errorf("version not found")
	}

	return v.Version, nil
}
