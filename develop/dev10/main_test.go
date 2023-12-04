package main

import (
	"bytes"
	"io"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		t.Fatalf("Error creating fake server: %v", err)
	}

	defer listener.Close()

	go func() {
		connection, _ := listener.Accept()

		defer connection.Close()

		connection.Write([]byte("hello from server"))
	}()

	stdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	defer func() {
		os.Stdout = stdout
	}()

	flags := &Flags{Timeout: 3 * time.Second}
	application := &Application{}

	go func() {
		time.Sleep(1 * time.Second)

		w.Write([]byte("hello from client"))
	}()

	done := make(chan struct{})

	go func() {
		application.Connection(listener.Addr().String(), flags)

		close(done)
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("test timed out")
	}

	w.Close()

	var buffer bytes.Buffer

	io.Copy(&buffer, r)

	output := buffer.String()

	expected := "hello from server"

	if !strings.Contains(output, expected) {
		t.Errorf("expected program output to contain %s got %s", expected, output)
	}
}
