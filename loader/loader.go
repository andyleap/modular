package loader

import (
	"io"
	"net/rpc"
	"os/exec"
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
	Process *exec.Cmd
	*rpc.Client
}

func Load(name string, arg ...string) (*Module, error) {
	proc := exec.Command(name, arg...)
	in, err := proc.StdinPipe()
	if err != nil {
		return nil, err
	}
	out, err := proc.StdoutPipe()
	if err != nil {
		return nil, err
	}
	pipe := &ReadWritePipe{in, out}
	proc.Start()
	client := rpc.NewClient(pipe)
	return &Module{pipe, proc, client}, nil
}

func (m *Module) Close() {
	m.Client.Close()
	m.Process.Wait()
}
