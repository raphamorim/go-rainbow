package main

import (
	"fmt"

	"github.com/raphamorim/go-rainbow"
)

func main() {
	fmt.Println(rainbow.Bold(rainbow.Hex("#8E44AD", "raphael")))
	fmt.Printf("%s %s %s %s %s %s %s %s", rainbow.Blue("blue"), rainbow.Black("black"), rainbow.Red("red"), rainbow.Green("green"), rainbow.Yellow("yellow"), rainbow.Magenta("magenta"), rainbow.Cyan("cyan"), rainbow.White("white"))

	fmt.Printf("\n%s", rainbow.Hex("#000080", "String from HEX"))
	fmt.Printf("\n%s", rainbow.FromInt32(0xCC66FFFF, "String from Int32"))
	fmt.Printf("\n%s %s %s %s %s %s %s %s", rainbow.BgBlue("bgBlue"), rainbow.BgBlack("BgBlack"), rainbow.BgRed("bgRed"), rainbow.BgGreen("bgGreen"), rainbow.BgYellow("bgYellow"), rainbow.BgMagenta("bgMagenta"), rainbow.BgCyan("bgCyan"), rainbow.BgWhite("bgWhite"))
	fmt.Printf("\n%s %s %s %s %s %s", rainbow.Bold("Bold"), rainbow.Dim("Dim"), rainbow.Underscore("Underscore"), rainbow.Blink("Blink"), rainbow.Reverse("Reverse"), rainbow.Hidden("Hidden"))
	fmt.Printf("\n%s %s %s", rainbow.Bold(rainbow.BgYellow(rainbow.Blue("Bold+BgYellow+Blue"))), rainbow.Dim(rainbow.Underscore("Dim+Underscore")), rainbow.BgMagenta(rainbow.Red("BgMagenta+Red")))

	fmt.Printf("\n%s %s %s %s %s %s %s %s", rainbow.Blue("blue"), rainbow.Black("black"), rainbow.Red("red"), rainbow.Green("green"), rainbow.Yellow("yellow"), rainbow.Magenta("magenta"), rainbow.Cyan("cyan"), rainbow.White("white"))
	fmt.Printf("\n%s %s %s %s %s %s %s %s", rainbow.BlueBright("blue"), rainbow.Gray("Gray"), rainbow.RedBright("red"), rainbow.GreenBright("green"), rainbow.YellowBright("yellow"), rainbow.MagentaBright("magenta"), rainbow.CyanBright("cyan"), rainbow.WhiteBright("white"))

	fmt.Printf("\n%s %s %s %s %s %s %s %s", rainbow.BgBlueBright("bgBlue"), rainbow.BgBlackBright("BgBlack"), rainbow.BgRedBright("bgRed"), rainbow.BgGreenBright("bgGreen"), rainbow.BgYellowBright("bgYellow"), rainbow.BgMagentaBright("bgMagenta"), rainbow.BgCyanBright("bgCyan"), rainbow.BgWhiteBright("bgWhite"))

	fmt.Printf("\n%s %s", rainbow.BgCyan(rainbow.Yellow("Green")), rainbow.BgCyan(rainbow.YellowBright("Bright!")))
	fmt.Printf("\n%s %s", rainbow.Green(rainbow.Underscore("UnderscoreBright")), rainbow.Underscore(rainbow.GreenBright("UnderscoreBright")))
	fmt.Printf("\n%s %s", rainbow.BgWhite(rainbow.Red("Red")), rainbow.BgWhite(rainbow.RedBright("Bright!")))
	fmt.Printf("\n%s %s", rainbow.BgWhiteBright(rainbow.Red("Red")), rainbow.BgWhiteBright(rainbow.RedBright("Bright!")))
	fmt.Printf("\n%s %s\n", rainbow.BgBlack(rainbow.Green("Green")), rainbow.BgBlack(rainbow.GreenBright("Bright!")))
}
