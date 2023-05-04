-- 1-txxt
INSERT INTO another_store_products (store_id, product_id, quantity)
SELECT 2 AS store_id, product_id, quantity
FROM stocks
WHERE store_id = 1;

-- 2-si

SELECT staffs.first_name || ' ' || staffs.last_name AS "Сотурник",
       categories.category_name AS "Категория",
       products.product_name AS "Продукт",
       order_items.quantity AS "Количество",
       order_items.list_price * order_items.quantity AS "Общая Цена"
FROM orders
         JOIN order_items ON orders.order_id = order_items.order_id
         JOIN products ON order_items.product_id = products.product_id
         JOIN categories ON products.category_id = categories.category_id
         JOIN staffs ON orders.staff_id = staffs.staff_id
WHERE orders.order_date = '2016-01-01';

-- 3-si

INSERT INTO promo_codes (name, discount, discount_type, order_limit_price)
VALUES ('JUBAJUBA', 10000, 'Фикс', 50000);

-- 4-si
SELECT SUM(oi.list_price * oi.quantity * (CASE WHEN pc.discount_type = 'Фикс' THEN 1 ELSE (100 - pc.discount) / 100 END)) AS total_sum
FROM orders o
         JOIN order_items oi ON o.order_id = oi.order_id
         LEFT JOIN promo_codes pc ON o.promo_code = pc.name
WHERE o.order_id = 1;


SELECT
    SUM(oi.quantity * oi.list_price * (1 - oi.discount / 100))
        AS order_total
FROM
    order_items oi
        JOIN orders o ON oi.order_id = o.order_id
        LEFT JOIN promo_codes pc ON o.promo_code = pc.name
WHERE
        oi.order_id = 1;
-- 5-si




-- 1-si
UPDATE stocks
SET quantity = quantity - 10
WHERE store_id = 1 AND product_id = 1 AND quantity >= 10;

UPDATE stocks
SET quantity = quantity + 10
WHERE store_id = 2 AND product_id = 1;
