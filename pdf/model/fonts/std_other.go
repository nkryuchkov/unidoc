/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */
/*
 * The embedded character metrics specified in this file are distributed under the terms listed in
 * ./testdata/afms/MustRead.html.
 */

package fonts

import (
	"github.com/unidoc/unidoc/pdf/internal/textencoding"
)

func init() {
	// The aliases seen for the standard 14 font names.
	// Most of these are from table 5.5.1 in
	// https://www.adobe.com/content/dam/acom/en/devnet/acrobat/pdfs/adobe_supplement_iso32000.pdf
	RegisterStdFont(SymbolName, NewFontSymbol, "Symbol,Italic", "Symbol,Bold", "Symbol,BoldItalic")
	RegisterStdFont(ZapfDingbatsName, NewFontZapfDingbats)
}

const (
	// SymbolName is a PDF name of the Symbol font.
	SymbolName = StdFontName("Symbol")
	// ZapfDingbatsName is a PDF name of the ZapfDingbats font.
	ZapfDingbatsName = StdFontName("ZapfDingbats")
)

// NewFontSymbol returns a new instance of the font with a default encoder set (SymbolEncoder).
func NewFontSymbol() StdFont {
	enc := textencoding.NewSymbolEncoder()
	desc := Descriptor{
		Name:        SymbolName,
		Family:      string(SymbolName),
		Weight:      FontWeightMedium,
		Flags:       0x0004,
		BBox:        [4]float64{-180, -293, 1090, 1010},
		ItalicAngle: 0,
		Ascent:      0,
		Descent:     0,
		CapHeight:   0,
		XHeight:     0,
		StemV:       85,
		StemH:       92,
	}
	return NewStdFontWithEncoding(desc, symbolCharMetrics, enc)
}

// NewFontZapfDingbats returns a new instance of the font with a default encoder set (ZapfDingbatsEncoder).
func NewFontZapfDingbats() StdFont {
	enc := textencoding.NewZapfDingbatsEncoder()
	desc := Descriptor{
		Name:        ZapfDingbatsName,
		Family:      string(ZapfDingbatsName),
		Weight:      FontWeightMedium,
		Flags:       0x0004,
		BBox:        [4]float64{-1, -143, 981, 820},
		ItalicAngle: 0,
		Ascent:      0,
		Descent:     0,
		CapHeight:   0,
		XHeight:     0,
		StemV:       90,
		StemH:       28,
	}
	return NewStdFontWithEncoding(desc, zapfDingbatsCharMetrics, enc)
}

