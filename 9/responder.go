package main

import (
	"fmt"
	"os"
	"strings"
)

// getReportSections returns the structured sections of a ServiceReport
func getReportSections(report *ServiceReport) []struct {
	title   string
	content string
} {
	return []struct {
		title   string
		content string
	}{
		{"Brief History", report.BriefHistory},
		{"Target Audience", report.TargetAudience},
		{"Core Features", report.CoreFeatures},
		{"Unique Selling Points", report.UniqueSellingPoints},
		{"Business Model", report.BusinessModel},
		{"Tech Stack Insights", report.TechStackInsights},
		{"Perceived Strengths", report.PerceivedStrengths},
		{"Perceived Weaknesses", report.PerceivedWeaknesses},
	}
}

// WriteResponseToFile writes a ServiceReport to file in markdown format
func WriteResponseToFile(input string, report *ServiceReport) error {
	// Open file in append mode, create if doesn't exist
	file, err := os.OpenFile("response.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Build markdown content using a more structured approach
	var markdownBuilder strings.Builder

	// Header
	markdownBuilder.WriteString("## Service Analysis Report\n\n")

	// Input section
	markdownBuilder.WriteString("**Input:** " + input + "\n\n")

	// Report sections
	sections := getReportSections(report)
	for _, section := range sections {
		markdownBuilder.WriteString("**" + section.title + ":**\n")
		markdownBuilder.WriteString(section.content + "\n\n")
	}

	// Separator
	markdownBuilder.WriteString("---\n\n")

	// Write to file
	_, err = file.WriteString(markdownBuilder.String())
	return err
}

// WriteResponseToConsole writes a ServiceReport to console in a formatted way
func WriteResponseToConsole(input string, report *ServiceReport) {
	fmt.Printf("\nInput: %s\n\n", input)

	// Report sections
	sections := getReportSections(report)
	for _, section := range sections {
		fmt.Printf("%s:\n%s\n\n", section.title, section.content)
	}
}
