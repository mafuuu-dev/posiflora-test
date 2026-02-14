package errorsx

type StackFrame struct {
	File    string `json:"file"`
	Line    int    `json:"line"`
	Message string `json:"message"`
	IsHuman bool   `json:"is_human,omitempty"`
}

type Error struct {
	stack      []StackFrame
	humanCode  string
	httpStatus int
}

func (e *Error) Error() string {
	if e.humanCode != "" {
		return e.humanCode
	}
	if len(e.stack) > 0 {
		return e.stack[0].Message
	}
	return "unknown error"
}

func (e *Error) Unwrap() error {
	return nil
}
