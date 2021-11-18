package bootstrap

import (
	"os"
	"os/signal"
	"syscall"

	"go.danale.net/be/be/go/logx"
)

func Run() error {
	p, err := setupProvider()
	if err != nil {
		logx.X.Error(err)
		return err
	}
	defer p.Close()
	pch := p.Run()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-ch:
		logx.X.With("signal", sig).Info("receive exit signal")
		return nil
	case err := <-pch:
		logx.X.With("error", err).Error("program stop")
		return err
	}
}
