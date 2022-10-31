package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/apex/log"
	"github.com/cakemarketing/go-common/v5/settings"
	"github.com/pborman/getopt"

	pb "go-calculatorchallenge/internal/gen/proto"

	"github.com/knetic/govaluate"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServiceServer
}

var (
	flagConfigPath  = "config"
	flagEnvironment = "local"
	// port            = settings.GetString("TCP_PORT")
	clients                              []string
	each_client_information              = make(map[string]string)
	client_how_many_calculations_perform = make(map[string]int)
	addition_result                      string
	subtration_result                    string
	multiplication_result                string
	division_result                      string
	expression_result                    string
	TotalExpression                      = 0
)

func init() {
	// log.Log = log.WithFields(log.Fields{
	// 	"methods": "Calculator",
	// })

	getopt.StringVarLong(&flagConfigPath, "config-directory", 'p', "path to the config file")
	getopt.StringVarLong(&flagEnvironment, "environment", 'e', "environment of running instance")

	getopt.Parse()

	if len(getopt.Args()) > 0 {
		flagEnvironment = getopt.Arg(0)
	}
}

func (s *server) Addition(ctx context.Context, req *pb.CalRequest) (*pb.CalResponse, error) {
	log.Info("Addition Service Called...")

	n1 := req.GetNumb1()
	n2 := req.GetNumb2()
	clientName := req.GetClientName()
	fmt.Println("client name:- ", clientName)

	// Add Clients into slice
	for i := 0; i <= len(clients); i++ {
		if func() int {
			for i, v := range clients {
				if v == clientName {
					return i
				}
			}
			return -1
		}() != -1 {
			// fmt.Println("present")
			break
		} else if func() int {
			for i, v := range clients {
				if v == clientName {
					return i
				}
			}
			return -1
		}() == -1 {
			// fmt.Println("not present")
			clients = append(clients, clientName)
			break
		} else if len(clients) == 0 {
			// fmt.Println("first entry")
			clients = append(clients, clientName)
			break
		}
	}

	result := n1 + n2
	fmt.Printf("Addition Result:- %d + %d = %v \n", n1, n2, result)
	TotalExpression++

	for index, _ := range clients {
		if clients[index] == clientName {
			client_how_many_calculations_perform[clientName] += 1
			addition_result = "Addition Service:- " + strconv.FormatInt(req.GetNumb1(), 10) + " + " + strconv.FormatInt(req.GetNumb2(), 10) + " = " + strconv.FormatInt(result, 10) + ", "
			each_client_information[clientName] += addition_result
		}
	}
	return &pb.CalResponse{
		Result: float64(result),
	}, nil
}

func (s *server) Subtraction(ctx context.Context, req *pb.CalRequest) (*pb.CalResponse, error) {
	log.Info("Subtraction Service Called...")

	n1 := req.GetNumb1()
	n2 := req.GetNumb2()
	clientName := req.GetClientName()
	fmt.Println("client name:- ", clientName)

	// Add Clients into slice
	for i := 0; i <= len(clients); i++ {
		if func() int {
			for i, v := range clients {
				if v == clientName {
					return i
				}
			}
			return -1
		}() != -1 {
			// fmt.Println("present")
			break
		} else if func() int {
			for i, v := range clients {
				if v == clientName {
					return i
				}
			}
			return -1
		}() == -1 {
			// fmt.Println("not present")
			clients = append(clients, clientName)
			break
		} else if len(clients) == 0 {
			// fmt.Println("first entry")
			clients = append(clients, clientName)
			break
		}
	}

	result := n1 - n2
	fmt.Printf("Subtraction Result:- %d - %d = %v \n", n1, n2, result)
	TotalExpression++

	for index, _ := range clients {
		if clients[index] == clientName {
			client_how_many_calculations_perform[clientName] += 1
			subtration_result = "Subtraction Service:- " + strconv.FormatInt(req.GetNumb1(), 10) + " - " + strconv.FormatInt(req.GetNumb2(), 10) + " = " + strconv.FormatInt(result, 10) + ", "
			each_client_information[clientName] += subtration_result
		}
	}

	return &pb.CalResponse{
		Result: float64(result),
	}, nil
}

func (s *server) Multiplication(ctx context.Context, req *pb.CalRequest) (*pb.CalResponse, error) {
	log.Info("Multiplication Service Called...")

	n1 := req.GetNumb1()
	n2 := req.GetNumb2()
	clientName := req.GetClientName()
	fmt.Println("client name:- ", clientName)

	// Add Clients into slice
	for i := 0; i <= len(clients); i++ {
		if func() int {
			for i, v := range clients {
				if v == clientName {
					return i
				}
			}
			return -1
		}() != -1 {
			// fmt.Println("present")
			break
		} else if func() int {
			for i, v := range clients {
				if v == clientName {
					return i
				}
			}
			return -1
		}() == -1 {
			// fmt.Println("not present")
			clients = append(clients, clientName)
			break
		} else if len(clients) == 0 {
			// fmt.Println("first entry")
			clients = append(clients, clientName)
			break
		}
	}

	result := n1 * n2
	fmt.Printf("Multiplication Result:- %d * %d = %v \n", n1, n2, result)
	TotalExpression++

	for index, _ := range clients {
		if clients[index] == clientName {
			client_how_many_calculations_perform[clientName] += 1
			multiplication_result = "Multiplication Service:- " + strconv.FormatInt(req.GetNumb1(), 10) + " * " + strconv.FormatInt(req.GetNumb2(), 10) + " = " + strconv.FormatInt(result, 10) + ", "
			each_client_information[clientName] += multiplication_result
		}
	}

	return &pb.CalResponse{
		Result: float64(result),
	}, nil
}

