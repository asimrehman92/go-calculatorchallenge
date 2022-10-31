package main

// Useful imports
import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	pb "go-calculatorchallenge/internal/gen/proto"

	"github.com/apex/log"
	"google.golang.org/grpc"
)

// Global Variables
var (
	user  string
	err   error
	count int
	sigs  = make(chan os.Signal, 1)

	clients                              []string
	Client_used_servies                  []string
	client_information                   = make(map[string]string)
	client_how_many_calculations_perform = make(map[string]int)
	addition_result                      string
	subtration_result                    string
	multiplication_result                string
	division_result                      string
	expression_result                    string
	// port                                 = settings.GetString("TCP_PORT")
	// address                              = settings.GetString("ADDRESS")
)

// when we run the client file... first of all init() called
func init() {
	// log.Log = log.WithFields(log.Fields{
	// 	"methods": "Calculator",
	// })
}

// call_Addition() is perform addition by taking two numbers
func call_Addition(c pb.CalculatorServiceClient, numb1, numb2, name string) {
	log.Info("Calling Calculator Api")

	number1, _ := strconv.Atoi(numb1)
	number2, _ := strconv.Atoi(numb2)

	// fmt.Println("N1 in integer ", "|", int64(number1), "|", err1, "|", reflect.TypeOf(number1))
	// fmt.Println("N2 in integer ", "|", int64(number2), "|", err2, "|", reflect.TypeOf(number2))

	req := &pb.CalRequest{
		ClientName: name,
		Numb1:      int64(number1),
		Numb2:      int64(number2),
	}
	res, err := c.Addition(context.Background(), req)
	if err != nil {
		log.WithError(err).Error("call_Addition err")
	}
	log.Info("Response from Server side")
	addition_result = "Addition Service:- " + numb1 + " + " + numb2 + " = " + strconv.FormatFloat(res.GetResult(), 'f', -1, 32) + ", "
	Client_used_servies = append(Client_used_servies, addition_result)
	log.WithField(numb1+" + "+numb2, res.GetResult()).Info("Server Response: Addition Service:- ")
}

// call_Subtraction() is perform subtraction by taking two numbers
func call_Subtraction(c pb.CalculatorServiceClient, numb1, numb2, name string) {
	log.Info("Calling Subtraction Api")

	number1, _ := strconv.Atoi(numb1)
	number2, _ := strconv.Atoi(numb2)

	// fmt.Println("N1 in integer ", "|", int64(number1), "|", err1, "|", reflect.TypeOf(number1))
	// fmt.Println("N2 in integer ", "|", int64(number2), "|", err2, "|", reflect.TypeOf(number2))

	req := &pb.CalRequest{
		ClientName: name,
		Numb1:      int64(number1),
		Numb2:      int64(number2),
	}
	res, err := c.Subtraction(context.Background(), req)
	if err != nil {
		log.WithError(err).Error("call_Subtraction err")
	}
	subtration_result = "Subtraction Service:- " + numb1 + " - " + numb2 + " = " + strconv.FormatFloat(res.GetResult(), 'f', -1, 32) + ", "
	Client_used_servies = append(Client_used_servies, subtration_result)
	log.Info("Response from Server side")
	log.WithField(numb1+" - "+numb2, res.GetResult()).Info("Server Response: Subtraction Service:- ")
}

// call_Multiplication() is perform multiplication by taking two numbers
func call_Multiplication(c pb.CalculatorServiceClient, numb1, numb2, name string) {
	log.Info("Calling Multiplication Api")

	number1, _ := strconv.Atoi(numb1)
	number2, _ := strconv.Atoi(numb2)

	// fmt.Println("N1 in integer ", "|", int64(number1), "|", err1, "|", reflect.TypeOf(number1))
	// fmt.Println("N2 in integer ", "|", int64(number2), "|", err2, "|", reflect.TypeOf(number2))

	req := &pb.CalRequest{
		ClientName: name,
		Numb1:      int64(number1),
		Numb2:      int64(number2),
	}
	res, err := c.Multiplication(context.Background(), req)
	if err != nil {
		log.WithError(err).Error("call_Multiplication err")
	}
	multiplication_result = "Multiplication Service:- " + numb1 + " * " + numb2 + " = " + strconv.FormatFloat(res.GetResult(), 'f', -1, 32) + ", "
	Client_used_servies = append(Client_used_servies, multiplication_result)
	log.Info("Response from Server side")
	log.WithField(numb1+" * "+numb2, res.GetResult()).Info("Server Response: Multiplication Service:- ")
}

