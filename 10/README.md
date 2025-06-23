# AI Product Search CLI

A command-line interface application that uses OpenAI's API to search through product catalogs using natural language queries.

## Prerequisites

- Go 1.21 or higher
- OpenAI API key

## Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd <project-directory>
   ```

2. **Set your OpenAI API key**
   ```bash
   export OPENAI_API_KEY="your-openai-api-key-here"
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

## Running the Application

```bash
go run .
```

The application will start an interactive CLI where you can:
- Type natural language queries to search for products
- Type `quit` or `exit` to close the application

## Example Usage

```
Welcome to the AI Product Search CLI Application!
Type 'quit' or 'exit' to leave the application.
--------------------------------------------------
> Find me wireless headphones under $100
> Show me kitchen appliances for baking
> I need a fitness tracker with heart rate monitoring
```
