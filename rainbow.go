package rainbow

import (
	"fmt"
)

var (
	reset = "\u001b[0m"

	colorBlack = "\x1b[30m"
	colorBlue = "\x1b[34m"
	colorRed = "\x1b[31m"
	colorGreen = "\x1b[32m"
	colorYellow = "\x1b[33m"
	colorMagenta = "\x1b[35m"
	colorCyan = "\x1b[36m"
	colorWhite ="\x1b[37m"

	bgBlack = "\x1b[40m"
	bgBlue = "\x1b[44m"
	bgRed = "\x1b[41m"
	bgGreen = "\x1b[42m"
	bgYellow = "\x1b[43m"
	bgMagenta = "\x1b[45m"
	bgCyan = "\x1b[46m"
	bgWhite = "\x1b[47m"

	modifierBold = "\x1b[1m"
	modifierDim = "\x1b[2m"
	modifierUnderscore = "\x1b[4m"
	modifierBlink = "\x1b[5m"
	modifierReverse = "\x1b[7m"
	modifierHidden = "\x1b[8m"
)

func Black(String string) string {
	return colorBlack + String + reset
}

func Blue(String string) string {
	return colorBlue + String + reset
}

func Red(String string) string {
	return colorRed + String + reset
}

func Green(String string) string {
	return colorGreen + String + reset
}

func Yellow(String string) string {
	return colorYellow + String + reset
}

func Magenta(String string) string {
	return colorMagenta + String + reset
}

func Cyan(String string) string {
	return colorCyan + String + reset
}

func White(String string) string {
	return colorWhite + String + reset
}

func BgBlack(String string) string {
	return bgBlack + String + reset
}

func BgBlue(String string) string {
	return bgBlue + String + reset
}

func BgRed(String string) string {
	return bgRed + String + reset
}

func BgGreen(String string) string {
	return bgGreen + String + reset
}

func BgYellow(String string) string {
	return bgYellow + String + reset
}

func BgMagenta(String string) string {
	return bgMagenta + String + reset
}

func BgCyan(String string) string {
	return bgCyan + String + reset
}

func BgWhite(String string) string {
	return bgWhite + String + reset
}

func Bold(String string) string {
	return modifierBold + String + reset
}

func Dim(String string) string {
	return modifierDim + String + reset
}

func Underscore(String string) string {
	return modifierUnderscore + String + reset
}

func Blink(String string) string {
	return modifierBlink + String + reset
}

func Reverse(String string) string {
	return modifierReverse + String + reset
}

func Hidden(String string) string {
	return modifierHidden + String + reset
}

// func main() {
// 	fmt.Printf("%s %s %s %s %s %s %s %s", Blue("blue"), Black("black"), Red("red"), Green("green"), Yellow("yellow"), Magenta("magenta"), Cyan("cyan"), White("white"))
// 	fmt.Printf("\n%s %s %s %s %s %s %s %s", BgBlue("bgBlue"), BgBlack("BgBlack"), BgRed("bgRed"), BgGreen("bgGreen"), BgYellow("bgYellow"), BgMagenta("bgMagenta"), BgCyan("bgCyan"), BgWhite("bgWhite"))
// 	fmt.Printf("\n%s %s %s %s %s %s", Bold("Bold"), Dim("Dim"), Underscore("Underscore"), Blink("Blink"), Reverse("Reverse"), Hidden("Hidden"))
// 	fmt.Printf("\n%s %s %s", Bold(BgYellow(Blue("Bold+BgYellow+Blue"))), Dim(Underscore("Dim+Underscore")), BgMagenta(Red("BgMagenta+Red")))
// }

