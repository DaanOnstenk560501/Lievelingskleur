package main

import (
	"bufio"
	"fmt"
	"os"
)

func schrijfGedicht(kleur string) {
	gedichten := map[string]string{
		"blauw":  "Blauw zoals de lucht.",
		"rood":   "Rood met passie.",
		"geel":   "Geel als de stralen van de zon.",
		"groen":  "Groen van de natuur.",
		"oranje": "Oranje, vrolijk en warm.",
		"paars":  "Paars, mysterieus en diep.",
	}

	gedicht, gevonden := gedichten[kleur]
	if !gevonden {
		fmt.Printf("Fout: Kleur '%s' niet herkend.\n", kleur)
		os.Exit(1)
	}

	fmt.Println(gedicht)

	fmt.Print("Druk op Enter om af te sluiten...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func main() {
	var kleur string
	fmt.Print("Voer een kleur in: ")
	fmt.Scanln(&kleur)

	schrijfGedicht(kleur)
}
