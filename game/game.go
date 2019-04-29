package game

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Trophy struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Achievers   string `json:"achievers"`
	Rare        string `json:"rare"`
}

type Game struct {
	Name       string   `json:"name"`
	Owners     string   `json:"owners"`
	Achievers  string   `json:"achievers"`
	Completion string   `json:"completion"`
	Trophies   []Trophy `json:"trophies"`
}

type SearchResult struct {
	name string
	url  string
}

// https://psnprofiles.com/search/games?q=nioh
func FindGameLink(name string) SearchResult {
	resp, err := http.Get(baseURI + "/search/games?q=" + name)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	s := string(body[:])

	reg, _ := regexp.Compile("<a class=\"title\" href=\"(.*?)\">(.*?)</a>")
	arr := reg.FindStringSubmatch(s)

	return SearchResult{arr[2], baseURI + arr[1]}
}

func CutTrophy(tr string) Trophy {
	regN, _ := regexp.Compile("<a class=\"title\" href=\"(.*?)\">(.*?)</a><br />(.*?)\\s+</td>")
	arrN := regN.FindStringSubmatch(tr)

	regP, _ := regexp.Compile("<picture[\\s\\S]*?<img src=\"(.*?)\" />[\\s\\S]*?</picture>")
	arrP := regP.FindStringSubmatch(tr)

	regR, _ := regexp.Compile("<span class=\"typo-top\">([0-9\\.%]*?)</span><br /><span class=\"typo-bottom\"><nobr>(.*?)</nobr></span>")
	arrR := regR.FindStringSubmatch(tr)

	return Trophy{arrN[2], arrN[1], arrN[3], arrP[1], arrR[1], arrR[2]}
}

func GetGame(name string) Game {
	searchResult := FindGameLink(name)
	resp, err := http.Get(searchResult.url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("RESPONSE: %+v\n", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	s := string(body[:])

	// Game Owners
	regO, _ := regexp.Compile("<span class=\"stat grow\">(.*?)<span>Game Owners</span></span>")
	arrO := regO.FindStringSubmatch(s)

	// Platinum Achievers
	regA, _ := regexp.Compile("<span class=\"stat grow\">(.*?)<span>Platinum Achievers</span></span>")
	arrA := regA.FindStringSubmatch(s)

	// Average Completion
	regC, _ := regexp.Compile("<span class=\"stat grow\">(.*?)<span>Average Completion</span></span>")
	arrC := regC.FindStringSubmatch(s)

	reg, _ := regexp.Compile("<tr class=\"\">[\\s\\S]*?<a class=\"title\"[\\s\\S]*?</tr>")
	trs := reg.FindAllString(s, -1)

	result := make([]Trophy, 0, len(trs))

	for i := 0; i < len(trs); i++ {
		tr := trs[i]
		trophy := CutTrophy(tr)
		result = append(result, trophy)
	}

	return Game{searchResult.name, arrO[1], arrA[1], arrC[1], result}
}
