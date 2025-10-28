-- SQL миграция для вставки данных из YouTube Music
-- Генерация: 2025-10-26T01:17:23.037426

DO $$
DECLARE
    v_artist_id UUID;
    v_album_id UUID;
    v_track_id UUID;
BEGIN
-- Артист: Amon Amarth
INSERT INTO artist (artist_name, description, avatar_url)
VALUES ('Amon Amarth', 'Amon Amarth — шведская метал-группа, играющая в жанре мелодичный дэт-метал. Название группы значит «Роковая гора» на синдарине, эльфийском языке, и отсылает к горе Ородруин из романа Дж. Р. Р. Толкина «Властелина Колец». Основной тематикой группы является германо-скандинавская мифология. Музыку Amon Amarth составляет сочетание гроулинга с использованием двойной бас-бочки, «тяжелыми» мелодиями и эпическими текстами.

Источник: Wikipedia (', 'https://lh3.googleusercontent.com/00tD_bn_OvaG6fvlHEmtg23ns_e3VWnk5sX7_W_rv8d-YoOFTUZd_sjHC9p3U7zUaltiHaBFDj0AhQ=w2880-h1200-p-l90-rj')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: The Great Heathen Army
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('The Great Heathen Army', 'https://lh3.googleusercontent.com/KhxpDlaRuNGDXlYwXsvLb6lAvYjUpTbm7HzEwix8aOqEN5xJ4VoQyoL2d5lDWaFOovOKARWuO5ZtlXpmug=w544-h544-l90-rj', v_artist_id, '', '2022-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 9
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Get in the Ring', 265, 'p2-_PW4kzQo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Great Heathen Army', 245, 'bK4MbGyCSXU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Heidrun', 283, 'tGeSEbYzhBU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Oden Owns You All', 258, 'B5HxkpDN4eA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Find a Way or Make One', 271, 'Je3EcJhxUP8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dawn of Norsemen', 333, '1kDgnsJ1RPI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Saxons and Vikings (feat. Saxon)', 296, 'PLzsWa20YIo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Skagul Rides with Me', 275, 'W3-maPjWwlQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Serpent''s Trail', 362, 'TBUjAKvuWuI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Berserker
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Berserker', 'https://lh3.googleusercontent.com/8Y8-1rHatnDC59DWy0kvzikey_e8APVDrkuFo2pQwkxPBTHp2-vkfOfHeis6HZSAJ0If26FsawvTEWtdiA=w544-h544-l90-rj', v_artist_id, '', '2019-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 12
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Fafner''s Gold', 301, 'bzmjCLKpwWE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Crack the Sky', 231, 'zetGwzBjVxM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Mjölner, Hammer of Thor', 284, 'cUW7dfoN5pE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Shield Wall', 228, 'DjdXNr1F-v8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Valkyria', 285, 'an7amlg5dLU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Raven''s Flight', 322, 'Xj0Lnf4MIsg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ironside', 272, '2-7uI_ADLdo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Berserker at Stamford Bridge', 316, 'sMPGEpvU_5o', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('When Once Again We Can Set Our Sails', 267, 'GTU6VSJPvH0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Skoll and Hati', 269, 'uip6w_sXjAI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Wings of Eagles', 245, '7tHC9-yjjWw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Into the Dark', 409, '03wLnq3ZDDQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Jomsviking
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Jomsviking', 'https://lh3.googleusercontent.com/MxYvhiXDndGKQvRCHSIYUVTeTssyraiXCnV84BuU3eqn7rLpwePi9YmSFu4pSP7k7G1SJHCiwTPENh0a=w544-h544-l90-rj', v_artist_id, '', '2016-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 11
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('First Kill', 262, 'qw5G6fF-wqQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Wanderer', 283, 'XUnIS3VoBTg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('On a Sea of Blood', 245, 'Q_WvMyEWMIs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('One Against All', 218, 'PnilK7ZRZB8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Raise Your Horns', 264, 'bGfgf-AmXfw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Way of Vikings', 312, '55OJ17cHeJA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('At Dawn''s First Light', 231, 'h6-krHfdmGg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('One Thousand Burning Arrows', 350, 'cgYNyWzPsu4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Vengeance Is My Name', 282, 'wpmo2s6Iz_U', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('A Dream That Cannot Be (feat. Doro Pesch)', 263, 'IINxwsSDf80', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Back on Northern Shores', 429, 'XyLDphQ7S4k', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Deceiver of the Gods
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Deceiver of the Gods', 'https://lh3.googleusercontent.com/2QWKgqfI1QpJg8WW3lZezTpeJOYPf8JgS_CV0qErvqtBNeELDEaRRQ2UlZvzC3v8mdr_NJI15yi0urQ=w544-h544-l90-rj', v_artist_id, '', '2013-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 10
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Deceiver of the Gods', 260, 'CpAcxbtXUgQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('As Loke Falls', 279, 'LpZiuh0MlpM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Father of the Wolf', 260, '9_qvTEg_F0w', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Shape Shifter', 243, 'gVsdHaGZSYI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Under Siege', 378, 'CvD5_bvEpXY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Blood Eagle', 196, 'DhhLPEVLcLk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('We Shall Destroy', 266, 're4TAtkvyRE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hel', 250, 'a5b6fmNl8Cg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Coming of the Tide', 257, '91uY8k5qnpI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Warriors of the North', 493, 'yzkedvnGXLs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Surtur Rising
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Surtur Rising', 'https://lh3.googleusercontent.com/MxQ-FauUCbjMqQCFb0E0-ZQsrRwC39lDZNiWcz-rBvsNkLBZrOrSYTom7ZuQRphKi0B5Kaent-h4yP0a=w544-h544-l90-rj', v_artist_id, '', '2011-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 11
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('War of the Gods', 274, 'FVAQQujgSxQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Tock''s Taunt - Loke''s Treachery Part II', 359, 'gGAsCpkFl_0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Destroyer of the Universe', 222, '5aaOqUYG8Tw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Slaves of Fear', 266, 'MyVHiHxPdjU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Live Without Regrets', 304, '4ky88quezSA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Last Stand of Frej', 338, 'dlrdzJDTlkE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('For Victory or Death', 271, 'ZLKRey-Zj64', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Wrath of the Norsemen', 225, 'AmV_cY_8J1M', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('A Beast Am I', 315, 'QsB93MrYzys', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Doom Over Dead Man', 356, 'DSq_30byIMA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Aerials', 220, 'VnoGQMJKDr4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Twilight Of The Thunder God
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Twilight Of The Thunder God', 'https://lh3.googleusercontent.com/Ljn5df4eQ2AErEH4tK8_a1waQq9cpGrjEH_hqUH_9weOOOd18DdwGs2zKUPl-lQueSY48EUSBj-qE12U=w544-h544-l90-rj', v_artist_id, '', '2008-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 10
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Twilight Of The Thunder God', 249, 'edBYB1VCV0k', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Free Will Sacrifice', 249, 'OjIYTgannxc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Guardians Of Asgaard', 264, 'ARnBgW5XgSo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Where Is Your God?', 192, '16CwvKD78aY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Varyags Of Miklagaard', 259, '9wo6lNn4KVE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Tattered Banners And Bloody Flags', 271, 'Ed1CB2DldGo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('No Fear For The Setting Sun', 233, 'pxgGws7mOP0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Hero', 242, 'I6SIHlK3LqQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Live For The Kill', 250, 'Bh_5ofa__pY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Embrace The Endless Ocean', 405, 'RKHBX4abO20', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: With Oden On Our Side
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('With Oden On Our Side', 'https://lh3.googleusercontent.com/01XlPVgN3PyXwiX4YhSGnjGVGPhVqKDGeQyF8rt1aa3QcbSsxOr7QzTG9GPtqtYyV6dYufqSWxHYqRF4eg=w544-h544-l90-rj', v_artist_id, '', '2006-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 9
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Valhall Awaits Me', 284, '3Or87hx0R7w', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Runes to My Memory', 273, '5S9iruQus1s', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Asator', 185, 'MufdQYcesg8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hermod''s Ride to Hel - Lokes Treachery Part 1', 281, 'fHey3HFynwo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gods of War Arise', 363, 'UFAJ--O4GCQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('With Oden on Our Side', 275, 'lDzjZrgrV3w', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Cry of the Black Birds', 230, '4vMHVOI32r4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Under the Northern Star', 258, 'woiPB5tjvZU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Prediction of Warfare', 399, 'lY4aopLf5sU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Fate of Norns
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Fate of Norns', 'https://lh3.googleusercontent.com/zmQLiKe1JrCV58DAW3xo-DnpfKynnzSL5tzV9_KGzNJ0BeS-H65JWWTzVvrG3W0XF2I0hRt5LhsMSu1Z0A=w544-h544-l90-rj', v_artist_id, '', '2004-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 8
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('An Ancient Sign of Coming Storm', 279, 'EGgaPDgibKI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Where Death Seems to Dwell', 299, 'KBF-K2YaVzQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Fate of Norns', 358, 'd7NXxfLkMCU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Pursuit of Vikings', 271, 'M-43pOqheMY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Valkyries Ride', 298, 'kUePeI_B1FA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Beheading of a King', 204, 'ESZPvDf9_Qc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Arson', 409, 'KuDMCwSMmLc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Once Sealed in Blood', 291, 'OiBViKa7itM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Versus the World
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Versus the World', 'https://lh3.googleusercontent.com/XNmGY7oEdW53efzU2yCGKIJDCgmoy6d4Ka9JCA-p5Y86EmRFFlmjHVhuaf6imdW7dHBMvNtCkN1kqCE=w544-h544-l90-rj', v_artist_id, '', '2002-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 23
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Death in Fire', 295, 'Z_ucAP9tRB4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('For the Stabwounds in Our Back', 297, 'TEaHPMeqdTc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Where Silent Gods Stand Guard', 347, 'nRbYPH2QCVk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Versus the World', 322, '2pW32SPLpic', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Across the Rainbow Bridge', 291, 'mhqn3QSedwY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Down the Slopes of Death', 249, 'Jpuw3Rg0Y1Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Thousand Years of Oppression', 342, 'poxWPyye3mQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Bloodshed', 314, '-B4Rwd567Rc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('...And Soon the World Will Cease to Be', 418, '4y7QVW3dscM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Siegreicher Marsch (Victorious March)', 474, 'EiLXpdFJyaU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Sorrow Throughout the Nine Worlds', 232, 'hQKRtODmL2A', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Arrival of the Fimbul Winter', 267, 'zt1H9DVIzAA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Burning Creation', 298, 'imjlIZ4wf9Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Mighty Doors of the Speargod''s Hall', 343, '4o13ga540nM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Under the Grey Clouded Winter Sky', 335, 'KZ6Hoya591w', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Burning Creation (Demo Version)', 288, '1HWZ2XtIpAA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Arrival of the Fimbul Winter', 277, 'nAHYMUyx3Iw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Without Fear', 282, 'vKB8ppJGCno', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Risen from the Sea', 343, '3liPrJIruOE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Atrocious Humanity', 354, 'pWIOx-44z6E', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Army of Darkness', 325, 'DaQaaB4ooGw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Thor Arise', 391, 'liKMtwrrBnM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Sabbath Bloody Sabbath', 263, 'xtg13SO2xpQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Versus the World (Bonus Edition)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Versus the World (Bonus Edition)', 'https://lh3.googleusercontent.com/Qhl_wa3kJpUmLVKQkB1F0Oet1fBCekq6soxTOexDSS99OowKAA2pXq5Rn1HTvASeoQHiGKg5Pb1XId6c=w544-h544-l90-rj', v_artist_id, '', '2002-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 18
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Death in Fire', 295, 'Z_ucAP9tRB4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('For the Stabwounds in Our Backs', 297, 'gV5lWFxC0N8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Where Silent Gods Stand Guard', 347, 'E2N89AreXrk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Versus the World', 322, 'rzFCDX3m0DY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Across the Rainbow Bridge', 290, 'K1LG2336WRc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Down the Slopes of Death', 249, '-zzFxEaN_X4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Thousand Years of Oppression', 342, 'dqTFSV9xRo0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Bloodshed', 314, 'K7zdr9hvrmo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('...And Soon the World Will Cease to Be', 421, 'CL8UjjpWeD4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Death in Fire (Live)', 309, 'jFmQoItSWKY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('For the Stabwounds in Our Backs (Live)', 320, '17sE7rvRSsg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Where Silent Gods Stand Guard (Live)', 350, 'atV4jBwHeds', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Versus the World (Live)', 358, 'DVeZcyfXMrc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Across the Rainbow Bridge (Live)', 308, '9M4qLGz14Bw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Down the Slopes of Death (Live)', 279, 'ZQ-B_AMIK8g', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Thousand Years of Oppression (Live)', 368, 'dAWF9HhYh8Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Bloodshed (Live)', 344, '2TvFp9pfY9Y', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('...And Soon the World Will Cease to Be (Live)', 428, 'UKi9LfzfrUw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Crusher
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('The Crusher', 'https://lh3.googleusercontent.com/y52ciOaoFKu0J4ecnoDDU8SqmbuZyu0AKoh3IB3X8jNQJvuW2nnjxl0TzCA7VLOzAqL-QT76xbSxPybdDA=w544-h544-l90-rj', v_artist_id, '', '2001-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 10
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Bastards Of A Lying Breed', 334, 'CyRxdBfkLIw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Masters Of War', 274, 'wcI0BNVvCHQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Sound Of Eight Hooves', 291, 'fIiNB8V5QMU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Risen From The Sea (2000)', 267, 'fOJT_V5_2jE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('As Long As The Raven Flies', 245, 'yzJVnGw0v_Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('A Fury Divine', 396, 'GVFYxzfwDSY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Annihilation Of Hammerfest', 303, 'TvvK-aPct7E', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Fall Through Ginnungagap', 322, 'hwAy58j1LaQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Releasing Surtur''s Fire', 327, 'ul4PqtxH7jE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Eyes Of Horror', 215, 'OCBbzow4S3g', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Avenger
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('The Avenger', 'https://lh3.googleusercontent.com/h9-Mapb417VXNwcsI5B9f03Wyz5ck7ydTQ0tRStilVE6sEUQTQQzwmtYSVI_pbeE0gfH3EGUs_swkChcLA=w544-h544-l90-rj', v_artist_id, '', '1999-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 7
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Bleed For Ancient Gods', 271, 'PKnQeVhiz3M', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Last With Pagan Blood', 340, 'Ju1wdQMqz9I', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('North Sea Storm', 296, '0YpZYQ_z9wA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Avenger', 432, 'GhLYMxRDytY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('God, His Son And Holy Whore', 241, '32DTKmax6mo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Metalwrath', 230, 'fiwwY3XJDQA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Legend Of A Banished Man', 365, 'W0ndajOQ0b0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Once Sent From The Golden Hall
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Once Sent From The Golden Hall', 'https://lh3.googleusercontent.com/6PS9yqHmVGNyImLaVuU5yN0IJ2Qe7Jysj0Qqh_lSJqJS577OxQOPqTTrBEQb4Mp-NqvboWkqdeobzQqAOg=w544-h544-l90-rj', v_artist_id, '', '1998-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 8
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ride For Vengeance', 269, 'MRqqfzxKw5E', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Dragons'' Flight Across The Waves', 274, 'LEy9JJQD6VM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Without Fear', 291, 'tfjCVWCr3fI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Victorious March', 477, 'WYTkNlxjA2c', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Friends Of The Suncross', 283, 'srSIraHx9LE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Abandoned', 361, '1AfkTzbFgf0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Amon Amarth', 487, 'UGK6Ro6Rn7E', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Once Sent From The Golden Hall', 252, 'SUmroHlFQuI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Pursuit of Vikings: Live at Summer Breeze
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('The Pursuit of Vikings: Live at Summer Breeze', 'https://lh3.googleusercontent.com/J2LpedqJxerWC6sTWhf64QtXvJ5P2vKyXzWvD-R6QEHEaDhiY_NFctmTj8scWuxNPgP35Qy-TI45dzWj=w544-h544-l90-rj', v_artist_id, '', '2018-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 30
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Twilight of the Thunder God (Live at Summer Breeze: T-Stage)', 263, 'Sw9bdwJmiW4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Free Will Sacrifice (Live at Summer Breeze: T-Stage)', 359, 'RX56wRFjtWU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('With Oden on Our Side (Live at Summer Breeze: T-Stage)', 311, 'dCuvr3_w3mk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Last with Pagan Blood (Live at Summer Breeze: T-Stage)', 308, 'PUIryD6eniY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('For the Stabwounds in Our Backs (Live at Summer Breeze: T-Stage)', 291, 'naoIbwhlBOY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Thousand Years of Oppression (Live at Summer Breeze: T-Stage)', 407, 'HPrgLzeeoec', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gods of War Arise (Live at Summer Breeze: T-Stage)', 386, 'MYL-p-qEcQ4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Versus the World (Live at Summer Breeze: T-Stage)', 349, 'KzoSumqiWIk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Asator (Live at Summer Breeze: T-Stage)', 196, '0sEnuWjNFAc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Under the Northern Star (Live at Summer Breeze: T-Stage)', 378, '5QsN-ugQIM4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Fate of Norns (Live at Summer Breeze: T-Stage)', 382, '_EKT1uVO0Xs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Varyags of Miklagaard (Live at Summer Breeze: T-Stage)', 270, 'KRsF5b_VnhM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Live for the Kill (Live at Summer Breeze: T-Stage)', 257, 'IyXURWB3NNk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Victorious March (Live at Summer Breeze: T-Stage)', 599, 'QiSy63UglBc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Pursuit of Vikings (Live at Summer Breeze: Main Stage)', 379, 'Bc4v-5l30Mc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('As Loke Falls (Live at Summer Breeze: Main Stage)', 265, 'tJo_y_F4ZEs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('First Kill (Live at Summer Breeze: Main Stage)', 268, 'TrEQ-Pzi9RI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Way of Vikings (Live at Summer Breeze: Main Stage)', 319, '8gMGONMgzkE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('At Dawn''s First Light (Live at Summer Breeze: Main Stage)', 225, '1ZJSmBFcmeo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Cry of the Black Birds (Live at Summer Breeze: Main Stage)', 283, '39K7GgcxpKo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Deceiver of the Gods (Live at Summer Breeze: Main Stage)', 254, 'Jy6PM_WWIts', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Destroyer of the Universe (Live at Summer Breeze: Main Stage)', 222, 'KwOIFyAJgb0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Death in Fire (Live at Summer Breeze: Main Stage)', 304, 'gSiELmqBysA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Father of the Wolf (Live at Summer Breeze: Main Stage)', 336, 'g39UzXafFHQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Runes to My Memory (Live at Summer Breeze: Main Stage)', 285, 'VA0XwuXokmY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('War of the Gods (Live at Summer Breeze: Main Stage)', 286, 'xi1FheXCH1I', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Raise Your Horns (Live at Summer Breeze: Main Stage)', 275, 'jVOluNDjc5s', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('A Dream That Cannot Be (Live at Summer Breeze: Main Stage)', 273, 'qLF4gMLpKhw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Guardians of Asgaard (Live at Summer Breeze: Main Stage)', 284, '1Cy_HpQavrE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Twilight of the Thunder God (Live at Summer Breeze: Main Stage)', 414, 'bW9U8C85gtA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: We Rule the Waves
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('We Rule the Waves', 'https://lh3.googleusercontent.com/lbjdSKJYYD6NCYeCQ8cmwe0n89U_u_2yR9782oT26VPqLOHDKGVTB0PTq7RnJLIxLtg1OzBf1RDIDHaI_w=w544-h544-l90-rj', v_artist_id, '', '2025-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('We Rule the Waves', 231, '5VLMKHikNrM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Heidrun
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Heidrun', 'https://lh3.googleusercontent.com/qlZcRiiFUNQjoIVZJ8nlMqd4V6Wv3ObgtRWAW3WS9K8QHy5ZwS1eBKt8rzeL47N4QtRn82t_MqfKddtY7w=w544-h544-l90-rj', v_artist_id, '', '2023-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 4
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Heidrun (2023 Remix)', 283, 'tGeSEbYzhBU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Heidrun (Live at Graspop, 2023)', 299, 'ZxCE8_7YSWk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Put Your Back into the Oar (Live at Hellfest, 2023)', 298, 'XOBDPiTiC9E', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Heidrun (Goat Remix)', 283, 'dZ2Jm3-r8G0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Great Heathen Army
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('The Great Heathen Army', 'https://lh3.googleusercontent.com/PyHXH4GXtkE3LmjzJgDN0Wn3f5ewgnvkEPfmuwtw5brrgG5HYQM-cP-7MHlceDVGb5eQX2hw9BtS5MKELg=w544-h544-l90-rj', v_artist_id, '', '2022-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Great Heathen Army', 245, 'bK4MbGyCSXU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Get in the Ring
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Get in the Ring', 'https://lh3.googleusercontent.com/PIm-shwg9VFNLGhqwH-o5YnnbsTBqgsUDAce0EAaXAeWER6EawoLvRhxTomRPYKFr63WiWVhVLLnGw8D=w544-h544-l90-rj', v_artist_id, '', '2022-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Get in the Ring', 265, 'p2-_PW4kzQo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Put Your Back Into The Oar
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Put Your Back Into The Oar', 'https://lh3.googleusercontent.com/BQZiZCbCm3Qcvy2fUkOYfhpc3VUM2x0je06QfgUNwxxAXWBt0bl5Nmj-wsh10120U7SczIq6vOfDZvE=w544-h544-l90-rj', v_artist_id, '', '2022-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Put Your Back Into The Oar', 278, 'vYNAEzgKNec', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Masters of War
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Masters of War', 'https://lh3.googleusercontent.com/w53cS4ezbEdj7KjP79zkSJyg8_7V9Q9sn-zXXIBlV8UW1UkfEVwWze1InyX88iyJQx0PpBkG_OsA77VT=w544-h544-l90-rj', v_artist_id, '', '2021-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Masters of War', 275, '0H9pJdjTGQw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Crack the Sky
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Crack the Sky', 'https://lh3.googleusercontent.com/6Zw5D6bplnzhs8B9zKnT1X3k8F-H61pO_zEqo_dTnNE_aluc3TgWDjQvYzYDT_xRS0kXFL1MiWNOmuLs=w544-h544-l90-rj', v_artist_id, '', '2019-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Crack the Sky', 231, 'zetGwzBjVxM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Raven''s Flight
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Raven''s Flight', 'https://lh3.googleusercontent.com/AyyW3_Z2iEDWac1RwHSBrbam-myNXyIT8ExJYJkjYUbEF8c_NfJPU368N8-AQwHnZo20iF9MY_DZK4y_=w544-h544-l90-rj', v_artist_id, '', '2019-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Raven''s Flight', 322, 'Xj0Lnf4MIsg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: At Dawn''s First Light
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('At Dawn''s First Light', 'https://lh3.googleusercontent.com/KXtF01kYMoDmY-r7gk0ycXI3vDrRuqlGNccac91whgvDFz4J3WF6utOfbpD7H6Jz8eqqJpV8UwBli20J=w544-h544-l90-rj', v_artist_id, '', '2016-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('At Dawn''s First Light', 231, 'h6-krHfdmGg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: First Kill
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('First Kill', 'https://lh3.googleusercontent.com/AnJXDh0F4cMvCWGzitaKTrYCJdaolZ1jExGdCO-Z3yYVibVdd4e4_XNSwVdi4jj_70wMVl4rL6i8S3Y=w544-h544-l90-rj', v_artist_id, '', '2016-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('First Kill', 262, 'qw5G6fF-wqQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Under the Influence - EP
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Under the Influence - EP', 'https://lh3.googleusercontent.com/nej3xh4GI2a7A0-fVmVxpv8QjB8OsV9DLoZ8q8cQ7ZEgoUAVXSWl_RG3aIJ3XaJzKCSCgR9RyIwpwdjm=w544-h544-l90-rj', v_artist_id, '', '2013-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 4
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Burning Anvil of Steel', 268, 'k5M6kiWy_Ng', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Satan Rising', 261, 'goPL7zjfyW8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Snake Eyes', 193, 'UeYY20-YCIU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Stand up to Go Down', 208, '3p5GsfVHkkQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Raise Your Horns (Live at Summer Breeze: Main Stage)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Raise Your Horns (Live at Summer Breeze: Main Stage)', 'https://lh3.googleusercontent.com/Xo44CNXlARs6ArmuJL-paicoENQGs3BbRLQB0urHmTDgyPK9xq9TjRS1AmbzuZeYmVOYUi4tYAiuuVT5og=w544-h544-l90-rj', v_artist_id, '', '2018-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Raise Your Horns (Live at Summer Breeze: Main Stage)', 275, 'jVOluNDjc5s', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Twilight of the Thunder God (Live at Summer Breeze: Main Stage)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Twilight of the Thunder God (Live at Summer Breeze: Main Stage)', 'https://lh3.googleusercontent.com/qQuxjH3WaaWZ7rG4iKBEGvl7Bs5M8gOfpIdxp3EauSzRbyhiPVwvOxV5cV4WIAjIqVWwNuvJ4C8Ccqnh=w544-h544-l90-rj', v_artist_id, '', '2018-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Twilight of the Thunder God (Live at Summer Breeze: Main Stage)', 414, '8KopXlUHLFc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
-- Артист: Burzum
INSERT INTO artist (artist_name, description, avatar_url)
VALUES ('Burzum', 'Burzum — музыкальный проект музыканта и писателя из Норвегии Варга Викернеса. Название проекта было взято из трилогии Дж. Р. Р. Толкина «Властелин колец». Слово «burzum» переводится как «тьма» на чёрном наречии — вымышленном языке, придуманном Толкином.
Викернес сочинял свою музыку ещё в 1988 году, до основания Burzum. Проект Burzum стал неотъемлемой частью ранней блэк-метал-сцены, повлияв на его развитие наряду с такими группами, как Mayhem, Darkthrone, Immortal, Emperor и Gorgoroth.
Свои первые 4 релиза Викернес записал в период с января 1992 по март 1993, однако они были распространены спустя долгие месяцы между записью и выпуском каждого альбома. В мае 1994 года Варг Викернес был приговорён к лишению свободы в тюрьме сроком на 21 год за убийство Эйстейна «Евронимуса» Ошета — гитариста группы Mayhem, а также за поджог трёх церквей.
Во время своего тюремного заключения Викернес продолжает свою музыкальную деятельность и записывает 2 студийных альбома в жанре дарк-эмбиент, используя только синтезатор ввиду того, что другие инструменты не были разрешены тюремной администрацией.

Источник: Wikipedia (', 'https://lh3.googleusercontent.com/XJIqLgQ0togNSv67SiBLwwMRns8Z_R8wSfc89XgYyHMGu5HUoG362_IRvBhAMG9HBFEVXToCHq50sBE=w700-h291-p-l90-rj')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Thulêan Mysteries
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Thulêan Mysteries', 'https://lh3.googleusercontent.com/S6ikTda2Aos2Ex9bAKDcD8hW9w739q5gYuf58L-3dd6t1bv9S8UKuTPXwTxzMPXuODIscuL8hRVy2k4=w544-h544-l90-rj', v_artist_id, '', '2020-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 23
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Sacred Well', 177, '9aK685uqsUM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Loss of a Hero', 56, 'A_6Hx6gDUsI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('ForeBears', 245, 'eyjvrkhPtgg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('A Thulêan Perspective', 243, 'bZhm3yGOAug', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gathering of Herbs', 76, 'v7lWBgDiI4Y', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Heill auk Sæll', 219, '-4oETClJyAk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Jötunnheimr', 100, 'GTsxcw9vwGY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Spell-Lake Forest', 68, 'SpxpCpSP8eI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Ettin Stone Heart', 77, 'b4vmiTeo6Jk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Great Sleep', 90, 'NlVLU-oGD-Y', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Land of Thulê', 135, 'tSUulYTqY5g', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Lord of the Dwarves', 316, 'dGyGy_f92WE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('A Forgotten Realm', 446, 'z8x2w_DZ2UU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Heill Óðinn, Sire', 80, 'x-JKWRkNoV4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Ruins of Dwarfmount', 92, 'QCjc3v79oWs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Road to Hel', 465, 'TUbH9ep8CYg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Thulêan Sorceryl', 132, '80UDfff0SZc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Descent into Niflheimr', 103, 'yVO23CB5xXc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Skin Traveller', 278, '25KvHg5NlTY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Dream Land', 525, 'kZNPXK9jnt8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Thulêan Mysteries', 265, '8z4UoEIn9Tw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Password', 915, 'fo_e0517iW4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Loss of Thulê', 305, 'jg6PViOB6J8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Ways of Yore
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('The Ways of Yore', 'https://lh3.googleusercontent.com/O-n31IZtCiURCtverMCfBhjaynGBG8ujsrLRT6b1u1NyRH5vCNzd_nuAgYQqBd-F-EqYoBbGO9KetXiH=w544-h544-l90-rj', v_artist_id, '', '2014-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 13
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('God from the Machine', 101, 'n2RUz3uiw0U', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Portal', 140, 'Sk5xZ1Tov60', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Heill Odinn', 191, 't--E1DtWBuA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Lady in the Lake', 280, '6LzAa6oNYms', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Coming of Ettins', 276, '2tZkSlgGnVg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Reckoning of Man', 436, '6k4e2Ir7BE4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Heil Freyja', 117, 'jSEnYFurJdY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Ways of Yore', 372, 'cgRT4P94Fng', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ek Fellr (I Am Falling)', 174, 'mCbuIJ0JMOo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hall of the Fallen', 306, 'oPW-yQBd7ik', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Autumn Leaves', 290, '6RSWgX5gJKg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Emptiness', 793, 'Pjw9E4RXZ7E', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('To Hel and Back Again', 645, 'lf5jNze9lbQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Cold Empire of Negativity
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Cold Empire of Negativity', 'https://lh3.googleusercontent.com/at_wINw9_-Py6-_uh984VpPPjSm8t9zOgofh2uR16wd4-aoY0eZZ0OUUJ7HUi47-ffs3NSBy66qM_AQ=w544-h544-l90-rj', v_artist_id, '', '2014-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 6
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Then so Be It', 174, 'tUInNGcsxWc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('If This Is to Be My End', 386, 'VC-DdovNivc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I Am Soil', 282, 'hj9YI9h5ns0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dunkelheit', 422, 'hmFfy9c3Ji4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Black Page of Suffering', 247, 'I7SgPJLlrwQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Endless Misery', 321, 'D5dQpz1VpK4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Umskiptar
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Umskiptar', 'https://lh3.googleusercontent.com/_-bV79qvmAAM9HhxCXWgR-ZhGGGm19QVhmQjD37APwGoBeJnGg39dumMUcxXG50NKx9xD8KuD2iLpON8=w544-h544-l90-rj', v_artist_id, '', '2012-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 11
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Blodstokkinn', 71, 'Z-U63yiqFAM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Joln', 347, 'oVcEtMiJH7s', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Alfadanz', 559, 'iFMjdgu3A4A', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hit helga Tre', 408, 'qwKIUxy5XbA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Aera', 235, '1leD2CbnA5E', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Heidr', 178, 'cuzl4HeqJxc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Valgaldr', 478, 'ZQKI3wKTYb8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Galgvidr', 434, 'R9-iImwHpEM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Surtr Sunnan', 252, 'xDq7FwYNSUQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gullaldr', 615, 'XbYzMINUNK8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Nidhoggr', 300, 'jG8RUjzT9jw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Fallen
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Fallen', 'https://lh3.googleusercontent.com/MX8WC53VTyEViMNdceUmlioDwgIzwhzVj84GcaH-p_MDki81UyGzYq1RRHzLTqYtC0YUsU0bIbnfCYA=w544-h544-l90-rj', v_artist_id, '', '2011-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 7
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Fraverdenstreet', 64, 's9_6f0LxmNI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Jeg faller', 471, '7jqAhRh6dmY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Valen', 562, 'dYaZEGF9-jc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Vanvidd', 427, 'IevGARbAsMo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Enhver til Sitt', 377, 'ARdQYuLa2RE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Budstikken', 611, '0vkNLjRXLyk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Til Hel og tilbake igjen', 358, 'HyawUThShxI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Belus
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Belus', 'https://lh3.googleusercontent.com/S25k-YddVaNs5HlFUdGCaOiR5_aTUGdUp1m3BB37-CsxhLTa28T93zzXbTlWhO27O3ejU-RhkMniiXjy=w544-h544-l90-rj', v_artist_id, '', '2010-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 8
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('i. Lukans Renkespill (Introduksjon)', 34, 'AySqIzXhCuQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('ii. Belus Doed', 384, 'jI5-XCETpSk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('iii. Glemselens Elv', 715, 'dJZmix_Jgzg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('iv. Kaimadalthas Nedstigning', 404, 'hwI-EaIUdeQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('v. Sverddans', 148, 'XpU3pWMy4Kc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('vi. Keilohesten', 346, 'F6CAXGZMDzI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('vii. Morgenroede', 535, 'M-MdBwq4mbE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('viii. Belus'' Tilbakekomst (Konklusjon)', 578, 'Ir1t8bARi-Y', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Draugen - Rarities
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Draugen - Rarities', 'https://lh3.googleusercontent.com/NTJnFOlgd2CbiyhJCAh4nUB3l1AxxNFAQVNvi2jDLzAFnke13RbKIqIUfK2x5xYP8RPmeeQ3rfRLCzU=w544-h544-l90-rj', v_artist_id, '', '2008-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 13
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('A Lost Forgotten Sad Spirit', 551, 'Syzf8pyh1EE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Steminen Fra Taarnet', 370, 'fefMxyuwV9Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dominus Sathanas', 185, 'SoAFfKC_s2M', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Lost Wisdom', 293, 'XLZ-ROPoR_k', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Spell Of Destruction', 299, 'NKrIbRu0BJw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Channelling The Power Of Souls Into A New God', 241, 'cacIxbc5BMY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Outro', 118, 'jzPrE_5qjzw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Et Hvitt Lys Over Skogen', 549, 'xt-vzBiLfuI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Once Emperor', 375, 'D8G5tM0FogU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Seven Harmonies Of The Unknown Truth', 190, '9ZUDsXCPCwY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('My Journey To The Stars', 491, '3jlZ-QUnITo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Lost Wisdom (Demo)', 279, 'u1273LqED5Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dunkelheit', 426, 'ZE12_ja9GK0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Hlidskjalf
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Hlidskjalf', 'https://lh3.googleusercontent.com/hpvgz8IUsKHBfx2l91lrwGstU6SNm0Q6PuSYErJwazs-H4ouW3qOewD1HX8G76t8YaTpOQnG1ery9NQ93Q=w544-h544-l90-rj', v_artist_id, '', '1999-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 8
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Tuistos Herz', 374, 'mbLJrkqbnm4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Der Tod Wuotans', 404, '2aIynRhhA-s', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ansuzgardaraiwo', 269, 'CxnYiY-jyM4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Die Liebe Nerpus', 135, 'aGwopNSlzPI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Frijos Einsames Trauern', 376, 'Ns4OvPi1KFw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Einfuhlungsvermogen', 236, 'u3cwkCaqat0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Frijos Goldene Tranen', 159, 'jyPS4HW2-P4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Der Weinende Hadnur', 77, '-i5dHTDTYAg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Filosofem
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Filosofem', 'https://lh3.googleusercontent.com/4Nz0kYoCOT1pKOWMO_tJN9E26_IDy4zqsZ4vgCpD_8KY_2YxsiL7PExCW-z-PkOjIqQQYXIAbRmB-t7K=w544-h544-l90-rj', v_artist_id, '', '1996-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 6
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dunkelheit', 426, '-ZENtivAi6I', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Jesus'' Tod', 520, 'vfBmYfnjLho', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Erblicket Die Tochter Des Firmaments', 474, 'oPp2mIv4CyA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gebrechlichkeit I', 473, 'dwlh1e0ffGI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Rundgang Um Die Transzendentale Saule Der Singularitat', 1512, 'bn6lfljADKs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gebrechlichkeit II', 473, 'WkLC2x1_LrY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Det Som Engang Var
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Det Som Engang Var', 'https://lh3.googleusercontent.com/yXDJCG2kM8G0W-sniyhEflLNlK1StjSOa1vpealvpUGVhyYS7pYStc20kxBrn22wD7hKKFVWNuHRPmPX=w544-h544-l90-rj', v_artist_id, '', '1993-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 8
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Den Onde Kysten', 141, 'MpyH7HeYkxQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Key To The Gate', 315, '_8ziktiI3jc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('En Ring Til Aa Herske', 431, 'codih99MoWo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Lost Wisdom', 279, 'Qz4me6YlI9w', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ham Som Reiste', 292, 'CckPL8fg3xk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Naar Himmelen Klarner', 231, 'eUBSHOeyabo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Snu Mikrokosmos Tegn', 577, 'UExvq548ByA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Svarte Troner', 137, '3ySA617b79A', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Sol Austan, Mani Vestan
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Sol Austan, Mani Vestan', 'https://lh3.googleusercontent.com/vXXijSzYJ9Bcw9pTH48Ftn1b0J2FGZhLwKRVDZkOMhzOGVErzByyrU87ZFCjb9PtFks4PCprkfqsoVNrJw=w544-h544-l90-rj', v_artist_id, '', '2013-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 11
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Sol Austan (East of the Sun)', 263, 'KHFPJFkgjp8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Runar Munt Bu Finna (You Shall Find Secrets)', 208, '0Y0m93HYwas', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Solarras (Sun-Journey)', 244, 'cdZFF2JXt3o', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Haugaeldr (Burial Mound Fire)', 447, 'oC-_Ktc7Pac', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Fedrahellir (Forebear-Cave)', 321, 'jpiF_irI2zs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Solargudi (Sun-God)', 432, 'PghjcyTg7KI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ganga at Solu (Deasil)', 358, 'ignHTmtlnNs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hid (Deasil)', 384, 'FRgUm8NZobw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Heljarmyrkr (Death''s Darkness)', 242, 'qaurUViqHho', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Mani Vestan (West of the Moon)', 354, 'hpIJoc_uISM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Solbjorg (Sunset)', 240, 'pkF9atEJAoY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: From The Depths of Darkness
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('From The Depths of Darkness', 'https://lh3.googleusercontent.com/HRraGkBPLki_Ey_sqPfjZ3YoCQ1fIuyXb82-OWuCHiuUudNZlwSoO1JX3StVFxD7XMAr09IerSg5XwzINg=w544-h544-l90-rj', v_artist_id, '', '2011-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 11
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Coming (Introduction)', 26, 'sd0sLrxI0Jo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Feeble Screams from Forests Unknown', 469, 'u5oejpovCBk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Sassu Wunnu (Introduction)', 45, 'tFRXMPFtrlA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ea. Lord of the Depths', 324, 'Na7jCsSh2Oc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Spell of Destruction', 408, 'Pf57yy74BH8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('A Lost Forgotten Sad Spirit', 691, '0zWn_4kq_4k', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('My Journey to the Stars', 472, 'Et3CFZstROg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Call of the Siren (introduction)', 121, 'OwlzxsSyCnw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Key to the Gate', 315, 'AQiIAb6Eip0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Turn the Sign of the Microcosm (Snu Mikrokosmos'' Tegn)', 591, 'yzHaUGfuHto', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Channeling the Power of Minds Into a New God', 297, 'kMkCjCggRYc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Anthology
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Anthology', 'https://lh3.googleusercontent.com/YFcQ-k7Nu7g6V3EaDQwlnbJ7rH8myKNI88mTK4UHKJkI-RDoM1vvmcQ-e7D5JC9wCvESTFRAQjswOauScg=w544-h544-l90-rj', v_artist_id, '', '2008-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 9
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Feeble Screams from Forests Unknown', 450, 'PpVbQIqMgH0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Steminen Fra Taarnet', 370, 'L2YUw2OWTbw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Lost Wisdom', 280, 'Kd64qrWUnnk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Svarte Troner', 138, 'Z2B_T5d_h30', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Det Som Er Gang Var', 863, 'nZPOrp-UiU8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Jesus'' Tod', 521, 'QoyvFowe8G8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gebrechlichkeit (Ii)', 475, 'FqM3my5pgYs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Balterd Baldrs', 365, 'VmJPnYXCHBA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ansuzgardaraiwo', 269, '0pNT9QQXD3Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Burzum/Aske
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Burzum/Aske', 'https://lh3.googleusercontent.com/L9w_FPmPuqI0l8gwKIyunSFfT_SrgXIPtE6fRx5l96pL4Z6WoWSKp5q-BtcYO1CvTjjbIvHyWgQbjk9k=w544-h544-l90-rj', v_artist_id, '', '1992-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 11
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Feeble Screams from Forests Unknown', 449, 'qCnj45m9-TU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ea, Lord of the Depths', 293, 'JX68guGkdWs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Spell of Destruction', 340, '2m1BP4B5y8U', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Channelling the Power of Souls into a New God', 208, 'NGOQdHJTLAM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('War', 151, 'ytA2MrPEZdA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Crying Orc', 58, 'l4wSqSKvYjg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('My Journey to the Stars', 491, 'QAC3uw-r2AI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dungeons of Darkness', 293, 'jneBP6sa-4o', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Stemmen Fra Taarnet', 370, 'cKytrkDXVaw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dominus Sathanas', 185, '442WnT4rETA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('A Lost Forgotten Sad Spirit', 653, 'lyFFpb3SZ3k', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Reincarnation of Ódinn
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('The Reincarnation of Ódinn', 'https://lh3.googleusercontent.com/sCoxpSiw34TmwMOMOnAIl7k3jwTMgnG7CF3kkHupI41Mb1WPGKeddxenAmxu32VbzOGpBGbyGjtg4QcLaA=w544-h544-l90-rj', v_artist_id, '', '2019-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Reincarnation of Ódinn', 153, 'Jjk8Eb7WA9U', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Thulean Mysteries
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Thulean Mysteries', 'https://lh3.googleusercontent.com/R_ga0-fXdzuE7EghuXmsm2oULUIMInvVJNgo9ARWPMtvlq4tbPnJky86Lzyx0fEy2zTh5u-BJq8bA5Ox=w544-h544-l90-rj', v_artist_id, '', '2002-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Thulean Mysteries', 262, 'Ge65b4Mno68', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Forgotten Realms
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Forgotten Realms', 'https://lh3.googleusercontent.com/HLx83KqqDuOm-ROMHaIlG9H_GEkJWJ2_qXlFDCZ_34aGj3Y2cqOC5rEqI-Y0w7iArc-QKQ1yUpHCYJHt=w544-h544-l90-rj', v_artist_id, '', '2007-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Forgotten Realms', 443, 'GNTARIvvdPg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Mythic Dawn
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Mythic Dawn', 'https://lh3.googleusercontent.com/C-OtRxSH3EmPveuMHuHd_i9LARyBHnqLCHPpvrsE8O6HL-2LVW0LJ2OTDsSBnasNBs1o2Btm23SF9DDT=w544-h544-l90-rj', v_artist_id, '', '2017-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Mythic Dawn', 313, 'bUSE-vEPkRM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Daudi Baldrs
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Daudi Baldrs', 'https://lh3.googleusercontent.com/Bv-m-DNHvFj6vZ4Vy-ga3YL8sH6qb7jJV06Jr1njdoRSmebWUmwwOnbcR5v9qLMUVZ6Mu1OoI01En_-v=w544-h544-l90-rj', v_artist_id, '', '1995-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 6
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Daudi Baldrs', 530, 'snM1zaUrZFA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hermodr A Helferd', 162, 'jPJVG0n0EPE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Balferd Baldrs', 366, '82n3Qd1KufY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I Heimr Heljar', 123, 'krFCpNPDfnk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Illa Tidandi', 630, '8zXHY26PXPE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Moti Ragnarokum', 545, 'c7WsgIF-DNs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Jesu Død (Cover)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Jesu Død (Cover)', 'https://lh3.googleusercontent.com/V9UgFUGzlnGnz3B-rC_np04JHBeHRg78yD_IecJW-cPE3bddnq515swT7Rsf3LHWvfE_VXBT06kztMefxg=w544-h544-l90-rj', v_artist_id, '', '2008-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Jesu Død (Cover)', 497, 'IWl-g9CmB3c', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Hvis Lyset Tar Oss
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Hvis Lyset Tar Oss', 'https://lh3.googleusercontent.com/e1rkpz7nuiOlP8HufNdPXqd4rmk_R8f2Dju0qOMj3snkenEOV8je5mDvCCH0bOYmPv8rudLIgWMp55Hm9w=w544-h544-l90-rj', v_artist_id, '', '2008-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 4
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Det Som En Gang Var', 862, 'jw6-LzuEvb4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hvis Lyset Tar Oss', 485, 'MkmCEtaHJcQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Inn I Slottet Fra Droemmen', 472, 'xuzsTrTc91A', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Tomhet', 852, 'aTtNe8E0AoQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
-- Артист: Валентин Стрыкало
INSERT INTO artist (artist_name, description, avatar_url)
VALUES ('Валентин Стрыкало', '«Валенти́н Стры́кало» — украинская рок-группа, основанная в 2010 году солистом Юрием Капланом, который получил известность после записи серии видеообращений к звёздам шоу-бизнеса от имени наивного провинциального парня «Валентина Стрыкало из села Бурильцево».
Каплан основал группу под влиянием музыки таких рок-групп, как «Сплин» и Radiohead. В 2012 году был выпущен дебютный альбом «Смирись и расслабься!», который по большей части был написан в жанре камеди-рок, пусть в нём и имелись песни с серьёзным характером. В следующем году вышел альбом «Часть чего-то большего», в котором упор был сделан больше на лирику, но также не обошлось и без завуалированного юмора в некоторых песнях. В 2016 году выходит альбом «Развлечение», в котором группа окончательно отошла от юмора и ушла к теме депрессии. Также в записи заметно, что группа вдохновлялась британской рок-группой Pink Floyd.
Летом 2018 года группа дала последние концерты и ушла в затишье, а в мае 2019 года Юрий Каплан сообщил, что группа распущена.

Источник: Wikipedia (', 'https://lh3.googleusercontent.com/HIrTtjqJy2GkMz7CPWfwMuZdxfAp12HWG0y8q_kKK25BnyZKkVZ3RJxqWrHFFhoGJZHCXwJQ2zLpZmSW=w2560-h1066-p-l90-rj')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Развлечение
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Развлечение', 'https://lh3.googleusercontent.com/5r_0gDrkPWPC3B8sjE8zxzfo36m8dP45iSeH1TnbFvO6ZQ0k0yrIKXvxaSTv6f3214kowuw_MHq4kKJVfg=w544-h544-l90-rj', v_artist_id, '', '2016-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 8
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ускользает', 306, 'PZZxcKYskpg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('О брат', 295, '8JAfQuzbOBM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('92', 230, 'e8mLnAVQ9mY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Решится само собой', 253, 'kOfBszWa0t4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Делать это трезвым', 194, 'nAKX-7_fbw8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Тени', 345, 'FiynOYlSfrk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Бесполезно', 425, 'KdqzgAKoHTU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Подворотня - мой дом', 202, 'dFuIywTOJxc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Часть чего-то большего
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Часть чего-то большего', 'https://lh3.googleusercontent.com/W20_TGvzJCDuOlE46_yZYDSr1HILhA8EUeAg6r6KXFIuVX_rZFbwOrwrUkmtgYo9zW3CyN0CM5XTpkGRdQ=w544-h544-l90-rj', v_artist_id, '', '2013-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 14
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Самый лучший друг', 166, 'a_NbJCJoUM4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Офисный стиляга', 134, 'CmHO2FKYGTs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Все мои друзья', 161, 'C7_oi97EzbM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Знаешь, Таня', 274, 'ryWtOu454VE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Космос нас ждет', 203, 'G8rBEpD2H3A', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кладбище самолетов', 353, 'd-iiIzio_7g', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ты не такая', 238, 'fB60qciZP_E', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Танцы', 179, '-OlxXpKVv8A', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ебашь, Альбина', 167, 'pMbejQLkQzM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Преждевременное семяизвержение', 290, '-g6GYCazaio', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Сега – винтовар', 153, 'tcoAdT9XitY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Я стараюсь быть лучше', 327, 'OtEIrljyKBc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Взрослые травмы', 169, 'A8-Y_kmwqUU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Улица Сталеваров', 259, '5mEMmCeAVGw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Смирись и расслабься
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Смирись и расслабься', 'https://lh3.googleusercontent.com/zuB4NkTSjEnMZcMWLYKBSmE9dQArzmItBEJbVI2-MHEsnXi118OduPzYCBMQBP2XcVZwmJbHFVtl3OGu=w544-h544-l90-rj', v_artist_id, '', '2012-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 17
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Отель Кооператор', 236, '6NuTIggdXIY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Фанк', 194, 'nsfAj5wDBA0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Песня для девочек', 255, 'hc1Ih1wQDN0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Рустем', 212, '9Q0srCckNQE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('45 лет', 143, '1696AtvSpQM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Всё решено', 192, '_Yq2QAhLt_w', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кайен', 182, '3A2x-7HawDc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Он постоянно', 244, 'uSXblthoxY8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Лишь однажды', 173, 'c-0S68_ZLL8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Наше лето', 209, 'aeGdhFHFj8Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Русский рок', 194, 'pl1RGGyTGls', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Серёжа', 231, 'xeiVCgYxXZA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Первомай', 216, 'eAqJiZzgokE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Я бью женщин и детей', 162, 'M1PrbQAkWP0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gay Porn', 320, 'UgSVLIzemYw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Так гріє', 279, 'H-394VBVvE0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Дешёвые драмы', 228, '1N0BaUxMfXU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Решится Само Собой
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Решится Само Собой', 'https://lh3.googleusercontent.com/lZOE32pgucfFt-VpTI5DlhzvNR6uJfHWsRZ1cV3pSLEqSBaNdv7dJZf8na5vcJchBqJMzg7pl2dbcaXAEg=w544-h544-l90-rj', v_artist_id, '', '2015-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Решится Само Собой', 252, 'OU56-kXpQRY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
-- Артист: Noize MC
INSERT INTO artist (artist_name, description, avatar_url)
VALUES ('Noize MC', 'Noize МС, последний альбом и видеоклип выпущены как Noi3e MC, — российский рэп‑рок‑исполнитель, автор песен и музыкант.

Источник: Wikipedia (', 'https://lh3.googleusercontent.com/EcFC0vIO-VE_Kr1U2RuOkI_1AsVhpp1NrOliCsJ7DklF6EESY7MDErQWfNXXAeDZndRmjtGObbgyghWQ=w2880-h1200-p-l90-rj')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Не все дома
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Не все дома', 'https://lh3.googleusercontent.com/vdJCCX7uXHRPWei3yySJDxwXrmPnMCSrpWNfq6PRBgvF79VL5FQp1WHeyy_phR0nvDUFhVJur3zI-lU=w544-h544-l90-rj', v_artist_id, '', '2025-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 16
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Не все дома', 267, '9imx4Nrg6Zo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Атлантида', 282, 'f2ISdM89EAg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Криокамеры', 217, 'xR8jXM60CpY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Страна Дождей', 141, 'PgYrf8C_C2U', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Причины для радости', 196, '1g2x7RXdsMg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Зубная фея', 227, '9Dl5uW4Jl0M', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Молчанка', 281, '9o3B03nztsY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Соловьи', 201, 'jCgVziPorJY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Самка богомола', 261, 'SWunn_x9m7g', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кооператив «Лебединое озеро»', 191, 'He1fA2oCuRs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Open Air', 251, '4BP92VHkxSA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Никто не пострадал', 153, 'de6QR_37Y5I', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Светлая полоса', 232, 'ZmjbWNiQ6OI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Обломки чувств', 211, 'TB5MfV4Qd6Y', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Светлая полоса (One-Man Band Live)', 328, 'KZagK23gGBM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кооператив «Лебединое озеро» (Remix)', 243, 'TTJO--St8ok', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Выход в город
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Выход в город', 'https://lh3.googleusercontent.com/VypWdmSjpdcLjVGUvJlhPQhq-ch2wula2IaTix7Bm04PgSsA8fSXli9CI459zL752E359zgRArSMEaJV=w544-h544-l90-rj', v_artist_id, '', '2021-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 20
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вояджер-1', 219, 'zX3KFKy9P54', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Выход в город', 189, 'NgUo5YwF5nM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Миокард', 183, 'h-mlrvPYN88', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Сельма Лагерлёф', 275, 'HlcrtzXg93Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Сопротивление воздуха', 177, 'df0o8OcgNI4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Двадцатые годы', 196, 'jxi7aO5S_JI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Век-волкодав (Трибьют Осипу Мандельштаму)', 179, 'FO0Tdu_mISo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Столетняя война', 194, 'rCm1SYrxA0c', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вуду', 210, 'm2GIF5_0RcM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Инстинкт', 234, 'rfiCt9a2dZ4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Трансгуманизм 2.0 (feat. White Punk)', 201, 'lLfqL78ZM6o', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Паучьими тенётами', 179, '0egPm_BUVK4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Песня предателя', 238, '7XUgU6kR884', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Бизнесмен, что продал мир', 233, 'cHaIWi3YhDk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Всё как у людей (Трибьют Егору Летову)', 271, 'HZ3o3tobQc8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Букет крапивы', 186, 'PRqokCzpD90', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('26.04', 231, 'BsOrfosWJ_8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Первые симптомы', 238, 'OUlRd5g_bas', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Опыт отсутствия', 195, 'nuSqvxkz5_U', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вереница (feat. ВЕРЕНИЦА)', 225, 'YxNi4FgebV4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Хипхопера: Орфей & Эвридика
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Хипхопера: Орфей & Эвридика', 'https://lh3.googleusercontent.com/OtS4R3-7iRPeEYDvvk_37aYxqjdOK2KBAzBVtmoFSoC3ELp0qLe3sgHA8sSHN8z95JnpCcY6sDsovAcm=w544-h544-l90-rj', v_artist_id, '', '2018-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 30
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Голос & струны (Орфей)', 183, 'GHHt3LdM3fw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('С нами (Орфей и Эвридика)', 197, 'ihHaguOzq0o', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Фанфары Эллады (Немезида Пафос)', 73, '59FWlr7IAi4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кто тот герой? (Фортуна и Прометей)', 166, 'NkDIJDY9ouo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вряд ли боги соблаговолят нам (Орфей и Эвридика)', 220, 'rI-VEBypiBc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('На вершине - мало места (Харон)', 174, 'K34yBbbTtSg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Спартак vs. Прометей', 101, 'sNvBhbsIPo4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Прометей vs. Спартак', 86, 'ky7snvIeKH4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Нарцисс vs. Орфей', 104, 'y46FsiHJumw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Орфей vs. Нарцисс', 119, 'qnEs1F_bHvQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Прометей vs. Орфей', 138, '6n1o6eTkyE0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Орфей vs. Прометей', 210, 'N2D8WVjg6HY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Подписание контракта (Аид, Фортуна, Орфей)', 160, 'U8ki9wjXpgc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Романс (Орфей и Эвридика)', 224, 'mEa2-IlXsjU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Камера, мотор! (Аид, Фортуна, Орфей)', 238, 'mj75uLlJqQI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('В конце альбома (Эвридика)', 125, 'Ucxp0gjZkro', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Мастер слов и мелодий (Репрезент Орфея)', 170, 'wsB8vQns7kA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Не тот (Эвридика)', 166, 'wmFbLPcQthk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Мелкие черепки (Орфей и Эвридика)', 147, 'TAnL1M9D9ao', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Соблазн (Фортуна и Орфей)', 201, '-FS7eThsjvk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Всё это - лишнее (Эвридика)', 211, 'pzrv-mOGXqU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('1000 москитов (Харон и Орфей)', 54, 'Tz01NDS49IY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Фанфары Эллады', 49, 'caO5mP_DWEc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Разрыв контракта (Аид и Орфей)', 125, '0knCrez_SDg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кто кого? (Харон)', 118, 'qOrRqVjYPZ4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Иллюзии, сны, миражи (Харон и Орфей)', 274, 'NJgs_mBZdLA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Мантра (Орфей)', 264, 'KjumMoodams', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кошмар Эвридики (Эвридика, Морфей, Орфей)', 268, 'zsNauEowywY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Фанфары Эллады (Немезида Пафос)', 112, 'Dqg9xGRdMyQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Без нас (Орфей и Эвридика)', 295, 'lblF3R-fyeA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Чайлдфри
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Чайлдфри', 'https://lh3.googleusercontent.com/_m1UFIKs-7b2pdvH5apuVOb4-qY8rFm3pUNLzrdmcrYBKhqPEi4O2-rjjrp77u4N3Rwl6FCg-Wl85DYL=w544-h544-l90-rj', v_artist_id, '', '2017-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 7
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Чайлдфри (feat. Монеточка)', 245, '_l0LVFRuHMk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Чайлдфри (версия Монеточки)', 185, 'LbJ_Hef2XWg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Марафон', 187, '_xBNb4-5oYE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Чайлдфри (инструментал)', 240, '9MuNZpSdZt4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Чайлдфри (акапелла)', 183, '29mfmcV-_30', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Марафон (инструментал)', 187, 'MKWDk6DmhXE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Марафон (акапелла)', 145, 'r_UZcOksZ0w', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Царь горы
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Царь горы', 'https://lh3.googleusercontent.com/cIrMHLFgR7CFsz27mJEhwNtbVggXJJsg2SxlaVJgiANh_AsP2szppuVf5W8e13crwo5gP7cnEioo8vQo=w544-h544-l90-rj', v_artist_id, '', '2016-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 13
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Make Some Noize', 214, 'yfsfDP0svf8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Стэнли запишет хит', 173, 'dup4M3lDBX8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Царь горы', 270, 'VK1wIdSD4KY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Эмпайр Стэйт', 170, 'OvRGYTkHAI0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Гвозди (feat. Atlantida Project)', 211, 'tCix6HQudc8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Грабли', 278, 'UKxfpyYbh_4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Чайлдфри (feat. Монеточка)', 244, '_l0LVFRuHMk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Всё ОК (feat. Miss Baas)', 243, 'DlFYTJnGe7I', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Джеймесон', 232, 'qu2ysNI6zMk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Питерские крыши', 215, 'djt3_R38MfE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кислотный дождь', 178, 'DuTZ-jZvzwI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Любимый цвет', 152, 'IxoifT1_HKQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('!L!VE!', 158, 'b4g7eWT6wto', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Hard Reboot 3.0
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Hard Reboot 3.0', 'https://lh3.googleusercontent.com/OjNhFGzZqr04i4YyX7AIUHvegjJjJu6IJ1_FnXG8MXpwtLnuA0cXzYRZm1nS1CpK5KCEcc36XOmHPqA=w544-h544-l90-rj', v_artist_id, '', '2015-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 18
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Роботы', 225, 'Seu1XiEF_4k', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ne2Da? (feat. Mewark)', 299, 'wYl4BLmuTeY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Фарыфуры', 222, 'PuqlDNZ7Y1k', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Говорящие головы', 272, 'eUJ25SIprlY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Хозяин леса', 275, 'CygZVEruzjM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Порвав поводок', 211, 'Qky3aFps0LY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Сгораю', 233, 'D6xslyeySKs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Сохрани мою речь (из к/ф «Сохрани мою речь навсегда»)', 251, 'QqtB90hsFLg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Come $ome All (Shoom remix)', 288, '-r8eVuGoR_M', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('абв&эюя (feat. Вера Полозкова)', 265, '5ftlgZVf-Hw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Старые шрамы', 219, '_5fpAolP864', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Снайпер', 270, 'jIgs-DS5kRI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Иордан (feat. Atlantida Project)', 222, 'AOrH0sq-tKY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('220', 219, 'z5GZWDyiymw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('М (feat. Марина Кацуба)', 224, 'Zx5KW5CIYcc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hard Reboot (feat. Astronautalis)', 205, 'FpOMCP-c5uY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hard Reboot (Rusted remix)', 215, 'pKqRa3V7UpY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Порвав поводок (RasKar remix)', 194, 'ZbpOACeFTMI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Hard Reboot
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Hard Reboot', 'https://lh3.googleusercontent.com/5JSS24iOYf4qC0hECZo50ujYiI5tIaTp0gaBoNyts6NqmfvenbCb8BBwkFxQSrvBQAe5wkBcpAKsAUx8=w544-h544-l90-rj', v_artist_id, '', '2014-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 15
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Роботы', 222, 'N_PNOEgS4c0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ne2Da? (feat. Mewark)', 296, 'w-qxj5ced2I', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Фарыфуры', 216, 'q6vw6a7D3tI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Говорящие головы', 264, 'eUJ25SIprlY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Хозяин леса', 271, 'RsffKo-A0Lo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Сгораю', 236, 'HvTZbrzzhp4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Сохрани мою речь (из фильма «Сохрани мою речь навсегда»)', 250, 'QqtB90hsFLg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Come $ome All (Тоталитарный Трэпъ)', 242, '-zsjZhYk0h4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('абв&эюя', 263, 'XO0J1CEKWro', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Старые шрамы', 219, 'FVtPTSTvgrg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Снайпер', 269, 'jIgs-DS5kRI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('220', 217, 'NcYbcmQvFgQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('М', 225, 'oesqhnItJc4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hard Reboot', 204, 'Sq9tPBKHkXY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Safe mode', 142, 'CaB1jbxOVEo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Неразбериха
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Неразбериха', 'https://lh3.googleusercontent.com/VeQmRH0dViy4YJyrQ_Cizl9ksohY2hRBZuOwtb8P1sw1uUu9IDfl84_kzswQ0ATLFPPA-sYrnMtf5eE=w544-h544-l90-rj', v_artist_id, '', '2013-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 18
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('+-0', 176, '2L5m7RJ6gZ4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Нету паспорта (feat. Вахтанг)', 240, 'niyFA24H5_E', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Настоящего', 175, 'wYYHYUpqQT8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Нам не понять (feat. Ёлка)', 205, 'W2orjp93lTc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Жирный чёрный пробел', 195, 'ny7rncsJ9NQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ненавижу', 220, 'a8s6TE0YLWA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Тёмную сторону силы (feat. 7000$)', 133, 'MUAhi_kkCCA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Влиятельные покровители', 223, 'VGXjKfsGPL8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Не надо было', 199, '6I4CUhXN8k0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Чебуречная', 232, 'o4SbFpQ6uyY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Капитан Америка (Не берёт трубу)', 191, 'TUdlt8FTJr8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Нероссия (feat. Влади)', 250, 'EjhYEaflBVA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Неразбериха', 252, 'GDROrxSfxH0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Безмозглая музыка', 262, 'W6VdUO85XNA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Маятник', 157, 'pnW_g2-pCzU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Гой еси', 228, 'ZldhDygE65Y', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Game Over (feat. Staisha)', 174, 'J0FypLe7YPc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Жвачка (Video Edit)', 184, '06ATS_XC_CI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Protivo Gunz
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Protivo Gunz', 'https://lh3.googleusercontent.com/FqiTXbQa_S4iwterkN0ZpL-xJY8SBj0GHlSqmiA8UdF53Vp10335ManuKUuY5Hi8L5rAR-UhCEvRSD2v=w544-h544-l90-rj', v_artist_id, '', '2013-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 17
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Рок - это кал! (feat. Вася Васин)', 242, 'AOXWfVIIPV8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Миша Козырев (feat. Миша Козырев)', 224, 'mZaa_-nbbiI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Накосячу!', 178, 'MHPo9pdoStM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Заговор', 221, 'mAbDZPJ3W2Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Бритни Спирс', 256, 'OQzy2k93_GA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Калифорния', 120, 'gI5xpjU2j7Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('С&Г (Live @ «Б2»)', 192, 'fMgioF3brhY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Общага', 281, 'XGk1ihwVWI4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Честное слово!', 226, 'R0JphSL2DAU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Она уходит', 251, 'h_oR84Akk1Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Жирная корова', 153, 'GMIiHQXrDCA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('У моей девчонки', 173, '5InXjo8TWw0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Под белым флагом', 177, 'LQJo4nSYGPI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Тонкий лёд', 242, 'IR61yrig9bI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Жвачка (feat. Mewark)', 208, 'P_063dM8_zo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Нету паспорта (feat. Давид Казарян)', 235, 'Fpp3NmWD5V4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Хуёвая песня', 451, 'KJHQHJFX2FY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Новый альбом
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Новый альбом', 'https://lh3.googleusercontent.com/CHoG4j9FavAhneHOFLaJ26g6QvHpwuT3rur7DtrizNT9UmTQqB-SJOTVbeZ_C9Bz0PhMjhrZPjX5Mhju=w544-h544-l90-rj', v_artist_id, '', '2012-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 21
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вьетнам', 243, '-2VOyo65N5U', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('ШлакваШаклассика!', 125, 'zX1RPAmZesA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Я глуп', 231, 'rOqFPdOVuW8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Эгоизм', 165, 'GFDuNIFqnjc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вселенная бесконечна?', 261, 'DBnwy46OPFU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Жадина', 138, '-e-cy3bHR00', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Yes Future!', 190, 'rFH4jCW_LzM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Сам (feat. RasKar)', 230, 'ykZ3VNIHRqY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Бассейн', 217, 'uTcu2ZyFT04', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Заебались!', 261, 'Pfz4mxuIvUw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Похуисты (feat. Anacondaz)', 279, 'RxYnRun0P-0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Болт (feat. Ляпис Трубецкой)', 217, '8g1dflMEpOU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Мы хотим танцевать', 285, 'zPxxZzQVc4s', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Школотой', 269, 't4ZQWAysTBY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Танцi (feat. Vоплi Viдоплясова)', 216, 'TY0lDBsXGcc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ток (feat. Вахтанг)', 256, '9a9J7qOCwLk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Брынь-брынь-брынь', 156, '1Mx8lN6LKrQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Друг подруги тёлки брата', 111, 'trjGL-_eKLU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Пушкинский рэп', 151, 'rnm3aieV1wg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Чёрное/Белое (feat. Subatomic Sound System и Nomadic Wax)', 191, 'Mcnd3AlGB5w', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Эдем 14/88', 170, 'w1lmMJaiWR8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Последний альбом
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Последний альбом', 'https://lh3.googleusercontent.com/-aaPxU8FYy6PNi9NX3nit-pquPlLYjGM6_qffvjkSokoAOaeuh6TwrhJtdSUBnN1tbxZhjjeKeyNN7Tw=w544-h544-l90-rj', v_artist_id, '', '2010-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 20
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Тыщатыщ', 212, 'jr41mPpaJQ0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Устрой дестрой (feat. Чача)', 240, 'EFD5cdTkeIg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Испортить вам пати!', 236, 'aKk6F7G62kg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Манки бизнес', 239, 'GVEAiuuzoyE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Бабки в шапку!', 217, 'liMscYcf38E', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Артист (feat. Staisha)', 187, 'vBCxacmy6N0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Певец и актриса (feat. Staisha)', 264, 'gfUNxJ-Ojk0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Бэктон #1', 216, 'xXqQMqPJSkA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ты не считаешь', 275, 'OHCxIuqXfhE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Красный октябрь', 232, '2LyU6Kalhsc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Пустые места', 258, 'AexNbHUHV-8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ругань из-за стены', 225, 'vr1qkx3hfx8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Антенны', 234, '15WZ67vSM0o', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Why The Dollar Falls (feat. Раскарандаш)', 215, 'q9mWlaKg_1k', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Денежный дождь (feat. Comme-il-faut)', 216, 'QkvW_A8si6I', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Гимн понаехавших провинциалов', 212, 'fhSeoS5PUEY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Мизантроп-рэп', 177, 'zCFnt_4IFSU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Жечь электричество!', 269, 'FZLo6ITpdUY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('На Марсе классно', 254, 'glb2Ga4Ww9Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вот и всё. Ну и что?', 178, 'zJtIrY_DY2Y', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Розыгрыш (Из к/ф «Розыгрыш»)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Розыгрыш (Из к/ф «Розыгрыш»)', 'https://lh3.googleusercontent.com/VfTDdyqOsq2TDuY-toFkCI7s8T9UqVtCqsbtD8Pr525OUs-06TpgFLLTckpKEA9YGTuHODg_KERdDwk=w544-h544-l90-rj', v_artist_id, '', '2009-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 15
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Лето в столице', 207, 'rnQbp-wUk9U', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Моё море', 147, 'PdIjcEVlbbM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Утро (Он, а не я)', 132, 'wdLsARRQ5e0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Выдыхай', 189, 'S8VuUv4hwEo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Не могу найти', 205, 'hLqjr1eWy0I', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Это был дождь', 263, 'wV6XDG4fNQE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Суицид', 234, '6PL6taBk2qA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Давай приколемся', 210, 'KCqwCzT_Mt8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Какая жалость', 155, 'YHui7UhdVm8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('У. Е.', 163, 'ZklqpQh8T8o', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Деньги', 155, '_r9ZXcIP7BQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Лето в столице (Maestro A-Sid Jungle RMX)', 254, 'T_VpAK_jRFM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Моё море (ЛевПрав Версия)', 149, 'vpJeJEsi4lk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Выдыхай (ЛевПрав Версия)', 189, '009V0bvaN-k', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Мы всего добились сами (Maestro A-Sid RMX)', 171, 'A9kirPw648Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Voyager-2 (Live at Stadium)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Voyager-2 (Live at Stadium)', 'https://lh3.googleusercontent.com/VWF0CC3xrsLu88V7yP-rfsuuk62Cov-OOGskAEgLUfB2HDP34FbLwliH0rrDRcjtImD6e7DRgMS0VE4R=w544-h544-l90-rj', v_artist_id, '', '2022-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 28
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Миокард (Live at Stadium)', 249, '-XZhvK628vE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Грабли (Live at Stadium)', 293, 'QpTCwTAn31s', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Yes Future! (Live at Stadium)', 280, 'pcPWDqpQkMo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Устрой дестрой (Live at Stadium) (feat. Чача)', 350, 'evuIQeufcd0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Инстинкт (Live at Stadium)', 300, 'hpPU_cEXoTU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Голос & струны (Live at Stadium)', 276, 'N7jGvh22_sY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Столетняя война (Live at Stadium)', 278, 'yWgmhejLyPk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Люди с автоматами (Live at Stadium)', 270, 'Fraq-Hb0r-Y', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Иордан (Live at Stadium)', 264, 'I7Y_v-rf6vU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('На Марсе — классно! (Live at Stadium)', 294, '63-zj9K_mEg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вселенная бесконечна? (Live at Stadium)', 345, 'EQCNOGHIs28', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вояджер-1 (Live at Stadium)', 318, 'rIXE9E1iwtQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('26.04 (Live at Stadium)', 275, 'e3_vCFkUSdM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Выдыхай (Live at Stadium)', 289, 'WwH24LusbjE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Это был дождь (Live at Stadium)', 345, 'wYoTrATQV8I', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Выход в город (Live at Stadium)', 227, 'ERUwa-28OjU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вереница (live at Stadium) (feat. ВЕРЕНИЦА)', 295, 'gDkP3ELbRNo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ругань из-за стены (Live at Stadium)', 271, '1EBWlBfHENo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Букет крапивы (Live at Stadium)', 321, 'wxFBfOCR0bs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Двадцатые годы (Live at Stadium)', 238, 'i4AgrvlnqJY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Почитай старших (Live at Stadium)', 220, '1d7yziT351g', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Make Some Noize (Live at Stadium)', 262, 'Lrp42DqY25o', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Бассейн (Live at Stadium)', 315, 'VRwU8JxPdiY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Паучьими тенётами (Live at Stadium)', 289, 'hWpeUkZ_-bg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Марафон (Live at Stadium)', 187, 'rfJSLCcBhrk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Из окна (Live at Stadium)', 313, 'YGhBqDwyJsg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Всё как у людей (Live at Stadium)', 339, '62J-MkJ-qnQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Сельма Лагерлёф (Live at Stadium)', 292, 'yrJrlG96Mhs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: XV (Live)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('XV (Live)', 'https://lh3.googleusercontent.com/sJSoWa5jTQogQnDVxECePuF-j4BbFrtjawnA4eGNHLYPZuhzw3kQBV4LjFHaAPQT4pRlp8iI0QC1Mgg=w544-h544-l90-rj', v_artist_id, '', '2019-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 27
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('За закрытой дверью (Live)', 259, 'ih74cqCzJr0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Make Some Noize (Live)', 214, 'OJPDD51XyGY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Песня для радио (Live)', 179, 'nha7AaOLq0k', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Люди с автоматами (Live) (feat. Монеточка)', 198, 'u5_J_S0d_Y4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Чайлдфри (Live) (feat. Монеточка)', 241, '72ActzxFmts', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Коррозия Хип-Хопа (Live)', 244, 'pPSVia0yuYQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Бассейн (Live)', 223, '8Rm6pq3ZmhE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вселенная бесконечна? (Live)', 255, '0SuX8so5VF0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Следы на спине (Live)', 183, 'nCEos6TYVYc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('На Марсе классно (Live)', 254, 'NomOgwVUuHE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Иордан (Live)', 249, 'gQr9tLpuANU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Жвачка (Live) (feat. Mewark)', 181, 'vFB5tKDtd1Y', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('В темноте (Live)', 230, 'Vzej1R81Ndg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Yes Future! (Live)', 192, 'm14xLWABgHk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Устрой дестрой! (Live)', 254, '-bP-QNLvR-E', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Танцi (Live) (feat. Иван Дорн)', 216, '9LRUOrbtkqw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Голос и струны (Live)', 225, 'dyvv4sSOuqQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('С нами (Live) (feat. Leila)', 202, 'TvZwOVX2XRk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Без нас (Live) (feat. Leila)', 317, 'crBY6gKgnpw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Питерские крыши (Live)', 228, 'XdewE9TzpeA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Моё море (Live)', 154, '5eIEmlELiBQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Грабли (Live)', 232, 'BNnvd9nELBw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Выдыхай (Live)', 214, 'DC5V3Bh3PVI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ругань из-за стены (Live)', 230, '2rI5E8guXok', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Антенны (Live)', 195, 'yCwBRju7Dw8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Из окна (Live)', 281, 'QJ639UF8GMg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Марафон (Live)', 199, 'aWv-30UsfH4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: *кустик* (Live)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('*кустик* (Live)', 'https://lh3.googleusercontent.com/2-BGxjslaXEkt951D6hkedYxdC0m4B2KRQTlEmVzEwYM9qGuPWMxRAL127felWgbMGvPT20yy2tJ-_fe=w544-h544-l90-rj', v_artist_id, '', '2015-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 16
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Мэйк сам нойз (Live)', 201, '6CHnz3NyunM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Бассейн (Live)', 222, 'zQ-2H2Vis-w', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ругань из-за стены (Live)', 236, '_UY-Ue8_owc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ниже нуля (Live)', 184, 'pa6RJ9xX6GU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вселенная бесконечна? (Live)', 268, '7L0CHy0QVuA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Yes Future! (Live)', 212, 'PgReLnInlZ4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Моё море (Live)', 156, 'OQpgmt_AAlM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Выдыхай (Live)', 216, 'uefJitDBTwY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Певец и актриса (Live)', 286, 'G2hNqBEGPSA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Мерин (Live)', 240, '1wvh_v0dxJM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кури бамбук (Live)', 124, '_QK-6r_EdVY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Песня для радио (Live)', 176, '4hjwEyynov8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Манки бизнес (Live)', 243, 'AecsFY2Ib4w', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Устрой дестрой! (Live)', 268, 'Cf5llVLqiCA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кантемировская (Live)', 277, 'oSCnbAP69Gk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('18:30 (Live)', 190, 'jv6thloj_2c', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Жечь электричество! (Live)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Жечь электричество! (Live)', 'https://lh3.googleusercontent.com/3M_5T958cuyAxN_4JNcsWNFSuqgJzsIRJihrCr3sbT0LwaxzM6fis3myBT_h8bDZVxjInWQDSH2HZxI9=w544-h544-l90-rj', v_artist_id, '', '2011-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 26
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Мы все еще Noize MC! (Фристайл) (Live)', 242, 'B_AEAbBtUq0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Тыщатыщ (Live)', 215, 'cn9nrCvQZio', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('За закрытой дверью (Live)', 225, 'j3ecbXUE_4w', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Антенны (Live)', 208, 'nwM_fyVCS_Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Жизнь без наркотиков (Live)', 268, 'ICeRJPeP2pY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Мизантроп-рэп (Live)', 174, 'n-hnoW9uiWw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Устрой дестрой! (Live)', 234, 'bak9NUbnrpU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Болт (Live)', 217, 'kzcX9XYeb-Y', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Испортить вам пати! (Live)', 283, 'iytUKMPOA7w', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Гимн понаехавших провинциалов (Live)', 224, 'UafiBolbT-c', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Бабки в шапку! (Live)', 215, 'Muduf3PR7LU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Певец & актриса (Live)', 261, 'MN3mZzRthwE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Жечь электричество! (Live)', 321, 'FJbYjvYvxMs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('На Марсе классно (Live)', 258, 'UtuK7eYKKzE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Мы хотим танцевать (Live)', 301, '6ry91XehIXA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Манки бизнес (Live)', 237, 'wRU-G1FCz_I', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ругань из-за стены (Live)', 222, 'x_eW5ver20k', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ты не считаешь (Live)', 267, 'WDubMZFufwY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Выдыхай (Live)', 238, 'KP5oLvDNknM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кури бамбук (Фристайл) (Live)', 319, '6zKx9KsxJx0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вот и всё! Ну и что? (Live)', 192, 'twDluptLp3Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('На работе (платят бабло) (Live)', 149, 'fhv2dG4Dip0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Честное слово! (Live)', 198, 'i5uI8G4HVNA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Артист (Live)', 198, 'jSGZeySqYoM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Аренби! (Live)', 72, 'e9mGPlGJKCU', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Последняя песня (Фристайл) (Live)', 281, '2chvO-w9-gw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Greatest Hits, Vol. 2
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('The Greatest Hits, Vol. 2', 'https://lh3.googleusercontent.com/2uELbuop1DD4FYHuOVjVmwBdvzmj-z-Mh8oMvQgOOKEVy0yoJIFSfl-94_zzeU3rR7lZ8FTY_83ZtSS2=w544-h544-l90-rj', v_artist_id, '', '2010-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 25
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Песня для радио (Полная версия)', 180, 'DVNIVjjESmA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('За закрытой дверью', 215, '99xQhYWX09s', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Из окна', 244, 'uW4UqmQQbRQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Москва - не резиновая', 188, 'ylVXoaec2gk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кантемировская', 201, 'n3_rn3QKc_s', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Жизнь без наркотиков (feat. Маша Макарова)', 222, '4h2b3FYvxvQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('На районе (3 недели нету дудки)', 217, '6m2YJ105Rt8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('На работе (платят бабло)', 150, 'XFJVovX-f60', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кури бамбук!', 152, 'aU7DpbE4VA0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Палево!', 197, 'gitIagUzLtg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Блатняк (feat. Саша "Кислый")', 199, 'Uh0a0BOh0AQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('5П (Песня похуиста про получение пиздюлей)', 186, 'fu5tiubmhCY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Девочка - скинхед', 219, 'jrmWHAAQAgY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Наше движение', 166, '8blsV_0U7SI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Капитал (feat. Ляпис Трубецкой)', 156, 'aIh86w-SjWk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Мы всего добились сами (feat. Адик 22ВО7)', 184, '81YF266wfCg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Аренби!!!', 84, '3b0lMbXKIeI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Поднимите руки', 146, 'pLapjrHey5g', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ctrl+Alt+Del (feat. Staisha)', 169, '6NA1paZKbPw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ниже нуля', 160, 'yvBcgcIPWGM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Моё море', 154, 'PdIjcEVlbbM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Выдыхай', 193, 'S8VuUv4hwEo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Из окна (Ска-панк-версия)', 263, 'ShuIwhGWKbw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кто убил?', 267, 'pBEGDPkRFWI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('3П (Правдивая песня пиздабола)', 189, 'PwbAgf6blJw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Greatest Hits, Vol. 1
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('The Greatest Hits, Vol. 1', 'https://lh3.googleusercontent.com/BRL1AZgafVqCs0zFm6yiiLt7sCLIuo3VL4owkV6NBEvWLKml_pLGvIoSXGWkELdy3W7sLYabwzp4l5nt=w544-h544-l90-rj', v_artist_id, '', '2008-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 20
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Песня для радио', 174, 'DVNIVjjESmA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('За закрытой дверью', 215, '99xQhYWX09s', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Из окна', 244, 'uW4UqmQQbRQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Москва - не резиновая', 187, 'QoBkdd9UqC4', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кантемировская', 201, 'n3_rn3QKc_s', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Жизнь без наркотиков (feat. Маша Макарова)', 222, 'am8Oy4OOOqo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('На районе (3 недели нету плана)', 216, '8putR2hHreM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('На работе (платят бабло)', 150, 'FUyNYUu3jhE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кури бамбук!', 152, '9BgIPvD3rp8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Палево!', 197, 'gitIagUzLtg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Блатняк', 199, 'mqKxVFMCVLc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('5П (Песня похуиста про получение пиздюлей)', 186, 'ulvXquSZ48E', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Девочка - скинхед', 219, 'RagOBBdf8Zc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Наше движение', 166, 'Uf4zlJOt95k', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Аренби!!!', 84, 'wvZ8rNGnBDY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Поднимите руки', 146, 'Iq37ngvZciw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Моё море', 154, 'PdIjcEVlbbM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Выдыхай', 192, 'S8VuUv4hwEo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('3П (Правдивая песня пиздабола)', 188, 'yp_nyD23NSc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Мы всего добились сами! (feat. Адик 22ВО7)', 182, 'UEZ5WwvfWXE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Планета Земля
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Планета Земля', 'https://lh3.googleusercontent.com/5tPGBdsN8W1-alBrQPzUZdMbijNosnByWtwAx9n1YenZCm87tXGjMoa4TjAWpkkG5KVb38xM6Bzy_5l7=w544-h544-l90-rj', v_artist_id, '', '2025-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Планета Земля', 123, '1kAc1lZl3sI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Обломки чувств
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Обломки чувств', 'https://lh3.googleusercontent.com/tUPc3Myo9nVb6d2GYKCsZ_apDnqtI-1rUfzolJ4HTyQzRGZ5Hz5hueOiYBKlJvZXECJY1luieeS-Ttgc=w544-h544-l90-rj', v_artist_id, '', '2025-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Обломки чувств', 211, 'yEpS8egTymA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Светлая полоса
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Светлая полоса', 'https://lh3.googleusercontent.com/TukvVCflt8Xo1G4NNZHkUscIjW2MLui-npq9jonMqM9RVTbstxyaICQwUdW7oWUZTLV6_2t0cFijVos=w544-h544-l90-rj', v_artist_id, '', '2024-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Светлая полоса', 232, 'KJDcZmVg1fA', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Kalinka
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Kalinka', 'https://lh3.googleusercontent.com/HK2qhrF9Pq0NURuA5bYaqwhVKBO-nf6NeMjkm_a_DIDDWrjEMMdSpMDDai-qy-Qad6lGlykCqlairjOA=w544-h544-l90-rj', v_artist_id, '', '2024-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Kalinka', 172, 'rW8IMFn0vDM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Open Air
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Open Air', 'https://lh3.googleusercontent.com/7veJ_I9TD24aAEagrq1UPN2leUw5U4OvljI2nNY-3LGDCTQla6XWXhVVYLmNIP69mL8qWyHBVt7UOtA_=w544-h544-l90-rj', v_artist_id, '', '2024-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Open Air', 249, 'CF8YLzEnLqs', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Грабли (Live Looping Version)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Грабли (Live Looping Version)', 'https://lh3.googleusercontent.com/nac05ZOhWJr0feKnK9VhF8u6KKiGiXdJ9jpyiOXfBMN0PgMkf56Z5UepJwT-f6KNYOi-GJJSa2MT144j=w544-h544-l90-rj', v_artist_id, '', '2023-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Грабли (Live Looping Version)', 358, '5poPueFBC3g', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Кооператив «Лебединое озеро»
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Кооператив «Лебединое озеро»', 'https://lh3.googleusercontent.com/GNVJdyOtM0JlM9ySYh8zSlkHMKUNGgyD0eX79nc8vqtYFcqJmCrYFaXalYm7_yDTPBoWBNQmhyswsUFk-Q=w544-h544-l90-rj', v_artist_id, '', '2023-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кооператив «Лебединое озеро»', 189, 'qU1XzMNnjvY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Век-волкодав (За гремучую доблесть грядущих веков)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Век-волкодав (За гремучую доблесть грядущих веков)', 'https://lh3.googleusercontent.com/Sonc_3T3UGkHYBc6uiql8g2vjeyyf7H50ek6l4LJWV16CZDUh6YGgMP5w7qHKLOYgcmohc49DLHuGso=w544-h544-l90-rj', v_artist_id, '', '2021-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Век-волкодав (За гремучую доблесть грядущих веков)', 180, 'FO0Tdu_mISo', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Вояджер-1
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Вояджер-1', 'https://lh3.googleusercontent.com/Nw2gnXjaHrburCny4GpMGoY4lCmvfQK57O4WjmeT0x5EGScF3KpU3hj8Xeu9U-vjvwJTaADqzUg3jwY=w544-h544-l90-rj', v_artist_id, '', '2020-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 4
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вояджер-1', 215, 'zX3KFKy9P54', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вояджер-1 (feat. Белгородский академический русский оркестр)', 225, '3o63YASPeB0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вояджер-1 (LAB с Антоном Беляевым)', 239, 'RYwbqzBBPB0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Вояджер-1 (Undergroover Rework)', 298, 'y2BxT6zJwcw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Живи без остатка (feat. Монеточка)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Живи без остатка (feat. Монеточка)', 'https://lh3.googleusercontent.com/1Z8rrXuQXzDzmDBy_FGsQ-x1wJujILxbgzMKCU1sBz0yQsMombvnkDegpcImy8w0D3zIkFg3ScTBbB0W=w544-h544-l90-rj', v_artist_id, '', '2020-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Живи без остатка (feat. Монеточка)', 213, 'qrwfBxlGwSM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Лига легенд
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Лига легенд', 'https://lh3.googleusercontent.com/OOYwZr09XV9WGpA2RR4mHC55Kpz_7BlZMu8QcqF0ot0dZXdRHOOvMZ9yj9mEmsAT5Da4h2-Guki3N1G4_Q=w544-h544-l90-rj', v_artist_id, '', '2020-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Лига легенд', 219, 'La8FkMzk6P0', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Катацумури (feat. Linda)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Катацумури (feat. Linda)', 'https://lh3.googleusercontent.com/1oPXscn_cCKlumKY3K638Tcu1u8gmE6pYA8Lv9CzruC9v1DyFe8VqLt1NtCTAlLeto-pe4NXUFmhyizIRA=w544-h544-l90-rj', v_artist_id, '', '2020-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Катацумури (feat. Linda)', 247, 'S13tuGslY2M', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Давай сбежим (feat. Damilola Karpow)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Давай сбежим (feat. Damilola Karpow)', 'https://lh3.googleusercontent.com/mAPmnXFNv0AyqYq7vdTf0ytGfJ18rc25fCg5twmB7FMWFEzWOm0bn7txIJESnmCQR8K0pMozH6YzSom2=w544-h544-l90-rj', v_artist_id, '', '2020-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Давай сбежим (feat. Damilola Karpow)', 232, 'UqCJtDt4Qsk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: 26.04
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('26.04', 'https://lh3.googleusercontent.com/fqTgbdY-GLhYzYJ8dl8l1Lvl4sqSgucep9PufqeOqEHz4y1wjXZi7xByEJCeAaSLU_E_pcTk1wCWCHfkPQ=w544-h544-l90-rj', v_artist_id, '', '2020-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('26.04', 230, 'CVY5YxluoIw', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Последний министр (из "Последний министр")
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Последний министр (из "Последний министр")', 'https://lh3.googleusercontent.com/yJovbvYXjhrS33wA-zWK_k3HYg1u58zDkRqRozdBwCHb5z0ijSvCOoRzEk2Lzf5ExhGER1kP39YdQKY=w544-h544-l90-rj', v_artist_id, '', '2020-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Последний министр (из "Последний министр")', 232, 'hwPFznw-Z4g', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Почитай старших
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Почитай старших', 'https://lh3.googleusercontent.com/KKAuHaC9bWq8WKjiw-Niql5V7winwT8IrqlhoUdwVYE1TBXyuixrfQ6OUmzf3jazhVa-SzFONZoxPzPJ=w544-h544-l90-rj', v_artist_id, '', '2019-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Почитай старших', 202, 'iA8nxSf-HAc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Chasing the Horizon (feat. Sonny Sandoval)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Chasing the Horizon (feat. Sonny Sandoval)', 'https://lh3.googleusercontent.com/avgATeyWc08RnLvIvldBsTqb3yjzd_GwyIZy06ykYK8UoCAGJhSBd0U5rGJVI2jbZjSE30SGa9a40psaLQ=w544-h544-l90-rj', v_artist_id, '', '2019-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Chasing the Horizon (feat. Sonny Sandoval)', 199, 'yNC0p2RXeXM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: 200+
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('200+', 'https://lh3.googleusercontent.com/DTD6U0e3cEQJBoEPlzftHZZ0sioHAYbsvkerMVBJDzJophwexzA9KBKrqQ5IRQf6zGOqYTR2dcbfWE0=w544-h544-l90-rj', v_artist_id, '', '2018-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('200+', 147, 'zsZdFHETJGg', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Следы на спине
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Следы на спине', 'https://lh3.googleusercontent.com/aEctcK-hwtfzcT-owJybsgrON-xTqMF0WlG-4SblGFkaaS8_Jj4VSH2Ka_VE0GD0iJm8oiFwsK4H1ZbWBA=w544-h544-l90-rj', v_artist_id, '', '2018-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 3
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Следы на спине', 182, 'KqGNA0n6DrI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Следы на спине (Video Edit)', 178, '3-SXMJKVXPM', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Следы на спине (Tsucore-Neopop Remix)', 184, 'ehW4Rr0VD9Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Люди с автоматами
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Люди с автоматами', 'https://lh3.googleusercontent.com/nr_0DQMGXEjWhyw4_5VvN0VBSkkU-h6_2e_UfJZWMqyoz9EDMuRVhP-K_L0da0Z9U7xVdzHR-pWIme9A=w544-h544-l90-rj', v_artist_id, '', '2018-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Люди с автоматами', 193, 'l7dogXP9Lmc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: No Comments
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('No Comments', 'https://lh3.googleusercontent.com/au0EmoJ-dlqhSSn6STTKYppMDaX9o2prcUoFCumAUZ0zWpn4oT49vICsFejEYDJe8cEtA5nddFq911I=w544-h544-l90-rj', v_artist_id, '', '2017-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 4
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('В темноте', 207, 'CAZWpTtMHsk', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Детка, послушай', 219, 'LTjJ43ZAq1s', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Зубы мудрости', 189, 'YUvMP9DYiuc', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('В темноте (Brodsky Version)', 207, 'oMseWzx4qHY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Коррозия хип-хопа
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Коррозия хип-хопа', 'https://lh3.googleusercontent.com/RC5W8jWbmEab7582aIrf654-Er_qxlFu6Zyd9SOTxY96L1QtywEYj3JRTUalznZB4ln4cyl7jyL9a8gi=w544-h544-l90-rj', v_artist_id, '', '2017-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 4
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Коррозия хип-хопа', 240, 'rLxoaeI8n0g', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Коррозия хип-хопа (Инструментал)', 240, '1xJ-kYXGb5I', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Коррозия хип-хопа (Акапелла)', 226, 'C3vx88YT7KQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Коррозия хип-хопа (Цензурная версия)', 240, 'rLxoaeI8n0g', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Лето в столице (feat. SunSay и MC FAME)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Лето в столице (feat. SunSay и MC FAME)', 'https://lh3.googleusercontent.com/39hS7WvNxxUu3VKmRyyd37x8885cbP2Ac8uUaRy6Fo1nrF6y8eLqIn4qiB6B7W16ir9eR-1FwrEcItg=w544-h544-l90-rj', v_artist_id, '', '2017-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Лето в столице (feat. SunSay и MC FAME)', 204, 'ah5yBKSrYCQ', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Make Some Noize
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Make Some Noize', 'https://lh3.googleusercontent.com/whHXdGV4UrjYCeupNx6qPNub4a2yLE3kMUUzgxe1fLM5n206haG6rx-2aJGNQTjEK6bTVo7HyTMy_PnuNQ=w544-h544-l90-rj', v_artist_id, '', '2016-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 5
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Make Some Noize (Single Edit)', 214, 'yfsfDP0svf8', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Make Some Noize (Acoustic Version)', 203, 's6gTym3sbww', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Make Some Noize (Live in Minsk)', 251, '3guUFQtLb38', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Make Some Noize (Instrumental)', 214, 'AFLUWiNk1mE', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Make Some Noize (Acapella)', 195, '3dwRn3yb5UI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Lenin Has Risen
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Lenin Has Risen', 'https://lh3.googleusercontent.com/pX2rpT3w0X-fO0fSKFu3PIbwMjecsiBEhhPduyhLK1mqaYIZ7sMW7P4KW6Gckt3QzwItgimEpAGnhm4s=w544-h544-l90-rj', v_artist_id, '', '2016-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Lenin Has Risen', 262, 'HBQR-oI5PjI', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Jingle Bellz
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Jingle Bellz', 'https://lh3.googleusercontent.com/0uV49VXcdh_9KcK7BOqWW5Qcsc4KGZkP07MNeRV4bKAvbszIcHBZ2p1RnjMe7jXNAXmF-g5Xin8nReAp=w544-h544-l90-rj', v_artist_id, '', '2015-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Jingle Bellz', 234, 'GYGUYGcJJ1Q', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: 12 Обезьян
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('12 Обезьян', 'https://lh3.googleusercontent.com/WFQ1YIqazP4lF7d4fk_t5SsXlp9YCxmE58PZTp1aCHOM3ihNbX_SVoiW-hqiZ-41N5d5529dBWfKJ0VS=w544-h544-l90-rj', v_artist_id, '', '2014-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('12 Обезьян', 164, 'q3saTXDaj9o', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Voyager 1 (Live in New York)
INSERT INTO album (title, avatar_url, artist_id, description, release_date)
VALUES ('Voyager 1 (Live in New York)', 'https://lh3.googleusercontent.com/KIdO9XiXeCE4I5IqZfg9cnTXJb78lsK_-ngEeENGEAfbqYz5mSfxCBftQbhBKx_qLPtpKxvDu31M5vr-=w544-h544-l90-rj', v_artist_id, '', '2024-01-01')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Voyager 1 (Live in New York)', 257, 'cEm0o5jfftY', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
END $$;
