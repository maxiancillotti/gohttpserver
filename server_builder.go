package gohttpserver

import (
	"net/http"
	"time"
)

// Methods returning interface can concatenate method calls
type HttpServerBuilder interface {

	// Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	// See net.Dial for details of the address format.
	SetAddr(string) HttpServerBuilder

	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body. A zero or negative value means
	// there will be no timeout.
	//
	// Because ReadTimeout does not let Handlers make per-request
	// decisions on each request body's acceptable deadline or
	// upload rate, most users will prefer to use
	// ReadHeaderTimeout. It is valid to use them both.
	SetReadTimeout(time.Duration) HttpServerBuilder

	// ReadHeaderTimeout is the amount of time allowed to read
	// request headers. The connection's read deadline is reset
	// after reading the headers and the Handler can decide what
	// is considered too slow for the body. If ReadHeaderTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	// Default 5 seconds.
	SetReadHeaderTimeout(time.Duration) HttpServerBuilder

	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	// A zero or negative value means there will be no timeout.
	// Default 15 seconds.
	SetWriteTimeout(time.Duration) HttpServerBuilder

	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	// Default 90 seconds.
	SetIdleTimeout(time.Duration) HttpServerBuilder

	// Build sets the previously configured parameters into our HTTP Server
	// and returns it so the user can call ListenAndServer() method.
	Build(handler http.Handler) HttpServer
}

type httpServerBuilder struct {
	addr              string
	readTimeout       time.Duration
	readHeaderTimeout time.Duration
	writeTimeout      time.Duration
	idleTimeout       time.Duration
}

// NewBuilder receives a handler and returns a HttpServerBuilder
// that you can configure to build a HttpServer.
func NewBuilder() HttpServerBuilder {
	return &httpServerBuilder{
		readHeaderTimeout: 5 * time.Second,
		writeTimeout:      15 * time.Second,
		idleTimeout:       90 * time.Second,
	}
}

func (b *httpServerBuilder) Build(handler http.Handler) HttpServer {
	return &httpServer{
		server: http.Server{
			Addr:              b.addr,
			ReadTimeout:       b.readTimeout,
			ReadHeaderTimeout: b.readHeaderTimeout,
			WriteTimeout:      b.writeTimeout,
			IdleTimeout:       b.idleTimeout,
			Handler:           handler,
		},
	}
}

func (b *httpServerBuilder) SetAddr(addr string) HttpServerBuilder {
	b.addr = addr
	return b
}

func (b *httpServerBuilder) SetReadTimeout(timeout time.Duration) HttpServerBuilder {
	b.readTimeout = timeout
	return b
}

func (b *httpServerBuilder) SetReadHeaderTimeout(timeout time.Duration) HttpServerBuilder {
	b.readHeaderTimeout = timeout
	return b
}

func (b *httpServerBuilder) SetWriteTimeout(timeout time.Duration) HttpServerBuilder {
	b.writeTimeout = timeout
	return b
}

func (b *httpServerBuilder) SetIdleTimeout(timeout time.Duration) HttpServerBuilder {
	b.idleTimeout = timeout
	return b
}
