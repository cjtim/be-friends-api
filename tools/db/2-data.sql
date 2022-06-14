-- Test credential
-- test-preview@gmail.com
-- test
INSERT INTO "user" (id, name, email, password, line_uid)
VALUES 
('097b2b93-0a66-4de2-a5a6-4f0b15acc54c', 'Test', 'test-preview@gmail.com', '$2a$04$x1jU9wX5Ab7fzyL.qG5CO.4CHB/t3lq0obSjdXkJ5.tmlwjVJZyRO', NULL),
('e15f64ec-d87e-4bc3-8680-3fb7b47d438d', 'Line User Test', NULL, NULL, 'aaaaaaaaaaaaaaaaaaaa');

INSERT INTO "pet" (id, name, description)
VALUES
(1, 'Sam', 'Sam is friendly dog'),
(2, 'John', 'Johnny is friendly cat');

INSERT INTO "tag" (id, name, is_internal)
VALUES
(1, 'Friendly', FALSE), (2, 'Girl', FALSE), (3, 'Boy', FALSE), (4, 'Need attention', TRUE), (5, 'Line User', TRUE), (6, 'From campaign', TRUE),
(7, 'Admin', TRUE);

INSERT INTO "tag_pet" (pet_id, tag_id)
VALUES
(1, 1), (1, 2), (1, 3), (2, 1), (2, 3), (2, 4);

INSERT INTO "tag_user" (user_id, tag_id)
VALUES
('097b2b93-0a66-4de2-a5a6-4f0b15acc54c', 1),
('097b2b93-0a66-4de2-a5a6-4f0b15acc54c', 2),
('097b2b93-0a66-4de2-a5a6-4f0b15acc54c', 4),
('097b2b93-0a66-4de2-a5a6-4f0b15acc54c', 6),
('e15f64ec-d87e-4bc3-8680-3fb7b47d438d', 1),
('e15f64ec-d87e-4bc3-8680-3fb7b47d438d', 3),
('e15f64ec-d87e-4bc3-8680-3fb7b47d438d', 4),
('e15f64ec-d87e-4bc3-8680-3fb7b47d438d', 7);
