package scraping

import (
	"context"
	"digital_detective/models"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"time"

	"golang.org/x/time/rate"
)

// ip-api.com allows 45 requests per minute maximum
// ipwhois & abuseipdb 1,000 request per day

type IpScraper struct {
	client         *http.Client
	ipApiLimiter   *rate.Limiter
	ipWhoIsLimiter *rate.Limiter
	ipAbuseLimiter *rate.Limiter
	abuseApiKey    string
}

func NewIpScraper() *IpScraper {
	return &IpScraper{
		client:         &http.Client{Timeout: 5 * time.Second},
		ipApiLimiter:   rate.NewLimiter(rate.Every(2*time.Second), 1),
		ipWhoIsLimiter: rate.NewLimiter(rate.Every(2*time.Second), 1),
		ipAbuseLimiter: rate.NewLimiter(rate.Every(2*time.Second), 1),
		abuseApiKey:    os.Getenv("ABUSEIPDB_API_KEY"),
	}
}

func (scrp *IpScraper) prepareRequest(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (scrp *IpScraper) IpApiCheck(ctx context.Context, ip string) (*models.IpApiResponse, error) {

	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)

	if err := scrp.ipApiLimiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limiter error: %w", err)
	}

	req, err := scrp.prepareRequest(ctx, url)
	if err != nil {
		return nil, err
	}

	resp, err := scrp.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data models.IpApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (scrp *IpScraper) IpWhoIsCheck(ctx context.Context, ip string) (*models.IpWhoIsResponse, error) {

	url := fmt.Sprintf("https://ipwhois.app/json/%s", ip)

	if err := scrp.ipWhoIsLimiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limiter error: %w", err)
	}

	req, err := scrp.prepareRequest(ctx, url)
	if err != nil {
		return nil, err
	}

	resp, err := scrp.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data models.IpWhoIsResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (scrp *IpScraper) IpAbuseCheck(ctx context.Context, ip string) (*models.IPAbuseReport, error) {

	url := fmt.Sprintf("https://api.abuseipdb.com/api/v2/check?ipAddress=%s", ip)

	if err := scrp.ipAbuseLimiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limiter error: %w", err)
	}

	req, err := scrp.prepareRequest(ctx, url)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Key", scrp.abuseApiKey)

	resp, err := scrp.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response models.IPAbuseResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func (scrp *IpScraper) ExecuteIPSearch(ctx context.Context, ip string) (models.IPSearchResults, error) {

	var failures []string
	var failCount int
	var confirmation bool = true
	var differences []string
	var isp string

	var location models.Coordinates
	var latitude float64
	var longitude float64

	ipApiResult, err := scrp.IpApiCheck(ctx, ip)
	if err != nil || ipApiResult == nil {
		failCount++
		failures = append(failures,
			fmt.Sprintf("IP-API lookup failed: %v", err))
	}

	ipWhoIsResult, err := scrp.IpWhoIsCheck(ctx, ip)
	if err != nil || ipWhoIsResult == nil {
		failCount++
		failures = append(failures,
			fmt.Sprintf("IPWhoIs lookup failed: %v", err))
	}

	var ipAbuseResult *models.IPAbuseReport

	if scrp.abuseApiKey != "" {
		ipAbuseResult, err = scrp.IpAbuseCheck(ctx, ip)
		if err != nil || ipAbuseResult == nil {
			failCount++
			failures = append(failures,
				fmt.Sprintf("AbuseIPDB lookup failed: %v", err))
		}
	} else {
		fmt.Println("AbuseIPDB skipped: API key not configured")
	}

	if failCount > 1 {
		return models.IPSearchResults{
			Summary: models.IPModel{
				IP:          ip,
				IsValidated: false,
				Warning:     differences,
				Errors:      failures,
			},
			Raw: models.RawIPResults{
				IPAPI:   ipApiResult,
				IPWhoIs: ipWhoIsResult,
				AbuseIP: ipAbuseResult,
			},
		}, fmt.Errorf("IP search failed: too many API lookups failed!")
	}

	if ipApiResult != nil && ipWhoIsResult != nil && ipApiResult.ISP != "" && ipWhoIsResult.ISP != "" {
		if ipApiResult.ISP == ipWhoIsResult.ISP {
			isp = ipApiResult.ISP
		} else {
			isp = "multiple"
			differences = append(differences,
				fmt.Sprintf(
					"ISP mismatch found: IP-API=%s, IPWhoIs=%s",
					ipApiResult.ISP,
					ipWhoIsResult.ISP,
				),
			)
		}
	} else if ipApiResult != nil && ipApiResult.ISP != "" {
		isp = ipApiResult.ISP
	} else if ipWhoIsResult != nil && ipWhoIsResult.ISP != "" {
		isp = ipWhoIsResult.ISP
	}

	//threshold of 55.6 km
	const allowedDelta = 0.5

	if ipApiResult != nil && ipWhoIsResult != nil && ipApiResult.Lat != 0.0 && ipWhoIsResult.Lat != 0.0 {
		if math.Abs(ipApiResult.Lat-ipWhoIsResult.Lat) <= allowedDelta {
			latitude = ipApiResult.Lat
		} else {
			latitude = 0.0
			differences = append(differences,
				fmt.Sprintf(
					"Latitude mismatch found: IP-API=%f, IPWhoIs=%f",
					ipApiResult.Lat,
					ipWhoIsResult.Lat,
				),
			)
		}
	} else if ipApiResult != nil && ipApiResult.Lat != 0.0 {
		latitude = ipApiResult.Lat
	} else if ipWhoIsResult != nil && ipWhoIsResult.Lat != 0.0 {
		latitude = ipWhoIsResult.Lat
	}

	if ipApiResult != nil && ipWhoIsResult != nil && ipApiResult.Lon != 0.0 && ipWhoIsResult.Lon != 0.0 {
		if math.Abs(ipApiResult.Lon-ipWhoIsResult.Lon) <= allowedDelta {
			longitude = ipApiResult.Lon
		} else {
			longitude = 0.0
			differences = append(differences,
				fmt.Sprintf(
					"Longitude mismatch found: IP-API=%f, IPWhoIs=%f",
					ipApiResult.Lon,
					ipWhoIsResult.Lon,
				),
			)
		}
	} else if ipApiResult != nil && ipApiResult.Lon != 0.0 {
		longitude = ipApiResult.Lon
	} else if ipWhoIsResult != nil && ipWhoIsResult.Lon != 0.0 {
		longitude = ipWhoIsResult.Lon
	}

	location = models.Coordinates{
		Latitude:  latitude,
		Longitude: longitude,
	}

	if len(failures) > 0 || len(differences) > 0 || failCount > 0 {
		confirmation = false
	}

	var acs *int
	var tr *int
	lra := "Not searched"

	if ipAbuseResult != nil {
		acs = ipAbuseResult.AbuseConfidenceScore
		tr = ipAbuseResult.TotalReports
		lra = ipAbuseResult.LastReportedAt
	}

	return models.IPSearchResults{
		Summary: models.IPModel{
			IP:                   ip,
			ISP:                  isp,
			Location:             location,
			IsValidated:          confirmation,
			Warning:              differences,
			AbuseConfidenceScore: acs,
			TotalReports:         tr,
			LastReportedAt:       lra,
			Errors:               failures,
		},
		Raw: models.RawIPResults{
			IPAPI:   ipApiResult,
			IPWhoIs: ipWhoIsResult,
			AbuseIP: ipAbuseResult,
		},
	}, nil

}
