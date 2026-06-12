# ASCII-ART-WEB Learning Path

## Quest 1: Echo Server

Goal:
Learn how a web server works.

Requirements:

* Create a server on port 8080.
* When visiting "/", return:

Hello from my server

Learn:

* net/http
* http.HandleFunc
* http.ListenAndServe

Connection to ASCII-Art-Web:
Every request in the final project goes through handlers.

---

## Quest 2: Understanding Request and Response

Goal:
Understand the request lifecycle.

Requirements:

* Visiting "/hello" displays:

Welcome to my website

Learn:

* Browser → Request → Handler → Response

Connection to ASCII-Art-Web:
The HomeHandler and AsciiHandler are just specialized handlers.

---

## Quest 3: HTML Templates

Goal:
Separate HTML from Go code.

Requirements:

* Create index.html.
* Render it using template.ParseFiles.

Learn:

* html/template
* Execute()

Connection:
ASCII-Art-Web uses templates instead of fmt.Println.

---

## Quest 4: HTML Forms

Goal:
Accept input from users.

Requirements:

* Create a textarea.
* Submit text using POST.
* Display submitted text.

Learn:

* Forms
* POST requests
* r.FormValue()

Connection:
ASCII-Art-Web receives text exactly this way.

---

## Quest 5: Multiple Inputs

Goal:
Handle more than one field.

Requirements:

* Textarea
* Banner selector

Learn:

* Form parsing
* Multiple values

Connection:
The final project accepts text and banner.

---

## Quest 6: Packages and Separation of Concerns

Goal:
Move logic outside main.go.

Requirements:

* Create handlers package.

Learn:

* Packages
* Exported functions

Connection:
HomeHandler and AsciiHandler belong in handlers.

---

## Quest 7: Reading Files

Goal:
Read external files.

Requirements:

* Read a file and display contents.

Learn:

* os.ReadFile

Connection:
Banner files are read this way.

---

## Quest 8: Building an ASCII Generator

Goal:
Convert text into ASCII art.

Requirements:

* Use banner files.

Learn:

* Rune indexing
* Character mapping

Connection:
This is the core ASCII engine.

---

## Quest 9: Combining Web + ASCII

Goal:
Connect form submission to the ASCII generator.

Requirements:

* Input text
* Generate ASCII
* Show output

Connection:
This is the first complete version of ASCII-Art-Web.

---

## Quest 10: Error Handling

Goal:
Handle invalid requests.

Requirements:

* 404
* 400
* 500

Learn:

* http.Error
* Status codes

Connection:
Required by the project.

---

## Quest 11: Static Files

Goal:
Serve CSS files.

Requirements:

* Add style.css.

Learn:

* http.FileServer
* http.StripPrefix

Connection:
Required for stylize.

---

## Quest 12: CSS and Layout

Goal:
Improve usability.

Learn:

* Flexbox
* gap
* width
* max-width

Connection:
Transforms a functional project into a user-friendly one.

---

## Quest 13: Multiline Input

Goal:
Support Enter key.

Learn:

* \r\n
* strings.Split()

Connection:
Supports "Hello↵There".

---

## Quest 14: Preserving Blank Lines

Goal:
Handle consecutive newlines.

Connection:
Supports:

Hello

There

---

## Quest 15: Full ASCII-Art-Web

Combine everything:

Server
+
Templates
+
Forms
+
Handlers
+
Packages
+
File reading
+
ASCII generator
+
Error handling
+
Static files
+
Styling

Result:
ASCII-Art-Web
