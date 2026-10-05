# OSINT — Open Source Intelligence

## What Is OSINT?

OSINT (Open Source Intelligence) is the practice of gathering public, open-source information from **lawfully accessible public sources** to answer a specific question. It is used in many different fields by various professionals:

- **Journalism & NGOs:** for investigation & fact-checking
- **Human Resources:** for background checks on potential candidates
- **Financial institutions:** as part of the due diligence process, before entering a new contract
- **Insurances:** to assess potential customers & investigate claims
- **Law enforcement & private detectives:** to find new evidence, connect criminal activity to a person of interest, or investigate someone
- **Cybersecurity companies / consultants:** to perform penetration testing on companies or investigate a malicious actor
- **Cybercriminals or threat actors:** to conduct initial reconnaissance on a target company or individual, identifying weak points like exposed employee credentials, vulnerable software versions or personal details used to craft highly convincing phishing & social engineering attacks.

## ⚠️ Clearing Up the Confusion: What OSINT is NOT

Because OSINT deals with data collection, it is frequently confused with other technical fields. To keep your investigations accurate & legally sound, it is vital to understand how these terms differ based on **data accessibility, legality & methods**.

### 🔍 1. Digital Forensics vs. OSINT

* **Digital Forensics** is the scientific recovery, preservation & analysis of **private, hidden or deleted data** directly from physical storage media (like hard drives, locked phones or server logs). It is heavily regulated and typically requires a legal warrant or explicit corporate authorization.
* **The Difference:** OSINT only looks at **publicly broadcasted information** on the surface of the internet. Digital Forensics dives into **private internal storage**, decrypting files and pulling data that the owner actively tried to hide or delete from their physical device.

### 🕸️ 2. Data Scraping vs. OSINT
* **Data Scraping** (or Web Scraping) is an automated programming technique used to extract massive amounts of raw data from websites quickly and save it into a structured format (like a spreadsheet or database). 
* **The Difference:** Data Scraping is merely a **mechanism or tool** used to gather raw building blocks. OSINT is the **analytical framework** that turns that raw data into meaningful intelligence. Furthermore, scraping can sometimes violate a website's Terms of Service (ToS) or cross into private API boundaries, whereas OSINT strictly adheres to analyzing public-facing information.

### 🕵️ 3. Cyber Reconnaissance (Recon) vs. OSINT
* **Cyber Reconnaissance** is the initial phase of a cyberattack or penetration test. It involves actively probing a specific target's network infrastructure to discover open ports, live IP addresses, server vulnerabilities & firewall weaknesses.
* **The Difference:** While Recon utilizes OSINT (like looking up domain registration records), it frequently crosses over into **Active Engagement**. Sending malicious packets to scan a network's defenses is active hacking; OSINT is purely **Passive**, meaning the investigator reads public records without ever directly interacting with or alerting the target infrastructure.

### 📊 Summary Comparison

| Concept | Data Type | Legality / Boundary | Core Purpose |
| :--- | :--- | :--- | :--- |
| **OSINT** | Publicly Available | 100% Legal & Passive | Finding context, patterns and threat intelligence. |
| **Digital Forensics** | Private & Internal | Requires Warrant / Ownership | Proving digital chain of custody in legal courts. |
| **Data Scraping** | Structured Public/Semi-Public | Subject to Website ToS | Automating bulk data collection into spreadsheets. |
| **Cyber Recon** | Infrastructure & Systems | Can border on unauthorized scanning | Mapping technical network vulnerabilities for entry. |

## Major OSINT Categories

### 1. Search Engine & Web OSINT

Start with the basics: search engines and the target's publicly available website.

**Useful information can include:**

- Company information and subsidiaries

- Employees and organizational structure

- News articles and public mentions

- Contact information

- Technology and infrastructure clues

- Public documents

- Historical versions of websites

- Domain registration information

**Useful Sources & Techniques:**

- Search engines: Google, Bing, DuckDuckGo

- Domain registration: ICANN Lookup

- Web archives: Internet Archive / Wayback Machine

- DNS & infrastructure research: ViewDNS and similar services

**Security Relevance**

