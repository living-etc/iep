CREATE TABLE IF NOT EXISTS exercises(
  id TEXT NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS tests(
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  exercise_id TEXT NOT NULL,
  resource_type TEXT NOT NULL,
  resource_name TEXT NOT NULL,
  resource_attribute TEXT NOT NULL,
  resource_attribute_value TEXT NOT NULL,
  negation INTEGER NOT NULL
);
