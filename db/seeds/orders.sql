TRUNCATE orders CASCADE;
ALTER SEQUENCE orders_id_seq RESTART WITH 1;

-- Order 1: 2 Waffle with Berries ($12.99 each) + 1 Vanilla Bean Creme Brulee ($9.99)
-- Original total: $35.97, with WELCOME10 discount (10%)
INSERT INTO orders (original_total, discount_code, final_total, created_at, updated_at) VALUES 
    (35.97, 'WELCOME10', 32.37, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (38.97, NULL, 38.97, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Order items
INSERT INTO order_items (order_id, product_id, quantity, created_at, updated_at) VALUES
    (1, 1, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (1, 2, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 1, 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);