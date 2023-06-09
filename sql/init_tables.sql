CREATE TABLE IF NOT EXISTS users (
     id SERIAL PRIMARY KEY,
     vk INTEGER NOT NULL,
     full_name VARCHAR(80) NOT NULL,
     user_group VARCHAR(20) NOT NULL,
     qr VARCHAR(80) NOT NULL DEFAULT 'nil',
     coins INTEGER NOT NULL DEFAULT 0,
     rating REAL NOT NULL DEFAULT 0,
     admin INTEGER NOT NULL DEFAULT 0,
     is_banned INTEGER NOT NULL DEFAULT 0,
     is_subscribed INTEGER NOT NULL DEFAULT 1,
     notifications INTEGER NOT NULL DEFAULT 1
);

CREATE INDEX IF NOT EXISTS coins_index ON users(id);
CREATE INDEX IF NOT EXISTS rating_index ON users(rating);
CREATE INDEX IF NOT EXISTS vk_index ON users(vk);
CREATE INDEX IF NOT EXISTS subscribe_index ON users(is_subscribed);
CREATE INDEX IF NOT EXISTS notifications_index ON users(notifications);

CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER REFERENCES users(id),
    token VARCHAR(128) NOT NULL,
    qr VARCHAR(80) NOT NULL DEFAULT 'nil',
    item_id INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS owner_items_index ON items(owner_id);
CREATE INDEX IF NOT EXISTS item_index ON items(item_id);

CREATE TABLE IF NOT EXISTS achievements (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER REFERENCES users(id),
    achievement_id INTEGER NOT NULL,
    progress INTEGER NOT NULL DEFAULT 0,
    completed INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS owner_achievement_index ON achievements(owner_id);
CREATE INDEX IF NOT EXISTS achievement_index ON achievements(achievement_id, progress);

CREATE TABLE IF NOT EXISTS market (
    id SERIAL PRIMARY KEY,
    hash VARCHAR(128) NOT NULL,
    item_id INTEGER REFERENCES users(id),
    price INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS hash_index ON market(hash);
CREATE INDEX IF NOT EXISTS item_market_index ON market(item_id);

CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY,
    weight REAL NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS events_votes (
    id SERIAL PRIMARY KEY,
    event_id INTEGER REFERENCES events(id),
    user_id INTEGER REFERENCES users(id),
    general INTEGER NOT NULL,
    organization INTEGER NOT NULL,
    conversion INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS votes_user_index ON events_votes(user_id);
CREATE INDEX IF NOT EXISTS votes_event_index ON events_votes(event_id);

CREATE TABLE IF NOT EXISTS events_ratings (
    id SERIAL PRIMARY KEY,
    event_id INTEGER REFERENCES events(id)
);

CREATE INDEX IF NOT EXISTS ratings_event_index ON events_ratings(event_id);

CREATE TABLE IF NOT EXISTS events_visits (
    id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL,
    user_id INTEGER REFERENCES users(id),
    position INTEGER NOT NULL DEFAULT 100
);

CREATE TABLE IF NOT EXISTS bonus (
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL,
    user_id INTEGER REFERENCES users(id),
    liked INTEGER NOT NULL,
    comment INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS like_bonus ON bonus(liked);
CREATE INDEX IF NOT EXISTS comment_bonus ON bonus(comment);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    admin_id INTEGER REFERENCES users(id),
    user_id INTEGER REFERENCES users(id),
    amount INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);