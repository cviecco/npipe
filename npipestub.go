// +build !windows
package npipe

import (
	"errors"
	//"fmt"
	"net"
	"time"
)

/// PipeAddr is a Helper stub
type PipeAddr string

type PipeConn struct {
	addr PipeAddr
}
type PipeListener struct {
	addr PipeAddr
}

var _ net.Conn = (*PipeConn)(nil)
var _ net.Listener = (*PipeListener)(nil)

var ErrNotImplemented = errors.New("Not implemented")

// PipeAddr as a net.Addr
func (a PipeAddr) Network() string { return "pipe" }

func (a PipeAddr) String() string {
	return string(a)
}

// PipeConn as a net.Conn
func Dial(address string) (*PipeConn, error) {
	return nil, ErrNotImplemented
}
func DialTimeout(address string, timeout time.Duration) (*PipeConn, error) {
	return nil, ErrNotImplemented
}
func (c *PipeConn) Read(b []byte) (int, error) {
	return 0, ErrNotImplemented
}
func (c *PipeConn) Write(b []byte) (int, error) {
	return 0, ErrNotImplemented
}
func (c *PipeConn) Close() error {
	return ErrNotImplemented
}
func (c *PipeConn) LocalAddr() net.Addr {
	return c.addr
}
func (c *PipeConn) RemoteAddr() net.Addr {
	return c.addr
}
func (c *PipeConn) SetDeadline(t time.Time) error {
	return ErrNotImplemented
}
func (c *PipeConn) SetReadDeadline(t time.Time) error {
	return ErrNotImplemented
}
func (c *PipeConn) SetWriteDeadline(t time.Time) error {
	return ErrNotImplemented
}

// PipeListener as a net.Listener
func Listen(address string) (*PipeListener, error) {
	return nil, ErrNotImplemented
}
func (l *PipeListener) Accept() (net.Conn, error) {
	return nil, ErrNotImplemented
}

// Addr returns the listener's network address, a PipeAddr.
func (l *PipeListener) Addr() net.Addr { return l.addr }

func (l *PipeListener) Close() error {
	return ErrNotImplemented
}
