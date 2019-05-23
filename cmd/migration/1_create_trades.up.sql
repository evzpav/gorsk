CREATE TABLE trades (
 id serial PRIMARY KEY,
 pair VARCHAR(20) NULL DEFAULT NULL,
 exchange VARCHAR(200) NULL,
 balance DECIMAL NULL,
 entry_price DECIMAL NULL,
 position_initial DECIMAL NULL,
 position_actual DECIMAL NULL,
 stop_loss_initial DECIMAL NULL,
 stop_loss_actual DECIMAL NULL,
 exit_price DECIMAL NULL,
 change_history JSON NULL,
 is_split BOOLEAN NOT NULL DEFAULT FALSE
);