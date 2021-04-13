--
-- PostgreSQL database dump
--

-- Dumped from database version 13.2
-- Dumped by pg_dump version 13.2

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: token; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.token (
    id bigint NOT NULL,
    id_user bigint NOT NULL,
    token character varying(255) NOT NULL
);


ALTER TABLE public.token OWNER TO postgres;

--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id bigint NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    username character varying(100) NOT NULL,
    status character varying(10) NOT NULL
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Data for Name: token; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.token (id, id_user, token) FROM stdin;
1618325974108	1618325974107	-
1618327058799	1618327058798	
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."user" (id, email, password, username, status) FROM stdin;
1618325974107	ronaka123@gmail.com	$2a$14$05ulnchDfDiIarLanKgiHuzQBTL3PCy3B4Pw5M6/1TXZXK6gZXub2	ronaka123	offline
1618327058798	dummy@tes.com	$2a$14$N5KOMInDk9XGwDSu5/ecH.NBVWCvYfn.QeNLySkJXHX7PLAG4DRDy	test123	offline
\.


--
-- Name: user email; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT email UNIQUE (email);


--
-- Name: user id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT id UNIQUE (id);


--
-- Name: token token_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.token
    ADD CONSTRAINT token_pkey PRIMARY KEY (id, id_user);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id, email, username);


--
-- Name: user username; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT username UNIQUE (username);


--
-- Name: token token_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.token
    ADD CONSTRAINT token_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id);


--
-- PostgreSQL database dump complete
--

