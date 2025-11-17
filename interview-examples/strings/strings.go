package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// String immutability - strings cannot be modified
	// str2[0] = 'h' // This would cause a compile error
	fmt.Println("\n--- String Immutability ---")
	original := "Hello"
	// To "modify" a string, you must create a new one
	modified := "h" + original[1:]
	fmt.Printf("Original: %s, Modified: %s\n", original, modified)

	// String concatenation
	fmt.Println("\n--- String Concatenation ---")
	firstName := "John"
	lastName := "Doe"
	// Using strings.Builder for efficient concatenation
	var builder strings.Builder
	builder.WriteString(firstName)
	builder.WriteString(" ")
	builder.WriteString(lastName)
	fmt.Printf("Builder result: %s\n", builder.String())

	// String indexing - returns byte values, not characters
	fmt.Println("\n--- String Indexing (Bytes) ---")
	text := "Hello"
	fmt.Printf("First byte: %v (character: %c)\n", text[0], text[0])
	fmt.Printf("Length in bytes: %d\n", len(text))

	// UTF-8 encoding and multi-byte characters
	fmt.Println("\n--- UTF-8 and Multi-byte Characters ---")
	emoji := "Hello, 世界! 🌍"
	fmt.Printf("String: %s\n", emoji)
	fmt.Printf("Length in bytes: %d\n", len(emoji))
	fmt.Printf("Length in runes: %d\n", utf8.RuneCountInString(emoji))

	// Demonstrating multi-byte character
	chinese := "世"
	fmt.Printf("Chinese character '世' - byte length: %d\n", len(chinese))
	fmt.Printf("First byte: %v, but character is: %s\n", chinese[0], chinese)

	// String iteration - byte vs rune
	fmt.Println("\n--- String Iteration: Bytes vs Runes ---")
	str := "Go语言"

	// Iterating as runes (correct way for Unicode)
	fmt.Println("Iterating as runes:")
	for i, r := range str {
		fmt.Printf("  Index %d: %U (%c)\n", i, r, r)
	}

	// String conversion
	fmt.Println("\n--- String Conversion ---")
	original2 := "Hello"

	// String to []byte
	bytes := []byte(original2)
	fmt.Printf("String to bytes: %v\n", bytes)
	bytes[0] = 'h'
	fmt.Printf("Modified bytes: %v\n", bytes)
	fmt.Printf("Original string unchanged: %s\n", original2)
	fmt.Printf("New string from bytes: %s\n", string(bytes))

	// String to []rune
	runes := []rune(original2)
	fmt.Printf("String to runes: %v\n", runes)
	runes[0] = 'h'
	fmt.Printf("New string from runes: %s\n", string(runes))

	// Common string operations using strings package
	fmt.Println("\n--- Common String Operations ---")
	sample := "  Go Programming Language  "

	fmt.Printf("Original: '%s'\n", sample)
	fmt.Printf("Trimmed: '%s'\n", strings.TrimSpace(sample))
	fmt.Printf("ToUpper: '%s'\n", strings.ToUpper(sample))
	fmt.Printf("ToLower: '%s'\n", strings.ToLower(sample))
	fmt.Printf("Contains 'Programming': %v\n", strings.Contains(sample, "Programming"))
	fmt.Printf("HasPrefix '  Go': %v\n", strings.HasPrefix(sample, "  Go"))
	fmt.Printf("HasSuffix 'Language  ': %v\n", strings.HasSuffix(sample, "Language  "))

	// Split and Join
	fmt.Println("\n--- Split and Join ---")
	csv := "apple,banana,cherry"
	parts := strings.Split(csv, ",")
	fmt.Printf("Split '%s': %v\n", csv, parts)
	joined := strings.Join(parts, " | ")
	fmt.Printf("Joined: %s\n", joined)

	//fullString := "Hello, World! 世"
	// WRONG: slicing by bytes can break multi-byte characters
	// Don't do: fullString[0:3] - might cut a character in half

}
