# Downloader
Downloads a list of URLs.

## Installation

```sh
go get github.com/esoastor/go-downloader/pkg
```

## Public API

| Function | Description |
| --- | --- |
| `Download(d Downloadable, dir string)` | Fetches one resource via `GET` and writes it inside `dir` using `d.GetName()`. |
| `DownloadDir(dir Dir, parent string)` | Downloads every file in `dir.Files` concurrently and stores them under `parent/dir.Name`. |

The package ships with two helpers that already satisfy `Downloadable`:

- `FileUrl` – wraps a raw URL string and derives the filename from the URL path.
- `File` – lets you set both `Name` and `Url` explicitly.

## Usage

```go
package main

import (
	downloader "github.com/esoastor/go-downloader/pkg"
)

func main() {
	files := []downloader.Downloadable{
		downloader.FileUrl("https://example.com/chapter/001.png"),
		downloader.FileUrl("https://example.com/chapter/002.png"),
	}

	dir := downloader.Dir{
		Name:  "chapter-001",
		Files: files,
	}

	// Downloads every page into downloads/chapter-001/
	downloader.DownloadDir(dir, "downloads")

	// Fetch a single cover image into downloads/covers/cover.png
	cover := downloader.FileUrl("https://example.com/chapter/cover.png")
	downloader.Download(cover, "downloads/covers")
}
```
