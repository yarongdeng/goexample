package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	bufferSize = 1024
)

func main() {
	var (
		src string
		dst string
	)

	flag.StringVar(&src, "src", "", "src file path")
	flag.StringVar(&dst, "dst", "", "dst file path")

	flag.Usage = func() {
		fmt.Println("usage: cp --src path --dst path")
	}

	flag.Parse()

	// 文件路径检查

	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer dstFile.Close()

	ctx := make([]byte, bufferSize)

	for {
		n, err := srcFile.Read(ctx)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
				return
			}
			break
		}
		dstFile.Write(ctx[:n])
	}

}
