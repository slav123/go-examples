package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	m := martini.Classic()

	m.Post("/send", retrieve)

	m.Run()

	log.Fatal(http.ListenAndServe(":8080", m))

}

func retrieve(r *http.Request) {

	const _24K = (1 << 20) * 24

	var (
		status int
		err    error
	)

	// parse multipart form from go
	if err := r.ParseMultipartForm(_24K); nil != err {
		status = http.StatusInternalServerError
		fmt.Println(err)
		return
	}

	// read field
	body := r.FormValue("field")

	fmt.Println(body)

	// prarse "file apart"
	for _, fheaders := range r.MultipartForm.File {
		for _, hdr := range fheaders {
			// open uploaded
			var infile multipart.File
			if infile, err = hdr.Open(); nil != err {
				status = http.StatusInternalServerError
				return
			}
			// open destination
			var outfile *os.File
			if outfile, err = os.Create("./uploaded/" + hdr.Filename); nil != err {
				status = http.StatusInternalServerError
				return
			}
			// 32K buffer copy
			var written int64
			if written, err = io.Copy(outfile, infile); nil != err {
				status = http.StatusInternalServerError
				return
			}
			fmt.Println("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written)))
			fmt.Println(status)
			//res.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))
		}
	}

	return
}
