package template_method

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	ftp := NewFtpDownloader()
	ftp.Download("a")
	fmt.Println("------------------")
	http := NewHttpDownloader()
	http.Download("b")
	fmt.Println("------------------")
	d := NewDefaultDownloader()
	d.Download("c")
}

/*
before download
ftp download a
after download
default save
------------------
before download
http download b
after download
http save
------------------
before download
default download c
after download
default save
*/
