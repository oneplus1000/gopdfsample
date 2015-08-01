package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("osaka", "../ttf/osaka.unicode.ttf")
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
		return
	}
	pdf.SetLineWidth(2)
	pdf.Line(10, 30, 585, 30)

	err = pdf.SetFont("osaka", "", 14)
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
		return
	}
	pdf.Cell(nil, "こんにちは")
	pdf.Br(20)
	pdf.WritePdf("hello2.pdf")

}
