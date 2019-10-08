package main

import (
	"fmt"
	"os"
	"log"
	"errors"
	"bufio"
	"strings"
	"io"
	"bytes"
	"os/exec"
	"github.com/spf13/pflag"
)

func selectPages() error {
	// flags
	startPage := pflag.IntP("start page", "s", 0, "Start page to print")
	endPage := pflag.IntP("end page", "e", 0, "End page to print")
	pageLen := pflag.IntP("page length", "l", 72, "The number of lines in each page")
	pageType := pflag.BoolP("page type", "f", false, "Whether '\\f' is page breaker or not")
	printDest := pflag.StringP("print dest", "d", "", "The destination to print")

	// parse flags
	pflag.Parse()

	// check flags
	if *startPage < 0 || *endPage < 0 {
		return errors.New("The page number should be positive.")
	}
	if *startPage > *endPage {
		return errors.New("The start page number should be smaller than the end page number.")
	}
	if *pageLen != 72 && *pageType == true {
		return errors.New("The \"-l\" flag and the \"-f\" flag shouldn't be used together.")
	}
	if *pageLen < 0 {
		return errors.New("The page size should be positive.")
	}
	if pflag.NArg() > 1 {
		return errors.New("Too many arguments.")
	}

	// create reader
	var reader *bufio.Reader
	if pflag.NArg() == 1 {
		// read from file
		file, err := os.Open(pflag.Arg(0))
		if err != nil {
			return err
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	} else {
		// read from stdin
		reader = bufio.NewReader(os.Stdin)
	}

	// write target pages to buffer
	var buf strings.Builder
	if *pageType == true {
		// read one page each time
		pageCount := 0
		for {
			page, err := reader.ReadBytes('\f')
			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			pageCount += 1
			if pageCount >= *startPage && pageCount <= *endPage {
				buf.Write(page)
			} else if pageCount > *endPage {
				break
			}
		}
	} else {
		// read one line each time
		pageCount := 0
		lineCount := 0
		for {
			line, err := reader.ReadBytes('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			lineCount += 1
			if lineCount > *pageLen {
				pageCount += 1
				lineCount = 1
			}
			if pageCount >= *startPage && pageCount <= *endPage {
				buf.Write(line)
			} else if pageCount > *endPage {
				break
			}
		}
	}

	// write buffer to destination
	if *printDest == "" {
		// write to stdout
		fmt.Printf("%s", buf.String())
	} else {
		// write to printer
		cmd := exec.Command("lp", "-d" + *printDest)
		cmd.Stdin = strings.NewReader(buf.String())
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("Write to printer failed with %s.\nMessage: %s", err, stderr.Bytes())
		}
	}

	return nil
}

func main() {
	// select pages
	err := selectPages()
	// handle error
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}