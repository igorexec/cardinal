package collect

import (
	"fmt"
	"github.com/go-chi/render"
	"github.com/igorexec/cardinal/app/store"
	"io"
	"log"
	"net/http"
	"time"
)

const googlePageSpeedAPI = "https://www.googleapis.com/pagespeedonline/v5/runPagespeed"

type PageSpeedCollector struct {
	Token string
}

func NewPageSpeedCollector(token string) *PageSpeedCollector {
	return &PageSpeedCollector{Token: token}
}

func (c *PageSpeedCollector) Do(page string) (result store.PageSpeed, err error) {
	log.Printf("[info] pagespeed collector started for %s", page)

	url := fmt.Sprintf("%s?url=%s&key=%s", googlePageSpeedAPI, page, c.Token)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("[error] failed to collect data for %s: %v", page, err)
		return result, err
	}

	score, err := parsePageSpeed(resp.Body)
	if err != nil {
		log.Fatalf("[error] failed to decode: %v", err)
		return result, err
	}

	log.Printf("[info] pagespeed collector for %s finished", page)

	return store.PageSpeed{
		Score: int(score * 100),
		Page:  page,
		Date:  time.Now().UTC(),
	}, nil
}

func parsePageSpeed(body io.ReadCloser) (float32, error) {
	apiData := struct {
		LighthouseResult struct {
			Categories struct {
				Performance struct {
					Score float32 `json:"score"`
				} `json:"performance"`
			} `json:"categories"`
		} `json:"lighthouseResult"`
	}{}

	if err := render.DecodeJSON(body, &apiData); err != nil {
		return 0, err
	}
	return apiData.LighthouseResult.Categories.Performance.Score, nil
}
