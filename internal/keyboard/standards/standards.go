/*
Package standards aggregates per-standard keyboard definitions, physical key
matrices for ANSI, ISO, ABNT, JIS, and KS, as well as their layout maps
(kana for JIS, hangeul for KS), accessible via the All map.
*/
package standards

import "github.com/arvingarciabtw/ditto/internal/keyboard/base"

type Data struct {
	Sizes    map[int][][]base.Key
	ShiftMap map[string]string
	AltGrMap map[string]string
}

var All = map[string]Data{
	"ansi": {Sizes: SizesANSI, ShiftMap: base.USShift},
	"iso":  {Sizes: SizesISO, ShiftMap: base.UKShift},
	"abnt": {Sizes: SizesABNT, ShiftMap: base.ABNTShift, AltGrMap: base.ABNTAltGr},
	"jis":  {Sizes: SizesJIS, ShiftMap: base.JISShift},
	"ks":   {Sizes: SizesKS, ShiftMap: base.KSShift},
}

var ListItems = []string{
	"ansi",
	"iso",
	"abnt",
	"jis",
	"ks",
}
