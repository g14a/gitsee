package color

import (
	"fmt"
	"gitsee/cache"
	"log"
	"math"
	"strconv"
)

var DarkThemedColors = []string{
	"#eb370c", "#f36721", "#f2c543", "#f24843",
	"#91dc32", "#5cb258", "#afd142", "#49707a",
	"#00ffaa", "#a8dadc", "#ffb4a2", "#00c6b4",
	"#5e60ce", "fca311", "#ef233c", "#52b788",
	"#7b2cbf", "#c44536",
}

var LanguageColors map[string]interface{}

func GetColorCodesForLanguages(user string, languages map[string]int) {

	if colorCodes, ok := cache.Get(user + "Colors"); ok {
		log.Println("Got", user, "Colors from cache")
		LanguageColors = colorCodes.(map[string]interface{})
	} else {
		colors := DarkThemedColors

		languageColors := make(map[string]interface{})
		
		i := 0
		for k, _ := range languages {
			languageColors[k] = colors[i]
			i++
		}
		
		fmt.Println(languageColors)

		LanguageColors = languageColors

		if len(LanguageColors) != 0 {
			if cache.Set(user+"Colors", LanguageColors) {
				log.Println(user + "Colors set in cache")
			}
		}
	}
}

func CompareSimilarColors(c1, c2 string) bool {

	r1, _ := strconv.ParseInt(c1[1:3], 16, 16)
	g1, _ := strconv.ParseInt(c1[3:5], 16, 16)
	b1, _ := strconv.ParseInt(c1[5:7], 16, 16)

	r2, _ := strconv.ParseInt(c2[1:3], 16, 16)
	g2, _ := strconv.ParseInt(c2[3:5], 16, 16)
	b2, _ := strconv.ParseInt(c2[5:7], 16, 16)

	r := 255 - math.Abs(float64(r1-r2))
	g := 255 - math.Abs(float64(g1-g2))
	b := 255 - math.Abs(float64(b1-b2))

	r /= 255
	g /= 255
	b /= 255

	return (r+g+b)/3 < 0.5
}
