# Text Completion/Editing/Auto-Correction Tool

## Overview

This project is a simple text completion/editing/auto-correction tool written in Go. It takes as input a file containing a text that needs to be modified, and outputs the modified text to a specified file.

## Usage

To use the tool, run the following command:

go run . input.txt output.txt

Where input.txt is the name of the file containing the text to be modified, and output.txt is the name of the file the modified text should be placed in.
Modifications

The tool performs the following modifications on the input text:

   - Every instance of (hex) is replaced with the decimal version of the word before it (which should always be a hexadecimal number).
   - Every instance of (bin) is replaced with the decimal version of the word before it (which should always be a binary number).
   - Every instance of (up) converts the word before it to uppercase.
   - Every instance of (low) converts the word before it to lowercase.
   - Every instance of (cap) converts the word before it to capitalized form.
   - If a number appears after (low), (up), or (cap), the specified number of words before it are converted to lowercase, uppercase, or capitalized form, respectively.
   - Every instance of the punctuation marks ., ,, !, ?, :, ; should be close to the previous word and with space apart from the next one, except if there are groups of punctuation like ... or !?, in which case the program should format the text accordingly.
   - The punctuation mark ' should always be found with another instance of it, and they should be placed to the right and left of the word in the middle of them, without any spaces.
   - Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or an h.

## Allowed Packages

Standard Go packages are allowed.

## Examples

Here are some examples of how the tool can be used:

1. $ cat sample.txt

2. it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

3. $ go run . sample.txt result.txt

4. $ cat result.txt

5.  It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.

6. $ cat sample.txt

7. Simply add 42 (hex) and 10 (bin) and you will see the result is 68.

8. $ go run . sample.txt result.txt

9. $ cat result.txt

10. Simply add 66 and 2 and you will see the result is 68.

11. $ cat sample.txt

12.  There is no greater agony than bearing a untold story inside you.

13. $ go run . sample.txt result.txt

14. $ cat result.txt

15. There is no greater agony than bearing an untold story inside you.

16. $ cat sample.txt

17. Punctuation tests are ... kinda boring ,don't you think !?

18. $ go run . sample.txt result.txt

19. $ cat result.txt

20. Punctuation tests are... kinda boring, don't you think!?

## Learning Objectives

This project will help you learn about:

- The Go file system (fs) API
- String and numbers manipulation

## Testing

It is recommended to have test files for unit testing.

## Auditing

This project will be audited by other students, and you will also be an auditor for other students' projects. Please create your own tests for yourself and for when you will correct your auditees.