package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // '-' directive to hide
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitzero"`
	Runtime   int32     `json:"runtime,omitzero"`
	Genres    []string  `json:"genres,omitzero"`
	Version   int32     `json:"version"`
}

// `omitzero directive to hide the Year, Runtime and Genres fields in the output if and only if they are the relevant zero value for their type.`
