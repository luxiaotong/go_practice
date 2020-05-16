package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	wk "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

var in = flag.String("in", "", "指定的 HTML 文件地址")
var out = flag.String("out", "", "指定的 PDF 文件地址")

func main() {

	flag.Parse()
	// fmt.Println("-f:", *f)

	// Create new PDF generator
	pdfg, err := wk.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.Dpi.Set(600)
	pdfg.NoCollate.Set(false)
	pdfg.PageSize.Set(wk.PageSizeA4)
	pdfg.MarginBottom.Set(40)
	pdfg.Orientation.Set(wk.OrientationLandscape)

	htmlfile, err := ioutil.ReadFile(*in)
	if err != nil {
		log.Fatal(err)
	}
	pdfg.AddPage(wk.NewPageReader(bytes.NewReader(htmlfile)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile(*out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
