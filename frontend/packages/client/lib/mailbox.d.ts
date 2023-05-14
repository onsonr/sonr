import { SendMessageResponse } from "./types";
import { MailboxMessage } from "./types/mailbox";
export default class Mailbox {
    private token;
    constructor(token: string);
    /**
     * This is an asynchronous function that reads mailbox messages for a given address using Axios.
     * @param {string} address - The `address` parameter is a string representing the email address of the
     * mailbox from which the messages are to be read.
     * @returns The function `readMessages` returns a Promise that resolves to an array of `MailboxMessage`
     * objects.
     */
    readMessages(address: string): Promise<MailboxMessage[]>;
    /**
     * This is an asynchronous function that sends a message to a specified recipient using Axios and
     * returns a Promise containing the response data.
     * @param {string} to - The recipient of the message, specified as a string.
     * @param {string} message - The message parameter is a string that represents the content of the
     * message that will be sent to the recipient.
     * @returns The `sendMessage` function returns a Promise that resolves to a `SendMessageResponse`
     * object.
     */
    sendMessage(to: string, message: string): Promise<SendMessageResponse>;
}
