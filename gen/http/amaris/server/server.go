// Code generated by goa v3.7.10, DO NOT EDIT.
//
// amaris HTTP server
//
// Command:
// $ goa gen testamaris/design

package server

import (
	"context"
	"net/http"
	amaris "testamaris/gen/amaris"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the amaris service endpoint HTTP handlers.
type Server struct {
	Mounts         []*MountPoint
	SortNames      http.Handler
	PokemonName    http.Handler
	FriendlyChains http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the amaris service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *amaris.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"SortNames", "POST", "/sortnames"},
			{"PokemonName", "GET", "/pokemonname/{id}"},
			{"FriendlyChains", "POST", "/friendlychains"},
		},
		SortNames:      NewSortNamesHandler(e.SortNames, mux, decoder, encoder, errhandler, formatter),
		PokemonName:    NewPokemonNameHandler(e.PokemonName, mux, decoder, encoder, errhandler, formatter),
		FriendlyChains: NewFriendlyChainsHandler(e.FriendlyChains, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "amaris" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.SortNames = m(s.SortNames)
	s.PokemonName = m(s.PokemonName)
	s.FriendlyChains = m(s.FriendlyChains)
}

// Mount configures the mux to serve the amaris endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountSortNamesHandler(mux, h.SortNames)
	MountPokemonNameHandler(mux, h.PokemonName)
	MountFriendlyChainsHandler(mux, h.FriendlyChains)
}

// Mount configures the mux to serve the amaris endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountSortNamesHandler configures the mux to serve the "amaris" service
// "SortNames" endpoint.
func MountSortNamesHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/sortnames", f)
}

// NewSortNamesHandler creates a HTTP handler which loads the HTTP request and
// calls the "amaris" service "SortNames" endpoint.
func NewSortNamesHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSortNamesRequest(mux, decoder)
		encodeResponse = EncodeSortNamesResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "SortNames")
		ctx = context.WithValue(ctx, goa.ServiceKey, "amaris")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountPokemonNameHandler configures the mux to serve the "amaris" service
// "PokemonName" endpoint.
func MountPokemonNameHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/pokemonname/{id}", f)
}

// NewPokemonNameHandler creates a HTTP handler which loads the HTTP request
// and calls the "amaris" service "PokemonName" endpoint.
func NewPokemonNameHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodePokemonNameRequest(mux, decoder)
		encodeResponse = EncodePokemonNameResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "PokemonName")
		ctx = context.WithValue(ctx, goa.ServiceKey, "amaris")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountFriendlyChainsHandler configures the mux to serve the "amaris" service
// "FriendlyChains" endpoint.
func MountFriendlyChainsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/friendlychains", f)
}

// NewFriendlyChainsHandler creates a HTTP handler which loads the HTTP request
// and calls the "amaris" service "FriendlyChains" endpoint.
func NewFriendlyChainsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeFriendlyChainsRequest(mux, decoder)
		encodeResponse = EncodeFriendlyChainsResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "FriendlyChains")
		ctx = context.WithValue(ctx, goa.ServiceKey, "amaris")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
