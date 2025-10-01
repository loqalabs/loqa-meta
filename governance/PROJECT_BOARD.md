# Project Board Operating Guide

## Purpose
The Ambiware Labs project board provides a single, cross-repo view of Loqa’s roadmap execution. It helps track progress toward the MVP and subsequent milestones while keeping community contributors aligned.

## Board Structure
Columns are mapped to delivery states:
1. **Backlog** – Groomed issues aligned with MVP backlog workstreams
2. **Ready** – Prioritized for the current sprint/iteration with clear acceptance criteria
3. **In Progress** – Actively being worked by an assignee
4. **Review** – Pull request open and awaiting review or validation
5. **Done** – Merged & deployed, validation complete
6. **Blocked** – Items requiring decisions, dependencies, or upstream fixes

## Issue Tagging
- `type:feature`, `type:bug`, `type:rfc`, `type:docs` classify the work
- `area:runtime`, `area:voice`, `area:skills`, `area:platform`, `area:community` map to Workstreams A–E
- `priority:p0/p1/p2` reflect urgency

## Lifecycle
1. **Intake** – New issues open in the repo they impact; triage labels + link to relevant RFC or epic
2. **Refinement** – During weekly triage, ensure description, acceptance criteria, and owner are assigned before moving to Ready
3. **Execution** – Move cards across the board as status changes; keep checklists current
4. **Review** – Require at least one approving review and passing CI prior to moving to Done
5. **Retrospective** – At month end review completed work, cycle time, and carry-over tasks

## Board Maintenance Checklist
- [ ] Weekly triage (Mondays) – backlog grooming, new issues labeled
- [ ] Mid-week sync (Thursdays) – unblock In Progress/Blocked items
- [ ] Monthly review (last Friday) – update roadmap, post summary to Discussions

## Automation Ideas
- Use `gh project item-add` via GitHub Actions to auto-add labeled issues
- Configure workflow to move cards when PRs merge (using project automation)
- Generate progress reports from board data for monthly updates

## Links
- GitHub Project: `https://github.com/orgs/loqalabs/projects/1` (create and update URL once live)
- MVP Backlog: [roadmap/MVP_BACKLOG.md](../roadmap/MVP_BACKLOG.md)
- RFC Index: [rfcs/](../rfcs)

## Ownership
- Board Maintainer: Anna Barnes (@annabarnes1138)
- Backup Maintainer: TBD (assign once additional maintainers join)

Reassess this guide quarterly and update as team/process evolves.
