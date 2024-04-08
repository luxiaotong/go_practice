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
		Protection: gopdf.PDFProtectionConfig{
			UseProtection: true,
			Permissions:   gopdf.PermissionsPrint,
			// OwnerPass:     []byte("123456"),
			// UserPass:      []byte("123456789"),
		},
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

	_ = pdf.Image(Logo, 299, 70, &gopdf.Rect{W: 236, H: 25})
	pdf.SetX(82)
	pdf.SetY(150)
	_ = pdf.SetFont("Heavy", "", 37)
	pdf.SetTextColor(6, 63, 131)
	_ = pdf.Text("数  据  要  素  登  记  证  书")
	pdf.SetX(61)
	pdf.SetY(86)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.SetFont("Regular", "", 14)
	_ = pdf.Text(fmt.Sprintf("证书编号:%s", "MA388608629841317"))
	pdf.SetX(76)
	pdf.SetY(202)
	_ = pdf.SetFont("Regular", "", 13)
	_ = pdf.Text(fmt.Sprintf("数  据  要  素  名  称 : %s", "河南省新乡市统计局数据2"))
	pdf.SetX(76)
	pdf.SetY(250)
	_ = pdf.Text(fmt.Sprintf("数据要素所有单位  : %s", " 测试单位名称777"))
	pdf.SetX(76)
	pdf.SetY(298)
	_ = pdf.Text(fmt.Sprintf("数  据  要  素  标  识 : %s", "MA.10000.900000.00000000/000070289717"))
	pdf.SetX(76)
	pdf.SetY(346)
	_ = pdf.Text(printer.Sprintf("数  据  要  素  总  量 : %d DRs", 64480))
	pdf.SetX(76)
	pdf.SetY(394)
	sources := []string{"程序爬虫取得", "其他，请注明testtesttest"}
	ss, _ := pdf.SplitText(strings.Join(sources, ";"), 327)
	_ = pdf.Text(fmt.Sprintf("获     取     方     式     : %s", ss[0]))
	if len(ss) > 1 {
		y := 418 // (394+442)/2
		for i, t := range ss[1:] {
			pdf.SetX(192)
			pdf.SetY(float64(y + i*15))
			_ = pdf.Text(t)
		}
	}
	pdf.SetX(76)
	pdf.SetY(442)
	_ = pdf.Text(fmt.Sprintf("申     请     时     间     : %s", time.Now().Format("2006年1月2日")))
	pdf.SetX(76)
	pdf.SetY(490)
	_ = pdf.Text("审     核     机     构     : 中国工业互联网研究院")
	pdf.SetX(76)
	pdf.SetY(538)
	_ = pdf.Text(fmt.Sprintf("审     核     时     间     : %s", time.Now().Format("2006年1月2日")))
	pdf.SetX(76)
	pdf.SetY(586)
	_ = pdf.Text(fmt.Sprintf("区     块       I       D     : %s", "45966274826631ae1c680fad0caeb1f47441ac9c"))
	pdf.SetX(76)
	pdf.SetY(634)
	_ = pdf.Text(fmt.Sprintf("生  成  区  块  时  间 : %s", time.Now().Format("2006年1月2日 15:04:05")))
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
	_ = pdf.Image(Stamp, 414, 653, &gopdf.Rect{W: 119, H: 119}) // TODO add province_id
	pdf.SetX(259)
	pdf.SetY(792)
	_ = pdf.Text("第 1 页（共 2 页）")

	// Add New Page
	pdf.AddPage()
	pdf.UseImportedTemplate(tplBg, 0, 0, 0, 0)

	pdf.SetX(80)
	pdf.SetY(130)
	_ = pdf.SetFont("Heavy", "", 37)
	pdf.SetTextColor(6, 63, 131)
	_ = pdf.Text("数  据  概  况")
	pdf.SetX(56)
	pdf.SetY(64)
	_ = pdf.SetFont("Regular", "", 10)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.Text(fmt.Sprintf("证书编号:%s", "MA388608629841317"))
	pdf.SetX(80)
	pdf.SetY(190)
	_ = pdf.SetFont("Regular", "", 13)
	_ = pdf.Text(fmt.Sprintf("数   据   表   数   量   : %d 张", 2))
	pdf.SetX(80)
	pdf.SetY(240)
	_ = pdf.Text(printer.Sprintf("数  据  字  段  数  量 : %d 个", 20))
	pdf.SetX(80)
	pdf.SetY(290)
	_ = pdf.Text(fmt.Sprintf("有  效  记  录  占  比 : %.2f%%", 100.00))

	pdf.SetX(259)
	pdf.SetY(792)
	_ = pdf.Text("第 2 页（共 2 页）")

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
