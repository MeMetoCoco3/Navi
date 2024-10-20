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
	favRoute = os.Getenv("FAVROUTE")
	if favRoute == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}
}

func AddFav(new_dir string) {
	index := CheckFav(new_dir)
	if index == -1 {
		favorites := LoadFavs()
		favorites.Dirs = append(favorites.Dirs, new_dir)
		SaveFavs(favorites)
	} else {
		log.Fatalln("Already in favs")
	}
}

func RemoveFav(new_dir string) {
	index := CheckFav(new_dir)

	if index >= 0 {
		favorites := LoadFavs()
		favorites.Dirs = append(favorites.Dirs[:index], favorites.Dirs[index+1:]...)
		SaveFavs(favorites)
	} else {
		log.Fatalln("Not in favs")
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
		log.Fatalln("Error loading favs: ", err)
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

/*
// Gets Paths and returns length of longest
func MaxLenPaths(paths []fs.DirEntry) int {
	if len(paths) == 0 {
		log.Fatalln("Length of paths == 0")
	}

	firstPath, err := paths[0].Info()
	if err != nil {
		log.Fatalln(err)
	}

	maxLenSoFar := firstPath.Size()
	for _, p := range paths {

		lengthP, err := p.Info()
		if err != nil {
			log.Fatalln(err)
		}

		if lengthP.Size() > maxLenSoFar {
			maxLenSoFar = lengthP.Size()
		}
	}
	return int(maxLenSoFar)
}
*/
