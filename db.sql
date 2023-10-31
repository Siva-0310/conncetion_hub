CREATE TABLE users IF NOT EXISTS{
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    user_email VARCHAR(256) NOT NULL UNIQUE,
    user_name VARCHAR(256) NOT NULL,
    user_password VARCHAR(256) NOT NULL,
    CONSTRAINT check_email CHECK(user_email REGEXP "^[a-zA-Z0-9][a-zA-Z0-9.!#$%&'*+-/=?^_`{|}~]*?[a-zA-Z0-9._-]?@[a-zA-Z0-9][a-zA-Z0-9._-]*?[a-zA-Z0-9]?\\.[a-zA-Z]{2,63}$")
};