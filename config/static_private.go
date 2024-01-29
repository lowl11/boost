package config

import (
	"github.com/lowl11/boost/data/errors"
	"github.com/lowl11/boost/pkg/io/file"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

func read(fileName string) (map[string]string, error) {
	envVariables := make(map[string]string)

	if !file.Exist(fileName) {
		return envVariables, nil
	}

	fileContent, err := file.Read(fileName)
	if err != nil {
		return envVariables, err
	}

	combinations := strings.Split(string(fileContent), "\n")

	for _, combo := range combinations {
		dividedCombo := strings.Split(combo, "=")
		if len(dividedCombo) == 1 {
			continue
		}

		envVariables[dividedCombo[0]] = dividedCombo[1]
	}

	return envVariables, nil
}

func replaceVariables(fileContent []byte, envVariables map[string]string) ([]byte, error) {
	if envVariables == nil {
		return nil, errors.New("env variables map is NULL")
	}

	if fileContent == nil {
		return nil, errors.New("file content is NULL")
	}

	// define variables
	variableRegex, err := regexp.Compile("{{(.*?)}}")
	if err != nil {
		return nil, err
	}

	// convert file content to string once
	fileContentString := string(fileContent)

	// replace variables
	variables := variableRegex.FindAllString(fileContentString, -1)
	for _, envVariable := range variables {
		if cleanVariableValue, isVariable := isVar(envVariable); isVariable {
			if value, ok := envVariables[cleanVariableValue]; ok {
				fileContentString = strings.ReplaceAll(fileContentString, envVariable, value)
			} else {
				if osValue := os.Getenv(cleanVariableValue); osValue != "" {
					fileContentString = strings.ReplaceAll(fileContentString, envVariable, osValue)
				} else {
					fileContentString = strings.ReplaceAll(fileContentString, envVariable, "")
				}
			}
		}
	}

	// convert file content to bytes once
	return []byte(fileContentString), nil
}

func isVar(value string) (string, bool) {
	length := utf8.RuneCountInString(value)
	if length < 4 {
		return "", false
	}

	if value[0] == '{' && value[1] == '{' && value[length-1] == '}' && value[length-2] == '}' {
		return value[2 : length-2], true
	}

	return "", false
}
