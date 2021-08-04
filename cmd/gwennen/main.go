package main

import ("context"
	"golang.design/x/clipboard"
	"github.com/gen2brain/dlgs"
	"strings"
)

func main() {

	items := getItems()

	ch := clipboard.Watch(context.Background(), clipboard.FmtText)

	for data := range ch {

		entities := strings.Split(string(data), "\n")
		
		for _, entity := range entities {
			if links, ok := items[strings.TrimSpace(entity)]; ok {
				_, err := dlgs.Info("Info", implode(links))
				if err != nil {
					panic(err)
				}
				resetClipboard()
			}
		}
	}
}

func implode(links []string) string {
	
	var sb strings.Builder

	for _, link := range links {
		println(link)
		sb.WriteString(link + "\n\n")
	}
	
	return sb.String()
}

func resetClipboard() {
	clipboard.Write(clipboard.FmtText, []byte("reset"))
}

func getItems() (map[string][]string) {
	items := map[string][]string{
		"Ezomyte staff": []string{"https://www.pathofexile.com/trade/search/Expedition/E9lGlzh5"},
		"Judgement Staff": []string{" https://www.pathofexile.com/trade/search/Expedition/8kmmb2fV"},
		"Siege Axe": []string{"https://www.pathofexile.com/trade/search/Expedition/RwL4X4i7"},
		"Assassin Bow": []string{"https://www.pathofexile.com/trade/search/Expedition/OKkvaQtE"},
		"Gavel": []string{"https://www.pathofexile.com/trade/search/Expedition/Xe7a7sP"},
		"Imperial Bow": []string{"https://www.pathofexile.com/trade/search/Expedition/dZv5yRFJ"},
		"Ezomyte Dagger": []string{"https://www.pathofexile.com/trade/search/Expedition/7nXrMPBT5"},
		"Highborn Staff": []string{"https://www.pathofexile.com/trade/search/Expedition/Q2zoMmHw"},
		"Ambush Mitts": []string{"https://pathofexile.fandom.com/wiki/Architect%27s_Hand"},
		"Imperial Skean": []string{"https://www.pathofexile.com/trade/search/Expedition/GoRrBbtb"},
		"Karui Maul": []string{"https://www.pathofexile.com/trade/search/Expedition/w0kp0Rsb"},
		"Jewelled Foil": []string{"https://www.pathofexile.com/trade/search/Expedition/BBnOXaT8"},
		"Prophecy Wand": []string{"https://www.pathofexile.com/trade/search/Expedition/954qJ8hK"},
		"Fiend Dagger": []string{"https://www.pathofexile.com/trade/search/Expedition/43d6nwu9"},
		"Demon Dagger": []string{"https://www.pathofexile.com/trade/search/Expedition/X5d864FP"},
		"Murder Mitts": []string{"https://www.pathofexile.com/trade/search/Expedition/4m04JoQt9"},
		"Zodiac Leather": []string{"https://www.pathofexile.com/trade/search/Expedition/ylLEPrFR"},
		"Champion Kite Shield": []string{"https://www.pathofexile.com/trade/search/Expedition/90JnG4HK"},
		"Full Dragonscale": []string{"https://www.pathofexile.com/trade/search/Expedition/58YawjTa"},
		"Lacquered Garb": []string{"https://www.pathofexile.com/trade/search/Expedition/58YnnnHa",
			"https://www.pathofexile.com/trade/search/Expedition/K8lP0Lc5"},
		"Painted Tower Shield": []string{"https://www.pathofexile.com/trade/search/Expedition/pw89JdI0"},
		"Two-Point Arrow Quiver": []string{"https://www.pathofexile.com/trade/search/Expedition/ZqXw88hQ"},
		"Pinnacle Tower Shield": []string{"https://www.pathofexile.com/trade/search/Expedition/ZqzJa6fQ"},
		"Glorious Plate": []string{"https://www.pathofexile.com/trade/search/Expedition/kWL7Wt5"},
		"Sinner Tricorne": []string{"https://www.pathofexile.com/trade/search/Expedition/8kmVYntV"},
		"Deerskin Gloves": []string{"https://www.pathofexile.com/trade/search/Expedition/r2n5dZSQ"},
		"Nightmare Bascinet": []string{"https://www.pathofexile.com/trade/search/Expedition/ylzz8eiR"},
		"Crusader Plate": []string{"https://www.pathofexile.com/trade/search/Expedition/n7WKWvu0"},
		"Sadist Garb": []string{"https://www.pathofexile.com/trade/search/Expedition/k78yPrt5",
			"https://www.pathofexile.com/trade/search/Expedition/RwzEBqU7"},
		"Tricorne": []string{"https://www.pathofexile.com/trade/search/Expedition/ZWWn5L9iQ"},
		"Gladiator Plate": []string{"https://www.pathofexile.com/trade/search/Expedition/XRnwjdcP"},
		"Occultist's Vestment": []string{"https://www.pathofexile.com/trade/search/Expedition/aVq58He"},
		"Ebony Tower Shield": []string{"https://www.pathofexile.com/trade/search/Expedition/elj0BoFL"},
		"Archon Kite Shield": []string{"https://www.pathofexile.com/trade/search/Expedition/lynPDrUV"},
		"Praetor Crown": []string{"https://www.pathofexile.com/trade/search/Expedition/Q2yE5MCw"},
		"Silk Gloves": []string{"https://www.pathofexile.com/trade/search/Expedition/d0ppoqUJ"},
		"Simple Robe": []string{"https://www.pathofexile.com/trade/search/Expedition/NK6Ec5"},
		"Amethyst Ring": []string{"https://www.pathofexile.com/trade/search/Expedition/ar3Y7R7fe"},
		"Studded Belt": []string{"https://www.pathofexile.com/trade/search/Expedition/pw04bvH0"},
		"Coral Amulet": []string{"https://www.pathofexile.com/trade/search/Expedition/20mMRVck"},
		"Gold Amulet": []string{"https://www.pathofexile.com/trade/search/Expedition/YblaLbhY"},
		"Sapphire Ring": []string{"https://www.pathofexile.com/trade/search/Expedition/32vaXkh5"},
		"Two-Stone Ring": []string{"https://www.pathofexile.com/trade/search/Expedition/RwzwmBs7"},
		"Agate Amulet": []string{"https://www.pathofexile.com/trade/search/Expedition/zaznDKH4"},
		"Paua Amulet": []string{"https://www.pathofexile.com/trade/search/Expedition/8kmR6nuV"},
		"Chain Belt": []string{"https://www.pathofexile.com/trade/search/Expedition/a9PD47He"},
		"Unset Ring": []string{"https://www.pathofexile.com/trade/search/Expedition/7887k03C5"},
		"Crystal Belt": []string{"https://www.pathofexile.com/trade/search/Expedition/b8BzZKtL"},
		"Moonstone Ring": []string{"https://www.pathofexile.com/trade/search/Expedition/EnggBLf5"},
		"Leather Belt": []string{"https://www.pathofexile.com/trade/search/Expedition/12bdck"},
		"Prismatic Jewel": []string{"https://www.pathofexile.com/trade/search/Expedition/vKyzHE"},
		"Crimson Jewel": []string{"https://www.pathofexile.com/trade/search/Expedition/9aD3QFK",
			"https://www.pathofexile.com/trade/search/Expedition/bkwGUL",
			"https://www.pathofexile.com/trade/search/Expedition/VZqootp",
			"https://www.pathofexile.com/trade/search/Expedition/wvRvMBsb"},
		"Viridian Jewel": []string{"https://www.pathofexile.com/trade/search/Expedition/O62GPlHE",
			"https://www.pathofexile.com/trade/search/Expedition/bGP2KHL",
			"https://www.pathofexile.com/trade/search/Expedition/Xmv3WtP",
			"https://www.pathofexile.com/trade/search/Expedition/yXLb8hR"},
		"Onyx Amulet":  []string{"https://www.pathofexile.com/trade/search/Expedition/EnX8EzC5",
			"https://www.pathofexile.com/trade/search/Expedition/jgErVZuX"},
		"Citrine Amulet": []string{"https://www.pathofexile.com/trade/search/Expedition/pJoKG3XH0",
			"https://www.pathofexile.com/trade/search/Expedition/el0PPbiL"},
		"Cobalt Jewel": []string{"https://www.pathofexile.com/trade/search/Expedition/Dl4vC5",
			"https://www.pathofexile.com/trade/search/Expedition/jv35LCX",
			"https://www.pathofexile.com/trade/search/Expedition/zbM8pYqc4",
			"https://www.pathofexile.com/trade/search/Expedition/rY2mhQ",
			"https://www.pathofexile.com/trade/search/Expedition/0YjD8ug",
			"https://www.pathofexile.com/trade/search/Expedition/DVm85f5",
			"https://www.pathofexile.com/trade/search/Expedition/lowjTV"},
		"Turquoise Amulet": []string{"https://www.pathofexile.com/trade/search/Expedition/G6nvgBWcb",
			"https://www.pathofexile.com/trade/search/Expedition/Mdvm8vyFJ"}}
	return items
}
