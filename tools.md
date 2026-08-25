# Andrew Gerrand
9 Nov 2012, adg@golang.org
https://go.dev/talks/2012/simple.slide#1
http://vimeo.com/53221558
## Design
> "Consensus drove the design. Nothing went into the language until [Ken Thompson, Robert Griesemer, and myself] all agreed that it was right. Some features didn’t get resolved until after a year or more of discussion." - Rob Pike

* Lightweight, avoids unnecessary repetition
* Object Oriented, but not in the usual way
* Concurrent, in a way that keeps you sane
* Designed for working programmers
## Example
``` Go
package main

import "fmt"

func main() {
    fmt.Println("Hello, go")
}
```
## Libraries
Packages contain type, function, variable, and constant declarations.

Case determines visibility: Foo is exported, foo is not
### io
The io package provides fundamental I/O interfaces that are used throughout most Go code.

The most ubiquitous are the Reader and Writer types, which describe streams of data.

``` Go
package io

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Reader interface {
    Read(p []byte) (n int, err error)
}
```
Reader and Writer implementations include files, sockets, (de)compressors, image and JSON codecs, and many more.
``` Go
// Chaining io.Readers
package main

import (
    "compress/gzip"
    "encoding/base64"
    "io"
    "os"
    "strings"
)

func main() {
    var r io.Reader
    r = strings.NewReader(data)
    r = base64.NewDecoder(base64.StdEncoding, r)
    r, _ = gzip.NewReader(r)
    io.Copy(os.Stdout, r)
}

const data = `
H4sIAAAJbogA/1SOO5KDQAxE8zlFZ5tQXGCjjfYIjoURoPKgcY0E57f4VZlQXf2e+r8yOYbMZJhoZWRxz3wkCVjeReETS0VHz5fBCzpxxg/PbfrT/gacCjbjeiRNOChaVkA9RAdR8eVEw4vxa0Dcs3Fe2ZqowpeqG79L995l3VaMBUV/02OS+B6kMWikwG51c8n5GnEPr11F2/QJAAD//z9IppsHAQAA
`
```
### net/http
The net/http package implements an HTTP server and client.

``` Go
package main

import (
    "fmt"
    "log"
    "net/http"
)

type Greeting string

func (g Greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, g)
}

func main() {
    err := http.ListenAndServe("localhost:4000", Greeting("Hello, go"))
    if err != nil {
        log.Fatal(err)
    }
}
```
### encoding/json
The encoding/json package converts JSON-encoded data to and from native Go data structures.

``` Go
const blob = `[
    {"Title":"Øredev", "URL":"http://oredev.org"},
    {"Title":"Strange Loop", "URL":"http://thestrangeloop.com"}
]`

type Item struct {
    Title string
    URL   string
}

func main() {
    var items []*Item
    json.NewDecoder(strings.NewReader(blob)).Decode(&items)
    for _, item := range items {
        fmt.Printf("Title: %v URL: %v\n", item.Title, item.URL)
    }
}
```
### time
The time package provides a representation of time and duration, and other time-related functions.

``` Go
    if time.Now().Hour() < 12 {
        fmt.Println("Good morning.")
    } else {
        fmt.Println("Good afternoon (or evening).")
    }
```

``` Go
    birthday, _ := time.Parse("Jan 2 2006", "Nov 10 2009") // time.Time
    age := time.Since(birthday)                            // time.Duration
    fmt.Printf("Go is %d days old\n", age/(time.Hour*24))
```
time.Time values also contain a time.Location (for display only):

``` Go
    t := time.Now()
    fmt.Println(t.In(time.UTC))
    home, _ := time.LoadLocation("Australia/Sydney")
    fmt.Println(t.In(home))
```
### flag
The flag package provides a simple API for parsing command-line flags.
``` bash
$ flag -message 'Hold on...' -delay 5m
```

``` Go
package main

import (
    "flag"
    "fmt"
    "time"
)

var (
    message = flag.String("message", "Hello!", "what to say")
    delay   = flag.Duration("delay", 2*time.Second, "how long to wait")
)

func main() {
    flag.Parse()
    fmt.Println(*message)
    time.Sleep(*delay)
}
```

## The go tool
The go tool is the de facto standard for building and installing Go code.

Compile and run a single-file program:
``` bash
$ go run hello.go
```
Build and install the package in the current directory (and its dependencies):
``` bash
$ go install
```
    
Build and install the fmt package (and its dependencies):
``` bash
$ go install fmt
```
This tool also acts as an interface for most of the Go tools.
### Import paths
The go tool is a "zero configuration" tool. No Makefiles or scripts. Just Go code.
Your build schema and code are always in sync; they are one and the same.

Package import paths mirror the code's location in the file system:

    src/
    github.com/nf/
        gosynth/
        main.go
        note.go
        osc.go
        wav/
        writer.go
The gosynth program imports the wav package:
``` Go
import "github.com/nf/wav"
```
Installing gosynth will automatically install the wav package:
``` bash
$ go install github.com/nf/gosynth
```
### Remote dependencies
The go tool also fetches Go code from remote repositories.

