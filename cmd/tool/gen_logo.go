package tool

import (
	"fmt"

	"github.com/fogleman/gg"
)

func genLogo(outPath string) {
	const width = 1000
	const height = 1000

	dc := gg.NewContext(width, height)

	// define background color
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// define colors
	colors := []string{"#FF0000", "#00FF00", "#0000FF", "#FFFF00"} // Rouge, Vert, Bleu, Jaune
	texts := []string{"R&D", "Sales", "Executives", "Marketing"}
	radius := 300.0

	for i, color := range colors {
		dc.SetHexColor(color)
		dc.DrawArc(float64(width/2), float64(height/2), radius, float64(i)*gg.Radians(90), float64(i+1)*gg.Radians(90))
		dc.Fill()

		// add text
		dc.SetRGB(0, 0, 0) // Couleur du texte
		dc.DrawStringAnchored(texts[i], float64(width/2), float64(height/2), 0.5, 0.5)
	}

	// save PNG
	err := dc.SavePNG(outPath)
	if err != nil {
		fmt.Println("erreur lors de l'enregistrement de l'image :", err)
	}
}
