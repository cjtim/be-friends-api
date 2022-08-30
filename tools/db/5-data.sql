INSERT INTO "status" (name)
VALUES ('new'), ('pending'), ('adoped');


-- -- Test credential
-- -- test-preview@gmail.com
-- -- test
-- INSERT INTO "user" (id, name, email, password, is_org, line_uid)
-- VALUES 
-- ('097b2b93-0a66-4de2-a5a6-4f0b15acc54c', 'Test', 'test-preview@gmail.com', '$2a$04$x1jU9wX5Ab7fzyL.qG5CO.4CHB/t3lq0obSjdXkJ5.tmlwjVJZyRO', TRUE, NULL),
-- ('e15f64ec-d87e-4bc3-8680-3fb7b47d438d', 'Line User Test', NULL, NULL, FALSE, 'aaaaaaaaaaaaaaaaaaaa');

-- INSERT INTO "pet" (name, description, lat, lng)
-- VALUES
-- ('Sam', 'Sam is friendly dog', 53.1, 64),
-- ('John', 'Johnny is friendly cat', 55, 60),
-- ('James', NULL, 56, 57);

-- INSERT INTO "pic_pet" (pet_id, picture_url)
-- VALUES
-- (1, 'https://www.collinsdictionary.com/images/full/dog_230497594.jpg'),
-- (1, 'https://mpng.subpng.com/20180505/tse/kisspng-havanese-dog-pet-sitting-puppy-cat-dog-daycare-dog-claw-free-buckle-chart-5aed4809b1f553.3338712315254999137289.jpg'),
-- (2, 'https://i.guim.co.uk/img/media/fe1e34da640c5c56ed16f76ce6f994fa9343d09d/0_174_3408_2046/master/3408.jpg?width=1200&height=900&quality=85&auto=format&fit=crop&s=0d3f33fb6aa6e0154b7713a00454c83d'),
-- (2, 'https://media.istockphoto.com/photos/pug-sitting-and-panting-1-year-old-isolated-on-white-picture-id450709593?k=20&m=450709593&s=612x612&w=0&h=82zzJc3Cz39B6LyrQ_N2b4zXxYzZIEH9aNDZWzrZspg='),
-- (3, 'https://www.akc.org/wp-content/uploads/2017/11/Golden-Retriever-Puppy.jpg');

-- INSERT INTO "tag" (name, is_internal)
-- VALUES
-- ('Friendly', FALSE),
-- ('Girl', FALSE),
-- ('Boy', FALSE),
-- ('Need attention', TRUE),
-- ('Line User', TRUE),
-- ('From campaign', TRUE),
-- ('Admin', TRUE);

-- INSERT INTO "tag_pet" (pet_id, tag_id)
-- VALUES
-- (1, 1), (1, 2), (1, 3), (2, 1), (2, 3), (2, 4);
