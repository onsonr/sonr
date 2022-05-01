import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "sonrio.sonr.object";
/** ObjectFieldType is the type of the field */
export declare enum ObjectFieldType {
    /** OBJECT_FIELD_TYPE_UNSPECIFIED - ObjectFieldTypeUnspecified is the default value */
    OBJECT_FIELD_TYPE_UNSPECIFIED = 0,
    /** OBJECT_FIELD_TYPE_STRING - ObjectFieldTypeString is a string or text field */
    OBJECT_FIELD_TYPE_STRING = 1,
    /** OBJECT_FIELD_TYPE_NUMBER - ObjectFieldTypeInt is an integer */
    OBJECT_FIELD_TYPE_NUMBER = 2,
    /** OBJECT_FIELD_TYPE_BOOL - ObjectFieldTypeBool is a boolean */
    OBJECT_FIELD_TYPE_BOOL = 3,
    /** OBJECT_FIELD_TYPE_ARRAY - ObjectFieldTypeArray is a list of values */
    OBJECT_FIELD_TYPE_ARRAY = 4,
    /** OBJECT_FIELD_TYPE_TIMESTAMP - ObjectFieldTypeDateTime is a datetime */
    OBJECT_FIELD_TYPE_TIMESTAMP = 5,
    /** OBJECT_FIELD_TYPE_GEOPOINT - ObjectFieldTypeGeopoint is a geopoint */
    OBJECT_FIELD_TYPE_GEOPOINT = 6,
    /** OBJECT_FIELD_TYPE_BLOB - ObjectFieldTypeBlob is a blob of data */
    OBJECT_FIELD_TYPE_BLOB = 7,
    /** OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS - ObjectFieldTypeETU is a pointer to an Ethereum account address. */
    OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS = 8,
    UNRECOGNIZED = -1
}
export declare function objectFieldTypeFromJSON(object: any): ObjectFieldType;
export declare function objectFieldTypeToJSON(object: ObjectFieldType): string;
/** ObjectDoc is a document for an Object stored in the graph. */
export interface ObjectDoc {
    /** Label is human-readable name of the bucket. */
    label: string;
    /** Description is a human-readable description of the bucket. */
    description: string;
    /** Did is the identifier of the object. */
    did: string;
    /** Bucket is the did of the bucket that contains this object. */
    bucketDid: string;
    /** Fields are the fields associated with the object. */
    fields: {
        [key: string]: ObjectField;
    };
}
export interface ObjectDoc_FieldsEntry {
    key: string;
    value: ObjectField | undefined;
}
/** ObjectField is a field of an Object. */
export interface ObjectField {
    /** Label is the name of the field */
    label: string;
    /** Type is the type of the field */
    type: ObjectFieldType;
    /** Did is the identifier of the field. */
    did: string;
    /** String is the value of the field */
    stringValue: ObjectFieldText | undefined;
    /** Number is the value of the field */
    numberValue: ObjectFieldNumber | undefined;
    /** Float is the value of the field */
    boolValue: ObjectFieldBool | undefined;
    /** Array is the value of the field */
    arrayValue: ObjectFieldArray | undefined;
    /** Time is defined by milliseconds since epoch. */
    timeStampValue: ObjectFieldTime | undefined;
    /** Geopoint is the value of the field */
    geopointValue: ObjectFieldGeopoint | undefined;
    /** Blob is the value of the field */
    blobValue: ObjectFieldBlob | undefined;
    /** Blockchain Address is the value of the field */
    blockchainAddrValue: ObjectFieldBlockchainAddress | undefined;
    /** Metadata is additional info about the field */
    metadata: {
        [key: string]: string;
    };
}
export interface ObjectField_MetadataEntry {
    key: string;
    value: string;
}
/** ObjectFieldArray is an array of ObjectFields to be stored in the graph object. */
export interface ObjectFieldArray {
    /** Label is the name of the field */
    label: string;
    /** Type is the type of the field */
    type: ObjectFieldType;
    /** Did is the identifier of the field. */
    did: string;
    /** Entries are the values of the field */
    items: ObjectField[];
}
/** ObjectFieldText is a text field of an Object. */
export interface ObjectFieldText {
    /** Label is the name of the field */
    label: string;
    /** Did is the identifier of the field. */
    did: string;
    /** Value is the value of the field */
    value: string;
    /** Metadata is additional info about the field */
    metadata: {
        [key: string]: string;
    };
}
export interface ObjectFieldText_MetadataEntry {
    key: string;
    value: string;
}
/** ObjectFieldNumber is a number field of an Object. */
export interface ObjectFieldNumber {
    /** Label is the name of the field */
    label: string;
    /** Did is the identifier of the field. */
    did: string;
    /** Value is the value of the field */
    value: number;
    /** Metadata is additional info about the field */
    metadata: {
        [key: string]: string;
    };
}
export interface ObjectFieldNumber_MetadataEntry {
    key: string;
    value: string;
}
/** ObjectFieldBool is a boolean field of an Object. */
export interface ObjectFieldBool {
    /** Label is the name of the field */
    label: string;
    /** Did is the identifier of the field. */
    did: string;
    /** Value is the value of the field */
    value: boolean;
    /** Metadata is additional info about the field */
    metadata: {
        [key: string]: string;
    };
}
export interface ObjectFieldBool_MetadataEntry {
    key: string;
    value: string;
}
/** ObjectFieldTime is a time field of an Object. */
export interface ObjectFieldTime {
    /** Label is the name of the field */
    label: string;
    /** Did is the identifier of the field. */
    did: string;
    /** Value is the value of the field */
    value: number;
    /** Metadata is additional info about the field */
    metadata: {
        [key: string]: string;
    };
}
export interface ObjectFieldTime_MetadataEntry {
    key: string;
    value: string;
}
/** ObjectFieldGeopoint is a field of an Object for geopoints. */
export interface ObjectFieldGeopoint {
    /** Label is the name of the field */
    label: string;
    /** Type is the type of the field */
    type: ObjectFieldType;
    /** Did is the identifier of the field. */
    did: string;
    /** Latitude is the geo-latitude of the point. */
    latitude: number;
    /** Longitude is the geo-longitude of the field */
    longitude: number;
    /** Metadata is additional info about the field */
    metadata: {
        [key: string]: string;
    };
}
export interface ObjectFieldGeopoint_MetadataEntry {
    key: string;
    value: string;
}
/** ObjectFieldBlob is a field of an Object for blobs. */
export interface ObjectFieldBlob {
    /** Label is the name of the field */
    label: string;
    /** Did is the identifier of the field. */
    did: string;
    /** Value is the value of the field */
    value: Uint8Array;
    /** Metadata is additional info about the field */
    metadata: {
        [key: string]: string;
    };
}
export interface ObjectFieldBlob_MetadataEntry {
    key: string;
    value: string;
}
/** ObjectFieldBlockchainAddress is a field of an Object for blockchain addresses. */
export interface ObjectFieldBlockchainAddress {
    /** Label is the name of the field */
    label: string;
    /** Did is the identifier of the field. */
    did: string;
    /** Value is the value of the field */
    value: string;
    /** Metadata is additional info about the field */
    metadata: {
        [key: string]: string;
    };
}
export interface ObjectFieldBlockchainAddress_MetadataEntry {
    key: string;
    value: string;
}
export declare const ObjectDoc: {
    encode(message: ObjectDoc, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectDoc;
    fromJSON(object: any): ObjectDoc;
    toJSON(message: ObjectDoc): unknown;
    fromPartial(object: DeepPartial<ObjectDoc>): ObjectDoc;
};
export declare const ObjectDoc_FieldsEntry: {
    encode(message: ObjectDoc_FieldsEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectDoc_FieldsEntry;
    fromJSON(object: any): ObjectDoc_FieldsEntry;
    toJSON(message: ObjectDoc_FieldsEntry): unknown;
    fromPartial(object: DeepPartial<ObjectDoc_FieldsEntry>): ObjectDoc_FieldsEntry;
};
export declare const ObjectField: {
    encode(message: ObjectField, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectField;
    fromJSON(object: any): ObjectField;
    toJSON(message: ObjectField): unknown;
    fromPartial(object: DeepPartial<ObjectField>): ObjectField;
};
export declare const ObjectField_MetadataEntry: {
    encode(message: ObjectField_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectField_MetadataEntry;
    fromJSON(object: any): ObjectField_MetadataEntry;
    toJSON(message: ObjectField_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<ObjectField_MetadataEntry>): ObjectField_MetadataEntry;
};
export declare const ObjectFieldArray: {
    encode(message: ObjectFieldArray, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldArray;
    fromJSON(object: any): ObjectFieldArray;
    toJSON(message: ObjectFieldArray): unknown;
    fromPartial(object: DeepPartial<ObjectFieldArray>): ObjectFieldArray;
};
export declare const ObjectFieldText: {
    encode(message: ObjectFieldText, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldText;
    fromJSON(object: any): ObjectFieldText;
    toJSON(message: ObjectFieldText): unknown;
    fromPartial(object: DeepPartial<ObjectFieldText>): ObjectFieldText;
};
export declare const ObjectFieldText_MetadataEntry: {
    encode(message: ObjectFieldText_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldText_MetadataEntry;
    fromJSON(object: any): ObjectFieldText_MetadataEntry;
    toJSON(message: ObjectFieldText_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<ObjectFieldText_MetadataEntry>): ObjectFieldText_MetadataEntry;
};
export declare const ObjectFieldNumber: {
    encode(message: ObjectFieldNumber, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldNumber;
    fromJSON(object: any): ObjectFieldNumber;
    toJSON(message: ObjectFieldNumber): unknown;
    fromPartial(object: DeepPartial<ObjectFieldNumber>): ObjectFieldNumber;
};
export declare const ObjectFieldNumber_MetadataEntry: {
    encode(message: ObjectFieldNumber_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldNumber_MetadataEntry;
    fromJSON(object: any): ObjectFieldNumber_MetadataEntry;
    toJSON(message: ObjectFieldNumber_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<ObjectFieldNumber_MetadataEntry>): ObjectFieldNumber_MetadataEntry;
};
export declare const ObjectFieldBool: {
    encode(message: ObjectFieldBool, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldBool;
    fromJSON(object: any): ObjectFieldBool;
    toJSON(message: ObjectFieldBool): unknown;
    fromPartial(object: DeepPartial<ObjectFieldBool>): ObjectFieldBool;
};
export declare const ObjectFieldBool_MetadataEntry: {
    encode(message: ObjectFieldBool_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldBool_MetadataEntry;
    fromJSON(object: any): ObjectFieldBool_MetadataEntry;
    toJSON(message: ObjectFieldBool_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<ObjectFieldBool_MetadataEntry>): ObjectFieldBool_MetadataEntry;
};
export declare const ObjectFieldTime: {
    encode(message: ObjectFieldTime, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldTime;
    fromJSON(object: any): ObjectFieldTime;
    toJSON(message: ObjectFieldTime): unknown;
    fromPartial(object: DeepPartial<ObjectFieldTime>): ObjectFieldTime;
};
export declare const ObjectFieldTime_MetadataEntry: {
    encode(message: ObjectFieldTime_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldTime_MetadataEntry;
    fromJSON(object: any): ObjectFieldTime_MetadataEntry;
    toJSON(message: ObjectFieldTime_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<ObjectFieldTime_MetadataEntry>): ObjectFieldTime_MetadataEntry;
};
export declare const ObjectFieldGeopoint: {
    encode(message: ObjectFieldGeopoint, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldGeopoint;
    fromJSON(object: any): ObjectFieldGeopoint;
    toJSON(message: ObjectFieldGeopoint): unknown;
    fromPartial(object: DeepPartial<ObjectFieldGeopoint>): ObjectFieldGeopoint;
};
export declare const ObjectFieldGeopoint_MetadataEntry: {
    encode(message: ObjectFieldGeopoint_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldGeopoint_MetadataEntry;
    fromJSON(object: any): ObjectFieldGeopoint_MetadataEntry;
    toJSON(message: ObjectFieldGeopoint_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<ObjectFieldGeopoint_MetadataEntry>): ObjectFieldGeopoint_MetadataEntry;
};
export declare const ObjectFieldBlob: {
    encode(message: ObjectFieldBlob, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldBlob;
    fromJSON(object: any): ObjectFieldBlob;
    toJSON(message: ObjectFieldBlob): unknown;
    fromPartial(object: DeepPartial<ObjectFieldBlob>): ObjectFieldBlob;
};
export declare const ObjectFieldBlob_MetadataEntry: {
    encode(message: ObjectFieldBlob_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldBlob_MetadataEntry;
    fromJSON(object: any): ObjectFieldBlob_MetadataEntry;
    toJSON(message: ObjectFieldBlob_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<ObjectFieldBlob_MetadataEntry>): ObjectFieldBlob_MetadataEntry;
};
export declare const ObjectFieldBlockchainAddress: {
    encode(message: ObjectFieldBlockchainAddress, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldBlockchainAddress;
    fromJSON(object: any): ObjectFieldBlockchainAddress;
    toJSON(message: ObjectFieldBlockchainAddress): unknown;
    fromPartial(object: DeepPartial<ObjectFieldBlockchainAddress>): ObjectFieldBlockchainAddress;
};
export declare const ObjectFieldBlockchainAddress_MetadataEntry: {
    encode(message: ObjectFieldBlockchainAddress_MetadataEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ObjectFieldBlockchainAddress_MetadataEntry;
    fromJSON(object: any): ObjectFieldBlockchainAddress_MetadataEntry;
    toJSON(message: ObjectFieldBlockchainAddress_MetadataEntry): unknown;
    fromPartial(object: DeepPartial<ObjectFieldBlockchainAddress_MetadataEntry>): ObjectFieldBlockchainAddress_MetadataEntry;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
