## === GROUPIE TRACKER STUDIES ===

### PROJECT BRIEF

-   In this project, there is a server already running out on the
    internet that holds data about artists, concert locations, and
    dates. You didn't built it. You can't change it, its just waiting
    for requests.
-   it is your go program that needs to go and fetch data, that means
    the go program will act as a client
-   so the groupie project will carry two tasks, it will act as a server
    by listening for requests coming from a user's browser, just like
    the ascii art web and it will also act as a client making its own
    requests outward to the external api to get the data it needs to
    respond.
-   so we talk of internal api when both the client and the server are
    running on the same computer and we talk of external api when the
    server belongs to someone else, running on their own machine,
    connected to the public internet.
    
### API
-   API is short for application programming interface.
-   It is defined as a digital intermediary or contract that allows two
    different software applications to securely communicate, exchange
    data, and trigger actions in each other. so its not really thing,
    its more like the rules that a client and a server agree to follow
    when talking to each other.
-   It is more like a contract btw two programs that specifies how one
    program can request something from another, and what it will get
    back.
-   Here's a real world scenario, when you open your phone and try to
    use a whether app, the app's server or program does not have the
    data as concerning the current whether info, but it will use an API
    to request data from an external weather database, the API then
    safely returns the whether info which the app thereafter displays
    the result to you.
-   You can think of an API as a waiter in a restaurant that receives
    your order and takes it to the kitchen unit where your order is
    arranged and then he or she returns back with your food.
-   It is very different from what a database is and what a
    user-interface is, it's just a secure courier system used to fetch
    data. 
    
### There are four elements in a standard API documentation:
E-O-D-S

1.  Endpoints This refers to the specific web addresses or urls where
    the request must be sent
2.  Operations This refers to the action method required, such as GET
    request or POST request.
3.  Data formats This is the structured data language the machines use
    to speak to each other. its usually JSON or XML(Extensible markup
    language)
4.  security and keys This refers to the verification systems ensuring
    only authorized applications can make requests. 
    
### REST API

-   Rest here stands for *** Representational State Transfer ***
-   REST is said to be a set of architectural principles for designing
    apis that communicate over http.
-   A rest API is one which uses HTTP requests and URLs to send and
    receive data. They use urls, the application sends request to that
    url and the server sends back data. 

### client and server
-   client here is the thing that is sending a request it can be a
    browser, a go program, a mobile app.
-   The server here is the thing that is receiving the request and
    sending back a response 
    
### here's a rule of thumb,
-   any time two programs are communicating over http, an api is
    involved. the question is who built what and who is playing which
    role?
-   the two roles include:
-   the client:= which is the one that makes request
-   the server := which is the one who listens and responds with a given
    response. *** note *** these are roles, not identities, the
    program can play both roles at different moments.
-   in the ascii art web project, an api did exist but it was literally
    internal, nobody else was involved, the communication is between
    parts of the same computer or system 

### base url
-   This is the root address of an API, its the common prefix that all
    its endpoints share. for instance:
    http://groupitraskers.herokuapp.com/api

all endpoints of this url must carry this line as a prefix

### endpoint

This is the specific URL path that represents a specific resource or
collection of resources. it is directly where is request if sent. for
instance: http://groupitraskers.herokuapp.com/api/artists
http://groupitraskers.herokuapp.com/api/locations
http://groupitraskers.herokuapp.com/api/dates
http://groupitraskers.herokuapp.com/api/relation

### making http GET requests in go: http.Get

-   In the go net/http package, there is a function that allows a go
    program to act as an Http client, the tool is http.Get()

-   when you use the http.Get() to send a request to a server, it
    returns two values, the response and then an error. here is a syntac
    format of hoe it is used: resp, err :=
    http.Get("https://groupietrackers.herokuapp.com/api/artists")

-   the resp object returns evrything that the server sent back: the
    status code, the headers, and the body. so here is how to read a
    particular field in a response: body, err := io.ReadAll(resp.Body)
    status, err := io.ReadAll(resp.StatusCode) etc. and always remember
    to check the status code. a status code could be 404 which means the
    request did go through but obviously there wont be content, so you
    must check for all status codes. something like this:

