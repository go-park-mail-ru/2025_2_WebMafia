INSERT INTO "user" (user_id, login, password_hash, email, avatar_url) VALUES
('33333333-3333-3333-3333-333333333333', 'hiphop_fan', '$2a$10$N9qo8uLOickgx2ZMRZoMye.MockHash001', 'hiphop@example.com', 'avatar_hiphop.jpg'),
('44444444-4444-4444-4444-444444444444', 'music_collector', '$2a$10$N9qo8uLOickgx2ZMRZoMye.MockHash002', 'collector@example.com', 'avatar_collector.jpg');

INSERT INTO genre (genre_id, genre_name) VALUES
('77777777-7777-7777-7777-777777777777', 'Hip-Hop'),
('88888888-8888-8888-8888-888888888888', 'Rock'),
('99999999-9999-9999-9999-999999999999', 'R&B'),
('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Alternative');

INSERT INTO artist (artist_id, artist_name, avatar_url) VALUES
('11111111-1111-1111-1111-111111111111', 'Tyler, the Creator', 'image1.jpg'),
('22222222-2222-2222-2222-222222222222', 'Mac Miller', 'image2.jpg'),
('33333333-3333-3333-3333-333333333333', 'Jpegmafia', 'image3.jpg'),
('44444444-4444-4444-4444-444444444444', 'The Roots', 'image4.jpg'),
('55555555-5555-5555-5555-555555555555', 'Iggy Pop', 'image5.jpg'),
('66666666-6666-6666-6666-666666666666', 'Молчат Дома', 'image6.jpg'),
('77777777-7777-7777-7777-777777777777', 'Arctic Monkeys', 'image7.jpg'),
('88888888-8888-8888-8888-888888888888', 'Kali Uchis', 'image8.jpg'),
('99999999-9999-9999-9999-999999999999', 'Playboi Carti', 'image9.jpg'),
('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'The Weeknd', 'image10.jpg');

INSERT INTO album (album_id, title, avatar_url, artist_id, release_date) VALUES
('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Flower Boy', 'image11.jpg', '11111111-1111-1111-1111-111111111111', '2017-07-21'),
('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Lust for Life', 'image12.jpg', '55555555-5555-5555-5555-555555555555', '1977-03-29'),
('dddddddd-dddd-dddd-dddd-dddddddddddd', 'All My Heroes Are Cornballs', 'image13.jpg', '33333333-3333-3333-3333-333333333333', '2019-09-13'),
('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 'Things Fall Apart', 'image14.jpg', '44444444-4444-4444-4444-444444444444', '1999-02-23'),
('ffffffff-ffff-ffff-ffff-ffffffffffff', 'Flower Boy (Deluxe)', 'image15.jpg', '11111111-1111-1111-1111-111111111111', '2017-12-01'),
('11111111-1111-1111-1111-111111111112', 'Veteran', 'image16.jpg', '33333333-3333-3333-3333-333333333333', '2018-01-19'),
('22222222-2222-2222-2222-222222222223', 'AM', 'image17.jpg', '77777777-7777-7777-7777-777777777777', '2013-09-09'),
('33333333-3333-3333-3333-333333333334', 'Isolation', 'image18.jpg', '88888888-8888-8888-8888-888888888888', '2018-04-06'),
('44444444-4444-4444-4444-444444444445', 'Die Lit', 'image19.jpg', '99999999-9999-9999-9999-999999999999', '2018-05-11'),
('55555555-5555-5555-5555-555555555556', 'After Hours', 'image20.jpg', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '2020-03-20');

INSERT INTO track (track_id, title, duration_s, file_url) VALUES
('66666666-6666-6666-6666-666666666666', 'See You Again', 180000, 'tyler_see_you_again.mp3'),
('77777777-7777-7777-7777-777777777777', 'Kenan vs Kel', 165000, 'peggy_kenan_vs_kel.mp3'),
('88888888-8888-8888-8888-888888888888', 'Rather Lie', 195000, 'carti_rather_lie.mp3'),
('99999999-9999-9999-9999-999999999999', 'See You Again (Remix)', 210000, 'tyler_see_you_again_remix.mp3'),
('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaab', '1539 N. Calvert', 172000, 'peggy_1539_calvert.mp3'),
('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbc', 'Shoota', 188000, 'carti_shoota.mp3'),
('cccccccc-cccc-cccc-cccc-cccccccccccd', 'Baby I''m Bleeding', 159000, 'peggy_baby_bleeding.mp3'),
('dddddddd-dddd-dddd-dddd-ddddddddddde', 'Long Time', 203000, 'carti_long_time.mp3'),
('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeef', 'Thug Tears', 176000, 'peggy_thug_tears.mp3'),
('ffffffff-ffff-ffff-ffff-ffffffffffff', 'FlatBed Freestyle', 192000, 'carti_flatbed.mp3'),

('11111111-1111-1111-1111-111111111113', 'Self Care', 348000, 'mac_self_care.mp3'),
('22222222-2222-2222-2222-222222222224', 'Do I Wanna Know?', 272000, 'arctic_do_i_wanna_know.mp3'),
('33333333-3333-3333-3333-333333333335', 'After The Storm', 213000, 'kali_after_the_storm.mp3'),
('44444444-4444-4444-4444-444444444446', 'Blinding Lights', 200000, 'weeknd_blinding_lights.mp3'),
('55555555-5555-5555-5555-555555555557', 'Судно', 178000, 'molchat_sudno.mp3');

INSERT INTO track_artist (track_id, artist_id) VALUES

('66666666-6666-6666-6666-666666666666', '11111111-1111-1111-1111-111111111111'),
('66666666-6666-6666-6666-666666666666', '88888888-8888-8888-8888-888888888888'),

('77777777-7777-7777-7777-777777777777', '33333333-3333-3333-3333-333333333333'),

('88888888-8888-8888-8888-888888888888', '99999999-9999-9999-9999-999999999999'),

('99999999-9999-9999-9999-999999999999', '11111111-1111-1111-1111-111111111111'),
('99999999-9999-9999-9999-999999999999', '88888888-8888-8888-8888-888888888888'),

('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaab', '33333333-3333-3333-3333-333333333333'),

('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbc', '99999999-9999-9999-9999-999999999999'),
('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbc', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'),

('cccccccc-cccc-cccc-cccc-cccccccccccd', '33333333-3333-3333-3333-333333333333'),

('dddddddd-dddd-dddd-dddd-ddddddddddde', '99999999-9999-9999-9999-999999999999'),

('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeef', '33333333-3333-3333-3333-333333333333'),

('ffffffff-ffff-ffff-ffff-ffffffffffff', '99999999-9999-9999-9999-999999999999'),

('11111111-1111-1111-1111-111111111113', '22222222-2222-2222-2222-222222222222'),
('22222222-2222-2222-2222-222222222224', '77777777-7777-7777-7777-777777777777'),
('33333333-3333-3333-3333-333333333335', '88888888-8888-8888-8888-888888888888'),
('44444444-4444-4444-4444-444444444446', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'),
('55555555-5555-5555-5555-555555555557', '66666666-6666-6666-6666-666666666666');

INSERT INTO track_genre (track_id, genre_id) VALUES

('66666666-6666-6666-6666-666666666666', '77777777-7777-7777-7777-777777777777'),
('77777777-7777-7777-7777-777777777777', '77777777-7777-7777-7777-777777777777'),
('88888888-8888-8888-8888-888888888888', '77777777-7777-7777-7777-777777777777'),
('99999999-9999-9999-9999-999999999999', '77777777-7777-7777-7777-777777777777'),
('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaab', '77777777-7777-7777-7777-777777777777'),
('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbc', '77777777-7777-7777-7777-777777777777'),
('cccccccc-cccc-cccc-cccc-cccccccccccd', '77777777-7777-7777-7777-777777777777'),
('dddddddd-dddd-dddd-dddd-ddddddddddde', '77777777-7777-7777-7777-777777777777'),
('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeef', '77777777-7777-7777-7777-777777777777'),
('ffffffff-ffff-ffff-ffff-ffffffffffff', '77777777-7777-7777-7777-777777777777'),
('11111111-1111-1111-1111-111111111113', '77777777-7777-7777-7777-777777777777'),

('22222222-2222-2222-2222-222222222224', '88888888-8888-8888-8888-888888888888'),

('33333333-3333-3333-3333-333333333335', '99999999-9999-9999-9999-999999999999'),
('44444444-4444-4444-4444-444444444446', '99999999-9999-9999-9999-999999999999'),

('55555555-5555-5555-5555-555555555557', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa');

INSERT INTO track_album (track_id, album_id) VALUES
('66666666-6666-6666-6666-666666666666', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb'),

('77777777-7777-7777-7777-777777777777', '11111111-1111-1111-1111-111111111112'),
('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeef', '11111111-1111-1111-1111-111111111112'),

('88888888-8888-8888-8888-888888888888', '44444444-4444-4444-4444-444444444445'),
('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbc', '44444444-4444-4444-4444-444444444445'),
('dddddddd-dddd-dddd-dddd-ddddddddddde', '44444444-4444-4444-4444-444444444445');

