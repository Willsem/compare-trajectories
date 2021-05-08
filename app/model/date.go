package model

import (
	"strings"
	"time"
)

type ParsedDate struct {
	time.Time
}

func (c *ParsedDate) UnmarshalJSON(b []byte) (err error) {
	const layout = "01/02/06 15:04:05.000000"

	s := strings.Trim(string(b), "\"") // remove quotes
	if s == "null" {
		return
	}

	c.Time, err = time.Parse(layout, s)
	return
}
