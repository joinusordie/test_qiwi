package app

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"

	"example.com/m/internal/xmls"
	"golang.org/x/net/html/charset"
)

func GetExchangeRate(data []string) ([]string, error) {
	_, code, _ := strings.Cut(data[1], "=")
	_, date, _ := strings.Cut(data[2], "=")
	splittedDate := strings.Split(date, "-")

	url := fmt.Sprintf("http://www.cbr.ru/scripts/XML_daily.asp?date_req=%s/%s/%s", splittedDate[2], splittedDate[1], splittedDate[0])

	curRates, err := getCurrenciesRatesFromURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get currencies rates: %v", err)
	}

	for _, elem := range curRates.Valute {
		if elem.CharCode == code {
			return []string{code, elem.Name, elem.Value}, nil
		}
	}

	return nil, fmt.Errorf("failed to find your currency")
}

func getCurrenciesRatesFromURL(url string) (*xmls.ValCurs, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Chrome/23.0.1271.64")

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get xml scheme: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", response.StatusCode)
	}

	exchangeRates := new(xmls.ValCurs)
	decoder := xml.NewDecoder(response.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&exchangeRates)

	return exchangeRates, err
}
