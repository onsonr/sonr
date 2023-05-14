import { DidDocument } from "./did";

export interface SonrRegisterProps {
    alias: string;
    onCredentialSet: (credential: PublicKeyCredential) => Promise<void>;
    onRegisterComplete: (did: string, didDocument: DidDocument, jwt: string) => Promise<void>;
}

export interface SonrLoginProps {
    alias: string;
    onCredentialSet: (credential: PublicKeyCredential) => Promise<void>;
    onLoginComplete: (did: string, didDocument: DidDocument, jwt: string) => Promise<void>;
}
