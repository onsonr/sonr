interface MailboxMessage {
    id?: string;
    type?: string;
    content?: string;
    sender?: string;
    receiver?: string;
    coinType?: string;
}
export type { MailboxMessage };
