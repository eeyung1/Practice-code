# ascii-art-web

## Description

A web-based GUI for generating ASCII art. Users enter text, choose a banner style (standard, shadow, or thinkertoy), and the server returns the input rendered as ASCII art using the selected character set.

## Authors

- Your Name Here

## Usage: How to Run

**Prerequisites:** Go 1.21 or higher installed.

1. Clone or copy the project into a directory.
2. Add the three banner files into the `banners/` folder:
   - `banners/standard.txt`
   - `banners/shadow.txt`
   - `banners/thinkertoy.txt`
3. From the project root, run:

```bash
go run .
```

4. Open your browser and navigate to:

```
http://localhost:8080
```

## Implementation Details: Algorithm

### Banner File Format
Each banner `.txt` file defines printable ASCII characters (space through tilde, codes 32–126). Every character occupies exactly **8 lines**, preceded by a single blank line separator. This means each character's block starts at index `(charCode - 32) * 9 + 1` within the file's line array.

### ASCII Art Generation (`ascii/ascii.go`)
1. The banner file is read line-by-line into a string slice using `bufio.Scanner`.
2. The user's input is split on `\n` to handle multi-line input.
3. For each line of input, the output is built **row by row** (8 rows per character):
   - For each row (0–7), every character in the input line contributes one row of its 8-line definition.
   - Rows are concatenated horizontally to produce the full ASCII art line.
4. Empty input lines produce a single newline in the output.

### HTTP Server (`main.go` + `handlers/handlers.go`)
- `GET /` — serves the homepage template with empty fields.
- `POST /ascii-art` — reads `text` and `banner` from the submitted form, validates inputs, calls the ASCII generator, and re-renders the template with the result or an appropriate error.
- Status codes `200`, `400`, `404`, and `500` are returned as appropriate.
- HTML is rendered using Go's `html/template` package, which auto-escapes values to prevent XSS.
