package internal

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const generatorHomeDirName = ".solution-generator"

var skipDirs = []string{
	".idea",
}

func SaveSolution(solutionPath, solutionName string) {
	savePath := getSolutionSavePath(solutionName)

	createDir(savePath)

	saveRecursive(savePath, solutionPath, readDirectory(solutionPath))
}

func GetCurrentSolutionPath() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("fialed to get current solution dir: %e", err)
	}

	return wd
}

func IsGeneratorDirExist() bool {
	if _, err := os.Stat(getGeneratorHomePath()); os.IsNotExist(err) {
		return false
	}

	return true
}

func MakeInstall() {
	createDir(getGeneratorHomePath())
}

func saveRecursive(savePath, sourcePath string, entries []os.DirEntry) {
	for _, v := range entries {
		if skip(v.Name()) {
			continue
		}

		fullSavePath := filepath.Join(savePath, v.Name())
		fullSourcePath := filepath.Join(sourcePath, v.Name())

		if v.IsDir() {
			createDir(fullSavePath)
			saveRecursive(fullSavePath, fullSourcePath, readDirectory(fullSourcePath))
			continue
		}

		err := copyFile(fullSourcePath, fullSavePath)
		if err != nil {
			log.Fatalf("failed to create file: %s. err: %e", fullSavePath, err)
		}
	}
}

func copyFile(source, dst string) error {
	sourceFileStat, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Fatalf("%s is not a file", source)
	}

	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, src)

	return err
}

func readDirectory(path string) []os.DirEntry {
	entry, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("failed to read solution directory: %e", err)
	}

	return entry
}

func createDir(path string) {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to install solution generator: %e", err)
	}
}

func skip(path string) bool {
	for _, v := range skipDirs {
		if strings.Contains(path, v) {
			return true
		}
	}

	return false
}

func getGeneratorHomePath() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get user home directory")
	}

	return filepath.Join(dir, generatorHomeDirName)
}

func getSolutionSavePath(solutionName string) string {
	return filepath.Join(getGeneratorHomePath(), solutionName)
}
