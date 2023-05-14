interface MailboxMessage {
    // The message id
    id?: string;

    // The message type
    type?: string;

    // The message content
    content?: string;

    // The message sender
    sender?: string;

    // The message receiver
    receiver?: string;

    // CoinType Name is for the blockchain network name
    coinType?: string;
}

export type { MailboxMessage };
