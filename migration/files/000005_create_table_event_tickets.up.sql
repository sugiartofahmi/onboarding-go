CREATE TABLE IF NOT EXISTS public.event_tickets (
    id               uuid           DEFAULT uuid_generate_v4() NOT NULL,
    event_id         uuid                                      NOT NULL,
    type             integer                                   NOT NULL,
    price            decimal(12, 2)                            NOT NULL,
    quota            integer                                   NOT NULL,
    registered_count integer        DEFAULT 0                  NOT NULL,
    created_at       timestamptz    DEFAULT CURRENT_TIMESTAMP,
    updated_at       timestamptz    DEFAULT CURRENT_TIMESTAMP,
    deleted_at       timestamptz,
    created_by       uuid,
    updated_by       uuid,
    deleted_by       uuid,
    CONSTRAINT event_tickets_pkey          PRIMARY KEY (id),
    CONSTRAINT fk_event_tickets_event_id   FOREIGN KEY (event_id) REFERENCES public.events (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_event_tickets_deleted_at ON public.event_tickets USING btree (deleted_at);
CREATE INDEX IF NOT EXISTS idx_event_tickets_event_id   ON public.event_tickets USING btree (event_id);
