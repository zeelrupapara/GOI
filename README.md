# Github Analytical Projects Tool (GPAT)

Welcome to Github Analytical Projects Tool (GPAT) - your go-to solution for managing members across multiple organizations on Github.

## Setup Guide

Follow these simple steps to set up and start using GPAT:

### Step 1: Clone the Repository

```bash
git clone git@github.com:Improwised/GPAT.git
cd GPAT
```

### Step 2: Set Up Github Token

To interact with the Github API, you'll need to set up a personal access token. Follow these steps to create one:

1. Visit [Github Developer Settings](https://github.com/settings/tokens) page.
2. Click on "Generate token" and select the necessary scopes (at least read access to organizations).
3. Copy the generated token.

Now, add `GITHUB_TOKEN` to `.env.docker` file for the api:

```bash
echo "GITHUB_TOKEN=<YOUR_GITHUB_TOKEN>" >> ./api/.env.example
```
### Step 3: Start All Service Using Docker

To start all server (api, ui, database) using docker-compose file
```bash
docker compose up
```
Click [GPAT](http://localhost) for access the UI 

### Step 4: Retrieve Data from Github
> This command is optional If you want the another time range data from github instad of currunt week
Fetch organization and member data from Github using:

```bash
go run main.go github
```

This command populates the database with the necessary information for managing members.

## Usage

With the setup complete, you can now use GPAT to manage members across your organizations efficiently.

Feel free to explore additional commands and functionalities provided by GPAT by checking the available options in `main.go`.
