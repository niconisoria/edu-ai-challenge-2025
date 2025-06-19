# Data Validation Library

## Running the Application

This project is written in Go. To run the application, use:

```
go run application/main.go
```

## Running the Tests

To run all tests in the project, use:

```
go test ./...
```

To run tests for a specific package (e.g., validation):

```
go test ./domain/validation -v
```

## Running Tests with Coverage

To run tests with coverage reporting, use:

```
go test ./... -cover
```

For detailed coverage information with percentage breakdown by function:

```
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

To generate an HTML coverage report for visual inspection:

```
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

To run tests with coverage for a specific package:

```
go test ./domain/validation -cover -v
```

## Requirements
- Go 1.21 or newer

## Project Structure

This project follows a clean architecture pattern with clear separation of concerns:

### `domain/validation/` - Core Validation Logic
The heart of the validation system containing all validator implementations:

- **`validator.go`** - Core validator interface definitions
- **`base_validator.go`** - Base validator implementation with common functionality
- **`validation_error.go`** - Error handling and validation result structures

#### Type-Specific Validators:
- **`string_validator.go`** - String validation with length, pattern, and format checks
- **`number_validator.go`** - Numeric validation with range and type checks
- **`boolean_validator.go`** - Boolean value validation
- **`date_validator.go`** - Date and time validation
- **`array_validator.go`** - Array validation with element type checking
- **`object_validator.go`** - Object validation with field schema definitions

#### Test Files:
Each validator has comprehensive test coverage with corresponding `*_test.go` files containing unit tests and integration tests.

### `infrastructure/schema/` - Schema Factory
Provides a fluent API for building validation schemas:

- **`schema_factory.go`** - Schema builder with methods for creating different validator types
- **`schema_factory_test.go`** - Tests for schema factory functionality

### `application/` - Application Layer
Contains the main application entry point and examples:

- **`main.go`** - Main application with example usage demonstrating complex schema validation

### Root Level Files:
- **`go.mod`** - Go module definition and dependencies
- **`README.md`** - Project documentation
- **`test_report.txt`** - Test execution results

## Architecture Overview

The project implements a validation library with:
- **Domain Layer**: Pure validation logic with no external dependencies
- **Infrastructure Layer**: Schema building utilities and factory patterns
- **Application Layer**: Usage examples and integration demonstrations
