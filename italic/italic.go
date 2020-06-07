package main

import (
	"log"

	"github.com/signintech/gopdf"
)

//gopdf had no mechanism for set italic ( pdf.SetFont("SOMEFONT", "I", 14) ), but you can do like this.
func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: gopdf.Unit_PT, PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()

	err := pdf.AddTTFFont("DejaVuSerif", "../ttf/DejaVuSerif.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.AddTTFFont("DejaVuSerif-Italic", "../ttf/DejaVuSerif-Italic.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("DejaVuSerif", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "Hi! This is nomal.")
	pdf.Br(20)

	err = pdf.SetFont("DejaVuSerif-Italic", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "Hi! This is italic.")
	pdf.WritePdf("italic.pdf")
}
