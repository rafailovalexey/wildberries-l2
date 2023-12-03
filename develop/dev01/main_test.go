package main

import (
	"github.com/beevik/ntp"
	"testing"
	"time"
)

func TestNtpTime(t *testing.T) {
	application := &Application{}

	ntpTime, err := application.GetTimeNtp()

	if err != nil {
		t.Errorf(err.Error())
	}

	mockNtpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		t.Errorf(err.Error())
	}

	if ntpTime.Format(time.RFC822) != mockNtpTime.Format(time.RFC822) {
		t.Errorf("got %s expected %s", ntpTime.Format(time.RFC822), mockNtpTime.Format(time.RFC822))
	}
}
