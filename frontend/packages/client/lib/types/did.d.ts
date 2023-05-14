interface DidDocument {
    context: string[];
    id: string;
    controller: string[];
    verificationMethod: VerificationMethod[];
    authentication: string[];
    assertionMethod: string[];
    capabilityInvocation: string[];
    capabilityDelegation: string[];
    keyAgreement: string[];
    alsoKnownAs: string[];
    metadata: string;
    owner: string;
}
interface VerificationMethod {
    id: string;
    type: string;
    controller: string;
    publicKeyJwk?: string;
    publicKeyMultibase?: string;
    blockchainAccountId?: string;
    metadata?: string;
}
interface VerificationRelationship {
    verificationMethod?: VerificationMethod;
    reference: string;
    type: string;
}
export type { DidDocument, VerificationMethod, VerificationRelationship };