Import paths can be URLs:
``` Go
import "golang.org/x/net/websocket"
```
To fetch, build and install a package:
``` bash
$ go get code.google.com/p/go.net/websocket
```    
To fetch, build, and install gosynth and its dependencies:
``` bash
$ go get github.com/nf/gosynth    
```
This simple design leads to other cool tools: 
go.pkgdoc.org (deprecated)
### Godoc
Godoc extracts documentation from Go code and presents it in a variety of forms.

Comments need no special format, they just need to precede what they document.
``` Go
    // Split slices s into all substrings separated by sep and returns a slice of
    // the substrings between those separators.
    // If sep is empty, Split splits after each UTF-8 sequence.
    // It is equivalent to SplitN with a count of -1.
    func Split(s, sep string) []string {
```
Documentation that lives with code is easy to keep up-to-date.
### Gofmt
The gofmt tool is a pretty-printer for Go source code.
All Go code in the core is gofmt'd, as is ~70% of open source Go code.
Ends boring formatting discussions.
Improves readability. Improves writability.
Saves a huge amount of time.

## Tests: writing
The go tool and the testing package provide a lightweight test framework.
``` Go
    func TestIndex(t *testing.T) {
        var tests = []struct {
            s   string
            sep string
            out int
        }{
            {"", "", 0},
            {"", "a", -1},
            {"fo", "foo", -1},
            {"foo", "foo", 0},
            {"oofofoofooo", "f", 2},
            // etc
        }
        for _, test := range tests {
            actual := strings.Index(test.s, test.sep)
            if actual != test.out {
                t.Errorf("Index(%q,%q) = %v; want %v", test.s, test.sep, actual, test.out)
            }
        }
    }
```
### Tests: running
The go tool runs tests.

    $ go test
    PASS

    $ go test -v
    === RUN TestIndex
    --- PASS: TestIndex (0.00 seconds)
    PASS
To run the tests for all my projects:

    $ go test github.com/nf/...

### Tests: benchmarks
The testing package also supports benchmarks.

A sample benchmark function:
``` Go
    func BenchmarkIndex(b *testing.B) {
        const s = "some_text=some☺value"
        for i := 0; i < b.N; i++ {
            strings.Index(s, "v")
        }
    }
```
The benchmark package will vary b.N until the benchmark function lasts long enough to be timed reliably.

    $ go test -test.bench=Index
    PASS
    BenchmarkIndex    50000000            37.3 ns/op

### Tests: doc examples
The testing package also supports testable examples.
``` Go
    func ExampleIndex() {
        fmt.Println(strings.Index("chicken", "ken"))
        fmt.Println(strings.Index("chicken", "dmr"))
        // Output:
        // 4
        // -1
    }
```
Examples and built and run as part of the normal test suite:

    $ go test -v
    === RUN: ExampleIndex
    --- PASS: ExampleIndex (0.00 seconds)
    PASS
The example is displayed in godoc alongside the thing it demonstrates:
    https://pkg.go.dev/strings#Index

## More tools
* vet: checks code for common programmer mistakes
* pprof: CPU and memory profiling
* fix: automatically migrate code as APIs change
* GDB support
* Editor support: Vim, Emacs, Eclipse, Sublime Text

## Example
Webfront is an HTTP server and reverse proxy. It reads a JSON-formatted rule file like this:
``` JSON
[
    {"Host": "example.com", "Serve": "/var/www"},
    {"Host": "example.org", "Forward": "localhost:8080"}
]
```
For all requests to the host example.com (or any name ending in ".example.com") it serves files from the /var/www directory.

For requests to example.org, it forwards the request to the HTTP server listening on localhost port 8080.
### The Rule type
A Rule value specifies what to do for a request to a specific host.

