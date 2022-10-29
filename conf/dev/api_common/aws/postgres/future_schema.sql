CREATE SCHEMA future
    CREATE TABLE
        appointments
    (
        appointment_id integer unique PRIMARY KEY,
        trainer_id     integer,
        user_id        integer,
        started_at     integer,
        ended_at       integer
    )