// symbolCharMetrics are the font metrics loaded from afms/Symbol.afm.
// See afms/MustRead.html for license information.
var symbolCharMetrics = map[rune]CharMetrics{
	' ':      {Wx: 250},
	'!':      {Wx: 333},
	'#':      {Wx: 500},
	'%':      {Wx: 833},
	'&':      {Wx: 778},
	'(':      {Wx: 333},
	')':      {Wx: 333},
	'+':      {Wx: 549},
	',':      {Wx: 250},
	'.':      {Wx: 250},
	'/':      {Wx: 278},
	'0':      {Wx: 500},
	'1':      {Wx: 500},
	'2':      {Wx: 500},
	'3':      {Wx: 500},
	'4':      {Wx: 500},
	'5':      {Wx: 500},
	'6':      {Wx: 500},
	'7':      {Wx: 500},
	'8':      {Wx: 500},
	'9':      {Wx: 500},
	':':      {Wx: 278},
	';':      {Wx: 278},
	'<':      {Wx: 549},
	'=':      {Wx: 549},
	'>':      {Wx: 549},
	'?':      {Wx: 444},
	'[':      {Wx: 333},
	']':      {Wx: 333},
	'_':      {Wx: 500},
	'{':      {Wx: 480},
	'|':      {Wx: 200},
	'}':      {Wx: 480},
	'¬':      {Wx: 713},
	'°':      {Wx: 400},
	'±':      {Wx: 549},
	'µ':      {Wx: 576},
	'×':      {Wx: 549},
	'÷':      {Wx: 549},
	'ƒ':      {Wx: 500},
	'Α':      {Wx: 722},
	'Β':      {Wx: 667},
	'Γ':      {Wx: 603},
	'Ε':      {Wx: 611},
	'Ζ':      {Wx: 611},
	'Η':      {Wx: 722},
	'Θ':      {Wx: 741},
	'Ι':      {Wx: 333},
	'Κ':      {Wx: 722},
	'Λ':      {Wx: 686},
	'Μ':      {Wx: 889},
	'Ν':      {Wx: 722},
	'Ξ':      {Wx: 645},
	'Ο':      {Wx: 722},
	'Π':      {Wx: 768},
	'Ρ':      {Wx: 556},
	'Σ':      {Wx: 592},
	'Τ':      {Wx: 611},
	'Υ':      {Wx: 690},
	'Φ':      {Wx: 763},
	'Χ':      {Wx: 722},
	'Ψ':      {Wx: 795},
	'α':      {Wx: 631},
	'β':      {Wx: 549},
	'γ':      {Wx: 411},
	'δ':      {Wx: 494},
	'ε':      {Wx: 439},
	'ζ':      {Wx: 494},
	'η':      {Wx: 603},
	'θ':      {Wx: 521},
	'ι':      {Wx: 329},
	'κ':      {Wx: 549},
	'λ':      {Wx: 549},
	'ν':      {Wx: 521},
	'ξ':      {Wx: 493},
	'ο':      {Wx: 549},
	'π':      {Wx: 549},
	'ρ':      {Wx: 549},
	'ς':      {Wx: 439},
	'σ':      {Wx: 603},
	'τ':      {Wx: 439},
	'υ':      {Wx: 576},
	'φ':      {Wx: 521},
	'χ':      {Wx: 549},
	'ψ':      {Wx: 686},
	'ω':      {Wx: 686},
	'ϑ':      {Wx: 631},
	'ϒ':      {Wx: 620},
	'ϕ':      {Wx: 603},
	'ϖ':      {Wx: 713},
	'•':      {Wx: 460},
	'…':      {Wx: 1000},
	'′':      {Wx: 247},
	'″':      {Wx: 411},
	'⁄':      {Wx: 167},
	'€':      {Wx: 750},
	'ℑ':      {Wx: 686},
	'℘':      {Wx: 987},
	'ℜ':      {Wx: 795},
	'Ω':      {Wx: 768},
	'ℵ':      {Wx: 823},
	'←':      {Wx: 987},
	'↑':      {Wx: 603},
	'→':      {Wx: 987},
	'↓':      {Wx: 603},
	'↔':      {Wx: 1042},
	'↵':      {Wx: 658},
	'⇐':      {Wx: 987},
	'⇑':      {Wx: 603},
	'⇒':      {Wx: 987},
	'⇓':      {Wx: 603},
	'⇔':      {Wx: 1042},
	'∀':      {Wx: 713},
	'∂':      {Wx: 494},
	'∃':      {Wx: 549},
	'∅':      {Wx: 823},
	'∆':      {Wx: 612},
	'∇':      {Wx: 713},
	'∈':      {Wx: 713},
	'∉':      {Wx: 713},
	'∋':      {Wx: 439},
	'∏':      {Wx: 823},
	'∑':      {Wx: 713},
	'−':      {Wx: 549},
	'∗':      {Wx: 500},
	'√':      {Wx: 549},
	'∝':      {Wx: 713},
	'∞':      {Wx: 713},
	'∠':      {Wx: 768},
	'∧':      {Wx: 603},
	'∨':      {Wx: 603},
	'∩':      {Wx: 768},
	'∪':      {Wx: 768},
	'∫':      {Wx: 274},
	'∴':      {Wx: 863},
	'∼':      {Wx: 549},
	'≅':      {Wx: 549},
	'≈':      {Wx: 549},
	'≠':      {Wx: 549},
	'≡':      {Wx: 549},
	'≤':      {Wx: 549},
	'≥':      {Wx: 549},
	'⊂':      {Wx: 713},
	'⊃':      {Wx: 713},
	'⊄':      {Wx: 713},
	'⊆':      {Wx: 713},
	'⊇':      {Wx: 713},
	'⊕':      {Wx: 768},
	'⊗':      {Wx: 768},
	'⊥':      {Wx: 658},
	'⋅':      {Wx: 250},
	'⌠':      {Wx: 686},
	'⌡':      {Wx: 686},
	'〈':      {Wx: 329},
	'〉':      {Wx: 329},
	'◊':      {Wx: 494},
	'♠':      {Wx: 753},
	'♣':      {Wx: 753},
	'♥':      {Wx: 753},
	'♦':      {Wx: 753},
	'\uf6d9': {Wx: 790},
	'\uf6da': {Wx: 790},
	'\uf6db': {Wx: 890},
	'\uf8e5': {Wx: 500},
	'\uf8e6': {Wx: 603},
	'\uf8e7': {Wx: 1000},
	'\uf8e8': {Wx: 790},
	'\uf8e9': {Wx: 790},
	'\uf8ea': {Wx: 786},
	'\uf8eb': {Wx: 384},
	'\uf8ec': {Wx: 384},
	'\uf8ed': {Wx: 384},
	'\uf8ee': {Wx: 384},
	'\uf8ef': {Wx: 384},
	'\uf8f0': {Wx: 384},
	'\uf8f1': {Wx: 494},
	'\uf8f2': {Wx: 494},
	'\uf8f3': {Wx: 494},
	'\uf8f4': {Wx: 494},
	'\uf8f5': {Wx: 686},
	'\uf8f6': {Wx: 384},
	'\uf8f7': {Wx: 384},
	'\uf8f8': {Wx: 384},
	'\uf8f9': {Wx: 384},
	'\uf8fa': {Wx: 384},
	'\uf8fb': {Wx: 384},
	'\uf8fc': {Wx: 494},
	'\uf8fd': {Wx: 494},
	'\uf8fe': {Wx: 494},
	'\uf8ff': {Wx: 790},
}

