package main

import (
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
	fontSize := 24
	err = pdf.SetFont("Roboto", "", fontSize)
	if err != nil {
		log.Print(err.Error())
		return
	}

	height, err := pdf.CurrentFontCapHeight()
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "Sphinx")
	pdf.Br(height)
	pdf.Cell(nil, "Sphinx2")
	pdf.Br(height)
	pdf.Cell(nil, "Sphinx2")
	pdf.Br(height)
	pdf.Cell(nil, "Sphinx2")
	pdf.Br(height)
	pdf.Cell(nil, "Sphinx2")
	pdf.Br(height)
	pdf.Cell(nil, "Sphinx2")
	pdf.Br(height)
	pdf.Cell(nil, "Sphinx2")
	pdf.Br(height)
	pdf.Cell(nil, "Sphinx2")
	pdf.Br(height)
	pdf.Cell(nil, "Sphinx2")
	pdf.Br(height)
	pdf.Cell(nil, "Sphinx2")
	pdf.Br(height)
	pdf.Cell(nil, "Sphinx2")

	pdf.WritePdf("hello.pdf")
}
