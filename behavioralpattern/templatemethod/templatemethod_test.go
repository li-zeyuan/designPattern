package templatemethod

import "testing"

func TestHTTPDownload(t *testing.T) {
	downloader := NewHTTPDownload()
	downloader.Download("http://lizeyuan.com")
}

func TestFTPDownload(t *testing.T) {
	downloader := NewFTPDownload()
	downloader.Download("http://lizeyuan.com")
}