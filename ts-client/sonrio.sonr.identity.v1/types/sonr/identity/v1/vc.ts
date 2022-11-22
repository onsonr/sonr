/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { ProofType, proofTypeFromJSON, proofTypeToJSON } from "./ssi";

export const protobufPackage = "sonrio.sonr.identity.v1";

/** Proof represents a credential/presentation proof as defined by the Linked Data Proofs 1.0 specification (https://w3c-ccg.github.io/ld-proofs/). */
export interface Proof {
  /** Type defines the specific proof type used. For example, an Ed25519Signature2018 type indicates that the proof includes a digital signature produced by an ed25519 cryptographic key. */
  type: ProofType;
  /** ProofPurpose defines the intent for the proof, the reason why an entity created it. Acts as a safeguard to prevent the proof from being misused for a purpose other than the one it was intended for. */
  proofPurpose: string;
  /** VerificationMethod points to the ID that can be used to verify the proof, eg: a public key. */
  verificationMethod: string;
  /** Created notes when the proof was created using a iso8601 string */
  created: string;
  /** Domain specifies the restricted domain of the proof */
  domain: string;
}

/** JSONWebSignature2020Proof is a VC proof with a signature according to JsonWebSignature2020 */
export interface JSONWebSignature2020Proof {
  proof: Proof | undefined;
  jws: string;
}

/** VerifiableCredential represents a credential as defined by the Verifiable Credentials Data Model 1.0 specification (https://www.w3.org/TR/vc-data-model/). */
export interface VerifiableCredential {
  /** ID is the unique identifier for the credential. */
  id: string;
  /** Context is a list of URIs that define the context of the credential. */
  context: string[];
  /** Type is a list of URIs that define the type of the credential. */
  type: string[];
  /** Issuer is the DID of the issuer of the credential. */
  issuer: string;
  /** IssuanceDate is the date the credential was issued. */
  issuanceDate: string;
  /** ExpirationDate is the date the credential expires. */
  expirationDate: string;
  /** CredentialSubject is the subject of the credential. */
  credentialSubject: { [key: string]: string };
  /** Proof is the proof of the credential. */
  proof: { [key: string]: string };
}

export interface VerifiableCredential_CredentialSubjectEntry {
  key: string;
  value: string;
}

export interface VerifiableCredential_ProofEntry {
  key: string;
  value: string;
}

function createBaseProof(): Proof {
  return { type: 0, proofPurpose: "", verificationMethod: "", created: "", domain: "" };
}

