# Gemini 2.0 Flash API — Current Rate Limits & Free Tier Quotas

> **Note:** The free tier quotas for Gemini 2.0 Flash were significantly reduced in December 2025.

## Free Tier (Google AI Studio) — Post-December 2025

| Limit | Value |
|-------|-------|
| Requests Per Minute (RPM) | **5 RPM** (reduced from 10) |
| Requests Per Day (RPD) | **100 RPD** (reduced from 500 — an 80% decrease) |
| Tokens Per Minute (TPM) | ~250,000 TPM (explicitly stated for Gemini 2.5 Flash; similar tier) |

The December 2025 update reduced quotas across multiple dimensions with minimal advance notice.

## Key Notes

- Rate limits are applied **per project**, not per API key — creating multiple API keys within the same Google Cloud project does not grant additional quota.
- Exceeding any limit triggers a **429 quota-exceeded error**.
- Since March 2024, linking a billing account is required to activate free tier quotas; without it, content generation is limited to 0 and returns `429 ResourceExhausted`.

## Context: Other Free Tier Models (early 2026)

| Model | RPM | RPD | TPM |
|-------|-----|-----|-----|
| Gemini 2.5 Pro | 5 | 100 | 250,000 |
| Gemini 2.5 Flash | 10 | 250 | 250,000 |
| Gemini 2.0 Flash | 5 | 100 | ~250,000 |

*(Sources retrieved via `ais` CLI — see sources.txt)*
