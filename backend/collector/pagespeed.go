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

type PageSpeed struct {
}

func (s *PageSpeed) Collect(page string) (result store.PageSpeed, err error) {
	lgr.Printf("[INFO] pagespeed collector started")

	url := fmt.Sprintf("https://www.googleapis.com/pagespeedonline/v5/runPagespeed?url=%s", page)
	resp, err := http.Get(url)
	if err != nil {
		lgr.Printf("[ERROR] failed to collect data from %s", url)
		return result, err
	}

	score, err := s.mapBodyToScore(resp.Body)
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

func (s *PageSpeed) mapBodyToScore(body io.ReadCloser) (float32, error) {
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
