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
	onDownloadSuccessCallback func(d Downloadable)
	onDownloadErrorCallback func(d Downloadable, err error)
	onFileRecieveCallback func(file *[]byte)
}  

func GetDownloader() Downloader {
	badReqCall := func(resp *http.Response) {
		log.Printf("StatusCode: %v", resp.StatusCode)
		panic("Bad response")
	}
	dwnSuccessCall := func(d Downloadable) {
		log.Printf("OK %v", d.GetUrl())
	}
	dwnErrorCall := func(d Downloadable, err error) {
		log.Printf("ERROR: %v, %v", err, d.GetUrl())
		panic("")
	}
	downloader := Downloader{
		onBadResponseCallback: badReqCall,
		onDownloadSuccessCallback: dwnSuccessCall,
		onDownloadErrorCallback: dwnErrorCall,
	}
	return downloader	
}

func (dwn *Downloader)OnBadResponse(callback func(resp *http.Response)) {
	dwn.onBadResponseCallback = callback
}

func (dwn *Downloader)OnDownloadSuccess(callback func(d Downloadable)) {
	dwn.onDownloadSuccessCallback = callback
}

func (dwn *Downloader)OnDownloadError(callback func(d Downloadable, err error)) {
	dwn.onDownloadErrorCallback = callback
}

func (dwn *Downloader)OnFileRecieve(callback func(file *[]byte)) {
	dwn.onFileRecieveCallback = callback
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
	
	dwn.onFileRecieveCallback(&body)

	error := utils.WriteContentToFile(body, dir + "/" + d.GetName())

	if (error != nil) {
		dwn.onDownloadErrorCallback(d, error)
	} else {
		dwn.onDownloadSuccessCallback(d)
	}
}

func (dwn *Downloader)downloadWithWait(d Downloadable, dir string, wg *sync.WaitGroup) {
	defer wg.Done()
	dwn.Download(d, dir)
}
