-- SQL миграция для вставки данных из YouTube Music
-- Генерация: 2025-11-03T22:30:09.914345

DO $$
DECLARE
    v_artist_id UUID;
    v_album_id UUID;
    v_track_id UUID;
BEGIN
-- Артист: Валентин Стрыкало
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Валентин Стрыкало', '«Валенти́н Стры́кало» — украинская рок-группа, основанная в 2010 году солистом Юрием Капланом, который получил известность после записи серии видеообращений к звёздам шоу-бизнеса от имени наивного провинциального парня «Валентина Стрыкало из села Бурильцево».
Каплан основал группу под влиянием музыки таких рок-групп, как «Сплин» и Radiohead. В 2012 году был выпущен дебютный альбом «Смирись и расслабься!», который по большей части был написан в жанре камеди-рок, пусть в нём и имелись песни с серьёзным характером. В следующем году вышел альбом «Часть чего-то большего», в котором упор был сделан больше на лирику, но также не обошлось и без завуалированного юмора в некоторых песнях. В 2016 году выходит альбом «Развлечение», в котором группа окончательно отошла от юмора и ушла к теме депрессии. Также в записи заметно, что группа вдохновлялась британской рок-группой Pink Floyd.
Летом 2018 года группа дала последние концерты и ушла в затишье, а в мае 2019 года Юрий Каплан сообщил, что группа распущена.

