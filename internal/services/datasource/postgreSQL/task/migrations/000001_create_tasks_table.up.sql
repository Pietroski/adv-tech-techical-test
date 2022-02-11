CREATE EXTENSION btree_gist;

CREATE TABLE IF NOT EXISTS tasks
(
    table_id        BIGSERIAL,
    task_id         uuid UNIQUE NOT NULL,
    user_id         uuid        NOT NULL,
    data_range      tsrange,
    reminder_period INTERVAL,
    created_at      timestamptz DEFAULT now(),

    CONSTRAINT pk_tasks PRIMARY KEY (task_id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (user_id),

    EXCLUDE USING gist (user_id WITH =, data_range WITH &&)
);

-- INSERT INTO tasks (task_id, user_id, data_range, reminder_period)
-- VALUES (
--            '72386c28-c70d-416f-b39f-e65f64b6e34e',
--            '72386c28-c70d-416f-b39f-e65f64b6e20e',
--            '[2022-02-10 14:30, 2022-02-10 15:30)', '00:30')
-- RETURNING *;
-- INSERT INTO tasks (task_id, user_id, data_range, reminder_period)
-- VALUES (
--            '72386c28-c70d-416f-b39f-e65f64b6e38f',
--            '72386c28-c70d-416f-b39f-e65f64b6e205',
--            '[2022-02-10 14:30, 2022-02-10 15:30)', '00:30')
-- RETURNING *;
-- INSERT INTO tasks (task_id, user_id, data_range, reminder_period)
-- VALUES (
--            'a2143e10-f90e-4a48-9767-611f313e139a',
--            '72386c28-c70d-416f-b39f-e65f64b6e20e',
--            '[2022-02-10 14:00, 2022-02-10 14:30)', '00:30')
-- RETURNING *;
-- INSERT INTO tasks (task_id, user_id, data_range, reminder_period)
-- VALUES (
--            '68a2bd21-e542-4510-ad30-bb484a3d7e6a',
--            '72386c28-c70d-416f-b39f-e65f64b6e20e',
--            '[2022-02-10 13:00, 2022-02-10 13:30)', '00:30')
-- RETURNING *;
