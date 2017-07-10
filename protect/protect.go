package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}

	pdf.Start(gopdf.Config{
		PageSize: gopdf.Rect{W: 595.28, H: 841.89}, //595.28, 841.89 = A4
		Protection: gopdf.PDFProtectionConfig{
			UseProtection: true,
			Permissions:   gopdf.PermissionsPrint | gopdf.PermissionsCopy | gopdf.PermissionsModify,
			OwnerPass:     []byte("1234"),
			UserPass:      []byte("5555")},
	})

	pdf.AddPage()
	err := pdf.AddTTFFont("TakaoPGothic", "../ttf/TakaoPGothic.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetLineWidth(2)
	pdf.SetLineType("dashed")
	pdf.SetLineType("dotted")
	pdf.Line(10, 30, 585, 30)
	err = pdf.SetFont("TakaoPGothic", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	//#1 pic
	pdf.Image("../imgs/gopher.jpg", 200, 50, nil)

	pdf.Cell(nil, "こんにちは")
	pdf.Br(20)
	pdf.WritePdf("protect.pdf")
}
