package scraping

import (
	"context"
	"digital_detective/models"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type UnScraper struct {
	client         *http.Client
	xLimiter       *rate.Limiter
	youtubeLimiter *rate.Limiter
	scLimiter      *rate.Limiter
	githubLimiter  *rate.Limiter
	giteaLimiter   *rate.Limiter
}

func NewUnScraper() *UnScraper {
	return &UnScraper{
		client: &http.Client{
			Timeout: 7 * time.Second,
		},

		// Conservative local pacing defaults
		xLimiter:       rate.NewLimiter(rate.Every(3*time.Second), 1),
		youtubeLimiter: rate.NewLimiter(rate.Every(3*time.Second), 1),
		scLimiter:      rate.NewLimiter(rate.Every(3*time.Second), 1),
		githubLimiter:  rate.NewLimiter(rate.Every(time.Minute), 1),
		giteaLimiter:   rate.NewLimiter(rate.Every(3*time.Second), 1),
	}
}

func (scrp *UnScraper) prepareRequest(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	return req, nil
}

func (scrp *UnScraper) checkPlatform(ctx context.Context, limiter *rate.Limiter, url string) (bool, error) {
	if err := limiter.Wait(ctx); err != nil {
		return false, fmt.Errorf("rate limiter error: %w", err)
	}

	req, err := scrp.prepareRequest(ctx, url)
	if err != nil {
		return false, err
	}

	resp, err := scrp.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return true, nil
	case http.StatusNotFound:
		return false, nil
	default:
		return false, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}
}

func (scrp *UnScraper) XCheck(ctx context.Context, username string) (bool, error) {
	return scrp.checkPlatform(ctx, scrp.xLimiter, fmt.Sprintf("https://x.com/%s", username))
}

func (scrp *UnScraper) YoutubeCheck(ctx context.Context, username string) (bool, error) {
	return scrp.checkPlatform(ctx, scrp.youtubeLimiter, fmt.Sprintf("https://www.youtube.com/user/%s", username))
}

func (scrp *UnScraper) SoundcloudCheck(ctx context.Context, username string) (bool, error) {
	return scrp.checkPlatform(ctx, scrp.scLimiter, fmt.Sprintf("https://soundcloud.com/%s", username))
}

func (scrp *UnScraper) GithubCheck(ctx context.Context, username string) (bool, error) {
	return scrp.checkPlatform(ctx, scrp.githubLimiter, fmt.Sprintf("https://github.com/%s", username))
}

func (scrp *UnScraper) GiteaCheck(ctx context.Context, username string) (bool, error) {
	return scrp.checkPlatform(ctx, scrp.giteaLimiter, fmt.Sprintf("https://gitea.com/%s", username))
}

func (scrp *UnScraper) ExecuteUsernameSearch(ctx context.Context, username string) (models.UserModel, error) {

	var failures []string

	// If the context is already canceled, bail out early
	if err := ctx.Err(); err != nil {
		return models.UserModel{}, err
	}

	x, err := scrp.XCheck(ctx, username)
	if err != nil {
		failures = append(failures, fmt.Sprintf("xCheck failed! %s", err))
	}

	youtube, err := scrp.YoutubeCheck(ctx, username)
	if err != nil {
		failures = append(failures, fmt.Sprintf("YoutubeCheck failed! %s", err))
	}

	soundcloud, err := scrp.SoundcloudCheck(ctx, username)
	if err != nil {
		failures = append(failures, fmt.Sprintf("SoundcloudCheck failed! %s", err))
	}

	github, err := scrp.GithubCheck(ctx, username)
	if err != nil {
		failures = append(failures, fmt.Sprintf("GithubCheck failed! %s", err))
	}

	gitea, err := scrp.GiteaCheck(ctx, username)
	if err != nil {
		failures = append(failures, fmt.Sprintf("GiteaCheck failed! %s", err))
	}

	result := models.UserModel{
		X:          x,
		Youtube:    youtube,
		Soundcloud: soundcloud,
		Github:     github,
		Gitea:      gitea,
		Errors:     failures,
	}

	return result, nil

}
