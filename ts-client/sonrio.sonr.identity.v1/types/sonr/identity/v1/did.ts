/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { KeyType, keyTypeFromJSON, keyTypeToJSON } from "./ssi";

export const protobufPackage = "sonrio.sonr.identity.v1";

export interface DidDocument {
  /** optional */
  context: string[];
  /** optional */
  creator: string;
  iD: string;
  /** optional */
  controller: string[];
  /** optional */
  verificationMethod:
    | VerificationMethods
    | undefined;
  /** optional */
  authentication:
    | VerificationRelationships
    | undefined;
  /** optional */
  assertionMethod:
    | VerificationRelationships
    | undefined;
  /** optional */
  capabilityInvocation:
    | VerificationRelationships
    | undefined;
  /** optional */
  capabilityDelegation:
    | VerificationRelationships
    | undefined;
  /** optional */
  keyAgreement:
    | VerificationRelationships
    | undefined;
  /** optional */
  service:
    | Services
    | undefined;
  /** optional */
  alsoKnownAs: string[];
}

export interface VerificationMethod {
  iD: string;
  type: KeyType;
  controller: string;
  publicKeyJwk: { [key: string]: string };
  /** optional */
  publicKeyMultibase: string;
}

export interface VerificationMethod_PublicKeyJwkEntry {
  key: string;
  value: string;
}

export interface VerificationRelationship {
  verificationMethod: VerificationMethod | undefined;
  reference: string;
}

export interface Service {
  iD: string;
  type: string;
  serviceEndpoint: string;
  /** optional */
  serviceEndpoints: { [key: string]: string };
}

export interface Service_ServiceEndpointsEntry {
  key: string;
  value: string;
}

export interface Services {
  data: Service[];
}

export interface VerificationMethods {
  data: VerificationMethod[];
}

export interface VerificationRelationships {
  data: VerificationRelationship[];
}

function createBaseDidDocument(): DidDocument {
  return {
    context: [],
    creator: "",
    iD: "",
    controller: [],
    verificationMethod: undefined,
    authentication: undefined,
    assertionMethod: undefined,
    capabilityInvocation: undefined,
    capabilityDelegation: undefined,
    keyAgreement: undefined,
    service: undefined,
    alsoKnownAs: [],
  };
}

