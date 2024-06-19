# GitHub Organization Insight (GOI)

Welcome to GitHub Organization Insight (GOI) - your go-to solution for managing members across multiple organizations on Github.

## Setup Guide

Follow these simple steps to set up and start using GOI:

### Step 1: Clone the Repository

```bash
git clone https://github.com/zeelrupapara/GOI.git
cd GOI
```

### Step 2: Set Up Github Token

To interact with the Github API, you'll need to set up a personal access token. Follow these steps to create one:

1. Visit [Github Developer Settings](https://github.com/settings/tokens) page.
2. Click on "Generate token" and select the necessary scopes (at least read access to organizations).
3. Copy the generated token.

Now, add `GITHUB_TOKEN` to `.env.docker` file for the api:

```bash
echo "GITHUB_TOKEN=<YOUR_GITHUB_TOKEN>" >> ./api/.env.docker
```
### Step 3: Start All Service Using Docker

To start all server (api, ui, database) using docker-compose file
```bash
docker compose up
```
See **GOI**: http://localhost:80

**API** (Optional): http://localhost:8080
**UI** (Optional): http://localhost:5000

---

### Retrieve Data from Github
> This command is optional If you want to get the custom range data from github then try `cd api && go run main.go github --help`

Fetch organization and member data from Github using:
> By default last one week of github data fetch by below command

```bash
cd api
# copy env from .env.example file
cp .env.example .env
# After create .env add GITHUB_TOKEN in that env
# get github data using this command
go run main.go github
```

## Usage

With the setup complete, you can now use GOI to manage members across your organizations efficiently.

Feel free to explore additional commands and functionalities provided by GOI by checking the available options in `go run main.go --help`.
