import { DidDocument } from "./did";
interface User {
    did: string;
    didDocument: DidDocument;
    username: string;
    address: string;
}
interface Account {
    address: string;
    name: string;
    did: string;
    coin_type: string;
    chain_id: string;
    public_key: string;
    type: string;
}
export declare function derivePrivateKeyFromWebAuthnCredentialAndPin(pin: string): Promise<CryptoKey>;
export type { User, Account };
