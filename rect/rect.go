package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	//pdf.AddPageWithOption(gopdf.PageOption{PageSize: gopdf.Rect{W: 841.89, H: 595.28}})
	err := pdf.AddTTFFont("wts11", "../ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	rectFillColor(&pdf, "Play", 14, 10, 10, 100, 100, 255, 0, 255, alignRight, valignBottom)
	rectFillColor(&pdf, "Hello", 14, 10, 130, 100, 100, 205, 0, 255, alignCenter, valignTop)
	rectFillColor(&pdf, "Nano", 14, 10, 250, 100, 100, 155, 0, 255, alignCenter, valignMiddle)
	pdf.WritePdf("rect.pdf")
}

const (
	valignTop    = 1
	valignMiddle = 2
	valignBottom = 3
)

const (
	alignLeft   = 4
	alignCenter = 5
	alignRight  = 6
)

func rectFillColor(pdf *gopdf.GoPdf,
	text string,
	fontSize int,
	x, y, w, h float64,
	r, g, b uint8,
	align, valign int,
) {
	pdf.SetLineWidth(0.1)
	pdf.SetFillColor(r, g, b) //setup fill color
	pdf.RectFromUpperLeftWithStyle(x, y, w, h, "FD")
	pdf.SetFillColor(0, 0, 0)

	if align == alignCenter {
		textw, _ := pdf.MeasureTextWidth(text)
		x = x + (w / 2) - (textw / 2)
	} else if align == alignRight {
		textw, _ := pdf.MeasureTextWidth(text)
		x = x + w - textw
	}

	pdf.SetX(x)

	if valign == valignMiddle {
		y = y + (h / 2) - (float64(fontSize) / 2)
	} else if valign == valignBottom {
		y = y + h - float64(fontSize)
	}

	pdf.SetY(y)
	pdf.Cell(nil, text)
}
