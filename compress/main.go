package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4

	pdf.SetCompressLevel(1) //compress content streams

	pdf.AddPage()
	err := pdf.AddTTFFont("wts", "../ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("wts", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetGrayFill(0.5)
	pdf.Cell(nil, "您好 Hello")
	/*
		pdf.Image("../imgs/gopher.jpg", 200, 50, nil)
		err = pdf.SetFont("wts", "", 14)
		if err != nil {
			log.Print(err.Error())
			return
		}
	*/
	pdf.WritePdf("compress.pdf")
}
