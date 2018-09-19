package types

type Request struct {
	Message    string                 `json:"message"`
	SourceFile string                 `json:"src_file"`
	SourceLine string                 `json:"src_line"`
	Context    map[string]interface{} `json:"context"`
	Level      string                 `json:"level"`
	Time       string                 `json:"time"`
	Backtrace  string                 `json:"backtrace"`
}
