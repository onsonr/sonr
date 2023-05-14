import { KeyValuePair } from ".";
type DidDocument = {
    context: string[];
    id: string;
    controller?: string[];
    verification_method?: VerificationMethod[];
    authentication?: string[];
    assertion_method?: string[];
    capability_invocation?: string[];
    capability_delegation?: string[];
    key_agreement?: string[];
    service?: Service[];
    also_known_as?: string[];
    metadata?: KeyValuePair[];
};
type VerificationMethod = {
    id: string;
    type: string;
    controller: string;
    public_key_jwk?: KeyValuePair[];
    public_key_multibase?: string;
    blockchain_account_id?: string;
    metadata: KeyValuePair[];
};
type Service = {
    id: string;
    controller: string;
    type: string;
    origin: string;
    name: string;
    service_endpoints?: KeyValuePair[];
    metadata?: KeyValuePair[];
};
type VerificationRelationship = {
    verification_method: VerificationMethod;
    reference: string;
    type: string;
};
export type { DidDocument, VerificationMethod, Service, VerificationRelationship };
