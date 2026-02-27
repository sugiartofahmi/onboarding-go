CREATE TABLE IF NOT EXISTS public.users (
    id         uuid         DEFAULT uuid_generate_v4() NOT NULL,
    role_id    uuid                                    NOT NULL,
    name       varchar(255)                            NOT NULL,
    email      varchar(255)                            NOT NULL,
    password   varchar(255)                            NOT NULL,
    is_active  boolean      DEFAULT TRUE               NOT NULL,
    created_at timestamptz  DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz  DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamptz,
    created_by uuid,
    updated_by uuid,
    deleted_by uuid,
    CONSTRAINT users_pkey      PRIMARY KEY (id),
    CONSTRAINT uq_users_email  UNIQUE (email),
    CONSTRAINT fk_users_role_id FOREIGN KEY (role_id) REFERENCES public.roles (id)
);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON public.users USING btree (deleted_at);
CREATE INDEX IF NOT EXISTS idx_users_role_id    ON public.users USING btree (role_id);
