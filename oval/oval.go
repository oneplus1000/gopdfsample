package main

import "github.com/signintech/gopdf"

/*
Thank https://github.com/miiton for write draw oval.
*/
func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	pdf.SetLineWidth(1)
	pdf.Oval(100, 200, 500, 500)
	pdf.WritePdf("oval.pdf")
	//fmt.Printf("")

}
