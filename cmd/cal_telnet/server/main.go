//go:build tagforspacing
// +build tagforspacing

package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"

	"go-calculatorchallenge/calculator"
	clients_connection "go-calculatorchallenge/cmd/cal_telnet/client"
	stores_postgres "go-calculatorchallenge/stores/postgres"
	stores_redis "go-calculatorchallenge/stores/redis"

	"github.com/apex/log"
	"github.com/cakemarketing/go-common/v5/settings"
	"github.com/pborman/getopt"
)

var (
	flagConfigPath  = "config"
	flagEnvironment = "local"
	TotalExpression = 0
	listener        net.Listener
	quit            chan interface{}
	wg              sync.WaitGroup
	err             error
)

func init() {

	// Take the variable values from flag in cmd
	getopt.StringVarLong(&flagConfigPath, "config-directory", 'p', "path to the config file")
	getopt.StringVarLong(&flagEnvironment, "environment", 'e', "environment of running instance")

	getopt.Parse()

	if len(getopt.Args()) > 0 {
		flagEnvironment = getopt.Arg(0)
	}

	// Create the database connection and table
	stores_postgres.DataMigration()
}

func main() {

	quit = make(chan interface{})

	// Channels used for Signals i.e: Ctrl + C
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Parse the environment file
	settings.SetConfigName(flagEnvironment)
	settings.AddConfigPath(flagConfigPath)
	if err := settings.ReadInConfig(); err != nil {
		fmt.Printf("Could not parse configuration file '%s/%s': %v", flagConfigPath, flagEnvironment, err)
		return
	}

	// Check the all settings
	log.WithFields(log.Fields{
		"settings": fmt.Sprintf("%#v", settings.AllSettings()),
	}).Info("Initializing configuration")

	// Start Redis client
	_ = stores_redis.GetRedisClient()

	// Telnet client is Listening
	log.Info("Telnet client is listening on port " + settings.GetString("TCP_PORT_TELNET"))
	listener, err = net.Listen("tcp", ":"+settings.GetString("TCP_PORT_TELNET"))
	if err != nil {
		log.Info("Failed to listen")
	}

	wg.Add(1)
	go Serve()
	// Stop()

	// Goroutine used for waiting CTRL+C signal
	wg.Add(1)
	go func() {
		// Exiting the server gracefully with CTRL+C
		if os.Interrupt == <-sigs {
			fmt.Println("\nI am here in goroutine, after pressing CTRL+C signal...")
			fmt.Printf("Total %v expressions have performed on server by all clients\n", calculator.TotalExpression)
			log.Info("Server Shutdown Gracefully...Bye")
			close(quit)
			listener.Close()
			// wg.Wait()
			wg.Done()
		}
	}()
	wg.Wait()
	fmt.Println("No. of Goroutines:- ", runtime.NumGoroutine())
}

func Serve() {
	defer wg.Done()

	for {
		conn, err := listener.Accept()
		fmt.Println("I accepted...") // just used for checking purpose
		if err != nil {
			// log.WithError(err).Error("Error while accepting connection from client")
			// continue
			select {
			case <-quit:
				return
			default:
				log.WithError(err).Error("accept error")
			}
		} else {
			wg.Add(1)
			go func() {
				clients_connection.HandleConnection(conn)
				wg.Done()
			}()
			// wg.Wait()
		}
	}
}

// func Stop() {
// 	// time.Sleep(60 * time.Second)
// 	fmt.Println("I am in stop function...")
// 	close(quit)
// 	listener.Close()
// 	wg.Wait()
// }
