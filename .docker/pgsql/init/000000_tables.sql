CREATE OR REPLACE FUNCTION set_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

--

CREATE TABLE shops
(
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

--

CREATE TABLE orders
(
    id BIGSERIAL PRIMARY KEY,
    shop_id BIGINT REFERENCES shops(id) ON DELETE CASCADE,
    number VARCHAR(255) NOT NULL,
    total INT,
    customer_name VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_orders_shop_number ON orders (shop_id, number);

--

CREATE TABLE telegram_integrations
(
    id         BIGSERIAL PRIMARY KEY,
    shop_id    BIGINT UNIQUE REFERENCES shops(id) ON DELETE CASCADE,
    bot_token  VARCHAR(255) NOT NULL DEFAULT '',
    chat_id    VARCHAR(255) NOT NULL DEFAULT '',
    is_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TRIGGER trigger_update_timestamp
    BEFORE UPDATE
    ON telegram_integrations
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

--

CREATE TYPE TELEGRAM_SEND_STATUS AS ENUM ('SENT', 'FAILED');

CREATE TABLE telegram_send_log
(
    id BIGSERIAL PRIMARY KEY,
    shop_id BIGINT REFERENCES shops(id) ON DELETE CASCADE,
    order_id BIGINT REFERENCES orders(id) ON DELETE CASCADE,
    message VARCHAR(255) NOT NULL,
    status TELEGRAM_SEND_STATUS NOT NULL,
    error VARCHAR(255),
    sent_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_telegram_send_log_shop_order ON telegram_send_log(shop_id, order_id);

--