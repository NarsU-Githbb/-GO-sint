# Data Gathering Information

This OSINT tool has 3 different search methods:

## **Full Name:** 

Method demonstrates how personal information such as names, residential addresses and phone numbers could be gathered / structured in an OSINT pipeline. For educational and legal reasons, it does not query external APIs or real databases. Instead, it uses a locally stored JSON file (`Fullname_mockAPI.json`) that acts as a simulated database containing predefined mock profiles. Public API is not used, because it's illegal to scrape, aggregate or programmatically query real people's phone numbers and addresses without their consent (under EU's GDPR / Finnish privacy laws). So no real individuals are searched, tracked or queried here (...and Fullname_mockAPI.json isn't a collection of real personal information ether just some made up stuff).

When a user inputs a full name, the application loads & decodes the local JSON dataset using Go's `encoding/json` package. It then normalizes the search input by removing leading / trailing whitespaces and converts it to lowercase before comparing it against the stored search queries. If a matching profile is found, the application extracts the corresponding name, address & phone number. If no match is found, it returns a result indicating that the requested information is unavailable.

## **IP Address:** 

Method retrieves geographical and network-related information about an IP address and checks for potentially malicious activity using three external APIs: IP-API, IPWhoIs & AbuseIPDB.

The application sends HTTP GET requests containing the IP address input to each API endpoint. The returned JSON responses are decoded using Go's `encoding/json` package into predefined structures within the application's models package. IP-API and IPWhoIs provide geographical coordinates (latitude & longitude) and Internet Service Provider (ISP) information, while AbuseIPDB provides threat intelligence data, including an abuse confidence score, total reports & the date of the last reported activity. AbuseIPDB requires an API key, which is loaded from the application's environment variables (`.env`).
In order to require the API key, you need to register in AbuseIPDB (here's a link to free registeration: [AbuseIPDB Free Register](https://www.abuseipdb.com/register?plan=free) ).

The application primarily cross-references the results from IP-API and IPWhoIs to identify inconsistencies in the returned ISP and geographical information. ISP names are compared between the two providers & geographical coordinates are checked against an allowed difference of 0.5 degrees (threshold of 55.6 km). If the returned information differs beyond the defined threshold, the application records a warning and marks the corresponding information as inconsistent.

AbuseIPDB operates independently of this geographical comparison and supplements the results with information about previously reported malicious activity associated with the queried IP address.

## **Social Media Username:** 

Method checks for the user's presence on 5 social media platforms - including X, Youtube, Soundcloud, Github and Gitea to determine if accounts exists on the platform. This one uses basic go http.requests (`net/http`) to check if the username can be found on the platform according to http status codes. An HTTP 200 (OK) response is treated as an indication that a profile may exist, while an HTTP 404 (Not Found) response indicates that the requested profile could not be found. Other HTTP status codes, such as 403 (Forbidden), 429 (Too Many Requests), and server errors, are treated as unexpected responses and recorded as errors rather than as evidence that an account exists or does not exist.

It's important to know that the HTTP status codes alone cannot reliably confirm whether a username belongs to an existing account. Platforms may redirect requests, restrict access, block automated requests, apply rate limits or return different responses depending on their own policies and infrastructure. The social media platforms chosen above I've personally tested with multiple different usernames & my tests have so far brought correct information (they could however change their approach to the statuscodes at some point - you never know). So the results should be interpreted as more of a indicator of possible account existing in the platform rather than definitive confirmation.

## 🛠️ Libraries and Methods

Instead of using heavy web scraping frameworks (like Playwright or Scrapy), this tool relies almost entirely on the **Go Standard Library** alongside one official sub-repository module.

### Core Architecture Components:

* **Native Network Clients (`net/http` & `context`):** HTTP connections utilize the built-in Go HTTP client configured with hard 5-second context timeouts to prevent slow servers or dead routes from hanging the program execution.
* **Token-Bucket Throttling (`golang.org/x/time/rate`):** Rather than standard library time delays, this external Go package controls precise event flow-rates to safely query target infrastructures without breaching free-tier API parameters.
* **Graceful Shutdowns (`os/signal` & `syscall`):** Implements system signal monitoring to catch exit intents (such as `Ctrl+C`), enabling the script to cleanly close active routines and release network resources safely.
* **CLI-Driven Configuration (`flag`):** The program handles execution variables natively via terminal arguments, providing a lightweight command-line user experience typical for OSINT utility scripts.

## 🔄 API Integrations & Data Validation

* **Integrated APIs:** `ip-api.com` (free tier), `IPWhois` and `AbuseIPDB`.
* **Data Cross-Referencing:** different geolocation services maintain their own proprietary BGP and IP databases, information can occasionally mismatch. 

## ⏱️ Rate-Limiting Approach

To prevent our tool from getting IP-banned or accidentally overloading third-party servers, we use strict request throttling. 

Instead of spamming endpoints, the tool implements token-bucket rate limiters using the `golang.org/x/time/rate` package. For instance, since `ip-api.com` restricts free tiers to 45 requests per minute, our `IpScraper` enforces a safe cooldown (e.g., 1 request every 2 seconds) for each target API (`ip-api`, `IPWhois`, and `AbuseIPDB`). This keeps the tool running smoothly without hitting `HTTP 429 Too Many Requests` limits.


