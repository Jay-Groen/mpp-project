# Branching Model Overview
I use a lightweight GitHub Flow model with hotfix support.

## Main Branch
- main
- Always stable and releasable
- Protected (PR required)

## Feature Branches
- feature/<short-description>
- Used for new user stories
- Merged into main via pull request

## Hotfix Branches
- hotfix/<version>
- Created from a released tag
- Used to fix production bugs
- Merged back into main
- Tagged with a patch version

## Workflow Diagram
Simple ASCII diagram:

```
main ──────●─────●────────●─────────●
            \     \        \
feature/A    ●─────●
feature/B          ●─────●
hotfix/v0.2.1             ●
```

Or:

```
            feature/xyz
                 ↓
main ────────────●─────────────●──────
                        ↑
                   hotfix/v0.2.1
```
