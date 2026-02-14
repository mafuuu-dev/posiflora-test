package tgl_dto

import "time"

type StatsDTO struct {
	LastSentAt  *time.Time `json:"last_sent_at" db:"last_sent_at"`
	SentCount   int64      `json:"sent_count" db:"sent_count"`
	FailedCount int64      `json:"failed_count" db:"failed_count"`
}
