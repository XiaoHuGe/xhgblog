package models

import "time"

type Archive struct {
	ArchiveDate time.Time //month
	Year        int       `json:"year"`
	Month       int       `json:"month"`
	Total       int       `json:"total"`
}
