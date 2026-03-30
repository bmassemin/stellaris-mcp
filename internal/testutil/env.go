package testutil

import (
	"bufio"
	"os"
	"strings"
)

// LoadEnv reads a .env file and returns the value for the given key.
// Walks up from startDir to find the .env file.
func LoadEnv(key string) string {
	// Try current dir and up to 4 parents
	for _, prefix := range []string{".", "..", "../..", "../../..", "../../../.."} {
		path := prefix + "/.env"
		if v, ok := readEnvKey(path, key); ok {
			return v
		}
	}
	return ""
}

func readEnvKey(path, key string) (string, bool) {
	f, err := os.Open(path)
	if err != nil {
		return "", false
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || line[0] == '#' {
			continue
		}
		k, v, ok := strings.Cut(line, "=")
		if ok && k == key {
			return v, true
		}
	}
	return "", false
}
