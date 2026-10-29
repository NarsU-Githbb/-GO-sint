# Data Gathering Information

This OSINT tool has 3 different search methods:

* **Full Name:** Method gathers the person's address and phone number, along with any additional personal information available. This method uses made up mock API & is meant for educational purposes only. Public API is not used, because it's illegal to scrape, aggregate or programmatically query real people's phone numbers and addresses without their consent (under EU's GDPR / Finnish privacy laws).

* **IP Address:** Method determine the location and Internet Service Provider (ISP) among some information on the possibility if the IP has been abused. This method integrates ip-api, ipwhois & abuseipdb APIs & cross-referances so that the information can be validated. Mostly the cross-referencing happens between ip-api & ipwhois APIs, while Abuseipdb focuses on the information if IP has been compromised. 

* **Social Media Username:** Method checks for the user's presence on 5 social media platforms - including X, Youtube, Soundcloud, Github and Gitea to determine if accounts exists on the platform. This one uses basic http.requests to check if the username can be found on the platform according to http status codes. 

## 🛠️ Libraries and Methods

Instead of using heavy web scraping frameworks (like Playwright or Scrapy), this tool relies almost entirely on the **Go Standard Library** alongside one official sub-repository module.

### Core Architecture Components:

* **Native Network Clients (`net/http` & `context`):** HTTP connections utilize the built-in Go HTTP client configured with hard 5-second context timeouts to prevent slow servers or dead routes from hanging the program execution.
* **Token-Bucket Throttling (`golang.org/x/time/rate`):** Rather than standard library time delays, this external Go package controls precise event flow-rates to safely query target infrastructures without breaching free-tier API parameters.
* **Graceful Shutdowns (`os/signal` & `syscall`):** Implements system signal monitoring to catch exit intents (such as `Ctrl+C`), enabling the script to cleanly close active routines and release network resources safely.
* **CLI-Driven Configuration (`flag`):** The program handles execution variables natively via terminal arguments, providing a lightweight command-line user experience typical for OSINT utility scripts.

## 🔄 API Integrations & Data Validation

* **Integrated APIs:** `ip-api.com` (free tier), `IPWhois`, and `AbuseIPDB`.
* **Data Cross-Referencing:** different geolocation services maintain their own proprietary BGP and IP databases, information can occasionally mismatch. 

## ⏱️ Rate-Limiting Approach

To prevent our tool from getting IP-banned or accidentally overloading third-party servers, we use strict request throttling. 

Instead of spamming endpoints, the tool implements token-bucket rate limiters using the `golang.org/x/time/rate` package. For instance, since `ip-api.com` restricts free tiers to 45 requests per minute, our `IpScraper` enforces a safe cooldown (e.g., 1 request every 2 seconds) for each target API (`ip-api`, `IPWhois`, and `AbuseIPDB`). This keeps the tool running smoothly without hitting `HTTP 429 Too Many Requests` limits.


