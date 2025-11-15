CREATE TABLE "ticket" (
    ticket_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('Открыто', 'Закрыто')),
    category TEXT NOT NULL CHECK (category IN ('Баг', 'Предложение', 'Жалоба')),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    rating INTEGER CHECK (rating BETWEEN 1 AND 5),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP

);

    COMMENT ON TABLE "ticket" IS 'Таблица хранения обращений пользователей';