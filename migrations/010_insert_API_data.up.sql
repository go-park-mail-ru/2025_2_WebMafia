-- SQL миграция для вставки данных из YouTube Music
-- Генерация: 2025-11-04T21:25:46.565268

DO $$
DECLARE
    v_artist_id UUID;
    v_album_id UUID;
    v_track_id UUID;
BEGIN
-- Артист: Kanye West
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Kanye West', '', 'UCs6eXM7s8Vl5WcECcRHc2qQ_avatar.webp', 'UCs6eXM7s8Vl5WcECcRHc2qQ_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Graduation
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Graduation', 'MPREb_KCNeTnK02S7.webp', v_artist_id, 'Graduation — третий студийный альбом американского исполнителя хип-хопа Канье Уэста, вышел 11 сентября 2007 года на лейбле Roc-A-Fella Records. Запись длилась около двух лет и проходила в разных студиях в США. Обложку для альбома оформил японский иллюстратор Такаси Мураками. Альбом дебютировал под первым номером в US Billboard 200. Продажи в первую неделю составили 957 000 копий. 50 Cent выпустил свой альбом «Curtis» в этот же день. Такое событие вызвало большой ажиотаж среди музыкального общества. Конкуренция привела к хорошим продажам обоих альбомов. «Graduation» получил в целом положительные отзывы от большинства музыкальных критиков и получил «Грэмми» в номинации «Лучший рэп-альбом», в том числе и первую награду от премии American Music Awards 2008 года в номинации «Лучший рэп/хип-хоп альбом». Альбом был продан в размере 2 166 000 копий в США и был сертифицирован дважды платиновым.

