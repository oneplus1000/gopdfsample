package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	var err error
	err = pdf.AddTTFFont("HDZB_5", "../ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.AddTTFFont("TakaoPGothic", "../ttf/TakaoPGothic.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.AddTTFFont("loma", "../ttf/Loma.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.AddTTFFont("namum", "../ttf/NanumBarunGothic.ttf")

	//china
	err = pdf.SetFont("HDZB_5", "", 14)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	err = pdf.AddTTFFont("Roboto", "../ttf/Roboto-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.Cell(nil, "Hello")
	pdf.Br(20)
	pdf.Cell(nil, "您好")
	pdf.Br(20)

	//japan
	err = pdf.SetFont("TakaoPGothic", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "こんにちは")
	pdf.Br(20)

	//thai
	err = pdf.SetFont("loma", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "สวัสดี")
	pdf.Br(20)

	//test composite glyph
	err = pdf.SetFont("Roboto", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "ÁÉÍÓÖŐÚÜŰ")
	pdf.Br(20)
	pdf.Cell(nil, "ÁÉÍÓÖŐÚÜŰ")
	pdf.Br(20)

	//korean
	err = pdf.SetFont("namum", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "안녕하세요")
	pdf.Br(20)

	pdf.WritePdf("hello2.pdf")

}
