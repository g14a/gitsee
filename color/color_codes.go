package color

import (
	"fmt"
	"gitsee/cache"
	"log"
	"math/rand"
	"time"
)

var DarkThemedColors = []string{
	"#eeeeee", "#00adb5", "#3c4bbf0", "#927fbf",
	"#0a91ab", "#ffc045", "#f30a49", "#04879c",
	"#e43f5a", "#f2a365", "#ffbd69", "#ff6363",
	"#29c7ac", "#bbe1fa", "#3282b8", "#515585",
	"#ee4540", "#c72c41", "#b030b0", "#a72693",
	"#ff6768", "#3ca3e47", "#d65a31", "#69779b",
	"#e14594", "#278ea5", "#616f39", "#85cfcb",
	"#3c6562", "#4ecca3", "#00818a", "#8d6262",
	"#a0204c", "#ff8ba0", "#ff5733", "#9a0f98",
	"#77abb7", "#b0a565", "#e47676", "#90b8f8",
}

var LanguageColors map[string]interface{}

func GetColorCodesForLanguages(user string, languages map[string]int) {

	if colorCodes, ok := cache.Get(user + "Colors"); ok {
		log.Println("Got", user, "Colors from cache")
		LanguageColors = colorCodes.(map[string]interface{})
		fmt.Println(LanguageColors)
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

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
