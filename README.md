# transaction-routine
This project is a challenge to analyze my code skills at the moment.

## About
This project was developed to solve a job interview problem. In summary, it is a ledger where we have records of transactions made, just like a bank statement. Some tools were used to assist in the development, thus allowing us to focus on the business rule itself. Further down, we'll be able to see a sequence diagram to better understand these rules and also how to execute and run this application.

## How to run it
```bash
# Install dependecies
$ go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
$ go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
$ go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
$ go install github.com/renanbastos93/boneless/cmd/boneless@latest

# Up database
$ docker-compose up

# Run migrations
$ boneless migrate app up

# Run server
$ boneless run

# For execute unit tests
$ make test
# output in last line: 
# total:                                                                          (statements)            64.7%
```

## Driagram
```mermaid
sequenceDiagram 
  title Transaction Routine
  
  actor User
  
  alt register user
  User-->>Service:createAccount(documentNumber);
  Service-->>Database: saveUser(data);
  Database-->>User: created;
  end
  
  alt purchase
  User-->>Service:buy(accountID, operationTypeID, amout);
  Service-->>Service: validateOperationsTypes();
  Service-->>Service: normalizeValues();
  Service-->>Database: saveTransaction();
  Database-->>User: saved your transaction;
  end
  
  alt get user
  User-->>Service:getUserByAccountID(id);
  Service-->>Database: getUser(accountID);
  Database-->>User: user data;
  end
```
<!-- 
![image](https://github.com/renanbastos93/transaction-routine/assets/8202898/26cd5c8c-927b-46b7-a987-228d9220ce7a)


![image](https://github.com/renanbastos93/transaction-routine/assets/8202898/5a73fd16-6282-43c1-a38f-0716db205202)


![image](https://github.com/renanbastos93/transaction-routine/assets/8202898/49bcc520-d941-4f24-b6b6-4899110217e2) -->
