CREATE TABLE IF NOT EXISTS users
(
    table_id   BIGSERIAL,
    user_id    uuid UNIQUE NOT NULL,
    email      text UNIQUE NOT NULL,
    first_name TEXT        NOT NULL,
    last_name  TEXT        NOT NULL,
    created_at timestamptz DEFAULT now(),

    CONSTRAINT pk_users PRIMARY KEY (user_id)
);

-- INSERT INTO users (user_id, first_name, last_name)
-- VALUES ('72386c28-c70d-416f-b39f-e65f64b6e20e', 'Augusto', 'Pietroski')
-- RETURNING *;
-- INSERT INTO users (user_id, first_name, last_name)
-- VALUES
--     ('72386c28-c70d-416f-b39f-e65f64b6e20e', 'Augusto', 'Pintroski'),
--     ('72386c28-c70d-416f-b39f-e65f64b6e205', 'Lucas', 'Bichinski'),
--     ('72386c28-c70d-416f-b39f-e65f64b6e207', 'Gabriel', 'Chupinski')
-- RETURNING *;
