package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("Roboto", "../ttf/Roboto-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	fontSize := 14
	err = pdf.SetFont("Roboto", "", fontSize)
	if err != nil {
		log.Print(err.Error())
		return
	}
	//pdf.SetGrayFill(0.5)
	pdf.Cell(nil, "ÁÉÍÓÖŐÚÜŰ")

	pdf.WritePdf("hello.pdf")
}
