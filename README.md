# Gator

I don't normally do this with guided projects, but I want you to take some time and document your program. Treat it as if you were releasing it for other developers to use.

## Assignment

- Create a README.md file in the root of your repo if you don't have one already (and track your changes with Git).
- Explain to the user that they'll need Postgres and Go installed to run the program.
- Explain to the user how to install the gator CLI using `go install`.
- Explain to the user how to set up the config file and run the program. Tell them about a few of the commands they can run.
- Push gator up to GitHub, then submit the link to your remote repo (e.g., https://github.com/github-username/gator).

**Note:** Go programs are statically compiled binaries. After running `go build` or `go install`, you should be able to run your code without needing the Go toolchain.  
`go run .` is just for development; `gator` is for production.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.16 or later)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/YourUsername/gator.git
cd gator
```

### 2.Install the Gator CLI

```bash
go install ./cmd/gator
```

### 3. Set Up the Configuration File

Gator uses a configuration file to store settings like your PostgreSQL connection string and the current user. By default, it looks for a file named .gatorconfig.json in your home directory.

Create the file in your home directory (e.g., ~/.gatorconfig.json on Linux/macOS or %USERPROFILE%\.gatorconfig.json on Windows) with the following content:

```json
{
  "DBURL": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "CurrentUserName": "your_username"
}
```

Adjust the DBURL and CurrentUserName values as needed for your environment.

## Database Setup

### 1. Create the Database

Create a PostgreSQL database named gator (or modify the connection string in your config file if you use a different name):

```bash

createdb gator

```

### 2. Run Migrations

Gator uses Goose to manage SQL migrations. Install Goose if you haven't already:

```bash

go install github.com/pressly/goose/v3/cmd/goose@latest
```

Then run the migrations located in the sql/migrations directory:

```bash
goose -dir sql/migrations postgres "$DBURL" up
```

This will set up your database schema according to the migration scripts provided in the project.

## Running Gator

### 1. For Development

During development, you can run the application with:

```bash

go run .
```

### 2. For Production

After building or installing the binary, run the app using the gator command:

```bash

gator <command> [arguments]
```

Since Go produces statically compiled binaries, after running go build or go install, you no longer need the Go toolchain to run your program.

### 3. Available Commands

Here are some commands you can run with Gator:

register: Create a new user.

```bash
gator register <username> <password>
```

login: Log in as an existing user.

bash

gator login <username> <password>
addfeed: Add a new RSS feed.
