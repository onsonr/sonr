import { arrayBufferDecode } from "./types/utils";


export default class Webauthn {
    private origin: string;

    constructor(origin: string) {
        this.origin = origin;
    }

    /**
     * This function authenticates a WebAuthn credential and returns it if successful, otherwise it returns
     * null.
     * @param {string} credentialRequestOptions - A string representing the options for a WebAuthn
     * credential request, which includes information such as the challenge and allowed credentials. This
     * parameter will be parsed into a JSON object within the function.
     * @returns a Promise that resolves to a PublicKeyCredential object or null.
     */
    async authenticateWebAuthnCredential(
        credentialRequestOptions: string
    ): Promise<PublicKeyCredential | null> {
        const abortController = new AbortController();
        try {
            const options = JSON.parse(credentialRequestOptions);
            options.publicKey.challenge = arrayBufferDecode(options.publicKey.challenge);

            if (options.publicKey.allowCredentials) {
                for (let i = 0; i < options.publicKey.allowCredentials.length; i++) {
                    options.publicKey.allowCredentials[i].id = arrayBufferDecode(
                        options.publicKey.allowCredentials[i].id
                    );
                }
            }
            const credential = await navigator.credentials.get(options);
            console.log('WebAuthn authentication successful:', credential);
            return credential as PublicKeyCredential;
        } catch (error) {
            console.error('WebAuthn authentication failed:', error);
            return null;
        }
    }

    /**
     * This function generates a WebAuthn credential using the provided options and returns it as a
     * PublicKeyCredential object.
     * @param {string} credentialCreationOptions - A string representing the options for creating a new
     * WebAuthn credential. This string needs to be parsed into a JSON object before it can be used to
     * generate the credential.
     * @returns The function `generateWebAuthnCredential` returns a Promise that resolves to a
     * `PublicKeyCredential` object if the WebAuthn registration is successful, or `null` if it fails.
     */
    async generateWebAuthnCredential(
        credentialCreationOptions: string,
    ): Promise<PublicKeyCredential | null> {

        try {
            // Generate WebAuthn credential creation options
            const options = JSON.parse(credentialCreationOptions);
            const challenge = arrayBufferDecode(
                options.publicKey.challenge
            );
            const userId = arrayBufferDecode(
                options.publicKey.user.id
            );
            options.publicKey.challenge = challenge;
            options.publicKey.user.id = userId;
            if (options.publicKey.excludeCredentials) {
                for (
                    var i = 0;
                    i < options.publicKey.excludeCredentials.length;
                    i++
                ) {
                    options.publicKey.excludeCredentials[i].id =
                        Uint8Array.from(
                            options.publicKey.excludeCredentials[i]
                                .id as string,
                            (c) => c.charCodeAt(0)
                        );
                }
            }
            // Request the creation of a new WebAuthn credential using the generated options
            const credential = await navigator.credentials.create(options);
            console.log('WebAuthn registration successful:', credential);
            return credential as PublicKeyCredential;
        } catch (error) {
            console.error('WebAuthn registration failed:', error);
            return null;
        }
    }
}
