package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Price struct {
	Dollar       int
	SekkeTamam   int
	SekketGhadim int
	SekkehNim    int
	RobeSekke    int
	Gold18       int
	GoldDast2    int
}

func main() {

	startTime := time.Now()
	p := new(Price)
	responseChannel := make(chan Price, 1024)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go p.httpGetUSD("https://www.tgju.org/profile/price_dollar_rl", "priceGold", responseChannel, wg)

	wg.Add(1)
	go p.httpGetFullCoin("https://www.tgju.org/profile/sekee", "priceGold", responseChannel, wg)

	wg.Add(1)
	go p.httpGetOldCoin("https://www.tgju.org/profile/sekeb", "priceGold", responseChannel, wg)

	wg.Add(1)
	go p.httpGetSemiCoin("https://www.tgju.org/profile/nim", "priceGold", responseChannel, wg)

	wg.Add(1)
	go p.httpGetQuarterGold("https://www.tgju.org/profile/rob", "priceGold", responseChannel, wg)

	wg.Add(1)
	go p.httpGet18New("https://www.tgju.org/profile/geram18", "priceGold", responseChannel, wg)

	wg.Add(1)
	go p.httpGet18Old("https://www.tgju.org/profile/gold_mini_size", "priceGold", responseChannel, wg)

	wg.Wait()
	close(responseChannel)
	finalPrice := new(Price)
	for responseChann := range responseChannel {
		finalPrice.Dollar = responseChann.Dollar + finalPrice.Dollar
		finalPrice.SekkeTamam = responseChann.SekkeTamam + finalPrice.SekkeTamam
		finalPrice.SekketGhadim = responseChann.SekketGhadim + finalPrice.SekketGhadim
		finalPrice.SekkehNim = responseChann.SekkehNim + finalPrice.SekkehNim
		finalPrice.RobeSekke = responseChann.RobeSekke + finalPrice.RobeSekke
		finalPrice.Gold18 = responseChann.Gold18 + finalPrice.Gold18
		finalPrice.GoldDast2 = responseChann.GoldDast2 + finalPrice.GoldDast2
	}

	log.Println("From channel is: ", finalPrice)
	log.Println("Total latency: ", time.Since(startTime))

}

func (p Price) httpGetUSD(url string, priceType string, responceChannel chan Price, wg *sync.WaitGroup) {
	var price int

	netClient := customHttpClient()

	responseByte, err := netClient.Get(url)

	httpErrorHandeler(err)

	responeBody, err := ioutil.ReadAll(responseByte.Body)
	byteReadErrorHandelete(err)

	responseString := string(responeBody)

	if priceType == "priceGold" {
		_, price = findSekkeTamam(responseString)
	}

	responseByte.Body.Close()

	p.Dollar = price

	responceChannel <- p
	wg.Done()

}

func (p Price) httpGetFullCoin(url string, priceType string, responceChannel chan Price, wg *sync.WaitGroup) {
	var price int

	netClient := customHttpClient()

	responseByte, err := netClient.Get(url)

	httpErrorHandeler(err)

	responeBody, err := ioutil.ReadAll(responseByte.Body)
	byteReadErrorHandelete(err)

	responseString := string(responeBody)

	if priceType == "priceGold" {
		_, price = findSekkeTamam(responseString)
	}

	responseByte.Body.Close()

	p.SekkeTamam = price

	responceChannel <- p
	wg.Done()

}

func (p Price) httpGetOldCoin(url string, priceType string, responceChannel chan Price, wg *sync.WaitGroup) {
	var price int

	netClient := customHttpClient()

	responseByte, err := netClient.Get(url)

	httpErrorHandeler(err)

	responeBody, err := ioutil.ReadAll(responseByte.Body)
	byteReadErrorHandelete(err)

	responseString := string(responeBody)

	if priceType == "priceGold" {
		_, price = findSekkeTamam(responseString)
	}

	responseByte.Body.Close()

	p.SekketGhadim = price

	responceChannel <- p
	wg.Done()

}

func (p Price) httpGetSemiCoin(url string, priceType string, responceChannel chan Price, wg *sync.WaitGroup) {
	var price int

	netClient := customHttpClient()

	responseByte, err := netClient.Get(url)

	httpErrorHandeler(err)

	responeBody, err := ioutil.ReadAll(responseByte.Body)
	byteReadErrorHandelete(err)

	responseString := string(responeBody)

	if priceType == "priceGold" {
		_, price = findSekkeTamam(responseString)
	}

	responseByte.Body.Close()

	p.SekkehNim = price

	responceChannel <- p
	wg.Done()

}

func (p Price) httpGetQuarterGold(url string, priceType string, responceChannel chan Price, wg *sync.WaitGroup) {
	var price int

	netClient := customHttpClient()

	responseByte, err := netClient.Get(url)

	httpErrorHandeler(err)

	responeBody, err := ioutil.ReadAll(responseByte.Body)
	byteReadErrorHandelete(err)

	responseString := string(responeBody)

	if priceType == "priceGold" {
		_, price = findSekkeTamam(responseString)
	}

	responseByte.Body.Close()

	p.RobeSekke = price

	responceChannel <- p
	wg.Done()

}

func (p Price) httpGet18New(url string, priceType string, responceChannel chan Price, wg *sync.WaitGroup) {
	var price int

	netClient := customHttpClient()

	responseByte, err := netClient.Get(url)

	httpErrorHandeler(err)

	responeBody, err := ioutil.ReadAll(responseByte.Body)
	byteReadErrorHandelete(err)

	responseString := string(responeBody)

	if priceType == "priceGold" {
		_, price = findSekkeTamam(responseString)
	}

	responseByte.Body.Close()

	p.Gold18 = price

	responceChannel <- p
	wg.Done()

}

func (p Price) httpGet18Old(url string, priceType string, responceChannel chan Price, wg *sync.WaitGroup) {
	var price int

	netClient := customHttpClient()

	responseByte, err := netClient.Get(url)

	httpErrorHandeler(err)

	responeBody, err := ioutil.ReadAll(responseByte.Body)
	byteReadErrorHandelete(err)

	responseString := string(responeBody)

	if priceType == "priceGold" {
		_, price = findSekkeTamam(responseString)
	}

	responseByte.Body.Close()

	p.GoldDast2 = price

	responceChannel <- p
	wg.Done()

}

func customHttpClient() http.Client {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: config,
	}

	netClient := &http.Client{
		Transport: transport,
	}
	return *netClient
}

func httpErrorHandeler(err error) error {
	if err != nil {
		fmt.Println("Cannot http call with error: ", err)
	}
	return err
}

func byteReadErrorHandelete(err error) error {
	if err != nil {
		fmt.Println("Cannot read as byte: ", err)
	}
	return err
}

func findSekkeTamam(html string) (string, int) {
	regex, _ := regexp.Compile("info.last_trade.PDrCotVal.*")
	price := regex.FindString(html)
	priceInt := priceCleaner(price)
	return price, priceInt
}

func priceCleaner(priceString string) int {
	regexInt, _ := regexp.Compile("[0-9].*")

	// Make Clean "info.last_trade.PDrCotVal">195,000,000</span>"
	priceByte := regexInt.FindString(priceString)
	someString := string(priceByte)

	// Make Clean "</span>" from "195,000,000</span>""
	someString2 := strings.Replace(someString, "</span>", "", -1)

	// Make Clean all "," in "195,000,000"
	priceInString := strings.Replace(someString2, ",", "", -1)

	// Make int format.
	price, _ := strconv.Atoi(priceInString)

	// Enable just for debug:
	// fmt.Println(price)

	return price
}
