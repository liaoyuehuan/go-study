package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"net/url"
	"reflect"
)

func get() {
	var url string = "http://47.107.37.69:9501/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func Post() {
	var url string = "http://192.168.50.74:9501/admin/index/index?aaa=1"

	err := ioutil.WriteFile("upload2", []byte("ok"), 0675)
	if err != nil {
		panic(err)
	}

	rd, err := ioutil.ReadFile("src/example/net/ewm.png");
	if err != nil {
		panic(err)
	}

	rs, err := http.Post(url, "image/png", strings.NewReader(string(rd)))
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	fmt.Printf("resp status %s,statusCode %d\n", rs.Status, rs.StatusCode)

	fmt.Printf("resp Proto %s\n", rs.Proto)

	fmt.Printf("resp content length %d\n", rs.ContentLength)

	fmt.Printf("resp transfer encoding %v\n", rs.TransferEncoding)

	fmt.Printf("resp Uncompressed %t\n", rs.Uncompressed)

	body, _ := ioutil.ReadAll(rs.Body);
	fmt.Printf("resp body %s\n", body)

	fmt.Printf("resp header %v \n", rs.Header)

	fmt.Printf("resp tls %v\n", rs.TLS)

	fmt.Printf("reflect type:  %v\n",reflect.TypeOf(rs.Body))


}

func PostForm()  {
	var url2 string = "http://192.168.50.74:9501/admin/index/index?aaa=1"
	resp,err := http.PostForm(url2,url.Values{
		"a": {"a"},
		"b": {"b"}})
	if err != nil {
		panic(nil)
	}
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(nil)
	}
	fmt.Println(string(body))

}

func main() {
	Post()
}
