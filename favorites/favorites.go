package favorites

import (
	"encoding/json"
	"log"
	"os"
)

// Type definition with json tag, helps with encoding.
// Says that Dirs should be kept as a JSON key called dirs.
type Favorites struct {
	Dirs []string `json:"dirs"`
}

var favRoute string

func init() {
	// Set on Bashscript
	favRoute = os.Getenv("FAVROUTE")
	if favRoute == "" {
		log.Fatal("(-) CONFIG_PATH environment variable not set")
	}
}

func AddFav(new_dir string) {
	index := CheckFav(new_dir)
	if index == -1 {
		favorites := LoadFavs()
		favorites.Dirs = append(favorites.Dirs, new_dir)
		SaveFavs(favorites)
	} else {
		log.Fatalln("(-) Already in favs")
	}
}

func RemoveFav(new_dir string) {
	index := CheckFav(new_dir)

	if index >= 0 {
		favorites := LoadFavs()
		favorites.Dirs = append(favorites.Dirs[:index], favorites.Dirs[index+1:]...)
		SaveFavs(favorites)
	} else {
		log.Fatalf("(-) Directory '%s' was not in favs", new_dir)
	}
}

func CheckFav(new_dir string) int {
	favorites := LoadFavs()
	for index, dir := range favorites.Dirs {
		if dir == new_dir {
			return index
		}
	}
	return -1
}

func SaveFavs(favorites Favorites) {
	data, _ := json.MarshalIndent(favorites, "", "  ")
	os.WriteFile(favRoute, data, 0644)
}

func ListFavs() []string {
	list := LoadFavs()
	return list.Dirs
}

func LoadFavs() Favorites {
	var favorites Favorites
	data, err := os.ReadFile(favRoute)
	if err != nil {
		log.Fatalln("(-) Error loading favs: ", err)
	}
	json.Unmarshal(data, &favorites)
	return favorites
}

func WriteOnTmp(content string) {
	f, err := os.Create("/tmp/GatorPath")
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}
