CREATE DATABASE snippetbox;

\c snippetbox

CREATE TABLE snippets (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX idx_snippets_created_at ON snippets (created_at);

INSERT INTO snippets (title, content, expires_at) VALUES (
  'An old silent pond',
  E'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
  CURRENT_TIMESTAMP + INTERVAL '365 day'
);

INSERT INTO snippets (title, content, expires_at) VALUES (
  'Over the wintry forest',
  E'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
  CURRENT_TIMESTAMP + INTERVAL '365 day'
);

INSERT INTO snippets (title, content, expires_at) VALUES (
  'First autumn morning',
  E'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
  CURRENT_TIMESTAMP + INTERVAL '7 day'
);

\set app_user `cat $APP_USER`
\set app_password `cat $APP_PASSWORD`

CREATE USER :app_user WITH PASSWORD :'app_password';
GRANT pg_read_all_data TO :app_user;
GRANT pg_write_all_data TO :app_user;
