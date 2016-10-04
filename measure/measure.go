package main

import (
	"fmt"
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
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
	pdf.SetX(0)
	text01 := "こんにちは"
	pdf.Cell(nil, text01)
	w01, _ := pdf.MeasureTextWidth(text01)
	pdf.SetY(20)
	pdf.SetX(w01)

	text02 := "i am a man."
	pdf.Cell(nil, text02)
	w02, _ := pdf.MeasureTextWidth(text02)
	pdf.SetY(30)
	pdf.SetX(w01 + w02)

	text03 := "done"
	pdf.Cell(nil, text03)

	fmt.Printf("MeasureTextWidth = %f\n", w01)
	pdf.WritePdf("m.pdf")
}
