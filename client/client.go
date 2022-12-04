package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"

	"github.com/BenJetson/aoc-2022/aoc"
	"github.com/BenJetson/aoc-2022/utilities"
)

const aocHost = "adventofcode.com"

// GetSessionToken retrieves the session token from disk.
func getSessionToken() (string, error) {
	lines, err := utilities.ReadLinesFromFile(".aoc-session")
	if err != nil {
		return "", fmt.Errorf("cannot read session file: %w", err)
	} else if len(lines) != 1 {
		return "", errors.New("session file should have exactly one line")
	}

	return lines[0], nil
}

type Client struct {
	httpClient interface {
		Do(req *http.Request) (*http.Response, error)
	}
}

// New creates a new AOC client.
func New() (*Client, error) {
	sessionToken, err := getSessionToken()
	if err != nil {
		return nil, fmt.Errorf("could not get session token: %w", err)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("could not create client cookie jar: %w", err)
	}

	jar.SetCookies(&url.URL{
		Scheme: "https",
		Host:   aocHost,
	}, []*http.Cookie{
		{
			Name:  "session",
			Value: sessionToken,

			Path: "/",

			Secure:   true,
			HttpOnly: true,
		},
	})

	return &Client{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
			Jar:     jar,
		},
	}, nil
}

var puzzleTitleExp = regexp.MustCompile(
	`^\\?-\\?-\\?- Day [0-9]{1,2}: (.*) \\?-\\?-\\?-$`)

func (c *Client) GetPuzzleMarkdown(day, part int) (*aoc.PuzzleText, error) {
	u := url.URL{
		Scheme: "https",
		Host:   aocHost,
		Path:   fmt.Sprintf("/2022/day/%d", day),
	}

	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to craft puzzle request: %w", err)
	}

	res, err := c.httpClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("client failed to do puzzle request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"received non-OK status for puzzle request: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf(
			"could not parse document from puzzle page response: %w", err)
	}

	selection := doc.Find("article.day-desc")
	if selection.Length() != part {
		return nil, fmt.Errorf(
			"expected %d article element(s) for puzzle part %d, found %d",
			part, part, selection.Length(),
		)
	}

	var puzzle aoc.PuzzleText

	selection = selection.Eq(part - 1)
	converter := md.NewConverter(aocHost, true, &md.Options{
		CodeBlockStyle: "fenced",
	})
	converter.AddRules(md.Rule{
		Filter: []string{"h2"},
		Replacement: func(
			content string,
			_ *goquery.Selection,
			_ *md.Options,
		) *string {
			if part != 1 {
				return md.String("")
			}

			matches := puzzleTitleExp.FindStringSubmatch(content)

			puzzle.Title = matches[1]
			return md.String("")
		},
	})

	puzzle.Body = strings.Split(
		converter.Convert(selection),
		"\n",
	)

	return &puzzle, nil
}

func (c *Client) GetPuzzleInput(day int) (string, error) {
	u := url.URL{
		Scheme: "https",
		Host:   aocHost,
		Path:   fmt.Sprintf("/2022/day/%d/input", day),
	}

	r, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to craft input request: %w", err)
	}

	res, err := c.httpClient.Do(r)
	if err != nil {
		return "", fmt.Errorf("client failed to do input request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"received non-OK status for input request: %w", err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("cannot read input respone body: %w", err)
	}

	return string(data), nil
}

func (c *Client) SubmitAnswer(day, part, answer int) (string, error) {
	u := url.URL{
		Scheme: "https",
		Host:   aocHost,
		Path:   fmt.Sprintf("/2022/day/%d/answer", day),
	}

	data := make(url.Values)
	data.Set("level", strconv.Itoa(part))
	data.Set("answer", strconv.Itoa(answer))

	reqBody := strings.NewReader(data.Encode())

	r, err := http.NewRequest(http.MethodPost, u.String(), reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to craft answer request: %w", err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(r)
	if err != nil {
		return "", fmt.Errorf("client failed to do answer request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"received non-OK status for answer request: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", fmt.Errorf(
			"could not parse document from answer response: %w", err)
	}

	resultText := strings.Join(
		strings.Split(
			doc.Find("article").Text(),
			"  ",
		),
		"\n",
	)

	return resultText, nil
}
