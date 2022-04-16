package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/signintech/gopdf"
)

const (
	yearX float64 = 334
	monX  float64 = 410
	dayX  float64 = 474
)

func main() {
	notification := "1509484018444275712.notification.pdf"

	cmd := exec.Command("pdftohtml", "-s", "-xml", "-stdout", "-zoom", "1", "-i", notification)
	// cmd := exec.Command("pdftohtml", "-v")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out)
	log.Println("cmd output: ", s)
	signExp := regexp.MustCompile(`<text top="([\d]+)" left="([\d]+)" width="([\d]+)" height="([\d]+)" font="([\d]+)">受送达人签名：</text>`)
	signPos := signExp.FindStringSubmatch(s)
	log.Println("sign pos: ", signPos)
	dateExp := regexp.MustCompile(`<text top="([\d]+)" left="([\d]+)" width="([\d]+)" height="([\d]+)" font="([\d]+)">.*年.*月.*日</text>`)
	datePos := dateExp.FindStringSubmatch(s)
	log.Println("date pos: ", datePos)

	signTop, _ := strconv.ParseFloat(signPos[1], 10)
	signLeft, _ := strconv.ParseFloat(signPos[2], 10)
	signWeight, _ := strconv.ParseFloat(signPos[3], 10)
	signHeight, _ := strconv.ParseFloat(signPos[4], 10)
	signX := signLeft + signWeight + 2
	signY := signTop - signHeight
	signW := float64(67)
	signH := float64(40)

	dateTop, _ := strconv.ParseFloat(datePos[1], 10)
	dateHeight, _ := strconv.ParseFloat(datePos[4], 10)
	dateY := dateTop + dateHeight

	signature := "signature.png"
	dst := "1509484018444275712.signature.pdf"
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeLetter,
	})
	tpl := pdf.ImportPage(notification, 1, "/MediaBox")
	pdf.AddPage()
	pdf.UseImportedTemplate(tpl, 0, 0, 0, 0)
	if err := pdf.Image(signature, signX, signY, &gopdf.Rect{W: signW, H: signH}); err != nil {
		log.Fatal(err)
	}
	if err := pdf.AddTTFFont("Regular", "SourceHanSansCN-Regular.ttf"); err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	pdf.SetY(dateY)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.SetFont("Regular", "", 15)
	pdf.SetX(yearX)
	_ = pdf.Text(fmt.Sprintf("%d", now.Year()))
	pdf.SetX(monX)
	_ = pdf.Text(fmt.Sprintf("%d", now.Month()))
	pdf.SetX(dayX)
	_ = pdf.Text(fmt.Sprintf("%02d", now.Day()))
	// log.Debug("dst: %s", dst)
	if err := pdf.WritePdf(dst); err != nil {
		log.Println(err)
		return
	}
}
