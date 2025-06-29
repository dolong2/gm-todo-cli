package model

import (
	"os"
	"path/filepath"
)

func getContextFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	cliDir := filepath.Join(configDir, "gm-todo")
	if err := os.MkdirAll(cliDir, 0755); err != nil { // 0755는 권한 설정 (읽기/쓰기/실행)
		return "", err
	}

	contextPath := filepath.Join(cliDir, "contexts.json")
	return contextPath, nil
}
