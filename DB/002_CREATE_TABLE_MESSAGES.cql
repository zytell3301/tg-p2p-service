USE tg_p2p_messages;

CREATE TABLE IF NOT EXISTS messages
(
    left_side  UUID,
    contact_id UUID,
    right_side UUID,
    text       VARCHAR, // Other fields like file or voice and etc. will be added when tg-fs implemented
    sender     BOOLEAN, // value 0 or false indicates that the sender is left side and true the opposite
    sent_at    TIMESTAMP,
    message_id UUID,    // This field is used to distinguish several messages and their logs in exact same nano-sec
    -- (Nearly impossible but we can't waiver that little possibility)
    PRIMARY KEY ( (left_side, contact_id), sent_at, message_id)
) WITH CLUSTERING ORDER BY (sent_at DESC);