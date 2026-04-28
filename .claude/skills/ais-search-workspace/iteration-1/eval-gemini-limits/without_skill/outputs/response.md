# Gemini 2.0 Flash API: Rate Limits and Free Tier Quotas

> **Important Caveat**: This information is based on training data with a knowledge cutoff of August 2025. Google frequently updates API pricing and quotas. Always refer to the official Google AI documentation for the most current limits.

## Free Tier (Google AI Studio / Gemini API via API Key)

As of my training knowledge, the free tier for Gemini 2.0 Flash includes:

### Requests Per Minute (RPM)
- **15 RPM** (requests per minute)

### Tokens Per Minute (TPM)
- **1,000,000 TPM** (tokens per minute)

### Requests Per Day (RPD)
- **1,500 RPD** (requests per day)

## Paid Tier Rate Limits

For the paid tier (billed usage), Gemini 2.0 Flash generally offers significantly higher limits:

### Requests Per Minute (RPM)
- **2,000 RPM**

### Tokens Per Minute (TPM)
- **4,000,000 TPM**

### Tokens Per Day (TPD)
- No hard daily token cap (metered billing)

## Notes on Gemini 2.0 Flash Variants

Gemini 2.0 Flash has several variants that may have different limits:
- **gemini-2.0-flash** — the main model
- **gemini-2.0-flash-lite** — a lighter, more cost-efficient variant, may have different quotas
- **gemini-2.0-flash-thinking** (experimental) — may have lower rate limits due to experimental status

## Context Window

- **Input context window**: 1,000,000 tokens
- **Output token limit**: 8,192 tokens (standard); some variants may support more

## Disclaimer

These figures reflect what was documented as of mid-2025. Google has a history of revising these numbers — sometimes increasing free tier limits as models mature. The authoritative source is:

- https://ai.google.dev/gemini-api/docs/rate-limits (for current limits)
- https://ai.google.dev/pricing (for pricing and quota details)