Источник: Wikipedia (', 'http://localhost:8099/avatars/artists/UCabtR67_U5O72yRXF7hiI-g_avatar.webp', 'http://localhost:8099/avatars/artists/UCabtR67_U5O72yRXF7hiI-g_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Развлечение
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Развлечение', 'http://localhost:8099/avatars/albums/MPREb_47cJSJKG9bx.webp', v_artist_id, '«Развлечение» — третий и последний студийный альбом украинской группы «Валентин Стрыкало», релиз которого состоялся 13 октября 2016 года.

Источник: Wikipedia (', '2016-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 8
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ускользает', 306, 'http://localhost:8099/music/tracks/PZZxcKYskpg.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('О брат', 295, 'http://localhost:8099/music/tracks/8JAfQuzbOBM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('92', 230, 'http://localhost:8099/music/tracks/e8mLnAVQ9mY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Решится само собой', 253, 'http://localhost:8099/music/tracks/kOfBszWa0t4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Делать это трезвым', 194, 'http://localhost:8099/music/tracks/nAKX-7_fbw8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Тени', 345, 'http://localhost:8099/music/tracks/FiynOYlSfrk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Бесполезно', 425, 'http://localhost:8099/music/tracks/KdqzgAKoHTU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Подворотня - мой дом', 202, 'http://localhost:8099/music/tracks/dFuIywTOJxc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Часть чего-то большего
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Часть чего-то большего', 'http://localhost:8099/avatars/albums/MPREb_UVzSEcp2hnu.webp', v_artist_id, '«Часть чего-то большего» — второй студийный альбом украинской группы «Валентин Стрыкало», вышедший 20 октября 2013 года.
Премьера «Часть чего-то большего» состоялась 20 октября 2013 года на трёх крупнейших российских электронных площадках: iTunes, «Яндекс. Музыка» и Trava.ru. В поддержку альбома был запущен одноимённый концертный тур по городам России, в том числе презентации 20 октября в Санкт-Петербурге и 27 октября в Москве.
В поддержку альбома вышло два клипа: «Знаешь, Таня» и «Космос нас ждет».
Юрий Каплан об альбоме:
Как и на первой пластинке, юмора будет предостаточно,
разница лишь в том, что он будет более завуалированным. Ведь каждый любит вуаль.
— Официальный сайт группы Валентин Стрыкало

Источник: Wikipedia (', '2013-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 14
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Самый лучший друг', 166, 'http://localhost:8099/music/tracks/a_NbJCJoUM4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Офисный стиляга', 134, 'http://localhost:8099/music/tracks/CmHO2FKYGTs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Все мои друзья', 161, 'http://localhost:8099/music/tracks/C7_oi97EzbM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Знаешь, Таня', 274, 'http://localhost:8099/music/tracks/ryWtOu454VE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Космос нас ждет', 203, 'http://localhost:8099/music/tracks/G8rBEpD2H3A.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кладбище самолетов', 353, 'http://localhost:8099/music/tracks/d-iiIzio_7g.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ты не такая', 238, 'http://localhost:8099/music/tracks/fB60qciZP_E.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Танцы', 179, 'http://localhost:8099/music/tracks/-OlxXpKVv8A.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ебашь, Альбина', 167, 'http://localhost:8099/music/tracks/pMbejQLkQzM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Преждевременное семяизвержение', 290, 'http://localhost:8099/music/tracks/-g6GYCazaio.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Сега – винтовар', 153, 'http://localhost:8099/music/tracks/tcoAdT9XitY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Я стараюсь быть лучше', 327, 'http://localhost:8099/music/tracks/OtEIrljyKBc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Взрослые травмы', 169, 'http://localhost:8099/music/tracks/A8-Y_kmwqUU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Улица Сталеваров', 259, 'http://localhost:8099/music/tracks/5mEMmCeAVGw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Смирись и расслабься
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Смирись и расслабься', 'http://localhost:8099/avatars/albums/MPREb_8WzKCTl3P8x.webp', v_artist_id, '«Смирись и расслабься!» — первый студийный альбом украинской группы «Валентин Стрыкало».
За три дня до официального релиза группа представила альбом на сайте Яндекс.Музыка для бесплатного прослушивания. На песни «Кайен», «Русский рок» и «Наше лето» были сняты клипы.
Юрий Каплан об альбоме:
В альбоме будет всем давно знакомый материал, который мы два года исполняем на концертах. Но его не было в качественных записях и люди не могли закинуть его себе в плеер. Мы сменили несколько студий прежде, чем остановились на удовлетворяющем нас варианте. Бывало, что материал накапливали, сводили и мастерили, но всё это переписывали опять. Фишка в том, что люди привыкли слышать эти песни на концертах, где всё звучит круто и драйвово. Поэтому в записи они ожидают того же драйва, что и на живых выступлениях. Но на концерте мы заряжены энергией зала и наша отдача гораздо круче. Очень сложно вывести себя на этот уровень в студии и добиться того же рок-н-ролла. Потому мы требовательно относились к выбору студии звукозаписи и звукорежиссёра.
— Официальный сайт Валентина Стрыкало

Источник: Wikipedia (', '2012-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 17
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Отель Кооператор', 236, 'http://localhost:8099/music/tracks/6NuTIggdXIY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Фанк', 194, 'http://localhost:8099/music/tracks/nsfAj5wDBA0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Песня для девочек', 255, 'http://localhost:8099/music/tracks/hc1Ih1wQDN0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Рустем', 212, 'http://localhost:8099/music/tracks/9Q0srCckNQE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('45 лет', 143, 'http://localhost:8099/music/tracks/1696AtvSpQM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Всё решено', 192, 'http://localhost:8099/music/tracks/_Yq2QAhLt_w.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кайен', 182, 'http://localhost:8099/music/tracks/3A2x-7HawDc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Он постоянно', 244, 'http://localhost:8099/music/tracks/uSXblthoxY8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Лишь однажды', 173, 'http://localhost:8099/music/tracks/c-0S68_ZLL8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Наше лето', 209, 'http://localhost:8099/music/tracks/aeGdhFHFj8Q.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Русский рок', 194, 'http://localhost:8099/music/tracks/pl1RGGyTGls.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Серёжа', 231, 'http://localhost:8099/music/tracks/xeiVCgYxXZA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Первомай', 216, 'http://localhost:8099/music/tracks/eAqJiZzgokE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Я бью женщин и детей', 162, 'http://localhost:8099/music/tracks/M1PrbQAkWP0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gay Porn', 320, 'http://localhost:8099/music/tracks/UgSVLIzemYw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Так гріє', 279, 'http://localhost:8099/music/tracks/H-394VBVvE0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Дешёвые драмы', 228, 'http://localhost:8099/music/tracks/1N0BaUxMfXU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Решится Само Собой
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Решится Само Собой', 'http://localhost:8099/avatars/albums/MPREb_Vnj7dnuAV3p.webp', v_artist_id, '', '2013-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Решится Само Собой', 252, 'http://localhost:8099/music/tracks/OU56-kXpQRY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
END $$;