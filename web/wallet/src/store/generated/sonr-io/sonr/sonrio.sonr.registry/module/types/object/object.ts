/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sonrio.sonr.object";

/** ObjectFieldType is the type of the field */
export enum ObjectFieldType {
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
  UNRECOGNIZED = -1,
}

export function objectFieldTypeFromJSON(object: any): ObjectFieldType {
  switch (object) {
    case 0:
    case "OBJECT_FIELD_TYPE_UNSPECIFIED":
      return ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED;
    case 1:
    case "OBJECT_FIELD_TYPE_STRING":
      return ObjectFieldType.OBJECT_FIELD_TYPE_STRING;
    case 2:
    case "OBJECT_FIELD_TYPE_NUMBER":
      return ObjectFieldType.OBJECT_FIELD_TYPE_NUMBER;
    case 3:
    case "OBJECT_FIELD_TYPE_BOOL":
      return ObjectFieldType.OBJECT_FIELD_TYPE_BOOL;
    case 4:
    case "OBJECT_FIELD_TYPE_ARRAY":
      return ObjectFieldType.OBJECT_FIELD_TYPE_ARRAY;
    case 5:
    case "OBJECT_FIELD_TYPE_TIMESTAMP":
      return ObjectFieldType.OBJECT_FIELD_TYPE_TIMESTAMP;
    case 6:
    case "OBJECT_FIELD_TYPE_GEOPOINT":
      return ObjectFieldType.OBJECT_FIELD_TYPE_GEOPOINT;
    case 7:
    case "OBJECT_FIELD_TYPE_BLOB":
      return ObjectFieldType.OBJECT_FIELD_TYPE_BLOB;
    case 8:
    case "OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS":
      return ObjectFieldType.OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ObjectFieldType.UNRECOGNIZED;
  }
}

export function objectFieldTypeToJSON(object: ObjectFieldType): string {
  switch (object) {
    case ObjectFieldType.OBJECT_FIELD_TYPE_UNSPECIFIED:
      return "OBJECT_FIELD_TYPE_UNSPECIFIED";
    case ObjectFieldType.OBJECT_FIELD_TYPE_STRING:
      return "OBJECT_FIELD_TYPE_STRING";
    case ObjectFieldType.OBJECT_FIELD_TYPE_NUMBER:
      return "OBJECT_FIELD_TYPE_NUMBER";
    case ObjectFieldType.OBJECT_FIELD_TYPE_BOOL:
      return "OBJECT_FIELD_TYPE_BOOL";
    case ObjectFieldType.OBJECT_FIELD_TYPE_ARRAY:
      return "OBJECT_FIELD_TYPE_ARRAY";
    case ObjectFieldType.OBJECT_FIELD_TYPE_TIMESTAMP:
      return "OBJECT_FIELD_TYPE_TIMESTAMP";
    case ObjectFieldType.OBJECT_FIELD_TYPE_GEOPOINT:
      return "OBJECT_FIELD_TYPE_GEOPOINT";
    case ObjectFieldType.OBJECT_FIELD_TYPE_BLOB:
      return "OBJECT_FIELD_TYPE_BLOB";
    case ObjectFieldType.OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS:
      return "OBJECT_FIELD_TYPE_BLOCKCHAIN_ADDRESS";
    default:
      return "UNKNOWN";
  }
}

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
  fields: { [key: string]: ObjectField };
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
  metadata: { [key: string]: string };
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
  metadata: { [key: string]: string };
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
  metadata: { [key: string]: string };
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
  metadata: { [key: string]: string };
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
  metadata: { [key: string]: string };
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
  metadata: { [key: string]: string };
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
  metadata: { [key: string]: string };
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
  metadata: { [key: string]: string };
}

export interface ObjectFieldBlockchainAddress_MetadataEntry {
  key: string;
  value: string;
}

const baseObjectDoc: object = {
  label: "",
  description: "",
  did: "",
  bucketDid: "",
};

