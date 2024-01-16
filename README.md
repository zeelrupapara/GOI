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

Now, create a `.env` file in the project directory and add the token:

```bash
echo "GITHUB_TOKEN=<YOUR_GITHUB_TOKEN>" > .env.example
cp .env.example .env
```
### Step 3: Start Database

To start postgres database server using docker-compose file
```bash
docker compose up
```

### Step 4: Apply Migrations

GPAT uses migrations to set up the necessary database structure. Apply the migrations using the following command:

```bash
go run main.go migrate up
```

### Step 5: Retrieve Data from Github

Fetch organization and member data from Github using:

```bash
go run main.go github
```

This command populates the database with the necessary information for managing members.

## Usage

With the setup complete, you can now use GPAT to manage members across your organizations efficiently.

Feel free to explore additional commands and functionalities provided by GPAT by checking the available options in `main.go`.
