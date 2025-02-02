# Logging Library

## Overview

This library provides a structured and configurable logging mechanism using [Uber's Zap](https://github.com/uber-go/zap). It supports different log levels, output formats, and includes correlation IDs for enhanced traceability.

## Features

- ✅ Fully configurable via environment variables
- ✅ Supports multiple log levels: `debug`, `info`, `warn`, `error`, `fatal`
- ✅ Supports different log formats: `json` and `console`
- ✅ Singleton logger instance for efficient logging
- ✅ Correlation ID support for better traceability (configurable key)

## Environment Variables

| Variable             | Default           | Description                                                                                                                                              |
| -------------------- | ---------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `LOG_LEVEL`         | `info`            | Defines the severity of log messages. Options: `debug`, `info`, `warn`, `error`, `fatal`. Lower levels (e.g., `debug`) include all higher levels.        |
| `LOG_ENV`           | `prod`            | Specifies the application environment. Options: `prod`, `dev`, `test`. This can help differentiate log sources in monitoring tools.                      |
| `LOG_FORMAT`        | `json`            | Determines the log output format. Options: `json` (structured logs) or `console` (human-readable logs). `json` is preferred for production environments. |
| `LOG_CORRELATION_ID_KEY` | `correlation_id` | Defines the key name used to store and retrieve the correlation ID in the context. This allows customization of the correlation identifier.               |

## Installation

```sh
go get github.com/bylucasqueiroz/loggo
```

## Usage

### Setting Environment Variables

Before running your application, set the desired environment variables:

```sh
export LOG_LEVEL=debug
export LOG_ENV=dev
export LOG_FORMAT=console
export LOG_CORRELATION_ID_KEY=request_id
```

Alternatively, you can use a `.env` file and load it in your application using [github.com/joho/godotenv](https://github.com/joho/godotenv):

1. Install the package:

```sh
go get github.com/joho/godotenv
```

2. Create a `.env` file:

```sh
LOG_LEVEL=debug
LOG_ENV=dev
LOG_FORMAT=console
LOG_CORRELATION_ID_KEY=request_id
```

3. Load it in your application:

```go
package main

import (
	"context"
	"log"
	"github.com/bylucasqueiroz/loggo"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	logger := loggo.GetLogger()
	ctx := loggo.WithCorrelationID(context.Background(), "12345-ABCDE")

	logger.Info(ctx, "Application started")
	logger.Debug(ctx, "Debugging some logic")
	logger.Error(ctx, "An error occurred")
}
```

## Project Structure

```
.
├── config.go       # Handles environment-based configuration
├── context.go      # Provides correlation ID management (configurable key)
├── examples
│   └── main.go    # Example usage of the logging library
├── go.mod          # Go module file
├── go.sum          # Dependency management file
└── logger.go       # Implements the singleton logger
```

## Implementation Details

### Configuration (`config.go`)

- Reads environment variables to configure log level, format, and output.
- Supports JSON and console log formats.

### Logger (`logger.go`)

- Implements a singleton logger instance.
- Provides structured logging with correlation ID support.

### Context Management (`context.go`)

- Handles correlation ID generation and retrieval from context.
- The correlation ID key is now configurable using `LOG_CORRELATION_ID_KEY`, allowing flexibility in log tracking.

## License

MIT License

