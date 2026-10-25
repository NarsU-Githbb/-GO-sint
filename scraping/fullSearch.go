package scraping

import (
	"context"
	"digital_detective/models"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

/**
 * ============================================================================
 * COMPLIANCE & LEGAL NOTICE: STUDENT DEMO ONLY
 * ============================================================================
 *
 * THIS IS A NON-FUNCTIONAL PROTOTYPE:
 *
 * It does not actually search the internet,
 * query any real databases or save any real data. All results here are faked or blocked.
 *
 * REGULATORY REASONING:
 *
 * Under the EU General Data Protection Regulation (GDPR) and the Finnish Data
 * Protection Act (Tietosuojalaki), phone numbers and residential addresses
 * associated with identifiable individuals are classified as personal data.
 *
 * Automated scraping, aggregation or programmatic querying of this data
 * without explicit data subject consent or a valid legal basis constitutes
 * illegal data processing (GDPR Article 6). Furthermore, it systematically
 * bypasses mandatory statutory data disclosure bans managed by the Finnish
 * Digital and Population Data Services Agency (DVV) and Traficom.
 *
 * INTENDED USE:
 *
 * This functionality is purely an educational example for my coding assignment.
 * It safely demonstrates how the code logic *would* look like, using mock data.
 * No real people are tracked, searched, or saved.
 *
 * FOR TESTING:
 *
 * The only names in the fake API are "john Doe" & "Dwight Schrute". Use those names to
 * quickly test out the functionality of the fullname search. You can also create a new
 * item using a template inside FullnameJSON_template.md inside Fullname_mockAPI.json
 *
 * ============================================================================
 */

type FnScraper struct {
	mockLimiter  *rate.Limiter
	mockDataPath string
}

func NewFnScraper() *FnScraper {
	jsonPath := "fullname_mockAPI.json"
	return &FnScraper{
		mockLimiter:  rate.NewLimiter(rate.Every(1*time.Second), 1),
		mockDataPath: jsonPath,
	}
}

// loadMockProfiles reads our local fake JSON file safely

func (scrp *FnScraper) loadMockProfiles() ([]models.MockProfile, error) {
	api, err := os.Open(scrp.mockDataPath)
	if err != nil {
		return nil, err
	}

	defer api.Close()

	decoder := json.NewDecoder(api)
	var profiles []models.MockProfile

	err = decoder.Decode(&profiles)
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func (scrp *FnScraper) SimulatedFnCheck(ctx context.Context, fullName string) (*models.FullModel, error) {

	if err := scrp.mockLimiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limiter error: %w", err)
	}

	profiles, err := scrp.loadMockProfiles()
	if err != nil {
		return nil, fmt.Errorf("failed to read mock database: %w", err)
	}

	// Create lowercase input query
	LowcaseInputQuery := strings.ToLower(strings.TrimSpace(fullName))

	for _, profile := range profiles {

		// Match against lowercase api 'search_query' field

		if strings.ToLower(profile.SearchQuery) == LowcaseInputQuery {

			var Phone string
			var Location models.MockAddress

			for _, source := range profile.ScrapedSources {

				if source.PhoneNumber != "" {
					Phone = source.PhoneNumber
				}

				if source.Address.Street != "" {
					Location = source.Address
				}
			}
			names := strings.Fields(profile.Person.FullName)
			firstName := ""
			lastName := ""

			if len(names) > 0 {
				firstName = names[0]
			}
			if len(names) > 1 {
				lastName = names[len(names)-1]
			}

			return &models.FullModel{
				FirstName:   firstName,
				LastName:    lastName,
				Address:     Location,
				PhoneNumber: Phone,
			}, nil
		}
	}
	return nil, nil
}

func (scrp *FnScraper) ExecuteFullSearch(ctx context.Context, fullName string) (models.FullModel, error) {

	profilePointer, err := scrp.SimulatedFnCheck(ctx, fullName)
	if err != nil || profilePointer == nil {
		names := strings.Fields(fullName)

		firstName := ""
		lastName := ""

		if len(names) > 0 {
			firstName = names[0]
		}

		if len(names) > 1 {
			lastName = names[len(names)-1]
		}

		// Return the fallback model with "not found!" placeholders
		return models.FullModel{
			FirstName: firstName,
			LastName:  lastName,
			Address: models.MockAddress{
				Street:     "not found!",
				PostalCode: "not found!",
				City:       "not found!",
			},
			PhoneNumber: "not found!",
		}, nil
	}

	result := *profilePointer
	return result, nil
}
