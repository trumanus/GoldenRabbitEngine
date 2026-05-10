# Security policy

## Scope

This repository is a **reference Go library** with **no production deployment** maintained by the authors. Security expectations:

- Do **not** commit secrets, API keys, or personal data.  
- Treat demos (`cmd/gr-demo`, `cmd/gr-full`) as **local laboratory** code.

## Reporting

If you believe you found a security issue **in this repository’s code** (e.g. a pattern that could mislead integrators into unsafe defaults), email **trumanus@gmail.com** with:

- Description and impact  
- Steps to reproduce (if any)  
- Suggested fix (optional)

For vulnerabilities in **your** deployment that uses this library, follow your own incident process.

## Disclosure

The maintainer will acknowledge receipt when possible; there is no SLA. Critical fixes may be released as patch tags.