export const DidDocument = {
  encode(message: DidDocument, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.context) {
      writer.uint32(10).string(v!);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.iD !== "") {
      writer.uint32(26).string(message.iD);
    }
    for (const v of message.controller) {
      writer.uint32(34).string(v!);
    }
    if (message.verificationMethod !== undefined) {
      VerificationMethods.encode(message.verificationMethod, writer.uint32(42).fork()).ldelim();
    }
    if (message.authentication !== undefined) {
      VerificationRelationships.encode(message.authentication, writer.uint32(50).fork()).ldelim();
    }
    if (message.assertionMethod !== undefined) {
      VerificationRelationships.encode(message.assertionMethod, writer.uint32(58).fork()).ldelim();
    }
    if (message.capabilityInvocation !== undefined) {
      VerificationRelationships.encode(message.capabilityInvocation, writer.uint32(66).fork()).ldelim();
    }
    if (message.capabilityDelegation !== undefined) {
      VerificationRelationships.encode(message.capabilityDelegation, writer.uint32(74).fork()).ldelim();
    }
    if (message.keyAgreement !== undefined) {
      VerificationRelationships.encode(message.keyAgreement, writer.uint32(82).fork()).ldelim();
    }
    if (message.service !== undefined) {
      Services.encode(message.service, writer.uint32(90).fork()).ldelim();
    }
    for (const v of message.alsoKnownAs) {
      writer.uint32(98).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DidDocument {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDidDocument();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.context.push(reader.string());
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.iD = reader.string();
          break;
        case 4:
          message.controller.push(reader.string());
          break;
        case 5:
          message.verificationMethod = VerificationMethods.decode(reader, reader.uint32());
          break;
        case 6:
          message.authentication = VerificationRelationships.decode(reader, reader.uint32());
          break;
        case 7:
          message.assertionMethod = VerificationRelationships.decode(reader, reader.uint32());
          break;
        case 8:
          message.capabilityInvocation = VerificationRelationships.decode(reader, reader.uint32());
          break;
        case 9:
          message.capabilityDelegation = VerificationRelationships.decode(reader, reader.uint32());
          break;
        case 10:
          message.keyAgreement = VerificationRelationships.decode(reader, reader.uint32());
          break;
        case 11:
          message.service = Services.decode(reader, reader.uint32());
          break;
        case 12:
          message.alsoKnownAs.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DidDocument {
    return {
      context: Array.isArray(object?.context) ? object.context.map((e: any) => String(e)) : [],
      creator: isSet(object.creator) ? String(object.creator) : "",
      iD: isSet(object.iD) ? String(object.iD) : "",
      controller: Array.isArray(object?.controller) ? object.controller.map((e: any) => String(e)) : [],
      verificationMethod: isSet(object.verificationMethod)
        ? VerificationMethods.fromJSON(object.verificationMethod)
        : undefined,
      authentication: isSet(object.authentication)
        ? VerificationRelationships.fromJSON(object.authentication)
        : undefined,
      assertionMethod: isSet(object.assertionMethod)
        ? VerificationRelationships.fromJSON(object.assertionMethod)
        : undefined,
      capabilityInvocation: isSet(object.capabilityInvocation)
        ? VerificationRelationships.fromJSON(object.capabilityInvocation)
        : undefined,
      capabilityDelegation: isSet(object.capabilityDelegation)
        ? VerificationRelationships.fromJSON(object.capabilityDelegation)
        : undefined,
      keyAgreement: isSet(object.keyAgreement) ? VerificationRelationships.fromJSON(object.keyAgreement) : undefined,
      service: isSet(object.service) ? Services.fromJSON(object.service) : undefined,
      alsoKnownAs: Array.isArray(object?.alsoKnownAs)
        ? object.alsoKnownAs.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: DidDocument): unknown {
    const obj: any = {};
    if (message.context) {
      obj.context = message.context.map((e) => e);
    } else {
      obj.context = [];
    }
    message.creator !== undefined && (obj.creator = message.creator);
    message.iD !== undefined && (obj.iD = message.iD);
    if (message.controller) {
      obj.controller = message.controller.map((e) => e);
    } else {
      obj.controller = [];
    }
    message.verificationMethod !== undefined && (obj.verificationMethod = message.verificationMethod
      ? VerificationMethods.toJSON(message.verificationMethod)
      : undefined);
    message.authentication !== undefined && (obj.authentication = message.authentication
      ? VerificationRelationships.toJSON(message.authentication)
      : undefined);
    message.assertionMethod !== undefined && (obj.assertionMethod = message.assertionMethod
      ? VerificationRelationships.toJSON(message.assertionMethod)
      : undefined);
    message.capabilityInvocation !== undefined && (obj.capabilityInvocation = message.capabilityInvocation
      ? VerificationRelationships.toJSON(message.capabilityInvocation)
      : undefined);
    message.capabilityDelegation !== undefined && (obj.capabilityDelegation = message.capabilityDelegation
      ? VerificationRelationships.toJSON(message.capabilityDelegation)
      : undefined);
    message.keyAgreement !== undefined
      && (obj.keyAgreement = message.keyAgreement ? VerificationRelationships.toJSON(message.keyAgreement) : undefined);
    message.service !== undefined && (obj.service = message.service ? Services.toJSON(message.service) : undefined);
    if (message.alsoKnownAs) {
      obj.alsoKnownAs = message.alsoKnownAs.map((e) => e);
    } else {
      obj.alsoKnownAs = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DidDocument>, I>>(object: I): DidDocument {
    const message = createBaseDidDocument();
    message.context = object.context?.map((e) => e) || [];
    message.creator = object.creator ?? "";
    message.iD = object.iD ?? "";
    message.controller = object.controller?.map((e) => e) || [];
    message.verificationMethod = (object.verificationMethod !== undefined && object.verificationMethod !== null)
      ? VerificationMethods.fromPartial(object.verificationMethod)
      : undefined;
    message.authentication = (object.authentication !== undefined && object.authentication !== null)
      ? VerificationRelationships.fromPartial(object.authentication)
      : undefined;
    message.assertionMethod = (object.assertionMethod !== undefined && object.assertionMethod !== null)
      ? VerificationRelationships.fromPartial(object.assertionMethod)
      : undefined;
    message.capabilityInvocation = (object.capabilityInvocation !== undefined && object.capabilityInvocation !== null)
      ? VerificationRelationships.fromPartial(object.capabilityInvocation)
      : undefined;
    message.capabilityDelegation = (object.capabilityDelegation !== undefined && object.capabilityDelegation !== null)
      ? VerificationRelationships.fromPartial(object.capabilityDelegation)
      : undefined;
    message.keyAgreement = (object.keyAgreement !== undefined && object.keyAgreement !== null)
      ? VerificationRelationships.fromPartial(object.keyAgreement)
      : undefined;
    message.service = (object.service !== undefined && object.service !== null)
      ? Services.fromPartial(object.service)
      : undefined;
    message.alsoKnownAs = object.alsoKnownAs?.map((e) => e) || [];
    return message;
  },
};

function createBaseVerificationMethod(): VerificationMethod {
  return { iD: "", type: 0, controller: "", publicKeyJwk: {}, publicKeyMultibase: "" };
}

export const VerificationMethod = {
  encode(message: VerificationMethod, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.iD !== "") {
      writer.uint32(10).string(message.iD);
    }
    if (message.type !== 0) {
      writer.uint32(16).int32(message.type);
    }
    if (message.controller !== "") {
      writer.uint32(26).string(message.controller);
    }
    Object.entries(message.publicKeyJwk).forEach(([key, value]) => {
      VerificationMethod_PublicKeyJwkEntry.encode({ key: key as any, value }, writer.uint32(34).fork()).ldelim();
    });
    if (message.publicKeyMultibase !== "") {
      writer.uint32(42).string(message.publicKeyMultibase);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VerificationMethod {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVerificationMethod();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.iD = reader.string();
          break;
        case 2:
          message.type = reader.int32() as any;
          break;
        case 3:
          message.controller = reader.string();
          break;
        case 4:
          const entry4 = VerificationMethod_PublicKeyJwkEntry.decode(reader, reader.uint32());
          if (entry4.value !== undefined) {
            message.publicKeyJwk[entry4.key] = entry4.value;
          }
          break;
        case 5:
          message.publicKeyMultibase = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VerificationMethod {
    return {
      iD: isSet(object.iD) ? String(object.iD) : "",
      type: isSet(object.type) ? keyTypeFromJSON(object.type) : 0,
      controller: isSet(object.controller) ? String(object.controller) : "",
      publicKeyJwk: isObject(object.publicKeyJwk)
        ? Object.entries(object.publicKeyJwk).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
      publicKeyMultibase: isSet(object.publicKeyMultibase) ? String(object.publicKeyMultibase) : "",
    };
  },

  toJSON(message: VerificationMethod): unknown {
    const obj: any = {};
    message.iD !== undefined && (obj.iD = message.iD);
    message.type !== undefined && (obj.type = keyTypeToJSON(message.type));
    message.controller !== undefined && (obj.controller = message.controller);
    obj.publicKeyJwk = {};
    if (message.publicKeyJwk) {
      Object.entries(message.publicKeyJwk).forEach(([k, v]) => {
        obj.publicKeyJwk[k] = v;
      });
    }
    message.publicKeyMultibase !== undefined && (obj.publicKeyMultibase = message.publicKeyMultibase);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VerificationMethod>, I>>(object: I): VerificationMethod {
    const message = createBaseVerificationMethod();
    message.iD = object.iD ?? "";
    message.type = object.type ?? 0;
    message.controller = object.controller ?? "";
    message.publicKeyJwk = Object.entries(object.publicKeyJwk ?? {}).reduce<{ [key: string]: string }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = String(value);
        }
        return acc;
      },
      {},
    );
    message.publicKeyMultibase = object.publicKeyMultibase ?? "";
    return message;
  },
};

function createBaseVerificationMethod_PublicKeyJwkEntry(): VerificationMethod_PublicKeyJwkEntry {
  return { key: "", value: "" };
}

export const VerificationMethod_PublicKeyJwkEntry = {
  encode(message: VerificationMethod_PublicKeyJwkEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VerificationMethod_PublicKeyJwkEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVerificationMethod_PublicKeyJwkEntry();
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

  fromJSON(object: any): VerificationMethod_PublicKeyJwkEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? String(object.value) : "" };
  },

  toJSON(message: VerificationMethod_PublicKeyJwkEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VerificationMethod_PublicKeyJwkEntry>, I>>(
    object: I,
  ): VerificationMethod_PublicKeyJwkEntry {
    const message = createBaseVerificationMethod_PublicKeyJwkEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseVerificationRelationship(): VerificationRelationship {
  return { verificationMethod: undefined, reference: "" };
}

export const VerificationRelationship = {
  encode(message: VerificationRelationship, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.verificationMethod !== undefined) {
      VerificationMethod.encode(message.verificationMethod, writer.uint32(10).fork()).ldelim();
    }
    if (message.reference !== "") {
      writer.uint32(18).string(message.reference);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VerificationRelationship {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVerificationRelationship();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.verificationMethod = VerificationMethod.decode(reader, reader.uint32());
          break;
        case 2:
          message.reference = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VerificationRelationship {
    return {
      verificationMethod: isSet(object.verificationMethod)
        ? VerificationMethod.fromJSON(object.verificationMethod)
        : undefined,
      reference: isSet(object.reference) ? String(object.reference) : "",
    };
  },

  toJSON(message: VerificationRelationship): unknown {
    const obj: any = {};
    message.verificationMethod !== undefined && (obj.verificationMethod = message.verificationMethod
      ? VerificationMethod.toJSON(message.verificationMethod)
      : undefined);
    message.reference !== undefined && (obj.reference = message.reference);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VerificationRelationship>, I>>(object: I): VerificationRelationship {
    const message = createBaseVerificationRelationship();
    message.verificationMethod = (object.verificationMethod !== undefined && object.verificationMethod !== null)
      ? VerificationMethod.fromPartial(object.verificationMethod)
      : undefined;
    message.reference = object.reference ?? "";
    return message;
  },
};

function createBaseService(): Service {
  return { iD: "", type: "", serviceEndpoint: "", serviceEndpoints: {} };
}

export const Service = {
  encode(message: Service, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.iD !== "") {
      writer.uint32(10).string(message.iD);
    }
    if (message.type !== "") {
      writer.uint32(18).string(message.type);
    }
    if (message.serviceEndpoint !== "") {
      writer.uint32(26).string(message.serviceEndpoint);
    }
    Object.entries(message.serviceEndpoints).forEach(([key, value]) => {
      Service_ServiceEndpointsEntry.encode({ key: key as any, value }, writer.uint32(34).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Service {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseService();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.iD = reader.string();
          break;
        case 2:
          message.type = reader.string();
          break;
        case 3:
          message.serviceEndpoint = reader.string();
          break;
        case 4:
          const entry4 = Service_ServiceEndpointsEntry.decode(reader, reader.uint32());
          if (entry4.value !== undefined) {
            message.serviceEndpoints[entry4.key] = entry4.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Service {
    return {
      iD: isSet(object.iD) ? String(object.iD) : "",
      type: isSet(object.type) ? String(object.type) : "",
      serviceEndpoint: isSet(object.serviceEndpoint) ? String(object.serviceEndpoint) : "",
      serviceEndpoints: isObject(object.serviceEndpoints)
        ? Object.entries(object.serviceEndpoints).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
    };
  },

  toJSON(message: Service): unknown {
    const obj: any = {};
    message.iD !== undefined && (obj.iD = message.iD);
    message.type !== undefined && (obj.type = message.type);
    message.serviceEndpoint !== undefined && (obj.serviceEndpoint = message.serviceEndpoint);
    obj.serviceEndpoints = {};
    if (message.serviceEndpoints) {
      Object.entries(message.serviceEndpoints).forEach(([k, v]) => {
        obj.serviceEndpoints[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Service>, I>>(object: I): Service {
    const message = createBaseService();
    message.iD = object.iD ?? "";
    message.type = object.type ?? "";
    message.serviceEndpoint = object.serviceEndpoint ?? "";
    message.serviceEndpoints = Object.entries(object.serviceEndpoints ?? {}).reduce<{ [key: string]: string }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = String(value);
        }
        return acc;
      },
      {},
    );
    return message;
  },
};

function createBaseService_ServiceEndpointsEntry(): Service_ServiceEndpointsEntry {
  return { key: "", value: "" };
}

export const Service_ServiceEndpointsEntry = {
  encode(message: Service_ServiceEndpointsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Service_ServiceEndpointsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseService_ServiceEndpointsEntry();
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

  fromJSON(object: any): Service_ServiceEndpointsEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? String(object.value) : "" };
  },

  toJSON(message: Service_ServiceEndpointsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Service_ServiceEndpointsEntry>, I>>(
    object: I,
  ): Service_ServiceEndpointsEntry {
    const message = createBaseService_ServiceEndpointsEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseServices(): Services {
  return { data: [] };
}

export const Services = {
  encode(message: Services, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.data) {
      Service.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Services {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseServices();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data.push(Service.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Services {
    return { data: Array.isArray(object?.data) ? object.data.map((e: any) => Service.fromJSON(e)) : [] };
  },

  toJSON(message: Services): unknown {
    const obj: any = {};
    if (message.data) {
      obj.data = message.data.map((e) => e ? Service.toJSON(e) : undefined);
    } else {
      obj.data = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Services>, I>>(object: I): Services {
    const message = createBaseServices();
    message.data = object.data?.map((e) => Service.fromPartial(e)) || [];
    return message;
  },
};

function createBaseVerificationMethods(): VerificationMethods {
  return { data: [] };
}

export const VerificationMethods = {
  encode(message: VerificationMethods, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.data) {
      VerificationMethod.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VerificationMethods {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVerificationMethods();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data.push(VerificationMethod.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VerificationMethods {
    return { data: Array.isArray(object?.data) ? object.data.map((e: any) => VerificationMethod.fromJSON(e)) : [] };
  },

  toJSON(message: VerificationMethods): unknown {
    const obj: any = {};
    if (message.data) {
      obj.data = message.data.map((e) => e ? VerificationMethod.toJSON(e) : undefined);
    } else {
      obj.data = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VerificationMethods>, I>>(object: I): VerificationMethods {
    const message = createBaseVerificationMethods();
    message.data = object.data?.map((e) => VerificationMethod.fromPartial(e)) || [];
    return message;
  },
};

function createBaseVerificationRelationships(): VerificationRelationships {
  return { data: [] };
}

export const VerificationRelationships = {
  encode(message: VerificationRelationships, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.data) {
      VerificationRelationship.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VerificationRelationships {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVerificationRelationships();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data.push(VerificationRelationship.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VerificationRelationships {
    return {
      data: Array.isArray(object?.data) ? object.data.map((e: any) => VerificationRelationship.fromJSON(e)) : [],
    };
  },

  toJSON(message: VerificationRelationships): unknown {
    const obj: any = {};
    if (message.data) {
      obj.data = message.data.map((e) => e ? VerificationRelationship.toJSON(e) : undefined);
    } else {
      obj.data = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VerificationRelationships>, I>>(object: I): VerificationRelationships {
    const message = createBaseVerificationRelationships();
    message.data = object.data?.map((e) => VerificationRelationship.fromPartial(e)) || [];
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
