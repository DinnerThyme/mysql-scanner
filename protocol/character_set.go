package protocol

const (
	CharacterSetUtf8 = 33
)

// CharacterSetMap maps the charset name (used in ConnParams) to the
// integer value.  Interesting ones have their own constant above.
var CharacterSetMap = map[uint8]string{
	1:                "big5",
	3:                "dec8",
	4:                "cp850",
	6:                "hp8",
	7:                "koi8r",
	8:                "latin1",
	9:                "latin2",
	10:               "swe7",
	11:               "ascii",
	12:               "ujis",
	13:               "sjis",
	16:               "hebrew",
	18:               "tis620",
	19:               "euckr",
	22:               "koi8u",
	24:               "gb2312",
	25:               "greek",
	26:               "cp1250",
	28:               "gbk",
	30:               "latin5",
	32:               "armscii8",
	CharacterSetUtf8: "utf8",
	35:               "ucs2",
	36:               "cp866",
	37:               "keybcs2",
	38:               "macce",
	39:               "macroman",
	40:               "cp852",
	41:               "latin7",
	45:               "utf8mb4",
	51:               "cp1251",
	54:               "utf16",
	56:               "utf16le",
	57:               "cp1256",
	59:               "cp1257",
	60:               "utf32",
	63:               "binary",
	92:               "geostd8",
	95:               "cp932",
	97:               "eucjpms",
	255:              "unknown",
}
