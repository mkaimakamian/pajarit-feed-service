CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id TEXT PRIMARY KEY,
    author_id TEXT NOT NULL,
    content TEXT NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    --, FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS followup (
		follower_id TEXT NOT NULL,
		followed_id TEXT NOT NULL,
		PRIMARY KEY(follower_id, followed_id)
		--, FOREIGN KEY(follower_id) REFERENCES users(id),
		-- FOREIGN KEY(followed_id) REFERENCES users(id)
	);


CREATE TABLE IF NOT EXISTS timelines (
    user_id TEXT PRIMARY KEY,
    posts TEXT NOT NULL  -- JSON con los posts de los followed
    --, FOREIGN KEY (user_id) REFERENCES users(id)
);