export const ObjectDoc = {
  encode(message: ObjectDoc, writer: Writer = Writer.create()): Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.did !== "") {
      writer.uint32(26).string(message.did);
    }
    if (message.bucketDid !== "") {
      writer.uint32(34).string(message.bucketDid);
    }
    Object.entries(message.fields).forEach(([key, value]) => {
      ObjectDoc_FieldsEntry.encode(
        { key: key as any, value },
        writer.uint32(42).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ObjectDoc {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObjectDoc } as ObjectDoc;
    message.fields = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.did = reader.string();
          break;
        case 4:
          message.bucketDid = reader.string();
          break;
        case 5:
          const entry5 = ObjectDoc_FieldsEntry.decode(reader, reader.uint32());
          if (entry5.value !== undefined) {
            message.fields[entry5.key] = entry5.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectDoc {
    const message = { ...baseObjectDoc } as ObjectDoc;
    message.fields = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = String(object.label);
    } else {
      message.label = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.bucketDid !== undefined && object.bucketDid !== null) {
      message.bucketDid = String(object.bucketDid);
    } else {
      message.bucketDid = "";
    }
    if (object.fields !== undefined && object.fields !== null) {
      Object.entries(object.fields).forEach(([key, value]) => {
        message.fields[key] = ObjectField.fromJSON(value);
      });
    }
    return message;
  },

  toJSON(message: ObjectDoc): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.description !== undefined &&
      (obj.description = message.description);
    message.did !== undefined && (obj.did = message.did);
    message.bucketDid !== undefined && (obj.bucketDid = message.bucketDid);
    obj.fields = {};
    if (message.fields) {
      Object.entries(message.fields).forEach(([k, v]) => {
        obj.fields[k] = ObjectField.toJSON(v);
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ObjectDoc>): ObjectDoc {
    const message = { ...baseObjectDoc } as ObjectDoc;
    message.fields = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = object.label;
    } else {
      message.label = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.bucketDid !== undefined && object.bucketDid !== null) {
      message.bucketDid = object.bucketDid;
    } else {
      message.bucketDid = "";
    }
    if (object.fields !== undefined && object.fields !== null) {
      Object.entries(object.fields).forEach(([key, value]) => {
        if (value !== undefined) {
          message.fields[key] = ObjectField.fromPartial(value);
        }
      });
    }
    return message;
  },
};

const baseObjectDoc_FieldsEntry: object = { key: "" };

export const ObjectDoc_FieldsEntry = {
  encode(
    message: ObjectDoc_FieldsEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      ObjectField.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ObjectDoc_FieldsEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObjectDoc_FieldsEntry } as ObjectDoc_FieldsEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = ObjectField.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectDoc_FieldsEntry {
    const message = { ...baseObjectDoc_FieldsEntry } as ObjectDoc_FieldsEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = ObjectField.fromJSON(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },

  toJSON(message: ObjectDoc_FieldsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value
        ? ObjectField.toJSON(message.value)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ObjectDoc_FieldsEntry>
  ): ObjectDoc_FieldsEntry {
    const message = { ...baseObjectDoc_FieldsEntry } as ObjectDoc_FieldsEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = ObjectField.fromPartial(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },
};

const baseObjectField: object = { label: "", type: 0, did: "" };

export const ObjectField = {
  encode(message: ObjectField, writer: Writer = Writer.create()): Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.type !== 0) {
      writer.uint32(16).int32(message.type);
    }
    if (message.did !== "") {
      writer.uint32(26).string(message.did);
    }
    if (message.stringValue !== undefined) {
      ObjectFieldText.encode(
        message.stringValue,
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.numberValue !== undefined) {
      ObjectFieldNumber.encode(
        message.numberValue,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.boolValue !== undefined) {
      ObjectFieldBool.encode(
        message.boolValue,
        writer.uint32(50).fork()
      ).ldelim();
    }
    if (message.arrayValue !== undefined) {
      ObjectFieldArray.encode(
        message.arrayValue,
        writer.uint32(58).fork()
      ).ldelim();
    }
    if (message.timeStampValue !== undefined) {
      ObjectFieldTime.encode(
        message.timeStampValue,
        writer.uint32(66).fork()
      ).ldelim();
    }
    if (message.geopointValue !== undefined) {
      ObjectFieldGeopoint.encode(
        message.geopointValue,
        writer.uint32(74).fork()
      ).ldelim();
    }
    if (message.blobValue !== undefined) {
      ObjectFieldBlob.encode(
        message.blobValue,
        writer.uint32(82).fork()
      ).ldelim();
    }
    if (message.blockchainAddrValue !== undefined) {
      ObjectFieldBlockchainAddress.encode(
        message.blockchainAddrValue,
        writer.uint32(98).fork()
      ).ldelim();
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ObjectField_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(106).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ObjectField {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObjectField } as ObjectField;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.type = reader.int32() as any;
          break;
        case 3:
          message.did = reader.string();
          break;
        case 4:
          message.stringValue = ObjectFieldText.decode(reader, reader.uint32());
          break;
        case 5:
          message.numberValue = ObjectFieldNumber.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.boolValue = ObjectFieldBool.decode(reader, reader.uint32());
          break;
        case 7:
          message.arrayValue = ObjectFieldArray.decode(reader, reader.uint32());
          break;
        case 8:
          message.timeStampValue = ObjectFieldTime.decode(
            reader,
            reader.uint32()
          );
          break;
        case 9:
          message.geopointValue = ObjectFieldGeopoint.decode(
            reader,
            reader.uint32()
          );
          break;
        case 10:
          message.blobValue = ObjectFieldBlob.decode(reader, reader.uint32());
          break;
        case 12:
          message.blockchainAddrValue = ObjectFieldBlockchainAddress.decode(
            reader,
            reader.uint32()
          );
          break;
        case 13:
          const entry13 = ObjectField_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry13.value !== undefined) {
            message.metadata[entry13.key] = entry13.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectField {
    const message = { ...baseObjectField } as ObjectField;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = String(object.label);
    } else {
      message.label = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = objectFieldTypeFromJSON(object.type);
    } else {
      message.type = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.stringValue !== undefined && object.stringValue !== null) {
      message.stringValue = ObjectFieldText.fromJSON(object.stringValue);
    } else {
      message.stringValue = undefined;
    }
    if (object.numberValue !== undefined && object.numberValue !== null) {
      message.numberValue = ObjectFieldNumber.fromJSON(object.numberValue);
    } else {
      message.numberValue = undefined;
    }
    if (object.boolValue !== undefined && object.boolValue !== null) {
      message.boolValue = ObjectFieldBool.fromJSON(object.boolValue);
    } else {
      message.boolValue = undefined;
    }
    if (object.arrayValue !== undefined && object.arrayValue !== null) {
      message.arrayValue = ObjectFieldArray.fromJSON(object.arrayValue);
    } else {
      message.arrayValue = undefined;
    }
    if (object.timeStampValue !== undefined && object.timeStampValue !== null) {
      message.timeStampValue = ObjectFieldTime.fromJSON(object.timeStampValue);
    } else {
      message.timeStampValue = undefined;
    }
    if (object.geopointValue !== undefined && object.geopointValue !== null) {
      message.geopointValue = ObjectFieldGeopoint.fromJSON(
        object.geopointValue
      );
    } else {
      message.geopointValue = undefined;
    }
    if (object.blobValue !== undefined && object.blobValue !== null) {
      message.blobValue = ObjectFieldBlob.fromJSON(object.blobValue);
    } else {
      message.blobValue = undefined;
    }
    if (
      object.blockchainAddrValue !== undefined &&
      object.blockchainAddrValue !== null
    ) {
      message.blockchainAddrValue = ObjectFieldBlockchainAddress.fromJSON(
        object.blockchainAddrValue
      );
    } else {
      message.blockchainAddrValue = undefined;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: ObjectField): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.type !== undefined &&
      (obj.type = objectFieldTypeToJSON(message.type));
    message.did !== undefined && (obj.did = message.did);
    message.stringValue !== undefined &&
      (obj.stringValue = message.stringValue
        ? ObjectFieldText.toJSON(message.stringValue)
        : undefined);
    message.numberValue !== undefined &&
      (obj.numberValue = message.numberValue
        ? ObjectFieldNumber.toJSON(message.numberValue)
        : undefined);
    message.boolValue !== undefined &&
      (obj.boolValue = message.boolValue
        ? ObjectFieldBool.toJSON(message.boolValue)
        : undefined);
    message.arrayValue !== undefined &&
      (obj.arrayValue = message.arrayValue
        ? ObjectFieldArray.toJSON(message.arrayValue)
        : undefined);
    message.timeStampValue !== undefined &&
      (obj.timeStampValue = message.timeStampValue
        ? ObjectFieldTime.toJSON(message.timeStampValue)
        : undefined);
    message.geopointValue !== undefined &&
      (obj.geopointValue = message.geopointValue
        ? ObjectFieldGeopoint.toJSON(message.geopointValue)
        : undefined);
    message.blobValue !== undefined &&
      (obj.blobValue = message.blobValue
        ? ObjectFieldBlob.toJSON(message.blobValue)
        : undefined);
    message.blockchainAddrValue !== undefined &&
      (obj.blockchainAddrValue = message.blockchainAddrValue
        ? ObjectFieldBlockchainAddress.toJSON(message.blockchainAddrValue)
        : undefined);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ObjectField>): ObjectField {
    const message = { ...baseObjectField } as ObjectField;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = object.label;
    } else {
      message.label = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.stringValue !== undefined && object.stringValue !== null) {
      message.stringValue = ObjectFieldText.fromPartial(object.stringValue);
    } else {
      message.stringValue = undefined;
    }
    if (object.numberValue !== undefined && object.numberValue !== null) {
      message.numberValue = ObjectFieldNumber.fromPartial(object.numberValue);
    } else {
      message.numberValue = undefined;
    }
    if (object.boolValue !== undefined && object.boolValue !== null) {
      message.boolValue = ObjectFieldBool.fromPartial(object.boolValue);
    } else {
      message.boolValue = undefined;
    }
    if (object.arrayValue !== undefined && object.arrayValue !== null) {
      message.arrayValue = ObjectFieldArray.fromPartial(object.arrayValue);
    } else {
      message.arrayValue = undefined;
    }
    if (object.timeStampValue !== undefined && object.timeStampValue !== null) {
      message.timeStampValue = ObjectFieldTime.fromPartial(
        object.timeStampValue
      );
    } else {
      message.timeStampValue = undefined;
    }
    if (object.geopointValue !== undefined && object.geopointValue !== null) {
      message.geopointValue = ObjectFieldGeopoint.fromPartial(
        object.geopointValue
      );
    } else {
      message.geopointValue = undefined;
    }
    if (object.blobValue !== undefined && object.blobValue !== null) {
      message.blobValue = ObjectFieldBlob.fromPartial(object.blobValue);
    } else {
      message.blobValue = undefined;
    }
    if (
      object.blockchainAddrValue !== undefined &&
      object.blockchainAddrValue !== null
    ) {
      message.blockchainAddrValue = ObjectFieldBlockchainAddress.fromPartial(
        object.blockchainAddrValue
      );
    } else {
      message.blockchainAddrValue = undefined;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseObjectField_MetadataEntry: object = { key: "", value: "" };

export const ObjectField_MetadataEntry = {
  encode(
    message: ObjectField_MetadataEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ObjectField_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseObjectField_MetadataEntry,
    } as ObjectField_MetadataEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectField_MetadataEntry {
    const message = {
      ...baseObjectField_MetadataEntry,
    } as ObjectField_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: ObjectField_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ObjectField_MetadataEntry>
  ): ObjectField_MetadataEntry {
    const message = {
      ...baseObjectField_MetadataEntry,
    } as ObjectField_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseObjectFieldArray: object = { label: "", type: 0, did: "" };

export const ObjectFieldArray = {
  encode(message: ObjectFieldArray, writer: Writer = Writer.create()): Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.type !== 0) {
      writer.uint32(16).int32(message.type);
    }
    if (message.did !== "") {
      writer.uint32(26).string(message.did);
    }
    for (const v of message.items) {
      ObjectField.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ObjectFieldArray {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObjectFieldArray } as ObjectFieldArray;
    message.items = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.type = reader.int32() as any;
          break;
        case 3:
          message.did = reader.string();
          break;
        case 4:
          message.items.push(ObjectField.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldArray {
    const message = { ...baseObjectFieldArray } as ObjectFieldArray;
    message.items = [];
    if (object.label !== undefined && object.label !== null) {
      message.label = String(object.label);
    } else {
      message.label = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = objectFieldTypeFromJSON(object.type);
    } else {
      message.type = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.items !== undefined && object.items !== null) {
      for (const e of object.items) {
        message.items.push(ObjectField.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: ObjectFieldArray): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.type !== undefined &&
      (obj.type = objectFieldTypeToJSON(message.type));
    message.did !== undefined && (obj.did = message.did);
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? ObjectField.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ObjectFieldArray>): ObjectFieldArray {
    const message = { ...baseObjectFieldArray } as ObjectFieldArray;
    message.items = [];
    if (object.label !== undefined && object.label !== null) {
      message.label = object.label;
    } else {
      message.label = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.items !== undefined && object.items !== null) {
      for (const e of object.items) {
        message.items.push(ObjectField.fromPartial(e));
      }
    }
    return message;
  },
};

const baseObjectFieldText: object = { label: "", did: "", value: "" };

export const ObjectFieldText = {
  encode(message: ObjectFieldText, writer: Writer = Writer.create()): Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ObjectFieldText_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ObjectFieldText {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObjectFieldText } as ObjectFieldText;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.value = reader.string();
          break;
        case 4:
          const entry4 = ObjectFieldText_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry4.value !== undefined) {
            message.metadata[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldText {
    const message = { ...baseObjectFieldText } as ObjectFieldText;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = String(object.label);
    } else {
      message.label = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: ObjectFieldText): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.did !== undefined && (obj.did = message.did);
    message.value !== undefined && (obj.value = message.value);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ObjectFieldText>): ObjectFieldText {
    const message = { ...baseObjectFieldText } as ObjectFieldText;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = object.label;
    } else {
      message.label = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseObjectFieldText_MetadataEntry: object = { key: "", value: "" };

export const ObjectFieldText_MetadataEntry = {
  encode(
    message: ObjectFieldText_MetadataEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ObjectFieldText_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseObjectFieldText_MetadataEntry,
    } as ObjectFieldText_MetadataEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldText_MetadataEntry {
    const message = {
      ...baseObjectFieldText_MetadataEntry,
    } as ObjectFieldText_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: ObjectFieldText_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ObjectFieldText_MetadataEntry>
  ): ObjectFieldText_MetadataEntry {
    const message = {
      ...baseObjectFieldText_MetadataEntry,
    } as ObjectFieldText_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseObjectFieldNumber: object = { label: "", did: "", value: 0 };

export const ObjectFieldNumber = {
  encode(message: ObjectFieldNumber, writer: Writer = Writer.create()): Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.value !== 0) {
      writer.uint32(25).double(message.value);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ObjectFieldNumber_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ObjectFieldNumber {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObjectFieldNumber } as ObjectFieldNumber;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.value = reader.double();
          break;
        case 4:
          const entry4 = ObjectFieldNumber_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry4.value !== undefined) {
            message.metadata[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldNumber {
    const message = { ...baseObjectFieldNumber } as ObjectFieldNumber;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = String(object.label);
    } else {
      message.label = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Number(object.value);
    } else {
      message.value = 0;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: ObjectFieldNumber): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.did !== undefined && (obj.did = message.did);
    message.value !== undefined && (obj.value = message.value);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ObjectFieldNumber>): ObjectFieldNumber {
    const message = { ...baseObjectFieldNumber } as ObjectFieldNumber;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = object.label;
    } else {
      message.label = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = 0;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseObjectFieldNumber_MetadataEntry: object = { key: "", value: "" };

export const ObjectFieldNumber_MetadataEntry = {
  encode(
    message: ObjectFieldNumber_MetadataEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ObjectFieldNumber_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseObjectFieldNumber_MetadataEntry,
    } as ObjectFieldNumber_MetadataEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldNumber_MetadataEntry {
    const message = {
      ...baseObjectFieldNumber_MetadataEntry,
    } as ObjectFieldNumber_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: ObjectFieldNumber_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ObjectFieldNumber_MetadataEntry>
  ): ObjectFieldNumber_MetadataEntry {
    const message = {
      ...baseObjectFieldNumber_MetadataEntry,
    } as ObjectFieldNumber_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseObjectFieldBool: object = { label: "", did: "", value: false };

export const ObjectFieldBool = {
  encode(message: ObjectFieldBool, writer: Writer = Writer.create()): Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.value === true) {
      writer.uint32(24).bool(message.value);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ObjectFieldBool_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ObjectFieldBool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObjectFieldBool } as ObjectFieldBool;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.value = reader.bool();
          break;
        case 4:
          const entry4 = ObjectFieldBool_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry4.value !== undefined) {
            message.metadata[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldBool {
    const message = { ...baseObjectFieldBool } as ObjectFieldBool;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = String(object.label);
    } else {
      message.label = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Boolean(object.value);
    } else {
      message.value = false;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: ObjectFieldBool): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.did !== undefined && (obj.did = message.did);
    message.value !== undefined && (obj.value = message.value);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ObjectFieldBool>): ObjectFieldBool {
    const message = { ...baseObjectFieldBool } as ObjectFieldBool;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = object.label;
    } else {
      message.label = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = false;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseObjectFieldBool_MetadataEntry: object = { key: "", value: "" };

export const ObjectFieldBool_MetadataEntry = {
  encode(
    message: ObjectFieldBool_MetadataEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ObjectFieldBool_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseObjectFieldBool_MetadataEntry,
    } as ObjectFieldBool_MetadataEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldBool_MetadataEntry {
    const message = {
      ...baseObjectFieldBool_MetadataEntry,
    } as ObjectFieldBool_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: ObjectFieldBool_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ObjectFieldBool_MetadataEntry>
  ): ObjectFieldBool_MetadataEntry {
    const message = {
      ...baseObjectFieldBool_MetadataEntry,
    } as ObjectFieldBool_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseObjectFieldTime: object = { label: "", did: "", value: 0 };

export const ObjectFieldTime = {
  encode(message: ObjectFieldTime, writer: Writer = Writer.create()): Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.value !== 0) {
      writer.uint32(24).int64(message.value);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ObjectFieldTime_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ObjectFieldTime {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObjectFieldTime } as ObjectFieldTime;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.value = longToNumber(reader.int64() as Long);
          break;
        case 4:
          const entry4 = ObjectFieldTime_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry4.value !== undefined) {
            message.metadata[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldTime {
    const message = { ...baseObjectFieldTime } as ObjectFieldTime;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = String(object.label);
    } else {
      message.label = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Number(object.value);
    } else {
      message.value = 0;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: ObjectFieldTime): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.did !== undefined && (obj.did = message.did);
    message.value !== undefined && (obj.value = message.value);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ObjectFieldTime>): ObjectFieldTime {
    const message = { ...baseObjectFieldTime } as ObjectFieldTime;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = object.label;
    } else {
      message.label = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = 0;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseObjectFieldTime_MetadataEntry: object = { key: "", value: "" };

export const ObjectFieldTime_MetadataEntry = {
  encode(
    message: ObjectFieldTime_MetadataEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ObjectFieldTime_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseObjectFieldTime_MetadataEntry,
    } as ObjectFieldTime_MetadataEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldTime_MetadataEntry {
    const message = {
      ...baseObjectFieldTime_MetadataEntry,
    } as ObjectFieldTime_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: ObjectFieldTime_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ObjectFieldTime_MetadataEntry>
  ): ObjectFieldTime_MetadataEntry {
    const message = {
      ...baseObjectFieldTime_MetadataEntry,
    } as ObjectFieldTime_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseObjectFieldGeopoint: object = {
  label: "",
  type: 0,
  did: "",
  latitude: 0,
  longitude: 0,
};

export const ObjectFieldGeopoint = {
  encode(
    message: ObjectFieldGeopoint,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.type !== 0) {
      writer.uint32(16).int32(message.type);
    }
    if (message.did !== "") {
      writer.uint32(26).string(message.did);
    }
    if (message.latitude !== 0) {
      writer.uint32(33).double(message.latitude);
    }
    if (message.longitude !== 0) {
      writer.uint32(41).double(message.longitude);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ObjectFieldGeopoint_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(50).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ObjectFieldGeopoint {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObjectFieldGeopoint } as ObjectFieldGeopoint;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.type = reader.int32() as any;
          break;
        case 3:
          message.did = reader.string();
          break;
        case 4:
          message.latitude = reader.double();
          break;
        case 5:
          message.longitude = reader.double();
          break;
        case 6:
          const entry6 = ObjectFieldGeopoint_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry6.value !== undefined) {
            message.metadata[entry6.key] = entry6.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldGeopoint {
    const message = { ...baseObjectFieldGeopoint } as ObjectFieldGeopoint;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = String(object.label);
    } else {
      message.label = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = objectFieldTypeFromJSON(object.type);
    } else {
      message.type = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.latitude !== undefined && object.latitude !== null) {
      message.latitude = Number(object.latitude);
    } else {
      message.latitude = 0;
    }
    if (object.longitude !== undefined && object.longitude !== null) {
      message.longitude = Number(object.longitude);
    } else {
      message.longitude = 0;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: ObjectFieldGeopoint): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.type !== undefined &&
      (obj.type = objectFieldTypeToJSON(message.type));
    message.did !== undefined && (obj.did = message.did);
    message.latitude !== undefined && (obj.latitude = message.latitude);
    message.longitude !== undefined && (obj.longitude = message.longitude);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ObjectFieldGeopoint>): ObjectFieldGeopoint {
    const message = { ...baseObjectFieldGeopoint } as ObjectFieldGeopoint;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = object.label;
    } else {
      message.label = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.latitude !== undefined && object.latitude !== null) {
      message.latitude = object.latitude;
    } else {
      message.latitude = 0;
    }
    if (object.longitude !== undefined && object.longitude !== null) {
      message.longitude = object.longitude;
    } else {
      message.longitude = 0;
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseObjectFieldGeopoint_MetadataEntry: object = { key: "", value: "" };

export const ObjectFieldGeopoint_MetadataEntry = {
  encode(
    message: ObjectFieldGeopoint_MetadataEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ObjectFieldGeopoint_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseObjectFieldGeopoint_MetadataEntry,
    } as ObjectFieldGeopoint_MetadataEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldGeopoint_MetadataEntry {
    const message = {
      ...baseObjectFieldGeopoint_MetadataEntry,
    } as ObjectFieldGeopoint_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: ObjectFieldGeopoint_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ObjectFieldGeopoint_MetadataEntry>
  ): ObjectFieldGeopoint_MetadataEntry {
    const message = {
      ...baseObjectFieldGeopoint_MetadataEntry,
    } as ObjectFieldGeopoint_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseObjectFieldBlob: object = { label: "", did: "" };

export const ObjectFieldBlob = {
  encode(message: ObjectFieldBlob, writer: Writer = Writer.create()): Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.value.length !== 0) {
      writer.uint32(26).bytes(message.value);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ObjectFieldBlob_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ObjectFieldBlob {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObjectFieldBlob } as ObjectFieldBlob;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.value = reader.bytes();
          break;
        case 4:
          const entry4 = ObjectFieldBlob_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry4.value !== undefined) {
            message.metadata[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldBlob {
    const message = { ...baseObjectFieldBlob } as ObjectFieldBlob;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = String(object.label);
    } else {
      message.label = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = bytesFromBase64(object.value);
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: ObjectFieldBlob): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.did !== undefined && (obj.did = message.did);
    message.value !== undefined &&
      (obj.value = base64FromBytes(
        message.value !== undefined ? message.value : new Uint8Array()
      ));
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ObjectFieldBlob>): ObjectFieldBlob {
    const message = { ...baseObjectFieldBlob } as ObjectFieldBlob;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = object.label;
    } else {
      message.label = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = new Uint8Array();
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseObjectFieldBlob_MetadataEntry: object = { key: "", value: "" };

export const ObjectFieldBlob_MetadataEntry = {
  encode(
    message: ObjectFieldBlob_MetadataEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ObjectFieldBlob_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseObjectFieldBlob_MetadataEntry,
    } as ObjectFieldBlob_MetadataEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldBlob_MetadataEntry {
    const message = {
      ...baseObjectFieldBlob_MetadataEntry,
    } as ObjectFieldBlob_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: ObjectFieldBlob_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ObjectFieldBlob_MetadataEntry>
  ): ObjectFieldBlob_MetadataEntry {
    const message = {
      ...baseObjectFieldBlob_MetadataEntry,
    } as ObjectFieldBlob_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseObjectFieldBlockchainAddress: object = {
  label: "",
  did: "",
  value: "",
};

export const ObjectFieldBlockchainAddress = {
  encode(
    message: ObjectFieldBlockchainAddress,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      ObjectFieldBlockchainAddress_MetadataEntry.encode(
        { key: key as any, value },
        writer.uint32(34).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ObjectFieldBlockchainAddress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseObjectFieldBlockchainAddress,
    } as ObjectFieldBlockchainAddress;
    message.metadata = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.value = reader.string();
          break;
        case 4:
          const entry4 = ObjectFieldBlockchainAddress_MetadataEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry4.value !== undefined) {
            message.metadata[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldBlockchainAddress {
    const message = {
      ...baseObjectFieldBlockchainAddress,
    } as ObjectFieldBlockchainAddress;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = String(object.label);
    } else {
      message.label = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        message.metadata[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: ObjectFieldBlockchainAddress): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.did !== undefined && (obj.did = message.did);
    message.value !== undefined && (obj.value = message.value);
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<ObjectFieldBlockchainAddress>
  ): ObjectFieldBlockchainAddress {
    const message = {
      ...baseObjectFieldBlockchainAddress,
    } as ObjectFieldBlockchainAddress;
    message.metadata = {};
    if (object.label !== undefined && object.label !== null) {
      message.label = object.label;
    } else {
      message.label = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      Object.entries(object.metadata).forEach(([key, value]) => {
        if (value !== undefined) {
          message.metadata[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseObjectFieldBlockchainAddress_MetadataEntry: object = {
  key: "",
  value: "",
};

export const ObjectFieldBlockchainAddress_MetadataEntry = {
  encode(
    message: ObjectFieldBlockchainAddress_MetadataEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ObjectFieldBlockchainAddress_MetadataEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseObjectFieldBlockchainAddress_MetadataEntry,
    } as ObjectFieldBlockchainAddress_MetadataEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ObjectFieldBlockchainAddress_MetadataEntry {
    const message = {
      ...baseObjectFieldBlockchainAddress_MetadataEntry,
    } as ObjectFieldBlockchainAddress_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: ObjectFieldBlockchainAddress_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ObjectFieldBlockchainAddress_MetadataEntry>
  ): ObjectFieldBlockchainAddress_MetadataEntry {
    const message = {
      ...baseObjectFieldBlockchainAddress_MetadataEntry,
    } as ObjectFieldBlockchainAddress_MetadataEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
