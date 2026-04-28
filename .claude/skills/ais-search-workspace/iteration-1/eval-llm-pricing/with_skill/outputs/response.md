# LLM API Pricing Comparison — March 2026

*Retrieved via `ais -temperature 0.2 -thinking-budget none` for fast, consistent factual results.*

## Per-Million-Token Prices (Input / Output)

| Model | Input ($/M) | Output ($/M) |
|-------|-------------|--------------|
| **Claude Sonnet 4.6** | $3.00 | $15.00 |
| **Claude Haiku 4.5** | $1.00 | $5.00 |
| **Claude Opus 4.6** | $5.00 | $25.00 |
| **GPT-4o** | $2.50–$5.00 | $10.00–$15.00 |
| **GPT-4o mini** | $0.15 | $0.60 |
| **Gemini 2.0 Flash** | $0.10 | $0.40 |
| **Gemini 2.5 Flash-Lite** | $0.10 | $0.40 |
| **Grok 4.1 (xAI)** | $0.20 | $0.50 |

> **Note:** Claude Sonnet 4.6 prices double beyond 200K token prompts.

## Cost at Your Scale (10M tokens/day, 70% input / 30% output)

Daily: 7M input tokens + 3M output tokens

| Model | Daily Cost | Monthly Cost |
|-------|-----------|--------------|
| **Gemini 2.0 Flash** | $1.90 | ~$57 |
| **GPT-4o mini** | $2.85 | ~$85 |
| **Claude Haiku 4.5** | $22.00 | ~$660 |
| **Claude Sonnet 4.6** | $66.00 | ~$1,980 |
| **GPT-4o** | ~$47.50 | ~$1,425 |

## Recommendation

**Gemini 2.0 Flash** is the clear winner at this scale — **~97% cheaper than Claude Sonnet 4.6** and cheaper than GPT-4o mini. At 10M tokens/day it's roughly $57/month vs ~$1,980/month.

If you need higher quality than Flash provides, **GPT-4o mini** is the next-cheapest option at ~$85/month.

## Context Window & Rate Limit Notes

- **Gemini 2.0 Flash**: 1M token context window; free tier now limited to 5 RPM / 100 RPD (reduced Dec 2025)
- **Claude Sonnet 4.6**: 200K token context window; pricing tiers above 200K
- **GPT-4o**: 128K context window

*(Sources retrieved via `ais` CLI — see sources.txt)*
