package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4

	pdf.AddPage()
	err := pdf.AddTTFFont("wts11", "../ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.Cell(nil, "A")
	//pdf.AddPageWithOption(gopdf.PageOption{PageSize: gopdf.Rect{W: 595.28, H: 841.89}})
	pdf.AddPageWithOption(gopdf.PageOption{PageSize: &gopdf.Rect{W: 841.89, H: 595.28}})
	//pdf.AddPage()
	pdf.Cell(nil, "BBBBBBB")
	pdf.Br(10)
	pdf.Cell(nil, "BBBBBBBB")
	pdf.AddPage()
	pdf.Cell(nil, "C")

	pdf.WritePdf("pages.pdf")
}
