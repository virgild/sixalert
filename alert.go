package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

const TTC_Timezone = "EDT"

type TTCAlert struct {
	Content   string
	Affecting string
	Timestamp time.Time
}

func (alert *TTCAlert) Checksum() string {
	buf := []byte(fmt.Sprintf("%s%s", alert.Content, alert.Timestamp))
	return fmt.Sprintf("%x", md5.Sum(buf))
}

func (alert *TTCAlert) String() string {
	s := fmt.Sprintf("%v: %s", alert.Timestamp, alert.Content)
	return s
}

func NewTTCAlert(content string, affecting string, timestamp time.Time) *TTCAlert {
	a := &TTCAlert{
		content,
		affecting,
		timestamp,
	}
	return a
}
