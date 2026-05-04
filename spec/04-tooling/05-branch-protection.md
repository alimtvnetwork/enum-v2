# 05 — Branch Protection

> ✅ **Status**: drafted at spec-v0.5.0 (2026-05-04, Asia/Kuala_Lumpur).
> **Audience**: repo admins configuring GitHub branch protection.
> **Companion**: workflows in `.github/workflows/` and CI guards in
> [`04-ci-guards.md`](./04-ci-guards.md).

---

## 1. Protected Branches

| Branch pattern | Purpose | Protection level |
|---|---|---|
| `main` | Stable trunk. Source of releases. | **Strict** |
| `release/**` | Release-candidate branches consumed by `release.yml`. | **Strict** |
| `feature/**`, `fix/**` | Short-lived dev branches. | None (PR target only) |

---

## 2. Required Status Checks

Configure these as **required** on `main` and `release/**` (GitHub →
Settings → Branches → Branch protection rule → *Require status checks
to pass before merging*):

| Check name | Workflow | Job |
|---|---|---|
| `lint` | `ci.yml` | `lint` |
| `vulncheck` | `ci.yml` | `vulncheck` |
| `build-check` | `ci.yml` | `build-check` |
| `test-summary` | `ci.yml` | `test-summary` (aggregates the 4 test-matrix suites) |
| `collision-audit` | `ci-guards.yml` | `collision-audit` |
| `lint-baseline-diff` | `ci-guards.yml` | `lint-baseline-diff` |

> ℹ️ The four matrix suites (`unit`, `integration`, `creation`, `tests`)
> intentionally aren't required individually — `test-summary` depends on
> all of them and surfaces a single, stable check name. This avoids
> branch-protection churn whenever the matrix is reshaped.

### Recommended toggles

- ☑ **Require branches to be up to date before merging** — forces a
  rebase/merge from `main` so the required checks run against the
  post-merge state.
- ☑ **Require linear history** — keeps `git log --oneline` readable for
  the changelog extractor in `release.yml`.
- ☑ **Require signed commits** *(optional but recommended)*.
- ☑ **Do not allow bypassing the above settings** — applies to admins
  too. Override via a single, audited "break glass" rule if needed.

---

## 3. Pull Request Rules

- ☑ **Require a pull request before merging**
  - Required approving reviews: **1** (raise to 2 before 1.0).
  - Dismiss stale approvals on new commits.
  - Require review from Code Owners (uses `CODEOWNERS`).
- ☑ **Require conversation resolution before merging**.

---

## 4. Tag Protection

Protect the `v*` tag pattern (Settings → Tags → New rule → `v*`). Only
maintainers may create or move release tags — `release.yml` triggers on
these tags and will publish a GitHub Release on success.

---

## 5. Verification Checklist

After enabling protection, verify by opening a throwaway PR:

1. Push a branch with a deliberate lint violation → `lint-baseline-diff`
   reports it as **NEW** and the PR is blocked.
2. Push a branch that introduces an exported identifier collision →
   `collision-audit` fails.
3. Push a branch with a failing test → the relevant matrix shard fails;
   `test-summary` rolls up to a red required check.
4. Approve + merge a clean PR → all six required checks turn green and
   the merge button enables.
5. Tag the merge commit `vX.Y.Z` → `release.yml` publishes the GitHub
   Release with the `[vX.Y.Z]` section of `CHANGELOG.md` as the body.

---

## See Also

- [`02-release-pipeline.md`](./02-release-pipeline.md) — what happens
  after a `v*` tag is pushed.
- [`03-vulnerability-scanning.md`](./03-vulnerability-scanning.md) — the
  classification rules that drive the `vulncheck` required check.
- [`04-ci-guards.md`](./04-ci-guards.md) — collision audit + lint
  baseline-diff internals.
- `CONTRIBUTING.md` — contributor-facing summary of the merge flow.
