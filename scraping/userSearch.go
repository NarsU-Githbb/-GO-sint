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
	client            *http.Client
	secondLifeLimiter *rate.Limiter
	xLimiter          *rate.Limiter
	linkedInLimiter   *rate.Limiter
	instagramLimiter  *rate.Limiter
	githubLimiter     *rate.Limiter
}

func NewUnScraper() *UnScraper {
	return &UnScraper{
		client: &http.Client{
			// Always assign a reasonable timeout so CLI tool doesn't hang indefinitely
			Timeout: 7 * time.Second,
		},

		// find the CORRECT cooldown time for every platform
		secondLifeLimiter: rate.NewLimiter(rate.Limit(5), 5),
		xLimiter:          rate.NewLimiter(rate.Every(2*time.Second), 1),
		linkedInLimiter:   rate.NewLimiter(rate.Every(2*time.Second), 1),
		instagramLimiter:  rate.NewLimiter(rate.Every(1*time.Second), 1),
		githubLimiter:     rate.NewLimiter(rate.Limit(3), 3),
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

func (scrp *UnScraper) SecondLifeCheck(ctx context.Context, username string) (bool, error) {

	url := fmt.Sprintf("https://my.secondlife.com/%s", username)

	//Wait for the rate limiter before making the request

	if err := scrp.secondLifeLimiter.Wait(ctx); err != nil {
		return false, fmt.Errorf("rate limiter error: %w", err)
	}

	req, err := scrp.prepareRequest(ctx, url)
	if err != nil {
		return false, err
	}

	// Execute the request
	resp, err := scrp.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close() // Protect against resource/connection leaks

	switch resp.StatusCode {
	case http.StatusOK: // 200
		return true, nil
	case http.StatusNotFound: // 404
		return false, nil
	default:
		return false, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}
}

func (scrp *UnScraper) XCheck(ctx context.Context, username string) (bool, error) {
	url := fmt.Sprintf("https://x.com/%s", username)

	if err := scrp.xLimiter.Wait(ctx); err != nil {
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

func (scrp *UnScraper) LinkedInCheck(ctx context.Context, username string) (bool, error) {
	url := fmt.Sprintf("https://linkedin.com/in/%s", username)

	if err := scrp.linkedInLimiter.Wait(ctx); err != nil {
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
	case 999: //find similar specific error codes for other platforms
		return false, fmt.Errorf("blocked by LinkedIn anti-scraping wall (Status 999)")
	default:
		return false, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}
}

func (scrp *UnScraper) InstagramCheck(ctx context.Context, username string) (bool, error) {
	url := fmt.Sprintf("https://www.instagram.com/%s", username)

	if err := scrp.instagramLimiter.Wait(ctx); err != nil {
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

func (scrp *UnScraper) GithubCheck(ctx context.Context, username string) (bool, error) {
	url := fmt.Sprintf("https://github.com/%s", username)

	if err := scrp.githubLimiter.Wait(ctx); err != nil {
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

func (scrp *UnScraper) ExecuteUsernameSearch(ctx context.Context, username string) (models.UserModel, error) {

	secondLife, _ := scrp.SecondLifeCheck(ctx, username)
	x, _ := scrp.XCheck(ctx, username)
	linkedin, _ := scrp.LinkedInCheck(ctx, username)
	instagram, _ := scrp.InstagramCheck(ctx, username)
	github, _ := scrp.GithubCheck(ctx, username)

	result := models.UserModel{
		SecondLife: secondLife,
		X:          x,
		LinkedIn:   linkedin,
		Instagram:  instagram,
		Github:     github,
	}

	return result, nil

}
