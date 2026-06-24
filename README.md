# textn8r

[![Go Report Card](https://goreportcard.com/badge/github.com/slashdevops/textn8r)](https://goreportcard.com/report/github.com/slashdevops/textn8r)
[![GoDoc](https://godoc.org/github.com/slashdevops/textn8r?status.svg)](https://godoc.org/github.com/slashdevops/textn8r)

A flexible and extensible Go library for text normalization. `textn8r` provides a comprehensive set of string normalizers that can be used individually or chained together to clean, transform, and standardize text data.

## Features

- **Case Conversion**: Convert text to uppercase, lowercase
- **Space Handling**: Trim, remove extra spaces, or remove all spaces
- **Character Removal**: Remove special characters, digits, punctuation, or non-alphanumeric characters
- **Accent & Diacritic Handling**: Remove or replace accented characters and diacritics
- **Whitespace Normalization**: Handle tabs, newlines, and carriage returns
- **Flexible Replacement**: Replace specific character types with custom strings
- **Chainable Operations**: Combine multiple normalizers for complex transformations
- **Custom Normalizers**: Create your own normalizers easily

## Installation

```bash
go get github.com/slashdevops/textn8r
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/slashdevops/textn8r"
)

func main() {
    // Basic usage
    text := "  Hello, World!  "
    result := textn8r.UpperCaseNormalizer(text)
    fmt.Println(result) // "  HELLO, WORLD!  "

    // Chain multiple normalizers
    normalizers := textn8r.Normalizers{
        textn8r.TrimSpaceNormalizer,
        textn8r.UpperCaseNormalizer,
        textn8r.RemoveSpecialCharactersNormalizer,
    }
    result = normalizers.Apply("  hello, world!  ")
    fmt.Println(result) // "HELLO  WORLD "
}
```

## Available Normalizers

### Case Conversion

- `UpperCaseNormalizer`: Converts text to uppercase
- `LowerCaseNormalizer`: Converts text to lowercase

### Space Handling

- `TrimSpaceNormalizer`: Removes leading and trailing whitespace
- `RemoveExtraSpaceNormalizer`: Collapses multiple spaces into single spaces
- `RemoveAllSpaceNormalizer`: Removes all spaces

### Character Removal

- `RemoveSpecialCharactersNormalizer`: Removes special characters (keeps alphanumeric and spaces)
- `RemoveDigitsNormalizer`: Removes all numeric digits
- `RemovePunctuationNormalizer`: Removes punctuation marks
- `RemoveNonAlphanumericNormalizer`: Removes all non-alphanumeric characters

### Whitespace Control

- `RemoveTabNormalizer`: Removes tab characters
- `RemoveNewLineNormalizer`: Removes newline characters
- `RemoveCarriageReturnNormalizer`: Removes carriage return characters

### Accent & Diacritic Handling

- `ReplaceAccentsNormalizer`: Replaces accented characters with base characters (café → cafe)
- `ReplaceTildesNormalizer`: Replaces tilde characters (ñ → n)
- `RemoveDiacriticsNormalizer`: Removes all diacritical marks
- `RemoveTildesNormalizer`: Removes tilde characters

### Replacement Normalizers

These normalizers replace characters with custom strings:

- `ReplaceTabNormalizer(replacement)`: Replaces tabs with specified string
- `ReplaceNewLineNormalizer(replacement)`: Replaces newlines with specified string
- `ReplaceCarriageReturnNormalizer(replacement)`: Replaces carriage returns
- `ReplaceSpaceNormalizer(replacement)`: Replaces spaces
- `ReplaceDigitsNormalizer(replacement)`: Replaces digits
- `ReplacePunctuationNormalizer(replacement)`: Replaces punctuation
- `ReplaceNonAlphanumericNormalizer(replacement)`: Replaces non-alphanumeric characters
- `ReplaceSpecialCharactersNormalizer(input, replacement)`: Replaces special characters
- `ReplaceDiacriticsNormalizer(replacement)`: Replaces diacritics

## Usage Examples

### Basic Normalizers

```go
text := "  Hello, World!  "

// Case conversion
fmt.Println(textn8r.UpperCaseNormalizer(text))     // "  HELLO, WORLD!  "
fmt.Println(textn8r.LowerCaseNormalizer(text))     // "  hello, world!  "

// Space handling
fmt.Println(textn8r.TrimSpaceNormalizer(text))     // "Hello, World!"
```

### Basic Character Removal Examples

```go
text := "Hello@#$%World123!?*"

fmt.Println(textn8r.RemoveSpecialCharactersNormalizer(text))    // "Hello World123 "
fmt.Println(textn8r.RemoveDigitsNormalizer(text))              // "Hello@#$%World!?*"
fmt.Println(textn8r.RemovePunctuationNormalizer(text))         // "Hello@#$%World123"
fmt.Println(textn8r.RemoveNonAlphanumericNormalizer(text))     // "HelloWorld123"
```

### Accent Handling

```go
text := "Café, résumé, naïve, jalapeño"

fmt.Println(textn8r.ReplaceAccentsNormalizer(text))  // "Cafe, resume, naive, jalapeno"
fmt.Println(textn8r.ReplaceTildesNormalizer("niño")) // "nino"
```

### Using Replacement Normalizers

```go
text := "Hello\tWorld\nNew Line"

tabReplacer := textn8r.ReplaceTabNormalizer("-")
fmt.Println(tabReplacer.Apply(text))  // "Hello-World\nNew Line"

newlineReplacer := textn8r.ReplaceNewLineNormalizer(" | ")
fmt.Println(newlineReplacer.Apply(text))  // "Hello\tWorld | New Line"
```

### Chaining Normalizers

```go
text := "  CAFÉ, résumé! @2023  "

normalizers := textn8r.Normalizers{
    textn8r.TrimSpaceNormalizer,
    textn8r.LowerCaseNormalizer,
    textn8r.ReplaceAccentsNormalizer,
    textn8r.RemoveSpecialCharactersNormalizer,
    textn8r.RemoveExtraSpaceNormalizer,
}

result := normalizers.Apply(text)
fmt.Println(result)  // "cafe resume 2023"
```

### Custom Normalizers

```go
// Create a custom normalizer
customNormalizer := textn8r.Normalizer(func(input string) string {
    return strings.ReplaceAll(input, "foo", "bar")
})

result := customNormalizer.Apply("foo is everywhere")
fmt.Println(result)  // "bar is everywhere"
```

## Real-World Use Cases

### URL Slug Generation

```go
title := "How to Create Amazing Web Apps in 2023!"

slugNormalizers := textn8r.Normalizers{
    textn8r.TrimSpaceNormalizer,
    textn8r.LowerCaseNormalizer,
    textn8r.RemovePunctuationNormalizer,
    textn8r.RemoveExtraSpaceNormalizer,
    textn8r.ReplaceSpaceNormalizer("-"),
}

slug := slugNormalizers.Apply(title)
fmt.Println(slug)  // "how-to-create-amazing-web-apps-in-2023"
```

### User Input Sanitization

```go
userInput := "  JoHn.DoE@123!  \t\n"

cleaningNormalizers := textn8r.Normalizers{
    textn8r.TrimSpaceNormalizer,
    textn8r.RemoveTabNormalizer,
    textn8r.RemoveNewLineNormalizer,
    textn8r.LowerCaseNormalizer,
    textn8r.RemoveSpecialCharactersNormalizer,
    textn8r.RemoveExtraSpaceNormalizer,
}

cleanedInput := cleaningNormalizers.Apply(userInput)
fmt.Println(cleanedInput)  // "john doe 123"
```

### Data Standardization

```go
// Standardize company names
companyNames := []string{
    "  ACME Corp.  ",
    "acme corporation",
    "A.C.M.E. Inc",
}

standardizer := textn8r.Normalizers{
    textn8r.TrimSpaceNormalizer,
    textn8r.LowerCaseNormalizer,
    textn8r.RemovePunctuationNormalizer,
    textn8r.RemoveExtraSpaceNormalizer,
}

for _, name := range companyNames {
    standardized := standardizer.Apply(name)
    fmt.Printf("'%s' -> '%s'\n", name, standardized)
}
// Output:
// '  ACME Corp.  ' -> 'acme corp'
// 'acme corporation' -> 'acme corporation'
// 'A.C.M.E. Inc' -> 'acme inc'
```

## Running the Examples

You can run the comprehensive examples included in this repository:

```bash
cd examples
go run main.go
```

## Testing

Run the test suite:

```bash
go test -v
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
