CREATE TABLE IF NOT EXISTS public.categories (
    id         uuid         DEFAULT uuid_generate_v4() NOT NULL,
    name       varchar(255)                            NOT NULL,
    slug       varchar(255)                            NOT NULL,
    created_at timestamptz  DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz  DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamptz,
    created_by uuid,
    updated_by uuid,
    deleted_by uuid,
    CONSTRAINT categories_pkey     PRIMARY KEY (id),
    CONSTRAINT uq_categories_name  UNIQUE (name),
    CONSTRAINT uq_categories_slug  UNIQUE (slug)
);

CREATE INDEX IF NOT EXISTS idx_categories_deleted_at ON public.categories USING btree (deleted_at);
