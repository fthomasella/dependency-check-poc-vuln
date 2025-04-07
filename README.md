# Dependency Check PoC Workflow

This GitHub Actions workflow automates dependency scanning for security vulnerabilities in a repository. It runs on every push event, detects programming languages, installs necessary dependencies, performs a dependency check, and reports critical vulnerabilities (CVSS score ≥ 7) by creating a GitHub issue if any are found.

## Overview

The workflow executes the following steps:

- **Repository Checkout**: Clones the repository to the runner environment.
- **Language Detection**: Identifies programming languages used in the project based on file extensions and configuration files (e.g., Java, Python, JavaScript, etc.).
- **Environment Setup**: Installs runtime environments and tools for detected languages (e.g., JDK, Node.js, Python, etc.).
- **Dependency Installation**: Installs project dependencies based on the detected languages and their package managers.
- **Java Setup**: Configures JDK 11 specifically for the dependency check tool.
- **Dependency Check**: Scans the project for vulnerabilities using the Dependency-Check tool, generating reports in multiple formats.
- **Report Storage**: Saves the scan reports to a designated directory.
- **Artifact Upload**: Uploads the generated reports as workflow artifacts for review.
- **Vulnerability Check**: Analyzes the report and creates a GitHub issue if critical vulnerabilities (CVSS ≥ 7) are detected, including details about affected files and libraries.

## Purpose

- Automate dependency vulnerability scanning.
- Ensure timely detection of critical security issues.
- Provide actionable feedback via GitHub issues and downloadable reports.

## Trigger

- Runs on every `push` event to the repository.

## Outputs

- Dependency scan reports (available as artifacts).
- GitHub issue with details if critical vulnerabilities are found.
