package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	application := &Application{}

	t.Run("OneChannelCloses", func(t *testing.T) {
		start := time.Now()

		<-application.or(
			signal(2*time.Hour),
			signal(5*time.Minute),
			signal(1*time.Second),
			signal(1*time.Hour),
			signal(1*time.Minute),
		)

		duration := time.Since(start)

		if duration < 1*time.Second || duration > 2*time.Hour {
			t.Errorf("expected duration between 1 second and 2 hours got %v", duration)
		}
	})

	t.Run("AllChannelsClose", func(t *testing.T) {
		start := time.Now()

		<-application.or(
			signal(5*time.Second),
			signal(10*time.Second),
			signal(15*time.Second),
		)

		duration := time.Since(start)

		if duration < 5*time.Second || duration > 15*time.Second {
			t.Errorf("expected duration between 5 seconds and 15 seconds got %v", duration)
		}
	})

	t.Run("NoChannels", func(t *testing.T) {
		start := time.Now()

		select {
		case <-application.or():
			t.Error("expected no channel to close but a channel closed")
		case <-time.After(1 * time.Second):
		}

		duration := time.Since(start)

		if duration < 1*time.Second || duration > 2*time.Second {
			t.Errorf("expected duration between 1 second and 2 seconds got %v", duration)
		}
	})
}

func signal(after time.Duration) <-chan interface{} {
	channel := make(chan interface{})

	go func() {
		defer close(channel)

		time.Sleep(after)
	}()

	return channel
}
