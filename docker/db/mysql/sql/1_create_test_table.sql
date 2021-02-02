CREATE TABLE test_table(
  id INT(10) AUTO_INCREMENT NOT NULL,
  name VARCHAR(30) NOT NULL,
  amount INT(10) NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);

INSERT INTO test_table (name, amount) VALUES ("apple", 50), ("orange", 100), ("lemon", 120);