export const Proof = {
  encode(message: Proof, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.proofPurpose !== "") {
      writer.uint32(18).string(message.proofPurpose);
    }
    if (message.verificationMethod !== "") {
      writer.uint32(26).string(message.verificationMethod);
    }
    if (message.created !== "") {
      writer.uint32(34).string(message.created);
    }
    if (message.domain !== "") {
      writer.uint32(42).string(message.domain);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Proof {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProof();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.int32() as any;
          break;
        case 2:
          message.proofPurpose = reader.string();
          break;
        case 3:
          message.verificationMethod = reader.string();
          break;
        case 4:
          message.created = reader.string();
          break;
        case 5:
          message.domain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Proof {
    return {
      type: isSet(object.type) ? proofTypeFromJSON(object.type) : 0,
      proofPurpose: isSet(object.proofPurpose) ? String(object.proofPurpose) : "",
      verificationMethod: isSet(object.verificationMethod) ? String(object.verificationMethod) : "",
      created: isSet(object.created) ? String(object.created) : "",
      domain: isSet(object.domain) ? String(object.domain) : "",
    };
  },

  toJSON(message: Proof): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = proofTypeToJSON(message.type));
    message.proofPurpose !== undefined && (obj.proofPurpose = message.proofPurpose);
    message.verificationMethod !== undefined && (obj.verificationMethod = message.verificationMethod);
    message.created !== undefined && (obj.created = message.created);
    message.domain !== undefined && (obj.domain = message.domain);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Proof>, I>>(object: I): Proof {
    const message = createBaseProof();
    message.type = object.type ?? 0;
    message.proofPurpose = object.proofPurpose ?? "";
    message.verificationMethod = object.verificationMethod ?? "";
    message.created = object.created ?? "";
    message.domain = object.domain ?? "";
    return message;
  },
};

function createBaseJSONWebSignature2020Proof(): JSONWebSignature2020Proof {
  return { proof: undefined, jws: "" };
}

export const JSONWebSignature2020Proof = {
  encode(message: JSONWebSignature2020Proof, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.proof !== undefined) {
      Proof.encode(message.proof, writer.uint32(10).fork()).ldelim();
    }
    if (message.jws !== "") {
      writer.uint32(18).string(message.jws);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): JSONWebSignature2020Proof {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseJSONWebSignature2020Proof();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proof = Proof.decode(reader, reader.uint32());
          break;
        case 2:
          message.jws = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): JSONWebSignature2020Proof {
    return {
      proof: isSet(object.proof) ? Proof.fromJSON(object.proof) : undefined,
      jws: isSet(object.jws) ? String(object.jws) : "",
    };
  },

  toJSON(message: JSONWebSignature2020Proof): unknown {
    const obj: any = {};
    message.proof !== undefined && (obj.proof = message.proof ? Proof.toJSON(message.proof) : undefined);
    message.jws !== undefined && (obj.jws = message.jws);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<JSONWebSignature2020Proof>, I>>(object: I): JSONWebSignature2020Proof {
    const message = createBaseJSONWebSignature2020Proof();
    message.proof = (object.proof !== undefined && object.proof !== null) ? Proof.fromPartial(object.proof) : undefined;
    message.jws = object.jws ?? "";
    return message;
  },
};

function createBaseVerifiableCredential(): VerifiableCredential {
  return {
    id: "",
    context: [],
    type: [],
    issuer: "",
    issuanceDate: "",
    expirationDate: "",
    credentialSubject: {},
    proof: {},
  };
}

export const VerifiableCredential = {
  encode(message: VerifiableCredential, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    for (const v of message.context) {
      writer.uint32(18).string(v!);
    }
    for (const v of message.type) {
      writer.uint32(26).string(v!);
    }
    if (message.issuer !== "") {
      writer.uint32(34).string(message.issuer);
    }
    if (message.issuanceDate !== "") {
      writer.uint32(42).string(message.issuanceDate);
    }
    if (message.expirationDate !== "") {
      writer.uint32(50).string(message.expirationDate);
    }
    Object.entries(message.credentialSubject).forEach(([key, value]) => {
      VerifiableCredential_CredentialSubjectEntry.encode({ key: key as any, value }, writer.uint32(58).fork()).ldelim();
    });
    Object.entries(message.proof).forEach(([key, value]) => {
      VerifiableCredential_ProofEntry.encode({ key: key as any, value }, writer.uint32(66).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VerifiableCredential {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVerifiableCredential();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.context.push(reader.string());
          break;
        case 3:
          message.type.push(reader.string());
          break;
        case 4:
          message.issuer = reader.string();
          break;
        case 5:
          message.issuanceDate = reader.string();
          break;
        case 6:
          message.expirationDate = reader.string();
          break;
        case 7:
          const entry7 = VerifiableCredential_CredentialSubjectEntry.decode(reader, reader.uint32());
          if (entry7.value !== undefined) {
            message.credentialSubject[entry7.key] = entry7.value;
          }
          break;
        case 8:
          const entry8 = VerifiableCredential_ProofEntry.decode(reader, reader.uint32());
          if (entry8.value !== undefined) {
            message.proof[entry8.key] = entry8.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VerifiableCredential {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      context: Array.isArray(object?.context) ? object.context.map((e: any) => String(e)) : [],
      type: Array.isArray(object?.type) ? object.type.map((e: any) => String(e)) : [],
      issuer: isSet(object.issuer) ? String(object.issuer) : "",
      issuanceDate: isSet(object.issuanceDate) ? String(object.issuanceDate) : "",
      expirationDate: isSet(object.expirationDate) ? String(object.expirationDate) : "",
      credentialSubject: isObject(object.credentialSubject)
        ? Object.entries(object.credentialSubject).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
      proof: isObject(object.proof)
        ? Object.entries(object.proof).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
    };
  },

  toJSON(message: VerifiableCredential): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    if (message.context) {
      obj.context = message.context.map((e) => e);
    } else {
      obj.context = [];
    }
    if (message.type) {
      obj.type = message.type.map((e) => e);
    } else {
      obj.type = [];
    }
    message.issuer !== undefined && (obj.issuer = message.issuer);
    message.issuanceDate !== undefined && (obj.issuanceDate = message.issuanceDate);
    message.expirationDate !== undefined && (obj.expirationDate = message.expirationDate);
    obj.credentialSubject = {};
    if (message.credentialSubject) {
      Object.entries(message.credentialSubject).forEach(([k, v]) => {
        obj.credentialSubject[k] = v;
      });
    }
    obj.proof = {};
    if (message.proof) {
      Object.entries(message.proof).forEach(([k, v]) => {
        obj.proof[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VerifiableCredential>, I>>(object: I): VerifiableCredential {
    const message = createBaseVerifiableCredential();
    message.id = object.id ?? "";
    message.context = object.context?.map((e) => e) || [];
    message.type = object.type?.map((e) => e) || [];
    message.issuer = object.issuer ?? "";
    message.issuanceDate = object.issuanceDate ?? "";
    message.expirationDate = object.expirationDate ?? "";
    message.credentialSubject = Object.entries(object.credentialSubject ?? {}).reduce<{ [key: string]: string }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = String(value);
        }
        return acc;
      },
      {},
    );
    message.proof = Object.entries(object.proof ?? {}).reduce<{ [key: string]: string }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseVerifiableCredential_CredentialSubjectEntry(): VerifiableCredential_CredentialSubjectEntry {
  return { key: "", value: "" };
}

export const VerifiableCredential_CredentialSubjectEntry = {
  encode(message: VerifiableCredential_CredentialSubjectEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VerifiableCredential_CredentialSubjectEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVerifiableCredential_CredentialSubjectEntry();
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

  fromJSON(object: any): VerifiableCredential_CredentialSubjectEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? String(object.value) : "" };
  },

  toJSON(message: VerifiableCredential_CredentialSubjectEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VerifiableCredential_CredentialSubjectEntry>, I>>(
    object: I,
  ): VerifiableCredential_CredentialSubjectEntry {
    const message = createBaseVerifiableCredential_CredentialSubjectEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseVerifiableCredential_ProofEntry(): VerifiableCredential_ProofEntry {
  return { key: "", value: "" };
}

export const VerifiableCredential_ProofEntry = {
  encode(message: VerifiableCredential_ProofEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VerifiableCredential_ProofEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVerifiableCredential_ProofEntry();
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

  fromJSON(object: any): VerifiableCredential_ProofEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? String(object.value) : "" };
  },

  toJSON(message: VerifiableCredential_ProofEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VerifiableCredential_ProofEntry>, I>>(
    object: I,
  ): VerifiableCredential_ProofEntry {
    const message = createBaseVerifiableCredential_ProofEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
