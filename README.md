<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dependency Check PoC Workflow</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            max-width: 800px;
            margin: 20px auto;
            padding: 0 20px;
        }
        h1 {
            color: #333;
            border-bottom: 2px solid #333;
            padding-bottom: 10px;
        }
        h2 {
            color: #555;
            margin-top: 20px;
        }
        ul {
            list-style-type: disc;
            margin-left: 20px;
        }
        li {
            margin-bottom: 10px;
        }
    </style>
</head>
<body>
    <h1>Dependency Check PoC Workflow</h1>
    <p>This GitHub Actions workflow automates dependency scanning for security vulnerabilities in a repository. It runs on every push event, detects programming languages, installs necessary dependencies, performs a dependency check, and reports critical vulnerabilities (CVSS score ≥ 7) by creating a GitHub issue if any are found.</p>

    <h2>Overview</h2>
    <p>The workflow executes the following steps:</p>
    <ul>
        <li><strong>Repository Checkout:</strong> Clones the repository to the runner environment.</li>
        <li><strong>Language Detection:</strong> Identifies programming languages used in the project based on file extensions and configuration files (e.g., Java, Python, JavaScript, etc.).</li>
        <li><strong>Environment Setup:</strong> Installs runtime environments and tools for detected languages (e.g., JDK, Node.js, Python, etc.).</li>
        <li><strong>Dependency Installation:</strong> Installs project dependencies based on the detected languages and their package managers.</li>
        <li><strong>Java Setup:</strong> Configures JDK 11 specifically for the dependency check tool.</li>
        <li><strong>Dependency Check:</strong> Scans the project for vulnerabilities using the Dependency-Check tool, generating reports in multiple formats.</li>
        <li><strong>Report Storage:</strong> Saves the scan reports to a designated directory.</li>
        <li><strong>Artifact Upload:</strong> Uploads the generated reports as workflow artifacts for review.</li>
        <li><strong>Vulnerability Check:</strong> Analyzes the report and creates a GitHub issue if critical vulnerabilities (CVSS ≥ 7) are detected, including details about affected files and libraries.</li>
    </ul>

    <h2>Purpose</h2>
    <ul>
        <li>Automate dependency vulnerability scanning.</li>
        <li>Ensure timely detection of critical security issues.</li>
        <li>Provide actionable feedback via GitHub issues and downloadable reports.</li>
    </ul>

    <h2>Trigger</h2>
    <ul>
        <li>Runs on every <code>push</code> event to the repository.</li>
    </ul>

    <h2>Outputs</h2>
    <ul>
        <li>Dependency scan reports (available as artifacts).</li>
        <li>GitHub issue with details if critical vulnerabilities are found.</li>
    </ul>
</body>
</html>
