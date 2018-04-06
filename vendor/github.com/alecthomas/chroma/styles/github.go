package styles

import (
	"github.com/alecthomas/chroma"
)

// GitHub style.
var GitHub = Register(chroma.MustNewStyle("github", chroma.StyleEntries{
	chroma.CommentMultiline:     "italic #999988",
	chroma.CommentPreproc:       "bold #999999",
	chroma.CommentSingle:        "italic #999988",
	chroma.CommentSpecial:       "bold italic #999999",
	chroma.Comment:              "italic #999988",
	chroma.Error:                "bg:#e3d2d2 #a61717",
	chroma.GenericDeleted:       "bg:#ffdddd #000000",
	chroma.GenericEmph:          "italic #000000",
	chroma.GenericError:         "#aa0000",
	chroma.GenericHeading:       "#999999",
	chroma.GenericInserted:      "bg:#ddffdd #000000",
	chroma.GenericOutput:        "#888888",
	chroma.GenericPrompt:        "#555555",
	chroma.GenericStrong:        "bold",
	chroma.GenericSubheading:    "#aaaaaa",
	chroma.GenericTraceback:     "#aa0000",
	chroma.KeywordType:          "bold #445588",
	chroma.Keyword:              "bold #000000",
	chroma.LiteralNumber:        "#009999",
	chroma.LiteralStringRegex:   "#009926",
	chroma.LiteralStringSymbol:  "#990073",
	chroma.LiteralString:        "#d14",
	chroma.NameAttribute:        "#008080",
	chroma.NameBuiltinPseudo:    "#999999",
	chroma.NameBuiltin:          "#0086B3",
	chroma.NameClass:            "bold #445588",
	chroma.NameConstant:         "#008080",
	chroma.NameDecorator:        "bold #3c5d5d",
	chroma.NameEntity:           "#800080",
	chroma.NameException:        "bold #990000",
	chroma.NameFunction:         "bold #990000",
	chroma.NameLabel:            "bold #990000",
	chroma.NameNamespace:        "#555555",
	chroma.NameTag:              "#000080",
	chroma.NameVariableClass:    "#008080",
	chroma.NameVariableGlobal:   "#008080",
	chroma.NameVariableInstance: "#008080",
	chroma.NameVariable:         "#008080",
	chroma.Operator:             "bold #000000",
	chroma.TextWhitespace:       "#bbbbbb",
	chroma.Background:           " bg:#ffffff",
}))
