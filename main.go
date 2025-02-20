package main

import (
	"bufio"
	"fmt"
	"os"
)

func schrijfGedicht(kleur string) {
	gedichten := map[string]string{
		"aliceblauw":       "Aliceblauw, als de lucht boven een bevroren wonderland.",
		"antiekwit":        "Antiekwit, de kleur van oude kant, verhalen fluisterend.",
		"aqua":             "Aqua, als het diepe water van een tropisch rif.",
		"aquamarijn":       "Aquamarijn, een zeemeermin haar verloren juweel.",
		"azuur":            "Azuur, de kleur van de hemel op een heldere dag.",
		"beige":            "Beige, als het zand op een verlaten strand.",
		"bisque":           "Bisque, een romige soep, zacht en warm.",
		"zwart":            "Zwart, de afwezigheid van licht, een leeg canvas.",
		"amandelwit":       "Amandelwit, de kleur van een pasgeboren lammetje.",
		"blauw":            "Blauw, de kleur van de nacht, diep en mysterieus.",
		"blauwviolet":      "Blauwviolet, een bloem in de schemering, geheimzinnig en mooi.",
		"bruin":            "Bruin, de kleur van de aarde, stevig en betrouwbaar.",
		"houtkleur":        "Houtkleur, als het hout van een oude schuur, verweerd door de tijd.",
		"cadetblauw":       "Cadetblauw, een militaire kleur, stoer en gedisciplineerd.",
		"chartreuse":       "Chartreuse, een explosie van lente, fris en energiek.",
		"chocoladebruin":   "Chocoladebruin, de kleur van verwennerij, rijk en decadent.",
		"koraal":           "Koraal, als de kleur van de zee, warm en levendig.",
		"korenbloemblauw":  "Korenbloemblauw, een zomerse kleur, licht en vrolijk.",
		"maïszijde":        "Maïszijde, als de kleur van maïsvelden in de zon.",
		"crimson":          "Crimson, een dieprode kleur, intens en gepassioneerd.",
		"cyaan":            "Cyaan, een kleur tussen blauw en groen, fris en helder.",
		"donkerblauw":      "Donkerblauw, als de diepste oceanen, vol geheimen.",
		"donkercyaan":      "Donkercyaan, een mysterieuze kleur, als de schaduw van de nacht.",
		"donkergoudgeel":   "Donkergoudgeel, als de kleur van oude schatten, kostbaar en glanzend.",
		"donkergrijs":      "Donkergrijs, een sombere kleur, als een bewolkte dag.",
		"donkergroen":      "Donkergroen, als de kleur van een diep bos, mysterieus en ondoordringbaar.",
		"donkerkaki":       "Donkerkaki, een aardse kleur, als de woestijn.",
		"donkermagenta":    "Donkermagenta, een dieppaarse kleur, mystiek en elegant.",
		"donkerolijfgroen": "Donkerolijfgroen, een natuurlijke kleur, als de bladeren van een olijfboom.",
		"donkeroranje":     "Donkeroranje, als de kleur van een herfstdag, warm en gloeiend.",
		"donkerorchidee":   "Donkerorchidee, een exotische kleur, als een bloem in de nacht.",
		"donkerrood":       "Donkerrood, als de kleur van bloed, krachtig en intens.",
		"donkerzalm":       "Donkerzalm, een delicate kleur, als de huid van een vis.",
		"donkerzeegroen":   "Donkerzeegroen, een kalme kleur, als de diepe zee.",
		"dieproze":         "Dieproze, een intense kleur, als de liefde zelf.",
		"diepluchtblauw":   "Diepluchtblauw, een heldere kleur, als de zomerse lucht.",
		"dimgrijs":         "Dimgrijs, een zachte kleur, als de schaduw van de wolken.",
		"dodgerblauw":      "Dodgerblauw, een sportieve kleur, als de snelheid van de wind.",
		"baksteenrood":     "Baksteenrood, een warme kleur, als een knappend haardvuur.",
		"bloemwit":         "Bloemwit, een delicate kleur, als de bloemblaadjes van een lelie.",
		"bosgroen":         "Bosgroen, een natuurlijke kleur, als het hart van het bos.",
		"fuchsia":          "Fuchsia, een levendige kleur, als een explosie van vreugde.",
		"goud":             "Goud, een kostbare kleur, als de zon zelf.",
		"goudstaaf":        "Goudstaaf, een warme kleur, als de stralen van de zon.",
		"grijs":            "Grijs, een neutrale kleur, als de kleur van de wolken.",
		"groen":            "Groen, de kleur van de natuur, leven en groei symboliserend.",
		"groengeel":        "Groengeel, een vrolijke kleur, als de eerste bloemen in de lente.",
		"honingdauw":       "Honingdauw, een zoete kleur, als de nectar van bloemen.",
		"helderroze":       "Helderroze, een opvallende kleur, als een explosie van energie.",
		"indischrood":      "Indischrood, een aardse kleur, als de kleur van de woestijn.",
		"indigo":           "Indigo, een diepblauwe kleur, als de nachtelijke hemel.",
		"ivoor":            "Ivoor, een elegante kleur, als de toetsen van een piano.",
		"kaki":             "Kaki, een praktische kleur, als de kleding van een avonturier.",
		"lavendel":         "Lavendel, een rustgevende kleur, als de geur van lavendelvelden.",
		"lavendelblos":     "Lavendelblos, een delicate kleur, als de bloos op de wangen.",
		"grasgroen":        "Grasgroen, een heldere kleur, als het gras in de lente.",
		"citroengeel":      "Citroengeel, een zachte kleur, als een zomerse taart.",
		"lichtblauw":       "Lichtblauw, een luchtige kleur, als de wolken aan de hemel.",
		"lichtgroen":       "Lichtgroen, een vrolijke kleur, als het ontluiken van de natuur.",
		"lichtgeel":        "Lichtgeel, een vrolijke kleur, als de eerste bloemen in de lente.",
		"limoen":           "Limoen, een frisse kleur, als de smaak van limoen.",
		"limoengroen":      "Limoengroen, een heldere kleur, als de bladeren in de lente.",
		"linnen":           "Linnen, een natuurlijke kleur, als het weefsel van linnen.",
		"magenta":          "Magenta, een krachtige kleur, als de kleur van passie.",
		"kastanjebruin":    "Kastanjebruin, een dieprode kleur, als de kleur van wijn.",
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
	fmt.Print("Voer je lievelingskleur in: ")
	fmt.Scanln(&kleur)

	schrijfGedicht(kleur)
}
