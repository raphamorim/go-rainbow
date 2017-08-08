package rainbow

import (
	"strconv"

	"github.com/tomnomnom/xtermcolor"
)

var (
	reset = "\u001b[0m"

	colorBlack   = "\x1b[30m"
	colorBlue    = "\x1b[34m"
	colorRed     = "\x1b[31m"
	colorGreen   = "\x1b[32m"
	colorYellow  = "\x1b[33m"
	colorMagenta = "\x1b[35m"
	colorCyan    = "\x1b[36m"
	colorWhite   = "\x1b[37m"

	bgBlack   = "\x1b[40m"
	bgBlue    = "\x1b[44m"
	bgRed     = "\x1b[41m"
	bgGreen   = "\x1b[42m"
	bgYellow  = "\x1b[43m"
	bgMagenta = "\x1b[45m"
	bgCyan    = "\x1b[46m"
	bgWhite   = "\x1b[47m"

	modifierBold       = "\x1b[1m"
	modifierDim        = "\x1b[2m"
	modifierUnderscore = "\x1b[4m"
	modifierBlink      = "\x1b[5m"
	modifierReverse    = "\x1b[7m"
	modifierHidden     = "\x1b[8m"
)

func Black(text string) string {
	return colorBlack + text + reset
}

func Blue(text string) string {
	return colorBlue + text + reset
}

func Red(text string) string {
	return colorRed + text + reset
}

func Green(text string) string {
	return colorGreen + text + reset
}

func Yellow(text string) string {
	return colorYellow + text + reset
}

func Magenta(text string) string {
	return colorMagenta + text + reset
}

func Cyan(text string) string {
	return colorCyan + text + reset
}

func White(text string) string {
	return colorWhite + text + reset
}

func BgBlack(text string) string {
	return bgBlack + text + reset
}

func BgBlue(text string) string {
	return bgBlue + text + reset
}

func BgRed(text string) string {
	return bgRed + text + reset
}

func BgGreen(text string) string {
	return bgGreen + text + reset
}

func BgYellow(text string) string {
	return bgYellow + text + reset
}

func BgMagenta(text string) string {
	return bgMagenta + text + reset
}

func BgCyan(text string) string {
	return bgCyan + text + reset
}

func BgWhite(text string) string {
	return bgWhite + text + reset
}

func Bold(text string) string {
	return modifierBold + text + reset
}

func Dim(text string) string {
	return modifierDim + text + reset
}

func Underscore(text string) string {
	return modifierUnderscore + text + reset
}

func Blink(text string) string {
	return modifierBlink + text + reset
}

func Reverse(text string) string {
	return modifierReverse + text + reset
}

func Hidden(text string) string {
	return modifierHidden + text + reset
}

func Hex(hexColor, text string) string {
	code, err := xtermcolor.FromHexStr(hexColor)

	if err != nil {
		return White(text)
	}

	return "\033[1;38;5;" + strconv.Itoa(int(code)) + "m" + text
}

func FromInt32(Int uint32, text string) string {
	code := xtermcolor.FromInt(Int)
	return "\033[1;38;5;" + strconv.Itoa(int(code)) + "m" + text
}
