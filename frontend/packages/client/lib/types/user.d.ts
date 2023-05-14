import { DidDocument } from "./did";
type User = {
    did: string;
    didDocument: DidDocument;
    username: string;
    address: string;
};
type Account = {
    address: string;
    name: string;
    did: string;
    coin_type: string;
    chain_id: string;
    public_key: string;
    type: string;
};
export declare function derivePrivateKeyFromWebAuthnCredentialAndPin(pin: string): Promise<CryptoKey>;
export type { User, Account };
