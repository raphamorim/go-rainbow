package rainbow

import (
	"fmt"
	"strconv"
	"time"

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

	colorGray          = "\x1b[90m"
	colorBlueBright    = "\x1b[94m"
	colorRedBright     = "\x1b[91m"
	colorGreenBright   = "\x1b[92m"
	colorYellowBright  = "\x1b[93m"
	colorMagentaBright = "\x1b[95m"
	colorCyanBright    = "\x1b[96m"
	colorWhiteBright   = "\x1b[97m"

	bgBlack   = "\x1b[40m"
	bgBlue    = "\x1b[44m"
	bgRed     = "\x1b[41m"
	bgGreen   = "\x1b[42m"
	bgYellow  = "\x1b[43m"
	bgMagenta = "\x1b[45m"
	bgCyan    = "\x1b[46m"
	bgWhite   = "\x1b[47m"

	bgBlackBright   = "\x1b[100m"
	bgBlueBright    = "\x1b[104m"
	bgRedBright     = "\x1b[101m"
	bgGreenBright   = "\x1b[102m"
	bgYellowBright  = "\x1b[103m"
	bgMagentaBright = "\x1b[105m"
	bgCyanBright    = "\x1b[106m"
	bgWhiteBright   = "\x1b[107m"

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

func Gray(text string) string {
	return colorGray + text + reset
}

func BlueBright(text string) string {
	return colorBlueBright + text + reset
}

func RedBright(text string) string {
	return colorRedBright + text + reset
}

func GreenBright(text string) string {
	return colorGreenBright + text + reset
}

func YellowBright(text string) string {
	return colorYellowBright + text + reset
}

func MagentaBright(text string) string {
	return colorMagentaBright + text + reset
}

func CyanBright(text string) string {
	return colorCyanBright + text + reset
}

func WhiteBright(text string) string {
	return colorWhiteBright + text + reset
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

func BgBlackBright(text string) string {
	return bgBlackBright + text + reset
}

func BgBlueBright(text string) string {
	return bgBlueBright + text + reset
}

func BgRedBright(text string) string {
	return bgRedBright + text + reset
}

func BgGreenBright(text string) string {
	return bgGreenBright + text + reset
}

func BgYellowBright(text string) string {
	return bgYellowBright + text + reset
}

func BgMagentaBright(text string) string {
	return bgMagentaBright + text + reset
}

func BgCyanBright(text string) string {
	return bgCyanBright + text + reset
}

func BgWhiteBright(text string) string {
	return bgWhiteBright + text + reset
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

func Animation(text string, animation string) {
	if animation == "neon" {
		go neonAnimation(text)
	}

	select {}
}

func neonAnimation(text string) {
	fmt.Println("\033[E")
	light := false
	for range time.Tick(time.Second) {
		if light {
			fmt.Println("\033[A\033[K" + Bold(Magenta(text)))
		} else {
			fmt.Println("\033[A\033[K" + Bold(Dim(text)))
		}
		light = !light
	}
}
