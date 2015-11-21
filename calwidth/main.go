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
	//pdf.Cell(nil, "Áa")

	text := "How can I cordinate the text that I want draw?"
	pdf.Cell(nil, text)
	sfont := pdf.Curr.Font_ISubset
	sumWidth := uint64(0)
	for _, r := range text {
		width, err := sfont.CharWidth(r)
		if err != nil {
			log.Print(err.Error())
			return
		}
		sumWidth = sumWidth + width
	}
	realWidth := float64(sumWidth) * (float64(fontSize) / 1000.0) //convert
	fmt.Printf("realWidth = %f", realWidth)

	pdf.WritePdf("hello.pdf")
}
