package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
		}
	}()

	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s {domain} {_session_id} {_mastodon_session}", os.Args[0])
		return
	}

	url := fmt.Sprintf("https://%s/api/v1/timelines/public", os.Args[1])
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.AddCookie(&http.Cookie{
		Name:  "_session_id",
		Value: os.Args[2],
	})
	req.AddCookie(&http.Cookie{
		Name:  "_mastodon_session",
		Value: os.Args[3],
	})
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		panic(res.Status + "\n" + string(b))
	}

	var tl []struct {
		Account struct {
			Display_name string
			Acct         string
			Url          string
		}
		Content string
	}
	if err := json.Unmarshal(b, &tl); err != nil {
		panic(err)
	}

	for _, status := range tl {
		account := status.Account
		fmt.Printf("%v (@%s)\n", account.Display_name, account.Acct)
		fmt.Println("    " + removeHtmlTag(status.Content))
	}
}

func removeHtmlTag(html string) string {
	return regexp.MustCompile(`</?[^>]*>`).ReplaceAllString(html, "")
}
