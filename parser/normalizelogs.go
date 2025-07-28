package parser

type NormalizedLog struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Level     string `json:"level"` // optional, fallback = "info"
	Raw       string `json:"raw"`   // original unstructured message (optional)
}


func NormalizeLog(input string) NormalizedLog {
	var temp map[string]interface{}
	err := json.Unmarshal([]byte(input), &temp)
	if err == nil {
		// input is JSON
		norm := NormalizedLog{
			Timestamp: getStringOrDefault(temp["timestamp"], time.Now().UTC().Format(time.RFC3339)),
			Message:   getStringOrDefault(temp["message"], ""),
			Level:     getStringOrDefault(temp["level"], "info"),
			Raw:       input,
		}
		return norm
	}
	// Plain text fallback
	return NormalizedLog{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Message:   input,
		Level:     "info",
		Raw:       input,
	}
}

func getStringOrDefault(val interface{}, def string) string {
	if str, ok := val.(string); ok {
		return str
	}
	return def
}
