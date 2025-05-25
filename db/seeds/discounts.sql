TRUNCATE discounts CASCADE;

INSERT INTO discounts (discount_code, discount_value, valid_from, valid_to, remaining_count) VALUES
('WELCOME10', 10.00, '2023-01-01', '2025-12-31', 100),
('SUMMER20', 20.00, '2023-06-01', '2023-08-31', 50),
('SPECIAL25', 25.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '30 days', 25);