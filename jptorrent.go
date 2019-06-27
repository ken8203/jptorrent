package jptorrent

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/ken8203/jptorrent/options"
)

const downloadURL = "http://www.jptorrent.org/download.php"

var (
	contentType = "application/x-www-form-urlencoded; param=value"
	accept      = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"
	host        = "www.jptorrent.org"
	origin      = "http://www.jptorrent.org"
	referer     = "http://www.jptorrent.org/link.php?ref="
	userAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"
)

func Download(ref string, opts ...options.Option) error {
	if ref == "" {
		return fmt.Errorf("ref can't be empty")
	}

	path := "./" + ref + ".torrent"
	if len(opts) > 0 {
		path = opts[0].Location
	}

	data := url.Values{
		"ref":    {ref},
		"Submit": {"Download"},
	}

	req, _ := http.NewRequest("POST", downloadURL, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", accept)
	req.Header.Set("Host", host)
	req.Header.Set("Origin", origin)
	req.Header.Set("Referer", referer+ref)
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
