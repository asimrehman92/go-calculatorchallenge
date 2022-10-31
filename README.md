## Run the code with Telnet
First go to "cmd" directory, and then go to the "cal_telnet" directory, and then go to the "server" directory and run the  file

## Run the code with gRPC
First go to "cmd" directory, and then go to the "cal_grpc" directory, 
and then go to the "server" directory and run the file,
and then also go to the "client" directory and run the file.

## Directory path
cd C:\Users\rehmaasi\project\go-calculatorchallenge


## For creating go mod file anywhere
go mod init projectName or path

## For run the telnet
telent localhost 9999

## For Redis Client
redis-cli run on the cmd

## For Redis Server
redis-server run on the cmd

## Server Shutdown Gracefully (waiting for client connections to shut down)
I write the code of shutdown the server gracefully (its working like:- server is waiting for all client connections to shut down then its closed shutdown gracefully)

## Go Testing Files
Go testing files are always located in the same folder, or package, where the code they are testing resides
so, go to the same folder and run the command "go test" or "go test -v"
