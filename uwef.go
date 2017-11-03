package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"log"
	"path"
	"strings"
	"errors"
	"encoding/json"
	"os"
)

type ConfigParams struct {
	folder 		*string
}

func main() {
	params := getConfigParams()

	configFileName, err := retrieveBuildConfigName(params)
	handleError(err)

	configFile, err := loadConfigFileContents(configFileName)
	handleError(err)

	fixedConfigFile := fixConfigFileProperties(configFile)

	err = overwriteConfigFile(configFileName, fixedConfigFile)
	handleError(err)

	fmt.Println(configFileName, ": fixed!")
}

func getConfigParams() *ConfigParams {
	configParams := &ConfigParams{}

	configParams.folder = flag.String("folder", "./", "Project root folder")

	flag.Parse()
	return configParams
}

func retrieveBuildConfigName(params *ConfigParams) (string, error) {
	buildFolder := path.Join(*params.folder, "Build")

	files, err := ioutil.ReadDir(buildFolder)

	if err != nil {
		return "", err
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".json") {
			return path.Join(buildFolder, file.Name()), nil
		}
	}

	return "", errors.New("build config file not found")
}

type ConfigFile struct {
	TotalMemory			int64		`json:"TOTAL_MEMORY"`
	DataUrl				string		`json:"dataUrl"`
	AsmCodeUrl			string		`json:"asmCodeUrl"`
	AsmMemoryUrl		string		`json:"asmMemoryUrl"`
	AsmFrameworkUrl		string		`json:"asmFrameworkUrl"`
	SplashScreenStyle	string		`json:"splashScreenStyle"`
	BackgroundColor		string		`json:"backgroundColor"`
}

func loadConfigFileContents(filePath string) (*ConfigFile, error) {
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		return &ConfigFile{}, err
	}

	var config ConfigFile
	err = json.Unmarshal(file, &config)

	if err != nil {
		return &ConfigFile{}, err
	}

	return &config, nil
}

func fixConfigFileProperties(config *ConfigFile) *ConfigFile {
	return &ConfigFile{
		config.TotalMemory,
		fixPropertyIfNeeded(config.DataUrl),
		fixPropertyIfNeeded(config.AsmCodeUrl),
		fixPropertyIfNeeded(config.AsmMemoryUrl),
		fixPropertyIfNeeded(config.AsmFrameworkUrl),
		config.SplashScreenStyle,
		config.BackgroundColor,
	}
}

func overwriteConfigFile(path string, config *ConfigFile) error {
	jsonConfig, err := json.Marshal(config)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, jsonConfig, 0755)
}

func fixPropertyIfNeeded(property string) string {
	if !strings.Contains(property, ".gz") {
		return property + ".gz"
	}
	return property
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}