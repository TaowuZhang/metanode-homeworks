# WQB Nonlinear Search Architecture

## Correction

A linear loop is not an architecture. A linear loop is only plumbing:

```text
proposal -> simulation -> score -> next proposal
```

That loop can be useful, but it cannot be the strategic layer. The market is not solved by a silver bullet expression, and WQB search space is not crossed by a single queue. The correct system is a nonlinear adaptive architecture.

## System thesis

The goal is not to prove that one formula will beat the market. The goal is to build a constrained scientific search machine:

```text
maximize expected information gain
under simulation budget, rate limits, gate constraints, and memory persistence
```

A failed alpha is useful only if it changes the belief graph and the next experiment distribution.

## Architecture overview

```text
                  +----------------------+
                  |  External Priors     |
                  |  docs / community    |
                  +----------+-----------+
                             |
                             v
+-------------------+   +----+----------------+   +------------------+
| Historical Ledger |-->| Belief Graph        |<--| Gate / PnL / Corr |
| Batches + Gates   |   | Star Map + Terrain  |   | Observations     |
+-------------------+   +----+----------------+   +------------------+
                             |
                             v
                 +-----------+------------+
                 | Multi-Agent Proposers  |
                 | Scout / Climb / Repair |
                 | Diagnostic / Adversary |
                 +-----------+------------+
                             |
                             v
                 +-----------+------------+
                 | Portfolio Scheduler    |
                 | constrained optimizer  |
                 +-----------+------------+
                             |
                             v
                 +-----------+------------+
                 | Experiment Factory     |
                 | legal FASTEXPR only    |
                 +-----------+------------+
                             |
                             v
                 +-----------+------------+
                 | Platform Executor      |
                 | authorized simulation  |
                 +-----------+------------+
                             |
                             v
                 +-----------+------------+
                 | Recovery Controller    |
                 | backoff + persistence  |
                 +-----------+------------+
                             |
                             v
                 +-----------+------------+
                 | Bayesian Terrain Update|
                 +------------------------+
```

## 1. Belief graph, not flat heatmap

A heatmap table is insufficient. The durable object should be a belief graph.

### Node types

- `dataset`
- `field_family`
- `field`
- `field_type`
- `operator_archetype`
- `horizon_bucket`
- `neutralization_family`
- `failure_mode`
- `regime`
- `candidate`
- `alpha_result`
- `constellation`

### Edge types

- `belongs_to`
- `shares_mechanism_with`
- `failed_by`
- `repaired_by`
- `decorrelates_from`
- `is_sibling_of`
- `inherits_from_parent`
- `contradicts`
- `confirms`

### Belief vector per star

```json
{
  "star_id": "USA_TOP3000_D1_behavioral_vector_crossfield_medium_slow",
  "belief": {
    "expected_signal": 0.0,
    "expected_stability": 0.0,
    "expected_originality": 0.0,
    "expected_tradability": 0.0,
    "uncertainty": 1.0,
    "repairability": 0.0,
    "crowding_risk": 0.0,
    "sample_count": 0,
    "last_information_gain": 0.0
  }
}
```

The scheduler reads the belief graph, not just the last batch.

## 2. Multi-agent proposer swarm

A single generator is too linear. The architecture needs several specialized proposers that compete for batch slots.

### Scout Agent

Purpose: search unsampled or undersampled stars.

Output: candidates with high uncertainty bonus and low overcrowding priors.

### Climber Agent

Purpose: exploit a promising parent or repair seed.

Output: population around a star with evidence.

### Repair Agent

Purpose: target one blocker, such as LOW_2Y_SHARPE or PROD_CORRELATION.

Output: candidates designed to change a specific gate vector component.

### Diagnostic Agent

Purpose: test if the system's hypothesis is wrong.

Output: inverse sign, operator sibling, field sibling, horizon diagnostic.

### Risk-Adversary Agent

Purpose: prevent seductive but bad local search.

Output: kill rules, repeated-blocker warnings, overfitting flags.

### Archivist Agent

Purpose: maintain durable memory.

Output: ledgers, heatmap snapshots, belief graph updates, context recovery state.

## 3. Portfolio scheduler, not next-step selector

The scheduler should allocate a portfolio of six experiments, not pick a single next idea.

```text
maximize sum(expected_information_gain(slot_i))
subject to:
  batch_size = 6
  platform_policy = satisfied
  at least 2 independent mechanisms unless near-gate
  at least 1 diagnostic after repeated blocker
  no more than 2 candidates from same field unless near-gate
  no more than 3 from same constellation unless near-gate
  no candidate without a kill rule
```

## 4. Expected information gain

Each proposed candidate must estimate what it will teach.

```text
EIG(candidate) = uncertainty_reduction
               + blocker_resolution_value
               + mechanism_discrimination_value
               + correlation_information_value
               + parent_validation_value
               - redundancy_penalty
               - platform_risk_penalty
```

A candidate that is merely another nearby parameter tweak has low EIG unless the parent is near-gate.

## 5. Nonlinear feedback

A result updates more than one thing.

Example: Batch E failing does not only mark Batch E failed. It updates:

- D02 parent belief;
- behavioral vector constellation;
- CROWDING neutralization risk;
- prod correlation blocker prior;
- future slot allocation;
- local repair depth budget;
- diagnostic need.

That is nonlinear feedback. One observation modifies many future probabilities.

## 6. Regime of operation

### If near-gate exists

Use focused repair:

```text
confirm: 2
repair: 2
diagnostic: 1
scout: 1
```

### If promising but blocked

Use mixed climb:

```text
climb: 2
repair: 1
diagnostic: 1
scout: 2
```

### If repeated local repair fails

Use star-map expansion:

```text
scout: 3
diagnostic: 2
climb: 1
```

### If no strong signal exists

Use broad scout:

```text
scout: 4
diagnostic: 2
```

## 7. What counts as progress

Progress is not number of simulations.

Progress is:

- a gate turns from FAIL to PASS;
- a blocker becomes quantified;
- a family is killed with evidence;
- a new constellation becomes promising;
- a heatmap region is proven weak;
- correlation risk is mapped;
- the next distribution changes rationally.

If a batch does not change the distribution, it was mostly wasted.

## 8. How this survives context reset

The following files are the memory boundary:

- `runs/wqb-current/run-state.json`
- `runs/wqb-current/starmap-grid.json`
- `runs/wqb-current/starmap-weight-model.json`
- `runs/wqb-current/heatmap-summary-*.json`
- `runs/wqb-current/continuation-bridge.json`
- `runs/wqb-current/authorization-ready-*.json`
- `runs/wqb-batches/BATCH-*.json`

A new session must read these and resume from the architecture state, not from chat memory.

## 9. Relationship to platform limits

The architecture cannot and should not bypass WQB constraints. It can exceed the limits in a different sense: by doing more thinking per simulation.

Platform limit:

```text
few authorized simulations, rate limits, captcha/auth boundaries, no submit tool
```

Architecture response:

```text
higher information per simulation, persistent memory, nonlinear belief updates, no empty stops
```

## 10. Current correction

The current state after Batch E should not be treated as a simple Batch F queue. It should be treated as a belief-graph correction:

- local behavioral D02 repair failed;
- pure behavioral local mutation is over-sampled;
- fnd65 has unspent repair budget;
- ai_factor may be operator-template failure rather than dataset failure;
- analyst has inverse/diagnostic value;
- next batch must be portfolio-style, not single-family linear continuation.

Therefore the recommended next batch is:

```text
BATCH-20260622-F-starmap-scout-and-climb
```

with mixed slots rather than one local ridge.
