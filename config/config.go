// Package config holds everything related to configuration
package config

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
)

// PercentAnswer is the probability of anwser to any update
var PercentAnswer = 20

// init token
func init() {
	tokenFile, err := os.Open("datafiles/token")
	if err != nil {
		log.Panic(err)
	}
	defer tokenFile.Close()
	scanner := bufio.NewScanner(tokenFile)
	if scanner.Scan() {
		Token = string(scanner.Text())
	} else {
		panic("invalid token in token file content")
	}
}

// global variables
var (
	// Enabled defines if the bot should answer
	Enabled = true
	// Token is the token value of the bot
	Token string
)

// init "cion" ended words list and phrases list
func init() {
	cionFile, err := os.Open("datafiles/cion.csv")
	if err != nil {
		panic(err)
	}
	defer cionFile.Close()
	phrasesFile, err := os.Open("datafiles/phrases.csv")
	if err != nil {
		panic(err)
	}
	defer phrasesFile.Close()
	// init cion ended words map
	rc := csv.NewReader(cionFile)
	cionSlice, err := rc.Read()
	if err != nil {
		panic(err)
	}
	CionList = make(map[string]bool)
	for _, word := range cionSlice {
		CionList[word] = true
	}
	// init phrases list
	rp := csv.NewReader(phrasesFile)
	Phrases, err = rp.Read()
	if err != nil {
		panic(err)
	}
}

var (
	// Phrases is the list of predefined phrases to say randomly
	Phrases = []string{}
	// CionList is the list of words ending with "cion"
	CionList map[string]bool
)
