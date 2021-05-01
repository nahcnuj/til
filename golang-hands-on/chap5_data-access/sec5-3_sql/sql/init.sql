CREATE TABLE mydata (
    id      INTEGER PRIMARY KEY AUTOINCREMENT,
    name    TEXT    NOT NULL,
    mail    TEXT,
    age     INTEGER
);

INSERT INTO mydata
    (name, mail, age)
VALUES
    ("jun",     "nahcnuj.work@gmail.com", 25),
    ("taro",    "taro@example.com",       39),
    ("hanako",  "hanako@example.com",     28),
    ("sachiko", "sachiko@example.com",    17),
    ("jiro",    "jiro@example.com",       6);