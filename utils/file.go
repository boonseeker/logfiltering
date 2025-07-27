package utils

import (
	"encoding/json"
	"io/ioutil"
	"squeezlogs/parser"
	"os"
)

func ReadLogs(path string) ([]parser.LogEntry, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var logs []parser.LogEntry
	err = json.Unmarshal(data, &logs)
	return logs, err
}

func WriteLogs(path string, logs []parser.FilteredLog) error {
	data, err := json.MarshalIndent(logs, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}