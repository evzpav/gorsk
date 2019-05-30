CREATE TABLE trades (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    id_parent_order INTEGER DEFAULT NULL,
    target_risk DECIMAL DEFAULT NULL,
    atr JSON NULL DEFAULT '{}',
    balance DECIMAL NULL DEFAULT NULL,
    pair VARCHAR(20) NULL DEFAULT NULL,
    exchange VARCHAR(100) NULL DEFAULT NULL,
    side VARCHAR(20) NULL DEFAULT NULL,
    initial_position DECIMAL NULL DEFAULT NULL,
    actual_position DECIMAL NULL DEFAULT NULL,
    total_position DECIMAL NULL DEFAULT NULL,
    entry_price DECIMAL NULL DEFAULT NULL,
    entry_timestamp TIMESTAMPTZ NULL,
    initial_target DECIMAL NULL DEFAULT NULL,
    actual_target DECIMAL NULL DEFAULT NULL,
    is_open BOOLEAN DEFAULT FALSE,
    is_closed BOOLEAN DEFAULT FALSE,
    initial_stop_loss DECIMAL NULL DEFAULT NULL,
    actual_stop_loss DECIMAL NULL DEFAULT NULL,
    exit_price DECIMAL NULL DEFAULT NULL,
    exit_timestamp TIMESTAMPTZ NULL,
	initial_risk DECIMAL DEFAULT NULL,
    actual_risk DECIMAL DEFAULT NULL,
    change_history JSON NULL DEFAULT '{}',
    is_child BOOLEAN DEFAULT FALSE,
    is_parent BOOLEAN DEFAULT TRUE
);