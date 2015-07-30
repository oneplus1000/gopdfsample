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
	err = pdf.AddTTFFont("HDZB_5", "../ttf/wts11.ttf")
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	err = pdf.AddTTFFont("osaka", "../ttf/osaka.unicode.ttf")
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
		return
	}

	err = pdf.AddTTFFont("loma", "../ttf/Loma.ttf")
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
		return
	}

	//china
	err = pdf.SetFont("HDZB_5", "", 14)
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
		return
	}
	pdf.Cell(nil, "Hello")
	pdf.Br(20)
	pdf.Cell(nil, "您好")
	pdf.Br(20)

	//japan
	err = pdf.SetFont("osaka", "", 14)
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
		return
	}
	pdf.Cell(nil, "こんにちは")
	pdf.Br(20)

	//thai
	err = pdf.SetFont("loma", "", 14)
	if err != nil {
		log.Printf("ERROR:%s\n", err.Error())
		return
	}
	pdf.Cell(nil, "สวัสดี")

	pdf.WritePdf("hello2.pdf")
}
