package http

import (
	"github.com/ohmygd/mgo/config"
	"github.com/ohmygd/mgo/merror"
	"github.com/ohmygd/mgo/pc"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DaoHttp struct {
	Module  string
	Timeout int
}

func (d *DaoHttp) SetTimeout(timeout int) {
	d.Timeout = timeout
}

func (d *DaoHttp) GetUriStr(uriStr string) (url string) {
	httpInfo := config.GetHttpMsg(d.Module)
	if httpInfo == nil {
		panic("http config lost")
	}

	info := httpInfo.(map[string]interface{})
	url = info["url"].(string)
	uri := info["uri"].(map[string]interface{})[uriStr].(string)

	url += uri

	return
}

func (d *DaoHttp) Post(uriStr string, params map[string]string, headers map[string]string) (resp string, err error) {
	client := &http.Client{
		Timeout: time.Duration(d.Timeout * int(time.Second)),
	}

	values := url.Values{}
	for k, v := range params {
		values.Add(k, v)
	}

	p := strings.NewReader(values.Encode())

	moduleUrl := d.GetUriStr(uriStr)

	req, _ := http.NewRequest("POST", moduleUrl, p)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	addHeader(req, headers)

	res, _ := client.Do(req)

	if res == nil {
		err = merror.New(pc.ErrorHttpGetPost)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	resp = string(body)

	return
}

func (d *DaoHttp) Get(uriStr string, headers map[string]string) (resp string, err error) {
	client := &http.Client{
		Timeout: time.Duration(d.Timeout * int(time.Second)),
	}

	req, _ := http.NewRequest("GET", d.GetUriStr(uriStr), nil)

	addHeader(req, headers)

	res, _ := client.Do(req)
	if res == nil {
		err = merror.New(pc.ErrorHttpGetPost)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	resp = string(body)

	return
}

func addHeader(r *http.Request, headers map[string]string) {
	if len(headers) == 0 {
		return
	}

	for k, v := range headers {
		r.Header.Set(k, v)
	}

	return
}
