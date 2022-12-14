CREATE SCHEMA IF NOT EXISTS wallets;
CREATE TABLE IF NOT EXISTS wallets.users
(
    id             bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name           text NOT NULL,
    creation_date  timestamptz DEFAULT now() NOT NULL,
    update_date    timestamptz DEFAULT now() NOT NULL
);