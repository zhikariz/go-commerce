BEGIN;

CREATE TABLE IF NOT EXISTS transactions(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL,
    qty INT NOT NULL,
    user_id UUID NOT NULL,
    discount NUMERIC(10,2),
    is_paid BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

COMMIT;