// zapfDingbatsCharMetrics are the font metrics loaded from afms/ZapfDingbats.afm.
// See afms/MustRead.html for license information.
var zapfDingbatsCharMetrics = map[rune]CharMetrics{
	' ':      {Wx: 278},
	'→':      {Wx: 838},
	'↔':      {Wx: 1016},
	'↕':      {Wx: 458},
	'①':      {Wx: 788},
	'②':      {Wx: 788},
	'③':      {Wx: 788},
	'④':      {Wx: 788},
	'⑤':      {Wx: 788},
	'⑥':      {Wx: 788},
	'⑦':      {Wx: 788},
	'⑧':      {Wx: 788},
	'⑨':      {Wx: 788},
	'⑩':      {Wx: 788},
	'■':      {Wx: 761},
	'▲':      {Wx: 892},
	'▼':      {Wx: 892},
	'◆':      {Wx: 788},
	'●':      {Wx: 791},
	'◗':      {Wx: 438},
	'★':      {Wx: 816},
	'☎':      {Wx: 719},
	'☛':      {Wx: 960},
	'☞':      {Wx: 939},
	'♠':      {Wx: 626},
	'♣':      {Wx: 776},
	'♥':      {Wx: 694},
	'♦':      {Wx: 595},
	'✁':      {Wx: 974},
	'✂':      {Wx: 961},
	'✃':      {Wx: 974},
	'✄':      {Wx: 980},
	'✆':      {Wx: 789},
	'✇':      {Wx: 790},
	'✈':      {Wx: 791},
	'✉':      {Wx: 690},
	'✌':      {Wx: 549},
	'✍':      {Wx: 855},
	'✎':      {Wx: 911},
	'✏':      {Wx: 933},
	'✐':      {Wx: 911},
	'✑':      {Wx: 945},
	'✒':      {Wx: 974},
	'✓':      {Wx: 755},
	'✔':      {Wx: 846},
	'✕':      {Wx: 762},
	'✖':      {Wx: 761},
	'✗':      {Wx: 571},
	'✘':      {Wx: 677},
	'✙':      {Wx: 763},
	'✚':      {Wx: 760},
	'✛':      {Wx: 759},
	'✜':      {Wx: 754},
	'✝':      {Wx: 494},
	'✞':      {Wx: 552},
	'✟':      {Wx: 537},
	'✠':      {Wx: 577},
	'✡':      {Wx: 692},
	'✢':      {Wx: 786},
	'✣':      {Wx: 788},
	'✤':      {Wx: 788},
	'✥':      {Wx: 790},
	'✦':      {Wx: 793},
	'✧':      {Wx: 794},
	'✩':      {Wx: 823},
	'✪':      {Wx: 789},
	'✫':      {Wx: 841},
	'✬':      {Wx: 823},
	'✭':      {Wx: 833},
	'✮':      {Wx: 816},
	'✯':      {Wx: 831},
	'✰':      {Wx: 923},
	'✱':      {Wx: 744},
	'✲':      {Wx: 723},
	'✳':      {Wx: 749},
	'✴':      {Wx: 790},
	'✵':      {Wx: 792},
	'✶':      {Wx: 695},
	'✷':      {Wx: 776},
	'✸':      {Wx: 768},
	'✹':      {Wx: 792},
	'✺':      {Wx: 759},
	'✻':      {Wx: 707},
	'✼':      {Wx: 708},
	'✽':      {Wx: 682},
	'✾':      {Wx: 701},
	'✿':      {Wx: 826},
	'❀':      {Wx: 815},
	'❁':      {Wx: 789},
	'❂':      {Wx: 789},
	'❃':      {Wx: 707},
	'❄':      {Wx: 687},
	'❅':      {Wx: 696},
	'❆':      {Wx: 689},
	'❇':      {Wx: 786},
	'❈':      {Wx: 787},
	'❉':      {Wx: 713},
	'❊':      {Wx: 791},
	'❋':      {Wx: 785},
	'❍':      {Wx: 873},
	'❏':      {Wx: 762},
	'❐':      {Wx: 762},
	'❑':      {Wx: 759},
	'❒':      {Wx: 759},
	'❖':      {Wx: 784},
	'❘':      {Wx: 138},
	'❙':      {Wx: 277},
	'❚':      {Wx: 415},
	'❛':      {Wx: 392},
	'❜':      {Wx: 392},
	'❝':      {Wx: 668},
	'❞':      {Wx: 668},
	'❡':      {Wx: 732},
	'❢':      {Wx: 544},
	'❣':      {Wx: 544},
	'❤':      {Wx: 910},
	'❥':      {Wx: 667},
	'❦':      {Wx: 760},
	'❧':      {Wx: 760},
	'❶':      {Wx: 788},
	'❷':      {Wx: 788},
	'❸':      {Wx: 788},
	'❹':      {Wx: 788},
	'❺':      {Wx: 788},
	'❻':      {Wx: 788},
	'❼':      {Wx: 788},
	'❽':      {Wx: 788},
	'❾':      {Wx: 788},
	'❿':      {Wx: 788},
	'➀':      {Wx: 788},
	'➁':      {Wx: 788},
	'➂':      {Wx: 788},
	'➃':      {Wx: 788},
	'➄':      {Wx: 788},
	'➅':      {Wx: 788},
	'➆':      {Wx: 788},
	'➇':      {Wx: 788},
	'➈':      {Wx: 788},
	'➉':      {Wx: 788},
	'➊':      {Wx: 788},
	'➋':      {Wx: 788},
	'➌':      {Wx: 788},
	'➍':      {Wx: 788},
	'➎':      {Wx: 788},
	'➏':      {Wx: 788},
	'➐':      {Wx: 788},
	'➑':      {Wx: 788},
	'➒':      {Wx: 788},
	'➓':      {Wx: 788},
	'➔':      {Wx: 894},
	'➘':      {Wx: 748},
	'➙':      {Wx: 924},
	'➚':      {Wx: 748},
	'➛':      {Wx: 918},
	'➜':      {Wx: 927},
	'➝':      {Wx: 928},
	'➞':      {Wx: 928},
	'➟':      {Wx: 834},
	'➠':      {Wx: 873},
	'➡':      {Wx: 828},
	'➢':      {Wx: 924},
	'➣':      {Wx: 924},
	'➤':      {Wx: 917},
	'➥':      {Wx: 930},
	'➦':      {Wx: 931},
	'➧':      {Wx: 463},
	'➨':      {Wx: 883},
	'➩':      {Wx: 836},
	'➪':      {Wx: 836},
	'➫':      {Wx: 867},
	'➬':      {Wx: 867},
	'➭':      {Wx: 696},
	'➮':      {Wx: 696},
	'➯':      {Wx: 874},
	'➱':      {Wx: 874},
	'➲':      {Wx: 760},
	'➳':      {Wx: 946},
	'➴':      {Wx: 771},
	'➵':      {Wx: 865},
	'➶':      {Wx: 771},
	'➷':      {Wx: 888},
	'➸':      {Wx: 967},
	'➹':      {Wx: 888},
	'➺':      {Wx: 831},
	'➻':      {Wx: 873},
	'➼':      {Wx: 927},
	'➽':      {Wx: 970},
	'➾':      {Wx: 918},
	'\uf8d7': {Wx: 390},
	'\uf8d8': {Wx: 390},
	'\uf8d9': {Wx: 317},
	'\uf8da': {Wx: 317},
	'\uf8db': {Wx: 276},
	'\uf8dc': {Wx: 276},
	'\uf8dd': {Wx: 509},
	'\uf8de': {Wx: 509},
	'\uf8df': {Wx: 410},
	'\uf8e0': {Wx: 410},
	'\uf8e1': {Wx: 234},
	'\uf8e2': {Wx: 234},
	'\uf8e3': {Wx: 334},
	'\uf8e4': {Wx: 334},
}
