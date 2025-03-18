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

	responseChannel := make(chan int, 2048)
	wg := &sync.WaitGroup{}

	// usdPrice :=
	go httpGetUSD("https://www.tgju.org/profile/price_dollar_rl", "priceGold", responseChannel, wg)
	wg.Add(1)

	// // sekkeTamamPrice :=
	go httpGetFullCoin("https://www.tgju.org/profile/sekee", "priceGold", responseChannel, wg)
	// // price.SekkeTamam = sekkeTamamPrice
	wg.Add(1)

	// // sekkeGhadimPrice :=
	go httpGetOldCoin("https://www.tgju.org/profile/sekeb", "priceGold", responseChannel, wg)
	// // price.SekketGhadim = sekkeGhadimPrice
	wg.Add(1)

	// // SekkehNimPrice :=
	go httpGetSemiCoin("https://www.tgju.org/profile/nim", "priceGold", responseChannel, wg)
	// // price.SekkehNim = SekkehNimPrice
	wg.Add(1)

	// // SekkehRobePrice :=
	go httpGetQuarterGold("https://www.tgju.org/profile/rob", "priceGold", responseChannel, wg)
	// // price.RobeSekke = SekkehRobePrice
	wg.Add(1)

	// // Gold18 :=
	go httpGet18New("https://www.tgju.org/profile/geram18", "priceGold", responseChannel, wg)
	// // price.Gold18 = Gold18
	wg.Add(1)

	// // GoldDast2 :=
	go httpGet18Old("https://www.tgju.org/profile/gold_mini_size", "priceGold", responseChannel, wg)
	// // price.GoldDast2 = GoldDast2
	wg.Add(1)

	wg.Wait()
	close(responseChannel)

	for responseChann := range responseChannel {
		fmt.Println(responseChann)
	}

	// fmt.Println(price)
	log.Println("Total latency: ", time.Since(startTime))

}

func httpGetUSD(url string, priceType string, responceChannel chan int, wg *sync.WaitGroup) {
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

	responceChannel <- price
	wg.Done()

}

func httpGetFullCoin(url string, priceType string, responceChannel chan int, wg *sync.WaitGroup) {
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

	responceChannel <- price
	wg.Done()

}

func httpGetOldCoin(url string, priceType string, responceChannel chan int, wg *sync.WaitGroup) {
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

	responceChannel <- price
	wg.Done()

}

func httpGetSemiCoin(url string, priceType string, responceChannel chan int, wg *sync.WaitGroup) {
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

	responceChannel <- price
	wg.Done()

}

func httpGetQuarterGold(url string, priceType string, responceChannel chan int, wg *sync.WaitGroup) {
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

	responceChannel <- price
	wg.Done()

}

func httpGet18New(url string, priceType string, responceChannel chan int, wg *sync.WaitGroup) {
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

	responceChannel <- price
	wg.Done()

}

func httpGet18Old(url string, priceType string, responceChannel chan int, wg *sync.WaitGroup) {
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

	responceChannel <- price
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
	fmt.Println(price)
	return price
}
