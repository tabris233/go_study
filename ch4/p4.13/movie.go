package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Movie struct
type Movie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []Rating
	Metascore  string
	imdbRating string
	imdbVotes  string
	imdbID     string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}

// Rating struct
type Rating struct {
	Source string
	Value  string
}

var reqURL = "http://www.omdbapi.com/"

func rq(params []string) (*Movie, error) {
	q := url.QueryEscape(strings.Join(params, " "))

	resp, err := http.Get(reqURL + "?i=tt3896198&apikey=6e91b90c&t=" + q)
	fmt.Println(resp.Request.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func download(name, imageURL string) error {
	resp, _ := http.Get(imageURL)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create(name + ".jpg")
	io.Copy(out, bytes.NewReader(body))

	return nil
}

func main() {
	name := os.Args[1:]

	movie, err := rq(name)
	if err != nil {
		log.Fatal(err)
	}

	posterURL := movie.Poster

	fmt.Println("movie posterURL:", posterURL)

	download("movie_poster", posterURL)
}
