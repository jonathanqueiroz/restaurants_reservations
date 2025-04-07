CREATE TABLE reservations (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    restaurant_id INT NOT NULL,
    date DATE NOT NULL,
    time TIME NOT NULL
);