Public web information can reveal an organization's technology, employees, historical infrastructure or outdated information. Attackers may use these clues to improve phishing, impersonation or reconnaissance.

### 2. Username, Email & SOCMINT

### Username Intelligence

A username can sometimes be correlated across multiple platforms.

**Useful techniques include:**

- Searching the username directly

- Checking username variations

- Comparing profile pictures

- Comparing biographies and interests

- Looking for consistent links or websites

**Useful Sources & Techniques:**

- Tools: Namechk, WhatsMyName, Sherlock

A username match alone is **not proof of identity**. Treat it as a lead that requires more investigation.

### Email Intelligence

An email address is a unique global identifier. Because no two people can own the exact same email address, finding one allows an analyst to pivot and map out a target's digital footprint with high accuracy.

**Information that can be found:**

- Account Registrations (on platforms like LinkedIn, Amazon, Adobe etc).

- Corporate Structure & Domains (professional email like john.doe@company.com for example - explicitly confirms a target's employer, their naming convention format & their corporate affiliation).

- Searching an email can uncover associated phone numbers, physical addresses or the target's real name (if they are not using it in email address).

⚠️ Important: PDF, Word or Excel documents found online often contain hidden metadata identifying the author's email address.

**Useful Sources & Techniques:**

- Breach & Leak Checkers: HaveIBeenPwned

- Email Verification Tools: Epieos, Holehe

- OSINT Search Engines: IntelX, Dehashed

### SOCMINT

**SOCMINT (Social Media Intelligence)** is the collection and analysis of information from publicly accessible social-media platforms.

Public social-media activity can reveal:

- Employment and professional relationships

- Interests and hobbies

- Events and locations

- Organizational structure

- Relationships between accounts

- Historical activity

**Security Risks**

Attackers can use public social information for:

- **Spear-phishing:** creating highly personalized messages

- **Social engineering:** exploiting known interests or relationships

- **Impersonation:** identifying people worth impersonating (like some of your close friends)

- **Business Email Compromise:** researching organizational roles and reporting structures

- **Account recovery attacks:** discovering information that may have been used in security questions

**Useful Sources & Techniques:**

- For account discovery & correlation: Namechk, WhatsMyName, Sherlock
- For archive & historical tracking: Wayback Machine (Internet Archive)
- For platform-specific scraping & analysis: Twint (for X/Twitter), Maltego, Social Links

⚠️ Important: Public information should be treated as evidence, not automatically as truth. A matching username, photo or name can belong to multiple people.

### 3. IMINT

IMINT (Imagery Intelligence) involves extracting information from photographs, satellite imagery, videos & other visual material.

**Common techniques include:**

- Metadata analysis

- Reverse image search

- Geolocation

- Image comparison

- Visual identification of landmarks

- Time and environmental analysis

### EXIF Metadata

EXIF (Exchangeable Image File Format) metadata can contain information such as:

- GPS coordinates

- Date and time

- Camera/device model

- Software used to edit the image

- Author/copyright information

Not every image contains all of this information & many platforms remove some metadata during upload (like social media & whatsapp in most cases).

**Security Risks**

EXIF information can potentially expose:

- Where a photograph was taken

- When it was taken

- What device created it

- Device Model and Software Version

- Information connecting the image to its creator

A series of publicly available images can sometimes reveal routines, locations or relationships even when individual images appear harmless. This might give criminal information about your home address & when you are away from it (like on holiday in foreign country), so that it's safe for them to do a burglary on your house (in worst case scenario). Also If they see you are running an outdated, unpatched version of Android for example - they know your device is vulnerable to specific, known security exploits & may target you with tailored malware or phishing links.

**Useful Sources & Techniques:**

- For EXIF data retrieval use: metadata2go, Exiftool, Preview on Mac OS
- For GPS location use: Google Maps

### Reverse Image Search

Reverse image search uses an image (or part of an image) as an search term to find visually similar or previously published images.

**It can help identify:**

- Where an image originally appeared

- Unmarked locations or landmarks in social media

- Other accounts using the same image

- Older versions of an image

- Potentially reused profile pictures

**Useful Sources & Techniques:**

- Reverse image search: Google images, tineye, Yandex
- Image enhancement: MyHeritage, Let’s Enhance, InVID & WeVerify Chrome extension

**Security Risks**

Reverse-image searching can be abused to:

- Connect pseudonymous accounts

- Discover old or forgotten profiles

- Identify locations

- Reuse photographs for impersonation or fraud

Again, a visual match is a lead, not proof of identity.

### 4. Website & Document OSINT

Websites and publicly accessible documents often contain more information than their main pages reveal.

**Look for:**

- PDFs and office documents

- Job advertisements

- Annual reports

- Press releases

- Public presentations

- Technical documentation

- Metadata

- Historical website versions

- Subdomains and related domains

- Publicly referenced technologies

- Organization names and contact details

### Document Metadata

Documents can contain metadata such as:

- Author

- Organization

- Creation/modification dates

- Software used

- Document title

- Revision information

Metadata should be treated as supporting evidence because it can be removed, modified or inaccurate.

**Security Relevance**

Poorly managed public documents may unintentionally reveal organizational information, internal naming conventions, employee names, technologies or historical information.

**Useful Sources & Techniques:**

- tools: ExifTool, PDF metadata viewers, Internet Archive, search engines

### 5. HUMINT & Social Engineering

**HUMINT (Human Intelligence)** is intelligence obtained from people.

In an OSINT context, this can involve analyzing public statements, interviews, forums, communities, comments & other openly available human sources.

**Social engineering** is different: it involves influencing or manipulating people.

OSINT can help an investigator understand:

- Who is connected to whom

- Organizational roles

- Public interests and expertise

- Communication patterns

- Publicly stated opinions or claims

**Security Relevance**

Attackers may combine OSINT with social engineering to make fraudulent communication appear legitimate.

For example, publicly available information about an employee's role, colleagues, projects or business relationships can make an impersonation attempt more convincing.

**Key principle:** collecting public information is different from deceiving, coercing or manipulating someone to obtain information.

## Verification & Corroboration

One of the most important OSINT skills is knowing when to not trust your first findings.

So remember to verify & corroborate. You may ask, what does Corroboration mean? It means the use of extra evidence, information or proof to support & confirm a statement, theory or claim.

A useful process flow in OSINT should be:

1. **Observe** —> What does the source actually show?

2. **Infer** —> What could this information suggest?

3. **Corroborate** —> Can another independent source support it?

4. **Conclude** —> What can reasonably be established?

### Example:

**Observation:**
An account uses the username example123.

**Inference:**
Some other websites have an account using example123. They might belong to the same person?

**Corroboration:**
All accounts independently reference the same public website and email address.

**Conclusion:**
It is highly likely that all accounts belong to the same person.

(Even with this evidence, you must label the conclusion as High Confidence proof, but not 100% fact as you never know if someone has copied or faked a profile.)

### Avoid Confirmation Bias

Do not search only for information that supports your initial theory. Instead ask:

```
"What evidence would prove me wrong?"
```

This is especially important when identifying people or connecting separate accounts.

## OSINT Investigation Workflow

This is common example of the OSINT Workflow: 

```
1. Define your question
          ↓
2. Identify starting information
          ↓
3. Search
          ↓
4. Find useful clues
          ↓
5. Pivot to new identifiers
          ↓
6. Corroborate
          ↓
7. Record your sources
          ↓
8. State what you know — and what you don't know
```

### Good OSINT Practice

- Define your objective first. Don't collect information simply because it exists.

- Stay within your authorization and applicable laws.

- Prefer primary sources when possible.

- Corroborate important claims.

- Record URLs, timestamps & relevant evidence.

- Separate facts from assumptions.

- Consider alternative explanations.

- Minimize unnecessary personal data collection.

- Clearly communicate uncertainty.

### The Core Idea

```
OSINT is not just finding information. It is turning publicly accessible information into reliable, verifiable intelligence.
```

The strongest OSINT investigation is therefore not the one that collects the most data - it is the one that answers the original question with **relevant, corroborated, and defensible evidence.**

