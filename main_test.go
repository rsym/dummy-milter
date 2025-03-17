package main

import (
	"net"
	"net/textproto"
	"testing"

	"github.com/emersion/go-milter"
)

func TestSimpleMilterConnect(t *testing.T) {
	m := &SimpleMilter{}
	host := "example.com"
	family := "tcp"
	var port uint16 = 25
	addr := net.ParseIP("192.0.2.1")
	mod := &milter.Modifier{}

	resp, err := m.Connect(host, family, port, addr, mod)
	if err != nil {
		t.Errorf("Connect returned an error: %v", err)
	}
	if resp != milter.RespContinue {
		t.Errorf("Connect returned %v, expected %v", resp, milter.RespContinue)
	}
}

func TestSimpleMilterHelo(t *testing.T) {
	m := &SimpleMilter{}
	name := "example.com"
	mod := &milter.Modifier{}

	resp, err := m.Helo(name, mod)
	if err != nil {
		t.Errorf("Helo returned an error: %v", err)
	}
	if resp != milter.RespContinue {
		t.Errorf("Helo returned %v, expected %v", resp, milter.RespContinue)
	}
}

func TestSimpleMilterMailFrom(t *testing.T) {
	m := &SimpleMilter{}
	from := "sender@example.com"
	mod := &milter.Modifier{}

	resp, err := m.MailFrom(from, mod)
	if err != nil {
		t.Errorf("MailFrom returned an error: %v", err)
	}
	if resp != milter.RespContinue {
		t.Errorf("MailFrom returned %v, expected %v", resp, milter.RespContinue)
	}
}

func TestSimpleMilterRcptTo(t *testing.T) {
	m := &SimpleMilter{}
	rcpt := "recipient@example.com"
	mod := &milter.Modifier{}

	resp, err := m.RcptTo(rcpt, mod)
	if err != nil {
		t.Errorf("RcptTo returned an error: %v", err)
	}
	if resp != milter.RespContinue {
		t.Errorf("RcptTo returned %v, expected %v", resp, milter.RespContinue)
	}
}

func TestSimpleMilterData(t *testing.T) {
	m := &SimpleMilter{}

	resp, err := m.Data()
	if err != nil {
		t.Errorf("Data returned an error: %v", err)
	}
	if resp != milter.RespContinue {
		t.Errorf("Data returned %v, expected %v", resp, milter.RespContinue)
	}
}

func TestSimpleMilterHeader(t *testing.T) {
	m := &SimpleMilter{}
	name := "Subject"
	value := "Test Email"
	mod := &milter.Modifier{}

	resp, err := m.Header(name, value, mod)
	if err != nil {
		t.Errorf("Header returned an error: %v", err)
	}
	if resp != milter.RespContinue {
		t.Errorf("Header returned %v, expected %v", resp, milter.RespContinue)
	}
}

func TestSimpleMilterHeaders(t *testing.T) {
	m := &SimpleMilter{}
	headers := textproto.MIMEHeader{
		"From":    []string{"sender@example.com"},
		"To":      []string{"recipient@example.com"},
		"Subject": []string{"Test Email"},
	}
	mod := &milter.Modifier{}

	resp, err := m.Headers(headers, mod)
	if err != nil {
		t.Errorf("Headers returned an error: %v", err)
	}
	if resp != milter.RespContinue {
		t.Errorf("Headers returned %v, expected %v", resp, milter.RespContinue)
	}
}

func TestSimpleMilterBody(t *testing.T) {
	m := &SimpleMilter{}
	mod := &milter.Modifier{}

	resp, err := m.Body(mod)
	if err != nil {
		t.Errorf("Body returned an error: %v", err)
	}
	if resp != milter.RespContinue {
		t.Errorf("Body returned %v, expected %v", resp, milter.RespContinue)
	}
}

func TestSimpleMilterBodyChunk(t *testing.T) {
	m := &SimpleMilter{}
	chunk := []byte("This is the body of the email.")
	mod := &milter.Modifier{}

	resp, err := m.BodyChunk(chunk, mod)
	if err != nil {
		t.Errorf("BodyChunk returned an error: %v", err)
	}
	if resp != milter.RespContinue {
		t.Errorf("BodyChunk returned %v, expected %v", resp, milter.RespContinue)
	}
}

func TestSimpleMilterClose(t *testing.T) {
	m := &SimpleMilter{}

	err := m.Close()
	if err != nil {
		t.Errorf("Close returned an error: %v", err)
	}
}

func TestSimpleMilterAbort(t *testing.T) {
	m := &SimpleMilter{}
	mod := &milter.Modifier{}

	err := m.Abort(mod)
	if err != nil {
		t.Errorf("Abort returned an error: %v", err)
	}
}

// Integration test that simulates a basic SMTP session
func TestSimpleMilterIntegration(t *testing.T) {
	m := &SimpleMilter{}
	mod := &milter.Modifier{}

	// Connect
	resp, err := m.Connect("example.com", "tcp", 25, net.ParseIP("192.0.2.1"), mod)
	if err != nil || resp != milter.RespContinue {
		t.Fatalf("Connect failed: %v, %v", resp, err)
	}

	// HELO
	resp, err = m.Helo("example.com", mod)
	if err != nil || resp != milter.RespContinue {
		t.Fatalf("Helo failed: %v, %v", resp, err)
	}

	// MAIL FROM
	resp, err = m.MailFrom("sender@example.com", mod)
	if err != nil || resp != milter.RespContinue {
		t.Fatalf("MailFrom failed: %v, %v", resp, err)
	}

	// RCPT TO
	resp, err = m.RcptTo("recipient@example.com", mod)
	if err != nil || resp != milter.RespContinue {
		t.Fatalf("RcptTo failed: %v, %v", resp, err)
	}

	// DATA
	resp, err = m.Data()
	if err != nil || resp != milter.RespContinue {
		t.Fatalf("Data failed: %v, %v", resp, err)
	}

	// Headers
	resp, err = m.Header("Subject", "Test Email", mod)
	if err != nil || resp != milter.RespContinue {
		t.Fatalf("Header failed: %v, %v", resp, err)
	}

	headers := textproto.MIMEHeader{
		"From":    []string{"sender@example.com"},
		"To":      []string{"recipient@example.com"},
		"Subject": []string{"Test Email"},
	}
	resp, err = m.Headers(headers, mod)
	if err != nil || resp != milter.RespContinue {
		t.Fatalf("Headers failed: %v, %v", resp, err)
	}

	// Body
	resp, err = m.Body(mod)
	if err != nil || resp != milter.RespContinue {
		t.Fatalf("Body failed: %v, %v", resp, err)
	}

	resp, err = m.BodyChunk([]byte("This is the body of the email."), mod)
	if err != nil || resp != milter.RespContinue {
		t.Fatalf("BodyChunk failed: %v, %v", resp, err)
	}

	// Close
	err = m.Close()
	if err != nil {
		t.Fatalf("Close failed: %v", err)
	}
}

// Test server creation
func TestServerCreation(t *testing.T) {
	server := &milter.Server{
		NewMilter: func() milter.Milter {
			return &SimpleMilter{}
		},
		Actions:  milter.OptAction(milter.OptAddRcpt),
		Protocol: milter.OptProtocol(0),
	}

	if server == nil {
		t.Fatal("Failed to create server")
	}

	// Check that NewMilter returns a SimpleMilter
	milter := server.NewMilter()
	if _, ok := milter.(*SimpleMilter); !ok {
		t.Error("NewMilter did not return a SimpleMilter")
	}
}
