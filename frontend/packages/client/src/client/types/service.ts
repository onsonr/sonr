// ServiceRecordTypes.ts
interface KeyValuePair {
    key: string;
    value: string;
}

interface ServiceRecord {
    id: string;
    controller: string;
    type: string;
    origin: string;
    name: string;
    serviceEndpoints: KeyValuePair[];
    metadata: KeyValuePair[];
}

export type { KeyValuePair, ServiceRecord };
