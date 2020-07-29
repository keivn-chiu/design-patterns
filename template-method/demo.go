package template_method

import "fmt"

// template method put general functions and steps into parent type, the actual implements in child

// IDownloader the parent of interface
type IDownloader interface {
	Download(url string)
}

// ITemplate the interface of template
type ITemplate interface {
	download()
	save()
}

// template the parent type
type template struct {
	ITemplate
	url string
}

func NewTemplate(t ITemplate) *template {
	return &template{ITemplate: t}
}
func (t *template) Download(url string) {
	t.url = url
	fmt.Println("before download")
	t.ITemplate.download()
	fmt.Println("after download")
	t.ITemplate.save()
}
func (t *template) save() {
	fmt.Println("default save")
}
func (t *template) download() {
	fmt.Println("default download", t.url)
}

// HttpDownloader the child the implement of ITemplate
type HttpDownloader struct {
	*template
}

func NewHttpDownloader() IDownloader {
	downloader := &HttpDownloader{}
	t := NewTemplate(downloader)
	downloader.template = t
	return downloader
}
func (h *HttpDownloader) download() {
	fmt.Println("http download", h.url)
}
func (h *HttpDownloader) save() {
	fmt.Println("http save")
}

// FtpDownloader the child the implement of ITemplate, use default save
type FtpDownloader struct {
	*template
}

func NewFtpDownloader() IDownloader {
	downloader := &FtpDownloader{}
	t := NewTemplate(downloader)
	downloader.template = t
	return downloader
}

func (f *FtpDownloader) download() {
	fmt.Println("ftp download", f.url)
}

// DefaultDownloader the child the implement of ITemplate, use default save and download
type DefaultDownloader struct {
	*template
}

func NewDefaultDownloader() IDownloader {
	downloader := &DefaultDownloader{}
	t := NewTemplate(downloader)
	downloader.template = t
	return downloader
}
