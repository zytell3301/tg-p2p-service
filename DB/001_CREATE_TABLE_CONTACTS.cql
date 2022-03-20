USE tg_p2p_messages;

CREATE TABLE IF NOT EXISTS contacts
(
    left_side  UUID,
    contact_id UUID,
    right_side UUID,
    PRIMARY KEY ( left_side, contact_id )
)