package homework_day4

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	BaseUrl = "https://gateway.chotot.com/v1/public/ad-listing"
	CateVeh = "2000"
	CatePty = "1000"
)

type AdsResponse struct {
	Total int  `json:"total"`
	Ads   []Ad `json:"ads"`
}

type Ad struct {
	AdId int `json:"ad_id"`
	//TODO #1 Define struct
	// list_id , account_name, subject, list_time
	ListId      int    `json:"list_id"`
	AccountName string `json:"account_name"`
	Subject     string `json:"subject"`
	ListTime    int    `json:"list_time"`
}

type client struct {
	httpClient *http.Client
	baseUrl    string
	retryTimes int
	logger     *log.Logger
}

type Option func(*client)

func homework_day4() {
	log := log.New(os.Stdout, "PREFIX: ", log.LstdFlags)
	client := NewClient(BaseUrl, 2, log)
	ctx := context.TODO()

	// Add corountine & waitgroup
	var wg sync.WaitGroup

	// Call xe
	wg.Add(1)
	go func() {
		defer wg.Done()
		adResp, err := client.GetAdByCate(ctx, "2010", 1)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(adResp)
	}()

	// Call nha
	wg.Add(1)
	go func() {
		defer wg.Done()
		adResp, err := client.GetAdByCate(ctx, "1010", 2)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(adResp)
	}()

	wg.Wait()
}

func (c *client) GetAdByCate(ctx context.Context, cate string, worker int) (*AdsResponse, error) {
	fmt.Printf("GetAdByCate Request %d starting\n", worker)
	now := time.Now()
	defer func() {
		c.logger.Printf("GetAdByCate Request Done - Cate %v, Duration: %v", cate, time.Since(now).String())
	}()

	url := fmt.Sprintf("%v?cg=%v&limit=10", BaseUrl, cate)

	// TODO #3 implement retry if StatusCode = 5xx
	var resp *http.Response
	var err error

	for attemp := 0; attemp < c.retryTimes; attemp++ {
		resp, err = c.httpClient.Get(url)
		if resp.StatusCode >= 200 || resp.StatusCode < 300 {
			break
		} else if resp.StatusCode >= 500 || resp.StatusCode < 600 {
			// Wait 300ms to try a gain
			time.Sleep(300)
			continue
		}

		if err != nil {
			return nil, err
		}
	}

	var adResp AdsResponse
	// TODO #2 unmarshal json
	err = json.NewDecoder(resp.Body).Decode(&adResp)
	if err != nil {
		return nil, err
	}

	return &adResp, nil
}

func NewClient(baseUrl string, retryTimes int, log *log.Logger) *client {
	// TODO #4 refactor NewClient using functional options
	return &client{
		httpClient: http.DefaultClient,
		baseUrl:    baseUrl,
		retryTimes: retryTimes,
		logger:     log,
	}

}

func NewClientOption(opt ...Option) *client {
	c := new(client)
	for _, o := range opt {
		o(c)
	}

	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}

	if c.baseUrl == "" {
		c.baseUrl = BaseUrl
	}

	if c.retryTimes == 0 {
		c.retryTimes = 3
	}

	if c.logger == nil {
		c.logger = log.Default()
	}

	return c
}
