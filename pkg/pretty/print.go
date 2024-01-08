package pretty

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/zyedidia/highlight"
)

func print(input string, def *highlight.Def) {
	h := highlight.NewHighlighter(def)
	matches := h.HighlightString(input)

	lines := strings.Split(input, "\n")
	for lineN, l := range lines {
		for colN, c := range l {
			// Check if the group changed at the current position
			if group, ok := matches[lineN][colN]; ok {
				// Check the group name and set the color accordingly (the colors chosen are arbitrary)
				if group == highlight.Groups["statement"] {
					color.Set(color.FgBlue)
				} else if group == highlight.Groups["preproc"] {
					color.Set(color.FgHiGreen)
				} else if group == highlight.Groups["special"] {
					color.Set(color.FgBlue)
				} else if group == highlight.Groups["constant.string"] {
					color.Set(color.FgHiMagenta)
				} else if group == highlight.Groups["constant.specialChar"] {
					color.Set(color.FgHiMagenta)
				} else if group == highlight.Groups["type"] {
					color.Set(color.FgCyan)
				} else if group == highlight.Groups["constant.number"] {
					color.Set(color.FgCyan)
				} else if group == highlight.Groups["comment"] {
					color.Set(color.FgHiBlack)
				} else {
					color.Unset()
				}
			}
			// Print the character
			fmt.Print(string(c))
		}
		// This is at a newline, but highlighting might have been turned off at the very end of the line so we should check that.
		if group, ok := matches[lineN][len(l)]; ok {
			if group == highlight.Groups["default"] || group == highlight.Groups[""] {
				color.Unset()
			}
		}

		color.Unset()
		fmt.Print("\n")
	}
}
