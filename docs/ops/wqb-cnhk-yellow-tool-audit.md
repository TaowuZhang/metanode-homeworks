# WQB CNHK MCP Yellow Tool Audit

## Correction

The CNHK MCP was underused. The earlier system mostly used the yellow3 connection for simulation, recovery, submission checks, correlations, and datafields. That was not enough.

The richer `mcpServer_cnhkmcp_wqb_cloudflare_yellow` connection exposes more read-only research context that should feed the architecture before choosing experiments.

## Connection comparison observed on 2026-06-22

### yellow3 / yellow2 / yellow4

Core tools available:

- `health`
- `policy`
- `tools`
- `config_status`
- `auth_status`
- `get_datasets`
- `get_datafields`
- `create_simulation`
- `create_multi_simulation`
- `get_simulation`
- `check_correlation`
- `get_submission_check`
- `get_alpha_pnl`
- `get_user_alphas`

These are sufficient for dispatch and recovery but incomplete for learning platform/community priors.

### yellow

Additional useful tools observed:

- `dataset_brief`
- `field_brief`
- `get_documentations`
- `get_documentation_page`
- `get_alpha_detail`
- `summarize_alpha`
- `get_alpha_yearly_stats`
- `get_pyramid_alphas`
- `get_user_competitions`
- `get_competition_details`
- `get_competition_agreement`
- `get_leaderboard`
- `get_pyramid_multipliers`
- `get_user_activities`
- `get_alpha_recordsets`
- `get_alpha_recordset`
- `get_alpha_before_after_performance`

This connection should be the default for read-only research planning and platform/community context. Simulation may still use the most reliable connected yellow endpoint, but planning should not ignore the richer yellow tools.

## Tool routing rule

1. Use `mcpServer_cnhkmcp_wqb_cloudflare_yellow` for:
   - WQB documentation/tutorial pages;
   - community / consultant guidance;
   - pyramid multipliers and incentives;
   - leaderboard aggregate priors;
   - alpha detail / yearly stats / recordsets / before-after diagnostics;
   - dataset / field briefs.

2. Use simulation-enabled endpoint for:
   - authorized `create_multi_simulation`;
   - recovery `get_simulation`;
   - submission / correlation / PnL reads if stable.

3. If one CNHK endpoint returns 429/captcha/auth/tool failure, verify with `auth_status` and use another available CNHK endpoint only for safe read-only recovery, not to bypass platform rules.

## Immediate architecture changes from Yellow audit

- Add documentation pages as first-class priors.
- Add consultant/community recommendations to the belief graph.
- Add anti-duplicate simulation cache before dispatch.
- Add pyramid multiplier priors to dataset/category allocation.
- Add leaderboard-derived breadth prior: top performers use many data fields; narrow local repair loops are structurally weak.
- Add yearly/PnL recordset diagnostics before deep repair where available.

## Read-only pages discovered

Important WQB documentation/tutorial pages available through Yellow:

- `list-must-read-posts-how-improve-your-alphas-are-submitted`
- `getting-started-finding-consultant-alphas-read-first`
- `consultant-dos-and-donts`
- `consultant-submission-tests`
- `how-can-you-avoid-duplicate-simulations`
- `simulation-settings`
- `vector-datafields`
- `neut-cons`
- `neut-users`
- `getting-started-risk-neutralized-alphas`
- `getting-started-crowding-risk-neutralized-alphas`
- `understanding-simulation-limits`
- `brain-api`
- `single-dataset-alphas`
- `brain-genius`

## Resulting principle

Do not claim to understand WQB from generic quant language. Read the platform's own tutorial/community/consultant pages, turn them into priors, then let the scheduler decide.
