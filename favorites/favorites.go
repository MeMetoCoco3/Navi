package favorites

import (
	"encoding/json"
	"errors"
	"os"
)

// Type definition with json tag, helps with encoding.
// Says that Dirs should be kept as a JSON key called dirs.
type Favorites struct {
	Dirs []string `json:"dirs"`
}

const favRoute = "/home/evildead20/Documents/Projects/GO/navi/favorites/favs.json"

func AddFav(new_dir string) error {
	favorites, err := LoadFavs()
	if err != nil {
		return err
	}

	// Check if it exists.
	for _, dir := range favorites.Dirs {
		if dir == new_dir {
			return errors.New("(-) Directory already in favorites")
		}
	}

	favorites.Dirs = append(favorites.Dirs, new_dir)
	SaveFavs(favorites)
	return nil
}

func ChangeDir(index int) (string, error) {
	dirs, err := ListFavs()
	if err != nil {
		return "", err
	}

	if index >= len(dirs) {
		return "", errors.New("(-) Error: Index out of bounds.")
	}
	newDir := dirs[index]
	return newDir, nil
}

func ListFavs() ([]string, error) {
	list, err := LoadFavs()
	if err != nil {
		return list.Dirs, err
	}

	return list.Dirs, nil
}

func LoadFavs() (Favorites, error) {
	var favorites Favorites
	data, err := os.ReadFile(favRoute)
	if err != nil {
		return favorites, err
	}
	json.Unmarshal(data, &favorites)
	return favorites, nil
}

func SaveFavs(favorites Favorites) {
	data, _ := json.MarshalIndent(favorites, "", "  ")
	os.WriteFile(favRoute, data, 0644)

}
