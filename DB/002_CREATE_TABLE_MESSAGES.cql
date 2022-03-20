USE tg_p2p_messages;

CREATE TABLE IF NOT EXISTS messages
(
    left_side  UUID,
    contact_id UUID,
    right_side UUID,
    text       UUID, // Other fields like file or voice and etc. will be added when tg-fs implemented
    sent_at    TIMESTAMP,
    PRIMARY KEY ( (left_side, contact_id), sent_at)
) WITH CLUSTERING ORDER BY (sent_at DESC);