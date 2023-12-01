package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
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

func main() {
	flag.Parse()

	site := flag.Arg(0)

	if site == "" {
		fmt.Printf("пожалуйста, укажите URL сайта для загрузки\n")

		os.Exit(1)
	}

	err := Download(site)

	if err != nil {
		fmt.Printf("ошибка загрузки сайта: %v\n", err)

		os.Exit(1)
	}
}

func Download(site string) error {
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
		return fmt.Errorf("ошибка: %s", response.Status)
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
			return tokenizer.Err()
		case html.StartTagToken, html.SelfClosingTagToken:
			temporary := tokenizer.Token()

			for _, attribute := range temporary.Attr {
				if attribute.Key == "href" || attribute.Key == "src" {
					resource := Resolve(parsed, attribute.Val)

					err = DownloadResource(directory, resource)

					if err != nil {
						fmt.Printf("ошибка загрузки ресурса %s: %v\n", resource, err)
					}
				}
			}
		}
	}
}

func Resolve(base *url.URL, resource string) string {
	parsed, err := url.Parse(resource)

	if err != nil {
		return resource
	}

	resolved := base.ResolveReference(parsed)

	return resolved.String()
}

func DownloadResource(directory string, resource string) error {
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
		return fmt.Errorf("ошибка при загрузки ресурса %s: %s", resource, response.Status)
	}

	filepath := path.Join(directory, path.Base(parsed.Path))

	if path.Base(parsed.Path) != "/" {
		filepath += ".html"
	}

	file, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		return err
	}

	fmt.Printf("ресурс %s успешно загружен\n", resource)

	return nil
}
