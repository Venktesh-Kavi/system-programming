## Functional Options

Ref: Dot Conference Dev Cheney (https://www.youtube.com/watch?v=24lFtGHWxAQ)

### Context

Assume that we have a simple server.

```go
package gserver

import "net"

type Server struct {
	listener net.Listener
}

func (s *Server) Addr() net.Addr
func (s *Server) Shutdown()

// NewServer returns a new server listening on Addr
func NewServer(addr string) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	srv := Server{listener: l}
	go srv.run()
	return &srv, nil
}
```

Assume that we receive a lot of features coming in

- slow clients keep using up all my resources (disconnect slow clients).
- lack of support for tls
- user running your server cannot handle the number of connections coming, they want to limit it.
- user wants to limit the number of concurrent cnx's, being attacked by botner.

Now we will have to modify the API to accommodate these features.

```go
// Typically improve the API like this
func NewServer(addr string, clientTimeout time.Duration, maxConns, maxConcurrent int, cer *tls.Cert)
```

- The above improvement if flaky, new comers of your package have no idea of which parameters are important/optional.
- If i want to use it as testing server, should i provide a tls value?
- May be I dont want the concurrent conns, should i put 0 value? (This will probably make no cnx available).

The responsibility lies on the user to provide it correct.

### Solutions

#### Many Functions Make Light Work

```
NewServer(addr string) (*Server, error)

NewTLSServer(addr string, cert *tls.Cert) (*Server, error)

NewServerWithTimeout(addr string, timeout time.Duration) (*Server, error)

NewTLSServerWithTimeout(addr string, cert *tls.Cert, timeout time.Duration) (*Server, error)
```

- With this the caller can choose which variant to use basis their requirement.
- As the number of options increases providing every permutation of functions becomes non-viable.

#### Passing a Configuration Struct

```go
type Config struct {
// Timeout sets the amount of time before closing
// idle connections or forever if not provided.
Timeout time.Duration

// The server will accept TLS connection is the
// certificate provided.
Cert *tls.Cert
}

func NewServer(addr string, config Config) (*Server, error)
```

Advantages

- Configuration struct can grow over time.
- Public API remains the same
- What used to big block comment, now becomes a nice go doc which tells about each field.
- Enables callers to use the default value.

Disadvantages

Example

- Trouble with defaults
-

```go
type Config struct {
// The port to listen on, if unset defaults to 8080
Port int
}

func main() {
svr, _ := NewServer("localhost", Config{
Port: 0, // opps, can't do this
})
}
```

- In the above example. NewServer returns the a new server listening to port 8080.
- The user cannot set explicity set the port to 0 and have the os to choose for us an ephemeral port to listen on.
  because there is no way to tell that explicit zero that you set from the fields default from the zero value

Most of the time users want default behaviour. "I just want a server, I don't want to have to think about it"

```go
func NewServer(addr string, config Config) (*Server, error)

func main() {
src, _ := NewServer(addr, Config{})
}
```

- Most of the time they don't want to change any configuration, but they would have to pass something to the second
  argument.
- Magic empty value, why should users of the api just construct an empty value and pass it along.

Solution to the above problem

- Passing a pointer to the value instead, thereby enabling callers to use nil, rather than passing an empty value. 
- Still it as worse as the above solution. Still we need to pass the second argument.
- More concering, this configuration value is now retained with the caller and what its passed into. What happens if the config struct is mutated?.
- Well written APIs, should not require caller to create these dummy value to satisfy those rare use-cases
- As a Go programmer ensure that nil is never passed as a parameter to a public API.

```go
func NewServer(addr string, config *Config) (*Server, error) {
}

func main() {
	srv, _ := NewServer("localhost", nil) // accept the defaults
    config := Config{Port: 9000}
    srv2, _ := NewServer("localhost", &config)
    config.Port = 9001 // what happens now?
}
```

#### Variadic Configuration (Better Solution)

- Instead of passing nil or an empty value. 
- The variadic nature of the nature allows not passing anything at all or something.
- This solves two big problems 
  - Invocation of default behaviour becomes as concise as possible.
  - New server can accept values and not pointers to config values. This removes nil as a possible argument. Ensures caller does not hold state of the configuration.

```go
func NewServer(addr string, config ...Config) (*Server, error) {
}

func main() {
	srv, _ := NewServer("localhost") // defaults
    
    srv2, _ := NewServer("locahost", Config{
		Timeout: 300 * time.Second,
        MaxConns: 10,
    })
}
```

Problems
- Because the function parameter is variadic.The implementation should cope with multiple values if passed by the user.
- The expectation is the user should pass only one config.
- Is there a way in which we can have variadic paramters and improve expressiveness of configuration parameters when we need them?

#### Functional Options

Self-Referential functions and design by Rob Pike (Blog Post)

```go
func NewServer(addr string, options ...func(*Server)) (*Server, error)

func main() {
	svr, _ := NewServer("localhost") // default
    timeout := func(srv *Server) {
		srv.timeout = 60 * time.Second
    }
    tls := func(srv *Server) {
		config := loadTLSConfig()
        srv.listener = tls.NewListener(svr.listener, &config)
    }
}
```

- Sensible defaults, Highly Configurable, Can grow over time, self documented.
- Its safe and discoverable for newcomer. Never require nil or empty value.

