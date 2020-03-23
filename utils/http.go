package utils

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/fanjq99/common/log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var MaxHttpBodySize = 2 * 1024 * 1024

// GetHTTPOriginalBody 返回未压缩的 Body
func GetHTTPOriginalBody(response *http.Response) ([]byte, error) {
	var err error
	bodyReader := io.LimitReader(response.Body, int64(MaxHttpBodySize))
	contentEncoding := strings.ToLower(response.Header.Get("Content-Encoding"))
	if !response.Uncompressed {
		if contentEncoding == "gzip" {
			bodyReader, err = gzip.NewReader(bodyReader)
			if err != nil {
				return nil, err
			}
		} else if contentEncoding == "deflate" {
			bodyReader, err = zlib.NewReader(bodyReader)
			if err != nil {
				return nil, err
			}
		}
	}

	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetHTTPUtf8Body for net/http response
func GetHTTPUtf8Body(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	body, err := GetHTTPOriginalBody(response)
	if err != nil {
		return nil, err
	}
	body, _ = ForceHtmlUtf8(body, string(response.Header.Get("Content-Type")))
	return body, nil
}

func ForceHtmlUtf8(body []byte, contentType string) ([]byte, string) {
	htmlCharset := getCharSet(contentType)
	if htmlCharset == "" {
		htmlCharset = detectHtmlCharset(body)
		if htmlCharset == "" {
			_, htmlCharset, _ = charset.DetermineEncoding(body, contentType)
		}
	}

	return ForceUtf8(body, htmlCharset)
}

func getCharSet(contentType string) string {
	defer func() {
		if err := recover(); err != nil {
			log.Error("getCharSet error", err, contentType)
		}
	}()
	content := strings.ToLower(contentType)

	pos := strings.Index(content, "charset=")
	if pos > 0 {
		begin := pos + 8
		if len(contentType) >= begin {
			return strings.TrimSpace(contentType[begin:])
		}
	}
	return ""
}

var charsetPattern = regexp.MustCompile(`(?i)<meta[^>]+charset\s*=\s*["]{0,1}([a-z0-9-]*)`)

func detectHtmlCharset(body []byte) string {
	if len(body) > 1024 {
		body = body[:1024]
	}
	match := charsetPattern.FindSubmatch(body)
	if match == nil {
		return ""
	}
	return string(match[1])
}

func ForceUtf8(body []byte, charsetName string) ([]byte, string) {
	if charsetName == "utf-8" && utf8.Valid(body) {
		return body, charsetName
	}

	reader := transform.NewReader(bytes.NewReader(body), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		log.Error(e)
		return body, charsetName
	}
	return d, charsetName
}
