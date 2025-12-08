package downloader

import (
	"net/http"
	"path/filepath"
	"sync"
	"log"
	"github.com/esoastor/go-downloader/internal/utils"
)

type Downloader struct {
	onBadResponseCallback func(resp *http.Response)
}  

func GetDownloader() Downloader {
	downloader := Downloader{onBadResponseCallback: func(resp *http.Response) {
		log.Printf("StatusCode: %v", resp.StatusCode)
		panic("Bad response")
	}}
	return downloader	
}

func (dwn *Downloader)OnBadResponse(callback func(resp *http.Response)) {
	dwn.onBadResponseCallback = callback
}

func (dwn *Downloader)DownloadDir(dir Dir, parentDir string) {
	finalDir := filepath.Join(parentDir, dir.Name)
	var wg sync.WaitGroup
	
	for _, file := range dir.Files {
		wg.Add(1)
		go dwn.downloadWithWait(file, finalDir, &wg)
	}
	wg.Wait()
}

func (dwn *Downloader)Download(d Downloadable, dir string) {
	body := utils.MakeGetRequest(d.GetUrl(), dwn.onBadResponseCallback)
	
	utils.WriteContentToFile(body, dir + "/" + d.GetName())
}

func (dwn *Downloader)downloadWithWait(d Downloadable, dir string, wg *sync.WaitGroup) {
	defer wg.Done()
	dwn.Download(d, dir)
}
