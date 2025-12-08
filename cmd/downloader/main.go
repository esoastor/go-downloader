package main

import (
	"net/http"
	"log"
	downloader "github.com/esoastor/go-downloader/pkg"
)

func main() {
	dwn := downloader.GetDownloader() 
	
	dwn.OnBadResponse(func(resp *http.Response) {
		log.Printf("Headers: %v", resp.Header)
		log.Printf("STATUS: %v", resp.Status)
		panic("test")
	})
	files := []downloader.Downloadable{
		downloader.FileUrl("https://scans.lastation.us/manga/Tower-Dungeon/0022-002.png"),
	}

	dir := downloader.Dir{
		Name: "subdir",
		Files: files,
	}

	dwn.DownloadDir(dir, "test")

}
