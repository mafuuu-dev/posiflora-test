package tgl_enum

type Type string
type DBType string

const (
	Sent    DBType = "SENT"
	Failed  DBType = "FAILED"
	Error   Type   = "ERROR"
	Skipped Type   = "SKIPPED"
)
