# (G)Osint CLI Tool

A Lightweight command-line intelligence tool written in Go, designed to demonstrate automated open-source data aggregation techniques. Built strictly for educational purposes, this application showcases how to safely perform username profile lookups, query public IP geolocation endpoints and cross-reference threat intelligence data using Go Standard library rather than resource-heavy scraping frameworks.

## Project Overview

Before you get started, you should read this documentation to understand the project & the concepts needed to grasp in order to understand "why osint is useful technique for cybersecurity"? Also reading the documentation includes answers to some of the testing questions inside kood/Sisu platform & has a detailed knowledge about the features inside the application / how they truly work.

**Concepts:**

* [What is OSINT?](docs/osint.md)
* [Digital Footprint](docs/digital_footprint.md)
* [CIA Triad](docs/cia_triad.md.md)

**Features:**

* [Data Gathering Documentation](docs/data-gathering.md)

**For kood/Sisu Reviewers:**

* [Information about the extra features](docs/testing.md)

## Prerequisites

Depends on your preferred installation method:

**Docker installation:**

- Docker installed and running.
- No separate Go installation required on the host machine.

**Manual installation:**

- Go version 1.26.0 (or later) installed on your system.
- Git (for cloning the repository).

## Setup Configuration

Clone the repository, navigate into the directory & set up your local environment file:

```bash
cp .env.example .env
```

*Open the newly created `.env` file & append your private AbuseIPDB APIKey if you want to implement it. It's an optional feature though, so you can leave .env as it is & the program will work as intended (it just skips the optional AbuseIPDB check).*

## Usage Guide

Run the core script using standard Go invocation combined with one of the search mode flags:

```bash
go run . [flag] <input>
```

### Available Flags:

* `-n` : Performs a simulated full-name query (Mock PII data).
* `-ip` : Performs a live cross-referenced IP lookup.
* `-un` : Performs a multi-platform username audit.
* `-raw` : Returns raw API responses if possible (add it before another flag).

## How to Install and Run with Docker (Recommended Method)

The project is a Go-based command-line application and does not require a graphical user interface or frontend. Docker simplifies deployment by automatically building and packaging the Go application into a lightweight runtime image. No separate Go installation or manual dependency compilation is required on the host machine.

### The repository includes:

* **Dockerfile** - Builds and packages the application.
* **docker-run.sh** - Builds and executes the application using a single command.
* **.dockerignore** - Optimizes the Docker build context.

### Execution:

```bash
./docker-run.sh --help
```

### Examples:

```bash
./docker-run.sh -n "Dwight Schrute"
```

```bash
./docker-run.sh -ip 8.8.8.8
```

```bash
./docker-run.sh -un torvalds
```

If you want quickly to test every available CLI command easily - just run this command:

```bash
./docker-run.sh -all
```

### Execution Output Examples

1. **Full-Name Search Example (Mock Environment)**

```bash
go run . -n "John Doe"
```

**Output Preview:**

```bash
First Name: John
Last Name: Doe
Address: Mango street 12 A 4, 00100 Helsinki
Phone Number: +358 40 123 4567
Result written to file: `n_dwight_schrute.txt`
```

2. **IP Address Lookup Example**

```bash
go run . -ip 8.8.8.1
```

**Output Preview:**

```bash
IP: 8.8.8.1
ISP: Google LLC
Location: 37.422000,-122.085000
Validated: true
Warnings: None
Abuse Confidence Score: Not available
Total Reports: Not available
Last Reported At: Not searched
Result written to file: ip_8_8_8_1.txt
```

3. **Social Media Username Check Example**

```bash
go run . -un torvalds
```


```bash
X: true
Youtube: false
Soundcloud: true
Github: true
Gitea: true
Result written to file: un_torvalds.txt
```

4. **Raw Responses Examples**

```bash
go run . -raw -n "Dwight Schrute" 
```

or 

```bash
go run . -raw -ip 1.1.1.1
```

Username search doesn't have `-raw` output as it doesn't utilize any APIs.

## 🔐 Security & API Configuration

The application securely interfaces with external data sources without exposing proprietary credentials inside the version control history:

* **Environment Separation:** Sensitive configuration parameters and private API keys (such as your AbuseIPDB token) are kept completely isolated from the codebase using a local `.env` configuration file.
* **Zero-Exposure Commits:** The root `.gitignore` ruleset explicitly blocks environment files and local configuration templates from being pushed to public repositories, preventing accidental supply-chain leaks.
* **On-Launch Bootstrapping:** System environment variables are pulled programmatically at runtime, ensuring your real API keys are only loaded directly into the ephemeral machine memory during execution.

## ⚖️ Ethical & Legal Considerations

As an educational OSINT tool, keeping our scraping ethical is a priority:

* **Publicly Accessible Data Sources:** The tool retrieves information from publicly accessible OSINT services and external API endpoints. Some services, such as AbuseIPDB, require API-key authentication to access their threat intelligence APIs. The application does not access private databases or unauthorized information sources.
* **Target-Specific Limits:** Our rate limiters are designed to respect the public terms of service of each API we target so we don't abuse their free infrastructure.
* **Privacy-by-Design (Mock Data):** Because PII (Personally Identifiable Information) like real addresses and phone numbers fall under strict EU GDPR / Finnish privacy laws, the Full Name search method relies on completely simulated mock data. This demonstrates how a real OSINT pipeline functions without illegally scraping or aggregating live human identities.
