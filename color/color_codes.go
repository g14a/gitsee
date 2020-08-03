package color

import (
	"gitsee/cache"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

var DarkThemedColors = []string{
	"#f4b89a", "#f36721", "#f2c543", "#f24843",
	"#afd142", "#ece74c", "#91dc32", "#5cb258",
	"#08f14b", "#00ffaa", "#00c6b4", "#0096c6",
	"#3f97fb", "#6b79fe", "#7737ff", "#f000ff",
	"#ff006c", "#b1a2a5", "#ac727d",
}

var LanguageColors map[string]interface{}

func GetColorCodesForLanguages(user string, languages map[string]int) {

	if colorCodes, ok := cache.Get(user + "Colors"); ok {
		log.Println("Got", user, "Colors from cache")
		LanguageColors = colorCodes.(map[string]interface{})
	} else {
		rand.Seed(time.Now().UnixNano())

		colors := DarkThemedColors

		languageColors := make(map[string]interface{})
		for k, _ := range languages {
			index := rand.Intn(len(colors))
			languageColors[k] = colors[index]
			colors = remove(colors, index)
		}

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

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