// call_Division() is perform division by taking two numbers
func call_Division(c pb.CalculatorServiceClient, numb1, numb2, name string) {
	log.Info("Calling Division Api")

	number1, _ := strconv.Atoi(numb1)
	number2, _ := strconv.Atoi(numb2)

	// fmt.Println("N1 in integer ", "|", int64(number1), "|", err1, "|", reflect.TypeOf(number1))
	// fmt.Println("N2 in integer ", "|", int64(number2), "|", err2, "|", reflect.TypeOf(number2))

	req := &pb.CalRequest{
		ClientName: name,
		Numb1:      int64(number1),
		Numb2:      int64(number2),
	}
	res, err := c.Division(context.Background(), req)
	if err != nil {
		log.WithError(err).Error("call_Division err")
	}
	division_result = "Division Service:- " + numb1 + " / " + numb2 + " = " + strconv.FormatFloat(res.GetResult(), 'f', -1, 32) + ", "
	Client_used_servies = append(Client_used_servies, division_result)
	log.Info("Response from Server side")
	log.WithField(numb1+" / "+numb2, res.GetResult()).Info("Server Response: Division Service:- ")
}

// Call_ExpressionCalculator() is perform the calculation of any expression by taking expression
func Call_ExpressionCalculator(c pb.CalculatorServiceClient, expr, name string) {
	log.Info("Calling Calculator Api")

	// Give Expression by using the Flag
	// expressionPtr = flag.String("exp", "(2 + 2) / 10", "give expression for calculation")
	// flag.Parse()

	stream, err := c.ExpressionCalculator(context.Background())
	if err != nil {
		log.Fatalf("Call Calculator err %v", err)
	}
	req := pb.CalExpressionRequest{
		ClientName: name,
		Expression: expr,
	}

	err = stream.Send(&req)
	if err != nil {
		log.WithError(err).Error("Send calculator Request err")
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.WithError(err).Error("Receive calculator Response err")
	}
	expression_result = "Expression Service:- " + expr + " = " + strconv.FormatFloat(resp.GetResult(), 'f', -1, 32) + ", "
	Client_used_servies = append(Client_used_servies, expression_result)
	// fmt.Printf("Server Response:  %s = %v \n", expr, resp.GetResult())
	log.WithField(expr, resp.GetResult()).Info("Server Response:")
}

