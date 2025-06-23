package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// Check if audio file path is provided as command line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <audio_file_path>")
		fmt.Println("Example: go run . audio.mp3")
		os.Exit(1)
	}

	audioFilePath := os.Args[1]

	// Step 1: Generate transcription from audio file using Whisper model
	fmt.Println("Step 1: Generating transcription...")
	transcription, err := generateTranscription(audioFilePath)
	if err != nil {
		log.Fatalf("Failed to generate transcription: %v", err)
	}

	// Write transcription to file
	err = writeToFile("transcription.md", transcription)
	if err != nil {
		log.Printf("Warning: Failed to write transcription to file: %v", err)
	} else {
		fmt.Println("Transcription saved to transcription.md")
	}

	// Step 2: Generate summary using chat completion
	fmt.Println("Step 2: Generating summary...")
	summary, err := generateSummary(transcription)
	if err != nil {
		log.Fatalf("Failed to generate summary: %v", err)
	}

	// Write summary to file
	err = writeToFile("summary.md", summary)
	if err != nil {
		log.Printf("Warning: Failed to write summary to file: %v", err)
	} else {
		fmt.Println("Summary saved to summary.md")
	}

	// Step 3: Generate speech analysis using the schema
	fmt.Println("Step 3: Generating speech analysis...")
	analysis, err := generateSpeechAnalysis(transcription)
	if err != nil {
		log.Fatalf("Failed to generate speech analysis: %v", err)
	}

	// Write analysis to JSON file
	err = writeAnalysisToJSON("analysis.json", analysis)
	if err != nil {
		log.Printf("Warning: Failed to write analysis to file: %v", err)
	} else {
		fmt.Println("Analysis saved to analysis.json")
	}
}

// writeToFile writes content to a file with the given filename
func writeToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

// writeAnalysisToJSON writes the analysis struct to a JSON file
func writeAnalysisToJSON(filename string, analysis *SpeechAnalysisResponse) error {
	jsonData, err := json.MarshalIndent(analysis, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal analysis to JSON: %w", err)
	}
	return os.WriteFile(filename, jsonData, 0644)
}

// generateTranscription creates a transcription from an audio file using Whisper
func generateTranscription(audioFilePath string) (string, error) {
	// Create Whisper client
	whisperClient, err := NewClient("whisper-1")
	if err != nil {
		return "", fmt.Errorf("failed to create Whisper client: %w", err)
	}
	defer whisperClient.Close()

	// Create transcription with options
	transcription, err := whisperClient.CreateAudioTranscriptionWithOptions(
		audioFilePath,
		"en", // language
		"This is a speech that may contain technical terms and proper nouns.", // prompt
	)
	if err != nil {
		return "", fmt.Errorf("failed to create transcription: %w", err)
	}

	return transcription.Text, nil
}

// generateSummary creates a summary of the transcription using chat completion
func generateSummary(transcription string) (string, error) {
	// Create GPT-4 client for summary generation
	gptClient, err := NewClient("gpt-4.1-mini")
	if err != nil {
		return "", fmt.Errorf("failed to create GPT client: %w", err)
	}
	defer gptClient.Close()

	systemMessage := `
		You are a helpful assistant that creates concise and informative summaries of transcribed speech. 
		Focus on the main points, key topics, and important insights from the content. 
		Keep the summary clear and well-structured.
	`

	userMessage := fmt.Sprintf("Provide a comprehensive summary of the following transcription:\n\n%s", transcription)

	// Create chat completion
	response, err := gptClient.CreateChatCompletion(systemMessage, userMessage)
	if err != nil {
		return "", fmt.Errorf("failed to create chat completion: %w", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response choices received")
	}

	return response.Choices[0].Message.Content, nil
}

// generateSpeechAnalysis creates a structured speech analysis using the schema
func generateSpeechAnalysis(transcription string) (*SpeechAnalysisResponse, error) {
	// Create GPT-4 client for speech analysis
	gptClient, err := NewClient("gpt-4.1-mini")
	if err != nil {
		return nil, fmt.Errorf("failed to create GPT client: %w", err)
	}
	defer gptClient.Close()

	systemMessage := `
		You are a speech analysis expert. Analyze the provided transcription and return structured data including:
			1. word count,
			2. speaking speed (WPM),
			3. top 3+ frequently mentioned topics (ordered by number of mentions)
		Be thorough in identifying topics and counting their mentions. Consider synonyms and related terms when grouping topics.
	`

	userMessage := fmt.Sprintf("Please analyze the following speech transcription:\n\n%s", transcription)

	// Get the speech analysis schema
	schema := GetSpeechAnalysisSchema()

	// Create chat completion with structured output
	response, err := gptClient.CreateChatCompletionWithStructuredOutput(systemMessage, userMessage, schema)
	if err != nil {
		return nil, fmt.Errorf("failed to create structured chat completion: %w", err)
	}

	if len(response.Choices) == 0 {
		return nil, fmt.Errorf("no response choices received")
	}

	// Parse the structured response
	var analysis SpeechAnalysisResponse
	content := response.Choices[0].Message.Content

	err = json.Unmarshal([]byte(content), &analysis)
	if err != nil {
		return nil, fmt.Errorf("failed to parse structured response: %w", err)
	}

	return &analysis, nil
}
