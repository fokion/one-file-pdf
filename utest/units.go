// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-22 02:55:11 8BF1F2                               [utest/units.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

func Units(t *testing.T) {

	// (pdf *PDF) Units() string
	//
	fmt.Println("utest.Units")

	func() {
		var ob pdf.PDF // uninitialized PDF
		TEqual(t, ob.Units(), "")
	}()

	func() {
		var ob = pdf.NewPDF("A4")
		TEqual(t, ob.Units(), "POINT")
	}()

	// (pdf *PDF) SetUnits(unitName string) *PDF
	//
	fmt.Println("utest.SetUnits")

	func() {
		var ob = pdf.NewPDF("A4")
		TEqual(t, len(ob.Errors()), 0)
		ob.SetUnits("cm")
		TEqual(t, len(ob.Errors()), 0)
		TEqual(t, ob.Units(), "CM")
	}()

	func() {
		var ob = pdf.NewPDF("A4")
		TEqual(t, len(ob.Errors()), 0)
		ob.SetUnits("fathoms")
		TEqual(t, len(ob.Errors()), 1)
		//
		if len(ob.Errors()) == 1 {
			TEqual(t,
				ob.Errors()[0],
				fmt.Errorf(`Unknown unit name "fathoms" @SetUnits`))
		}
		TEqual(t, ob.Units(), "POINT")
	}()

} //                                                                       Units

//end
