/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "sonrio.sonr.object";

export enum TypeKind {
  TypeKind_Invalid = 0,
  TypeKind_Map = 1,
  TypeKind_List = 2,
  TypeKind_Unit = 3,
  TypeKind_Bool = 4,
  TypeKind_Int = 5,
  TypeKind_Float = 6,
  TypeKind_String = 7,
  TypeKind_Bytes = 8,
  TypeKind_Link = 9,
  TypeKind_Struct = 10,
  TypeKind_Union = 11,
  TypeKind_Enum = 12,
  TypeKind_Any = 13,
  UNRECOGNIZED = -1,
}

export function typeKindFromJSON(object: any): TypeKind {
  switch (object) {
    case 0:
    case "TypeKind_Invalid":
      return TypeKind.TypeKind_Invalid;
    case 1:
    case "TypeKind_Map":
      return TypeKind.TypeKind_Map;
    case 2:
    case "TypeKind_List":
      return TypeKind.TypeKind_List;
    case 3:
    case "TypeKind_Unit":
      return TypeKind.TypeKind_Unit;
    case 4:
    case "TypeKind_Bool":
      return TypeKind.TypeKind_Bool;
    case 5:
    case "TypeKind_Int":
      return TypeKind.TypeKind_Int;
    case 6:
    case "TypeKind_Float":
      return TypeKind.TypeKind_Float;
    case 7:
    case "TypeKind_String":
      return TypeKind.TypeKind_String;
    case 8:
    case "TypeKind_Bytes":
      return TypeKind.TypeKind_Bytes;
    case 9:
    case "TypeKind_Link":
      return TypeKind.TypeKind_Link;
    case 10:
    case "TypeKind_Struct":
      return TypeKind.TypeKind_Struct;
    case 11:
    case "TypeKind_Union":
      return TypeKind.TypeKind_Union;
    case 12:
    case "TypeKind_Enum":
      return TypeKind.TypeKind_Enum;
    case 13:
    case "TypeKind_Any":
      return TypeKind.TypeKind_Any;
    case -1:
    case "UNRECOGNIZED":
    default:
      return TypeKind.UNRECOGNIZED;
  }
}

export function typeKindToJSON(object: TypeKind): string {
  switch (object) {
    case TypeKind.TypeKind_Invalid:
      return "TypeKind_Invalid";
    case TypeKind.TypeKind_Map:
      return "TypeKind_Map";
    case TypeKind.TypeKind_List:
      return "TypeKind_List";
    case TypeKind.TypeKind_Unit:
      return "TypeKind_Unit";
    case TypeKind.TypeKind_Bool:
      return "TypeKind_Bool";
    case TypeKind.TypeKind_Int:
      return "TypeKind_Int";
    case TypeKind.TypeKind_Float:
      return "TypeKind_Float";
    case TypeKind.TypeKind_String:
      return "TypeKind_String";
    case TypeKind.TypeKind_Bytes:
      return "TypeKind_Bytes";
    case TypeKind.TypeKind_Link:
      return "TypeKind_Link";
    case TypeKind.TypeKind_Struct:
      return "TypeKind_Struct";
    case TypeKind.TypeKind_Union:
      return "TypeKind_Union";
    case TypeKind.TypeKind_Enum:
      return "TypeKind_Enum";
    case TypeKind.TypeKind_Any:
      return "TypeKind_Any";
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
  bucket_did: string;
  /** Fields are the fields associated with the object. */
  fields: TypeField[];
}

export interface TypeField {
  /** Name is the name of the field. */
  name: string;
  /** Type is the type of the field. */
  kind: TypeKind;
}

const baseObjectDoc: object = {
  label: "",
  description: "",
  did: "",
  bucket_did: "",
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
    if (message.bucket_did !== "") {
      writer.uint32(34).string(message.bucket_did);
    }
    for (const v of message.fields) {
      TypeField.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ObjectDoc {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObjectDoc } as ObjectDoc;
    message.fields = [];
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
          message.bucket_did = reader.string();
          break;
        case 5:
          message.fields.push(TypeField.decode(reader, reader.uint32()));
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
    message.fields = [];
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
    if (object.bucket_did !== undefined && object.bucket_did !== null) {
      message.bucket_did = String(object.bucket_did);
    } else {
      message.bucket_did = "";
    }
    if (object.fields !== undefined && object.fields !== null) {
      for (const e of object.fields) {
        message.fields.push(TypeField.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: ObjectDoc): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.description !== undefined &&
      (obj.description = message.description);
    message.did !== undefined && (obj.did = message.did);
    message.bucket_did !== undefined && (obj.bucket_did = message.bucket_did);
    if (message.fields) {
      obj.fields = message.fields.map((e) =>
        e ? TypeField.toJSON(e) : undefined
      );
    } else {
      obj.fields = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ObjectDoc>): ObjectDoc {
    const message = { ...baseObjectDoc } as ObjectDoc;
    message.fields = [];
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
    if (object.bucket_did !== undefined && object.bucket_did !== null) {
      message.bucket_did = object.bucket_did;
    } else {
      message.bucket_did = "";
    }
    if (object.fields !== undefined && object.fields !== null) {
      for (const e of object.fields) {
        message.fields.push(TypeField.fromPartial(e));
      }
    }
    return message;
  },
};

const baseTypeField: object = { name: "", kind: 0 };

export const TypeField = {
  encode(message: TypeField, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.kind !== 0) {
      writer.uint32(16).int32(message.kind);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TypeField {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTypeField } as TypeField;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.kind = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TypeField {
    const message = { ...baseTypeField } as TypeField;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.kind !== undefined && object.kind !== null) {
      message.kind = typeKindFromJSON(object.kind);
    } else {
      message.kind = 0;
    }
    return message;
  },

  toJSON(message: TypeField): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.kind !== undefined && (obj.kind = typeKindToJSON(message.kind));
    return obj;
  },

  fromPartial(object: DeepPartial<TypeField>): TypeField {
    const message = { ...baseTypeField } as TypeField;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.kind !== undefined && object.kind !== null) {
      message.kind = object.kind;
    } else {
      message.kind = 0;
    }
    return message;
  },
};

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
