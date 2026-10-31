## Quality

**The student has implemented data visualization for the gathered data.**

* Yes, normal outputs has some ASCII art attached to it. 

**The chosen visualization format makes the data easier to analyze compared to raw text output.**

* ¯\ _(ツ) _/¯ 

**The enhanced data representation is consistent across different types of searches.**

* Every normal search output uses ASCII art, but the information inside those outputs may differ depending what they return.

**The student has integrated additional APIs beyond the minimum requirements.**

* The program uses 3 APIs (minimum was 2 APIs).

**The additional APIs provide more comprehensive information.**

* The additional API used in IP Search (AbuseIPDB) brings completely different information
than other APIs, which checks if the IP has been used in malicious activities before. 

**The tool can perform searches using multiple parameters in a single command.**

* No, but there is a test feature that runs all the commands in sequence using a script.
This though only uses the same example inputs everytime. 

**The output file is named correctly based on the parameters used and the full name.**

* yes!

**Student has implemented additional technologies, security enhancements and/or features beyond the core requirements.**

* Not much, but there is a -raw flag that returns raw api responses instead of the output  summary. This can be used in full name searches & ip searches (not username search as it
doesn't use any APIs). Just add it before the other flag like "-raw -n 'input'" or 
"-raw -ip 'input'" & it should print you the raw response. 

**Project runs with a single Docker command or script**

* yes, just use "./docker-run.sh -ip 8.8.8.8" for example (or whatever search you want to
make)
