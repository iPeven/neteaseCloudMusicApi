package ncm

import (
	"net/http"
	"net/url"
	"strings"
	"log"
	"io/ioutil"
)

const(
	Referer = "http://music.163.com/"
	ContentType = "application/x-www-form-urlencoded"
	UserAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0"
)

func doPost(postUrl string,params string) string{
	c := &http.Client{}

	fromData := url.Values{}
	fromData.Set("params",params)
	fromData.Set("encSecKey",encSecKey)
	body := strings.NewReader(fromData.Encode())

	req, err := http.NewRequest("POST",postUrl,body)
	req.Header.Set("Referer",Referer)
	req.Header.Set("User-Agent",UserAgent)
	req.Header.Set("Content-Type",ContentType)
	if err != nil{
		log.Fatal(err)
	}
	resp, err := c.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil{
		log.Fatal(err)
	}
	return string(respBody)
}