if resp.StatusCode != 200 { // something went wrong --- handle it }

### reading the response body

the body can be read using io.ReadAll() here is the syntax: body, err :=
io.ReadAll(resp.Body) - after this, body contains raw response which is
basically json text as bytes, you then convert those bytes to a string
or decode the json directly.

### closing the response body:

always remember to defer resp.Body.Close() after reading the response
object. 

### JSON Json is short for Javascript object notation. despite
the name, it has nothing to do with javascript in practice. - it is a
universal language that programs use to exchange data over the internet,
regardless of what language those programs are written in. - The reason
why json is the major languageg that programs use to transfer their data
is because it is human readable and it maps the data structures that
almost every programming language has. so if yuor program needs to
communincate with another one and transfer data between each other, such
data will be transfered using json because regardless of the language
your program is written in, such a program can still understand and work
with json. - for the groupie-tracker project, the api i am going to be
using in fetching data from another server will return the data in json
format, it is now my responsibility to convert such data into what my go
program can understand and work with.

### The shape of json - json has
6 value types and they include: 

### string Which is a text wrapped in double quotes. for instance: "Freddi Mercury" 

### number Integer or decimal, no quotes. for instance: 42 

### Boolean for instance: true or false (all in lowercases, no quotes) 

### Null This depicts the absence of a value, for instance: null 

### Object 
The object value injson is a key-value collection that has curly braces wrapped around its
pairs. The keys are always strings while the values can take any form of
the value types available in json. for instance: { "name": "Queen",
"members": 4, "active": true } 

### Array 
This is an ordered list of values, wrapped in square brackets. Elements can be any json value type.
The structure of an Array in json is like that of a slice in go. The
values must maintain the order, if an element is on 0, then it is always
first, if it's on 1, then it is always second and so on ... - in arrays
in json, there is something that happens that is called nesting, whereby
you are allowed to put any of the data types as an element or as a value
in the array. - here is a quick defference btw the object and the array
value types in json. - an object consists of key-pair values, where the
keys can only be strings, but their values can take any type at all
while an array consits of an ordered list of elements that never change
their order but they can have any type they want. - an object has its
elements wrapped by curly braces while an array has its element wrapped
by square brackets. and also an array can take any element of any value
type at all, even objects. 

### Nesting in json arrays 
An object can contain an array, which contains more objects, which contains more
arrays. There's no limit to how deep the nesting goes. the major thing
to do is, when you look at a json data, you are able to refine and make
it useable to your go program. *** note *** json is not the only way
that programs transfer data, before json, there has bening XML which is
short for: extensible markup language, which program used to transfer
data among themselves, it was similar to html. so if a program and
another agree to share their data using this, its fine, the data will be
shared, your program will only have to decode the data and turn it into
what you actully want to work with. also, its not all programming
languages that have a json library in them, for such you will have to
write the parsing logic yourself, its painfull but its not impossible.
There is also another one that does the same work, its called the
protocols buffers, it transfers data but in computer understable form
alone and not text form like json, so converting it is a bit of a
painful work. json is agreed upon and used because its structures maps
those upon which other programming languages are built and those
programming languages happen to be the major ones in the tech space.

### why json is the universal format for transfering data between two
programs - it is language-neutral all major programming languages have
json libraries. No negotiation is needed btw the client and the server
about data format, its strictly json. - it is self-describing or easy to
understand When you recieve a json data, the keys and values tell you
exactly what the data is all about and its content as well. you would
not need external documentation to read it. - It travels as plain text
it moves over http just like a web page does. 

### the translation problem: json to go lang 
There is a common problem when trying to work with
json data in go, the reason being that go lang is a statically typed language
where all variables have fixed types known at compile time. but json is
just text, it has not types in go sense. - when my go program receives a
response from an external API, what arrives is just a stream of text
bytes, and go cant really use or understand such a text, it needs it to
take a type recognised in go for it to be used, and this is where a
process known as unmarshalling comes in. *** UNMARSHALLING *** is the
process of reading a raw json text and transferring its values into a go
struct, field by field, with proper types. You define a struct that
mirrors the shape of the json: type Artist struct { Name string
`json:"name"` Formed int `json:"formed"` Active bool `json:"active"` }

then you call json.Unmarshal: var artist Artist json.Unmarshal(body,
&artist)

go reads the json text, finds the keys and sees their values and puts
them together field by field. - the backticks is go's way of telling the
json decoder which json key maps to which struct field. and note that go
requires all exported fields to start with a capital letter, that is why
the struct fields all starts with capital letters. 

### why it needs a pointer you pass &artist and not artist because json.Unmarshal needs to
modify the struct directly, but if you pass artist without the pointer,
go would hand the function a copy, the function would fill in the copy
and the original artist will remain empty, the & gives the function the
actual address of the struct so it can write to it dirrectly. 

### designing go structs with json struct tags the struct tags here is the
backticks that wraps the second part of the struct. for instance: type
Artist struct { ID int `json:"id"` Name string `json:"name"` Members
\[\]string `json:"members"` CreationDate int `json:"creationDate"`
FirstAlbum string `json:"firstAlbum"` Image string `json:"image"` }

#### json.Unmarshal --- The Core Tool

this function takes a \[\]bytes of json and a pointer to a go value, and
fills that value with the decoded data, it returns an error if something
went wrong, so you must always handle errors. it is this function that
decodes a json data into go usable info. 

### json.NewDecoder --- 
The Streaming Alternative This function is a different alternative or
approach to the same problem, but it differs a bit, instead of reading
all the bytes first and then decoding them, it reads and decodes in a
single steaming pass. here is the syntax for it: decoder :=
json.NewDecoder(resp.Body) var artist Artist decoder.Decode(&artist)

### differences btw json.Unmarshal and json.NewDecoder\
the first one requires you to have all the bytes in memory before
decoding starts. so for large responses, this means holding the entire
json in memory twice, once as raw bytes, once as the decoded struct.
while json.NewDecoder reads the bytes and decodes as it goes, using less
memory for large payloads. but for this groupie-tracker project, either
of them will work fine but its is recommndable to use the
json.NewDecoder. 

#### Nested structs: json objects inside json objects

#### json arrays mapping to go slices

-   whenever you see a json array, you use a go slice. The element type
    of the slice matches whatever type the array contains. *** note *** a single object maps to a struct while a collection of
    objects maps to a slice of structs. 
    ### the full mental model: json
    as a contract: the api owner from which you are going to be reading
    your data decides what their json looks like. the structure of the
    json is a contract. and your job is to write go structs that mirrow
    that contract exactly.

## === ADDITIONAL CONCEPTS DISCOVERED DURING IMPLEMENTATION ===

### View Models

-   As the project grows, a single template may require data coming from
    multiple API endpoints.
-   A view model is a struct created specifically for templates so that
    multiple independent pieces of data can be rendered together.
-   Example:

``` go
type ArtistPageData struct {
    Artist   models.Artist
    Relation models.Relation
}
```

-   This allows the handler to aggregate data before passing it to the
    template.

### Caching

-   Making repeated requests to the same external API endpoint wastes
    time and bandwidth.
-   Caching stores previously fetched data in memory so it can be reused
    without making another HTTP request.
-   Example:

``` go
var artistsCache []models.Artist
var relationCache = map[int]models.Relation{}
```

-   The first request fetches from the API while later requests read
    from memory.

### Concurrency

-   Two independent operations do not need to wait for each other.
-   If two requests can run at the same time, they should.
-   In Groupie Tracker, fetching artists and fetching relation data can
    happen concurrently.

### Goroutines

-   A goroutine is a lightweight thread managed by the Go runtime.
-   It is started with the `go` keyword:

``` go
go services.GetRelation(id)
```

-   The caller does not wait for the function to finish automatically.

### WaitGroups

-   Since goroutines run independently, the main goroutine sometimes
    needs to wait for workers to finish.
-   The `sync.WaitGroup` package provides this functionality.
-   Three operations exist:

``` go
wg.Add(n)
wg.Done()
wg.Wait()
```

### Cache Stampede

-   A cache stampede happens when several requests arrive at the same
    time before the cache has been populated.
-   All requests miss the cache and all of them make the same external
    API call.
-   This is usually solved with mutexes or request coalescing in
    production systems.

### Separation of Concerns

-   Templates should display data.
-   Go code should prepare and transform data.
-   Business logic should not be embedded inside templates.

### Mental Model of Groupie Tracker

``` text
Browser
↓
Handlers
↓
Services
↓
External API
↓
JSON Response
↓
Go Structs
↓
Templates
↓
Browser
```
