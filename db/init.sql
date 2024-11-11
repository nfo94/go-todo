CREATE TABLE todos (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  completed BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP WITH TIME ZONE
);

INSERT INTO todos (title, description, completed) VALUES
  ('Learn Go', 'Awesome language, but some terrible docs. Maybe is the Python/JavaScript community in me.', FALSE),
  ('Try Demon Souls', 'Every From Software fan needs to play this game, trully remarkable.', TRUE),
  ('Learn GraphQL', 'Because I tried Demon Souls and loved it, I couldn''t study GraphQL. A shame.', FALSE),
  ('Watch Matrix again', '... For the 50th time, on a sunday night. Yes.', TRUE);
