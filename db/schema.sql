CREATE TABLE IF NOT EXISTS todos (
     id INTEGER PRIMARY KEY AUTOINCREMENT,
     message TEXT NOT NULL,
     priority INTEGER NOT NULL,
     status TEXT CHECK(status IN ('todo', 'doing', 'done')) DEFAULT 'todo',
     created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- TODO: add a intra-list ordering
