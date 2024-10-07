// Copyright (c) 2024 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
)

// Define the set of Hiragana syllables in Romaji.
var hiraganaSyllables = []string{
	// "a", "i", "u", "e", "o",
	"ka", "ki", "ku", "ke", "ko",
	"sa" /*"shi",*/, "su", "se", "so",
	"ta" /*"chi", "tsu",*/, "te", "to",
	"na", "ni", "nu", "ne", "no",
	"ha", "hi", "fu", "he", "ho",
	"ma", "mi", "mu", "me", "mo",
	"ya", "yu", "yo",
	"ra", "ri", "ru", "re", "ro",
	"wa", "wo",
	// "n",
}

// Function to generate a random string of Hiragana syllables.
func generateHiraganaString(syllableCount int) string {
	result := ""

	for i := 0; i < syllableCount; i++ {
		randomIndex := rand.IntN(len(hiraganaSyllables))
		result += hiraganaSyllables[randomIndex]
	}

	return result
}

// Function to display help information.
func displayHelp() {
	fmt.Println("Usage: nunino [options]")
	fmt.Println("Options:")
	fmt.Println("  -lines, -l int")
	fmt.Println("        Number of lines to output (default 1)")
	fmt.Println("  -syllables, -s int")
	fmt.Println("        Number of syllables per line (default 3)")
}

// Function to display error messages.
func displayError(message string) {
	fmt.Fprintln(os.Stderr, message)
}

func main() {
	fs := flag.NewFlagSet("nunino", flag.ContinueOnError)

	lines := fs.Int("lines", 1, "Number of lines to output")
	syllables := fs.Int("syllables", 3, "Number of syllables per line")

	// Add shorthand flags.
	fs.IntVar(lines, "l", 1, "Number of lines to output")
	fs.IntVar(syllables, "s", 3, "Number of syllables per line")

	fs.Usage = displayHelp

	if err := fs.Parse(os.Args[1:]); err != nil {
		displayHelp()

		return
	}

	if *lines < 1 {
		displayError("Error: The number of lines must be at least 1.")
		displayHelp()

		return
	}

	if *syllables < 1 {
		displayError("Error: The number of syllables must be at least 1.")
		displayHelp()

		return
	}

	for i := 0; i < *lines; i++ {
		hiraganaString := generateHiraganaString(*syllables)
		fmt.Println(hiraganaString)
	}
}
