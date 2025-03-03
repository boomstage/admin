package gmiddleware

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

/*
业务服务注册:
h.SetCustomSignalWaiter(gmiddleware.WaitSignal)
*/

// WaitSignal custom wait signal
func WaitSignal(errCh chan error) error {
	signalToNotify := []os.Signal{syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM}
	if signal.Ignored(syscall.SIGHUP) {
		signalToNotify = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, signalToNotify...)

	select {
	case sig := <-signals:
		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM: // 目前k8s环境关闭发送信号: SIGTERM
			log.Printf("Received signal, sig:%s\n", sig.String())
			// graceful shutdown
			return nil
		}
	case err := <-errCh:
		// error occurs, exit immediately
		return err
	}

	return nil
}
