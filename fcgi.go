package main

import (
	"github.com/flashmob/go-fastcgi-client"
	"net/url"
	"fmt"
	"flag"
	"strings"
	"strconv"
)

func main() {
	uri := flag.String("url", "fastcgi://172.16.22.101/tmp/a.php", "url to request")
	e := flag.String("env", "a=A&b=B", "enviroment")

	flag.Parse()

	u, err := url.Parse(*uri)
	if err != nil {
		panic(err)
	}
	env := make(map[string]string)
	if err != nil {
		panic(err)
	}

	arr := strings.Split(u.Host, ":")
	host := arr[0]
	var port int = 9000
	if len(arr) > 1 {
		_port,err := strconv.ParseInt(arr[1], 10 ,32)
		if err != nil {
			panic(err)
		}
		port = int(_port)
	}
	mapEnv,err  := url.ParseQuery(*e)
	env["REQUEST_METHOD"] = "GET"
	env["SCRIPT_FILENAME"] = u.Path
	env["SERVER_SOFTWARE"] = "go / fcgiclient "
	env["REMOTE_ADDR"] = "127.0.0.1"
	env["SERVER_PROTOCOL"] = "HTTP/1.1"
	env["QUERY_STRING"] = u.RawQuery
	for k,v := range mapEnv {
		env[k] = v[0]
	}
	fcgi, err := fcgiclient.New(host, port)
	if err != nil {
		panic(err)
	}
	content,  err := fcgi.Request(env, "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", content)
}
