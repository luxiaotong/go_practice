package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/webengine"
	"github.com/therecipe/qt/widgets"
)

func main() {
	in := flag.String("uri", "https://www.fyianlai.com/", "a URI to convert to PDF")
	out := flag.String("out", "output.pdf", "a file path for the generated PDF output")
	flag.Parse()

	start := time.Now()

	app := widgets.NewQApplication(len(os.Args), os.Args)

	view := webengine.NewQWebEngineView(nil)
	page := view.Page()

	view.ConnectLoadStarted(func() {
		fmt.Printf("loading : %q\n", *in)
	})

	view.ConnectLoadFinished(func(ok bool) {
		defer fmt.Printf("PDF generated : %q (%s)\n", *out, time.Since(start))
		fmt.Printf("load finished : %t\n", ok)
		layout := gui.NewQPageLayout2(gui.NewQPageSize2(0), 0, core.NewQMarginsF2(0.0, 0.0, 0.0, 0.0), 0, core.NewQMarginsF2(0.0, 0.0, 0.0, 0.0))

		page.PrintToPdf(*out, layout)

		// app.Exit(0)
	})

	view.Load(core.NewQUrl3(*in, 0))
	app.Exec()
}
