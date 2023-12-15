--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4
-- Dumped by pg_dump version 15.4

-- Started on 2023-12-15 17:30:28

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3340 (class 0 OID 16779)
-- Dependencies: 214
-- Data for Name: genres; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.genres OVERRIDING SYSTEM VALUE VALUES (1, 'Comedy', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.genres OVERRIDING SYSTEM VALUE VALUES (2, 'Sci-Fi', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.genres OVERRIDING SYSTEM VALUE VALUES (3, 'Horror', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.genres OVERRIDING SYSTEM VALUE VALUES (4, 'Romance', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.genres OVERRIDING SYSTEM VALUE VALUES (5, 'Action', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.genres OVERRIDING SYSTEM VALUE VALUES (6, 'Thriller', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.genres OVERRIDING SYSTEM VALUE VALUES (7, 'Drama', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.genres OVERRIDING SYSTEM VALUE VALUES (8, 'Mystery', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.genres OVERRIDING SYSTEM VALUE VALUES (9, 'Crime', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.genres OVERRIDING SYSTEM VALUE VALUES (10, 'Animation', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.genres OVERRIDING SYSTEM VALUE VALUES (11, 'Adventure', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.genres OVERRIDING SYSTEM VALUE VALUES (12, 'Fantasy', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.genres OVERRIDING SYSTEM VALUE VALUES (13, 'Superhero', '2022-09-23 00:00:00', '2022-09-23 00:00:00');


--
-- TOC entry 3342 (class 0 OID 16783)
-- Dependencies: 216
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.movies OVERRIDING SYSTEM VALUE VALUES (1, 'Highlander', '1986-03-07', 116, 'R', 'He fought his first battle on the Scottish Highlands in 1536. He will fight his greatest battle on the streets of New York City in 1986. His name is Connor MacLeod. He is immortal.', '/8Z8dptJEypuLoOQro1WugD855YE.jpg', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.movies OVERRIDING SYSTEM VALUE VALUES (3, 'The Godfather', '1972-03-24', 175, '18A', 'The aging patriarch of an organized crime dynasty in postwar New York City transfers control of his clandestine empire to his reluctant youngest son.', '/3bhkrj58Vtu7enYsRolD1fZdja1.jpg', '2022-09-23 00:00:00', '2022-09-23 00:00:00');
INSERT INTO public.movies OVERRIDING SYSTEM VALUE VALUES (2, 'Raiders of the Lost Ark', '1981-06-12', 115, 'PG13', 'Archaeology professor Indiana Jones ventures to seize a biblical artefact known as the Ark of the Covenant. While doing so, he puts up a fight against Renee and a troop of Nazis.', '/ceG9VzoRAVGwivFU403Wc3AHRys.jpg', '2022-09-23 00:00:00', '2023-11-13 15:35:52.375921');
INSERT INTO public.movies OVERRIDING SYSTEM VALUE VALUES (4, 'Avengers', '2020-09-11', 120, 'PG13', 'Los Vengadores', '/rDzig50dj7VpLwJ7SThbamETK1G.jpg', '2023-11-13 19:40:41.434992', '2023-11-13 19:40:41.434992');
INSERT INTO public.movies OVERRIDING SYSTEM VALUE VALUES (5, 'Sabrina', '1920-11-11', 100, 'PG13', 'Brujas', '/z1oNjotUI7D06J4LWQFQzdIuPnf.jpg', '2023-11-13 20:16:16.879633', '2023-11-13 20:16:16.879633');


--
-- TOC entry 3343 (class 0 OID 16788)
-- Dependencies: 217
-- Data for Name: movies_genres; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.movies_genres OVERRIDING SYSTEM VALUE VALUES (1, 1, 5);
INSERT INTO public.movies_genres OVERRIDING SYSTEM VALUE VALUES (2, 1, 12);
INSERT INTO public.movies_genres OVERRIDING SYSTEM VALUE VALUES (5, 3, 9);
INSERT INTO public.movies_genres OVERRIDING SYSTEM VALUE VALUES (6, 3, 7);
INSERT INTO public.movies_genres OVERRIDING SYSTEM VALUE VALUES (7, 2, 5);
INSERT INTO public.movies_genres OVERRIDING SYSTEM VALUE VALUES (8, 2, 11);
INSERT INTO public.movies_genres OVERRIDING SYSTEM VALUE VALUES (9, 4, 5);
INSERT INTO public.movies_genres OVERRIDING SYSTEM VALUE VALUES (10, 4, 11);
INSERT INTO public.movies_genres OVERRIDING SYSTEM VALUE VALUES (11, 4, 2);
INSERT INTO public.movies_genres OVERRIDING SYSTEM VALUE VALUES (12, 4, 13);
INSERT INTO public.movies_genres OVERRIDING SYSTEM VALUE VALUES (13, 5, 11);
INSERT INTO public.movies_genres OVERRIDING SYSTEM VALUE VALUES (14, 5, 1);


--
-- TOC entry 3346 (class 0 OID 16793)
-- Dependencies: 220
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users OVERRIDING SYSTEM VALUE VALUES (1, 'Admin', 'User', 'admin@example.com', '$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOIEHpyh2DAMmso1wy', '2022-09-23 00:00:00', '2022-09-23 00:00:00');


--
-- TOC entry 3353 (class 0 OID 0)
-- Dependencies: 215
-- Name: genres_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.genres_id_seq', 13, true);


--
-- TOC entry 3354 (class 0 OID 0)
-- Dependencies: 218
-- Name: movies_genres_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.movies_genres_id_seq', 14, true);


--
-- TOC entry 3355 (class 0 OID 0)
-- Dependencies: 219
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.movies_id_seq', 5, true);


--
-- TOC entry 3356 (class 0 OID 0)
-- Dependencies: 221
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


-- Completed on 2023-12-15 17:30:28

--
-- PostgreSQL database dump complete
--

