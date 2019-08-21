package collector

import (
	"fmt"
	"github.com/go-chi/render"
	"github.com/go-pkgz/lgr"
	"github.com/icheliadinski/cardinal/store"
	"io"
	"net/http"
	"time"
)

type PageSpeedAPI struct {
	URL   string
	Token string
}

type Collector struct {
	PageSpeedAPI PageSpeedAPI
}

func (c *Collector) CollectPageSpeed(page string) (result store.PageSpeed, err error) {
	lgr.Printf("[INFO] pagespeed collector started")

	url := fmt.Sprintf("%s?url=%s", c.PageSpeedAPI, page)
	resp, err := http.Get(url)
	if err != nil {
		lgr.Printf("[ERROR] failed to collect data from %s", url)
		return result, err
	}

	score, err := c.mapBodyToScore(resp.Body)
	if err != nil {
		lgr.Printf("[INFO] can't map google API response to score")
		return result, err
	}

	lgr.Printf("[INFO] pagespeed collector finished")

	result = store.PageSpeed{
		Page:  page,
		Score: int(score * 100),
		Date:  time.Now().UTC(),
	}
	return result, nil
}

func (c *Collector) mapBodyToScore(body io.ReadCloser) (float32, error) {
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
