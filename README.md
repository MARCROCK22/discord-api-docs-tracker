# Discord API Docs Tracker

Sends Discord webhook notifications when pull requests are opened, closed, or merged in a target GitHub repository (defaults to [discord/discord-api-docs](https://github.com/discord/discord-api-docs), configurable via `-repo-source`).

## Prerequisites

- [Go](https://go.dev/dl/) 1.26.1+
- A GitHub personal access token with permission to edit issues on the target repo
- A GitHub issue whose body contains a Unix timestamp in milliseconds (used to track the last check time)
- One or more Discord webhook URLs

## Setup

1. Clone the repository:

   ```sh
   git clone https://github.com/MARCROCK22/discord-api-docs-tracker.git
   cd discord-api-docs-tracker
   ```

2. Build:

   ```sh
   go build -o discord-api-docs-tracker .
   ```

## Usage

Configuration is passed as command-line flags:

```sh
./discord-api-docs-tracker \
  -token "$GITHUB_TOKEN" \
  -webhooks "https://discord.com/api/webhooks/...,https://discord.com/api/webhooks/..." \
  -repo-target owner/repo \
  -repo-source discord/discord-api-docs \
  -issue 1
```

| Flag           | Description                                                                            |
| -------------- | ------------------------------------------------------------------------------------- |
| `-token`       | GitHub PAT with permission to edit issues on the target repo                          |
| `-webhooks`    | Discord webhook URLs to post notifications (comma-separated for multiple)             |
| `-repo-target` | GitHub repo that holds the tracker issue (e.g. `MARCROCK22/discord-api-docs-tracker`) |
| `-repo-source` | GitHub repo to watch for pull requests (e.g. `discord/discord-api-docs`)              |
| `-issue`       | Issue number used to store the last check timestamp                                   |

Run `./discord-api-docs-tracker -h` to list the flags.

Each execution performs a single check-and-notify cycle. To run it on a schedule, use a cron job or a task scheduler:

```sh
# Example: run every 5 minutes via cron
*/5 * * * * /path/to/discord-api-docs-tracker -token ... -webhooks ... -repo-target owner/repo -repo-source discord/discord-api-docs -issue 1
```
