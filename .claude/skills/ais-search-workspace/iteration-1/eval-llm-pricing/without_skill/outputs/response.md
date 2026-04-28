# LLM API Cost Comparison for 10M Tokens/Day

> **Disclaimer:** Pricing data is based on training knowledge (cutoff ~August 2025). Always verify current prices on each provider's official pricing page.

## Pricing (as of ~mid-2025)

| Model | Input ($/M tokens) | Output ($/M tokens) |
|---|---|---|
| Claude Sonnet (latest) | $3.00 | $15.00 |
| GPT-4o (latest) | $5.00 | $15.00 |
| Gemini 2.0 Flash | $0.075 | $0.30 |

## Cost Calculation for 10M tokens/day (70% input, 30% output)

That is 7M input tokens + 3M output tokens per day.

| Model | Input Cost/Day | Output Cost/Day | Total/Day | Total/Month |
|---|---|---|---|---|
| Claude Sonnet | $21.00 | $45.00 | **$66.00** | ~$1,980 |
| GPT-4o | $35.00 | $45.00 | **$80.00** | ~$2,400 |
| Gemini 2.0 Flash | $0.525 | $0.90 | **$1.425** | ~$42.75 |

**Winner: Gemini 2.0 Flash** — approximately 46x cheaper than Claude Sonnet and 56x cheaper than GPT-4o.

## Context Windows

| Model | Context Window |
|---|---|
| Claude Sonnet | 200,000 tokens |
| GPT-4o | 128,000 tokens |
| Gemini 2.0 Flash | 1,000,000 tokens |

## Key Caveats

- **Gemini 2.0 Flash** pricing doubles for prompts exceeding 128K tokens. Quality is lower than the other two for complex reasoning tasks.
- **Claude Sonnet** offers a Batch API at 50% discount (~$33/day) for async workloads.
- **GPT-4o** also has a Batch API at 50% discount (~$40/day). GPT-4o mini ($0.15/$0.60 per M) is far cheaper for simpler tasks.
- Rate limits: all three comfortably support 10M tokens/day on paid tiers. Gemini 2.0 Flash has the highest TPM limits (~4M TPM).
