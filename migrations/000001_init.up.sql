CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE video_files (
    ID SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(ID),
    filename VARCHAR(255) NOT NULL,
    upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    processing_status VARCHAR(50)
);

CREATE TABLE processing_results (
    ID SERIAL PRIMARY KEY,
    video_file_id INT NOT NULL REFERENCES video_files(ID),
    gestures_description TEXT,
    processing_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
