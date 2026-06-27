package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning the Url ya API or Link ")
	myurl := "http://localhost:8080/health"
	// fmt.Printf("%T",myurl)

	parseUrl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("cannot pass url ", err)
		return
	}
	fmt.Printf("url type is %T\n", parseUrl)
	fmt.Println("url Host", parseUrl.Host)
	fmt.Println("url  Schema", parseUrl.Path)
	fmt.Println("url  RawQuery", parseUrl.RawQuery)

	// modify the url components
	parseUrl.Path = "/newPath"
	parseUrl.RawQuery = "username=iamjacky"

	// constructing a url String from a url object
	newURL := parseUrl.String()
	fmt.Println("new Url : ", newURL)
}
