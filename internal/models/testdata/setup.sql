CREATE TABLE cards (
                       id SERIAL PRIMARY KEY NOT NULL,
                       title VARCHAR(100) NOT NULL,
                       content TEXT NOT NULL,
                       created_at     timestamp DEFAULT NOW() NOT NULL,
                       updated_at     timestamp DEFAULT NOW() NOT NULL
);

CREATE INDEX idx_cards_created_at ON cards(created_at);

CREATE TABLE users (
                       id SERIAL PRIMARY KEY NOT NULL,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL,
                       hashed_password CHAR(60) NOT NULL,
                       created_at     timestamp DEFAULT NOW() NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

INSERT INTO users (name, email, hashed_password) VALUES (
                                                            'Alice Jones',
                                                            'alice@example.com',
                                                            '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG'
                                                        );