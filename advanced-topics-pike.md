# Advanced Topics in Programming Languages
## Concurrency/message passing Newsqueak (May 9, 2007)
by Rob Pike@https://www.youtube.com/watch?v=hB05UFqOtFA

The world is concurrent but computers are not.

This is a profound mismatch.

Two approaches:
1. Make (interface to) the world sequential.
2. Make software concurrent.

A common approach to concurrency is to consider is a necessary evil.

"The world is concurrent, computers are becoming multicore, it's all ugly, better learn it!"--NO!

Concurrent programming provides a model for interface and design that simplifies even non-parallel software.
### State
An execution process has:
- program counter (PC)
- stack

Write this as (PC, stack). This is a powerful way to represent state. (Consider stack traceback.)

A concurrent program is (PC, stack)+ -- linearly growing aglomeration of counters and stacks.

Exponentially more powerful, yet conceptually just as easy to understand--and only linearly harder to debug. 
### The Model
Think about software as 
- interacting 
- independently executing
- processes

In Hoare's original forumaltion: Communicating Sequential Processes (no threads, shared memory, locks, semaphores etc.)

This isn't Andrew Birrell's model, which is too low-level for easy programming.
### Selected history
Hoare, 1978, CSP
- processes and communication

Cardelli & Pike, 1985, Squeak
- CSP with channels (but really just a toy)
- for interactive programs

Pike, 1988, Newsqueak
- a full, interpreted, experimental language

Winterbottom, 1995, Alef
- a full, compiled, systems language

Dorward, Pike, Winterbottom, 1996, Limbo
- approximately JIT-ed Newsqueak
### An overview of Newsqueak
- Looks like Sawzall (but richer)
- Pascal-like type declarations
- C-like control structures
- Functions are just variables (lambdas, `prog`)
- Process controls (`begin`, `become`)
- Channels are first-class citizens
- Memory: value semantics (no sharing)
### Examples
``` Go
hello: array of char = "hello";
i := 23 + 24;
b := mk(array[2] of array of char={s, "go"});
j: int;
for (i = 0; i < 10; i = i + 1)
    print("i is ", i, "\n");
```
#### Akermann's function
``` Go
rec ack := prog(a, b: int of int) {
    if (a == 0)
        become b + 1;
    if (b == 0)
        become ack(a - 1, 1);
    become ack(a - 1, ack(a, b - 1));
};

ack(3, 4);
```
### Functions and prog()
Programs are lambdas; this is a value:
``` Go
prog(a: int, b: int) of int {
    become a + b;
}
```
Four things to do with a prog:
1. Treat it like a value
2. Call it
3. Replace yourself with its execution
4. Start it as an independent process
#### prog as a value
Function defintiion is assignment:
``` Go
sum := prog(a, b: int) of int {
    become a + b;
};

diff := prog(a, b: int) of int {
    become a - b;
};

sum = diff;

a := sum(24, 23);
// Same result as
a := prog(a, b: int) of int {
    become a - b;
}(24, 23);
```
#### Replace yourself with a prog()
The keyword is `become`: it generalizes tail recursion. Any execution can become another of equal return type.

Any expression:
``` Go
prog(a, b: int) of int { become a + b; }
```
Any function invocation:
``` Go
diff := prog(a, b: int) of int {
    become sum(a, -b);
};
```
When aggressively possible, reuses stack space.
#### Start a prog() as a process
The keyword `begin` launches an independent process:
``` Go
begin sum(23, 5);   // no return value

begin prog(a, b: int) {
    for (; a > 0; a = a - 1) print(b);
}(10, sum(23, 5));  // print in background
```
Issue: how do we jnow when it's done?
There's no `join`.
Want some sort of notification.
### Channels
A channel (`chan`) is an unbuffered synchronous communication port.
- typed
- half duplex
- unbuffered
- first-class value
- send and receive operators
### Send and receive
``` Go
c : chan of int;        // declare
c = mk(chan of int);    // initialize
```
The communication operator is `<-`.
Direction is mnemonic.
``` Go
c<- = 23;   // send; usually written <-=
x = <-c;    // receive
```
Communication is synchronous:
- sender blocks until there is a receiver
- receiver blocks until there is a sender
- when both are ready, value is transferred

Asynchrony is rarely needed in practice...
### Asynchrony can be simulated
Asynchronous send, given channel `ch` and value `v`:
``` Go
begin prog(c: chan of int, a: int) {
    c <-= a;
}(ch, v);
```
Wrap it for prettiness:
``` Go
async_send := prog(ch: chan of int, v: int) {
    begin prog(c: chan of int, a: int) {
        c <-= a;
    }(ch, v);
};

async_send(ch, v); // a useful pattern: an asynchronous send done as a function call
```
> Asynchronous receive is trickier because you have to know when to receive. There's signaling involved, it involves multiple channels. 

Hint: channel can serve aas a signal as well as a value.

To discover when a process is done, pass it a channel with which to signal.
### Channels communicate
Some pointes about channels:
- they are used for communication
- they should be used for all communication
    - avoid shared memory; pass the data on a `chan`
    - ties signaling with data
- they're like capabilities
    - file descriptors, not files
        - original CSP missed this point
- they're first-class values
    - `chan of chan of int;`
