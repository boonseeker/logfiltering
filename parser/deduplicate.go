// File: parser/deduplicator.go
package parser

// import (
// 	"time"
// )

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Level	 string `json:"level"`
}

type FilteredLog struct {
	Message    string `json:"message"`
	Count      int    `json:"count"`
	Level      string `json:"level,omitempty"` // Optional, only if present in the original log
	FirstSeen  string `json:"first_seen"`
	LastSeen   string `json:"last_seen"`
}

func Deduplicate(logs []LogEntry) []FilteredLog {
	logMap := make(map[string]*FilteredLog)
	resp := dropInfoLogs(logs)

	for _, entry := range resp {
		ts := entry.Timestamp
		if logMap[entry.Message] == nil {
			if entry.Level != "" {
			logMap[entry.Message] = &FilteredLog{
				Message:   entry.Message,
				Count:     1,
				Level:    entry.Level,
				FirstSeen: ts,
				LastSeen:  ts,
			}
		}else{
			logMap[entry.Message] = &FilteredLog{
				Message:   entry.Message,
				Count:     1,
				FirstSeen: ts,
				LastSeen:  ts,
		}
		}
		} else {
			logMap[entry.Message].Count++
			logMap[entry.Message].LastSeen = ts
		}
	}
	result := make([]FilteredLog, 0, len(logMap))
	for _, log := range logMap {
		result = append(result, *log)
	}
	return result
}

func dropInfoLogs(logs []LogEntry) []LogEntry {
	filtered := make([]LogEntry, 0, len(logs))
	for _, log := range logs {
		if log.Level != "info" {
			filtered = append(filtered, log)
		}
	}
	return filtered
}