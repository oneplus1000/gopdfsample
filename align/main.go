package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {

	pdf := gopdf.GoPdf{}
	docWidth := 590.0
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: docWidth, H: 840}})
	pdf.AddPage()

	// Add the font
	err := pdf.AddTTFFont("times-bold", "../ttf/times.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("times-bold", "", 15)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.CellWithOption(&gopdf.Rect{
		W: 200,
		H: 200,
	}, "Center", gopdf.CellOption{Align: gopdf.Center,
		Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,
	})

	pdf.SetX(10)
	pdf.SetY(250)

	pdf.CellWithOption(&gopdf.Rect{
		W: 200,
		H: 200,
	}, "Left", gopdf.CellOption{Align: gopdf.Left,
		Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,
	})

	pdf.SetX(10)
	pdf.SetY(500)

	pdf.CellWithOption(&gopdf.Rect{
		W: 200,
		H: 200,
	}, "Center+Middle", gopdf.CellOption{Align: gopdf.Center | gopdf.Middle,
		Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,
	})

	pdf.WritePdf("center_text.pdf")

}
