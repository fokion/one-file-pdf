// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-19 23:06:24 540D71                      [utest/draw_unit_grid.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// DrawUnitGrid is the unit test for PDF.DrawUnitGrid()
func DrawUnitGrid(t *testing.T) {
	fmt.Println("utest.DrawUnitGrid")
	//
	var result = func() []byte {
		var pdf = pdf.NewPDF("A4")
		pdf.SetCompression(false).
			SetUnits("cm").
			DrawUnitGrid()
		return pdf.Bytes()
	}()
	const expect = `
	%PDF-1.4
	1 0 obj<</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj<</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
	endobj
	3 0 obj<</Type/Page/Parent 2 0 R\
		/Contents 4 0 R/Resources<</Font <</FNT1 5 0 R>>>>>>
	endobj
	4 0 obj <</Length 7417>>stream
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.100 w
	0.000 841.890 m 0.000 0.000 l S
	BT /FNT1 8 Tf ET
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 833 Td (0) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	28.346 841.890 m 28.346 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 31 833 Td (1) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	56.693 841.890 m 56.693 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 59 833 Td (2) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	85.039 841.890 m 85.039 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 87 833 Td (3) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	113.386 841.890 m 113.386 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 116 833 Td (4) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	141.732 841.890 m 141.732 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 144 833 Td (5) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	170.079 841.890 m 170.079 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 172 833 Td (6) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	198.425 841.890 m 198.425 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 201 833 Td (7) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	226.772 841.890 m 226.772 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 229 833 Td (8) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	255.118 841.890 m 255.118 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 257 833 Td (9) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	283.465 841.890 m 283.465 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 286 833 Td (10) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	311.811 841.890 m 311.811 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 314 833 Td (11) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	340.157 841.890 m 340.157 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 342 833 Td (12) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	368.504 841.890 m 368.504 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 371 833 Td (13) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	396.850 841.890 m 396.850 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 399 833 Td (14) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	425.197 841.890 m 425.197 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 428 833 Td (15) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	453.543 841.890 m 453.543 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 456 833 Td (16) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	481.890 841.890 m 481.890 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 484 833 Td (17) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	510.236 841.890 m 510.236 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 513 833 Td (18) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	538.583 841.890 m 538.583 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 541 833 Td (19) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	566.929 841.890 m 566.929 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 569 833 Td (20) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 841.890 m 595.276 841.890 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 833 Td (0) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 813.543 m 595.276 813.543 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 805 Td (1) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 785.197 m 595.276 785.197 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 776 Td (2) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 756.850 m 595.276 756.850 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 748 Td (3) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 728.504 m 595.276 728.504 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 720 Td (4) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 700.157 m 595.276 700.157 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 691 Td (5) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 671.811 m 595.276 671.811 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 663 Td (6) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 643.465 m 595.276 643.465 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 634 Td (7) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 615.118 m 595.276 615.118 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 606 Td (8) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 586.772 m 595.276 586.772 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 578 Td (9) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 558.425 m 595.276 558.425 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 549 Td (10) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 530.079 m 595.276 530.079 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 521 Td (11) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 501.732 m 595.276 501.732 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 493 Td (12) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 473.386 m 595.276 473.386 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 464 Td (13) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 445.039 m 595.276 445.039 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 436 Td (14) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 416.693 m 595.276 416.693 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 408 Td (15) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 388.346 m 595.276 388.346 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 379 Td (16) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 360.000 m 595.276 360.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 351 Td (17) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 331.654 m 595.276 331.654 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 323 Td (18) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 303.307 m 595.276 303.307 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 294 Td (19) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 274.961 m 595.276 274.961 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 266 Td (20) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 246.614 m 595.276 246.614 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 238 Td (21) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 218.268 m 595.276 218.268 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 209 Td (22) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 189.921 m 595.276 189.921 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 181 Td (23) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 161.575 m 595.276 161.575 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 153 Td (24) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 133.228 m 595.276 133.228 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 124 Td (25) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 104.882 m 595.276 104.882 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 96 Td (26) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 76.535 m 595.276 76.535 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 68 Td (27) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 48.189 m 595.276 48.189 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 39 Td (28) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 19.843 m 595.276 19.843 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 11 Td (29) Tj ET
	endstream
	5 0 obj<</Type/Font/Subtype/Type1/Name/F1/\
		BaseFont/Helvetica/Encoding/WinAnsiEncoding>>
	endobj
	xref
	0 6
	0000000000 65535 f
	0000000009 00000 n
	0000000053 00000 n
	0000000125 00000 n
	0000000217 00000 n
	0000007676 00000 n
	trailer
	<</Size 6/Root 1 0 R>>
	startxref
	7771
	%%EOF
	`
	pdfCompare(t, result, expect)
} //                                                                DrawUnitGrid

//end
