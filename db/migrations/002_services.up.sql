create table if not exists services(
  id INTEGER NOT NULL PRIMARY KEY,
  name VARCHAR,
  hash VARCHAR,
  version VARCHAR,
  created DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated DATETIME DEFAULT CURRENT_TIMESTAMP 
);
