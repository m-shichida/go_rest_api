USE go_rest_api;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS fishes;

CREATE TABLE fishes (
  id int AUTO_INCREMENT NOT NULL,
  name varchar(50) NOT NULL COMMENT "名前",
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(id)
) COMMENT = "魚";

CREATE TABLE posts (
  id int AUTO_INCREMENT NOT NULL,
  image_url varchar(255) NOT NULL COMMENT "画像URL",
  address varchar(255) NOT NULL COMMENT "住所",
  description TEXT NOT NULL COMMENT "釣果の説明",
  fish_id int NOT NULL COMMENT "魚のID",
  fising_at datetime NOT NULL COMMENT "釣行日",
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  FOREIGN KEY (fish_id)
    REFERENCES fishes(id)
) COMMENT = "投稿";
