package textn8r_test

import (
	"fmt"
	"strings"

	"github.com/slashdevops/textn8r"
)

// Example demonstrates basic usage of individual normalizers
func ExampleUpperCaseNormalizer() {
	result := textn8r.UpperCaseNormalizer("hello world")
	fmt.Println(result)

	// Output:
	// HELLO WORLD
}

// Example demonstrates trim space normalizer
func ExampleTrimSpaceNormalizer() {
	result := textn8r.TrimSpaceNormalizer("  hello world  ")
	fmt.Println(result)

	// Output:
	// hello world
} // Example demonstrates chaining multiple normalizers
func ExampleNormalizers_Apply() {
	text := "  CAFÉ, résumé! @2023  "

	normalizers := textn8r.Normalizers{
		textn8r.TrimSpaceNormalizer,
		textn8r.LowerCaseNormalizer,
		textn8r.ReplaceAccentsNormalizer,
		textn8r.RemoveSpecialCharactersNormalizer,
		textn8r.RemoveExtraSpaceNormalizer,
	}

	result := normalizers.Apply(text)
	fmt.Println(result)

	// Output:
	// cafe resume 2023
}

// Example demonstrates replacement normalizers
func ExampleReplaceTabNormalizer() {
	text := "Hello\tWorld\tTest"

	tabReplacer := textn8r.ReplaceTabNormalizer("-")
	result := tabReplacer.Apply(text)
	fmt.Println(result)

	// Output:
	// Hello-World-Test
}

// Example demonstrates creating a custom normalizer
func ExampleNormalizer_custom() {
	text := "foo is everywhere, foo here, foo there"

	customNormalizer := textn8r.Normalizer(func(input string) string {
		return strings.ReplaceAll(input, "foo", "bar")
	})

	result := customNormalizer.Apply(text)
	fmt.Println(result)

	// Output:
	// bar is everywhere, bar here, bar there
}

// Example demonstrates accent handling
func ExampleReplaceAccentsNormalizer() {
	text := "Café, résumé, naïve, jalapeño"

	result := textn8r.ReplaceAccentsNormalizer(text)
	fmt.Println(result)

	// Output:
	// Cafe, resume, naive, jalapeno
}

// Example demonstrates creating URL slugs
func Example_urlSlug() {
	title := "How to Create Amazing Web Apps in 2023!"

	slugNormalizers := textn8r.Normalizers{
		textn8r.TrimSpaceNormalizer,
		textn8r.LowerCaseNormalizer,
		textn8r.RemovePunctuationNormalizer,
		textn8r.RemoveExtraSpaceNormalizer,
		textn8r.ReplaceSpaceNormalizer("-"),
	}

	slug := slugNormalizers.Apply(title)
	fmt.Println(slug)

	// Output:
	// how-to-create-amazing-web-apps-in-2023
}
