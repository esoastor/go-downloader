package downloader

import (
	"net/url"
	"strings"
)

type Downloadable interface {
	GetName() string
	GetUrl() string
}

type FileUrl string

func (fUrl FileUrl) GetName() string {
	parsed, err := url.Parse(string(fUrl))
    if err != nil {
        panic(err)
    }
	path := strings.Split(parsed.Path, "/")
	return  path[len(path)-1]
}

func (fUrl FileUrl) GetUrl() string {
	return string(fUrl)
}


type File struct {
	Name string
	Url string
}

func (file File) GetName() string {
	return  file.Name
}

func (file File) GetUrl() string {
	return file.Url
}

type Dir struct {
	Name string
	Files []Downloadable
}

