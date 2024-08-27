# Andrew Gerrand
9 Nov 2012, https://go.dev/talks/2012/simple.slide#46

## The go tool
The go tool is the de facto standard for building and installing Go code.

Compile and run a single-file program:

    $ go run hello.go
Build and install the package in the current directory (and its dependencies):

    $ go install
Build and install the fmt package (and its dependencies):

    $ go install fmt
This tool also acts as an interface for most of the Go tools.

## Import paths
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

    import "github.com/nf/wav"
Installing gosynth will automatically install the wav package:

    $ go install github.com/nf/gosynth

## Remote dependencies
The go tool also fetches Go code from remote repositories.

Import paths can be URLs:

    import "golang.org/x/net/websocket"
To fetch, build and install a package:

    $ go get code.google.com/p/go.net/websocket
To fetch, build, and install gosynth and its dependencies:

    $ go get github.com/nf/gosynth
This simple design leads to other cool tools: go.pkgdoc.org (deprecated)

## Godoc
Godoc extracts documentation from Go code and presents it in a variety of forms.

Comments need no special format, they just need to precede what they document.

    // Split slices s into all substrings separated by sep and returns a slice of
    // the substrings between those separators.
    // If sep is empty, Split splits after each UTF-8 sequence.
    // It is equivalent to SplitN with a count of -1.
    func Split(s, sep string) []string {
Documentation that lives with code is easy to keep up-to-date.

## Gofmt
The gofmt tool is a pretty-printer for Go source code.

All Go code in the core is gofmt'd, as is ~70% of open source Go code.

Ends boring formatting discussions.

Improves readability. Improves writability.

Saves a huge amount of time.

## Tests: writing
The go tool and the testing package provide a lightweight test framework.

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

## Tests: running
The go tool runs tests.

    $ go test
    PASS

    $ go test -v
    === RUN TestIndex
    --- PASS: TestIndex (0.00 seconds)
    PASS
To run the tests for all my projects:

    $ go test github.com/nf/...

## Tests: benchmarks
The testing package also supports benchmarks.

A sample benchmark function:

    func BenchmarkIndex(b *testing.B) {
        const s = "some_text=some☺value"
        for i := 0; i < b.N; i++ {
            strings.Index(s, "v")
        }
    }
The benchmark package will vary b.N until the benchmark function lasts long enough to be timed reliably.

    $ go test -test.bench=Index
    PASS
    BenchmarkIndex    50000000            37.3 ns/op

## Tests: doc examples
The testing package also supports testable examples.

    func ExampleIndex() {
        fmt.Println(strings.Index("chicken", "ken"))
        fmt.Println(strings.Index("chicken", "dmr"))
        // Output:
        // 4
        // -1
    }
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