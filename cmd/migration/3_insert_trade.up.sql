INSERT INTO trades (
    pair, 
    exchange, 
    balance,
    position_initial,
    position_actual,
    entry_price,
    stop_loss_initial,
    stop_loss_actual,
    exit_price,
    change_history,
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
    '[
        { "datetime": "23/05/19 12:00", "comment":"This is my comment 1", "stop_loss": "10", "split_siblings": []},
        { "datetime": "24/05/19 13:00", "comment":"This is my comment 2", "stop_loss": "12", "split_siblings": []}
    ]',
    FALSE);