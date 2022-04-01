package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	notification := "1509484018444275712.notification.pdf"
	signature := "signature.png"
	dst := "1509484018444275712.signature.pdf"
	x := float64(378)
	y := float64(479)
	w := float64(67)
	h := float64(40)
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeLetter,
	})
	tpl := pdf.ImportPage(notification, 1, "/MediaBox")
	pdf.AddPage()
	pdf.UseImportedTemplate(tpl, 0, 0, 0, 0)
	if err := pdf.Image(signature, x, y, &gopdf.Rect{W: w, H: h}); err != nil {
		log.Println(err)
		return
	}
	// log.Debug("dst: %s", dst)
	if err := pdf.WritePdf(dst); err != nil {
		log.Println(err)
		return
	}
}
