package model

import (
	"bufio"
	"encoding/json"
	"google.golang.org/genai"
	"os"
	"path/filepath"
)

func saveContext(context []*genai.Content) error {
	contextToJson, err := json.Marshal(context)
	if err != nil {
		return err
	}

	contextPath, err := getContextFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(contextPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(string(contextToJson))
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func getContext() ([]*genai.Content, error) {
	contextPath, err := getContextFilePath()
	if err != nil {
		return nil, err
	}

	rawContextInfo, err := os.ReadFile(contextPath)
	if os.IsNotExist(err) {
		return []*genai.Content{}, nil
	} else if err != nil {
		return nil, err
	}

	contexts := []*genai.Content{}
	err = json.Unmarshal(rawContextInfo, &contexts)
	if err != nil {
		return nil, err
	}

	return contexts, nil
}

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
