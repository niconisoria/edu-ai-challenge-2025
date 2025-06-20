# Service Analyzer

## Prerequisites
- Go 1.21 or higher
- OpenAI API key

## Installation

1. Install Go dependencies:
   ```bash
   go mod download
   ```

2. Set your OpenAI API key:
   ```bash
   export OPENAI_API_KEY=your_api_key_here
   ```

3. Run the application:
   ```bash
   go run .
   ```

## Configuration Files

The application automatically loads these files:
- `system.txt` - System prompt for analysis instructions
- `context.txt` - Additional context for analysis
- `examples.json` - Few-shot learning examples
- `schema.json` - JSON schema defining the structure for service analysis responses

## Usage

- Enter queries when prompted with `>`
- Responses are saved to `response.md`
- Type `quit` or `exit` to close the application
