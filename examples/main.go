package main

import (
	"fmt"
	"strings"

	"github.com/slashdevops/textn8r"
)

func main() {
	fmt.Println("=== textn8r Examples ===")
	fmt.Println()

	// Basic normalizers
	fmt.Println("1. Basic Case Normalizers:")
	text := "  Hello, World!  "
	fmt.Printf("Original: '%s'\n", text)
	fmt.Printf("Uppercase: '%s'\n", textn8r.UpperCaseNormalizer(text))
	fmt.Printf("Lowercase: '%s'\n", textn8r.LowerCaseNormalizer(text))
	fmt.Printf("Trimmed: '%s'\n", textn8r.TrimSpaceNormalizer(text))
	fmt.Println()

	// Space normalization
	fmt.Println("2. Space Normalizers:")
	spaceText := "  Hello    world   with   extra    spaces  "
	fmt.Printf("Original: '%s'\n", spaceText)
	fmt.Printf("Remove extra spaces: '%s'\n", textn8r.RemoveExtraSpaceNormalizer(spaceText))
	fmt.Printf("Remove all spaces: '%s'\n", textn8r.RemoveAllSpaceNormalizer(spaceText))
	fmt.Println()

	// Character removal
	fmt.Println("3. Character Removal Normalizers:")
	specialText := "Hello@#$%World123!?*"
	fmt.Printf("Original: '%s'\n", specialText)
	fmt.Printf("Remove special chars: '%s'\n", textn8r.RemoveSpecialCharactersNormalizer(specialText))
	fmt.Printf("Remove digits: '%s'\n", textn8r.RemoveDigitsNormalizer(specialText))
	fmt.Printf("Remove punctuation: '%s'\n", textn8r.RemovePunctuationNormalizer(specialText))
	fmt.Printf("Remove non-alphanumeric: '%s'\n", textn8r.RemoveNonAlphanumericNormalizer(specialText))
	fmt.Println()

	// Accent and diacritic handling
	fmt.Println("4. Accent and Diacritic Normalizers:")
	accentText := "Café, résumé, naïve, jalapeño, niño"
	fmt.Printf("Original: '%s'\n", accentText)
	fmt.Printf("Remove accents: '%s'\n", textn8r.ReplaceAccentsNormalizer(accentText))
	fmt.Printf("Remove tildes: '%s'\n", textn8r.ReplaceTildesNormalizer(accentText))
	fmt.Println()

	// Whitespace handling
	fmt.Println("5. Whitespace Normalizers:")
	whitespaceText := "Line 1\nLine 2\tTabbed\rCarriage Return"
	fmt.Printf("Original: '%s'\n", whitespaceText)
	fmt.Printf("Remove newlines: '%s'\n", textn8r.RemoveNewLineNormalizer(whitespaceText))
	fmt.Printf("Remove tabs: '%s'\n", textn8r.RemoveTabNormalizer(whitespaceText))
	fmt.Printf("Remove carriage returns: '%s'\n", textn8r.RemoveCarriageReturnNormalizer(whitespaceText))
	fmt.Println()

	// Replacement normalizers
	fmt.Println("6. Replacement Normalizers:")
	replaceText := "Hello\tWorld\nNew Line\rCarriage 123!@#"
	fmt.Printf("Original: '%s'\n", replaceText)

	tabReplacer := textn8r.ReplaceTabNormalizer("-")
	fmt.Printf("Replace tabs with '-': '%s'\n", tabReplacer.Apply(replaceText))

	newlineReplacer := textn8r.ReplaceNewLineNormalizer(" | ")
	fmt.Printf("Replace newlines with ' | ': '%s'\n", newlineReplacer.Apply(replaceText))

	digitReplacer := textn8r.ReplaceDigitsNormalizer("X")
	fmt.Printf("Replace digits with 'X': '%s'\n", digitReplacer.Apply(replaceText))

	spaceReplacer := textn8r.ReplaceSpaceNormalizer("_")
	fmt.Printf("Replace spaces with '_': '%s'\n", spaceReplacer.Apply("Hello World Test"))
	fmt.Println()

	// Chaining normalizers
	fmt.Println("7. Chaining Multiple Normalizers:")
	chainText := "  CAFÉ, résumé! @2023  "
	fmt.Printf("Original: '%s'\n", chainText)

	normalizers := textn8r.Normalizers{
		textn8r.TrimSpaceNormalizer,
		textn8r.LowerCaseNormalizer,
		textn8r.ReplaceAccentsNormalizer,
		textn8r.RemoveSpecialCharactersNormalizer,
		textn8r.RemoveExtraSpaceNormalizer,
	}

	result := normalizers.Apply(chainText)
	fmt.Printf("After chaining normalizers: '%s'\n", result)
	fmt.Println()

	// Custom normalizer example
	fmt.Println("8. Custom Normalizer:")
	customText := "foo is everywhere, foo here, foo there"
	fmt.Printf("Original: '%s'\n", customText)

	customNormalizer := textn8r.Normalizer(func(input string) string {
		return strings.ReplaceAll(input, "foo", "bar")
	})

	fmt.Printf("Replace 'foo' with 'bar': '%s'\n", customNormalizer.Apply(customText))
	fmt.Println()

	// Real-world example: Clean user input
	fmt.Println("9. Real-world Example - Clean User Input:")
	userInput := "  JoHn.DoE@123!  \t\n"
	fmt.Printf("User input: '%s'\n", userInput)

	cleaningNormalizers := textn8r.Normalizers{
		textn8r.TrimSpaceNormalizer,
		textn8r.RemoveTabNormalizer,
		textn8r.RemoveNewLineNormalizer,
		textn8r.LowerCaseNormalizer,
		textn8r.RemoveSpecialCharactersNormalizer,
		textn8r.RemoveExtraSpaceNormalizer,
	}

	cleanedInput := cleaningNormalizers.Apply(userInput)
	fmt.Printf("Cleaned input: '%s'\n", cleanedInput)
	fmt.Println()

	// URL slug example
	fmt.Println("10. URL Slug Example:")
	title := "How to Create Amazing Web Apps in 2023!"
	fmt.Printf("Title: '%s'\n", title)

	slugNormalizers := textn8r.Normalizers{
		textn8r.TrimSpaceNormalizer,
		textn8r.LowerCaseNormalizer,
		textn8r.RemovePunctuationNormalizer,
		textn8r.RemoveExtraSpaceNormalizer,
		textn8r.ReplaceSpaceNormalizer("-"),
	}

	slug := slugNormalizers.Apply(title)
	fmt.Printf("URL slug: '%s'\n", slug)
}
