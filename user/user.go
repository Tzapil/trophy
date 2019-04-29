package user

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type User struct {
	Name        string `json:"name"`
	Played      string `json:"played"`
	Completed   string `json:"completed"`
	Completion  string `json:"completion"`
	Unearned    string `json:"unearned"`
	PerDay      string `json:"per_day"`
	Rank        string `json:"rank"`
	CountryRank string `json:"country_rank"`
}

func GetUser(name string) User {
	resp, err := http.Get(baseURI + "/" + name)
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

	// Games Played
	regP, _ := regexp.Compile("<span class=\".*?\">(.*?)<span>Games Played</span></span>")
	arrP := regP.FindStringSubmatch(s)

	// Completed Games
	regC, _ := regexp.Compile("<span class=\".*?\">(.*?)<span>Completed Games</span></span>")
	arrC := regC.FindStringSubmatch(s)

	// Completion
	regCP, _ := regexp.Compile("<span class=\".*?\">(.*?)<span>Completion</span></span>")
	arrCP := regCP.FindStringSubmatch(s)

	// Unearned Trophies
	regT, _ := regexp.Compile("<span class=\".*?\">(.*?)<span>Unearned Trophies</span></span>")
	arrT := regT.FindStringSubmatch(s)

	// Trophies Per Day
	regTPD, _ := regexp.Compile("<span class=\".*?\">(.*?)<span>Trophies Per Day</span></span>")
	arrTPD := regTPD.FindStringSubmatch(s)

	// World Rank
	regR, _ := regexp.Compile("rel=\"nofollow\">\\s+([\\s\\S]*?)<span>World Rank</span>")
	arrR := regR.FindStringSubmatch(s)

	// Country Rank
	regCR, _ := regexp.Compile("country-rank[\\s\\S]*?rel=\"nofollow\">\\s+([\\s\\S]*?)<span>Country Rank</span>")
	arrCR := regCR.FindStringSubmatch(s)

	result := User{name, arrP[1], arrC[1], arrCP[1], arrT[1], arrTPD[1], arrR[1], arrCR[1]}

	return result
}
