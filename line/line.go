package main

import (
	"log"

	"github.com/signintech/gopdf"
)

/*
Thank https://github.com/miiton for setLineType
*/
func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: gopdf.Unit_PT, PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("TakaoPGothic", "../ttf/TakaoPGothic.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetLineWidth(2)
	pdf.SetLineType("dashed")
	//pdf.SetLineType("dotted")
	pdf.Line(10, 30, 585, 30)

	err = pdf.SetFont("TakaoPGothic", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "こんにちは")
	pdf.Br(20)
	pdf.WritePdf("line.pdf")

}
