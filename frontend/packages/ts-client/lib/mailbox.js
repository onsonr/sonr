"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const axios_1 = __importDefault(require("axios"));
const utils_1 = require("./types/utils");
class Mailbox {
    constructor(token) {
        this.token = token;
    }
    /**
     * This is an asynchronous function that reads mailbox messages for a given address using Axios.
     * @param {string} address - The `address` parameter is a string representing the email address of the
     * mailbox from which the messages are to be read.
     * @returns The function `readMessages` returns a Promise that resolves to an array of `MailboxMessage`
     * objects.
     */
    async readMessages(address) {
        let api = axios_1.default.create({
            baseURL: (0, utils_1.baseUrl)(),
            headers: { 'Authorization': `Bearer ${this.token}` },
        });
        try {
            const response = await api.get(`/mailbox/${address}/read`);
            return response.data;
        }
        catch (error) {
            console.error(`Error fetching mailbox messages: ${error}`);
            throw error;
        }
    }
    /**
     * This is an asynchronous function that sends a message to a specified recipient using Axios and
     * returns a Promise containing the response data.
     * @param {string} to - The recipient of the message, specified as a string.
     * @param {string} message - The message parameter is a string that represents the content of the
     * message that will be sent to the recipient.
     * @returns The `sendMessage` function returns a Promise that resolves to a `SendMessageResponse`
     * object.
     */
    async sendMessage(to, message) {
        let api = axios_1.default.create({
            baseURL: (0, utils_1.baseUrl)(),
            headers: { 'Authorization': `Bearer ${this.token}` },
        });
        try {
            const resp = await api.post(`/mailbox/${to}/send`, { message: message });
            return resp.data;
        }
        catch (error) {
            console.error(`Error sending message: ${error}`);
            throw error;
        }
    }
}
exports.default = Mailbox;
