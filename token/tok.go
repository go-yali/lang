package token

import (
	"fmt"
)

type Pos int // threw this in there so it'll work.

// item represents a token or text string returned from the scanner.
type Token struct {
	Typ ItemType // The type of this item.
	Pos Pos      // The starting position, in bytes, of this item in the input string.
	Val string   // The value of this item.
}

func (i Token) String() string {
	switch {
	case i.Typ == ItemEOF:
		return "EOF"
	case i.Typ == ItemError:
		return i.Val
	case len(i.Val) > 10:
		return fmt.Sprintf("%.10q...", i.Val)
	}
	return fmt.Sprintf("%q", i.Val)
}

// itemType identifies the type of lex items.
type ItemType int

const (
	ItemError     ItemType = iota // error occurred; value is text of error
	ItemEOF                       // end of file
	ItemBeginList                 // starts a list
	ItemEndList                   // ends a list
	beginOperation
	ItemAdd   // add
	ItemAdc   // add with carry
	ItemSub   // subtract
	ItemSbc   // subtract with carry
	ItemMul   // multiply
	ItemAnd   // bitwise and
	ItemOrr   // bitwise or
	ItemEor   // bitwise xor
	ItemBic   // bitwise bit clear
	ItemTeq   // bitwise test equality
	ItemColon // assgnment
	endOperation
	beginCompare
	ItemEq // Z set: test equality
	ItemNe // Z clear: test inequality
	ItemVs // V set: overflow
	ItemVc // V clear: no overflow
	ItemMi // N set: negative
	ItemPl // N clear: positive or zero
	ItemGt // greater than
	ItemGe // greater than or equal to
	ItemLt // less than
	ItemLe // less than or equal to
	endCompare
	beginConstant
	ItemBool      // boolean true/false
	ItemChar      // character 'c'
	ItemComplex   // complex number
	ItemNumber    // number
	ItemString    // "string"
	ItemRawString // `raw string`
	endConstant
	ItemVariable    // variable
	ItemLineComment // comment '//' style
	ItemNewline     // newline
	ItemSpace       // spaces
)

var key = map[string]ItemType{
	// Operations (instructions)
	"+":  ItemAdd,
	"+c": ItemAdc,
	"-":  ItemSub,
	"-c": ItemSbc,
	"*":  ItemMul,
	"&":  ItemAnd,
	"|":  ItemOrr,
	"^":  ItemEor,
	"&!": ItemBic,
	"=":  ItemTeq,
	// Conditionals (conditional instruction prefixes)
	cond + "=":  ItemEq,
	cond + "!=": ItemNe,
	cond + "#":  ItemVs,
	cond + "!#": ItemVc,
	cond + "-":  ItemMi,
	cond + "+":  ItemPl,
	cond + ">":  ItemGt,
	cond + ">=": ItemGe,
	cond + "<":  ItemLt,
	cond + "<=": ItemLe,
}

const cond = "?"

const Eof = -1

func IsKeyword(word string) bool {
	return key[word] > beginOperation && key[word] < endCompare
}

func Constant(typ ItemType) bool {
	return typ > beginConstant && typ < endConstant
}

func Keyword(typ ItemType) bool {
	return typ > beginOperation && typ < endCompare
}

// returns the token the word is the key to
func Lookup(word string) ItemType {
	return key[word]
}