CREATE USER gorsk;
ALTER USER gorsk WITH ENCRYPTED PASSWORD 'gorskpass';
CREATE DATABASE gorskdb;
GRANT ALL PRIVILEGES ON DATABASE gorskdb TO gorsk;
\c gorskdb
CREATE EXTENSION pgcrypto;