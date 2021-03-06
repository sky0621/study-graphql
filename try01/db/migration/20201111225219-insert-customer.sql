
-- +migrate Up
INSERT INTO customer(name, age) VALUES
  ('Sato', 30),
  ('Suzuki', 22),
  ('Takahashi', 27),
  ('Tanaka', 19),
  ('Ito', 14),
  ('Watanabe', 33),
  ('Yamamoto', 25),
  ('Nakamura', 29),
  ('Kobayashi', 17),
  ('Kato', 20),
  ('Yoshida', 35),
  ('Yamada', 20),
  ('Sasaki', 24),
  ('Yamaguchi', 22),
  ('Matsumoto', 38);

INSERT INTO todo(task, user_id) VALUES
  ('さとうのタスク01', 1),
  ('さとうのタスク02', 1),
  ('さとうのタスク03', 1),
  ('すずきのタスク01', 2),
  ('すずきのタスク02', 2),
  ('すずきのタスク03', 2),
  ('たかはしのタスク01', 3),
  ('たかはしのタスク02', 3),
  ('たかはしのタスク03', 3),
  ('たなかのタスク01', 4),
  ('たなかのタスク02', 4),
  ('たなかのタスク03', 4),
  ('いとうのタスク01', 5),
  ('いとうのタスク02', 5),
  ('いとうのタスク03', 5),
  ('わたなべのタスク01', 6),
  ('わたなべのタスク02', 6),
  ('わたなべのタスク03', 6),
  ('やまもとのタスク01', 7),
  ('やまもとのタスク02', 7),
  ('やまもとのタスク03', 7),
  ('なかむらのタスク01', 8),
  ('なかむらのタスク02', 8),
  ('なかむらのタスク03', 8),
  ('こばやしのタスク01', 9),
  ('こばやしのタスク02', 9),
  ('こばやしのタスク03', 9),
  ('かとうのタスク01', 10),
  ('かとうのタスク02', 10),
  ('かとうのタスク03', 10),
  ('よしだのタスク01', 11),
  ('よしだのタスク02', 11),
  ('よしだのタスク03', 11),
  ('やまだのタスク01', 12),
  ('やまだのタスク02', 12),
  ('やまだのタスク03', 12),
  ('ささきのタスク01', 13),
  ('ささきのタスク02', 13),
  ('ささきのタスク03', 13),
  ('やまぐちのタスク01', 14),
  ('やまぐちのタスク02', 14),
  ('やまぐちのタスク03', 14),
  ('まつもとのタスク01', 15),
  ('まつもとのタスク02', 15),
  ('まつもとのタスク03', 15);

-- +migrate Down
TRUNCATE table todo;
TRUNCATE table customer;
