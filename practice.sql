SELECT
    categories.id,
    categories.name,
    categories.description,
    categories.created_at,
    categories.updated_at,
    COUNT(items.id) AS item_count
FROM categories
         LEFT JOIN
     items ON categories.id = items.category_id
GROUP BY categories.id, categories.name, categories.description, categories.created_at, categories.updated_at;


SELECT
    categories.id,
    categories.name,
    categories.description,
    categories.created_at,
    categories.updated_at,
    SUM(items.amount) AS sum_amount
FROM categories
         LEFT JOIN
     items ON categories.id = items.category_id
GROUP BY categories.id, categories.name, categories.description, categories.created_at, categories.updated_at;


SELECT
    categories.id,
    categories.name,
    categories.description,
    categories.created_at,
    categories.updated_at
FROM categories
         LEFT JOIN
     items ON categories.id = items.category_id
WHERE items.amount > 40
GROUP BY categories.id, categories.name, categories.description, categories.created_at, categories.updated_at;


DELETE FROM categories
WHERE id NOT IN (
    SELECT DISTINCT category_id
    FROM items
    WHERE category_id IS NOT NULL
);