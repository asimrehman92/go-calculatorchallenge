package cal_telnet

import (
	"bufio"
	"io"
	"net"
	"strings"

	"go-calculatorchallenge/calculator"
	stores_redis "go-calculatorchallenge/stores/redis"

	"github.com/apex/log"
	"github.com/go-redis/redis"
)

var (
	clients     []string
	redisClient *redis.Client
)

func HandleConnection(c net.Conn) {
	// defer c.Close()

	// Redis Client
	redisClient = stores_redis.GetRedisClient()

	reader := bufio.NewReader(c)

	io.WriteString(c, "Client connected to the server\n")
	io.WriteString(c, "\rEnter your name:- ")
	clientName, _ := reader.ReadString('\n')
	clientName = strings.Replace(clientName, "\r\n", "", -1)
	clients = append(clients, clientName)

	for {
		io.WriteString(c, "\n\rIn our Client Server Calculator we have these services: \n")
		io.WriteString(c, "\r1. Addition Service\n")

		io.WriteString(c, "\r2. Subtraction Service\n")
		io.WriteString(c, "\r3. Multiplication Service\n")
		io.WriteString(c, "\r4. Division Service\n")
		io.WriteString(c, "\r5. Expression Service\n")
		io.WriteString(c, "\r6. Display the Previous Redis Data\n")
		io.WriteString(c, "\r7. Close the client Connection\n")
		io.WriteString(c, "\rChoose the any Calculator Service: ")
		option, err := reader.ReadString('\n')
		if err != nil {
			log.WithError(err).Error("ReadString err")
		}
		option = strings.Replace(option, "\r\n", "", -1)
		if option == "1" {
			io.WriteString(c, "\nEnter Two Numbers for Addition: \n")
			io.WriteString(c, "\rEnter Number 1: ")
			text1, err := reader.ReadString('\n')
			if err != nil {
				log.WithError(err).Error("ReadString err")
			}
			text1 = strings.Replace(text1, "\r\n", "", -1)
			io.WriteString(c, "Enter Number 2: ")
			text2, err := reader.ReadString('\n')
			if err != nil {
				log.WithError(err).Error("ReadString err")
			}
			text2 = strings.Replace(text2, "\r\n", "", -1)
			calculator.AdditionService(c, text1, text2, clientName, redisClient)

		} else if option == "2" {
			io.WriteString(c, "\nEnter Two Numbers for Subtraction: \n")
			io.WriteString(c, "\rEnter Number 1: ")
			text1, err := reader.ReadString('\n')
			if err != nil {
				log.WithError(err).Error("ReadString err")
			}
			text1 = strings.Replace(text1, "\r\n", "", -1)
			io.WriteString(c, "Enter Number 2: ")
			text2, err := reader.ReadString('\n')
			if err != nil {
				log.WithError(err).Error("ReadString err")
			}
			text2 = strings.Replace(text2, "\r\n", "", -1)
			calculator.SubtractionService(c, text1, text2, clientName, redisClient)
		} else if option == "3" {
			io.WriteString(c, "\nEnter Two Numbers for Multiplication: \n")
			io.WriteString(c, "\rEnter Number 1: ")
			text1, err := reader.ReadString('\n')
			if err != nil {
				log.WithError(err).Error("ReadString err")
			}
			text1 = strings.Replace(text1, "\r\n", "", -1)
			io.WriteString(c, "Enter Number 2: ")
			text2, err := reader.ReadString('\n')
			if err != nil {
				log.WithError(err).Error("ReadString err")
			}
			text2 = strings.Replace(text2, "\r\n", "", -1)
			calculator.MultiplicationService(c, text1, text2, clientName, redisClient)
		} else if option == "4" {
			io.WriteString(c, "\nEnter Two Numbers for Division: \n")
			io.WriteString(c, "\rEnter Number 1: ")
			text1, err := reader.ReadString('\n')
			if err != nil {
				log.WithError(err).Error("ReadString err")
			}
			text1 = strings.Replace(text1, "\r\n", "", -1)
			io.WriteString(c, "Enter Number 2: ")
			text2, err := reader.ReadString('\n')
			if err != nil {
				log.WithError(err).Error("ReadString err")
			}
			text2 = strings.Replace(text2, "\r\n", "", -1)
			calculator.DivisionService(c, text1, text2, clientName, redisClient)
		} else if option == "5" {
			io.WriteString(c, "\nEnter any Expression for Calculation: ")
			text, err := reader.ReadString('\n')
			if err != nil {
				log.WithError(err).Error("ReadString err")
			}
			text = strings.Replace(text, "\r\n", "", -1)
			calculator.ExpressionService(c, text, clientName, redisClient)
		} else if option == "6" {
			stores_redis.DisplayDataRedis(clientName, redisClient)
		} else if option == "7" {
			log.Info("Good Bye! " + clientName + " client connection closed")
			c.Close()
			break
		} else {
			io.WriteString(c, "\nWrong Input! Please Enter the option again...\n")
		}
	}
}
