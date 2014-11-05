package module

import (
	"io"
	"net/rpc"
	"os"
)

type ReadWritePipe struct {
        in io.WriteCloser
        out io.ReadCloser
}

func (rwp *ReadWritePipe) Close() error {
        rwp.in.Close()
        rwp.out.Close()
	return nil
}

func (rwp *ReadWritePipe) Read(p []byte) (n int, err error) {
	return rwp.out.Read(p)
}

func (rwp *ReadWritePipe) Write(p []byte) (n int, err error) {
	return rwp.in.Write(p)
}

type Module struct {
	Pipe *ReadWritePipe
	*rpc.Server
}

func New() *Module {
	pipe := &ReadWritePipe{os.Stdout, os.Stdin}
	server := rpc.NewServer()
	return &Module{pipe, server}
}

func (m *Module) Serve() {
	m.Server.ServeConn(m.Pipe)
}
