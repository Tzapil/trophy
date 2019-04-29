package user

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Game struct {
	Url  string `json:"url"`
	Icon string `json:"icon"`
}

type Trophy struct {
	Url         string `json:"url"`
	Icon        string `json:"icon"`
	Ttype       string `json:"type"`
	Percent     string `json:"percent"`
	Rare        string `json:"rare"`
	Owners      string `json:"owners"`
	Achievers   string `json:"achievers"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Log struct {
	Number string `json:"number"`
	Trophy Trophy `json:"trophy"`
	Game   Game   `json:"game"`
}

func CutLog(tr string) Log {
	reg, _ := regexp.Compile("<b>#(.*?)</b>")
	arr := reg.FindStringSubmatch(tr)

	return Log{arr[1], CutTrophy(tr), CutGame(tr)}
}

func CutTrophy(tr string) Trophy {
	reg, _ := regexp.Compile("<a href=\"(.*?)\" class=\".*?\">[\\s\\S]*?<img class=\"trophy\" src=\"(.*?)\"[\\s\\S]*?</a>")
	arr := reg.FindStringSubmatch(tr)

	regT, _ := regexp.Compile("<img title=\"(.*?)\"")
	arrT := regT.FindStringSubmatch(tr)

	regP, _ := regexp.Compile("<span class=\".*?\">(.*?%)</span><br /><span class=\"typo-bottom\"><nobr>(.*?)</nobr>")
	arrP := regP.FindStringSubmatch(tr)

	regO, _ := regexp.Compile("<span class=\".*?\">(.*?)</span><br /><span class=\"typo-bottom\"><nobr>(?:DLC )?Owners</nobr>")
	arrO := regO.FindStringSubmatch(tr)

	regA, _ := regexp.Compile("<span class=\".*?\">(.*?)</span><br /><span class=\"typo-bottom\"><nobr>Achievers</nobr>")
	arrA := regA.FindStringSubmatch(tr)

	regD, _ := regexp.Compile("<a class=\"title\" href=\"(.*?)\">(.*?)</a><br />(.*?)\\s+</td>")
	arrD := regD.FindStringSubmatch(tr)

	return Trophy{baseURI + arr[1], arr[2], arrT[1], arrP[1], arrP[2], arrO[1], arrA[1], arrD[2], arrD[3]}
}

func CutGame(tr string) Game {
	reg, _ := regexp.Compile("<a href=\"(.*?)\">[\\s\\S]*?<img class=\"game\" src=\"(.*?)\"[\\s\\S]*?</a>")
	arr := reg.FindStringSubmatch(tr)

	return Game{baseURI + arr[1], arr[2]}
}

func GetLogs(name string) []Log {
	resp, err := http.Get(baseURI + "/" + name + "/log")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	reg, _ := regexp.Compile("<tr\\s?>[\\s\\S]*?</tr>")

	s := string(body[:])
	trs := reg.FindAllString(s, -1)

	result := make([]Log, 0, len(trs))

	for i := 0; i < len(trs); i++ {
		tr := trs[i]
		log := CutLog(tr)
		result = append(result, log)
	}

	return result
}
