package main

import (
	"log"
	"time"

	"github.com/adityasunny1189/scrappy/notifier"
	"github.com/adityasunny1189/scrappy/scrappy"
)

const (
	target = "https://www.flipkart.com/apple-2020-macbook-air-m1-8-gb-512-gb-ssd-mac-os-big-sur-mgn73hn-a/p/itm37da92e833fa3?pid=COMFXEKM8GQJW5DE&lid=LSTCOMFXEKM8GQJW5DEWZL2XD&marketplace=FLIPKART&q=macbook+air&store=6bo%2Fb5g&srno=s_1_1&otracker=AS_QueryStore_OrganicAutoSuggest_1_7_na_na_na&otracker1=AS_QueryStore_OrganicAutoSuggest_1_7_na_na_na&fm=organic&iid=490649e0-3bf6-4620-b541-871acb63ef21.COMFXEKM8GQJW5DE.SEARCH&ppt=hp&ppn=homepage&ssid=bqdaes07q80000001664110771842&qH=b61d62051d5441f9"
)

func main() {
	log.Println("Hello World")

	log.Println("Initiating Scrappy")

	oldPrice := scrappy.Scrap(target)
	SendNotification(oldPrice)

	for {
		newPrice := scrappy.Scrap(target)
		if newPrice != oldPrice {
			log.Println(newPrice)
			SendNotification(newPrice)
			oldPrice = newPrice
		}
		time.Sleep(5 * time.Second)
	}
}

func SendNotification(price string) {
	log.Println("Sending Notification")

	msg := "Price of macbook air: " + price

	notifier.Notify(msg)
}
