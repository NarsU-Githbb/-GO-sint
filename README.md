# (g)OSINT CLI Tool

A Lightweight command-line intelligence tool written in Go, designed to demonstrate automated open-source data aggregation techniques. Built strictly for educational purposes, this application showcases how to safely perform username profile lookups, query public IP geolocation endpoints and cross-reference threat intelligence data using Go Standard library rather than resource-heavy scraping frameworks.

## Documentation

Before you get started, you should read this documentation to understand the project & the concepts needed to grasp in order to understand why osint is useful technique for cybersecurity? (also they include answers to the testing questions inside kood/Sisu platform)

[What is OSINT?](docs/osint.md).
[Digital Footprint](docs/digital_footprint.md)
[CIA Triad](docs/cia_triad.md.md).
[Data Gathering Documentation](docs/data-gathering.md).

## 🚀 How to Install and Run

### 1. Prerequisites

Make sure you have [Go](https://go.dev) installed on your system.

### 2. Setup Configuration

Clone the repository, navigate into the directory, and set up your local environment file:

```bash
cp .env.example .env
```

*(Open the newly created `.env` file and append your private AbuseIPDB or other API tokens where requested).*

### 3. Usage Syntax

Run the core script using standard Go invocation combined with one of the search mode flags:

```bash
go run . [options] <input>
```

#### Available Flags:

* `-n` : Performs a simulated full-name query (Mock PII data).
* `-ip` : Performs a live cross-referenced IP lookup.
* `-un` : Performs a multi-platform username audit.

### 💡 Execution Examples

#### 1. Full-Name Search (Mock Environment)
```bash
go run . -n "Dwight Schrute"
```
**Output Preview:**
* **First Name:** Dwight
* **Last Name:** Schrute
* **Address:** Mango street 12 A 4, 00100 Helsinki
* **Phone Number:** +358 40 123 4567
* **Result written to file:** `n_dwight_schrute.txt`

#### 2. IP Address Lookup
```bash
go run . -ip 1.1.1.1
```
**Output Preview:**
* **ISP:** Cloudflare, Inc.
* **Location:** -33.4940, 143.2104
* **Result written to file:** `ip_1_1_1_1.txt`

#### 3. Social Media Username Probing
```bash
go run . -un @recyclops
```
**Output Preview:**
* **Second Life:** yes
* **X:** no
* **LinkedIn:** no
* **Instagram:** no
* **GitHub:** no
* **Result written to file:** `un_recyclops.txt`

## 🔐 Security & API Configuration

The application securely interfaces with external data sources without exposing proprietary credentials inside the version control history:

* **Environment Separation:** Sensitive configuration parameters and private API keys (such as your AbuseIPDB token) are kept completely isolated from the codebase using a local `.env` configuration file.
* **Zero-Exposure Commits:** The root `.gitignore` ruleset explicitly blocks environment files and local configuration templates from being pushed to public repositories, preventing accidental supply-chain leaks.
* **On-Launch Bootstrapping:** System environment variables are pulled programmatically at runtime, ensuring your real API keys are only loaded directly into the ephemeral machine memory during execution.

## ⚖️ Ethical & Legal Considerations

As an educational OSINT tool, keeping our scraping ethical is a priority:

* **Public Data Only:** The tool exclusively queries publicly available OSINT data and endpoint services, avoiding any authenticated or private databases.
* **Target-Specific Limits:** Our rate limiters are designed to respect the public terms of service of each API we target so we don't abuse their free infrastructure.
* **Privacy-by-Design (Mock Data):** Because PII (Personally Identifiable Information) like real addresses and phone numbers fall under strict EU GDPR and Finnish privacy laws, the Full Name search method relies on completely simulated mock data. This demonstrates how a real OSINT pipeline functions without illegally scraping or aggregating live human identities.
