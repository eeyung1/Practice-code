# ASCII Art Color

## Description

ASCII Art Color is an extension of the ASCII Art project that allows users to generate ASCII art with colored output in the terminal. The program supports coloring either the entire input string or only specified substrings using ANSI escape codes while preserving the selected ASCII banner style.

## Features

* Convert text into ASCII art.
* Color the entire output.
* Color specific substrings.
* Highlight multiple occurrences of the same substring.
* Support multiple banner styles:

  * `standard`
  * `shadow`
  * `thinkertoy`
* Validate command-line arguments.
* Handle invalid banners, colors, and input gracefully.

---

## Project Structure

```text
.
├── ascii_art/
│   ├── color.go
│   ├── generate.go
│   ├── load_banner.go
│   ├── render.go
│   └── validate.go
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── main.go
└── README.md
```

---

## Requirements

* Go 1.20 or later
* A terminal that supports ANSI color escape sequences

---

## Usage

### Print ASCII art

```bash
go run . "Hello"
```

### Print using another banner

```bash
go run . "Hello" shadow
```

### Color the entire string

```bash
go run . --color=red "Hello"
```

### Color the entire string with another banner

```bash
go run . --color=green "Hello" thinkertoy
```

### Color a substring

```bash
go run . --color=blue ell "Hello"
```

### Color a substring using another banner

```bash
go run . --color=yellow ell "Hello" shadow
```

---

## Supported Colors

The application supports the following ANSI colors:

* black
* red
* green
* yellow
* blue
* purple
* cyan
* white

---

## Behavior

* If no substring is provided, the entire string is colored.
* Every occurrence of the specified substring is highlighted.
* Banner names must be valid.
* Color names must be supported.
* Input must contain only printable ASCII characters.

---

## Error Handling

The program validates:

* Invalid command-line usage
* Unsupported colors
* Invalid banner names
* Missing or corrupted banner files
* Non-printable ASCII characters

When invalid input is detected, a usage message or appropriate error is displayed.

---

## Example

Command:

```bash
go run . --color=red kit "a king kitten have kit"
```

Behavior:

* Every occurrence of `kit` is displayed in red.
* The remaining text is rendered normally.
* The output is generated using the default `standard` banner unless another banner is specified.

---

## Authors

Developed as part of the ASCII Art optional project using Go.
By:
##### Akilozi Samuel,
##### Elaigwu Emmanuel, and
##### Eyung Sunday

