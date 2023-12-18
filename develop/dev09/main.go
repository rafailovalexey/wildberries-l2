package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

/*
	=== Утилита wget ===

	Реализовать утилиту wget с возможностью скачивать сайты целиком

	Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type FlagsInterface interface {
	InitializeFlags()
}

type Flags struct{}

var _ FlagsInterface = (*Flags)(nil)

func (f *Flags) InitializeFlags() {
	flag.Parse()
}

type ApplicationInterface interface {
	Download(site string) error
	DownloadResource(directory string, resource string) error
	Resolve(base *url.URL, resource string) string
}

type Application struct{}

var _ ApplicationInterface = (*Application)(nil)

func main() {
	flags := &Flags{}
	application := &Application{}

	flags.InitializeFlags()

	site := flag.Arg(0)

	if site == "" {
		log.Printf("indicate the url of the download site\n")

		os.Exit(1)
	}

	err := application.Download(site)

	if err != nil {
		log.Printf("site loading error %v\n", err)

		os.Exit(1)
	}
}

func (a *Application) Download(site string) error {
	parsed, err := url.Parse(site)

	if err != nil {
		return err
	}

	response, err := http.Get(site)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error %s", response.Status)
	}

	directory := parsed.Host

	err = os.Mkdir(directory, os.ModePerm)

	if err != nil && !os.IsExist(err) {
		return err
	}

	tokenizer := html.NewTokenizer(response.Body)

	for {
		token := tokenizer.Next()

		switch token {
		case html.ErrorToken:
			if tokenizer.Err().Error() == "EOF" {
				return nil
			}

			return tokenizer.Err()
		case html.StartTagToken, html.SelfClosingTagToken:
			temporary := tokenizer.Token()

			for _, attribute := range temporary.Attr {
				if attribute.Key == "href" || attribute.Key == "src" {
					resource := a.Resolve(parsed, attribute.Val)

					err = a.DownloadResource(directory, resource)

					if err != nil {
						log.Printf("resource loading error %s: %v\n", resource, err)
					}
				}
			}
		}
	}
}

func (a *Application) Resolve(base *url.URL, resource string) string {
	parsed, err := url.Parse(resource)

	if err != nil {
		return resource
	}

	resolved := base.ResolveReference(parsed)

	return resolved.String()
}

func (a *Application) DownloadResource(directory string, resource string) error {
	parsed, err := url.Parse(resource)

	if err != nil {
		return err
	}

	response, err := http.Get(resource)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return log.Errorf("resource loading error %s %s", resource, response.Status)
	}

	filepath := path.Join(directory, path.Base(parsed.Path))

	file, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		return err
	}

	log.Printf("resource %s loaded successfully\n", resource)

	return nil
}
