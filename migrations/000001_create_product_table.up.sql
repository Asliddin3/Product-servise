CREATE TABLE IF NOT EXISTS categories(
  id SERIAL PRIMARY KEY,
  name TEXT
);
CREATE TABLE IF NOT EXISTS types(
  id SERIAL PRIMARY KEY,
  name TEXT
);
CREATE Table IF NOT EXISTS products(
  id SERIAL PRIMARY KEY,
  name TEXT,
  categoryid INT REFERENCES categories(id),
  typeid INT REFERENCES types(id)
);