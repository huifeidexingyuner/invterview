package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	root   = flag.String("root", "./*", "for example:invterview.exe --root=PATH\\*\\*.go")
	record = flag.String("record", "record.txt", "filter regular")
)

func genSha1(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	h := sha1.New()
	_, erro := io.Copy(h, file)
	if erro != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func scanDir(path string, output string) error {
	var fd *os.File
	var err error
	var list []string
	var fileinfo os.FileInfo
	fd, err = os.Create(output)

	list, err = filepath.Glob(path)

	if err != nil {
		return err
	}
	for _, v := range list {
		fmt.Println(v)
		fileinfo, err = os.Stat(v)
		if err != nil {
			return err
		}
		sh1, _ := genSha1(v)

		if err != nil {
			return err
		}
		fmt.Println(fileinfo.Name(), fileinfo.IsDir())
		if fileinfo.IsDir() {
			continue
		}
		value := fmt.Sprintf("%s,%d,%x\n", v, fileinfo.Size(), sh1)
		fd.WriteString(value)
	}

	return nil
}

func main() {
	flag.Parse()
	err := scanDir(*root, *record)
	fmt.Println(err)
}
