export default class Webauthn {
    private origin;
    constructor(origin: string);
    /**
     * This function authenticates a WebAuthn credential and returns it if successful, otherwise it returns
     * null.
     * @param {string} credentialRequestOptions - A string representing the options for a WebAuthn
     * credential request, which includes information such as the challenge and allowed credentials. This
     * parameter will be parsed into a JSON object within the function.
     * @returns a Promise that resolves to a PublicKeyCredential object or null.
     */
    authenticateWebAuthnCredential(credentialRequestOptions: string): Promise<PublicKeyCredential | null>;
    /**
     * This function generates a WebAuthn credential using the provided options and returns it as a
     * PublicKeyCredential object.
     * @param {string} credentialCreationOptions - A string representing the options for creating a new
     * WebAuthn credential. This string needs to be parsed into a JSON object before it can be used to
     * generate the credential.
     * @returns The function `generateWebAuthnCredential` returns a Promise that resolves to a
     * `PublicKeyCredential` object if the WebAuthn registration is successful, or `null` if it fails.
     */
    generateWebAuthnCredential(credentialCreationOptions: string): Promise<PublicKeyCredential | null>;
}
