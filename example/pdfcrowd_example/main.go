package main

import (
	"fmt"
	"os"

	"github.com/pdfcrowd/pdfcrowd-go"
)

func main() {
	// create the API client instance
	client := pdfcrowd.NewHtmlToPdfClient("demo", "ce544b6ea52a5621fb9d55f8b542d14d")

	// run the conversion and write the result to a file
	// err := client.ConvertUrlToFile("http://127.0.0.1:8081/SetDatassetApply.html", "apply.pdf")

	err := client.ConvertFileToFile("./SetDatassetApply.html", "apply.pdf")

	// check for a conversion error
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		// report the error
		why, ok := err.(pdfcrowd.Error)
		if ok {
			os.Stderr.WriteString(fmt.Sprintf("Pdfcrowd Error: %s\n", why))
		} else {
			os.Stderr.WriteString(fmt.Sprintf("Generic Error: %s\n", err))
		}

		// rethrow or handle the exception
		panic(err.Error())
	}
}
