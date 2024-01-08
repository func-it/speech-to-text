package pretty

import "github.com/pterm/pterm"

func Banner(header string) {
	pterm.DefaultHeader. // Use DefaultHeader as base
				WithMargin(30).
				WithBackgroundStyle(pterm.NewStyle(pterm.BgDarkGray)).
				WithTextStyle(pterm.NewStyle(pterm.FgYellow)).
				Println(header)
	//newHeader := pterm.HeaderPrinter{
	//	TextStyle:       pterm.NewStyle(pterm.FgBlack),
	//	BackgroundStyle: pterm.NewStyle(pterm.BgRed),
	//	Margin:          20,
	//}
	//newHeader.Println("This is a custom header!")
}
