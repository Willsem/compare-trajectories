// Package model ...
package model

import (
	"fmt"
	"strings"
	"time"
)

type ParsedDate struct {
	time.Time
}

const layout = "01/02/06 15:04:05.000000"

func (d *ParsedDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return
	}

	d.Time, err = time.Parse(layout, s)
	return
}

func (d ParsedDate) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return nil, nil
	}

	return []byte(fmt.Sprintf(`"%s"`, d.Time.Format(layout))), nil
}
