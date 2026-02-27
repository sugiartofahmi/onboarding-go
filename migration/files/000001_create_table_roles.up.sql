CREATE TABLE IF NOT EXISTS public.roles (
    id         uuid        DEFAULT uuid_generate_v4() NOT NULL,
    name       varchar(255)                           NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamptz,
    created_by uuid,
    updated_by uuid,
    deleted_by uuid,
    CONSTRAINT roles_pkey    PRIMARY KEY (id),
    CONSTRAINT uq_roles_name UNIQUE (name)
);

CREATE INDEX IF NOT EXISTS idx_roles_deleted_at ON public.roles USING btree (deleted_at);
