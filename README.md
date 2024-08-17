# Star Trek Database

This project provides a database for Star Trek series and episodes. It includes functionality to seed the database with data from JSON files.

## Project Structure

```
star-trek-database/
│
├── data/
│   ├── series.json
│   └── episodes/
│       ├── <series-name>-episodes.json
│
├── main.go            # Main application entry point
├── seed.go            # Database seeding logic
├── .env               # Environment variables
├── go.mod             # Go module file
├── go.sum             # Go module checksum file
└── README.md          # Project documentation
```

## Setup

### Prerequisites

- Go (1.22+)
- PostgreSQL

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/yourusername/star-trek-database.git
   cd star-trek-database
   ```

2. **Install dependencies:**

   ```sh
   go mod tidy
   ```

3. **Create a `.env` file:**

   Rename `.env.example` to `.env` and edit it with your PostgreSQL credentials:

   ```env
   DB_HOST=localhost
   DB_USER=yourusername
   DB_PASS=yourpassword
   DB_NAME=yourdbname
   DB_PORT=5432
   ```

4. **Prepare the database:**

   Ensure PostgreSQL is running and the database exists.

### Running the Application

To seed the database, run:

```sh
go run main.go
```

This will:
- Drop existing tables
- Recreate tables
- Seed the database with data from `data/series.json` and `data/episodes/`

## Data Files

- **`data/series.json`**: Series information.
- **`data/episodes/`**: Episode data files for each series.

## License

MIT License.

## Contact

For questions, contact [eric@ericthomas.ca](eric@ericthomas.ca).

---

Feel free to replace placeholders with actual information specific to your project.