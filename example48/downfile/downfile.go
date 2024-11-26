package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/jony-lee/go-progress-bar"
	"io"
	"net/http"
	"os"
	"strings"
)

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	bar := progress.New(int64(wc.Total))
	for i := 0; i < (int(wc.Total)); i++ {
		//time.Sleep(time.Second / 10)
		bar.Done(1)
	}
	bar.Finish()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s complete\n", humanize.Bytes(wc.Total))

}

func main() {
	fmt.Println("Download Started")

	fileUrl := "http://topgoer.com/static/2/9.png"
	err := DownloadFile("12.png", fileUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Download Finished")
}
func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}
	fmt.Print("\n")
	out.Close()
	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}
