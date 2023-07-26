# Wordgame API

Wordgame API is a simple guessing game built in Go. This service exposes an HTTP API for starting a new game and making guesses.

## The game

When a new game begins, the server chooses a random word from the list of words and sets the guesses remaining to 6. Let's say the server has chosen `APPLE`. The server will return `_____` (6 guesses remaining) to the player, indicating that the chosen word is 5 characters long. Now play progresses as follows:

1. The player guesses a character that they believe to be in the chosen word.

2. a) The player's guess matches a letter in the chosen word: The server returns a new string revealing the location(s) of that character.

   b) The player's guess does not match a letter in the chosen word: The server returns the existing string, and decrements the remaining guesses.

In the above example, imagine the player guesses `P`, now the server returns `_PP__` (6 guesses remaining). Now the player guesses `I`, the server returns `_PP__` (5 guesses remaining).

The game ends either when there are no guesses remaining (player loses), or the player has guessed all of the characters in the word (player wins).

## Project Structure

This project follows standard Go project layout conventions. All the application logic resides inside the `internal` directory. The `store` package contains the definition for the game data store interface and its implementations (e.g., `redis` package for Redis-based storage).
The `api` package has handlers for API routes and their respective request/response definitions.

The project folder structure is organized as follows:

- `assets/`: Directory containing game assets such as the word list (`words.txt`).
- `bin/`: Directory where the compiled application binaries are stored.
- `cmd/`: Contains the main application (`main.go`) and configuration code (`config/config.go`).
- `docker/`: Contains Docker files for building and running the application in a containerized environment.
- `docs/`: Contains documentation for the application, including Swagger UI files for API documentation.
- `internal/`: Contains the main business logic of the application, organized by different domains.
  - `api/`: API handlers, request, and response definitions.
  - `game/`: Game-related logic, includes mock objects for testing.
  - `models/`: Data models used by the application.
  - `server/`: Server initialization and configuration code.
  - `store/`: This directory contains all the logic related to data storage of the game state. This ensures separation of concerns where the storage logic is isolated from the rest of the application. It includes:
    - `game_store.go`: This file defines the `GameStore` interface. The `GameStore` interface defines the methods that any data storage should implement in order to be used by this application. This design is very flexible and allows the application to switch to a different database or storage system easily, only requiring the implementation of this interface. The methods include operations to save and load game instances.
    - `redis/`: This directory contains the implementation of the `GameStore` interface for Redis.
- `pkg/`: Libraries and packages that could be used by other services.
  - `identifier/`: Unique identifier generation code.
  - `words/`: Code related to word handling and processing.
- `go.mod` and `go.sum`: Go module files.
- `Makefile`: Contains commands for building, testing, and running the application.
- `config.local.yml`: Configuration file with local settings.

## Prerequisites

- Go 1.20 or later
- Docker and Docker Compose
- Make

## Setup

1. Start Redis using Docker Compose:

    ```bash
    make run-local-dept
    ```

   This will start a Redis instance which will be used as a store for the game data.

2. Build the application:

    ```bash
    make build
    ```

3. Run the application locally:

    ```bash
    make run
    ```

## API Documentation

API documentation is available via Swagger UI and is served at the `/docs/` endpoint when the application is running.

## Game Rules

You can start a new game via a POST request to `/new`. A game is identified by a UUID, a word to guess, and a masked version of this word. When a game starts, the word to guess is replaced by underscores and is revealed partially according to each correct guess.

You can make a guess in a game via a POST request to `/guess`, providing an ID and a guessed letter.

The game ends when all letters have been guessed or the guesses chances are over, the default limit of guesses is 6.


## Testing and Quality Assurance

In this project, we value the importance of testing as part of the software development process. As a result, every key component inside the `/internal` directory has corresponding test files to ensure the functionality and integrity of the codebase.

- Every component in the `/internal` directory has corresponding `_test.go` files. This means that all functionalities implemented in the project, including the API handlers, game logic, server setup, and storage interfaces, are covered by unit tests.

- These tests can be easily run using the command `make test`, facilitating regular checks on the status and health of the codebase during the development process.

    ```bash
    make test
    ```

Apart from unit tests, this project also makes use of Go's in-built tools for maintaining a high standard of code quality:

- `go fmt`: Is used for automatically formatting our Go code in a standard way. It helps keep the code consistent and readable.

    ```bash
    make fmt
    ```

- `go vet`: Is used for analyzing our Go source code and reporting suspicious constructs. It helps catch issues that may not have been caught by the compiler.

    ```bash
    make vet
    ```

- `go lint`: Is a linter for Go source code. It flags style mistakes and programming errors that legal Go code can have.

    ```bash
    make lint
    ```

## Documentation

This project documentations uses the OpenAPI 2.0 specification (swagger).

Notations/comments directly in the source code are used to describe the endpoints, requests, and responses and models of the API.

These notations are then interpreted by the swagger to generate the OpenAPI 2.0 specification as a `swagger.json` file.

This helps to ensure that the documentation stays synchronized with the source code.

The generated `swagger.json` is served via a user-friendly UI on the `/docs` route.

To generate the latest API documentation, you can use the following command:

 ```bash
    make generate-swagger
   ```

## Improvements

Here are some areas where you could consider making improvements:

1. **Word list asset loading**: Instead of loading words from a file into memory, consider loading them into a database like Redis, fetching a word when needed to reduce memory usage.

2. **Error Handling**: The application currently returns generic error messages. You could improve this by returning more detailed error messages and HTTP status codes.

3. **Input Validation**: The application currently does not validate the input data. You could add validation checks to ensure that the input data is valid.

4. **Security**: The application currently does not have any security measures in place. You could add authentication and authorization to the API endpoints.

5. **Automated Testing**: While your project already has unit tests, you can also introduce integration and end-to-end tests for a more holistic test coverage.

6. **Continuous Integration/Continuous Deployment (CI/CD)**: If not already in place, you could setup CI/CD pipelines to automatically build, test, and deploy your application whenever changes are pushed to the codebase.

7. **Environment Variables**: Use environment variables for sensitive data or data that changes between deployment environments. This includes database credentials, API keys, or any configuration options you want to tune between environments.

8. **Logging and Monitoring**: Consider implementing more detailed logging within the application. This could help debug any issues that arise in production. Tools like Prometheus or Grafana can help with monitoring your application's health and performance.

9. **Swagger Annotations**: The swagger specification is currently generated from the source code. Some models are not being correctly interpreted by the swagger generator.


## License

This project is licensed under the MIT License - see the `LICENSE.md` file for details.
