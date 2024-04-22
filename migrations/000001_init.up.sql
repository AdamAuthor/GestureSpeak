CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE video_files (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    s3_key VARCHAR(255) NOT NULL,
    upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    processing_status VARCHAR(50)
);

CREATE TABLE processing_results (
    id SERIAL PRIMARY KEY,
    video_file_id INT NOT NULL REFERENCES video_files(id),
    gestures_description TEXT,
    processing_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);