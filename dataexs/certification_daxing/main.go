package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"strings"
	"time"

	"github.com/signintech/gopdf"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"rsc.io/qr"
)

const (
	TplBackground = "background_tpl.pdf"
	Stamp         = "stamp.png"
	Logo          = "logo.png"
	FontRegular   = "SourceHanSansCN-Regular.ttf"
	FontHeavy     = "SourceHanSansCN-Heavy.ttf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4, //595.28, 841.89 = A4
	})

	if err := pdf.AddTTFFont("Regular", FontRegular); err != nil {
		log.Fatal(err)
	}
	if err := pdf.AddTTFFont("Heavy", FontHeavy); err != nil {
		log.Fatal(err)
	}

	printer := message.NewPrinter(language.English)
	tplBg := pdf.ImportPage(TplBackground, 1, "/MediaBox")
	pdf.AddPage()
	pdf.UseImportedTemplate(tplBg, 0, 0, 0, 0)

	_ = pdf.Image(Logo, 267, 70, &gopdf.Rect{W: 260, H: 12})
	pdf.SetX(59)
	pdf.SetY(135)
	_ = pdf.SetFont("Heavy", "", 32)
	pdf.SetTextColor(6, 63, 131)
	_ = pdf.Text("国际数据合作合规编码与登记证书")
	pdf.SetX(60)
	pdf.SetY(159)
	_ = pdf.SetFont("Heavy", "", 12)
	pdf.SetTextColor(6, 63, 131)
	_ = pdf.Text("International Data Transfer&Compliance Coding and Registration Certificate")
	pdf.SetX(61)
	pdf.SetY(80)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.SetFont("Regular", "", 14)
	_ = pdf.Text(fmt.Sprintf("NO.%s", "IDT27807206071729"))
	pdf.SetX(76)
	pdf.SetY(197)
	_ = pdf.SetFont("Regular", "", 13)
	_ = pdf.Text(fmt.Sprintf("数  据  要  素  名  称 : %s", "国际生物医药公司动物基因检测数据"))
	pdf.SetX(76)
	pdf.SetY(212)
	name := "International BioPharmaceutical Company Animal Genetic Test Data"
	ss, _ := pdf.SplitText(name, 327)
	_ = pdf.Text(fmt.Sprintf("(  D a t a    N a m e  ) %s", ss[0]))
	if len(ss) > 1 {
		y := 225 // 212+13
		for i, t := range ss[1:] {
			pdf.SetX(192)
			pdf.SetY(float64(y + i*13))
			_ = pdf.Text(t)
		}
	}
	pdf.SetX(76)
	pdf.SetY(250)
	_ = pdf.Text(fmt.Sprintf("数据要素所有单位  : %s", "国际生物医药公司"))
	pdf.SetX(76)
	pdf.SetY(267)
	_ = pdf.Text(fmt.Sprintf("(   O    w    n    e    r   ) %s", "International BioPharmaceutical Company"))
	pdf.SetX(76)
	pdf.SetY(298)
	_ = pdf.Text(fmt.Sprintf("数  据  要  素  标  识 : %s", "IDT99.0086.000000000"))
	pdf.SetX(76)
	pdf.SetY(315)
	_ = pdf.Text("(   D   a   t   a   I   D   )")
	pdf.SetX(76)
	pdf.SetY(346)
	_ = pdf.Text(printer.Sprintf("数  据  要  素  总  量 : %d DRs", 64480))
	pdf.SetX(76)
	pdf.SetY(363)
	_ = pdf.Text("(Amount  of   Data)")
	pdf.SetX(76)
	pdf.SetY(389)
	sources := []string{"程序爬虫取得", "其他，请注明testtesttest"}
	ss, _ = pdf.SplitText(strings.Join(sources, ";"), 327)
	_ = pdf.Text(fmt.Sprintf("获     取     方     式     : %s", ss[0]))
	if len(ss) > 1 {
		y := 403 // 389+13
		for i, t := range ss[1:] {
			pdf.SetX(192)
			pdf.SetY(float64(y + i*13))
			_ = pdf.Text(t)
		}
	}
	pdf.SetX(76)
	pdf.SetY(406)
	_ = pdf.Text("( A c q u i r e d  B y )")
	pdf.SetX(76)
	pdf.SetY(437)
	_ = pdf.Text(fmt.Sprintf("申     请     时     间     : %s", time.Now().Format("2006年1月2日")))
	pdf.SetX(76)
	pdf.SetY(454)
	_ = pdf.Text("(Application  Date)")
	pdf.SetX(76)
	pdf.SetY(482)
	_ = pdf.Text("审     核     机     构     : 北京大兴临空国际数据合作编码与登记中心(筹)")
	pdf.SetX(76)
	pdf.SetY(499)
	_ = pdf.Text("( V e r i f i e d    B y )  Beijing Daxing IAEZ International Data")
	pdf.SetX(194)
	pdf.SetY(513)
	_ = pdf.Text("Transfer&Compliance Coding and Registration Center")
	pdf.SetX(76)
	pdf.SetY(538)
	_ = pdf.Text(fmt.Sprintf("审     核     时     间     : %s", time.Now().Format("2006年1月2日")))
	pdf.SetX(76)
	pdf.SetY(555)
	_ = pdf.Text("(Verification  Date)")
	pdf.SetX(76)
	pdf.SetY(586)
	_ = pdf.Text(fmt.Sprintf("区     块       I       D     : %s", "45966274826631ae1c680fad0caeb1f47441ac9c"))
	pdf.SetX(76)
	pdf.SetY(603)
	_ = pdf.Text("(  Blockchain    ID  )")
	pdf.SetX(76)
	pdf.SetY(634)
	_ = pdf.Text(fmt.Sprintf("生  成  区  块  时  间 : %s", time.Now().Format("2006年1月2日 15:04:05")))
	pdf.SetX(76)
	pdf.SetY(651)
	_ = pdf.Text("(Blockchain ")
	pdf.SetX(76)
	pdf.SetY(664)
	_ = pdf.Text("Generation Time)")
	url := "http://datassests.cn"
	b, err := genQR(url, 150)
	if err != nil {
		log.Fatal(err)
	}
	img, err := gopdf.ImageHolderByBytes(b)
	if err != nil {
		log.Fatal(err)
	}
	if err := pdf.ImageByHolder(img, 76, 677, nil); err != nil {
		log.Fatal(err)
	}
	pdf.SetX(344)
	pdf.SetY(744)
	_ = pdf.Text(fmt.Sprintf("发  证  日  期  : %s", time.Now().Format("2006年01月02日")))
	pdf.SetX(344)
	pdf.SetY(761)
	_ = pdf.Text("(Issued Date)")
	_ = pdf.Image(Stamp, 414, 653, &gopdf.Rect{W: 119, H: 119})
	pdf.SetX(259)
	pdf.SetY(792)
	_ = pdf.Text("1/2")

	// Add New Page
	pdf.AddPage()
	pdf.UseImportedTemplate(tplBg, 0, 0, 0, 0)

	pdf.SetX(80)
	pdf.SetY(130)
	_ = pdf.SetFont("Heavy", "", 37)
	pdf.SetTextColor(6, 63, 131)
	_ = pdf.Text("数  据  概  况")
	pdf.SetX(80)
	pdf.SetY(160)
	_ = pdf.SetFont("Heavy", "", 14)
	pdf.SetTextColor(6, 63, 131)
	_ = pdf.Text("SUMMARY")
	pdf.SetX(67)
	pdf.SetY(71)
	_ = pdf.SetFont("Regular", "", 14)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.Text(fmt.Sprintf("NO.%s", "IDT278072060717290"))
	pdf.SetX(80)
	pdf.SetY(190)
	_ = pdf.SetFont("Regular", "", 13)
	_ = pdf.Text(fmt.Sprintf("数       据      表       数       量   : %d 张", 2))
	pdf.SetX(80)
	pdf.SetY(207)
	_ = pdf.Text("(  D  a  t  a       S  h  e  e  t  s  )")
	pdf.SetX(80)
	pdf.SetY(240)
	_ = pdf.Text(printer.Sprintf("数     据     字     段     数     量 : %d 个", 20))
	pdf.SetX(80)
	pdf.SetY(257)
	_ = pdf.Text("(  D  a  t  a         F  i  e  l  d  s  )")
	pdf.SetX(80)
	pdf.SetY(290)
	_ = pdf.Text(fmt.Sprintf("有    效    记    录    占    比      : %.2f%%", 100.00))
	pdf.SetX(80)
	pdf.SetY(307)
	_ = pdf.Text("(  Valid     Records     Ratio  )")

	pdf.SetX(80)
	pdf.SetY(340)
	hash := "40cf58a16853ba0b519ad26b643c2164d609f57da3f35e6331755b1380fcb89366ad15bb93ab2c647a11a10b572a08aac655a91d5f85d949b40599406b00ddda91dd3dd25a629068a471943fa37cb442e376d2e2691184a97f56a445bed7ad945bb67238e5af4445bababf7d65176e89d45432498f07903e2f26e5fe04b5be96efcbebd1dce85b5b9d142c79b0e3d33c4eec81d66b79adbad56e74448b9337c5d19ccd6c80412895fcd0ef4128fb1a889619d45de80fecc89aba0e97d5e8a58ed19ccd6c80412895fcd0ef4128fb1a889619d45de80fecc89aba0e97d5e8a58ef976bcf90cd0c0f917f3e5d4b51090e1f40de8e046c14c34cf0b77983f29d0dbf976bcf90cd0c0f917f3e5d4b51090e1f40de8e046c14c34cf0b77983f29d0db6115d01......7d1e0794ed636dde7e858e13012b7a00ac89f31972742cf60e8bc7d6dbbf11e3c1a3f8b159ca2b8b2a18e1f573e561faa93ccac282a34ae942acc1a0e516e19d507b204becd0517d4d438dfa18390244f52217c282a34ae942acc1a0e516e19d507b204becd0517d4d438dfa18390244f52217c282a34ae942acc1a0e516e19d507b204becd0517d4d438dfa18390244f522174496dcf237fdb680bfb272bb8beda69c7e7de70600c625b887ce1bc4e1cc17264496dcf237fdb680bfb272bb8beda69c7e7de70600c625b887ce1bc4e1cc1726e309d4896cc6f7d3c4fcae98ac7f9a788e53165e16f004ddb273d21d57976d73e309d4896cc6f7d3c4fcae98ac7f9a788e53165e16f004ddb273d21d57976d73"
	hh, _ := pdf.SplitText(hash, 372)
	_ = pdf.Text("数  据  要  素  特  征  属  性    :")
	pdf.SetX(80)
	pdf.SetY(357)
	_ = pdf.Text("(Data Characteristic Value)")
	lineSpace := 15
	n := len(hh)
	if n > 1 && n <= 27 {
		y := 374
		for _, t := range hh[:] {
			pdf.SetX(81)
			y += lineSpace
			pdf.SetY(float64(y))
			_ = pdf.Text(t)
		}
	} else if n > 27 {
		y := 374
		for _, t := range hh[:13] {
			pdf.SetX(81)
			y += lineSpace
			pdf.SetY(float64(y))
			_ = pdf.Text(t)
		}
		y += lineSpace
		pdf.SetX(81)
		pdf.SetY(float64(y))
		_ = pdf.Text("......")
		for _, t := range hh[n-13:] {
			pdf.SetX(81)
			y += lineSpace
			pdf.SetY(float64(y))
			_ = pdf.Text(t)
		}
	}

	pdf.SetX(259)
	pdf.SetY(792)
	_ = pdf.Text("2/2")

	// log.Debug("path: %s", global.ProductDir()+env.PathSeparator+certName)
	_ = pdf.WritePdf("1.certification.pdf")
}

func genQR(url string, size int) ([]byte, error) {
	c, _ := qr.Encode(url, qr.L)
	realSize := c.Size
	rect := image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{size, size}}
	bgColor := color.Transparent
	fgColor := color.Black
	p := color.Palette([]color.Color{bgColor, fgColor})
	img := image.NewPaletted(rect, p)
	fgClr := uint8(img.Palette.Index(fgColor))
	modulesPerPixel := float64(realSize) / float64(size)
	for y := 0; y < size; y++ {
		y2 := int(float64(y) * modulesPerPixel)
		for x := 0; x < size; x++ {
			x2 := int(float64(x) * modulesPerPixel)
			if c.Black(x2, y2) {
				pos := img.PixOffset(x, y)
				img.Pix[pos] = fgClr
			}
		}
	}
	var b bytes.Buffer
	encoder := png.Encoder{CompressionLevel: png.BestCompression}
	if err := encoder.Encode(&b, img); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
