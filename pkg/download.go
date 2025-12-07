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

func downloadWithWait(d Downloadable, dir string, wg *sync.WaitGroup) {
	defer wg.Done()
	Download(d, dir)
}

func Download(d Downloadable, dir string) {
	body := utils.MakeGetRequest(d.GetUrl())
	
	utils.WriteContentToFile(body, dir + "/" + d.GetName())
}

