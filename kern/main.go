package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFontWithOption("times", "../ttf/times.ttf", gopdf.TtfOption{
		UseKerning: true,
	})
	err = pdf.AddTTFFontWithOption("times2", "../ttf/times.ttf", gopdf.TtfOption{
		UseKerning: false,
	})

	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("times", "", 40)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "WowAVAVWA")

	err = pdf.SetFont("times", "", 12)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "     with kerning")

	pdf.Br(30)

	err = pdf.SetFont("times2", "", 40)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "WowAVAVWA")
	err = pdf.SetFont("times", "", 12)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "      without kerning")

	pdf.WritePdf("kern.pdf")
	//test
}
