# WQB Dispatch / Recovery Controller v1

状态：解决“抛出后马上撞 polling 导致 429”的控制层  
生成时间：2026-06-22T10:30:00+08:00

## Problem

Batch dispatch returns `201 Created` plus `Retry-After`. If the agent immediately polls several simulations, WQB may return `429 API rate limit exceeded`. A good flywheel needs a persistent recovery queue, not an eager polling loop.

## Contract

After authorized dispatch:

1. Create a batch ledger with simulation ids/urls.
2. Create `runs/wqb-current/recovery-queue.json`.
3. Stagger `next_poll_at` by the returned `Retry-After`.
4. Poll only due simulations, usually one per cycle.
5. On `429`, stop the polling cycle immediately, back off that item, and persist queue state.
6. Do not dispatch a new batch until the current batch is recovered or explicitly abandoned.

## Modes

- `init-from-dispatch`: convert a dispatched batch + create_multi_simulation response into a recovery queue.
- `next-poll`: return whether to wait or which simulation is due.
- `ingest-poll`: update the queue from one poll result.
- `self-test`: deterministic local test for wait/stagger/backoff behavior.

## Backoff

Default 429 backoff ladder:

- 5 min
- 15 min
- 30 min
- 60 min
- 2 h

If WQB returns a longer `Retry-After`, use the longer value.

## Efficiency rule

Efficiency is not “poll immediately”; efficiency is “never waste platform calls while preserving exact recovery state.”

## Agent operating rule

The agent should not ask the user to type “continue” just to recover an already-dispatched batch. In a live turn, the agent should either:

- wait until the next due poll and poll one due simulation; or
- if the next due poll is outside the practical turn window, persist the queue and report the exact next recovery time.

No new batch may be dispatched until the current recovery queue is ledgered.
