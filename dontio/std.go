package dontio

import (
	"context"
	"fmt"
	"io"
	"os"
)

type ctxkey string

var stdkey = "std"

type notInCtx string // what wasn't present

func (what notInCtx) Error() string {
	return fmt.Sprintf("%s was not found in context", string(what))
}

type Std struct {
	Out io.Writer
	Err io.Writer
	In  io.Reader
}

func AddStdToContext(ctx context.Context, s *Std) context.Context {
	return AddNamedStdToContext(ctx, stdkey, s)
}

func AddNamedStdToContext(ctx context.Context, name string, s *Std) context.Context {
	return context.WithValue(ctx, ctxkey(name), s)
}

func StdFromContext(ctx context.Context) (*Std, error) {
	return NamedStdFromContext(ctx, stdkey)
}

func NamedStdFromContext(ctx context.Context, name string) (*Std, error) {
	v := ctx.Value(stdkey)
	if v == nil {
		return nil, notInCtx("stdio")
	}
	return v.(*Std), nil
}

func StdIo() Std {
	return Std{
		Out: os.Stdout,
		Err: os.Stderr,
		In:  os.Stdin,
	}
}

func Closed() Std {
	r := AlwaysErrReader{io.ErrClosedPipe}
	w := AlwaysErrWriter{io.ErrClosedPipe}
	return Std{
		Out: w,
		Err: w,
		In:  r,
	}
}

func (s Std) WriteLineOut(msg string, v ...any) {
	fmt.Fprintf(s.Out, msg+"\n", v...)
}

func (s Std) WriteLineErr(msg string, v ...any) {
	fmt.Fprintf(s.Err, msg+"\n", v...)
}

func WriteLineOut(ctx context.Context, tpl string, v ...any) error {
	stdio, stdErr := StdFromContext(ctx)
	if stdErr != nil {
		return stdErr
	}
	stdio.WriteLineOut(tpl, v...)
	return nil
}

func WriteLineErr(ctx context.Context, tpl string, v ...any) error {
	stdio, stdErr := StdFromContext(ctx)
	if stdErr != nil {
		return stdErr
	}
	stdio.WriteLineErr(tpl, v...)
	return nil
}

type AlwaysErrReader struct {
	Err error
}

func (this AlwaysErrReader) Read([]byte) (int, error) {
	return 0, this.Err
}

type AlwaysErrWriter struct {
	Err error
}

func (this AlwaysErrWriter) Write([]byte) (int, error) {
	return 0, this.Err
}
