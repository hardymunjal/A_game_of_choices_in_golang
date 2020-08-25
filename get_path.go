package A_game_of_choices_in_golang

import (
	"path/filepath"
	"runtime"
)

type path struct {
	basePath          string
	pathToDataFolder  string
	PathToJsonFile    string
	PathToHtmlFile    string
	PathToStaticFiles string
}

func InitialisePath() *path {
	var paths path
	_, fileInfo, _, _ := runtime.Caller(0)
	paths.basePath = filepath.Dir(fileInfo)
	paths.pathToDataFolder = filepath.Join(paths.basePath, "data_files")
	paths.PathToJsonFile = filepath.Join(paths.pathToDataFolder, "gopher.json")
	paths.PathToHtmlFile = filepath.Join(paths.basePath, "game", "templates", "index.html")
	paths.PathToStaticFiles = filepath.Join(paths.basePath, "game")
	return &paths
}
