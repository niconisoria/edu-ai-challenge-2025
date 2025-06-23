package main

import (
	"github.com/openai/openai-go"
)

// Topic represents a frequently mentioned topic with its mention count
type Topic struct {
	Topic    string `json:"topic"`
	Mentions int    `json:"mentions"`
}

// SpeechAnalysisResponse represents the structured response format for speech analysis
type SpeechAnalysisResponse struct {
	WordCount                 int     `json:"word_count"`
	SpeakingSpeedWPM          int     `json:"speaking_speed_wpm"`
	FrequentlyMentionedTopics []Topic `json:"frequently_mentioned_topics"`
}

// SpeechAnalysisSchema defines the JSON schema for speech analysis responses
var SpeechAnalysisSchema = openai.ResponseFormatJSONSchemaJSONSchemaParam{
	Name: "SpeechAnalysis",
	Schema: map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"word_count": map[string]interface{}{
				"type":        "integer",
				"description": "The total number of words in the speech",
			},
			"speaking_speed_wpm": map[string]interface{}{
				"type":        "integer",
				"description": "The speaking speed in words per minute",
			},
			"frequently_mentioned_topics": map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"topic": map[string]interface{}{
							"type":        "string",
							"description": "The name of the topic",
						},
						"mentions": map[string]interface{}{
							"type":        "integer",
							"description": "The number of times this topic was mentioned",
						},
					},
					"required": []string{"topic", "mentions"},
				},
				"description": "Array of topics with their mention counts",
			},
		},
		"required": []string{"word_count", "speaking_speed_wpm", "frequently_mentioned_topics"},
	},
}

// GetSpeechAnalysisSchema returns the speech analysis schema
func GetSpeechAnalysisSchema() openai.ResponseFormatJSONSchemaJSONSchemaParam {
	return SpeechAnalysisSchema
}
