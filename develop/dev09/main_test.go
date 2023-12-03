package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"
)

func TestDownload(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<a href=\"/page1\">Page 1</a><img src=\"/image.jpg\">"))
	}))

	defer s.Close()

	application := &Application{}

	err := application.Download(s.URL)

	if err != nil {
		t.Errorf("download failed %v", err)
	}

	host := strings.Split(s.URL, "http://")[1]

	_, err = os.Stat(host)

	if os.IsNotExist(err) {
		t.Errorf("directory not created %v", err)
	}

	files := []string{path.Join(host, "page1", "image.jpg")}

	for _, file := range files {
		_, err = os.Stat(file)

		if os.IsNotExist(err) {
			t.Errorf("file not created %v", err)
		}
	}
}

func TestResolve(t *testing.T) {
	url, _ := url.Parse("https://example.com")
	application := &Application{}

	absolute := "https://example.com/path/to/resource"

	result := application.Resolve(url, absolute)

	if result != absolute {
		t.Errorf("resolve failed for absolute url expected %s got %s", absolute, result)
	}

	relative := "/path/to/resource"
	expected := "https://example.com/path/to/resource"

	result = application.Resolve(url, relative)

	if result != expected {
		t.Errorf("resolve failed for relative url expected %s got %s", expected, result)
	}
}
