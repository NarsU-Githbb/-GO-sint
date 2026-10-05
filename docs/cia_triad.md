# CIA TRIAD

The CIA Triad is a simple checklist used to protect your digital information. To have a secure system - you must satisfy all 3 core concepts: **Confidentiality**, **Integrity** & **Availability**. If even one part is missing, your security will be compromised in the future.

# CIA Triad - 3 Core Parts:

## 1. 🔒  Confidentiality  🔒  (Keep it Secret): 

Ensuring only authorized people can access specific data. It prevents data leakage using tools like passwords, encryption & multi-factor authentication (MFA).

### **How can it be compromised?**

There are several ways confidentiality can be compromised. This may involve direct attacks aimed at gaining access to systems the attacker does not have the rights to see. It can also involve an attacker making a direct attempt to infiltrate an application or database so they can take data or alter it. 

One example of this would be *man-in-the-middle (MITM)* attack, where an attacker positions themselves in the stream of information to intercept data and then either steal or alter it.

Not all violations of confidentiality are intentional as human error or insufficient security controls may be to blame as well. For example, An employee could accidentally leave their computer unlocked or organization fails to enforce an automatic screen-lock timeout policy. Also, a thief may steal hardware, whether an entire computer or a device used in the login process and use it to access confidential information.

### **Ways to secure?**

• Using strong passwords & Multi-Factor Authentication (MFA).

• Encrypt your data.

• Train staff never to share login details / being extra careful with company owned hardware that it doesn't slip into criminals posession.

## 2. 🔄  Integrity  🔄  (Keep it Correct): 

Ensuring data remains accurate, complete & untampered with. It stops data from being secretly changed by using safeguards like digital signatures & version tracking.

### **How can it be compromised?**

Let's say your company provides information about senior managers on your website, this information needs to be truthful. If it's inaccurate, those visiting your website for information may feel your organization is not trustworthy. If someone is on a mission of damaging the reputation of your organization they may try to hack your website & alter the descriptions, photographs or titles of the executives to hurt their / your whole organizations reputation.

An attacker may bypass an *intrusion detection system (IDS)* (a passive cybersecurity tool designed to monitor network traffic or device activity for signs of malicious behavior, policy violations or known threat patterns) to change file configurations to allow unauthorized access & alter the logs to hide the fact that they broke in.

Integrity may also be violated by accident. Someone may enter wrong code or make another kind of careless mistake. Also, if the company’s security policies, protections, and procedures are inadequate, integrity can be violated without any one person in the organization accountable for the blame.

### **Ways to secure?**

• Use digital signatures to prove exactly who sent a file via email for example.

• Employ trustworthy certificate authorities (CAs) that can verify the authenticity of your website.

• Use hashing or encyption (a digital fingerprint for files that changes completely if even a single letter is modified).

• Set up version tracking to see who changed what & when.


## 3. ⚡  Availability  ⚡  (Keep it Accessible): 

Ensuring systems / data are reliable & ready whenever authorized users need them. This is maintained through routine backups, system updates & disaster recovery plans to minimize downtime.

### **How can it be compromised?**

Suppose there's a natural disaster, which would cause a power outage & there is no *disaster recovery system* in place to help users regain access to critical systems, availability will be compromised. That compromise could prevent users from getting to the office, which can interrupt the availability of their workspace / other devices that provide business-critical information or applications.

Availability can also be compromised through deliberate acts of sabotage, such as the use of *denial-of-service (DoS) attacks*, which flood networks or servers with fake traffic to crash systems or worst-case scenario, someone could install *ransomware* to your environment that will encrypt critical data & hold it hostage until a ransom is paid.

### **Ways to secure?**

• Create full disaster recovery plan that helps a company regain availability soon after a negative event has happened.

• Use redundancy (having backup servers ready to take over instantly if the main one goes down).

• Keep systems updated so they don't crash from old bugs.

• Create backups of your data & keep them in a safe, separate location.

## The Balance Dilemma 

In the real world, you cannot easily have 100% of all three. Increasing one often hurts another so most cybersecurity professionals battle with Security vs. Usability balancing.

• **Too much Confidentiality hurts Availability:** If you require a 20-character password, a fingerprint scan & a text message code just to open a file - it becomes slow & difficult for employees to do their jobs.

• **Too much Availability hurts Confidentiality:** If you make data incredibly easy to access from any device anywhere in the world - you increase the risk of an unauthorized person stealing it.

## Basically, use CIA TRIAD when...

• ...deciding who gets access to what data / files.

• ...looking for weak points in a company's computer network.

• ...trying to figure out what hackers want so you can protect it first.