CREATE TABLE IF NOT EXISTS Accounts (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    document_number VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL,
    deleted_at BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS OperationTypes (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    description TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS Transactions (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    account_id VARCHAR(36) NOT NULL,
    operation_type_id VARCHAR(36) NOT NULL,
    amout FLOAT,
    event_date BIGINT NOT NULL,

    FOREIGN KEY(account_id) REFERENCES Accounts(id) ON UPDATE CASCADE,
    FOREIGN KEY(operation_type_id) REFERENCES OperationTypes(id) ON UPDATE CASCADE
);