// Instruction() give the instruction to client and take values by ReadString from console
func Instruction() {
	reader := bufio.NewReader(os.Stdin)
	log.Info("Enter your name: ")
	user, err = reader.ReadString('\n')
	if err != nil {
		log.WithError(err).Error("ReadString err")
	}
	user = strings.Replace(user, "\r\n", "", -1)
	clients = append(clients, user)
	// fmt.Println("These are clients: ", clients)
	// fmt.Println("length:- ", len(clients))
	// fmt.Println("This is Map1:- ", client_information)

	// log.Info("I think you want to Connection with Server: ")
	log.Info(`Say "yes" for Connection otherwise Say "no"`)
	inputRequest, err := reader.ReadString('\n')
	if err != nil {
		log.WithError(err).Error("ReadString err")
	}
	inputRequest = strings.Replace(inputRequest, "\r\n", "", -1)
	if inputRequest == "yes" || inputRequest == "YES" || inputRequest == "Yes" {
		conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
		if err != nil {
			log.WithError(err).Error("Error while Dial with Server")
		}
		log.Info("Client Connected to Server")
		defer conn.Close()
		client := pb.NewCalculatorServiceClient(conn)
		for {
			log.Info("--------------------------------------------------------")
			log.Info("In our Client Server Calculator we have these services: ")
			log.Info("1. Addition Service")
			log.Info("2. Subtraction Service")
			log.Info("3. Multiplication Service")
			log.Info("4. Division Service")
			log.Info("5. Expression Service")
			log.Info("6. Connection Close")

			log.Info("Choose the any Calculator Service: ")
			text, err := reader.ReadString('\n')
			if err != nil {
				log.WithError(err).Error("ReadString err")
			}
			text = strings.Replace(text, "\r\n", "", -1)
			if text == "1" {
				log.Info("Enter Two Numbers for Addition: ")
				log.Info("Enter Number 1: ")
				text1, err := reader.ReadString('\n')
				if err != nil {
					log.WithError(err).Error("ReadString err")
				}
				text1 = strings.Replace(text1, "\r\n", "", -1)
				log.Info("Enter Number 2: ")
				text2, err := reader.ReadString('\n')
				if err != nil {
					log.WithError(err).Error("ReadString err")
				}
				text2 = strings.Replace(text2, "\r\n", "", -1)
				call_Addition(client, text1, text2, user)
				count++
			} else if text == "2" {
				log.Info("Enter Two Numbers for Subtraction: ")
				log.Info("Enter Number 1: ")
				text1, err := reader.ReadString('\n')
				if err != nil {
					log.WithError(err).Error("ReadString err")
				}
				text1 = strings.Replace(text1, "\r\n", "", -1)
				log.Info("Enter Number 2: ")
				text2, err := reader.ReadString('\n')
				if err != nil {
					log.WithError(err).Error("ReadString err")
				}
				text2 = strings.Replace(text2, "\r\n", "", -1)
				call_Subtraction(client, text1, text2, user)
				count++
			} else if text == "3" {
				log.Info("Enter Two Numbers for Multiplication: ")
				log.Info("Enter Number 1: ")
				text1, err := reader.ReadString('\n')
				if err != nil {
					log.WithError(err).Error("ReadString err")
				}
				text1 = strings.Replace(text1, "\r\n", "", -1)
				log.Info("Enter Number 2: ")
				text2, err := reader.ReadString('\n')
				if err != nil {
					log.WithError(err).Error("ReadString err")
				}
				text2 = strings.Replace(text2, "\r\n", "", -1)
				call_Multiplication(client, text1, text2, user)
				count++
			} else if text == "4" {
				log.Info("Enter Two Numbers for Division: ")
				log.Info("Enter Number 1: ")
				text1, err := reader.ReadString('\n')
				if err != nil {
					log.WithError(err).Error("ReadString err")
				}
				text1 = strings.Replace(text1, "\r\n", "", -1)
				log.Info("Enter Number 2: ")
				text2, err := reader.ReadString('\n')
				if err != nil {
					log.WithError(err).Error("ReadString err")
				}
				text2 = strings.Replace(text2, "\r\n", "", -1)
				call_Division(client, text1, text2, user)
				count++
			} else if text == "5" {
				log.Info("Enter any Expression for Calculation: ")
				text, err := reader.ReadString('\n')
				if err != nil {
					log.WithError(err).Error("ReadString err")
				}
				text = strings.Replace(text, "\r\n", "", -1)
				Call_ExpressionCalculator(client, text, user)
				count++
			} else if text == "6" {
				log.Info("Close the connection on Client Request....Good Bye")
				break
			} else {
				log.Info("Wrong Input! Please Enter the option again...")
			}
		}
	} else if inputRequest == "no" || inputRequest == "NO" || inputRequest == "No" {
		log.Info("Close the connection on your request....Good Bye")
	} else {
		log.Info("Wrong Input!!!")
	}

}

func main() {
	io.WriteString(os.Stdout, "Client Side...\n")
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Error("Error while Dial with Server")
	}
	log.Info("Client Connected to Server")
	defer conn.Close()

	sigs = make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// sig := <-sigs
		// fmt.Println("this is called...1", sig)
		done <- true
		if os.Interrupt == <-sigs {
			log.Info("We are here in Goroutine...Receive a signal")

		}
	}()

	Instruction()
	for index, _ := range clients {
		if clients[index] == user {
			log.Info("--------------------------------------------")
			log.Info(`  "Client Information"  `)
			client_how_many_calculations_perform[user] = count
			for i := 0; i < len(Client_used_servies); i++ {
				client_information[user] += Client_used_servies[i]
			}
		}
	}
	fmt.Println(user, " performs ", count, "services")
	fmt.Println("This client has performed how many calculations:- ", client_how_many_calculations_perform)
	fmt.Println("This is the client information of using different services: ", client_information)
	fmt.Printf("These all services have used by %v:- %v", user, Client_used_servies)

}
