USE tg_p2p_messages;

CREATE TABLE IF NOT EXISTS messages_log
(
    left_side  UUID,
    contact_id UUID,
    right_side UUID,
    text       VARCHAR,   // Other fields like file or voice and etc. will be added when tg-fs implemented
    sender     BOOLEAN,   // value 0 or false indicates that the sender is left side and true the opposite
    sent_at    TIMESTAMP,
    action     VARCHAR,   // The action that resulted into log like editing a message or deleting it
    log_time   TIMESTAMP, // Maybe several edits happen to a message that will result in several logs.
    // This field will distinguish does logs
    message_id UUID,
    PRIMARY KEY ( (left_side, contact_id), sent_at, log_time, message_id)
) WITH CLUSTERING ORDER BY (sent_at DESC) AND
        CLUSTERING ORDER BY (log_time DESC);