### Select: a control structure
`select` is like `switch` but with communication:
``` Go
c1, c2: chan of int;
i: int;
select {
    case i = <-c1: print("A: ", i);
    case i = <-c2: print("B: ", i);
    case i = <-c1: print("C: ", i);
    case i = <-c2: print("D: ", i);
    case c2 <-= 7: print("S");
}
```
Satement blocks until a case can proceed.

If several are able, makes a random choice.

Compare with Dijkstra's *guarded commands*.
#### Array selection
For buidling *muxes*, can `select` on an array:
``` Go
a: array[N] of chan of int = mk();
i, j: int;
select {
    case <-a[]:         print("any");
    case i = <-a[]:     print("A: ", i);
    case <-a[i=]:       print("B: ", i);
    case i = <-a[j=]:   print("C: ", i, j);
}
```
Bracketed assignment reports mux id.
### A simple program
``` Go
counter := prog(c: chan of int) {
    i := 1;
    for (;;) c <-= (i = i + 1);
};

c := mk(chan of int);
begin counter(c);

<-c; // 2
<-c; // 3
```
#### Prime sieve
``` Go
filter := prog(prime: int, recv: chan of int, send: chan of int) {
    i: int;
    for (;;)
        if ((i = <-recv) % prime)
            send <-= i;
};

sieve := prog(prime: chan of int) {
    c := mk(chan of int);
    begin counter(c);
    p: int;
    for (;;) {
        prime <-= p <-c;
        newc := mk(chan of int);
        begin filter(p, c, newc);
        c = news;
    }
};

prime := mk(chan of int);
begin sieve(prime);    // begin with channel -- it is a useful pattern
<-prime; // 2
<-prime; // 3
<-prime; // 5
<-prime; // 7
```
A channel is a *capability*. Given a channel, one has an implicit contract to (say) receive some data value or sequence of data values, maybe later.
#### Power series
Sum of all a<sub>i</sub>x<sup>i</sup> e.g. the Taylor series (for the exponential function):

e<sup>x</sup> = Sum<sub>n=0</sub>(x<sup>n</sup>/n! = 1 + x + 1/2 x<sup>2</sup> + 1/6 x<sup>3</sup> ...)

Define a (rational) power series like:
``` Go
type rat: struct of {num: int; den: int; };
type ps: chan of rat;
```
Assume we have operators for manipulating rationals (`ratadd`, `ratmul` etc.)

Addition:
``` Go
do_psadd := prog(F: ps, G: ps, S: ps) {
    for (;;) S <-= ratadd(<-F, <-G);
};
```
Notice that without a return value the `prog` looks clunky.
``` Go
psadd := prog(F: ps, G: ps) of ps {
    S := mk(chan of rat);
    begin prog() {
        for (;;)
            S <-= ratadd(<-F, <-G);
    }();
    become S;
};
```
Pattern: begin a process, return a new channel. Leads to easy definition for operators:
- derivative:
    - drop the 0th term; `for(;;) P <-= ratmul(<-F, i)`
- integration:
    - emit constant; `for(;;) P <-= ratdiv(<-F, i)`

Example: Series for tangent is very messy to calculate; this method makes it quite concise and easy:
``` Go
psrev(psinteg(psmsubst(Ones, ratmk(-1,1), 2), 0));
// 0 1 0 1/3 0 2/15 0 17/315 0 62/2835 0 ...
```
See Doug McIlroy's [paper](https://swtch.com/~rsc/thread/squint.pdf) for details.
### Interfaces
A channel represents a contract--it's an interface.

Gather a set of channels together and one can express a rich interface to an abstract, independently executing process.

Like a class, but with communication rather than function calls:
``` Go
type Interface: struct of {
    c1: chan of T1;
    c2: chan of T2;
};
```
Instantiate a process, use channels to talk to it:
``` Go
in := mk(Interface);
begin prog(i: Interface) { /* code here */ } (in);
```
#### Window system: the client
A client has an environment defined as channels:
``` Go
type Env: struct of {
    G: chan of grahics; // graphical commands
    M: chan of Mouse;   // mouse
    K: chan of int;     // keyboard
};
```
Note: no events; channels of values.

Its interface is define as process type wrapping Env: `type Client: prog(env: Env);`

Its execution manages the client (or its local proxy).
#### Window system: the mux
A window system is just a mux for Envs: http://go/concurrentwindowsystem. Main loop is an array select.

The Newsqueak window system was just a toy but it clarified thinking about parallele interfaces. Many benefits including window system as client of itself.

The same fundamental design went into 8-and-a-half and then, pushing even farther, `rio`, and then Acme, production window systems in Plan 9. Alef and then a C library provided the mechanism. All major Plan 9 services were written this way.

Also see, https://www.usenix.org/legacy/publications/library/proceedings/sf94/full_papers/pike.pdf
### The system model
Define components as interfaces with all data flow and sharing done as communication over channels. 

The interface is a type; implementations of that interface just honor the protocol. 

Composition is linear in complexity of design but superlinear in expressibility. (The opposite of composition of state machines.) Interleaving is free. Compose interfaces, not state machines.

Parallelism is not the point, but falls out for free. Networking and remote execution are not the point, but also can fall out (although not quite for free).
### Conculsions
Concurrent processes with message passing can be a powerful model for programming, whether the problem is intrinsically parallel (server design) or not (power series).

You need both concurrency and communication, and you want high-level support. The epxressiveness--notation--is important.

Clarifies thinking about
- interfaces
- sharing
- comms