package cmd

import (
	"log/slog"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
)

// Run starts running the plugin
func Run() error {
	protogen.Options{}.Run(
		func(gen *protogen.Plugin) error {
			for _, file := range gen.Files {
				if !file.Generate {
					continue
				}
				if err := generateFile(gen, file); err != nil {
					return err
				}
			}
			return nil
		},
	)
	return nil
}

const templ = `
// Code generated by protoc-gen-go-nats. DO NOT EDIT.
// source: {{.GeneratedFilenamePrefix}}.proto

package {{.GoPackageName}}

import (
    "context"
    "log/slog"
    "strings"
    "errors"
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

{{ range .Services }}

// NewNATS{{ .GoName }}Server returns the gRPC server as a NATS micro service.
func NewNATS{{ .GoName }}Server(ctx context.Context, nc *nats.Conn, server {{ .GoName }}Server, version, subject, queue string) (micro.Service, error) {
	serviceName := "{{ .GoName }}Server"

    cfg := micro.Config{
    	Name: serviceName,
        Version: version,
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

    {{ range .Methods }}
    endpointSubject = strings.ToLower("svc.{{ .Parent.GoName }}.{{ .GoName }}")

    mlogger = logger.With(
     	slog.Group(
      		"endpoint",
       		slog.String("subject", endpointSubject),
      	),
    )

    mlogger.Info("registring endpoint")
    err = srv.AddEndpoint(
        "{{ .Parent.GoName }}",
        micro.ContextHandler(
        	ctx,
        	func(ctx context.Context, req micro.Request){
         		r := &{{ .Input.GoIdent.GoName }}{}

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
                resp, err := server.{{ .GoName }}(ctx, r)
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
    {{ end }}

    return srv, nil
}

// NATS{{ .GoName }}Client is a client connecting to a NATS {{ .GoName }}Server.
type NATS{{ .GoName }}Client struct {
	nc *nats.Conn
	subject string
	queue string
}

// NewNATS{{ .GoName }}Client returns a new {{ .GoName }}Server client.
func NewNATS{{ .GoName }}Client(nc *nats.Conn, queue string) *NATS{{ .GoName }}Client {
	return &NATS{{ .GoName }}Client{
		nc: nc,
		queue: queue,
	}
}

{{ range .Methods }}
{{ .Comments.Leading }}func (c *NATS{{ .Parent.GoName }}Client) {{ .GoName }}(ctx context.Context, req *{{ .Input.GoIdent.GoName }}) (*{{ .Output.GoIdent.GoName }}, error) {
	subject := strings.ToLower("svc.{{ .Parent.GoName }}.{{ .GoName }}")

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

	resp := &{{ .Output.GoIdent.GoName }}{}
	if err := proto.Unmarshal(respPayload.Data, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
{{ end }}

{{ end }}

`

// generateFile generates a .pb.go file.
func generateFile(gen *protogen.Plugin, file *protogen.File) error {

	tmpl, err := template.New("nats-micro-service").Parse(templ)
	if err != nil {
		return err
	}

	filename := file.GeneratedFilenamePrefix + ".nats.pb.go"

	logger := slog.With("filename", filename)
	logger.Info("generating the files")

	if len(file.Services) == 0 {
		// nothing to do here - no services found in this file.
		return nil
	}

	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	if err := tmpl.Execute(g, file); err != nil {
		logger.Error("failed to execute the template", slog.String("reason", err.Error()))
		return err
	}

	return nil
}
