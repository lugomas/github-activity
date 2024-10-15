# GitHub Activity CLI
`github-activity-cli` is a command-line interface (CLI) tool that allows you to fetch and display the recent activity of any GitHub user directly from the terminal. The application uses the GitHub API to retrieve and display activities such as push events, issues, stars, and more.

## Features
- Fetch and display a GitHub user's recent activity.
- Display activities such as push events, issues opened, stars, forks, etc.
- Gracefully handle errors such as invalid GitHub usernames or API failures.
- No external libraries or frameworks are used to fetch the data.

## Installation
To install github-activity, clone the repository and build the Go binary:
```
git clone https://github.com/lugomas/github-activity-cli.git
cd github-activity-cli
go build -o github-activity
```

You can now use the github-activity binary to fetch and view GitHub activity.

## Usage
```
### Fetching the activity of a GitHub user
./github-activity <username>
```

Notes:
- The tool uses the following GitHub API endpoint to fetch the user’s activity:
https://api.github.com/users/<username>/events
- The GitHub API limits the number of requests from an IP without authentication. Be mindful of the rate limits.

## License
This project is licensed under the MIT License.

## Project Inspiration
This project was developed based on the guidelines provided by the Project [GitHub User Activity](https://roadmap.sh/projects/github-user-activity).
It allows users to easily track the recent activity of any GitHub user from the command line, following best practices for CLI tools.