TRUNCATE products CASCADE;
ALTER SEQUENCE products_id_seq RESTART WITH 1;

INSERT INTO products (name, price, category, image_url) VALUES 
    ('Waffle with Berries', 12.99, 'Waffle', '/images/d85f60cd-dd67-43c7-b525-c378ccc9e57c.png'),
    ('Vanilla Bean Creme Brulee', 9.99, 'French', '/images/33de0504-4619-4e37-bc85-e080ddd4e354.png'),
    ('Macaron Mix of Five', 15.99, 'French', '/images/d003e638-30e8-4a1f-bf80-d3927abb38d0.png'),
    ('Classic Tiramisu', 11.99, 'Italian', '/images/78e956d6-6da7-4163-aceb-f225623b7ab8.png'),
    ('Pistachio Baklava', 8.99, 'Middle Eastern', '/images/6e635dd2-682f-45c0-bf4c-ad7537956629.png'),
    ('Lemon Meringue Pie', 13.99, 'Pie', '/images/e9563fab-3e54-4c00-ac52-50a30fce3afb.png'),
    ('Red Velvet Cake', 14.99, 'Cake', '/images/01f75f76-68c2-4e8f-875d-15623b9ace41.png'),
    ('Salted Caramel Brownie', 7.99, 'American', '/images/c32c5878-46ff-4460-ac51-5a86779006ae.png'),
    ('Vanilla Panna Cotta', 10.99, 'Italian', '/images/d236165a-df1b-4925-b6fa-213b3e1c8f0e.png');