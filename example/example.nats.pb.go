// Code generated by protoc-gen-go-nats. DO NOT EDIT.
// source: example.proto

package example

import (
	"context"
	"log/slog"
	"strings"
	"errors"
	proto "google.golang.org/protobuf/proto"
	nats "github.com/nats-io/nats.go"
	micro "github.com/nats-io/nats.go/micro"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("example.proto")

// handleError is a helper which response with the error.
func handleError(req micro.Request, err error) {
	if sendErr := req.Error("500", err.Error(), nil); sendErr != nil {
		slog.Error(
			"error sending response error",
			slog.String("reason", sendErr.Error()),
			slog.String("subject", req.Subject()),
		)
	}
}

// NewNATSGreeterServer returns the gRPC server as a NATS micro service.
func NewNATSGreeterServer(ctx context.Context, nc *nats.Conn, server GreeterServer, version, queueGroup string) (micro.Service, error) {
	cfg := micro.Config{
		Name:        "GreeterServer",
		Version:     version,
		QueueGroup:  queueGroup,
		Description: "NATS micro service adaptor wrapping GreeterServer",
	}

	srv, err := micro.AddService(nc, cfg)
	if err != nil {
		return nil, err
	}

	logger := slog.With(
		slog.Group(
			"service",
			slog.String("name", cfg.Name),
			slog.String("version", cfg.Version),
			slog.String("queue-group", cfg.QueueGroup),
		),
	)

	logger.Info(
		"registring endpoint",
		slog.Group(
			"endpoint",
			slog.String("subject", strings.ToLower("svc.Greeter.SayHello")),
		),
	)

	err = srv.AddEndpoint(
		"Greeter",
		micro.ContextHandler(
			ctx,
			func(ctx context.Context, req micro.Request) {
				endpointSubject := strings.ToLower("svc.Greeter.SayHello")

				ctx, span := tracer.Start(ctx, "SayHello", trace.WithAttributes(attribute.String("subject", endpointSubject)))
				defer span.End()

				hlogger := logger.With(
					slog.Group(
						"endpoint",
						slog.String("subject", endpointSubject),
					),
				)

				r := &HelloRequest{}

				/*
					Unmarshal the request.
				*/
				if err := proto.Unmarshal(req.Data(), r); err != nil {
					hlogger.Error("unmarshaling request", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Forward on the original request to the original gRPC service.
				*/
				resp, err := server.SayHello(ctx, r)
				if err != nil {
					hlogger.Error("service error", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Take the response from the gRPC service and dump it as a byte array.
				*/
				respDump, err := proto.Marshal(resp)
				if err != nil {
					hlogger.Error("marshaling response", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Finally response with the original response from the gRPC service.
				*/
				if err := req.Respond(respDump); err != nil {
					hlogger.Error("sending response", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}
			},
		),
		micro.WithEndpointSubject(strings.ToLower("svc.Greeter.SayHello")),
	)

	if err != nil {
		logger.Error(
			"registering endpoint",
			slog.Group(
				"endpoint",
				slog.String("subject", strings.ToLower("svc.Greeter.SayHello")),
			),
			slog.String("reason", err.Error()),
		)
		return nil, err
	}

	logger.Info(
		"registring endpoint",
		slog.Group(
			"endpoint",
			slog.String("subject", strings.ToLower("svc.Greeter.SayHelloAgain")),
		),
	)

	err = srv.AddEndpoint(
		"Greeter",
		micro.ContextHandler(
			ctx,
			func(ctx context.Context, req micro.Request) {
				endpointSubject := strings.ToLower("svc.Greeter.SayHelloAgain")

				ctx, span := tracer.Start(ctx, "SayHelloAgain", trace.WithAttributes(attribute.String("subject", endpointSubject)))
				defer span.End()

				hlogger := logger.With(
					slog.Group(
						"endpoint",
						slog.String("subject", endpointSubject),
					),
				)

				r := &HelloRequest{}

				/*
					Unmarshal the request.
				*/
				if err := proto.Unmarshal(req.Data(), r); err != nil {
					hlogger.Error("unmarshaling request", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Forward on the original request to the original gRPC service.
				*/
				resp, err := server.SayHelloAgain(ctx, r)
				if err != nil {
					hlogger.Error("service error", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Take the response from the gRPC service and dump it as a byte array.
				*/
				respDump, err := proto.Marshal(resp)
				if err != nil {
					hlogger.Error("marshaling response", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Finally response with the original response from the gRPC service.
				*/
				if err := req.Respond(respDump); err != nil {
					hlogger.Error("sending response", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}
			},
		),
		micro.WithEndpointSubject(strings.ToLower("svc.Greeter.SayHelloAgain")),
	)

	if err != nil {
		logger.Error(
			"registering endpoint",
			slog.Group(
				"endpoint",
				slog.String("subject", strings.ToLower("svc.Greeter.SayHelloAgain")),
			),
			slog.String("reason", err.Error()),
		)
		return nil, err
	}

	logger.Info(
		"registring endpoint",
		slog.Group(
			"endpoint",
			slog.String("subject", strings.ToLower("svc.Greeter.SayGoodbye")),
		),
	)

	err = srv.AddEndpoint(
		"Greeter",
		micro.ContextHandler(
			ctx,
			func(ctx context.Context, req micro.Request) {
				endpointSubject := strings.ToLower("svc.Greeter.SayGoodbye")

				ctx, span := tracer.Start(ctx, "SayGoodbye", trace.WithAttributes(attribute.String("subject", endpointSubject)))
				defer span.End()

				hlogger := logger.With(
					slog.Group(
						"endpoint",
						slog.String("subject", endpointSubject),
					),
				)

				r := &SayGoodbyeRequest{}

				/*
					Unmarshal the request.
				*/
				if err := proto.Unmarshal(req.Data(), r); err != nil {
					hlogger.Error("unmarshaling request", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Forward on the original request to the original gRPC service.
				*/
				resp, err := server.SayGoodbye(ctx, r)
				if err != nil {
					hlogger.Error("service error", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Take the response from the gRPC service and dump it as a byte array.
				*/
				respDump, err := proto.Marshal(resp)
				if err != nil {
					hlogger.Error("marshaling response", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Finally response with the original response from the gRPC service.
				*/
				if err := req.Respond(respDump); err != nil {
					hlogger.Error("sending response", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}
			},
		),
		micro.WithEndpointSubject(strings.ToLower("svc.Greeter.SayGoodbye")),
	)

	if err != nil {
		logger.Error(
			"registering endpoint",
			slog.Group(
				"endpoint",
				slog.String("subject", strings.ToLower("svc.Greeter.SayGoodbye")),
			),
			slog.String("reason", err.Error()),
		)
		return nil, err
	}

	return srv, nil
}

// NATSGreeterClient is a client connecting to a NATS GreeterServer.
type NATSGreeterClient struct {
	nc *nats.Conn
}

// NewNATSGreeterClient returns a new GreeterServer client.
func NewNATSGreeterClient(nc *nats.Conn) *NATSGreeterClient {
	return &NATSGreeterClient{
		nc: nc,
	}
}

// Sends a greeting
func (c *NATSGreeterClient) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	subject := strings.ToLower("svc.Greeter.SayHello")

	ctx, span := tracer.Start(ctx, "SayHello", trace.WithAttributes(attribute.String("subject", subject)))
	defer span.End()

	payload, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	respPayload, err := c.nc.RequestWithContext(ctx, subject, payload)
	if err != nil {
		return nil, err
	}

	rpcError := respPayload.Header.Get(micro.ErrorHeader)
	if rpcError != "" {
		return nil, errors.New(rpcError)
	}

	resp := &HelloReply{}
	if err := proto.Unmarshal(respPayload.Data, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// Sends another greeting
func (c *NATSGreeterClient) SayHelloAgain(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	subject := strings.ToLower("svc.Greeter.SayHelloAgain")

	ctx, span := tracer.Start(ctx, "SayHelloAgain", trace.WithAttributes(attribute.String("subject", subject)))
	defer span.End()

	payload, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	respPayload, err := c.nc.RequestWithContext(ctx, subject, payload)
	if err != nil {
		return nil, err
	}

	rpcError := respPayload.Header.Get(micro.ErrorHeader)
	if rpcError != "" {
		return nil, errors.New(rpcError)
	}

	resp := &HelloReply{}
	if err := proto.Unmarshal(respPayload.Data, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *NATSGreeterClient) SayGoodbye(ctx context.Context, req *SayGoodbyeRequest) (*SayGoodbyeReply, error) {
	subject := strings.ToLower("svc.Greeter.SayGoodbye")

	ctx, span := tracer.Start(ctx, "SayGoodbye", trace.WithAttributes(attribute.String("subject", subject)))
	defer span.End()

	payload, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	respPayload, err := c.nc.RequestWithContext(ctx, subject, payload)
	if err != nil {
		return nil, err
	}

	rpcError := respPayload.Header.Get(micro.ErrorHeader)
	if rpcError != "" {
		return nil, errors.New(rpcError)
	}

	resp := &SayGoodbyeReply{}
	if err := proto.Unmarshal(respPayload.Data, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
