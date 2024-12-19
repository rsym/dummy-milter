package main

import (
	"context"
	"log/slog"
	"net"
	"net/textproto"

	"github.com/emersion/go-milter"
)

type SimpleMilter struct{}

func (m *SimpleMilter) Connect(host string, family string, port uint16, addr net.IP, mod *milter.Modifier) (milter.Response, error) {
	slog.InfoContext(context.Background(), "Connect from ", slog.String("host", host), ":", slog.Int("port", int(port)))
	return milter.RespContinue, nil
}

func (m *SimpleMilter) Helo(name string, mod *milter.Modifier) (milter.Response, error) {
	slog.InfoContext(context.Background(), "HELO ", slog.String("name", name))
	return milter.RespContinue, nil
}

func (m *SimpleMilter) MailFrom(from string, esmtpArgs *milter.Modifier) (milter.Response, error) {
	slog.InfoContext(context.Background(), "MAIL FROM: ", slog.String("from", from))
	return milter.RespContinue, nil
}

func (m *SimpleMilter) RcptTo(rcpt string, esmtpArgs *milter.Modifier) (milter.Response, error) {
	slog.InfoContext(context.Background(), "RCPT TO: ", slog.String("rcpt", rcpt))
	return milter.RespContinue, nil
}

func (m *SimpleMilter) Data() (milter.Response, error) {
	slog.InfoContext(context.Background(), "DATA")
	return milter.RespContinue, nil
}

func (m *SimpleMilter) Header(name string, value string, mod *milter.Modifier) (milter.Response, error) {
	slog.InfoContext(context.Background(), "Header: ", slog.String("name", name), ": ", slog.String("value", value))
	return milter.RespContinue, nil
}

func (m *SimpleMilter) Headers(headers textproto.MIMEHeader, mod *milter.Modifier) (milter.Response, error) {
	slog.InfoContext(context.Background(), "Headers: ", slog.Any("headers", headers))
	return milter.RespContinue, nil
}

func (m *SimpleMilter) Body(esmtpArgs *milter.Modifier) (milter.Response, error) {
	slog.InfoContext(context.Background(), "Body")
	return milter.RespContinue, nil
}

func (m *SimpleMilter) BodyChunk(chunk []byte, esmtpArgs *milter.Modifier) (milter.Response, error) {
	slog.InfoContext(context.Background(), "Body chunk: ", slog.String("chunk", string(chunk)))
	return milter.RespContinue, nil
}

func (m *SimpleMilter) Close() error {
	slog.InfoContext(context.Background(), "Connection closed")
	return nil
}

func (m *SimpleMilter) Abort(esmtpArgs *milter.Modifier) error {
	slog.InfoContext(context.Background(), "Abort")
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":17357")
	if err != nil {
		slog.ErrorContext(context.Background(), "Failed to listen on port 17357", "error", err)
		return
	}
	defer listener.Close()

	server := &milter.Server{
		NewMilter: func() milter.Milter {
			return &SimpleMilter{}
		},
		Actions:  milter.OptAction(milter.OptAddRcpt),
		Protocol: milter.OptProtocol(milter.OptNoHelo),
	}
	server.Serve(listener)
	defer server.Close()
}
