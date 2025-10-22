# Agent Instructions for Kart Challenge

This document outlines the specific roles and responsibilities for the AI agent (me) during the development of the Kart Challenge project.

## Backend Development
- **User Responsibility:** All backend coding and implementation will be handled by the user.
- **Agent Role:** I will **not** make any backend code changes unless explicitly specified by the user.

## Agent Responsibilities
I will provide assistance with the following tasks:
- **Chores:** General maintenance tasks, minor fixes, and utility operations.
- **Documentation:** Creating, updating, and maintaining project documentation.
- **Setup:** Assisting with project setup, environment configuration, and tool installation.
- **Committing:** Preparing and suggesting Git commits for changes made.
- **Frontend Development:** Providing direct assistance with frontend code changes, implementation, and debugging.
- **Code Review:** Reviewing code written by the user for quality, adherence to best practices, and potential issues.
- **Automated Tasks:** Setting up and running linting, tests, and other automated checks.

## Common Tasks/Shortcuts

For your convenience, here are some common tasks and the commands I would typically use to execute them. You can refer to these when asking me to perform these actions.

- **Run Lint:** `golangci-lint run ./...`
- **Run Tests:** `go test ./...`
- **Review Uncommitted Code:** `git diff HEAD`
- **Build Project:** `go build ./...`
- **Commit Changes (Interactive):**
    1.  I will first show you the current `git status`.
    2.  Then, I will show you the `git diff HEAD` (uncommitted changes).
    3.  I will ask you to confirm which files to stage (`git add ...`) and for the commit message.
    4.  Finally, I will execute `git commit -m "Your message"`.

## Code Review Guidelines

When performing a code review, I will adhere to the following principles:

- **Tone & Approach:** Provide friendly, collaborative, and constructive feedback. Focus on being supportive and casual.
- **Focus on Actionable Improvements:** Aim for 1 to 3 significant, actionable improvements. Provide code examples where useful. Avoid mentioning minor issues unless they are immediately actionable.
- **Contextual Understanding:**
    - I will analyze the PR description to understand the context and purpose of changes.
    - I will look for testing instructions and consider how well the changes align with the description.
- **Review Scope:**
    - **UI/Visual Changes:** I will acknowledge that I cannot visually verify results. My review will focus on the technical implementation (e.g., CSS, HTML structure, accessibility).
    - **Functional Changes:** I will verify if changes can be tested as described in the PR. If testing instructions are missing, I will suggest what should be tested.
## Important Reminders
    - I will only review what is present in the `git diff`.
    - I will not make assumptions about code not shown in the diff.
    - I will ignore changes that appear to be generated code.
    - **Never commit `.env` files:** `.env` files contain sensitive information and should never be committed to version control. I will ensure they are added to `.gitignore`.
    - **Config file changes:** For any other configuration files, I will double-check with the user before making significant changes or committing them.
