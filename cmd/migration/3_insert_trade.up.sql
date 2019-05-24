INSERT INTO trades (
    pair, 
    exchange, 
    balance,
    initial_position,
    actual_position,
    entry_price,
    initial_stop_loss,
    actual_stop_loss,
    exit_price,
    is_split) 
VALUES (
    'BTC/USD',
    'Bitfinex',
    2000,
    1.00,
    1.00,
    8000,
    7900,
    7900,
    null,
    FALSE
);