package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/signintech/gopdf"
)

type Pos struct {
	Sign []string
	Recv []string
	Date [][]string
}

func main() {
	notification := "1558011508921733120.document.pdf"
	signs := []int{2}
	ss := make(map[int]struct{})
	for _, s := range signs {
		ss[s] = struct{}{}
	}
	i := 0
	cmd := exec.Command("pdftohtml", "-s", "-xml", "-stdout", "-zoom", "1", "-i", notification)
	// cmd := exec.Command("pdftohtml", "-v")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	// log.Println("cmd output: ", string(out))
	lines := strings.Split(string(out), "\n")

	signExp := regexp.MustCompile(`<text top="([\d]+)" left="([\d]+)" width="([\d]+)" height="([\d]+)" font="([\d]+)">申请人签名:</text>`)
	recvExp := regexp.MustCompile(`<text top="([\d]+)" left="([\d]+)" width="([\d]+)" height="([\d]+)" font="([\d]+)">收件人签名:</text>`)
	dateExp := regexp.MustCompile(`<text top="([\d]+)" left="([\d]+)" width="([\d]+)" height="([\d]+)" font="([\d]+)">签名日期:</text>`)
	prntExp := regexp.MustCompile(`\<text top="([\d]+)" left="([\d]+)" width="([\d]+)" height="([\d]+)" font="([\d]+)">打印日期:.*</text>`)
	pageExp := regexp.MustCompile(`<page number="([\d]+)" position="absolute" top="[\d]+" left="[\d]+" height="[\d]+" width="[\d]+">`)

	var pageNo uint64
	mm := make(map[uint64]*Pos)
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
		i++
		if len(ss) > 0 {
			if _, ok := ss[i]; !ok {
				continue
			}
		}
		log.Println("page no : ", pageNo)
		log.Println("sign pos: ", signPos)
		mm[pageNo] = &Pos{Sign: signPos}
	}
	for _, line := range lines {
		pagePos := pageExp.FindStringSubmatch(line)
		if len(pagePos) > 0 {
			pageNo, _ = strconv.ParseUint(pagePos[1], 10, 64)
			continue
		}
		_, ok := mm[pageNo]

		recvPos := recvExp.FindStringSubmatch(line)
		if len(recvPos) > 0 && ok {
			mm[pageNo].Recv = recvPos
			log.Println("recv pos: ", recvPos)
			continue
		}
		datePos := dateExp.FindStringSubmatch(line)
		if len(datePos) > 0 && ok {
			mm[pageNo].Date = append(mm[pageNo].Date, datePos)
			log.Println("date pos: ", datePos)
			continue
		}
		prntPos := prntExp.FindStringSubmatch(line)
		if len(prntPos) > 0 && ok {
			mm[pageNo].Date = append(mm[pageNo].Date, prntPos)
			log.Println("print pos: ", prntPos)
		}
	}
	log.Println("max page: ", pageNo)

	signSrc := "./signature.png"
	signDst := "./transparent.signature.png"
	transparentize(signSrc, signDst)
	dst := "1558011508921733120.signature.pdf"
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeLetter,
	})
	for i := 1; i <= int(pageNo); i++ {
		tpl := pdf.ImportPage(notification, i, "/MediaBox")
		pdf.AddPage()
		pdf.UseImportedTemplate(tpl, 0, 0, 0, 0)

		m, ok := mm[uint64(i)]
		if !ok {
			continue
		}
		x, y := getXY(m.Sign)
		signW := float64(90)
		signH := float64(60)
		if err := pdf.Image(signDst, x, y-40, &gopdf.Rect{W: signW, H: signH}); err != nil {
			log.Fatal(err)
		}
		if err := pdf.AddTTFFont("Regular", "stfangsong.ttf"); err != nil {
			log.Fatal(err)
		}
		if len(m.Recv) > 0 {
			x, y := getXY(m.Recv)
			pdf.SetX(x)
			pdf.SetY(y - 5)
			pdf.SetTextColor(0, 0, 0)
			_ = pdf.SetFont("Regular", "", 10)
			_ = pdf.Text("测试收件人")
		}
		for _, d := range m.Date {
			x, y := getXY(d)
			pdf.SetX(x)
			pdf.SetY(y - 5)
			pdf.SetTextColor(0, 0, 0)
			_ = pdf.SetFont("Regular", "", 10)
			_ = pdf.Text(time.Now().Format("2006年1月02日"))
		}
	}
	// log.Debug("dst: %s", dst)
	if err := pdf.WritePdf(dst); err != nil {
		log.Println(err)
		return
	}
}

func getXY(pos []string) (float64, float64) {
	top, _ := strconv.ParseFloat(pos[1], 10)
	left, _ := strconv.ParseFloat(pos[2], 10)
	weight, _ := strconv.ParseFloat(pos[3], 10)
	height, _ := strconv.ParseFloat(pos[4], 10)
	x := left + weight + 2
	y := top + height
	return x, y
}

func transparentize(src, dst string) {
	imgfile, err := os.Open(src)
	defer imgfile.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}

	decodedImg, err := png.Decode(imgfile)
	img := image.NewRGBA(decodedImg.Bounds())

	cc := decodedImg.At(0, 0)
	//r, g, b, a := cc.RGBA()
	//fmt.Printf("color: %v, (%v, %v, %v, %v)\n", cc, r, g, b, a)
	size := img.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			c := decodedImg.At(x, y)
			if !similar(c, cc) {
				img.Set(x, y, decodedImg.At(x, y))
			}
		}
	}

	// I change some pixels here with img.Set(...)

	outFile, _ := os.Create(dst)
	defer outFile.Close()
	_ = png.Encode(outFile, img)
}

func similar(a, b color.Color) bool {
	r1, g1, b1, _ := a.RGBA()
	r2, g2, b2, _ := b.RGBA()
	if math.Abs((float64(r1)-float64(r2))/65535.0) < 0.05 && math.Abs((float64(g1)-float64(g2))/65535.0) < 0.05 &&
		math.Abs((float64(b1)-float64(b2))/65535.0) < 0.05 {
		return true
	}
	//fmt.Println(float64(r1-r2)/65535.0, "=", r2, r1-r2, r1, g1, b1)
	return false
}
