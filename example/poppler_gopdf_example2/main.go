package main

import (
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/signintech/gopdf"
)

func main() {
	notification := "1547887167022305280.document.pdf"

	cmd := exec.Command("pdftohtml", "-s", "-xml", "-stdout", "-zoom", "1", "-i", notification)
	// cmd := exec.Command("pdftohtml", "-v")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out)
	// log.Println("cmd output: ", s)
	lines := strings.Split(s, "\n")

	signExp := regexp.MustCompile(`<text top="([\d]+)" left="([\d]+)" width="([\d]+)" height="([\d]+)" font="([\d]+)">申请人签名:</text>`)
	pageExp := regexp.MustCompile(`<page number="([\d]+)" position="absolute" top="[\d]+" left="[\d]+" height="[\d]+" width="[\d]+">`)

	var pageNo uint64
	mm := make(map[uint64][]string)
	for _, line := range lines {
		pagePos := pageExp.FindStringSubmatch(line)
		if len(pagePos) > 0 {
			pageNo, _ = strconv.ParseUint(pagePos[1], 10, 64)
			continue
		}
		signPos := signExp.FindStringSubmatch(line)
		if len(signPos) == 0 {
			continue
		}
		log.Println("page no : ", pageNo)
		log.Println("sign pos: ", signPos)
		mm[pageNo] = signPos
	}
	log.Println("max page: ", pageNo)

	signature := "signature.png"
	dst := "1547887167022305280.signature.pdf"
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeLetter,
	})
	for i := 1; i <= int(pageNo); i++ {
		tpl := pdf.ImportPage(notification, i, "/MediaBox")
		pdf.AddPage()
		pdf.UseImportedTemplate(tpl, 0, 0, 0, 0)

		signPos, ok := mm[uint64(i)]
		if !ok {
			continue
		}
		signTop, _ := strconv.ParseFloat(signPos[1], 10)
		signLeft, _ := strconv.ParseFloat(signPos[2], 10)
		signWeight, _ := strconv.ParseFloat(signPos[3], 10)
		signHeight, _ := strconv.ParseFloat(signPos[4], 10)
		signX := signLeft + signWeight + 2
		signY := signTop - signHeight - 30
		signW := float64(90)
		signH := float64(60)

		if err := pdf.Image(signature, signX, signY, &gopdf.Rect{W: signW, H: signH}); err != nil {
			log.Fatal(err)
		}
	}
	// log.Debug("dst: %s", dst)
	if err := pdf.WritePdf(dst); err != nil {
		log.Println(err)
		return
	}
}
