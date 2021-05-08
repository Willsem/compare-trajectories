package model

import (
	"strings"
	"time"
)

type ParsedDate struct {
	time.Time
}

func (d *ParsedDate) UnmarshalJSON(b []byte) (err error) {
	const layout = "01/02/06 15:04:05.000000"

	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return
	}

	d.Time, err = time.Parse(layout, s)
	return
}
