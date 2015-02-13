package main

import (
	"crypto/md5"
    "fmt"
    "os"
    "io"
    "io/ioutil"
    "strings"
)

func main() {
	var md []byte
	var dir, path, calc string
	
	// create map for existing files
	list := make(map[string]bool)
	
	// see of args
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		dir = "./test"
	}

	// display name of parsed directory
	fmt.Printf("Reading %s\n", dir)

	// get list of files
    files, _ := ioutil.ReadDir(dir)

    // iterate
    for _, f := range files {
		
		// create path to file (dir + name)
		path = dir + "/" + f.Name()

		// check if it's not directory
		if info, err := os.Stat(path); err == nil && !info.IsDir() {

			// calc md5
       		md,_ = ComputeMd5(dir + "/" + f.Name())    	
		
			// create string to use in map
			calc = string(md[:])
       
       		// show name and md5 of file
        	fmt.Printf("File: %s, md5: %x\n", f.Name(), md)

        	// if file is already on the list, mark as a duplicate
        	if list[calc] {

        		fmt.Printf("Possible duplicate: %s\n", f.Name())
        		
        		// check if is not already .duplicate
        		if !strings.HasSuffix(path, ".duplicate") {
        			os.Rename(path, path + ".duplicate")
        		}
       		} else {
				list[calc] = true
			}
        
   	 	} else {
   	 		fmt.Printf("Dir: %s - skipped\n", path)
   	 	} // end if

	
    } // end for
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