Источник: Wikipedia (', '2007-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 14
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Good Morning', 196, '6CHs4x2uqcQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Champion', 168, 'jKT4ArZCkso.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Stronger', 312, 'PsO6ZnUZI0g.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I Wonder', 244, 'MxEjnYdfLXU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Good Life (feat. T-Pain)', 207, 'FEKEjpTzB0Q.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Can''t Tell Me Nothing', 272, 'hqvcww4ydh8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Barry Bonds (feat. Lil Wayne)', 205, 'LYa5q2AjWNQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Drunk and Hot Girls (feat. Mos Def)', 314, '9uOg6LyIAUo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Flashing Lights (feat. Dwele)', 238, 'ila-hAUXR5U.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Everything I Am (feat. DJ Premier)', 228, 'ZtkNfC5Oymw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Glory', 213, 'e-IAGmTuUmw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Homecoming (feat. Chris Martin)', 204, 'LQ488QrqGE4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Big Brother', 288, 'HInIGGXhJHs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Good Night (feat. Mos Def и Al Be Back)', 186, '_J-OaFVdTKE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: My Beautiful Dark Twisted Fantasy
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('My Beautiful Dark Twisted Fantasy', 'MPREb_TeXtApchIhq.webp', v_artist_id, 'My Beautiful Dark Twisted Fantasy — пятый студийный альбом американского хип-хоп-артиста Канье Уэста, вышедший 22 ноября 2010 года. Запись альбома проходила на острове Даймонд-Хед с 2008 по 2010 год, при участии нескольких продюсеров, среди которых сам Уэст, Джефф Баскер, RZA, Кен Льюис, No I.D., Майк Дин и другие.
Четыре песни из альбома были выпущены на синглах: «Power», «Runaway», «Monster» и «All of the Lights». Все четыре сингла попали в различные чарты, некоторые из них заняли там верхние строчки.
Сразу после выхода альбом получил положительные отзывы критиков, которые отмечали в песнях разнообразие стилей, а также лирические темы, затронутые Уэстом.
Альбом возглавлял чарты Billboard в шести категориях одновременно.

Источник: Wikipedia (', '2010-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 13
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dark Fantasy', 281, 'UTH1VNHLjng.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gorgeous (feat. Kid Cudi и Raekwon)', 358, 'miJAfs7jhak.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('POWER', 293, 'L53gjP-TtGE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('All Of The Lights (Interlude)', 63, 'WHxRd_va950.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('All Of The Lights', 300, 'HAfFfqiYLp0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Monster (feat. JAY-Z, Rick Ross, Nicki Minaj и Bon Iver)', 379, 'pS6HRKZQLFA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('So Appalled (feat. Kanye West, JAY-Z, Pusha T и Prynce Cy Hi)', 398, '0o9HzQ3zAcE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Devil In A New Dress (feat. Rick Ross)', 352, 'sk3rpYkiHe8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Runaway (feat. Pusha T)', 548, 'Bm5iA4Zupek.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hell Of A Life', 328, 'tJKNcI6jC6A.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Blame Game (feat. John Legend)', 470, '6mp72xUirfs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Lost In The World (feat. Bon Iver)', 257, 'ofaRvNOV4SI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Who Will Survive In America', 99, 'UB6sXiZ1ldw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: 808s & Heartbreak
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('808s & Heartbreak', 'MPREb_jvd4HPJ3O2X.webp', v_artist_id, '808s & Heartbreak — четвёртый студийный альбом американского хип-хоп артиста Канье Уэста, изданный 24 ноября 2008 года под лейблом Roc-A-Fella Records. Уэст проводил запись между сентябрем и октябрем 2008 года в студии Glenwood, в Бербанк и студии звукозаписи Аvex в Гонолулу. В продюсировании альбома принимали участие такие продюсеры, как No I.D., Jeff Bhasker и другие. В записи песен также участвовали Кид Кади, Young Jeezy, Mr Hudson и Lil Wayne.
Сам Канье Уэст квалифицирует жанр альбома как поп, но 808s & Heartbreak содержит также элементы жанров синти-поп, R&B и электропоп, а в большинстве треков альбома Канье поёт, в отличие от предыдущих альбомов. В песнях затрагиваются лирические темы, такие как любовь и одиночество. Также в альбоме широко используется эффект Auto-Tune и драм-машина Roland TR-808, что делает звук электронным. Альбом выполнен в стиле минимализма, а его звучание отличается от привычного звучания хип-хопа.
Альбом дебютировал под номером один в американском чарте Billboard 200, с 450 145 проданных копий в первую неделю. Также все 4 сингла попали в Billboard 200.

Источник: Wikipedia (', '2008-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 12
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Say You Will', 378, 'd9BMPmfxaoM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Welcome To Heartbreak (feat. Kid Cudi)', 263, 'wMH0e8kIZtE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Heartless', 211, 'Co0tTeuUVhU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Amazing (feat. Young Jeezy)', 239, 'PH4JPgVD2SM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Love Lockdown', 271, 'HZwMX6T5Jhk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Paranoid (feat. Mr Hudson)', 278, 'CiY8-LYkCEk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('RoboCop', 275, 'kVl__NgDAdw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Street Lights', 190, 'TUfuDKKGQxU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Bad News', 239, '1BlH1JZBeXI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('See You In My Nightmares (feat. Lil Wayne)', 259, 'T5e-nhk4HTQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Coldest Winter', 165, 'tpT7H7qIHIo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Pinocchio Story', 362, 'OeCdG0Mzrkw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Yeezus
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Yeezus', 'MPREb_uR76QxVTBOB.webp', v_artist_id, 'Yeezus — шестой студийный альбом американского рэпера и продюсера Канье Уэста. Он был выпущен 18 июня 2013 года на лейблах Def Jam Recordings и Roc-A-Fella Records. Для записи пластинки Уэст привлёк ряд исполнителей, включая Майка Дина, Daft Punk, Ноя Голдштейна, Арку, Hudson Mohawke и Трэвиса Скотта. Альбом также содержит гостевые участия от Джастина Вернона, Chief Keef, Kid Cudi, Assassin, King L, Чарли Уилсона и Фрэнка Оушена.
За пятнадцать дней до выхода Уэст заручился помощью продюсера Рика Рубина, чтобы сделать звучание Yeezus более минималистичным. Пластинка вдохновлена множеством жанров, включая индастриал, эйсид-хаус, электро, панк и чикагский дрилл. На треках звучат искажённые сэмплы, например, отрывок из «Strange Fruit» Нины Симон в «Blood on the Leaves».
Yeezus получил широкое признание критиков, многие из которых назвали пластинку одной из лучших работ Уэста и высоко оценили её «дерзость». Альбом был номинирован на премию «Грэмми» в категории «Лучший рэп-альбом» в 2014 году. Проект дебютировал под номером один в американском чарте Billboard 200, было продано 327 000 копий за первую неделю.

Источник: Wikipedia (', '2013-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 10
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('On Sight', 157, 'uU9Fe-WXew4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Black Skinhead', 189, 'q604eed4ad0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I Am a God', 232, 'KuQoQgL63Xo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('New Slaves', 257, 'vQ0u09mFodw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hold My Liquor', 327, 'bvBfiRWLj_0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I''m in It', 235, '_jZuz3NEr18.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Blood on the Leaves', 360, 'KEA0btSNkpw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Guilt Trip', 244, '5hthMeEqf40.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Send It Up', 179, 'vUFiVwa6U_c.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Bound 2', 230, 'BBAtAM7vtgc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Late Registration
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Late Registration', 'MPREb_1RQBOijcxbh.webp', v_artist_id, 'Late Registration — второй студийный альбом американского рэпера Канье Уэста, выпущенный 30 августа 2005 года. Запись альбома проводилась в течение года на различных студиях в Нью-Йорке и Голливуде. Во время записи Уэст сотрудничал с продюсером Джоном Брайоном. В записи альбома приняли участие Jay-Z, Lupe Fiasco, Jamie Foxx, Nas, Brandy и Адам Левин из группы Maroon 5.
Альбом дебютировал на первой строке Billboard 200, продавшись в первую неделю количеством 860 000 копий. Всего в США было продано 3 миллиона копий. С альбома было выпущено 5 синглов, которые попали в чарты. После выпуска альбом Late Registration получил признание большинства музыкальных критиков и позволил получить Уэсту несколько наград, среди которых премия «Грэмми» за «Лучший рэп-альбом». Журнал Rolling Stone, назвавший Late Registration лучшим альбомом 2005 года, расположил альбом на 40-м месте в списке лучших альбомов 2000-х.

Источник: Wikipedia (', '2005-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 21
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Wake Up Mr. West', 42, 'Bwyu-SZ7g_E.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Heard ''Em Say (feat. Adam Levine)', 204, 'elVF7oG0pQs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Touch The Sky (feat. Lupe Fiasco)', 237, 'YkwQbuAGLj4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gold Digger (feat. Jamie Foxx)', 208, '6vwNcNOTVzY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Skit #1', 34, 'G4qTNRbAp-c.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Drive Slow (feat. Kanye West и Paul Wall)', 273, 'Q1ViJEYNki4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('My Way Home (feat. Common)', 104, 'TgAomHGqKUM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Crack Music (feat. The Game)', 271, '2tmPSK-w90o.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Roses', 246, 'Qxlnb1lEdEs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Bring Me Down (feat. Brandy)', 199, 'CZ_-O31R3p4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Addiction', 268, 'YuCwP-NbY0s.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Skit #2', 32, 'vRBOIbTyTnU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Diamonds From Sierra Leone (Remix) (feat. JAY-Z)', 234, '4q7OpvvfjWs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('We Major (feat. Kanye West и Really Doe)', 448, '_fr4SV4fGAw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Skit #3', 25, 'HyXEzp85RGE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hey Mama', 306, 'B3NmMKfl3Ic.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Celebration', 199, 'FZjlP-N7Hl4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Skit #4', 79, 'Y4r6lS04RpQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gone (feat. Kanye West)', 334, 'TwPCaWQIJME.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Diamonds From Sierra Leone (Bonus Track)', 239, '92FCRmggNqQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Late', 231, 'YRwTaWWK3dI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: ye
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('ye', 'MPREb_vWD8jVytxdF.webp', v_artist_id, 'Ye — восьмой сольный студийный альбом американского хип-хоп-музыканта и продюсера Канье Уэста, изданный 1 июня 2018 года на лейблах звукозаписи GOOD Music и Def Jam.
В записи альбома приняли участие такие музыканты, как Ty Dolla Sign, PartyNextDoor, Кид Кади, Jeremih, 070 Shake, Charlie Wilson, Ники Минаж и другие. Продюсировал альбом сам Уэст при участии Mike Dean в качестве сопродюсера.
Ye дебютировал на первом месте американского хит-парада Billboard 200, став восьмым для Уэста чарттоппером в карьере.

Источник: Wikipedia (', '2018-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 7
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I Thought About Killing You', 275, 'no1YszVVybo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Yikes', 189, 'kPPyUO6m3-4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('All Mine', 146, 'TrQ7w1bdNvY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Wouldn''t Leave (feat. PARTYNEXTDOOR)', 206, 'nMkXJohQiuQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('No Mistakes', 124, '4I8gDpuvZt4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ghost Town (feat. PARTYNEXTDOOR)', 272, '5S6az6odzPI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Violent Crimes', 216, 'DSY7u8Jg9c0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: I Love It (Freaky Girl Edit)
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('I Love It (Freaky Girl Edit)', 'MPREb_cCrw0VV76Hp.webp', v_artist_id, '', '2018-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I Love It (Freaky Girl Edit)', 128, 'cwQgjq0mCdE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: I Love It
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('I Love It', 'MPREb_XhLKiOfJBQ0.webp', v_artist_id, '', '2018-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I Love It', 128, 'cwQgjq0mCdE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: FourFiveSeconds
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('FourFiveSeconds', 'MPREb_7K82R4fQFdl.webp', v_artist_id, '', '2015-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('FourFiveSeconds', 189, 'kt0g4dWxEBo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Forever
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Forever', 'MPREb_MM8zGSv9zTW.webp', v_artist_id, '', '2009-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: Mercy (feat. Big Sean, Pusha T и 2 Chainz)
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Mercy (feat. Big Sean, Pusha T и 2 Chainz)', 'MPREb_gfzhON8bnJU.webp', v_artist_id, '', '2012-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Mercy (feat. Big Sean, Pusha T и 2 Chainz)', 330, '7Dqgr0wNyPo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
-- Артист: Rammstein
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Rammstein', 'Rammstein — немецкая метал-группа, образованная в январе 1994 года в Берлине. Музыкальный стиль группы относится к жанру индастриал-метала.
Основные черты творчества группы: специфический ритм, в котором выдержана большая часть композиций, и эпатирующие тексты песен.
Особую известность группе принесли сценические выступления, часто сопровождаемые использованием пиротехники, получившие признание в музыкальной среде. Состав группы ни разу не менялся. По состоянию на 2018 год она продала около 20 млн копий альбомов.

Источник: Wikipedia (', 'UCYp3rk70ACGXQ4gFAiMr1SQ_avatar.webp', 'UCYp3rk70ACGXQ4gFAiMr1SQ_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Made In Germany 1995 - 2011
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Made In Germany 1995 - 2011', 'MPREb_AAiq8bxCEw6.webp', v_artist_id, 'Made in Germany 1995—2011 — сборник немецкой металл-группы Rammstein, выпущенный 2 декабря 2011 в Германии, Австрии и Швейцарии. Мировой релиз — 5 декабря 2011. В него вошли лучшие песни группы, а также новая композиция «Mein Land». Концертный тур в поддержку сборника начался 6 ноября 2011 года. На обложках сборника изображены «посмертные» маски участников группы.
Сборник доступен в трёх вариантах: Standard Edition, Deluxe Edition, Super Deluxe Edition. Видеоверсия сборника получила название Videos 1995–2012.

Источник: Wikipedia (', '2011-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 16
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Pussy', 239, 'zU9V9QMXeyc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Mutter
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Mutter', 'MPREb_KtV3PgMEVJL.webp', v_artist_id, 'Mutter — третий студийный альбом немецкой метал-группы Rammstein. Вышел 2 апреля 2001 года. Альбом записывался в Германии, Франции, Швеции и Америке. Журнал Metal Hammer включил Mutter в 200 лучших рок-альбомов всех времён. В России альбом получил премию «Рекордъ» в номинации «Зарубежный альбом года». В интервью Noizr Zine шведский продюсер и музыкант Петер Тэгтгрен посоветовал Mutter в качестве ориентира всем начинающим метал-продюсерам: «Он очень хорош, потому что в нём много различных элементов — в нём есть оркестровые партии, тяжёлые гитары, хорошее звучание барабанов — он может быть хорошим ориентиром».

Источник: Wikipedia (', '2001-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 11
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Spieluhr', 287, 'bWtYzU40N-M.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Zwitter', 258, 'LJhQd_G5WfQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Rein raus', 190, 'O2FZnh_bmIM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Adios', 229, 'rvwftBpqiiw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Nebel', 295, 'E3hh37KbY18.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Sehnsucht
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Sehnsucht', 'MPREb_1aVRMvtQMe2.webp', v_artist_id, 'Sehnsucht — второй альбом группы Rammstein. Выпущен в 1997 году. Почти сразу после выхода стал платиновым. Альбом отличала особая динамика и жёсткий, слегка рваный гитарный рифф. Выпускался ограниченный тираж альбома в диджипаке. Также выпускался Limited tour Box Edition, включавший футболку и постер.

Источник: Wikipedia (', '1997-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 11
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Sehnsucht', 245, 'ZXjNjuLOi_k.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Tier', 227, 'QX3FNE6xbco.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Bestrafe mich', 217, 'lY-GrC5ZixQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Spiel mit mir', 286, '_vVejzHNGUk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Klavier', 263, '6A8bV_IEgyI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Alter Mann', 263, 'Pu1Iuq0G-gI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Eifersucht', 216, '7opB3Fniyh4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Küss mich (Fellfrosch)', 211, 'ZKU3qOU9F74.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Reise, Reise
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Reise, Reise', 'MPREb_lOwEMAJYuQh.webp', v_artist_id, 'Reise, Reise — четвёртый студийный альбом немецкой индастриал-метал-группы Rammstein, выпущенный 27 сентября 2004 года на лейбле Universal. Он был записан в 2003—2004 годах в различных звукозаписывающих студиях в Италии, Франции и Швеции при участии продюсера Якоба Хелльнера. Благоприятное разрешение кризиса, наступившего в Rammstein во время записи предыдущего альбома, Mutter, повлияло на атмосферу записи Reise, Reise, сделав её менее напряжённой. Записав в непринуждённой обстановке 17 песен, Rammstein отобрали для альбома только 11, остальные 6 вошли в следующий альбом, Rosenrot.
Альбом был положительно воспринят критиками и попал в первые десятки хит-парадов ряда стран, таких как Германия, Франция, Бельгия, а также занял 61 место в Billboard 200. Reise, Reise получил «золотой» статус в Дании, Германии и Финляндии, «платиновый» — в Австрии и Швейцарии, а в России альбом стал дважды платиновым. В 2006 году Rammstein с песней «Mein Teil» были номинированы на премию «Грэмми» за лучшее метал-исполнение.

Источник: Wikipedia (', '2004-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 11
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dalai Lama', 339, 'LeEPvlARStw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Los', 264, 'TXYFFXZWp24.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Morgenstern', 240, 'pFBgHMbrQK4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Stein um Stein', 233, 'IKPbjkZOsiA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Amour', 291, 'bIPCVaLHZng.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Rammstein
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Rammstein', 'MPREb_pHn32Q6RzS3.webp', v_artist_id, 'Безымянный седьмой студийный альбом немецкой NDH-группы Rammstein, также называемый как Rammstein, одноимённо названию группы, был выпущен 17 мая 2019 года. Официально альбом был анонсирован 28 марта 2019 года, одновременно с выпуском сингла «Deutschland». Концертный тур в поддержку альбома начался 27 мая 2019 года.

Источник: Wikipedia (', '2019-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 11
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Diamant', 155, '3s0w9d7oMwQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Weit weg', 261, 'N9AalJuwLyQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Tattoo', 252, 'rWQ8TjsuH9g.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hallomann', 252, 'HYTXWZQVPyM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Rosenrot
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Rosenrot', 'MPREb_ohpbefaOCEe.webp', v_artist_id, 'Rosenrot — пятый студийный альбом немецкой Neue Deutsche Härte-группы Rammstein, выпущенный 28 октября 2005 года на звукозаписывающем лейбле Universal Music Group.
Rosenrot содержит шесть песен, не вошедших в альбом Reise, Reise, одну песню, написанную для альбома Mutter, и четыре новых композиции. Изначально диск должен был называться Reise, Reise Volume Two, но 18 августа 2005 года был анонсирован под названием Rosenrot.
В то время как Rammstein не занималась активной рекламой альбома, группа создавала ожидание от нового релиза различными способами;
Первый сингл с альбома, «Benzin», премьера которого состоялась в Berliner Wuhlheide, впоследствии был выпущен на CD. Эту и 5 других песен были представлены на официальном сайте группы в виде одноминутных семплов за восемь дней до выхода альбома.

Источник: Wikipedia (', '2005-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 11
-- Артист: Валентин Стрыкало
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Валентин Стрыкало', '«Валенти́н Стры́кало» — украинская рок-группа, основанная в 2010 году солистом Юрием Капланом, который получил известность после записи серии видеообращений к звёздам шоу-бизнеса от имени наивного провинциального парня «Валентина Стрыкало из села Бурильцево».
Каплан основал группу под влиянием музыки таких рок-групп, как «Сплин» и Radiohead. В 2012 году был выпущен дебютный альбом «Смирись и расслабься!», который по большей части был написан в жанре камеди-рок, пусть в нём и имелись песни с серьёзным характером. В следующем году вышел альбом «Часть чего-то большего», в котором упор был сделан больше на лирику, но также не обошлось и без завуалированного юмора в некоторых песнях. В 2016 году выходит альбом «Развлечение», в котором группа окончательно отошла от юмора и ушла к теме депрессии. Также в записи заметно, что группа вдохновлялась британской рок-группой Pink Floyd.
Летом 2018 года группа дала последние концерты и ушла в затишье, а в мае 2019 года Юрий Каплан сообщил, что группа распущена.

Источник: Wikipedia (', 'UCabtR67_U5O72yRXF7hiI-g_avatar.webp', 'UCabtR67_U5O72yRXF7hiI-g_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Смирись и расслабься
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Смирись и расслабься', 'MPREb_8WzKCTl3P8x.webp', v_artist_id, '«Смирись и расслабься!» — первый студийный альбом украинской группы «Валентин Стрыкало».
За три дня до официального релиза группа представила альбом на сайте Яндекс.Музыка для бесплатного прослушивания. На песни «Кайен», «Русский рок» и «Наше лето» были сняты клипы.
Юрий Каплан об альбоме:
В альбоме будет всем давно знакомый материал, который мы два года исполняем на концертах. Но его не было в качественных записях и люди не могли закинуть его себе в плеер. Мы сменили несколько студий прежде, чем остановились на удовлетворяющем нас варианте. Бывало, что материал накапливали, сводили и мастерили, но всё это переписывали опять. Фишка в том, что люди привыкли слышать эти песни на концертах, где всё звучит круто и драйвово. Поэтому в записи они ожидают того же драйва, что и на живых выступлениях. Но на концерте мы заряжены энергией зала и наша отдача гораздо круче. Очень сложно вывести себя на этот уровень в студии и добиться того же рок-н-ролла. Потому мы требовательно относились к выбору студии звукозаписи и звукорежиссёра.
— Официальный сайт Валентина Стрыкало

Источник: Wikipedia (', '2012-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 17
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Отель Кооператор', 236, '6NuTIggdXIY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Фанк', 194, 'nsfAj5wDBA0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Песня для девочек', 255, 'hc1Ih1wQDN0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Рустем', 212, '9Q0srCckNQE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('45 лет', 143, '1696AtvSpQM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Всё решено', 192, '_Yq2QAhLt_w.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кайен', 182, '3A2x-7HawDc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Он постоянно', 244, 'uSXblthoxY8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Лишь однажды', 173, 'c-0S68_ZLL8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Наше лето', 209, 'aeGdhFHFj8Q.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Русский рок', 194, 'pl1RGGyTGls.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Первомай', 216, 'eAqJiZzgokE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Я бью женщин и детей', 162, 'M1PrbQAkWP0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gay Porn', 320, 'UgSVLIzemYw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Так гріє', 279, 'H-394VBVvE0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Дешёвые драмы', 228, '1N0BaUxMfXU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Часть чего-то большего
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Часть чего-то большего', 'MPREb_UVzSEcp2hnu.webp', v_artist_id, '«Часть чего-то большего» — второй студийный альбом украинской группы «Валентин Стрыкало», вышедший 20 октября 2013 года.
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
VALUES ('Самый лучший друг', 166, 'a_NbJCJoUM4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Офисный стиляга', 134, 'CmHO2FKYGTs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Все мои друзья', 161, 'C7_oi97EzbM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Знаешь, Таня', 274, 'ryWtOu454VE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Космос нас ждет', 203, 'G8rBEpD2H3A.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Кладбище самолетов', 353, 'd-iiIzio_7g.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ты не такая', 238, 'fB60qciZP_E.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Танцы', 179, '-OlxXpKVv8A.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Преждевременное семяизвержение', 290, '-g6GYCazaio.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Сега – винтовар', 153, 'tcoAdT9XitY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Я стараюсь быть лучше', 327, 'OtEIrljyKBc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Взрослые травмы', 169, 'A8-Y_kmwqUU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Улица Сталеваров', 259, '5mEMmCeAVGw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Развлечение
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Развлечение', 'MPREb_47cJSJKG9bx.webp', v_artist_id, '«Развлечение» — третий и последний студийный альбом украинской группы «Валентин Стрыкало», релиз которого состоялся 13 октября 2016 года.

Источник: Wikipedia (', '2016-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 8
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ускользает', 306, 'PZZxcKYskpg.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('О брат', 295, '8JAfQuzbOBM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('92', 230, 'e8mLnAVQ9mY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Решится само собой', 253, 'kOfBszWa0t4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Делать это трезвым', 194, 'nAKX-7_fbw8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Тени', 345, 'FiynOYlSfrk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Бесполезно', 425, 'KdqzgAKoHTU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Подворотня - мой дом', 202, 'dFuIywTOJxc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Решится Само Собой
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Решится Само Собой', 'MPREb_Vnj7dnuAV3p.webp', v_artist_id, '', '2016-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Решится Само Собой', 252, 'OU56-kXpQRY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
-- Артист: Black Magick SS
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Black Magick SS', 'Black Magick SS — австралийская инди-группа, играющая в стилях блэк-метал, психоделический рок и синтезаторная музыка. Тематика песен преимущественно оккультная. Отличительной чертой группы является смесь визуальной эстетики нацистской Германии и психоделического искусства. Из-за этого, иногда группу относят к национал-социалистическому блэк-металу.

Источник: Wikipedia (', 'UCsrEMJwC_oq97_VS1U2tJ8Q_avatar.webp', 'UCsrEMJwC_oq97_VS1U2tJ8Q_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Rainbow Nights
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Rainbow Nights', 'MPREb_6KTEHzN3THy.webp', v_artist_id, '', '2025-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 6
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Endless Hallucinations', 361, 'NEaoTD2BDOs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Rainbow Nights', 296, '0t2WobVXC74.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Get Out', 301, '8Wpjm46y24c.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Kali', 318, '8H_rwspVAdU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Mother''s Lullaby', 309, 'AF50jIqlczg.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Truth', 264, '9qqU3nEF_HE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Burning Bridges
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Burning Bridges', 'MPREb_0ozxTuExzxB.webp', v_artist_id, '', '2025-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 7
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Burning Bridges', 416, '9dLhmRd1OmY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Breaking Free', 294, 'U72dK2D1TZ0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dinosaurs', 300, 'QjwVDRgyWUw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Let Go', 435, 'q6hcLGEOqJ0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Mystery', 374, 'Gcge-rqjPc4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Let The Magick In', 364, 'zmrY3TlpgWI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Bomb (Hidden Track)', 258, 'jGSZR9YDIXw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Spectral Ecstasy
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Spectral Ecstasy', 'MPREb_OBZ8eI2qwI1.webp', v_artist_id, '', '2025-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 6
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Black Hand', 287, 's-3oR442B4Y.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Spectral Ecstasy', 281, '9CLUzzz2Dk8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('My Love', 334, 'BFaG_kKhE-s.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Fallen Tale', 291, 'e6SJu7A6wTA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Oath', 361, 'HY6g_myMfIM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hymn Of Pride', 264, 'SKyldBrv-ZY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Black Abyss
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('The Black Abyss', 'MPREb_abPNzWSG24T.webp', v_artist_id, '', '2025-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 16
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Black Abyss', 324, 'Hg9udVAoaww.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Step Into The Night', 211, 'dtBNeSeaU04.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Crystal Eyes', 305, 'Nr9yX-T2u3U.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hidden In Plain Sight', 233, 'pKwbmk7p_0w.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Black Magick Army', 546, 'AS9Icc9H278.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Wisdom Tree', 236, 'zK234GPtq-8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('War Of The Sorcerer', 224, 'i7D26wm7Z6k.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dark Lord', 210, 'exCbvePeX7U.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Shall Prevail', 219, '9z2cXL86Cws.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Panzerwitch', 357, 'TmORC7jVPxc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Symbols Of Great Power', 87, 'TXc5v9NbqkU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Fire (Burns Forever)', 232, 'M6TwOqaxbA8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('In The Circle', 220, 'szs8r8onHDU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Gospel', 251, '-jzQzyOABa8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Owls Of Winter', 360, '2aTgkIHN51o.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Talisman', 257, 'WrhzdgLzpXw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Kaleidoscope Dreams
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Kaleidoscope Dreams', 'MPREb_iD1OesA2Itr.webp', v_artist_id, '', '2025-01-01', 'EP')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 6
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Kaleidoscope Dreams', 322, 'ZFeTSYyA4Rc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Crusader', 290, 'DENh9fktPtE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Tåget', 285, 'ujxWZyRt6fk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Power', 265, 'zfIe-ppDmwE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Follow You', 209, 'Z_oImpDyAQ0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Eclipse', 281, 'zJiurQp9vfs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
-- Артист: Bloodhound Gang
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Bloodhound Gang', '«Bloodhound Gang» — американская камеди-рок-группа из Филадельфии, начинавшая с хип-хопа под влиянием «Beastie Boys» и с ростом популярности начавшая использовать в музыке элементы поп-панка и электроник-рока. Группа продала по всему миру более 6 миллионов копий альбомов.

Источник: Wikipedia (', 'UCqLqkJrUJ36z5bm0erMDkjQ_avatar.webp', 'UCqLqkJrUJ36z5bm0erMDkjQ_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Show Us Your Hits (International Version)
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Show Us Your Hits (International Version)', 'MPREb_4AbL63LKBxS.webp', v_artist_id, '', '2010-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 14

-- Альбом/сингл: Hooray for Boobies
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Hooray for Boobies', 'MPREb_xh1bprIWhIf.webp', v_artist_id, 'Hooray for Boobies — третий студийный альбом американской группы Bloodhound Gang. Выход альбома состоялся 4 октября 1999 года в Соединённом Королевстве и 29 февраля 2000 года в Соединённых Штатах. Спродюсированный Джимми Попом и Ричардом Гавалисом, это второй альбом группы на лейбле Geffen Records после One Fierce Beer Coaster. Это был второй и последний альбом, в записи которого участвовал барабанщик Spanky G, покинувший группу, чтобы закончить учёбу.
Альбом вывел группу в мейнстрим, имел коммерческий успех и был благосклонно принят музыкальными критиками. В США он первоначально дебютировал на втором месте в чарте Top Heatseekers и достиг 14-й позиции в чарте Billboard 200. Он занял первое место в Австрии и Германии.
С альбома было выпущено пять синглов: «Along Comes Mary», «The Bad Touch», «The Ballad of Chasey Lain», «Mope» и «The Inevitable Return of the Great White Dope». «The Bad Touch» попал в 14 музыкальных чартов, в пяти из которых занял 1-ю позицию.

Источник: Wikipedia (', '1999-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 19
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Inevitable Return Of The Great White Dope', 243, 'ZxlMX6cDjpw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Mama''s Boy', 35, 'XxFqFGxGVQw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Yummy Down On This', 229, 'cnM_tJ32K2s.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('R.S.V.P.', 16, '-x3UkMehAqs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Bad Touch', 261, '523Lqlk4fm8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('That Cough Came With A Prize', 15, 'uKCvJiV7v6s.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Right Turn Clyde', 325, '9oda4um412E.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('This Is Stupid', 11, '_RuY2nzmA1w.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('A Lap Dance Is So Much Better When The Stripper Is Crying', 338, 'OKjgf5mBGoQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Ten Coolest Things About New Jersey', 11, 'ilbbaNqivQ8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Along Comes Mary', 202, 'hQt89O94BfU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hidden Track', 293, 'n2j5lfhuOOA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Hooray For Boobies (Expanded Edition)
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Hooray For Boobies (Expanded Edition)', 'MPREb_O4N5tKzLDCT.webp', v_artist_id, 'Hooray for Boobies — третий студийный альбом американской группы Bloodhound Gang. Выход альбома состоялся 4 октября 1999 года в Соединённом Королевстве и 29 февраля 2000 года в Соединённых Штатах. Спродюсированный Джимми Попом и Ричардом Гавалисом, это второй альбом группы на лейбле Geffen Records после One Fierce Beer Coaster. Это был второй и последний альбом, в записи которого участвовал барабанщик Spanky G, покинувший группу, чтобы закончить учёбу.
Альбом вывел группу в мейнстрим, имел коммерческий успех и был благосклонно принят музыкальными критиками. В США он первоначально дебютировал на втором месте в чарте Top Heatseekers и достиг 14-й позиции в чарте Billboard 200. Он занял первое место в Австрии и Германии.
С альбома было выпущено пять синглов: «Along Comes Mary», «The Bad Touch», «The Ballad of Chasey Lain», «Mope» и «The Inevitable Return of the Great White Dope». «The Bad Touch» попал в 14 музыкальных чартов, в пяти из которых занял 1-ю позицию.

Источник: Wikipedia (', '1999-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 30

-- Альбом/сингл: Hefty Fine
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Hefty Fine', 'MPREb_dc2FXOz4eVi.webp', v_artist_id, 'Hefty Fine — четвёртый студийный альбом американской альтернативной рок-группы Bloodhound Gang, выпущенный 27 сентября 2005 года. Спродюсированный Джимми Попом, он стал третьим альбомом Bloodhound Gang, изданным на лейбле Geffen Records после Hooray for Boobies, который разошёлся тиражом более миллиона копий в Соединённых Штатах и Европе. Кроме того, это был последний альбом, в записи которого принимал участие Люпус Тандер и единственный альбом с барабанщиком Willie the New Guy и впоследствии были заменены Дэниелом Картером и Адамом Перри.
В целом Hefty Fine получил отрицательные отзывы от критиков; на сайте-агрегаторе Metacritic оценка альбома составляет 28 баллов из 100. Несмотря на негативную реакцию критиков, альбом хорошо продавался, особенно в Европе, где он дебютировал в первой десятке в Австрии, Нидерландах и Германии. В США альбом достиг 24-го места в Billboard 200.
С альбома были выпущены три сингла «Foxtrot Uniform Charlie Kilo», «Uhn Tiss Uhn Tiss Uhn Tiss» и «No Hard Feelings». Первые два сингла альбома стали хитами, первый достиг позиций в шести чартах, а второй — в четырёх.

Источник: Wikipedia (', '2005-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 13
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Strictly For The Tardcore', 9, 'WsL2eFBOt20.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Balls Out', 260, 'zMUQGzjFBCs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Foxtrot Uniform Charlie Kilo', 172, 'bTQPkysSaRQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I''m The Least You Could Do', 239, 'HkUmwNMxUSw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Farting With A Walkman On', 207, 'RiFJJYO_SsU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Diarrhea Runs In The Family', 24, 'mKLbrjeDby8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ralph Wiggum', 173, 'DBgFAOf2Z84.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Something Diabolical', 311, 'DlzA2oipz3s.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Overheard In A Wawa Parking Lot', 5, 'wWaWi3SXhR0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Pennsylvania', 178, 'MbUkOQSeqWo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Uhn Tiss Uhn Tiss Uhn Tiss', 261, 'CYHhwqR-8lk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('No Hard Feelings', 312, 'gyV_e0FFe38.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hefty Fine', 5, 'y5hRBgqMNt8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: One Fierce Beer Coaster
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('One Fierce Beer Coaster', 'MPREb_7d5WvGzhOhS.webp', v_artist_id, 'One Fierce Beer Coaster — второй студийный альбом американской рэп-рок-группы Bloodhound Gang, выпущенный в 1996 году. Продюсером альбома был Джимми Поп. Это был первый релиз группы, изданный лейблом Geffen Records и первый, в котором участвовали в записи барабанщик Майкл «Spanky G» Гётер, басист Джаред Хассельхофф, и DJ Q-Ball. Музыкальный стиль One Fierce Beer Coaster основан на альтернативном роке, с рэп-речитативом, в большинстве песен исполненным Джимми Попом, и лирикой, связанной с туалетным юмором.

Источник: Wikipedia (', '1996-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 12
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Kiss Me Where It Smells Funny', 185, 'CIkPOhWS6fg.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Lift Your Head Up High (And Blow Your Brains Out)', 299, '5OvlrIwmt7k.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Fire Water Burn', 292, 'zEtV7PUrvPs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I Wish I Was Queer So I Could Get Chicks', 229, 'IQFDdZgkTKM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Why''s Everybody Always Pickin'' On Me?', 203, 'IzrJ4SaEOLs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('It''s Tricky', 157, 'E-h5oguZ8OA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Asleep At The Wheel', 246, 'dY2DdcnH1Fw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Shut Up', 195, '_2_mJ3N-SW8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Your Only Friends Are Make Believe', 423, '7QUgdeuwEek.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Boom', 246, '6Ldh-Ogt9HI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Going Nowhere Slow', 261, 'q9Vw15nDVKI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Reflections Of Remoh', 52, 'IZvAMkmSY3o.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Hard-Off
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Hard-Off', 'MPREb_L3MablXzM4F.webp', v_artist_id, '', '2015-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 11
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('My Dad Says That''s For Pussies', 172, 'wWPcz3wsPO4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dimes', 267, 'DzFDmqbDWdU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('American Bitches', 226, 'sPrWae5CtOc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Chew Toy', 195, 'qIqXnpcxGCA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Uncool As Me (feat. Joey Fatone)', 220, 'RnoEzxkhZrw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Clean Up In Aisle Sexy', 172, 'vicMVF8CHgI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Diary Of A Stranger', 226, '292m2iyhZ0Q.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Socially Awkward Penguin', 197, 'DAbuWSraUyM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Think Outside The Box', 256, 'mEJFB-wTDDo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('We''re Gonna Bring The Party To You', 300, 'LPPuVJ6kaJk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Bumblebees', 80, 'cMFCDGfg3w4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Bad Touch
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('The Bad Touch', 'MPREb_CbacQdBsMqZ.webp', v_artist_id, '', '1999-01-01', 'EP')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 6
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Bad Touch', 261, 'xat1GVnl8-k.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Uhn Tiss Uhn Tiss Uhn Tiss
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Uhn Tiss Uhn Tiss Uhn Tiss', 'MPREb_hO5r16ZCb4M.webp', v_artist_id, '', '2005-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 2
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Uhn Tiss Uhn Tiss Uhn Tiss', 261, '89tgpzE4qkY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Uhn Tiss Uhn Tiss Uhn Tiss (The Scooter Remix)', 428, 'ewS1Hy3QApI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Foxtrot
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Foxtrot', 'MPREb_iNnva6xOhFw.webp', v_artist_id, '', '2005-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 3
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Foxtrot Uniform Charlie Kilo', 172, 'JZpxaiNV_sM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Foxtrot Uniform Charlie Kilo (The Jason Nevins Mix)', 185, 'j6hbN-2yA14.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Foxtrot Uniform Charlie Kilo (The M.I.K.E. Mix)', 480, '7HuyV_0Cre4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Altogether Ooky
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Altogether Ooky', 'MPREb_EGc3w76e65w.webp', v_artist_id, '', '2010-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: Uncool As Me (feat. Joey Fatone)
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Uncool As Me (feat. Joey Fatone)', 'MPREb_RQZ41oNsvMK.webp', v_artist_id, '', '2015-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
-- Артист: Tyler, The Creator
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Tyler, The Creator', 'Та́йлер Гре́гори Око́нма, более известный под псевдонимом Tyler, The Creator — американский хип-хоп-исполнитель и музыкальный продюсер. Он является одной из самых влиятельных фигур массовой культуры своего поколения благодаря эклектичному творчеству. Ранние релизы рэпера имели мрачное звучание, вдохновлённое хорроркором, однако позже музыкант стал выпускать более лёгкую музыку, находящуюся под влиянием джаза и соула.
Оконма стал широко известен в конце 2000-х в качестве лидера музыкального коллектива Odd Future. Тайлер выпустил свой дебютный альбом Bastard в 2009 году и быстро привлёк внимание музыкальной прессы благодаря своему тяжёлому звучанию, вдохновлённому хорроркором, а также жестокому, трансгрессивному лирическому содержанию. Его вторая пластинка Goblin принесла известность в мейнстриме, чему способствовала популярность сингла «Yonkers».
Начиная со своего второго альбома Wolf, рэпер отходит от хорроркора в сторону более альтернативного звучания. Так, его третья пластинка Cherry Bomb содержит больше мелодичных и джазовых звуков, что сохраняется на протяжении последующей дискографии музыканта.

Источник: Wikipedia (', 'UCsQBsZJltmLzlsJNG7HevBg_avatar.webp', 'UCsQBsZJltmLzlsJNG7HevBg_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Flower Boy
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Flower Boy', 'MPREb_8DCHgYVRKEg.webp', v_artist_id, 'Flower Boy — пятый студийный альбом американского рэп-исполнителя Tyler, the Creator. Альбом был выпущен 21 июля 2017 года на лейбле Columbia Records. Будучи полностью спродюсированным Тайлером, альбом включает в себя гостевые участия Фрэнка Оушена, ASAP Rocky, Anna of the North, Лила Уэйна, Кали Учис, Стива Лейси, Эстеля, Смита Джейдена и Rex Orange Country. Flower Boy был поддержан выпуском четырёх синглов: «Who Dat Boy» / «911», «Boredom», «I Ain’t Got Time!» и «See You Again». Альбом получил положительные оценки критиков за его мягкое и плавное звучание, уникальную смесь жанров и продакшн. Альбом был назван среди лучших альбомов 2017 года и десятилетия, также был номинирован на премию «Лучший рэп альбом» в ходе церемонии награждения Грэмми 2018.

Источник: Wikipedia (', '2017-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 14
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Foreword (feat. Rex Orange County)', 195, 'MvEtKc8-n3s.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Where This Flower Blooms (feat. Frank Ocean)', 195, '1gAHhLb6tjA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Sometimes...', 37, 'rJAhkWa4UR4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Pothole (feat. Jaden Smith)', 237, '_wvhrRb3oU0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Garden Shed (feat. Estelle)', 224, '-VgLwAjz4oA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Droppin'' Seeds (feat. Lil'' Wayne)', 60, 'MebxCTlHKto.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('November', 226, 'JaXoDWSQQrk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Glitter', 225, 'iNP8_xtq8YU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Enjoy Right Now, Today', 236, 'nJDwwB1JCxM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: IGOR
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('IGOR', 'MPREb_TDKJNChbDy4.webp', v_artist_id, 'Igor — шестой студийный альбом американского рэпера Tyler, the Creator. Он был выпущен 17 мая 2019 года на лейбле Columbia. Альбом был полностью спродюсирован Тайлером. В записи принимали участие исполнители Playboi Carti, Lil Uzi Vert, Соланж, Канье Уэст и Джеррод Кармайкл, а также Сантиголд, Джесси Уилсон, La Roux, Си Ло Грин, Чарли Уилсон, Slowthai, Фаррелл Уильямс и другие в качестве бэк-вокалистов.
Igor получил многочисленные положительные отзывы от критиков и дебютировал под № 1 в чарте Billboard 200. Ведущий сингл альбома «Earfquake» достиг места № 13 в чарте Billboard Hot 100. Альбом получил премию «Грэмми» в категории «Лучший рэп-альбом» в 2020 году.

Источник: Wikipedia (', '2019-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 12
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('IGOR''S THEME', 201, '6S20mJvr4vs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I THINK', 213, 'm91Vq-Yd3BA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('EXACTLY WHAT YOU RUN FROM YOU END UP CHASING', 15, 'dqZ8vr_Q4UI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('RUNNING OUT OF TIME', 178, 'Uyf_lImpdRw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('NEW MAGIC WAND', 196, '2w8KUgIkAu8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('PUPPET', 180, 'OZzfUagtyPE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('WHAT''S GOOD', 206, 'on7pfd91cKc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('GONE, GONE / THANK YOU', 376, 'pVInBRkoKgY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I DON''T LOVE YOU ANYMORE', 162, 'ZJsJ07vk23o.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('ARE WE STILL FRIENDS?', 266, 'Gb76TgCUqAY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Goblin
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Goblin', 'MPREb_Wi24pEqUHHd.webp', v_artist_id, 'Goblin — второй студийный альбом американского рэпера Tyler, The Creator. Релиз состоялся 10 мая 2011 года на лейбле XL Recordings.

Источник: Wikipedia (', '2011-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 15
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Goblin', 409, 'yFJG5_FShNo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Radicals', 439, '4QjcYIJBfnc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Transylvania', 193, '7IhQjL-RyvA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Nightmare', 322, 'uGEKqadrUvQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Tron Cat', 254, 'q0tzR-RrGxM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Her', 212, 'it-quHqbcn8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Sandwitches (feat. Hodgy Beats)', 292, 'MYWA3Vx6k5Y.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Fish', 380, '11ClPoc3PYo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Analog (feat. Hodgy Beats)', 175, 'U_Zyucrhf34.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Window (feat. Domo Genesis, Frank Ocean, Hodgy Beats и Mike G)', 481, 'hVRB821xnvQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('AU79', 221, '8eOlul8Z7U4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Golden', 344, '9l0833Vn-VA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Wolf
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Wolf', 'MPREb_ZEurlrvAC2F.webp', v_artist_id, '', '2013-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 18
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('WOLF', 110, 'AknjeiApwhg.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Jamba (feat. Hodgy)', 213, 'bjTlmQCB020.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Cowboy', 196, 'kIrWCBIvXSA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Awkward', 228, 'aG3mbIZY92E.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Answer', 231, 'QIuIdZ_L4iY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Slater (feat. Frank Ocean)', 234, 'FD-0K8G3Nyk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('48', 248, '4DZuvcBj70w.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Colossus', 214, 'Nn0tZl3Plcs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('PartyIsntOver/Campfire/Bimmer (feat. Frank Ocean и Laetitia Sadier)', 439, 'ivI0J1Z1uK4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Pigs', 255, 'gmr4iZ5dimY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Parking Lot (feat. Casey Veggies и Mike G)', 234, 'FRaBSQlGCus.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Trashwang (feat. Na'' kel, Джаспер Долфин, Lucas и L-Boy)', 283, '6msCJ1rmeEE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Treehome95 (feat. Coco O. и Erykah Badu)', 181, 'kpVpEnP7ELM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Lone', 238, 'QICNyBMod6M.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: CALL ME IF YOU GET LOST
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('CALL ME IF YOU GET LOST', 'MPREb_rMqWqQe5N3L.webp', v_artist_id, 'Call Me If You Get Lost — седьмой студийный альбом американского рэпера, певца, автора песен и продюсера Tyler, the Creator. Он был выпущен 25 июня 2021 года на лейбле Columbia Records. Альбом содержит гостевые участия от 42 Dugg, Ty Dolla Sign, YoungBoy Never Broke Again, Лила Уэйна, Teezo Touchdown, Domo Genesis, Brent Faiyaz, Fana Hues, Daisy World, Lil Uzi Vert и Фаррелла Уильямса.
Альбом знаменует собой отход от более лёгкой и соуловой эстетики Igor и Flower Boy в пользу более жёстких инструменталов и «сырых» рифм, вдохновлённых серией микстейпов DJ Drama Gangsta Grillz, а также старыми работами самого Тайлера. Жанр пластинки варьируется от хип-хопа до регги. На обложке показано удостоверение личности персонажа альбома по имени «Тайлер Бодлер», чьё имя было вдохновлено французским поэтом Шарлем Бодлером.
Call Me If You Get Lost был поддержан двумя синглами: «Lumberjack» и «WusYaName». Работа получила широкое признание от критиков. Пластинка вошла в десятку лучших альбомов 2021 года по версии нескольких изданий. Call Me If You Get Lost дебютировал под номером один в американском чарте Billboard 200, став второй пластинкой рэпера, занявшей первое место в США.

Источник: Wikipedia (', '2021-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 16

-- Альбом/сингл: Cherry Bomb
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Cherry Bomb', 'MPREb_L5lmoTeD4CA.webp', v_artist_id, 'Cherry Bomb — четвёртый студийный альбом американского рэпера Tyler, The Creator. Релиз состоялся 13 апреля 2015 года. 9 апреля 2015 года альбом был неофициально анонсирован в iTunes, наряду с двумя другими треками. В записи альбома приняли участие Скулбой Кью, Чарли Уилсон, Кали Учис, Канье Уэст, Лил Уэйн, Фаррелл Уильямс и другие.

Источник: Wikipedia (', '2015-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 13

-- Альбом/сингл: Yonkers
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Yonkers', 'MPREb_EUyxlLjl8xa.webp', v_artist_id, '', '2011-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: Potato Salad
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Potato Salad', 'MPREb_scdO8zXLk9w.webp', v_artist_id, '', '2018-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: Who Dat Boy / 911
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Who Dat Boy / 911', 'MPREb_wKBXpyYdtjj.webp', v_artist_id, '', '2017-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 2
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Who Dat Boy (feat. A$AP Rocky)', 206, 'FUXX55WqYZs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('911 / Mr. Lonely (feat. Frank Ocean и Steve Lacy)', 256, 'khMb3k-Wwvg.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: BEST INTEREST
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('BEST INTEREST', 'MPREb_T0xl9CSgxup.webp', v_artist_id, '', '2020-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: OKRA
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('OKRA', 'MPREb_7HOBxDt5K2s.webp', v_artist_id, '', '2018-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
-- Артист: madk1d
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('madk1d', '', 'UCvJLQlVdWiIU0qYmowdKh6g_avatar.webp', 'UCvJLQlVdWiIU0qYmowdKh6g_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: sexyswag
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('sexyswag', 'MPREb_61Uy57PiM6Q.webp', v_artist_id, '', '2025-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 9
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('цена', 139, 'IwF94S-T9RA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('ты че обиделась', 90, 'z39pUlKXkaE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('супермаркет', 120, '0xNI38yfb4k.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('засосы', 100, 'XZhS8rSdcoI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('sexyswag2010', 90, 'ce29K68xVFI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('сдвг', 133, 'fdpjNzyyVJU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('последним летом', 124, 'sjcug8dL3yA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('танцор', 89, 'jimlL3ObVCA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Он сказал поехали!
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Он сказал поехали!', 'MPREb_0XVYaP8YdVa.webp', v_artist_id, '', '2024-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 8
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('типичная ситуация', 131, '9HuO63khusE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('1 мая', 128, 'WZmxsBTMglo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('самый лучший трудовик', 119, 'yU1LZhc6C2w.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('грустная история', 93, 'WFt9indFGjw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('незаконно', 95, 'Uo0eyGPeing.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('number 9', 103, 'xs0-K0YTi5c.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('заправка', 97, 'R1E2WS4ZaoY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('типичная ситуация (Acoustic)', 89, '1TJWiu2BOOI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: ты че обиделась
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('ты че обиделась', 'MPREb_bWqv8OrhdLV.webp', v_artist_id, '', '2025-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('ты че обиделась', 90, 'GaGOwFAlwkU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Толпы
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Толпы', 'MPREb_7Z4fGoT4cys.webp', v_artist_id, '', '2022-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Толпы', 91, 'rHIer9aa6PI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: цена
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('цена', 'MPREb_iHWoIVfBevK.webp', v_artist_id, '', '2025-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('цена', 139, 'k2uiDdvtxks.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Толпы (Speed Up)
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Толпы (Speed Up)', 'MPREb_R37Q6pWXd91.webp', v_artist_id, '', '2023-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Толпы (Speed Up)', 82, 'bqk447IzpaU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: дырки в штанах
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('дырки в штанах', 'MPREb_MkqXHOvLXD8.webp', v_artist_id, '', '2025-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('дырки в штанах', 107, '3KVuiRBk5RI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
-- Артист: CUPSIZE
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('CUPSIZE', '', 'UCNpdKmV1hHFuKM6DUxGMOBw_avatar.webp', 'UCNpdKmV1hHFuKM6DUxGMOBw_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Как испортить вечеринку?
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Как испортить вечеринку?', 'MPREb_z441XRKFkNg.webp', v_artist_id, '', '2023-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 15
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Юра, Юра', 129, 'jcrA-YE65fE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('По улице иду я', 153, 'vdR06ZtG-98.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Они все дрочат на тебя в интернете', 109, 'eiSs2jvDq0k.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Стенки моего подъезда', 150, 'WPT_wl-96ac.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Василий', 147, 'A8oZFXK4WUI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Травматика', 161, 'Pjm5u23Cor4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('И это прекрасно', 206, 'ef0TI8QgzgU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Клей', 146, 'bruZDcRPOD0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Целовались', 153, 'ZgkG7-xOBIU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Пьяные', 129, 'mIidmzOT8hE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Высокий градус', 154, '__cFrzeIdfU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Но им не смешно', 144, 'ikeivvIk2BA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Семнадцатилетняя', 153, '1RwysRh343c.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Я схожу с ума', 188, 'dJFImLaSuno.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('ДПП (Аутро)', 109, 'NbRwXSn0Q94.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: кажется, в аду прикольно, но меня выгнали б утром
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('кажется, в аду прикольно, но меня выгнали б утром', 'MPREb_MQkQ1erFtpG.webp', v_artist_id, '', '2024-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 10
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('влечение', 137, 'eAu6k-E7QuM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('привет, если ты мне не ответишь', 123, 'xiJdF73HWac.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('фура', 134, 'MojRcq-pG-A.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('мой врач думает, что у меня шизофрения', 134, 'CwqpI8cxGwA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('маршрутка', 190, 'FK2Ld4Czq-s.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('ну почему', 172, '6kt3mezkNHM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('я тупая, моя жизнь тупая', 187, 'oHe_lH8jNks.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('пока-пока', 173, 'Rw8SA_GJPDw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('нам это нравится', 178, 'zJHlMlziotA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('больше, чем творчество', 155, 'sOZbxAXG44I.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Еби меня, малышка
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Еби меня, малышка', 'MPREb_hvHO1zItSrB.webp', v_artist_id, '', '2023-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 12

-- Альбом/сингл: в моих легких выросли цветы
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('в моих легких выросли цветы', 'MPREb_T9DiLjhYMO3.webp', v_artist_id, '', '2025-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 10
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('107.1', 119, 'JGs67DvT6YI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('печаль', 130, 'nLu0B1LBuhM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('минус, плюс', 225, 'ryfq5MMBVOQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('переломай мои кости', 168, 'BU-eVgBpxyc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('давай увидимся', 241, 'Vtc0poIr6VQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('твои поцелуи', 152, 'j73y9jpvjGs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('кислород', 143, '63yKCcJS488.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('или хотя бы завтра...', 105, 'nwiBb8EazS0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('самокрутки', 154, 'VqRhjv-RzWA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('улыбнись', 270, 'mqKk7sYg-BU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: неуравновешеннолетниепесни Pt.1
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('неуравновешеннолетниепесни Pt.1', 'MPREb_pWAYNewGsxi.webp', v_artist_id, '', '2025-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 7
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('оригами', 144, 'KEviVqwhLVw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('шАхАшАхА', 174, 'oxv5WWJ0oAo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('песня про спид', 171, 'cTWbBiZinpA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('конъюнктивит', 208, 'mvfzBLpaFws.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('тварьтварьтварьтварь...', 186, 'HJ5lHPtYhis.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('злой отчим', 225, 'Z9og1UByvoc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: дели на два
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('дели на два', 'MPREb_INP2iwQG88n.webp', v_artist_id, '', '2024-01-01', 'EP')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 4
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('ты любишь танцевать', 145, 'LHiSYbn2pUU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('пятый элемент', 134, 's3HlHnZ3CpQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('целую тебя', 130, 'cWaPuzn0TEs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('воздух', 125, 'pi7Mv5ZNfYU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: ВШБ
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('ВШБ', 'MPREb_HRNK4vaXqSf.webp', v_artist_id, '', '2025-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('ВШБ', 110, 'PjZzsCrMOW8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: LSD
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('LSD', 'MPREb_N5HIEjEUMMd.webp', v_artist_id, '', '2025-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('LSD', 117, 'mn0kbXinT0U.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: В окно с тобой
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('В окно с тобой', 'MPREb_ynm7KzRghpz.webp', v_artist_id, '', '2024-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 2
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('В окно с тобой', 120, 'F5D04udwezc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('В окно с тобой (Slow Version)', 149, 'LgeXnsuW8fM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: ЦВЕТОФОБИЯ
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('ЦВЕТОФОБИЯ', 'MPREb_7Bkbmuees4R.webp', v_artist_id, '', '2023-01-01', 'EP')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 5
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Цветофобия', 110, '7Y1csVBy4lU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Паранойя', 118, 'LtDzAtauEJY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Оттенки', 98, 'I0XqEh95Tv0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Бред', 113, 'sb8xAyZoMsY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Километры', 109, '_QPyoV6OzeA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
-- Артист: Пошлая Молли
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Пошлая Молли', '«Пошлая Молли» — украинская рок-группа, основанная музыкантом Кириллом Тимошенко, более известным под псевдонимом Кирилл Бледный. Группа является одним из самых ярких представителей синти-панка на Украине, совмещая поп-панк с электронной музыкой. «Пошлая Молли» набрала свою популярность в 2017 году, в частности в социальной сети «ВКонтакте»

Источник: Wikipedia (', 'UCDVnp5x53g5L-kBvQsGOW9w_avatar.webp', 'UCDVnp5x53g5L-kBvQsGOW9w_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: 8 способов как бросить...
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('8 способов как бросить...', 'MPREb_a0CQisTzcSr.webp', v_artist_id, '«8 способов как бросить дрочить» — дебютный студийный альбом украинской рок-группы Пошлая Молли, выпущенный 24 февраля 2017 года на лейбле Poshlaya Molly.

Источник: Wikipedia (', '2017-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 8
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Даже моя бэйби не знает', 208, 'Lp_euDcDQ40.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Нон стоп', 214, 'x7FGjJ6xk9Y.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Ханнамонтана', 191, 'q6iafAecXeA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Супермаркет', 215, 'WmC66JLxYYw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Паки пуси', 144, 'Q-wb3fG_sCE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Молли', 230, '9VJmsGWuWw4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Тмстс', 161, 'og5k3cro3F4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Контракт
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Контракт', 'MPREb_PUW5laM4Mo7.webp', v_artist_id, '', '2021-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Контракт', 205, 'kPkft7RLX9s.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Дом Периньон
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Дом Периньон', 'MPREb_qH4hpOjShhY.webp', v_artist_id, '', '2021-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: #HABIBATI
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('#HABIBATI', 'MPREb_5tFV9wq3A49.webp', v_artist_id, '', '2022-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('#HABIBATI', 154, 'oNr5-rQnhiM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Адская колыбельная
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Адская колыбельная', 'MPREb_kIZd6xYM0Fk.webp', v_artist_id, '', '2024-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Адская колыбельная', 179, 'U-xw6e-62fw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: DON''T PLAY, BAE
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('DON''T PLAY, BAE', 'MPREb_JfdC7li4ftt.webp', v_artist_id, '', '2022-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('DON''T PLAY, BAE', 182, 'euXAv2aeOhs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
-- Артист: Linkin Park
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Linkin Park', 'Linkin Park — американская рок-группа, основанная в 1996 году. Нынешний состав группы включает в себя вокалиста/ритм-гитариста/клавишника Майка Шиноду, соло-гитариста Брэда Делсона, басиста Дэйва Фаррелла, диджея Джо Хана, вокалистку Эмили Армстронг и барабанщика Колина Бриттэна. В период выхода первых семи студийных альбомов в составе группы были вокалист Честер Беннингтон и барабанщик Роб Бурдон. Linkin Park приостановили деятельность после самоубийства Беннингтона 20 июля 2017 года. В сентябре 2024 года было объявлено о воссоединении группы с Армстронг и Бриттэном в составе.
В целом творчество Linkin Park характеризуется стилями альтернативный рок и ню-метал; ранний период представляет собой смешение хэви-метала и хип-хопа, поздний включает больше элементов электроники и поп-музыки. Дебютный альбом Hybrid Theory принес группе мировую известность и приобрел «бриллиантовый» статус RIAA. Следующий альбом, Meteora, закрепил успех, дебютировав на #1 строчке чарта Billboard 200. На третьем альбоме, Minutes to Midnight, группа начала эксперименты со звучанием. В конце 2000-х коллектив был в числе самых успешных и популярных рок-групп.

Источник: Wikipedia (', 'UCZU9T1ceaOgwfLRq7OKFU4Q_avatar.webp', 'UCZU9T1ceaOgwfLRq7OKFU4Q_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Meteora
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Meteora', 'MPREb_qMlbe7gLeuH.webp', v_artist_id, 'Meteora — второй студийный альбом американской рок-группы Linkin Park, спродюсированный Доном Гилмором и впервые выпущенный 25 марта 2003 года. Альбом последовал за проектом Reanimation, который содержал ремиксы с дебютного альбома Hybrid Theory. В течение примерно года после выхода альбома Linkin Park выпускали синглы с Meteora, включая «Somewhere I Belong», «Faint», «Breaking The Habit» и «Numb». Трек «Lying from You» был выпущен в качестве промосингла.
Meteora стал самым успешным альбомом в истории чарта Modern Rock Tracks, который специализируется на радио-ротации песен альтернативного рока. «Numb» стала, по результатам чарта, песней года. В США продано 6 200 000 копий альбома и 27 миллионов копий по всему миру. Инструментальный трек «Session» был номинирован на «Грэмми» как «Лучшее инструментальное рок-исполнение» 2003 года.

Источник: Wikipedia (', '2003-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 13
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Foreword', 14, 'U6R-twDkrcI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Don''t Stay', 188, 'oWfGOVWrueo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Somewhere I Belong', 214, 'zsCD5XCu6CM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Lying from You', 176, 'NjdgcHdzvac.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hit the Floor', 165, 'oMals9XXQY8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Easier to Run', 205, 'U5zdmjVeQzE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Faint', 163, 'LYU-8IFcDPw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Figure.09', 198, '6dEAeCHQrBs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Breaking the Habit', 197, 'v2H4l9RpkwM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('From the Inside', 176, 'YLHpvjrFpe0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Nobody''s Listening', 179, 'QJ87793QXes.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Session', 145, 'J1KqQYsUYIk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Numb', 188, 'kXYiU_JCYtU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Hybrid Theory
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Hybrid Theory', 'MPREb_odUae150mq3.webp', v_artist_id, 'Hybrid Theory — дебютный студийный альбом американской рок-группы Linkin Park, спродюсированный Доном Гилмором и выпущенный 24 октября 2000 года лейблом Warner Bros., на следующий год после выпуска мини-альбома Hybrid Theory EP, состоявшего из шести эксклюзивных композиций, не вошедших в полноценный Hybrid Theory. Название для альбома было взято из предыдущего названия группы. Диск имел огромный коммерческий успех. В 2001 году в США было продано почти 5 миллионов копий альбома, что сделало Hybrid Theory самым продаваемым альбомом года в стране. Альбом поднялся на вторую позицию в американском чарте Billboard 200, а также занял высокие позиции в других мировых чартах. Hybrid Theory был записан на студии NRG Recording Studios, расположенной в Северном Голливуде, и спродюсирован Доном Гилмором. Тексты песен затрагивают проблемы вокалиста группы Честера Беннингтона, которые случились с ним во времена его юности, среди них: наркомания, постоянные ссоры и развод его родителей.
Четыре песни из альбома были выпущены в качестве синглов: «One Step Closer», «Crawling», «Papercut», и «In the End», все четыре сингла имели успех в различных хит-парадах и принесли группе большую известность.

Источник: Wikipedia (', '2000-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 12
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Papercut', 185, 'vjVkXlxsO8Q.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('One Step Closer', 158, '4qlCC1GOwFw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('With You', 204, 'M8UTS2iFXOo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Points of Authority', 201, 'jZSPAp8kCl4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Crawling', 209, 'Gd9OhYroLN0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('By Myself', 190, 'wWBp-nlGX1o.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('In the End', 217, 'eVTXPUF4Oz4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('A Place for My Head', 185, '3t2WkCudwfY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Forgotten', 195, 'HNCgBuI2eJc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Cure for the Itch', 158, 'qqC5sdsHLq8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Pushing Me Away', 192, 'Ve1LNJEIKUE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: LIVING THINGS
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('LIVING THINGS', 'MPREb_1tfBwXcIkmu.webp', v_artist_id, 'Living Things — пятый студийный альбом американской рок-группы Linkin Park. Он был записан в 2011—2012 годах в Северном Голливуде, на студии NRG Studios в сотрудничестве с продюсером Риком Рубином и выпущен в США 26 июня 2012 года на лейбле Warner Bros. Records. Пытаясь отделаться от статуса ню-метал-группы и избавиться от своего прежнего, привычного звучания, в новом альбоме Linkin Park добавили к своему звучанию такие жанры, как фолк и электронная музыка. Тематически альбом отличается от своего предшественника, A Thousand Suns, он повествует о проблемах личного характера и основан на собственном опыте музыкантов.
Уже в первую неделю после начала продаж альбом занял первое место в Billboard 200 и возглавил хит-парады ряда других стран, таких как Великобритания, Германия, Канада. В поддержку альбома Linkin Park выпустили 4 сингла, организовали два концертных тура: всемирный Living Things Tour и американский Honda Civic Tour.

Источник: Wikipedia (', '2012-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 12
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('LOST IN THE ECHO', 206, 'co4YpHTqmfQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('IN MY REMAINS', 201, 'QLFiuNdQrzI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('BURN IT DOWN', 231, 'dxytyRy-O1k.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('LIES GREED MISERY', 147, '9Dq9q6afIP8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I''LL BE GONE', 212, 'Y1wM5ljye28.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('VICTIMIZED', 107, '-6eUCOFXuxo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('ROADS UNTRAVELED', 230, 'KLgQKiUk7ms.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('SKIN TO BONE', 169, 'NwK4mxK7c2w.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('UNTIL IT BREAKS', 224, 'ZEQJm49_9OE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('TINFOIL', 72, 'M2YPuL-niQg.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('POWERLESS', 225, '32BOmle7Z6w.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: One More Light
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('One More Light', 'MPREb_t7vewLEKKPi.webp', v_artist_id, 'One More Light — седьмой студийный альбом американской мультиплатиновой рок-группы Linkin Park, вышедший 19 мая 2017 года на лейблах Warner Bros. Records и Machine Shop Recordings. Альбом сильно отличается от предшествующего The Hunting Party — One More Light не содержит тяжёлых гитарных риффов и записан в жанрах поп и поп-рок.
Первым синглом с альбома является трек «Heavy», вышедший 16 февраля 2017 года, 9 марта на официальном YouTube-канале группы появился официальный видеоклип. Группа также выпустила несколько промосинглов — «Battle Symphony», «Good Goodbye» и «Invisible». На каждый из синглов было выпущено лирик-видео, 5 мая в сеть был выложен официальный клип на песню «Good Goodbye».
В записи альбома участвовали такие музыканты, как Kiiara, Pusha T и Stormzy.
Это также последний альбом группы, в записи которого приняли участие двое её давних участников: со-ведущий вокалист Честер Беннингтон, который покончил жизнь самоубийством, повесившись через два месяца после выхода альбома, и барабанщик и соучредитель группы Роб Бурдон, который решил не возвращаться для воссоединения в 2024 году.

Источник: Wikipedia (', '2017-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 10
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Nobody Can Save Me', 226, 'FY9v147BZuE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Good Goodbye (feat. Pusha T)', 212, 'phVQZrb2AdA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Talking to Myself', 232, 'lvs68OKOquM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Battle Symphony', 217, 'D7ab595h0AU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Invisible', 215, 'n9_gWiwAWrA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Heavy (feat. Kiiara)', 170, '5dmQ3QWpy1Q.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Sorry for Now', 204, 'ylj0Xiiw1pM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Halfway Right', 218, 'ivSihne3rO8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Sharp Edges', 179, 'M5Ni_LskhFc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Minutes to Midnight
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Minutes to Midnight', 'MPREb_VKZiZdjp9B6.webp', v_artist_id, 'Minutes to Midnight — третий студийный альбом мультиплатиновой американской рок-группы Linkin Park, релиз которого состоялся 14 мая 2007 года, за исключением Северной Америки, где он был выпущен 15 мая. Название альбома связано с Часами Судного дня. Первым синглом стал трек «What I’ve Done», вторым — «Bleed It Out», третьим — «Shadow of the Day», четвёртым — «Given Up». Последним же синглом стала песня «Leave Out All the Rest». На все синглы были сняты клипы.
В то время как Шинода признаёт, что был определенный соблазн воспользоваться хорошо проверенной формулой успеха первых двух альбомов, он говорит, что Linkin Park находится в более благоприятной обстановке, чтобы пересмотреть свой звук под руководством нового продюсера Рика Рубина.
Альбом стал «Номером Один» в 28 странах, где был выпущен в 2007 году, его синглы попали в десятку лучших практически на всех территориях, а общее число проданных альбомов превысило отметку в 5 миллионов во всём мире.
В Tour Edition этого альбома бонусом вошли три трека:
«No Roads Left», «What I’ve Done » и «Given Up »

Источник: Wikipedia (', '2007-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 12
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Wake', 101, 'Me7TJDHCELk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Given Up', 190, '0xyxtzD54rM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Leave Out All The Rest', 210, 'yZIummTz9mM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Bleed It Out', 165, 'OnuuYcqhzCE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Shadow of the Day', 290, 'n1PCW0C1aiM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('What I''ve Done', 206, '8sgycukafqQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hands Held High', 234, 'gG4P3ayBzVY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('No More Sorrow', 222, 'rW4uBvP2Dqc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Valentine''s Day', 197, 'KAFOpywZbMM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('In Between', 197, 'YgVzhgygYfs.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('In Pieces', 219, 'NaRBn6QIMcQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Little Things Give You Away', 384, 'Gs0t8LXH6lw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: A Thousand Suns
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('A Thousand Suns', 'MPREb_uZFklMDQNMM.webp', v_artist_id, 'A Thousand Suns — четвёртый студийный альбом американской рок-группы Linkin Park, релиз которого состоялся 14 сентября 2010 года; альбом был слит в сеть за неделю до официального релиза — 7 сентября. Альбом был спродюсирован вокалистом группы Майком Шинодой и Риком Рубином, которые вместе работали над продюсированием предыдущего студийного альбома группы Minutes to Midnight. Сеансы записи для A Thousand Suns проходили на студии NRG Recording Studios в Северном Голливуде с конца 2008 по начало 2010 года.
A Thousand Suns — это концептуальный альбом, посвященный человеческим страхам, таким как ядерная война. Группа заявила, что альбом резко отличается от их предыдущей работы; они экспериментировали с разными и новыми звуками. Шинода сказал MTV, что альбом затрагивает множество социальных проблем и сочетает в себе человеческие идеи и технологии. Название является отсылкой к Бхагавад Гите, строка в которой была впервые популяризирована в 1945 году Робертом Оппенгеймером, который описал атомную бомбу как «яркую, как тысяча солнц». Он также появляется в строчке из первого сингла альбома «The Catalyst».

Источник: Wikipedia (', '2010-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 15
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Waiting for the End', 232, '5qF_qbaWt3Q.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Bleed It Out (Int''l DMD Single)
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Bleed It Out (Int''l DMD Single)', 'MPREb_i48WbODC7VY.webp', v_artist_id, '', '2007-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Bleed It Out', 165, 'OnuuYcqhzCE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Numb / Encore: MTV Ultimate Mash-Ups Presents Collision Course
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Numb / Encore: MTV Ultimate Mash-Ups Presents Collision Course', 'MPREb_wxK8avWNddb.webp', v_artist_id, '', '2004-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 2

-- Альбом/сингл: Final Masquerade
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Final Masquerade', 'MPREb_iOrokVLj6q1.webp', v_artist_id, '', '2014-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: Collision Course
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Collision Course', 'MPREb_Qlm86odf6dX.webp', v_artist_id, 'Collision Course — совместный мини-альбом рэпера Jay-Z и рок-группы Linkin Park. Был выпущен 30 ноября 2004 года в двух версиях: на CD и DVD. После выпуска достиг первого места в чарте Billboard 200. В августе 2009 года продажи альбома только в США достигли 1 934 000 экземпляров. По всему миру было продано около пяти миллионов экземпляров.

Источник: Wikipedia (', '2004-01-01', 'EP')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 6
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Jigga What / Faint', 212, 'LYU-8IFcDPw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Waiting for the End
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Waiting for the End', 'MPREb_mu6R9T7GEbZ.webp', v_artist_id, '', '2010-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 2
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Waiting for the End', 232, '5qF_qbaWt3Q.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
-- Артист: Red Hot Chili Peppers
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Red Hot Chili Peppers', 'Red Hot Chili Peppers — американская рок-группа, образованная в 1983 году в Калифорнии вокалистом Энтони Кидисом, басистом Майклом Бэлзари, гитаристом Хиллелом Словаком и барабанщиком Джеком Айронсом. Обладает 7 премиями «Грэмми». Во всём мире проданы более 80 миллионов копий их альбомов. По версии VH1 «100 Greatest Artists of Hard Rock» заняли 30-е место. 14 апреля 2012 года группа была включена в Зал славы рок-н-ролла. Группа заняла третье место в символическом списке «Лучшие исполнители за 10 лет скробблинга» портала Last.fm.

Источник: Wikipedia (', 'UCEuOwB9vSL1oPKGNdONB4ig_avatar.webp', 'UCEuOwB9vSL1oPKGNdONB4ig_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Greatest Hits
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Greatest Hits', 'MPREb_VDa115c6SZZ.webp', v_artist_id, 'Greatest Hits — второй официальный сборник суперхитов рок-группы Red Hot Chili Peppers, был выпущен в ноябре 2003 года. Помимо кавер-версии песни «Higher Ground», все треки были записаны в период сотрудничества с лейблом Warner Bros. — с 1991 по 2002 годы. Помимо этого, в сборник были включены две новые песни.
Также, отдельно был выпущен сборник Greatest Videos, содержащий большую часть музыкальных клипов группы того же временного периода.

Источник: Wikipedia (', '2003-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 16
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Under the Bridge', 266, 'GLvohMXgcBo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Give It Away', 285, 'Mr_uHJPUlO8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Californication', 330, 'YlUKcNNmywk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Scar Tissue', 216, 'mzJj5-lubeM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Soul to Squeeze', 290, '0XcN12uVHeQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Otherside', 256, 'rn_YodiJO6k.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Suck My Kiss', 216, 'C6jElKMMOWM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('By the Way', 216, 'JnfyjwChuNU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Parallel Universe', 270, 'mjHDqq1pWYw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Breaking the Girl', 295, 'iyu04pqC8lE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('My Friends', 250, '0kT5w27YxyI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Universally Speaking', 257, 'CoOibAstPJ4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Road Trippin''', 205, '11GYvfYjyV0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Fortune Faded', 202, 'TYYW_WwYHuM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Save the Population', 246, 'tSjw2LMtl9U.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: By the Way
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('By the Way', 'MPREb_uTi9HLPMEFd.webp', v_artist_id, 'By the Way — восьмой студийный альбом американской рок-группы Red Hot Chili Peppers, был выпущен в июле 2002 года на лейбле Warner Bros. Records. Спродюсирован Риком Рубином, как и три предыдущие альбома коллектива — Blood Sugar Sex Magik, One Hot Minute, Californication.
За первую неделю продаж By the Way разошёлся в количестве более 286.000 копий и к концу месяца достиг второй строчки хит-парада Billboard 200. В поддержку альбома были выпущены синглы: «By the Way», «The Zephyr Song», «Can''t Stop», «Dosed» и «Universally Speaking». В отличие от предыдущих пластинок группы, вокалист Энтони Кидис использовал более откровенный и рефлексивный подход к сочинению лирики.
By the Way получил положительные отзывы критиков, которые приветствовали изменения стиля, мелодизм и сдержанные эмоции. Автором большинства мелодий, басовых партий и гитарных прогрессий был гитарист Джон Фрушанте, что существенно повлияло на содержание альбома. По оценке издания PopMatters, «тёплая, негромкая гитара Джона Фрушанте и его вокальные партии а-ля ду-воп поистине королевские». Лонгплей содержит очень мало материала в стиле фанк-панк, который принёс известность группе в начале 1990-х.

Источник: Wikipedia (', '2002-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 16
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('By the Way', 217, 'JnfyjwChuNU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Universally Speaking', 257, 'CoOibAstPJ4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('This Is the Place', 258, 'gqgm7ViA2Ag.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dosed', 312, 'WeMXdaId60U.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Don''t Forget Me', 278, 'SnQ0E-UVt1g.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Zephyr Song', 232, '0fcRa5Z6LmU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Can''t Stop', 269, '8DyziWtkfBw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I Could Die for You', 193, '5hEjkH2DF5c.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Midnight', 296, '4QhFuJWqTAc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Throw Away Your Television', 225, 'Z2f6JmTssLQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Cabron', 219, '5wWd1a5U3vE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Tear', 318, 'etQKBq0Op4s.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('On Mercury', 208, 'X4ahMG3Iu8w.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Minor Thing', 218, 'iuCWDFNyyU4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Warm Tape', 256, 'iiyQWOSpOEE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Venice Queen', 368, '3s86rJvMIS0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Californication
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Californication', 'MPREb_vXDfRZDGgb3.webp', v_artist_id, 'Californication — седьмой студийный альбом американской рок-группы Red Hot Chili Peppers, выпущенный 8 июня 1999 года. Как и две предыдущие пластинки группы, альбом был спродюсирован Риком Рубином и выпущен на мейджор-лейбле Warner Bros. Records. Californication ознаменовал возвращение гитариста Джона Фрушанте, покинувшего Red Hot Chili Peppers в 1992 году во время тура в поддержку альбома Blood Sugar Sex Magik. Фрушанте официально вернулся в состав группы в 1998 году, заменив Дэйва Наварро. Помимо привычных для группы сексуальных тем, альбом затрагивает такие темы как смерть, самоубийство, наркотики, глобализация, а также путешествия.
В альбом вошли такие хиты как «Californication», «Around the World», «Otherside», а также «Scar Tissue», выигравшая премию «Грэмми» в категории «Лучшая рок-песня».
Californication является самым коммерчески успешным альбомом группы. Он добрался до 3-й строчки хит-парада Billboard 200. На сегодняшний день Californication продан тиражом более 16 миллионов копий, включая 5 миллионов в США.

Источник: Wikipedia (', '1999-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 15
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Around the World', 239, 'a9eNQZbjpJk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Parallel Universe', 270, 'N5Vgm6-KO3o.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Scar Tissue', 216, 'mzJj5-lubeM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Otherside', 256, 'rn_YodiJO6k.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Get on Top', 199, 'NJnOl3trYW0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Californication', 330, 'YlUKcNNmywk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Easily', 232, 'C-2xtTpgQXM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Porcelain', 164, 'j0tHXHcq724.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Emit Remmus', 241, 'bR6lbN40-Ho.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I Like Dirt', 158, '5fdJdfXid70.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('This Velvet Glove', 226, 'TqoDFpLyio0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Savior', 293, 'UijW9hGpnzc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Purple Stain', 254, 'H1LF-qyoNjo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Right on Time', 113, 'PSqH-4Wubq0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Road Trippin''', 205, '11GYvfYjyV0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Getaway
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('The Getaway', 'MPREb_0p4Lr89w8y0.webp', v_artist_id, 'The Getaway — одиннадцатый студийный альбом американской рок-группы Red Hot Chili Peppers, выпущенный 17 июня 2016 года на лейбле Warner Bros. Это первый за 25 лет альбом RHCP, в работе над которым не принимал участие продюсер Рик Рубин. Ему на смену был приглашён Danger Mouse. 5 мая был выпущен дебютный сингл альбома — «Dark Necessities». Это второй и последний альбом записанный при участии Джоша Клингхоффера, который пробыл в составе коллектива до 15 декабря 2019 года, когда музыканты решили воссоединиться с Джоном Фрушанте.

Источник: Wikipedia (', '2016-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 13
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Getaway', 251, '5w4d_mmJOGM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dark Necessities', 302, 'Q0oIoR9mLwc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('We Turn Red', 201, 'I7rlCyvHejY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Longest Wave', 212, 'HWOblzU6-hM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Goodbye Angels', 269, 'txTqtm58AqM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Go Robot', 264, 'HI-8CVixZ5o.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Feasting on the Flowers', 203, '34aEZpGIRDU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Detroit', 227, '_Zm6Iy0wMWk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('This Ticonderoga', 216, '773bV7xNQOQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Encore', 255, '63xPXEB4fjA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Hunter', 241, 's2TO6-zct04.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dreams of a Samurai', 370, '86zMV7U30nE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Stadium Arcadium
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Stadium Arcadium', 'MPREb_6NxJve9rqGu.webp', v_artist_id, 'Stadium Arcadium — девятый студийный альбом американской рок-группы Red Hot Chili Peppers. Этот двойной альбом впервые вышел в Германии 5 мая 2006 года, а 9 мая 2006 года пластинка была издана в США на Warner Bros. Records. Он включает пять синглов: «Dani California», «Tell Me Baby», «Snow », «Desecration Smile» и «Hump de Bump». Также был опубликован первый клип группы, сделанный фанатами, — на песню «Charlie». В США Stadium Arcadium стал первым альбомом группы, возглавившим национальные чарты. Первоначально Stadium Arcadium планировался как трилогия альбомов, которые должны были выходить с интервалом в шесть месяцев, но в итоге был сокращён до двойного альбома.
За первую неделю «Stadium Arcadium» разошёлся тиражом в 442 000 экземпляров в США и впервые в истории группы дебютировал на первом месте в Billboard 200.
Альбом хвалили за объединение музыкальных стилей, отражающих разные этапы карьеры группы. Он принёс группе семь номинаций на премию «Грэмми» в 2007 году, включая награду за лучший рок-альбом и одну за лучшее оформление бокс-сета или специального ограниченного издания.

Источник: Wikipedia (', '2006-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 29
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dani California', 283, 'Sb5aq5HcS1A.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Snow (Hey Oh)', 335, 'yuFI5KSPAt4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Charlie', 278, 'wNvOUkRTkz8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Stadium Arcadium', 315, 'j9qfClVvfIw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hump de Bump', 214, 'OM9uMJWtNww.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('She''s Only 18', 206, 'ZEKSY7XempM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Slow Cheetah', 320, '-877RlLrhJA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Torture Me', 225, 'jf5GSKFdrOA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Strip My Mind', 260, 'Gp7rGy0UkYU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Especially in Michigan', 241, '16zvMPHNAn4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Warlocks', 206, 'd-Ez8JMxDqc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('C''mon Girl', 229, 'HjbY3a1W7DY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Wet Sand', 310, 'oabjND9QW8Q.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hey', 340, 'U6eFQDaJmnk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Desecration Smile', 302, 'v-nNksBWxNI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Tell Me Baby', 248, 'oDNcL1VP3rY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hard to Concentrate', 242, 'G2dR2DV-eGc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('21st Century', 263, 'hYWGhQJbJ_c.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('She Looks to Me', 246, '1_yirYhYLDU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Readymade', 271, 'RZ83DshL4LM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('If', 173, 'NNXgvAVRflg.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Make You Feel Better', 232, 'I3-TO3CE2dY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Animal Bar', 326, 'HeSatuQQpI8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('So Much I', 225, 'JZktx4bKZNc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Storm in a Teacup', 225, '6odVqa_Q2Hw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('We Believe', 216, 'yqaxwen5Kj0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Turn It Again', 366, 'Vka1W1d7G9A.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Death of a Martian', 265, '44S9JcqkioY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Audio Commentary for Stadium Arcadium (Short Version)', 1754, '2godlilX3nc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Studio Album Collection 1991 - 2011
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('The Studio Album Collection 1991 - 2011', 'MPREb_TYSqP8qBN0T.webp', v_artist_id, '', '2014-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 103
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Power of Equality', 244, 'RVMykgBN8-Q.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('If You Have to Ask', 217, '7z4bZSa09Uw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Breaking the Girl', 295, 'iyu04pqC8lE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Funky Monks', 324, 'KdYj_CbUqaY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Suck My Kiss', 216, 'C6jElKMMOWM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('I Could Have Lied', 245, 'inGcYraci2A.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Mellowship Slinky in B Major', 240, 'YLeUi0FFDmA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Give It Away', 285, 'Mr_uHJPUlO8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Blood Sugar Sex Magik', 272, 'uLTEEndUoPk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Under the Bridge', 266, 'GLvohMXgcBo.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Naked in the Rain', 266, '7Hbxnze29oU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Apache Rose Peacock', 282, 'mrr1wzjjgTw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Greeting Song', 194, 'Ou-Vl-2gSiE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('My Lovely Man', 280, 'ZplAKsst608.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('They''re Red Hot', 72, 'M8yiBYrEcjQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Warped', 305, 'xmyuJZH3RAc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Aeroplane', 286, 'vV8IAOojoAA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Deep Kick', 394, 'VtwSM9eSZb8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('My Friends', 250, '0kT5w27YxyI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Coffee Shop', 189, 'WkkKStRwokQ.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Pea', 108, 'eqTwMykA-OA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('One Big Mob', 363, 'l4qfHKbkjEg.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Walkabout', 308, 'AffnBvnTzGw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Around the World', 241, 'a9eNQZbjpJk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Scar Tissue', 216, 'mzJj5-lubeM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Otherside', 256, 'rn_YodiJO6k.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Californication', 330, 'YlUKcNNmywk.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Road Trippin''', 205, '11GYvfYjyV0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('By the Way', 216, 'JnfyjwChuNU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Universally Speaking', 257, 'CoOibAstPJ4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Zephyr Song', 232, '0fcRa5Z6LmU.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Can''t Stop', 269, '8DyziWtkfBw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dani California', 283, 'Sb5aq5HcS1A.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Snow (Hey Oh)', 336, 'yuFI5KSPAt4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Charlie', 278, 'wNvOUkRTkz8.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Hump de Bump', 214, 'OM9uMJWtNww.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Desecration Smile', 303, 'v-nNksBWxNI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Tell Me Baby', 248, 'oDNcL1VP3rY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Dark Necessities
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Dark Necessities', 'MPREb_i6ho3skj4JZ.webp', v_artist_id, '', '2016-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dark Necessities', 302, 'Q0oIoR9mLwc.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Snow (Hey Oh)
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Snow (Hey Oh)', 'MPREb_6GeHwecKVoa.webp', v_artist_id, '', '2006-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 4
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Snow (Hey Oh)', 335, 'yuFI5KSPAt4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Dani California
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Dani California', 'MPREb_fC5hKrjaveN.webp', v_artist_id, '', '2006-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 4
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Dani California', 283, 'Sb5aq5HcS1A.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Getaway
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('The Getaway', 'MPREb_P1XwF6V7i2D.webp', v_artist_id, '', '2016-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('The Getaway', 251, '5w4d_mmJOGM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Tell Me Baby
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Tell Me Baby', 'MPREb_V2oL3jZDjL8.webp', v_artist_id, '', '2006-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 4
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Tell Me Baby', 248, 'oDNcL1VP3rY.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
-- Артист: Гражданская Оборона
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Гражданская Оборона', '«Гражда́нская оборо́на» — советская и российская рок-группа, основанная 8 ноября 1984 года в Омске Егором Летовым и Константином Рябиновым, наиболее яркий представитель сибирского панк-рока.
Музыка коллектива на начальном этапе представляла собой панк-рок с сильным влиянием гаражного рока, сохранявшимся на протяжении всей творческой деятельности группы, а в 1990-х годах её стилистика сместилась в сторону психоделического рока. «Гражданская оборона» являлась одной из наиболее значительных панк-рок-групп СССР и России. Группа распалась в 2008 году после смерти Егора Летова.

Источник: Wikipedia (', 'UCeMsJJOE6avjyvbqP4Kf24g_avatar.webp', 'UCeMsJJOE6avjyvbqP4Kf24g_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Лунный переворот
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Лунный переворот', 'MPREb_7wrYD8LxZ09.webp', v_artist_id, '«Солнцеворо́т» — студийный альбом группы «Гражданская оборона». Записывался в домашней студии Егора Летова с мая 1995 по октябрь 1996 года вместе с альбомом «Невыносимая лёгкость бытия». Оба альбома похожи по звучанию и стилю. В альбом «Солнцеворот» вошли песни, записанные раньше, «Невыносимую лёгкость бытия» — позже. В 2005 году был переиздан под названием «Лунный переворот».

Источник: Wikipedia (', '1996-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 14

-- Альбом/сингл: Реанимация
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Реанимация', 'MPREb_DlUv3HGVJfI.webp', v_artist_id, '«Реанима́ция» — двадцать второй студийный альбом группы «Гражданская оборона». Вторая часть дилогии «Долгая счастливая жизнь/Реанимация»

Источник: Wikipedia (', '2017-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 13

-- Альбом/сингл: The Best, Pt. 2
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('The Best, Pt. 2', 'MPREb_F6Rc2pFjD0o.webp', v_artist_id, '', '2017-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 25
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Свобода', 178, 'JXkxMfTe6N0.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Долгая счастливая жизнь
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Долгая счастливая жизнь', 'MPREb_bvDzbQNM3f3.webp', v_artist_id, '«До́лгая счастли́вая жизнь» — двадцать первый альбом российской рок-группы «Гражданская оборона». Первая часть дилогии «Долгая счастливая жизнь/Реанимация».
Обложкой альбома послужила картина хорватского художника наивного искусства Ивана Веченая «Napustino Grobje».

Источник: Wikipedia (', '2004-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 13

-- Альбом/сингл: Легенды русского рока. ГРАЖДАНСКАЯ ОБОРОНА
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Легенды русского рока. ГРАЖДАНСКАЯ ОБОРОНА', 'MPREb_nSwzrGzqCAF.webp', v_artist_id, '', '2022-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 28

-- Альбом/сингл: The Best, Pt. 1
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('The Best, Pt. 1', 'MPREb_UyLxpApOUlO.webp', v_artist_id, '', '2017-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 25
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Поганая молодёжь', 149, 'AifwyuGT-eI.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Про дурачка', 267, 'EjiVLwxV5fw.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Репетиции в Ленинграде
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Репетиции в Ленинграде', 'MPREb_Riwyl4Ljbx8.webp', v_artist_id, '', '2008-01-01', 'EP')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 6
-- Артист: Ramones
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Ramones', 'Ramones — американская панк-рок-группа, образованная в Нью-Йорке в 1974 году. Одни из самых первых исполнителей панк-рока, оказавшие влияние как в целом на этот жанр, так и на многие другие течения альтернативного рока.
Название группа получила по взятым себе псевдонимам, заканчивающимся фамилией Рамон, хотя биологически участники не были родственниками. Происхождение фамилии Рамон связано, во-первых, с Полом Маккартни, который в ранние годы группы The Beatles выступал под псевдонимом Пол Рамон, во-вторых, с тем фактом, что в Нью-Йорке фамилия «Рамон» ассоциировалась с членами латиноамериканских банд.

Источник: Wikipedia (', 'UCSrA5JaXpR21z_E1FDA03RA_avatar.webp', 'UCSrA5JaXpR21z_E1FDA03RA_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Best of The EMI Years
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Best of The EMI Years', 'MPREb_Uai2dmLm5ya.webp', v_artist_id, '', '2002-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 18
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Poison Heart', 244, 'OfIfzVf8t6E.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Essential
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Essential', 'MPREb_tavA03vlyA4.webp', v_artist_id, '', '2007-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 20
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Poison Heart', 244, 'OfIfzVf8t6E.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Brain Drain
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Brain Drain', 'MPREb_WvdiDbTNCGY.webp', v_artist_id, 'Brain Drain — одиннадцатый студийный альбом американской панк-рок-группы Ramones, выпущенный 23 марта 1989 года. Последний альбом с участием басиста Ди Ди Рамона и первый за шесть лет с участием Марки Рамона с момента его ухода из группы после записи Subterranean Jungle. Также это последний альбом, записанный на лейбле Sire Records.
В автобиографии Lobotomy: Surviving the Ramones Ди Ди писал:
Brain Drain было тяжело записывать, потому что каждый решил вылить своё дерьмо на меня. Меня пугало нахождение рядом с чужими проблемами. Это отвлекало меня, я даже не закончил свою часть работы над альбомом. Каждый член группы был занят своими проблемами: с отношениями, деньгами или со психикой.

Источник: Wikipedia (', '1989-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 12
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Learn to Listen', 111, 'B3MUC5LY1D4.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Mondo Bizarro
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Mondo Bizarro', 'MPREb_sQDEJrXqmMw.webp', v_artist_id, 'Mondo Bizarro — двенадцатый студийный альбом американской панк-группы Ramones, выпущенный в 1992 году. Это первый альбом с участием нового бас-гитариста Си Джея Рамона, заменившего ушедшего Ди Ди Рамона. 10 августа 2004 года в Великобритании альбом был перевыпущен лейблом Captain Oi!, с бонус-треком «Spider-Man».

Источник: Wikipedia (', '1992-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 13
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Poison Heart', 244, 'OfIfzVf8t6E.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: All the Best
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('All the Best', 'MPREb_PZlzdvaGXhh.webp', v_artist_id, '', '2012-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 40
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Poison Heart', 244, 'OfIfzVf8t6E.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Road to Ruin
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Road to Ruin', 'MPREb_Phf4F3PaEdh.webp', v_artist_id, 'Road to Ruin [roud tu: ruin] — четвёртый студийный альбом американской панк-рок-группы Ramones, вышедший в 1978 году на звукозаписывающем лейбле Sire Records.

Источник: Wikipedia (', '1978-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 12

-- Альбом/сингл: Hey Ho Let''s Go
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Hey Ho Let''s Go', 'MPREb_JUFPAKEZQDb.webp', v_artist_id, '', '2017-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 2

-- Альбом/сингл: Spider-Man
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Spider-Man', 'MPREb_kDlzmzu0Dj6.webp', v_artist_id, '', '2019-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 3

-- Альбом/сингл: Gabba Gabba Hey
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Gabba Gabba Hey', 'MPREb_5MSz4oTqFn3.webp', v_artist_id, '', '2017-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 2

-- Альбом/сингл: Leathers from New York
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Leathers from New York', 'MPREb_4K0O6GOlKfK.webp', v_artist_id, '', '1997-01-01', 'EP')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 4

-- Альбом/сингл: S.L.U.G. (2018 Mix)
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('S.L.U.G. (2018 Mix)', 'MPREb_O50xbB8qdOb.webp', v_artist_id, '', '2018-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
-- Артист: Misfits
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Misfits', 'Misfits — американская панк-рок-группа, образованная Гленном Данцигом в 1977 году в городке Лоди, штат Нью-Джерси, стала основателем хоррор-панка, а также оказала большое влияние на хэви-метал и на рок в целом. Коллектив не раз распадался, менял состав участников; бессменный басист группы — Джерри Онли.
В 2016 году произошло воссоединение классического состава.

Источник: Wikipedia (', 'UCUAa1fv7JsIFuBUWtH0ZaOA_avatar.webp', 'UCUAa1fv7JsIFuBUWtH0ZaOA_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: American Psycho
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('American Psycho', 'MPREb_DaijbRMDgLU.webp', v_artist_id, 'American Psycho — третий студийный альбом группы The Misfits, выпущен 13 мая 1997 года, первый альбом группы после их возвращения на сцену, без участия духовного лидера Гленна Данцига. Один из основателей группы, бессменный басист Джерри Онли, спустя 12 лет после распада тяжело достиг соглашения с Данцигом и получил права на использование имени группы, логотипа и ранних записей.

Источник: Wikipedia (', '1997-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 17

-- Альбом/сингл: Famous Monsters
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Famous Monsters', 'MPREb_h9aRc0M4T0Q.webp', v_artist_id, 'Famous Monsters — пятый студийный альбом группы The Misfits, выпущенный в 1999 году.
Альбом отметился 138-й позицией в The Billboard 200 и включён журналом Kerrang! в список «40 лучших панк-альбомов 1977—2017».
Famous Monsters — второй альбом группы после их возрождения в 1995 году; он же стал и последним, в записи которого принимали участие новый фронтмен Майкл Грэйвс, гитарист Дойл и ударник Др. Чад, покинувшие The Misfits в 2000 году.
В звучании альбома присутствуют элементы ду-вопа и сайкобилли. На песню «Scream!» был снят видеоклип известным режиссёром фильмов ужасов Джорджем Ромеро, снявшим «Ночь живых мертвецов». Кроме того, группа засветилась в его фильме «Вышибала».
Рецензент российского издания журнала Classic Rock отметил, что альбом несколько слабоват в сравнении со своим предшественником, но так же отметил зажигательное панково-хардкоровое звучание, классические тексты-ужастики и отличный вокал Майкла Грэйвса.

Источник: Wikipedia (', '1999-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 18

-- Альбом/сингл: Static Age
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Static Age', 'MPREb_IJy7KFrCpcM.webp', v_artist_id, 'Static Age — студийный альбом американской хоррор-панк группы The Misfits. Был записан в 1978 году, но выпущен только в 1996 году.

Источник: Wikipedia (', '1996-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 17

-- Альбом/сингл: Collection
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Collection', 'MPREb_cBS8fcdMM3x.webp', v_artist_id, 'Collection I — сборник американской хоррор-панк группы The Misfits, выпущенный в 1986 году на CD, а в 1988 году — на виниле и кассетах. В 1995 году последовало продолжение Collection II, а в 1996 году оба сборника вошли в бокс-сет «The Misfits».

Источник: Wikipedia (', '1986-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 20

-- Альбом/сингл: Walk Among Us
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Walk Among Us', 'MPREb_GxgiXxNr5J0.webp', v_artist_id, 'Walk Among Us — первый лонгплей американской хоррор-панк-группы The Misfits, издан в 1982 году. Однако это второй подготовленный альбом и восьмой релиз.

Источник: Wikipedia (', '1982-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 13

-- Альбом/сингл: Collection 2
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Collection 2', 'MPREb_wZTuNDKzbzV.webp', v_artist_id, '', '1995-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 20

-- Альбом/сингл: Vampire Girl
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Vampire Girl', 'MPREb_8AGmQxmbZwc.webp', v_artist_id, '', '1992-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 2

-- Альбом/сингл: Friday The 13th
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Friday The 13th', 'MPREb_c3P1C98cgUH.webp', v_artist_id, '', '1997-01-01', 'EP')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 4

-- Альбом/сингл: Day the Earth Caught Fire
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Day the Earth Caught Fire', 'MPREb_iUazfhaMw44.webp', v_artist_id, '', '1993-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: Horror Xmas
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Horror Xmas', 'MPREb_5JmlfSgGdp7.webp', v_artist_id, '', '1989-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 3
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('You’re a Mean One, Mr. Grinch', 133, '0PrB5elT1XM.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Descending Angel / Science Fiction/Double Feature
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Descending Angel / Science Fiction/Double Feature', 'MPREb_WPX6yH18iB6.webp', v_artist_id, '', '1994-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 2
-- Артист: Noize MC
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Noize MC', 'Noize МС, последний альбом и видеоклип выпущены как Noi3e MC, — российский рэп‑рок‑исполнитель, автор песен и музыкант.

Источник: Wikipedia (', 'UCgzshmpXAc1T30PHQ3Yw2lw_avatar.webp', 'UCgzshmpXAc1T30PHQ3Yw2lw_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Царь горы
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Царь горы', 'MPREb_nzZ4CqIOTdu.webp', v_artist_id, '«Царь горы» — седьмой студийный альбом российского рэп-исполнителя Noize MC, релиз которого состоялся 16 декабря 2016 года.
На композиции «Make Some Noize», «Питерские крыши», «Грабли», «Чайлдфри» были выпущены видеоклипы.

Источник: Wikipedia (', '2016-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 13

-- Альбом/сингл: Чайлдфри
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Чайлдфри', 'MPREb_xDJFzLswyyO.webp', v_artist_id, '', '2017-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 7

-- Альбом/сингл: The Greatest Hits, Vol. 1
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('The Greatest Hits, Vol. 1', 'MPREb_VkrlZhJmwuE.webp', v_artist_id, 'The Greatest Hits Vol. 1 — дебютный студийный альбом российского рэп-рок-исполнителя Noize MC. В официальной продаже с 17 июня 2008 года. Презентация состоялась 30 мая 2008 в клубе «Город».

Источник: Wikipedia (', '2008-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 20
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('На работе (платят бабло)', 150, 'FUyNYUu3jhE.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: Розыгрыш (Из к/ф «Розыгрыш»)
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Розыгрыш (Из к/ф «Розыгрыш»)', 'MPREb_90Pi4wK5VQJ.webp', v_artist_id, '', '2009-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 15
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Суицид', 234, '6PL6taBk2qA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;

-- Альбом/сингл: The Greatest Hits, Vol. 2
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('The Greatest Hits, Vol. 2', 'MPREb_9vSQ4AThlEA.webp', v_artist_id, '', '2010-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 25

-- Альбом/сингл: Новый альбом
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Новый альбом', 'MPREb_77nwukDVGle.webp', v_artist_id, '«Новый альбом» — третий студийный альбом российского рэп-исполнителя Noize MC. Альбом выпущен эксклюзивно на ThankYou.ru 30 марта 2012 года, где каждый желающий может сказать «Спасибо» исполнителю в виде денежного взноса. В песне «Вселенная бесконечна?» использован семпл из песни «Empty» группы «The Cranberries» с альбома «No Need to Argue».

Источник: Wikipedia (', '2012-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 21

-- Альбом/сингл: Люди с автоматами
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Люди с автоматами', 'MPREb_ez3Jm6ueMjw.webp', v_artist_id, '', '2018-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: Почитай старших
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Почитай старших', 'MPREb_XJa8Lvvj0Sx.webp', v_artist_id, '', '2019-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: 26.04
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('26.04', 'MPREb_ySNWttPHrdX.webp', v_artist_id, '', '2020-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: Век-волкодав (За гремучую доблесть грядущих веков)
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Век-волкодав (За гремучую доблесть грядущих веков)', 'MPREb_dhfDZeNdzdF.webp', v_artist_id, '', '2021-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: Обломки чувств
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Обломки чувств', 'MPREb_tpRAxTR1AbK.webp', v_artist_id, '', '2025-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1
INSERT INTO track (title, duration_s, file_url, description)
VALUES ('Обломки чувств', 211, 'yEpS8egTymA.webm', '')
RETURNING track_id INTO v_track_id;
INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;
INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;
-- Артист: Amon Amarth
INSERT INTO artist (artist_name, description, avatar_url, header_url)
VALUES ('Amon Amarth', 'Amon Amarth — шведская метал-группа, играющая в жанре мелодичный дэт-метал. Название группы значит «Роковая гора» на синдарине, эльфийском языке, и отсылает к горе Ородруин из романа Дж. Р. Р. Толкина «Властелина Колец». Основной тематикой группы является германо-скандинавская мифология. Музыку Amon Amarth составляет сочетание гроулинга с использованием двойной бас-бочки, «тяжелыми» мелодиями и эпическими текстами.

Источник: Wikipedia (', 'UCS8-ccff2oGbG0Mx7UOuKYw_avatar.webp', 'UCS8-ccff2oGbG0Mx7UOuKYw_header.webp')
RETURNING artist_id INTO v_artist_id;

-- Альбом/сингл: Twilight Of The Thunder God
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Twilight Of The Thunder God', 'MPREb_hzr0NTV2jJ5.webp', v_artist_id, 'Twilight of the Thunder God — седьмой студийный альбом шведской группы Amon Amarth. Релиз альбома состоялся 17 сентября 2008 года.

Источник: Wikipedia (', '2008-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 10

-- Альбом/сингл: Versus the World (Bonus Edition)
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Versus the World (Bonus Edition)', 'MPREb_6svILusKhju.webp', v_artist_id, '', '2002-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 18

-- Альбом/сингл: Once Sent From The Golden Hall
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Once Sent From The Golden Hall', 'MPREb_iNahZ36i4uF.webp', v_artist_id, '', '1998-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 8

-- Альбом/сингл: The Avenger
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('The Avenger', 'MPREb_5YQihpFXslm.webp', v_artist_id, '', '1999-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 7

-- Альбом/сингл: The Crusher
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('The Crusher', 'MPREb_mVC5kJh0keo.webp', v_artist_id, '', '2001-01-01', 'Альбом')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 10

-- Альбом/сингл: Put Your Back Into The Oar
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Put Your Back Into The Oar', 'MPREb_jQ2a9xUPQYU.webp', v_artist_id, '', '2002-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: Masters of War
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Masters of War', 'MPREb_TyF7TV2Qvei.webp', v_artist_id, '', '2004-01-01', 'Сингл')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 1

-- Альбом/сингл: Under the Influence - EP
INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)
VALUES ('Under the Influence - EP', 'MPREb_kRCZE6gz3mq.webp', v_artist_id, '', '2004-01-01', 'EP')
RETURNING album_id INTO v_album_id;
-- Найдено треков: 4
END $$;