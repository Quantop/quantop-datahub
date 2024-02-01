package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readSecretFromFile(filePath, secretKey string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")

		if len(parts) == 2 && strings.TrimSpace(parts[0]) == secretKey {
			return strings.TrimSpace(parts[1]), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("secret key %s not found in file", secretKey)
}
