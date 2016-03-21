package main

import (
	"fmt"
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("TakaoPGothic", "../ttf/TakaoPGothic.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("TakaoPGothic", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	i := 0
	for i < 2 {
		text01 := "こんにちは"
		measureTextWidth, err := pdf.MeasureTextWidth(text01)
		if err != nil {
			log.Printf("Err:%s\n", err.Error())
			return
		}
		fmt.Printf("MeasureTextWidth = %f\n", measureTextWidth)

		pdf.Cell(nil, text01)
		pdf.SetX(pdf.GetX() + measureTextWidth)
		pdf.Cell(nil, text01)
		pdf.SetX(pdf.GetX() + measureTextWidth)
		i++
	}
	pdf.WritePdf("m.pdf")
}
