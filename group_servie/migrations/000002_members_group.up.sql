CREATE TABLE IF NOT EXISTS members_group (
    group_id UUID NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    soldier_id UUID NOT NULL UNIQUE
);
