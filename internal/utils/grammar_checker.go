package utils

import (
	"strings"

	"github.com/jdkato/prose/v2"
)

// GrammarIssue represents a grammar issue in the text.
type GrammarIssue struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// CheckGrammar checks the grammar of the given text and returns a list of issues.
func CheckGrammar(text string) ([]GrammarIssue, error) {
	// Create a new document with the given text
	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, err
	}

	// Iterate over the sentences and check for grammar issues
	var issues []GrammarIssue
	for _, sent := range doc.Sentences() {
		// Check for long sentences
		if len(sent.Text) > 100 {
			issues = append(issues, GrammarIssue{
				Type:    "Long Sentence",
				Message: "Sentence is too long: " + sent.Text,
			})
		}

		// Check for passive voice
		for _, tok := range doc.Tokens() {
			if tok.Tag == "VBN" && strings.Contains(sent.Text, tok.Text) {
				issues = append(issues, GrammarIssue{
					Type:    "Passive Voice",
					Message: "Passive voice detected: " + sent.Text,
				})
				break
			}
		}

		// Check for repeated words
		words := strings.Fields(sent.Text)
		wordCount := make(map[string]int)
		for _, word := range words {
			wordCount[strings.ToLower(word)]++
			if wordCount[strings.ToLower(word)] > 1 {
				issues = append(issues, GrammarIssue{
					Type:    "Repeated Word",
					Message: "Repeated word detected: " + word,
				})
				break
			}
		}
	}

	// Implementation of more grammar checks can be added here

	return issues, nil
}
