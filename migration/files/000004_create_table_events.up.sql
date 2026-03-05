CREATE TABLE IF NOT EXISTS public.events (
    id                uuid         DEFAULT uuid_generate_v4() NOT NULL,
    organizer_user_id uuid                                    NOT NULL,
    category_id       uuid                                    NOT NULL,
    title             varchar(255)                            NOT NULL,
    slug              varchar(255)                            NOT NULL,
    description       text,
    location          varchar(255),
    start_date        timestamptz                             NOT NULL,
    end_date          timestamptz                             NOT NULL,
    status            integer                                 NOT NULL DEFAULT 1,
    created_at        timestamptz  DEFAULT CURRENT_TIMESTAMP,
    updated_at        timestamptz  DEFAULT CURRENT_TIMESTAMP,
    deleted_at        timestamptz,
    created_by        uuid,
    updated_by        uuid,
    deleted_by        uuid,
    CONSTRAINT events_pkey                 PRIMARY KEY (id),
    CONSTRAINT uq_events_slug              UNIQUE (slug),
    CONSTRAINT fk_events_organizer_user_id FOREIGN KEY (organizer_user_id) REFERENCES public.users (id)      ON DELETE CASCADE,
    CONSTRAINT fk_events_category_id       FOREIGN KEY (category_id)       REFERENCES public.categories (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_events_deleted_at        ON public.events USING btree (deleted_at);
CREATE INDEX IF NOT EXISTS idx_events_status_start_date ON public.events USING btree (status, start_date);
CREATE INDEX IF NOT EXISTS idx_events_organizer_status  ON public.events USING btree (organizer_user_id, status);
CREATE INDEX IF NOT EXISTS idx_events_category_status   ON public.events USING btree (category_id, status);
