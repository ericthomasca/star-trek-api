# Star Trek Database

[![Go Reference](https://pkg.go.dev/badge/github.com/ericthomasca/star-trek-api.svg)](https://pkg.go.dev/github.com/ericthomasca/star-trek-api)

This project provides a database for Star Trek series and episodes. It includes functionality to seed the database with data from JSON files.

## Project Structure

```
star-trek-api/
│
├── config/
│ └── config.go # Configuration settings
│
├── data/
│ ├── data.go # General data handling
│ ├── ds9.go # Data for Deep Space Nine
│ ├── ent.go # Data for Enterprise
│ ├── tas.go # Data for The Animated Series
│ ├── tng.go # Data for The Next Generation
│ ├── tos.go # Data for The Original Series
│ └── voy.go # Data for Voyager
│
├── models/
│ └── models.go # Database models
|
├── seed/
│ └── seed.go # Database seeding logic
|
├── go.mod # Go module file
├── go.sum # Go module checksum file
├── LICENSE # Project license
├── main.go # Main application entry point
├── README.md # Project documentation
└── star-trek-api # Executable binary (if built)
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
- Seed the database with data from `data/` files

## Data Files

- **`data/data.go`**: General data handling functions.
- **`data/ds9.go`**: Data for Deep Space Nine.
- **`data/ent.go`**: Data for Enterprise.
- **`data/tas.go`**: Data for The Animated Series.
- **`data/tng.go`**: Data for The Next Generation.
- **`data/tos.go`**: Data for The Original Series.
- **`data/voy.go`**: Data for Voyager.

## License

MIT License.

## Contact

For questions, contact [eric@ericthomas.ca](eric@ericthomas.ca).

---

Feel free to replace placeholders with actual information specific to your project.
