CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  name varchar(50) not null,
  nick varchar(50) not null unique,
  email varchar(50) not null unique,
  password varchar(255) not null,
  created_at timestamp default current_timestamp
);

CREATE TABLE IF NOT EXISTS followers (
  user_id int not null,
  follower_id int not null,

  foreign key (user_id) 
  references users (id)
  on delete cascade,

  foreign key (follower_id) 
  references users (id)
  on delete cascade,

  PRIMARY key (user_id, follower_id)
);