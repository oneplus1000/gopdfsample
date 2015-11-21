package main

import (
	"fmt"
	"log"

	"github.com/signintech/gopdf"
	"github.com/signintech/gopdf/fontmaker/core"
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
	fontSize := 24
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

	var parser core.TTFParser
	err = parser.Parse("../ttf/Roboto-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	//get  CapHeight (https://en.wikipedia.org/wiki/Cap_height)
	cap := float64(float64(parser.CapHeight()) * 1000.00 / float64(parser.UnitsPerEm()))
	//convert
	realHeight := cap * (float64(fontSize) / 1000.0)
	fmt.Printf("realHeight = %f", realHeight)
	//test
	pdf.Br(realHeight)
	pdf.Cell(nil, "How can I cordinate the text that I want draw?")
	pdf.Br(realHeight)
	pdf.Cell(nil, "How can I cordinate the text that I want draw?123")
	pdf.Br(realHeight)
	pdf.Cell(nil, "How can I cordinate the text that I want draw?456")

	pdf.WritePdf("hello.pdf")
}