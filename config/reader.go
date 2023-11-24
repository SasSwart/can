package config

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func ReadConfigs(reader io.Reader) <-chan []byte {
	configs := make(chan []byte)
	scanner := bufio.NewScanner(reader)
	scanner.Split(splitAtDashDashDash)

	go func() {
		defer close(configs)
		for scanner.Scan() {
			if Debug {
				fmt.Println("Reading config")
			}
			configBytes := scanner.Bytes()
			if Debug {
				fmt.Println("Read config:", len(configBytes), "bytes")
			}
			configs <- configBytes
		}
	}()

	return configs
}

// splitAtDashDashDash is used by bufio.Scanner to scan a file until it finds the first occurrence of `---`
func splitAtDashDashDash(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Look for the "---" separator
	if idx := bytes.Index(data, []byte("---")); idx >= 0 {
		// return the data before the separator, and advance the scanner
		// to the start of the next section (after the separator)
		return idx + 3, data[:idx], nil
	}

	if atEOF && len(data) != 0 {
		return len(data), data, nil
	}

	return 0, nil, nil
}
