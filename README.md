# transaction-routine
This project is a challenge to analyze my code skills at the moment.

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

![image](https://github.com/renanbastos93/transaction-routine/assets/8202898/26cd5c8c-927b-46b7-a987-228d9220ce7a)


![image](https://github.com/renanbastos93/transaction-routine/assets/8202898/5a73fd16-6282-43c1-a38f-0716db205202)


![image](https://github.com/renanbastos93/transaction-routine/assets/8202898/49bcc520-d941-4f24-b6b6-4899110217e2)
