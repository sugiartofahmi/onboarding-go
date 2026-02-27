CREATE TABLE IF NOT EXISTS public.event_registrations (
    id              uuid        DEFAULT uuid_generate_v4() NOT NULL,
    user_id         uuid                                   NOT NULL,
    event_ticket_id uuid                                   NOT NULL,
    status          integer                                NOT NULL,
    created_at      timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at      timestamptz DEFAULT CURRENT_TIMESTAMP,
    deleted_at      timestamptz,
    created_by      uuid,
    updated_by      uuid,
    deleted_by      uuid,
    CONSTRAINT event_registrations_pkey                  PRIMARY KEY (id),
    CONSTRAINT uq_event_registrations_user_ticket        UNIQUE (user_id, event_ticket_id),
    CONSTRAINT fk_event_registrations_user_id            FOREIGN KEY (user_id)         REFERENCES public.users (id)         ON DELETE CASCADE,
    CONSTRAINT fk_event_registrations_event_ticket_id    FOREIGN KEY (event_ticket_id) REFERENCES public.event_tickets (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_event_registrations_deleted_at    ON public.event_registrations USING btree (deleted_at);
CREATE INDEX IF NOT EXISTS idx_event_registrations_user_status   ON public.event_registrations USING btree (user_id, status);
CREATE INDEX IF NOT EXISTS idx_event_registrations_ticket_status ON public.event_registrations USING btree (event_ticket_id, status);
