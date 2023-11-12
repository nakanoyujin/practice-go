package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	Month    = time.Now().AddDate(0, -1, 0) //一ヶ月前
	HalfYear = time.Now().AddDate(0, -6, 0) // 半年前
	Year     = time.Now().AddDate(-1, 0, 0) // 1年前
)

const IssuesURL = "https://api.github.com/search/issues"

//example https://api.github.com/search/issues?q=quic-go/quic-go

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*issues
}

func (Isr IssuesSearchResult) String() string {
	return fmt.Sprintf("%d", Isr.TotalCount)
}

type issues struct {
	Number    int
	Html_url  string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login string
	Id    int `json:"id"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	//unmarshalの使い分けに関しては以下を参照　https://qiita.com/Coolucky/items/44f2bc6e32ca8e9baa96
	request_url := IssuesURL + "?q=" + q
	fmt.Println(request_url)
	resp, err := http.Get(request_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < result.TotalCount; i++ {

		switch {
		case result.Items[i].CreatedAt.After(Month):
			fmt.Println("一ヶ月以内:", result.Items[i].User.Id)
		case result.Items[i].CreatedAt.After(HalfYear):
			fmt.Println("半年以内:", result.Items[i].Html_url)
		case result.Items[i].CreatedAt.After(Year):
			fmt.Println("一年以内:", result.Items[i].Html_url)
		default:
			fmt.Println("一年以上前:", result.Items[i].Html_url)
		}

	}

}
