package templatemethod

import "fmt"

// 抽象模板接口
type Downloader interface {
	Download(url string)
}

// 抽象子类接口
type implementer interface {
	download()
	save()
}

// 具体模板类
type template struct {
	implementer // 指向子类的接口
	url string
}

func newTemplate(impl implementer) *template {
	return &template{
		implementer: impl,
	}
}

// 暴露出去供调用
func (t *template)Download(url string)  {
	t.url = url
	t.implementer.download()
	t.implementer.save()
}

// 具体子类1
type HTTPDownload struct {
	*template // 继承父类
}

func NewHTTPDownload() Downloader {
	downloader := &HTTPDownload{}
	downloader.template = newTemplate(downloader)

	return downloader
}

func (h *HTTPDownload) download() {
	fmt.Printf("HTTP download：%s\n", h.url)
}

func (h *HTTPDownload) save() {
	fmt.Printf("HTTP save: %s\n", h.url)
}

// 具体子类2
type FTPDownload struct {
	*template
}

func NewFTPDownload() Downloader {
	downloader := &FTPDownload{}
	downloader.template = newTemplate(downloader)

	return downloader
}

func (f *FTPDownload) download() {
	fmt.Printf("FTP download: %s\n", f.url)
}

func (f *FTPDownload) save() {
	fmt.Printf("FTP save: %s\n", f.url)
}
