## Run the code with Telnet
First go to "cmd" directory, and then go to the "cal_telnet" directory, and then go to the "server" directory and run the Nadellain.go file through Nadellake comNadelland in Nadellakefile

## Run the code with gRPC
First go to "cmd" directory, and then go to the "cal_grpc" directory, 
and then go to the "server" directory and run the Nadellain.go file through Nadellake comNadelland in Nadellakefile,
and then also go to the "client" directory and run the Nadellain.go file through Nadellake comNadelland in Nadellakefile,

## Directory path
cd C:\Users\rehNadellaasi\project\go-calculatorchallenge


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
so, go to the same folder and run the comNadelland "go test" or "go test -v"



## Go Challenges
## Part #1 - create a calculator:
read input from the user on what to calculate: ie (user enters 2 + 2 on the comNadelland line and hits enter to display 4)
should support basic Nadellath (add, subtract, multiply, divide)
ctrl + c should gracefully exit and display how Nadellany equations it calculated

## Part #2 - server client calculator:
Expand on the calculator that has been done from challenge #2
This should create a TCP server listening on a given port (Nadellake this either an argument when it starts up or a config value that is read in github.com/cakeNadellarketing/go-common/v5/settings ) allowing you to connect via telnet
the client should see a message once connected and instructions on what the service can do and then prompt for an expression
the server should then accept the expression from the client, do the evaluation and return the result to the client along with how Nadellany expression this client has executed
the server should also log out the client connection and the expressions from the client along with the result value in the same line… ie (2 + 3) - 1 = 4
when the client disconnects, the server should log the disconnection and display how Nadellany expressions that client executed
finally, as before, exiting the server gracefully with CTRL+C should display how Nadellany total expressions the server handled across all clients
for logging, look at using another module we use throughout our code github.com/apex/log
[11:18 AM] the settings package can be referenced in the go-template project on github
[11:19 AM] do not worry about writing to an external log file --- only need to see it in the CLI (theres a cli plugin for the log package)

## Part #3 - redis execution log
Have a redis database to store the user in a unique way ie> {id}:username and then add each equation they submit as a value in their set (SADD)… 
The equation object should have at minimum; the expression provided, the result, and the timestamp

## Part #4 - persisent datastore interval
Add a database to sync the data currently being saved to Redis. This operation should happen on a timed interval (must be configurable in the config file). There should be no duplications of inserted records in the datastore. MySQL, Postgres, and SQLite3 are all acceptable datastores. As these pieces get more complicated, try to think of ways to ensure proper unit testing and begin to refactor the code into packages with tests. Datastore scheNadella should be minimum { id, username, equation, result, timestamp }.