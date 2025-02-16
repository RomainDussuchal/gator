# Gator

Gator is a command-line application written in Go that aggregates and manages RSS feeds. It allows you to register users, add new feeds, and schedule feed aggregation, all from your terminal. This project is perfect for anyone who wants a lightweight, customizable feed aggregator.

## Features

- **User Management:** Register, log in, and list users.
- **Feed Management:** Add new RSS feeds and view posts from your feeds.
- **Aggregation:** Schedule and run feed aggregation at custom intervals.
- **Database Migrations:** Use Goose for easy database schema management.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.16 or later)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Goose](https://github.com/pressly/goose) (for running SQL migrations)

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/YourUsername/gator.git
cd gator
```
