package RedirectClient

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"time"
)

type RedirectClient struct {
	*http.Client
}

func (c *RedirectClient) BanRedirectDo(req *http.Request) (*http.Response, error) {
	c.SetCheckRedirect(false)
	defer c.SetCheckRedirect(true)
	rep, err := c.Client.Do(req)
	return rep, err
}
func (c *RedirectClient) BanRedirectGet(url string) (*http.Response, error) {
	c.SetCheckRedirect(false)
	defer c.SetCheckRedirect(true)
	rep, err := c.Client.Get(url)
	return rep, err
}

func defaultCheckRedirect(req *http.Request, via []*http.Request) error {
	if len(via) >= 10 {
		return errors.New("stopped after 10 redirects")
	}
	return nil
}
func newCheckRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func (c *RedirectClient) SetCheckRedirect(i bool) {
	if i {
		c.CheckRedirect = defaultCheckRedirect
	} else {
		c.CheckRedirect = newCheckRedirect
	}
}

func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	rand.Seed(time.Now().UnixNano())
	request, err := http.NewRequest(method, url, body)
	fakeip := fmt.Sprintf("59.79.%d,%d", rand.Intn(250)+2, rand.Intn(250)+2)
	request.Header.Add("X-Forwarded-For", fakeip)
	return request, err
}
func Addmultipartkvs(bw *multipart.Writer, kvs map[string]string) {
	for k, v := range kvs {
		fw, _ := bw.CreateFormField(k)
		fw.Write([]byte(v))
	}
}
