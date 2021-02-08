package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

const (
	IN_ONE_MONTH = iota
	IN_ONE_YEAR
	OVER_ONE_YEAR
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	tagName := map[int]string{
		IN_ONE_MONTH:  "one_month",
		IN_ONE_YEAR:   "one_year",
		OVER_ONE_YEAR: "over_one_year",
	}

	for _, item := range result.Items {
		var tag int
		if time.Now().Sub(item.CreatedAt) <= time.Hour*24*30 {
			tag = IN_ONE_MONTH
		} else if time.Now().Sub(item.CreatedAt) <= time.Hour*24*365 {
			tag = IN_ONE_YEAR
		} else {
			tag = OVER_ONE_YEAR
		}

		fmt.Printf("#%-5d [%s] %9.9s %.55s\n",
			item.Number, tagName[tag], item.User.Login, item.Title)
		// fmt.Println(item)
		// fmt.Println("---------------------")
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
}
