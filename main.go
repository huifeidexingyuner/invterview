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

func getFilelist(path string, fd *os.File) error {
	var err error
	var sh1 []byte

	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		sh1, err = genSha1(path)

		value := fmt.Sprintf("%s,%x,%d\n", path, sh1, f.Size())
		fd.WriteString(value)
		return nil
	})
	return err
}

func scanDir(path string, output string) error {
	var fd *os.File
	var err error
	var list []string

	fd, err = os.Create(output)

	if err != nil {
		return err
	}
	defer fd.Close()

	list, err = filepath.Glob(path)

	if err != nil {
		return err
	}
	for _, v := range list {
		getFilelist(v, fd)
	}

	return nil
}

func main() {
	flag.Parse()
	err := scanDir(*root, *record)
	fmt.Println(err)
}
