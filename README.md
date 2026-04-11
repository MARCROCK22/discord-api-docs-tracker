# Discord API Docs Tracker

Sends Discord webhook notifications when pull requests are opened, closed, or merged in the [discord/discord-api-docs](https://github.com/discord/discord-api-docs) repository.

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

2. Install dependencies:

   ```sh
   go mod download
   ```

3. Create a `.env` file in the project root:

   ```env
   GITHUB_TOKEN=your_github_personal_access_token
   WEBHOOK_URLS=https://discord.com/api/webhooks/...,https://discord.com/api/webhooks/...
   REPO_TARGET=owner/repo
   NUMBER_OF_ISSUE=1
   ```

   | Variable          | Description                                                        |
   | ----------------- | ------------------------------------------------------------------ |
   | `GITHUB_TOKEN`    | GitHub PAT with permission to edit issues on the target repo       |
   | `WEBHOOK_URLS`    | Discord webhook URLs to post notifications (comma-separated for multiple) |
   | `REPO_TARGET`     | GitHub repo that holds the tracker issue (e.g. `MARCROCK22/discord-api-docs-tracker`) |
   | `NUMBER_OF_ISSUE` | Issue number used to store the last check timestamp                |

## Usage

Run the application:

```sh
go run .
```

Each execution performs a single check-and-notify cycle. To run it on a schedule, use a cron job or a task scheduler:

```sh
# Example: run every 5 minutes via cron
*/5 * * * * cd /path/to/discord-api-docs-tracker && go run .
```
