package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"unicode"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

const (
	logo = `
██╗   ██╗██████╗ ██╗ ██████╗████████╗
╚██╗ ██╔╝██╔══██╗██║██╔════╝╚══██╔══╝
 ╚████╔╝ ██║  ██║██║██║        ██║   
  ╚██╔╝  ██║  ██║██║██║        ██║   
   ██║   ██████╔╝██║╚██████╗   ██║   
   ╚═╝   ╚═════╝ ╚═╝ ╚═════╝   ╚═╝   

YDict V0.8
https://github.com/TimothyYe/ydict

`
)

func displayUsage() {
	color.Cyan(logo)
	color.Cyan("Usage: ydict <word(s) to query>")
}

func isChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func isMacOS() bool {
	return runtime.GOOS == "darwin"
}

func getExecutePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return filepath.Dir(ex)
}

func loadEnv() {
	exPath := getExecutePath()
	envPath := fmt.Sprintf("%s/.env", exPath)

	// if .env file doesn't exist, just return
	if _, err := os.Stat(fmt.Sprintf("%s/.env", exPath)); os.IsNotExist(err) {
		return
	}

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	proxy = os.Getenv("SOCKS5")
}

func parseArgs(args []string) ([]string, bool) {
	//match argument: -v
	if args[len(args)-1] == "-v" {
		return args[1 : len(args)-1], true
	}

	return args[1:], false
}
