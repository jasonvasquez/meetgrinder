package types

import (
	"bytes"
	"time"
)

type MGDate struct {
	time.Time
}

const mgdLayout = "2006-01-02"

func (mgd *MGDate) UnmarshalJSON(b []byte) (err error) {
	//strip surrounding quotes off of the value
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	mgd.Time, err = time.Parse(mgdLayout, string(b))
	return
}

func (mgd *MGDate) MarshalJSON() ([]byte, error) {
	val := mgd.Time.Format(mgdLayout)[0:10]
	var b bytes.Buffer
	b.Grow(len(mgdLayout) + 2)

	//surround the date value in quotes
	b.WriteString("\"")
	b.WriteString(val)
	b.WriteString("\"")

	return b.Bytes(), nil
}

var nilTime = (time.Time{}).UnixNano()

func (mgd *MGDate) IsSet() bool {
	return mgd.UnixNano() != nilTime
}
