package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jennxsierra/echo-chamber/internal/logger"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return "unknown"
}

func Start(port string) {
	ip := getLocalIP()
	startMessage := fmt.Sprintf("Echo Chamber server started at [%s:%s]", ip, port)
	logger.LogToTerminal(startMessage)
	logger.LogMessage(startMessage)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	defer func() {
		listener.Close()
		fmt.Println()
		stopMessage := "Echo Chamber server has been stopped."
		logger.LogToTerminal(stopMessage)
		logger.LogMessage(stopMessage)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Err.Error() == "use of closed network connection" {
				break
			}
			logger.LogToTerminal(fmt.Sprintf("Error accepting connection: %v", err))
			logger.LogMessage(fmt.Sprintf("Error accepting connection: %v", err))
			continue
		}
		go handleConnection(conn)
	}
}
