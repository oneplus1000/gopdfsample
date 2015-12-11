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

	text01 := "こんにちは"
	pdf.Cell(nil, text01)
	w01, _ := pdf.MeasureTextWidth(text01)
	fmt.Printf("MeasureTextWidth = %f\n", w01)
	pdf.WritePdf("m.pdf")
}
