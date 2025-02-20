package main

import (
	"fmt"
	"os"
)

func schrijfGedicht(kleur string) {
	gedichten := map[string]string{
		"blauw":  "Blauw zoals de lucht.",
		"rood":   "Rood met passie.",
		"geel":   "Geel als de stralen van de zon.",
		"groen":  "Groen van de natuur.",
	}

	gedicht, gevonden := gedichten[kleur]
	if !gevonden {
		fmt.Printf("Fout: Kleur '%s' niet herkend.\n", kleur)
		os.Exit(1)
	}

    fmt.Println(gedicht)
}

func main() {
	var kleur string
	fmt.Print("Voer een kleur in: ")
	fmt.Scanln(&kleur)

	schrijfGedicht(kleur)
}