CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE songs (
                      id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                      group_name VARCHAR(100) NOT NULL,
                      song VARCHAR(100) NOT NULL,
                      release_date DATE NOT NULL,
                      text TEXT NOT NULL,
                      link VARCHAR(100) NOT NULL,
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_songs_group_and_name ON songs (group_name, song);