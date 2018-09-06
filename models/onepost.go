package models

import "time"

// TPost one post information.
type TPost struct {
	FileName   string
	Time       time.Time
	Title      string
	Author     string
	Categories string
	Tags       []string
	Content    []byte
}
