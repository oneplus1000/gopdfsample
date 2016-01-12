package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	var err error
	err = pdf.AddTTFFont("loma", "../ttf/Loma.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("loma", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "hello")
	//#1 pic
	pdf.Image("../imgs/gopher.jpg", 200, 50, nil)

	pdf.SetX(250)
	pdf.SetY(200)
	pdf.Cell(nil, "gopher and gopher")

	pdf.WritePdf("image.pdf")
}
