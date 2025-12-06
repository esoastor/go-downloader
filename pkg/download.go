package downloader

import (
	"github.com/esoastor/go-downloader/internal/utils"
	"path/filepath"
	"sync"
)

func DownloadDir(dir Dir, parentDir string) {
	finalDir := filepath.Join(parentDir, dir.Name)
	
	var wg sync.WaitGroup
	
	for _, file := range dir.Files {
		wg.Add(1)
		go downloadWithWait(file, finalDir, &wg)
	}
	wg.Wait()
}

func downloadWithWait[D Downloadable](d D, dir string, wg *sync.WaitGroup) {
	defer wg.Done()
	Download(d, dir)
}

func Download[D Downloadable](d D, dir string) {
	body := utils.MakeGetRequest(d.GetUrl())
	
	utils.WriteContentToFile(body, dir + "/" + d.GetName())
}

