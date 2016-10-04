package main

import (
	"log"

	"github.com/signintech/gopdf"
)

/*
Thank https://github.com/miiton for write text color supported.
*/
func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("HDZB_5", "../ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("HDZB_5", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	//
	pdf.SetGrayFill(0.5)
	pdf.Cell(nil, "gray")

	pdf.SetTextColor(255, 0, 0)
	pdf.Br(20)
	pdf.Cell(nil, "red")

	pdf.SetTextColor(0, 0, 0)
	pdf.Br(20)
	pdf.Cell(nil, "gray")

	pdf.SetGrayFill(0)
	pdf.Br(20)
	pdf.Cell(nil, "black")

	pdf.SetTextColor(0, 0, 255)
	pdf.Br(20)
	pdf.Cell(nil, "blue")

	pdf.WritePdf("color.pdf")
}
