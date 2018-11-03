// Copyright © 2018 ehlxr <ehlxr.me@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ehlxr/hex-convert/gen"

	"github.com/ehlxr/hex-convert/converter"
)

func decimal(w http.ResponseWriter, req *http.Request) {
	scale, _ := strconv.Atoi(req.FormValue("scale"))
	data := req.FormValue("data")

	result, err := converter.ToDecimal(scale, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Println(result)
	fmt.Fprintf(w, "<a href='/'>首页</a><br> %s", strconv.Itoa(result))
}

func binary(w http.ResponseWriter, req *http.Request) {
	scale, _ := strconv.Atoi(req.FormValue("scale"))
	data := req.FormValue("data")

	result, err := converter.ToBinary(scale, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Println(result)
	fmt.Fprintf(w, "<a href='/'>首页</a><br> %s", result)
}

func index(w http.ResponseWriter, r *http.Request) {
	f, err := gen.Assets.Open("/index.tpl")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(fd))
}

func Start(host string, port int) error {
	http.HandleFunc("/", index)
	http.HandleFunc("/d", decimal)
	http.HandleFunc("/b", binary)

	addr := fmt.Sprintf("%s:%d", host, port)
	if strings.Contains(addr, "0.0.0.0") {
		addr = strings.Replace(addr, "0.0.0.0", "", 1)
		host = strings.Replace(host, "0.0.0.0", "127.0.0.1", 1)
	}
	fmt.Printf("server start on: %s\n", fmt.Sprintf("http://%s:%d", host, port))

	if err := http.ListenAndServe(addr, nil); err != nil {
		return fmt.Errorf("ListenAndServe: : %v", err)
	}
	return nil
}