func (s *server) Division(ctx context.Context, req *pb.CalRequest) (*pb.CalResponse, error) {
	log.Info("Division Service Called...")

	n1 := req.GetNumb1()
	n2 := req.GetNumb2()
	clientName := req.GetClientName()
	fmt.Println("client name:- ", clientName)

	// Add Clients into slice
	for i := 0; i <= len(clients); i++ {
		if func() int {
			for i, v := range clients {
				if v == clientName {
					return i
				}
			}
			return -1
		}() != -1 {
			// fmt.Println("present")
			break
		} else if func() int {
			for i, v := range clients {
				if v == clientName {
					return i
				}
			}
			return -1
		}() == -1 {
			// fmt.Println("not present")
			clients = append(clients, clientName)
			break
		} else if len(clients) == 0 {
			// fmt.Println("first entry")
			clients = append(clients, clientName)
			break
		}
	}

	result := n1 / n2
	fmt.Printf("Division Result:- %d / %d = %v \n", n1, n2, result)
	TotalExpression++

	for index, _ := range clients {
		if clients[index] == clientName {
			client_how_many_calculations_perform[clientName] += 1
			division_result = "Division Service:- " + strconv.FormatInt(req.GetNumb1(), 10) + " / " + strconv.FormatInt(req.GetNumb2(), 10) + " = " + strconv.FormatInt(result, 10) + ", "
			each_client_information[clientName] += division_result
		}
	}

	return &pb.CalResponse{
		Result: float64(result),
	}, nil
}

func (s *server) ExpressionCalculator(stream pb.CalculatorService_ExpressionCalculatorServer) error {
	log.Info("Expression Service Method called...")

	req, err := stream.Recv()
	if err != nil {
		log.WithError(err).Error("Error while Receving request from client")
	}
	// log.WithField("Request", req).Info("Request from Client")

	inputStringExpression := req.GetExpression()
	clientName := req.GetClientName()
	fmt.Println("client name:- ", clientName)

	// Add Clients into slice
	for i := 0; i <= len(clients); i++ {
		if func() int {
			for i, v := range clients {
				if v == clientName {
					return i
				}
			}
			return -1
		}() != -1 {
			// fmt.Println("present")
			break
		} else if func() int {
			for i, v := range clients {
				if v == clientName {
					return i
				}
			}
			return -1
		}() == -1 {
			// fmt.Println("not present")
			clients = append(clients, clientName)
			break
		} else if len(clients) == 0 {
			// fmt.Println("first entry")
			clients = append(clients, clientName)
			break
		}
	}

	// log.WithField("Request", inputStringExpression).Info("GetExpression from Client Request")
	expression, err := govaluate.NewEvaluableExpression(inputStringExpression)
	if err != nil {
		log.WithError(err).Error("Failed to Handle Expression")
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		log.WithError(err).Error("Failed to Evaluate Expression")
	}
	// log.WithField("Result", result).Info("Expression Result")

	val, ok := result.(float64)
	if !ok {
		return nil
	}

	for {
		// log.Info("Waiting to receive more data")
		_, err := stream.Recv()
		// log.WithField("Request", req).Info("Recv request from client")
		if err == io.EOF {
			// log.Info("NO more data")
			break
		}
		if err != nil {
			log.Fatalf("Error while Recv Calculator %v", err)
		}
	}
	resp := &pb.CalExpressionResponse{
		Result: val,
	}
	err = stream.SendAndClose(resp)
	if err != nil {
		log.WithField("Error", err).Info("Cannot send Response")
	}
	TotalExpression++

	for index, _ := range clients {
		if clients[index] == clientName {
			client_how_many_calculations_perform[clientName] += 1
			expression_result = "Expression Service:- " + inputStringExpression + " = " + strconv.FormatInt(int64(val), 10) + ", "
			each_client_information[clientName] += expression_result
		}
	}

	fmt.Printf("Expression Result:- %s = %v \n", inputStringExpression, result)

	return nil
}

func handleConnection(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	log.Info("Server Side...")

	// Channels used for signals of key press by client i.e: Ctrl + c
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		done <- true
		if os.Interrupt == <-sigs {
			log.Info(`We are here in Goroutine...after Receiving a "CTRL + C" signal`)
			fmt.Println("All clients:- ", clients)
			fmt.Println("client_how_many_calculations_perform:- ", client_how_many_calculations_perform)
			fmt.Println("each_client_information:- ", each_client_information)
			fmt.Printf("Total expressions on server are %v by all clients", TotalExpression)
		}
	}()

	// parse the environment file
	settings.SetConfigName(flagEnvironment)
	settings.AddConfigPath(flagConfigPath)
	if err := settings.ReadInConfig(); err != nil {
		fmt.Printf("Could not parse configuration file '%s/%s': %v", flagConfigPath, flagEnvironment, err)
		return
	}

	// gRPC client Listening
	go func() {
		log.Info("gRPC is listening on port " + settings.GetString("TCP_PORT_GRPC"))
		listener, err := net.Listen("tcp", ":"+settings.GetString("TCP_PORT_GRPC"))
		if err != nil {
			log.WithError(err).Error("Failed to Listen")
		}
		defer listener.Close()
		grpcServer := grpc.NewServer()
		pb.RegisterCalculatorServiceServer(grpcServer, &server{})

		if err = grpcServer.Serve(listener); err != nil {
			log.WithError(err).Error("Error while Serve")
		}
	}()

	// Telnet client Listening
	log.Info("Telnet is listening " + settings.GetString("TCP_PORT_TELNET"))
	ln, err := net.Listen("tcp", ":"+settings.GetString("TCP_PORT_TELNET"))
	if err != nil {
		log.WithError(err).Error("Failed to Listen telnet connection")
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.WithError(err).Error("Failed")
		}
		go handleConnection(conn)
	}
}
