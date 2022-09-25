package scrappy

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Scrap(target string) string {
	res, err := http.Get(target)

	if err != nil {
		log.Println("error occured: ", err)
		return ""
	}

	// log.Println("Response: ", res)
	return Load(res)
}

func Load(res *http.Response) string {
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Println("Error: ", err)
		return ""
	}

	// log.Println("Response: ", doc)

	var data []string

	doc.Find("._25b18c").Each(func(i int, s *goquery.Selection) {
		price := s.Find("div").Text()
		data = append(data, price)
		// log.Println(price)
	})

	first_entry := data[0]
	// log.Println(first_entry)

	price := ExtractPrice(first_entry)
	log.Println(price)

	return price
}

func ExtractPrice(data string) string {
	start, end := 0, 0
	count := 0
	for i, v := range data {
		if v == '₹' {
			count++
			if count%2 == 0 {
				end = i
				break
			}
			start = i
		}
	}
	return data[start:end]
}
