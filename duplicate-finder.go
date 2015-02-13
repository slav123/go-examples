package main

import (
	"crypto/md5"
    "fmt"
    "os"
    "io"
    "io/ioutil"
)

func main() {
	var md []byte
	var dir, calc string
	

	list := make(map[string]bool)
	
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		dir = "./test"
	}

	fmt.Printf("Reading %s\n", dir)

    files, _ := ioutil.ReadDir(dir)

    for _, f := range files {
		
		md,_ = ComputeMd5(dir + "/" + f.Name())    	
		
		calc = string(md[:])
       
        fmt.Printf("File: %s, md5: %x\n", f.Name(), md)

        if list[calc] {
        	fmt.Printf("Possible duplicate: %s\n", f.Name())
        }

        list[calc] = true
    }
}

func ComputeMd5(filePath string) ([]byte, error) {
  var result []byte
  file, err := os.Open(filePath)
  if err != nil {
    return result, err
  }
  defer file.Close()

  hash := md5.New()
  if _, err := io.Copy(hash, file); err != nil {
    return result, err
  }

  return hash.Sum(result), nil
}