``` Go
// Rule represents a rule in a configuration file.
type Rule struct {
    Host    string // to match against request Host header
    Forward string // non-empty if reverse proxy
    Serve   string // non-empty if file server
}
```
It corresponds directly with the entries in the JSON configuration file.
``` JSON
[
    {"Host": "example.com", "Serve": "/var/www"},
    {"Host": "example.org", "Forward": "localhost:8080"}
]
```
### Rule methods
``` Go
// Match returns true if the Rule matches the given Request.
func (r *Rule) Match(req *http.Request) bool {
    return req.Host == r.Host || strings.HasSuffix(req.Host, "."+r.Host)
}

// Handler returns the appropriate Handler for the Rule.
func (r *Rule) Handler() http.Handler {
    if h := r.Forward; h != "" {
        return &httputil.ReverseProxy{
            Director: func(req *http.Request) {
                req.URL.Scheme = "http"
                req.URL.Host = h
            },
        }
    }
    if d := r.Serve; d != "" {
        return http.FileServer(http.Dir(d))
    }
    return nil
}
```
### The Server type
The Server type is responsible for loading (and refreshing) the rules from the rule file and serving HTTP requests with the appropriate handler.
``` Go
// Server implements an http.Handler that acts as either a reverse proxy or
// a simple file server, as determined by a rule set.
type Server struct {
    mu    sync.RWMutex // guards the fields below
    mtime time.Time    // when the rule file was last modified
    rules []*Rule
}
// ServeHTTP matches the Request with a Rule and, if found, serves the
// request with the Rule's handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if h := s.handler(r); h != nil {
        h.ServeHTTP(w, r)
        return
    }
    http.Error(w, "Not found.", http.StatusNotFound)
}
```
### The handler method
``` Go
// handler returns the appropriate Handler for the given Request,
// or nil if none found.
func (s *Server) handler(req *http.Request) http.Handler {
    s.mu.RLock()
    defer s.mu.RUnlock()
    for _, r := range s.rules {
        if r.Match(req) {
            return r.Handler()
        }
    }
    return nil
}
```
### Parsing rules
The parseRules function uses the encoding/json package to read the rule file into a Go data structure.
``` Go
// parseRules reads rule definitions from file returns the resultant Rules.
func parseRules(file string) ([]*Rule, error) {
    f, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    var rules []*Rule
    err = json.NewDecoder(f).Decode(&rules)
    if err != nil {
        return nil, err
    }
    return rules, nil
}
```
### The loadRules method
``` Go
// loadRules tests whether file has been modified
// and, if so, loads the rule set from file.
func (s *Server) loadRules(file string) error {
    fi, err := os.Stat(file)
    if err != nil {
        return err
    }
    mtime := fi.ModTime()
    if mtime.Before(s.mtime) && s.rules != nil {
        return nil // no change
    }
    rules, err := parseRules(file)
    if err != nil {
        return fmt.Errorf("parsing %s: %v", file, err)
    }
    s.mu.Lock()
    s.mtime = mtime
    s.rules = rules
    s.mu.Unlock()
    return nil
}
```
### Constructing the server
``` Go
// NewServer constructs a Server that reads rules from file with a period
// specified by poll.
func NewServer(file string, poll time.Duration) (*Server, error) {
    s := new(Server)
    if err := s.loadRules(file); err != nil {
        return nil, err
    }
    go s.refreshRules(file, poll)
    return s, nil
}
```
This constructor function launches a goroutine running the refreshRules method.
### Refreshing the rules
``` Go
// refreshRules polls file periodically and refreshes the Server's rule
// set if the file has been modified.
func (s *Server) refreshRules(file string, poll time.Duration) {
    for {
        if err := s.loadRules(file); err != nil {
            log.Println(err)
        }
        time.Sleep(poll)
    }
}
```
### Bringing it all together
The main function parses command-line flags, constructs a Server, and launches an HTTP server that serves all requests with the Server.
``` Go
var (
    httpAddr     = flag.String("http", ":80", "HTTP listen address")
    ruleFile     = flag.String("rules", "", "rule definition file")
    pollInterval = flag.Duration("poll", time.Second*10, "file poll interval")
)

func main() {
    flag.Parse()

    s, err := NewServer(*ruleFile, *pollInterval)
    if err != nil {
        log.Fatal(err)
    }

    err = http.ListenAndServe(*httpAddr, s)
    if err != nil {
        log.Fatal(err)
    }
}
```
### Testing 
The Server integration test uses the httptest package to construct a dummy HTTP server, synthesizes a set of rules, and constructs a Server instance that uses those rules.

Each test case in the table specifies a request URL and the expected response code and body.

For each test case, construct an http.Request for the url and an httptest.ResponseRecorder to capture the response, and pass them to the Server.ServeHTTP method. Then check that the response matches the test case.

``` Go
func testHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
}

func TestServer(t *testing.T) {
    dummy := httptest.NewServer(http.HandlerFunc(testHandler))
    defer dummy.Close()

    ruleFile := writeRules([]*Rule{
        {Host: "example.com", Forward: dummy.Listener.Addr().String()},
        {Host: "example.org", Serve: "testdata"},
    })
    defer os.Remove(ruleFile)

    s, err := NewServer(ruleFile, time.Hour)
    if err != nil {
        t.Fatal(err)
    }

    var tests = []struct {
        url  string
        code int
        body string
    }{
        {"http://example.com/", 200, "OK"},
        {"http://foo.example.com/", 200, "OK"},
        {"http://example.org/", 200, "contents of index.html\n"},
        {"http://example.net/", 404, "Not found.\n"},
        {"http://fooexample.com/", 404, "Not found.\n"},
    }

    for _, test := range tests {
        req, _ := http.NewRequest("GET", test.url, nil)
        rw := httptest.NewRecorder()
        rw.Body = new(bytes.Buffer)
        s.ServeHTTP(rw, req)
        if g, w := rw.Code, test.code; g != w {
            t.Errorf("%s: code = %d, want %d", test.url, g, w)
        }
        if g, w := rw.Body.String(), test.body; g != w {
            t.Errorf("%s: body = %q, want %q", test.url, g, w)
        }
    }
}
```
## Extra
github.com/nf/webfront

## Go offline
https://go.dev/tour/welcome/3

This tour is also available as a stand-alone program that you can use without access to the internet. It builds and runs the code samples on your own machine.

To run the tour locally, you'll need to first install Go and then run:
``` bash
$ go install golang.org/x/website/tour@latest
```
This will place a tour binary in your GOPATH's bin directory. When you run the tour program, it will open a web browser displaying your local version of the tour.