// Code generated by protoc-gen-go-nats. DO NOT EDIT.
// source: example/example.proto

package example

import (
	"context"
	"log/slog"
	"strings"
	proto "google.golang.org/protobuf/proto"
	nats "github.com/nats-io/nats.go"
	micro "github.com/nats-io/nats.go/micro"
)

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
func NewNATSGreeterServer(ctx context.Context, nc *nats.Conn, server GreeterServer, version, subject, queue string) (micro.Service, error) {
	serviceName := "GreeterServer"

	cfg := micro.Config{
		Name:       serviceName,
		Version:    version,
		QueueGroup: queue,
	}

	srv, err := micro.AddService(nc, cfg)
	if err != nil {
		return nil, err
	}

	logger := slog.With(
		slog.Group(
			"service",
			slog.String("name", serviceName),
			slog.String("version", version),
			slog.String("queue", queue),
		),
	)

	var endpointSubject string
	var mlogger *slog.Logger

	endpointSubject = strings.ToLower("svc.Greeter.SayHello")

	mlogger = logger.With(
		slog.Group(
			"endpoint",
			slog.String("subject", endpointSubject),
		),
	)

	mlogger.Info("registring endpoint")
	err = srv.AddEndpoint(
		"Greeter",
		micro.ContextHandler(
			ctx,
			func(ctx context.Context, req micro.Request) {
				r := &HelloRequest{}

				/*
					Unmarshal the request.
				*/
				if err := proto.Unmarshal(req.Data(), r); err != nil {
					mlogger.Error("unmarshaling request", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Forward on the original request to the original gRPC service.
				*/
				resp, err := server.SayHello(ctx, r)
				if err != nil {
					mlogger.Error("service error", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Take the response from the gRPC service and dump it as a byte array.
				*/
				respDump, err := proto.Marshal(resp)
				if err != nil {
					mlogger.Error("marshaling response", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Finally response with the original response from the gRPC service.
				*/
				if err := req.Respond(respDump); err != nil {
					mlogger.Error("sending response", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}
			},
		),
		micro.WithEndpointSubject(endpointSubject),
	)

	if err != nil {
		mlogger.Error("registering endpoint", slog.String("reason", err.Error()))
		return nil, err
	}

	endpointSubject = strings.ToLower("svc.Greeter.SayHelloAgain")

	mlogger = logger.With(
		slog.Group(
			"endpoint",
			slog.String("subject", endpointSubject),
		),
	)

	mlogger.Info("registring endpoint")
	err = srv.AddEndpoint(
		"Greeter",
		micro.ContextHandler(
			ctx,
			func(ctx context.Context, req micro.Request) {
				r := &HelloRequest{}

				/*
					Unmarshal the request.
				*/
				if err := proto.Unmarshal(req.Data(), r); err != nil {
					mlogger.Error("unmarshaling request", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Forward on the original request to the original gRPC service.
				*/
				resp, err := server.SayHelloAgain(ctx, r)
				if err != nil {
					mlogger.Error("service error", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Take the response from the gRPC service and dump it as a byte array.
				*/
				respDump, err := proto.Marshal(resp)
				if err != nil {
					mlogger.Error("marshaling response", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Finally response with the original response from the gRPC service.
				*/
				if err := req.Respond(respDump); err != nil {
					mlogger.Error("sending response", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}
			},
		),
		micro.WithEndpointSubject(endpointSubject),
	)

	if err != nil {
		mlogger.Error("registering endpoint", slog.String("reason", err.Error()))
		return nil, err
	}

	endpointSubject = strings.ToLower("svc.Greeter.SayGoodbye")

	mlogger = logger.With(
		slog.Group(
			"endpoint",
			slog.String("subject", endpointSubject),
		),
	)

	mlogger.Info("registring endpoint")
	err = srv.AddEndpoint(
		"Greeter",
		micro.ContextHandler(
			ctx,
			func(ctx context.Context, req micro.Request) {
				r := &SayGoodbyeRequest{}

				/*
					Unmarshal the request.
				*/
				if err := proto.Unmarshal(req.Data(), r); err != nil {
					mlogger.Error("unmarshaling request", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Forward on the original request to the original gRPC service.
				*/
				resp, err := server.SayGoodbye(ctx, r)
				if err != nil {
					mlogger.Error("service error", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Take the response from the gRPC service and dump it as a byte array.
				*/
				respDump, err := proto.Marshal(resp)
				if err != nil {
					mlogger.Error("marshaling response", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}

				/*
					Finally response with the original response from the gRPC service.
				*/
				if err := req.Respond(respDump); err != nil {
					mlogger.Error("sending response", slog.String("reason", err.Error()))
					handleError(req, err)
					return
				}
			},
		),
		micro.WithEndpointSubject(endpointSubject),
	)

	if err != nil {
		mlogger.Error("registering endpoint", slog.String("reason", err.Error()))
		return nil, err
	}

	return srv, nil
}

// NATSGreeterClient is a client connecting to a NATS GreeterServer.
type NATSGreeterClient struct {
	nc      *nats.Conn
	subject string
	queue   string
}

// NewNATSGreeterClient returns a new GreeterServer client.
func NewNATSGreeterClient(nc *nats.Conn, queue string) *NATSGreeterClient {
	return &NATSGreeterClient{
		nc:    nc,
		queue: queue,
	}
}

// Sends a greeting
func (c *NATSGreeterClient) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	subject := strings.ToLower("svc.Greeter.SayHello")

	payload, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	respPayload, err := c.nc.RequestWithContext(ctx, subject, payload)
	if err != nil {
		return nil, err
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

	payload, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	respPayload, err := c.nc.RequestWithContext(ctx, subject, payload)
	if err != nil {
		return nil, err
	}

	resp := &HelloReply{}
	if err := proto.Unmarshal(respPayload.Data, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *NATSGreeterClient) SayGoodbye(ctx context.Context, req *SayGoodbyeRequest) (*SayGoodbyeReply, error) {
	subject := strings.ToLower("svc.Greeter.SayGoodbye")

	payload, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	respPayload, err := c.nc.RequestWithContext(ctx, subject, payload)
	if err != nil {
		return nil, err
	}

	resp := &SayGoodbyeReply{}
	if err := proto.Unmarshal(respPayload.Data, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
