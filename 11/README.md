# Audio Processing Application

This Go application processes audio files through three main steps:

1. **Transcription**: Converts audio to text using OpenAI's Whisper model
2. **Summary**: Generates a concise summary of the transcribed content
3. **Speech Analysis**: Performs structured analysis including word count, speaking speed, and frequently mentioned topics

## Prerequisites

- Go 1.21 or higher
- OpenAI API key

## Setup

1. Set your OpenAI API key as an environment variable:
   ```bash
   export OPENAI_API_KEY="your-api-key-here"
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

Run the application with an audio file as an argument:

```bash
go run . <audio_file_path>
```

Example:
```bash
go run . audio.mp3
```

## Supported Audio Formats

The application supports various audio formats that are compatible with OpenAI's Whisper model, including:
- MP3
- MP4
- Mpeg
- MPGA
- M4A
- WAV
- WEBM

## Output

The application will output:
1. The complete transcription of the audio
2. A summary of the content
3. Structured speech analysis including:
   - Word count
   - Speaking speed (words per minute)
   - Frequently mentioned topics with mention counts

## Dependencies

- `github.com/openai/openai-go v1.6.0` - OpenAI Go SDK
