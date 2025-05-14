-- Table: public.user

-- DROP TABLE IF EXISTS public."user";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public."user"
(   
    id UUID PRIMARY KEY NOT NULL
);

-- Table: public.channel

-- DROP TABLE IF EXISTS public.channel;

CREATE TABLE IF NOT EXISTS public.channel
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    can_subscribe_anyone BOOLEAN NOT NULL DEFAULT false,
    can_publish_anyone BOOLEAN NOT NULL DEFAULT false
);

-- Table: public.user_channel

-- DROP TABLE IF EXISTS public.user_channel;

CREATE TABLE IF NOT EXISTS public.user_channel
(
    user_id uuid NOT NULL,
    channel_id uuid NOT NULL,
    can_publish boolean NOT NULL DEFAULT false,
    CONSTRAINT user_chat_pkey PRIMARY KEY (user_id, channel_id),
    CONSTRAINT user_chat_channel_id_fkey FOREIGN KEY (channel_id)
        REFERENCES public.channel (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT user_chat_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
);