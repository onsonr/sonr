/* eslint-disable */
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import * as Long from "long";
import { MIME, Profile } from "../../common/v1/core";

export const protobufPackage = "common.v1";

export enum Direction {
  DIRECTION_UNSPECIFIED = 0,
  DIRECTION_INCOMING = 1,
  DIRECTION_OUTGOING = 2,
  UNRECOGNIZED = -1,
}

export function directionFromJSON(object: any): Direction {
  switch (object) {
    case 0:
    case "DIRECTION_UNSPECIFIED":
      return Direction.DIRECTION_UNSPECIFIED;
    case 1:
    case "DIRECTION_INCOMING":
      return Direction.DIRECTION_INCOMING;
    case 2:
    case "DIRECTION_OUTGOING":
      return Direction.DIRECTION_OUTGOING;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Direction.UNRECOGNIZED;
  }
}

export function directionToJSON(object: Direction): string {
  switch (object) {
    case Direction.DIRECTION_UNSPECIFIED:
      return "DIRECTION_UNSPECIFIED";
    case Direction.DIRECTION_INCOMING:
      return "DIRECTION_INCOMING";
    case Direction.DIRECTION_OUTGOING:
      return "DIRECTION_OUTGOING";
    default:
      return "UNKNOWN";
  }
}

/** For Transfer File Payload */
export interface FileItem {
  /** Standard Mime Type */
  mime: MIME | undefined;
  /** File Name without Path */
  name: string;
  /** File Location */
  path: string;
  /** File Size in Bytes */
  size: number;
  /** Thumbnail of File */
  thumbnail: Thumbnail | undefined;
  /** Last Modified Time in Seconds */
  lastModified: number;
}

/** For Transfer Text Payload */
export interface MessageItem {
  /** Standard Mime Type */
  mime: MIME | undefined;
  /** Subject of Message */
  subject: string;
  /** Body of Message */
  body: string;
  /** Created Time in Seconds */
  createdAt: number;
  /** Attachments of Message */
  attachments: FileItem[];
}

/** Sonr Url Link Contains metadata of provided URL */
export interface UrlItem {
  /** Standard Mime Type */
  mime: MIME | undefined;
  /** OG URL Link */
  link: string;
  /** Meta Field for Title */
  title: string;
  /** Meta field for site */
  site: string;
  /** Meta field for sitename */
  siteName: string;
  /** Meta field for description */
  description: string;
  /** OpenGraph Object */
  openGraph: OpenGraph | undefined;
}

/** OpenGraph is a generic OpenGraph object */
export interface OpenGraph {
  /** Primary OpenGraph Object */
  primary: OpenGraph_Primary | undefined;
  /** Images */
  images: OpenGraph_Image[];
  /** Videos */
  videos: OpenGraph_Video[];
  /** Audios */
  audios: OpenGraph_Audio[];
  /** Twitter Card */
  twitter: OpenGraph_Twitter | undefined;
}

/** Url Opengraph Preview Type - In order of Priority */
export enum OpenGraph_Type {
  TYPE_UNSPECIFIED = 0,
  /** TYPE_IMAGE - Image Preview */
  TYPE_IMAGE = 1,
  /** TYPE_VIDEO - Video Preview */
  TYPE_VIDEO = 2,
  /** TYPE_TWITTER - Twitter Card Preview */
  TYPE_TWITTER = 3,
  /** TYPE_AUDIO - Audio Preview */
  TYPE_AUDIO = 4,
  /** TYPE_NONE - No Type, Preview not set. */
  TYPE_NONE = 5,
  UNRECOGNIZED = -1,
}

export function openGraph_TypeFromJSON(object: any): OpenGraph_Type {
  switch (object) {
    case 0:
    case "TYPE_UNSPECIFIED":
      return OpenGraph_Type.TYPE_UNSPECIFIED;
    case 1:
    case "TYPE_IMAGE":
      return OpenGraph_Type.TYPE_IMAGE;
    case 2:
    case "TYPE_VIDEO":
      return OpenGraph_Type.TYPE_VIDEO;
    case 3:
    case "TYPE_TWITTER":
      return OpenGraph_Type.TYPE_TWITTER;
    case 4:
    case "TYPE_AUDIO":
      return OpenGraph_Type.TYPE_AUDIO;
    case 5:
    case "TYPE_NONE":
      return OpenGraph_Type.TYPE_NONE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return OpenGraph_Type.UNRECOGNIZED;
  }
}

export function openGraph_TypeToJSON(object: OpenGraph_Type): string {
  switch (object) {
    case OpenGraph_Type.TYPE_UNSPECIFIED:
      return "TYPE_UNSPECIFIED";
    case OpenGraph_Type.TYPE_IMAGE:
      return "TYPE_IMAGE";
    case OpenGraph_Type.TYPE_VIDEO:
      return "TYPE_VIDEO";
    case OpenGraph_Type.TYPE_TWITTER:
      return "TYPE_TWITTER";
    case OpenGraph_Type.TYPE_AUDIO:
      return "TYPE_AUDIO";
    case OpenGraph_Type.TYPE_NONE:
      return "TYPE_NONE";
    default:
      return "UNKNOWN";
  }
}

/** Primary Opengraph Preview */
export interface OpenGraph_Primary {
  /** Type of Primary */
  type: OpenGraph_Type;
  /** Image */
  image: OpenGraph_Image | undefined;
  /** Video */
  video: OpenGraph_Video | undefined;
  /** Audio */
  audio: OpenGraph_Audio | undefined;
  /** Twitter Card */
  twitter: OpenGraph_Twitter | undefined;
}

/** OpenGraph Image */
export interface OpenGraph_Image {
  /** `meta:"og:image,og:image:url"` */
  url: string;
  /** `meta:"og:image:secure_url"` */
  secureUrl: string;
  /** `meta:"og:image:width"` */
  width: number;
  /** `meta:"og:image:height"` */
  height: number;
  /** `meta:"og:image:type"` */
  type: string;
}

/** OpenGraph Video */
export interface OpenGraph_Video {
  /** `meta:"og:video,og:video:url"` */
  url: string;
  /** `meta:"og:video:secure_url"` */
  secureUrl: string;
  /** `meta:"og:video:width"` */
  width: number;
  /** `meta:"og:video:height"` */
  height: number;
  /** `meta:"og:video:type"` */
  type: string;
}

/** OpenGraph Audio */
export interface OpenGraph_Audio {
  /** `meta:"og:audio,og:audio:url"` */
  url: string;
  /** `meta:"og:audio:secure_url"` */
  secureUrl: string;
  /** `meta:"og:audio:type"` */
  type: string;
}

/** Twitter Card */
export interface OpenGraph_Twitter {
  /** `meta:"twitter:card"` */
  card: string;
  /** `meta:"twitter:site"` */
  site: string;
  /** `meta:"twitter:site:id"` */
  siteId: string;
  /** `meta:"twitter:creator"` */
  creator: string;
  /** `meta:"twitter:creator:id"` */
  creatorId: string;
  /** `meta:"twitter:description"` */
  description: string;
  /** `meta:"twitter:title"` */
  title: string;
  /** `meta:"twitter:image,twitter:image:src"` */
  image: string;
  /** `meta:"twitter:image:alt"` */
  imageAlt: string;
  /** `meta:"twitter:url"` */
  url: string;
  /** Twitter Item Player */
  player: OpenGraph_Twitter_Player | undefined;
  /** Twitter iPhone Data */
  iphone: OpenGraph_Twitter_IPhone | undefined;
  /** Twitter iPad Data */
  ipad: OpenGraph_Twitter_IPad | undefined;
  /** Twitter Android Data */
  googlePlay: OpenGraph_Twitter_GooglePlay | undefined;
}

export interface OpenGraph_Twitter_Player {
  /** `meta:"twitter:player"` */
  url: string;
  /** `meta:"twitter:width"` */
  width: number;
  /** `meta:"twitter:height"` */
  height: number;
  /** `meta:"twitter:stream"` */
  stream: string;
}

export interface OpenGraph_Twitter_IPhone {
  /** `meta:"twitter:app:name:iphone"` */
  name: string;
  /** `meta:"twitter:app:id:iphone"` */
  id: string;
  /** `meta:"twitter:app:url:iphone"` */
  url: string;
}

export interface OpenGraph_Twitter_IPad {
  /** `meta:"twitter:app:name:ipad"` */
  name: string;
  /** `meta:"twitter:app:id:ipad"` */
  id: string;
  /** `meta:"twitter:app:url:ipad"` */
  url: string;
}

export interface OpenGraph_Twitter_GooglePlay {
  /** `meta:"twitter:app:name:googleplay"` */
  name: string;
  /** `meta:"twitter:app:id:googleplay"` */
  id: string;
  /** `meta:"twitter:app:url:googleplay"` */
  url: string;
}

/** Thumbnail of File */
export interface Thumbnail {
  /** Thumbnail Buffer */
  buffer: Uint8Array;
  /** Mime Type */
  mime: MIME | undefined;
}

/** Payload is Data thats being Passed */
export interface Payload {
  /** Payload Items */
  items: Payload_Item[];
  /** PROFILE: General Sender Info */
  owner: Profile | undefined;
  /** Payload Size in Bytes */
  size: number;
  /** Payload Creation Time in Seconds */
  createdAt: number;
}

/** Item in Payload */
export interface Payload_Item {
  /** MIME of the Item */
  mime: MIME | undefined;
  /** Size of the Item in Bytes */
  size: number;
  /** FILE: File Item */
  file: FileItem | undefined;
  /** URL: Url Item */
  url: UrlItem | undefined;
  /** MESSAGE: Message Item */
  message: MessageItem | undefined;
  /** Thumbnail of the Item */
  thumbnail: Thumbnail | undefined;
  /** Open Graph Image */
  openGraph: OpenGraph_Primary | undefined;
}

/** PayloadItemList is a list of Payload.Item's for Persistent Store */
export interface PayloadList {
  /** Payload List */
  payloads: Payload[];
  /** Key of the Payload List */
  key: string;
  /** Last Modified Time in Seconds */
  lastModified: number;
}

/** SupplyItem is an item supplied to be a payload */
export interface SupplyItem {
  /** Supply Path */
  path: string;
  /** Supply Path of the Thumbnail */
  thumbnail?: Uint8Array | undefined;
}

function createBaseFileItem(): FileItem {
  return {
    mime: undefined,
    name: "",
    path: "",
    size: 0,
    thumbnail: undefined,
    lastModified: 0,
  };
}

export const FileItem = {
  encode(message: FileItem, writer: Writer = Writer.create()): Writer {
    if (message.mime !== undefined) {
      MIME.encode(message.mime, writer.uint32(10).fork()).ldelim();
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.path !== "") {
      writer.uint32(26).string(message.path);
    }
    if (message.size !== 0) {
      writer.uint32(32).int64(message.size);
    }
    if (message.thumbnail !== undefined) {
      Thumbnail.encode(message.thumbnail, writer.uint32(42).fork()).ldelim();
    }
    if (message.lastModified !== 0) {
      writer.uint32(48).int64(message.lastModified);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): FileItem {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFileItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mime = MIME.decode(reader, reader.uint32());
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.path = reader.string();
          break;
        case 4:
          message.size = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.thumbnail = Thumbnail.decode(reader, reader.uint32());
          break;
        case 6:
          message.lastModified = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FileItem {
    return {
      mime: isSet(object.mime) ? MIME.fromJSON(object.mime) : undefined,
      name: isSet(object.name) ? String(object.name) : "",
      path: isSet(object.path) ? String(object.path) : "",
      size: isSet(object.size) ? Number(object.size) : 0,
      thumbnail: isSet(object.thumbnail)
        ? Thumbnail.fromJSON(object.thumbnail)
        : undefined,
      lastModified: isSet(object.lastModified)
        ? Number(object.lastModified)
        : 0,
    };
  },

  toJSON(message: FileItem): unknown {
    const obj: any = {};
    message.mime !== undefined &&
      (obj.mime = message.mime ? MIME.toJSON(message.mime) : undefined);
    message.name !== undefined && (obj.name = message.name);
    message.path !== undefined && (obj.path = message.path);
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.thumbnail !== undefined &&
      (obj.thumbnail = message.thumbnail
        ? Thumbnail.toJSON(message.thumbnail)
        : undefined);
    message.lastModified !== undefined &&
      (obj.lastModified = Math.round(message.lastModified));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FileItem>, I>>(object: I): FileItem {
    const message = createBaseFileItem();
    message.mime =
      object.mime !== undefined && object.mime !== null
        ? MIME.fromPartial(object.mime)
        : undefined;
    message.name = object.name ?? "";
    message.path = object.path ?? "";
    message.size = object.size ?? 0;
    message.thumbnail =
      object.thumbnail !== undefined && object.thumbnail !== null
        ? Thumbnail.fromPartial(object.thumbnail)
        : undefined;
    message.lastModified = object.lastModified ?? 0;
    return message;
  },
};

function createBaseMessageItem(): MessageItem {
  return {
    mime: undefined,
    subject: "",
    body: "",
    createdAt: 0,
    attachments: [],
  };
}

export const MessageItem = {
  encode(message: MessageItem, writer: Writer = Writer.create()): Writer {
    if (message.mime !== undefined) {
      MIME.encode(message.mime, writer.uint32(10).fork()).ldelim();
    }
    if (message.subject !== "") {
      writer.uint32(18).string(message.subject);
    }
    if (message.body !== "") {
      writer.uint32(26).string(message.body);
    }
    if (message.createdAt !== 0) {
      writer.uint32(32).int64(message.createdAt);
    }
    for (const v of message.attachments) {
      FileItem.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MessageItem {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMessageItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mime = MIME.decode(reader, reader.uint32());
          break;
        case 2:
          message.subject = reader.string();
          break;
        case 3:
          message.body = reader.string();
          break;
        case 4:
          message.createdAt = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.attachments.push(FileItem.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MessageItem {
    return {
      mime: isSet(object.mime) ? MIME.fromJSON(object.mime) : undefined,
      subject: isSet(object.subject) ? String(object.subject) : "",
      body: isSet(object.body) ? String(object.body) : "",
      createdAt: isSet(object.createdAt) ? Number(object.createdAt) : 0,
      attachments: Array.isArray(object?.attachments)
        ? object.attachments.map((e: any) => FileItem.fromJSON(e))
        : [],
    };
  },

  toJSON(message: MessageItem): unknown {
    const obj: any = {};
    message.mime !== undefined &&
      (obj.mime = message.mime ? MIME.toJSON(message.mime) : undefined);
    message.subject !== undefined && (obj.subject = message.subject);
    message.body !== undefined && (obj.body = message.body);
    message.createdAt !== undefined &&
      (obj.createdAt = Math.round(message.createdAt));
    if (message.attachments) {
      obj.attachments = message.attachments.map((e) =>
        e ? FileItem.toJSON(e) : undefined
      );
    } else {
      obj.attachments = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MessageItem>, I>>(
    object: I
  ): MessageItem {
    const message = createBaseMessageItem();
    message.mime =
      object.mime !== undefined && object.mime !== null
        ? MIME.fromPartial(object.mime)
        : undefined;
    message.subject = object.subject ?? "";
    message.body = object.body ?? "";
    message.createdAt = object.createdAt ?? 0;
    message.attachments =
      object.attachments?.map((e) => FileItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseUrlItem(): UrlItem {
  return {
    mime: undefined,
    link: "",
    title: "",
    site: "",
    siteName: "",
    description: "",
    openGraph: undefined,
  };
}

export const UrlItem = {
  encode(message: UrlItem, writer: Writer = Writer.create()): Writer {
    if (message.mime !== undefined) {
      MIME.encode(message.mime, writer.uint32(10).fork()).ldelim();
    }
    if (message.link !== "") {
      writer.uint32(18).string(message.link);
    }
    if (message.title !== "") {
      writer.uint32(26).string(message.title);
    }
    if (message.site !== "") {
      writer.uint32(34).string(message.site);
    }
    if (message.siteName !== "") {
      writer.uint32(42).string(message.siteName);
    }
    if (message.description !== "") {
      writer.uint32(50).string(message.description);
    }
    if (message.openGraph !== undefined) {
      OpenGraph.encode(message.openGraph, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UrlItem {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUrlItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mime = MIME.decode(reader, reader.uint32());
          break;
        case 2:
          message.link = reader.string();
          break;
        case 3:
          message.title = reader.string();
          break;
        case 4:
          message.site = reader.string();
          break;
        case 5:
          message.siteName = reader.string();
          break;
        case 6:
          message.description = reader.string();
          break;
        case 7:
          message.openGraph = OpenGraph.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UrlItem {
    return {
      mime: isSet(object.mime) ? MIME.fromJSON(object.mime) : undefined,
      link: isSet(object.link) ? String(object.link) : "",
      title: isSet(object.title) ? String(object.title) : "",
      site: isSet(object.site) ? String(object.site) : "",
      siteName: isSet(object.siteName) ? String(object.siteName) : "",
      description: isSet(object.description) ? String(object.description) : "",
      openGraph: isSet(object.openGraph)
        ? OpenGraph.fromJSON(object.openGraph)
        : undefined,
    };
  },

  toJSON(message: UrlItem): unknown {
    const obj: any = {};
    message.mime !== undefined &&
      (obj.mime = message.mime ? MIME.toJSON(message.mime) : undefined);
    message.link !== undefined && (obj.link = message.link);
    message.title !== undefined && (obj.title = message.title);
    message.site !== undefined && (obj.site = message.site);
    message.siteName !== undefined && (obj.siteName = message.siteName);
    message.description !== undefined &&
      (obj.description = message.description);
    message.openGraph !== undefined &&
      (obj.openGraph = message.openGraph
        ? OpenGraph.toJSON(message.openGraph)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UrlItem>, I>>(object: I): UrlItem {
    const message = createBaseUrlItem();
    message.mime =
      object.mime !== undefined && object.mime !== null
        ? MIME.fromPartial(object.mime)
        : undefined;
    message.link = object.link ?? "";
    message.title = object.title ?? "";
    message.site = object.site ?? "";
    message.siteName = object.siteName ?? "";
    message.description = object.description ?? "";
    message.openGraph =
      object.openGraph !== undefined && object.openGraph !== null
        ? OpenGraph.fromPartial(object.openGraph)
        : undefined;
    return message;
  },
};

function createBaseOpenGraph(): OpenGraph {
  return {
    primary: undefined,
    images: [],
    videos: [],
    audios: [],
    twitter: undefined,
  };
}

export const OpenGraph = {
  encode(message: OpenGraph, writer: Writer = Writer.create()): Writer {
    if (message.primary !== undefined) {
      OpenGraph_Primary.encode(
        message.primary,
        writer.uint32(10).fork()
      ).ldelim();
    }
    for (const v of message.images) {
      OpenGraph_Image.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.videos) {
      OpenGraph_Video.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.audios) {
      OpenGraph_Audio.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.twitter !== undefined) {
      OpenGraph_Twitter.encode(
        message.twitter,
        writer.uint32(42).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OpenGraph {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenGraph();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.primary = OpenGraph_Primary.decode(reader, reader.uint32());
          break;
        case 2:
          message.images.push(OpenGraph_Image.decode(reader, reader.uint32()));
          break;
        case 3:
          message.videos.push(OpenGraph_Video.decode(reader, reader.uint32()));
          break;
        case 4:
          message.audios.push(OpenGraph_Audio.decode(reader, reader.uint32()));
          break;
        case 5:
          message.twitter = OpenGraph_Twitter.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OpenGraph {
    return {
      primary: isSet(object.primary)
        ? OpenGraph_Primary.fromJSON(object.primary)
        : undefined,
      images: Array.isArray(object?.images)
        ? object.images.map((e: any) => OpenGraph_Image.fromJSON(e))
        : [],
      videos: Array.isArray(object?.videos)
        ? object.videos.map((e: any) => OpenGraph_Video.fromJSON(e))
        : [],
      audios: Array.isArray(object?.audios)
        ? object.audios.map((e: any) => OpenGraph_Audio.fromJSON(e))
        : [],
      twitter: isSet(object.twitter)
        ? OpenGraph_Twitter.fromJSON(object.twitter)
        : undefined,
    };
  },

  toJSON(message: OpenGraph): unknown {
    const obj: any = {};
    message.primary !== undefined &&
      (obj.primary = message.primary
        ? OpenGraph_Primary.toJSON(message.primary)
        : undefined);
    if (message.images) {
      obj.images = message.images.map((e) =>
        e ? OpenGraph_Image.toJSON(e) : undefined
      );
    } else {
      obj.images = [];
    }
    if (message.videos) {
      obj.videos = message.videos.map((e) =>
        e ? OpenGraph_Video.toJSON(e) : undefined
      );
    } else {
      obj.videos = [];
    }
    if (message.audios) {
      obj.audios = message.audios.map((e) =>
        e ? OpenGraph_Audio.toJSON(e) : undefined
      );
    } else {
      obj.audios = [];
    }
    message.twitter !== undefined &&
      (obj.twitter = message.twitter
        ? OpenGraph_Twitter.toJSON(message.twitter)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OpenGraph>, I>>(
    object: I
  ): OpenGraph {
    const message = createBaseOpenGraph();
    message.primary =
      object.primary !== undefined && object.primary !== null
        ? OpenGraph_Primary.fromPartial(object.primary)
        : undefined;
    message.images =
      object.images?.map((e) => OpenGraph_Image.fromPartial(e)) || [];
    message.videos =
      object.videos?.map((e) => OpenGraph_Video.fromPartial(e)) || [];
    message.audios =
      object.audios?.map((e) => OpenGraph_Audio.fromPartial(e)) || [];
    message.twitter =
      object.twitter !== undefined && object.twitter !== null
        ? OpenGraph_Twitter.fromPartial(object.twitter)
        : undefined;
    return message;
  },
};

function createBaseOpenGraph_Primary(): OpenGraph_Primary {
  return {
    type: 0,
    image: undefined,
    video: undefined,
    audio: undefined,
    twitter: undefined,
  };
}

export const OpenGraph_Primary = {
  encode(message: OpenGraph_Primary, writer: Writer = Writer.create()): Writer {
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.image !== undefined) {
      OpenGraph_Image.encode(message.image, writer.uint32(18).fork()).ldelim();
    }
    if (message.video !== undefined) {
      OpenGraph_Video.encode(message.video, writer.uint32(26).fork()).ldelim();
    }
    if (message.audio !== undefined) {
      OpenGraph_Audio.encode(message.audio, writer.uint32(34).fork()).ldelim();
    }
    if (message.twitter !== undefined) {
      OpenGraph_Twitter.encode(
        message.twitter,
        writer.uint32(42).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OpenGraph_Primary {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenGraph_Primary();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.int32() as any;
          break;
        case 2:
          message.image = OpenGraph_Image.decode(reader, reader.uint32());
          break;
        case 3:
          message.video = OpenGraph_Video.decode(reader, reader.uint32());
          break;
        case 4:
          message.audio = OpenGraph_Audio.decode(reader, reader.uint32());
          break;
        case 5:
          message.twitter = OpenGraph_Twitter.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OpenGraph_Primary {
    return {
      type: isSet(object.type) ? openGraph_TypeFromJSON(object.type) : 0,
      image: isSet(object.image)
        ? OpenGraph_Image.fromJSON(object.image)
        : undefined,
      video: isSet(object.video)
        ? OpenGraph_Video.fromJSON(object.video)
        : undefined,
      audio: isSet(object.audio)
        ? OpenGraph_Audio.fromJSON(object.audio)
        : undefined,
      twitter: isSet(object.twitter)
        ? OpenGraph_Twitter.fromJSON(object.twitter)
        : undefined,
    };
  },

  toJSON(message: OpenGraph_Primary): unknown {
    const obj: any = {};
    message.type !== undefined &&
      (obj.type = openGraph_TypeToJSON(message.type));
    message.image !== undefined &&
      (obj.image = message.image
        ? OpenGraph_Image.toJSON(message.image)
        : undefined);
    message.video !== undefined &&
      (obj.video = message.video
        ? OpenGraph_Video.toJSON(message.video)
        : undefined);
    message.audio !== undefined &&
      (obj.audio = message.audio
        ? OpenGraph_Audio.toJSON(message.audio)
        : undefined);
    message.twitter !== undefined &&
      (obj.twitter = message.twitter
        ? OpenGraph_Twitter.toJSON(message.twitter)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OpenGraph_Primary>, I>>(
    object: I
  ): OpenGraph_Primary {
    const message = createBaseOpenGraph_Primary();
    message.type = object.type ?? 0;
    message.image =
      object.image !== undefined && object.image !== null
        ? OpenGraph_Image.fromPartial(object.image)
        : undefined;
    message.video =
      object.video !== undefined && object.video !== null
        ? OpenGraph_Video.fromPartial(object.video)
        : undefined;
    message.audio =
      object.audio !== undefined && object.audio !== null
        ? OpenGraph_Audio.fromPartial(object.audio)
        : undefined;
    message.twitter =
      object.twitter !== undefined && object.twitter !== null
        ? OpenGraph_Twitter.fromPartial(object.twitter)
        : undefined;
    return message;
  },
};

function createBaseOpenGraph_Image(): OpenGraph_Image {
  return { url: "", secureUrl: "", width: 0, height: 0, type: "" };
}

export const OpenGraph_Image = {
  encode(message: OpenGraph_Image, writer: Writer = Writer.create()): Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    if (message.secureUrl !== "") {
      writer.uint32(18).string(message.secureUrl);
    }
    if (message.width !== 0) {
      writer.uint32(24).int32(message.width);
    }
    if (message.height !== 0) {
      writer.uint32(32).int32(message.height);
    }
    if (message.type !== "") {
      writer.uint32(42).string(message.type);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OpenGraph_Image {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenGraph_Image();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.url = reader.string();
          break;
        case 2:
          message.secureUrl = reader.string();
          break;
        case 3:
          message.width = reader.int32();
          break;
        case 4:
          message.height = reader.int32();
          break;
        case 5:
          message.type = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OpenGraph_Image {
    return {
      url: isSet(object.url) ? String(object.url) : "",
      secureUrl: isSet(object.secureUrl) ? String(object.secureUrl) : "",
      width: isSet(object.width) ? Number(object.width) : 0,
      height: isSet(object.height) ? Number(object.height) : 0,
      type: isSet(object.type) ? String(object.type) : "",
    };
  },

  toJSON(message: OpenGraph_Image): unknown {
    const obj: any = {};
    message.url !== undefined && (obj.url = message.url);
    message.secureUrl !== undefined && (obj.secureUrl = message.secureUrl);
    message.width !== undefined && (obj.width = Math.round(message.width));
    message.height !== undefined && (obj.height = Math.round(message.height));
    message.type !== undefined && (obj.type = message.type);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OpenGraph_Image>, I>>(
    object: I
  ): OpenGraph_Image {
    const message = createBaseOpenGraph_Image();
    message.url = object.url ?? "";
    message.secureUrl = object.secureUrl ?? "";
    message.width = object.width ?? 0;
    message.height = object.height ?? 0;
    message.type = object.type ?? "";
    return message;
  },
};

function createBaseOpenGraph_Video(): OpenGraph_Video {
  return { url: "", secureUrl: "", width: 0, height: 0, type: "" };
}

export const OpenGraph_Video = {
  encode(message: OpenGraph_Video, writer: Writer = Writer.create()): Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    if (message.secureUrl !== "") {
      writer.uint32(18).string(message.secureUrl);
    }
    if (message.width !== 0) {
      writer.uint32(24).int32(message.width);
    }
    if (message.height !== 0) {
      writer.uint32(32).int32(message.height);
    }
    if (message.type !== "") {
      writer.uint32(42).string(message.type);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OpenGraph_Video {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenGraph_Video();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.url = reader.string();
          break;
        case 2:
          message.secureUrl = reader.string();
          break;
        case 3:
          message.width = reader.int32();
          break;
        case 4:
          message.height = reader.int32();
          break;
        case 5:
          message.type = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OpenGraph_Video {
    return {
      url: isSet(object.url) ? String(object.url) : "",
      secureUrl: isSet(object.secureUrl) ? String(object.secureUrl) : "",
      width: isSet(object.width) ? Number(object.width) : 0,
      height: isSet(object.height) ? Number(object.height) : 0,
      type: isSet(object.type) ? String(object.type) : "",
    };
  },

  toJSON(message: OpenGraph_Video): unknown {
    const obj: any = {};
    message.url !== undefined && (obj.url = message.url);
    message.secureUrl !== undefined && (obj.secureUrl = message.secureUrl);
    message.width !== undefined && (obj.width = Math.round(message.width));
    message.height !== undefined && (obj.height = Math.round(message.height));
    message.type !== undefined && (obj.type = message.type);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OpenGraph_Video>, I>>(
    object: I
  ): OpenGraph_Video {
    const message = createBaseOpenGraph_Video();
    message.url = object.url ?? "";
    message.secureUrl = object.secureUrl ?? "";
    message.width = object.width ?? 0;
    message.height = object.height ?? 0;
    message.type = object.type ?? "";
    return message;
  },
};

function createBaseOpenGraph_Audio(): OpenGraph_Audio {
  return { url: "", secureUrl: "", type: "" };
}

export const OpenGraph_Audio = {
  encode(message: OpenGraph_Audio, writer: Writer = Writer.create()): Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    if (message.secureUrl !== "") {
      writer.uint32(18).string(message.secureUrl);
    }
    if (message.type !== "") {
      writer.uint32(26).string(message.type);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OpenGraph_Audio {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenGraph_Audio();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.url = reader.string();
          break;
        case 2:
          message.secureUrl = reader.string();
          break;
        case 3:
          message.type = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OpenGraph_Audio {
    return {
      url: isSet(object.url) ? String(object.url) : "",
      secureUrl: isSet(object.secureUrl) ? String(object.secureUrl) : "",
      type: isSet(object.type) ? String(object.type) : "",
    };
  },

  toJSON(message: OpenGraph_Audio): unknown {
    const obj: any = {};
    message.url !== undefined && (obj.url = message.url);
    message.secureUrl !== undefined && (obj.secureUrl = message.secureUrl);
    message.type !== undefined && (obj.type = message.type);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OpenGraph_Audio>, I>>(
    object: I
  ): OpenGraph_Audio {
    const message = createBaseOpenGraph_Audio();
    message.url = object.url ?? "";
    message.secureUrl = object.secureUrl ?? "";
    message.type = object.type ?? "";
    return message;
  },
};

function createBaseOpenGraph_Twitter(): OpenGraph_Twitter {
  return {
    card: "",
    site: "",
    siteId: "",
    creator: "",
    creatorId: "",
    description: "",
    title: "",
    image: "",
    imageAlt: "",
    url: "",
    player: undefined,
    iphone: undefined,
    ipad: undefined,
    googlePlay: undefined,
  };
}

export const OpenGraph_Twitter = {
  encode(message: OpenGraph_Twitter, writer: Writer = Writer.create()): Writer {
    if (message.card !== "") {
      writer.uint32(10).string(message.card);
    }
    if (message.site !== "") {
      writer.uint32(18).string(message.site);
    }
    if (message.siteId !== "") {
      writer.uint32(26).string(message.siteId);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    if (message.creatorId !== "") {
      writer.uint32(42).string(message.creatorId);
    }
    if (message.description !== "") {
      writer.uint32(50).string(message.description);
    }
    if (message.title !== "") {
      writer.uint32(58).string(message.title);
    }
    if (message.image !== "") {
      writer.uint32(66).string(message.image);
    }
    if (message.imageAlt !== "") {
      writer.uint32(74).string(message.imageAlt);
    }
    if (message.url !== "") {
      writer.uint32(82).string(message.url);
    }
    if (message.player !== undefined) {
      OpenGraph_Twitter_Player.encode(
        message.player,
        writer.uint32(90).fork()
      ).ldelim();
    }
    if (message.iphone !== undefined) {
      OpenGraph_Twitter_IPhone.encode(
        message.iphone,
        writer.uint32(98).fork()
      ).ldelim();
    }
    if (message.ipad !== undefined) {
      OpenGraph_Twitter_IPad.encode(
        message.ipad,
        writer.uint32(106).fork()
      ).ldelim();
    }
    if (message.googlePlay !== undefined) {
      OpenGraph_Twitter_GooglePlay.encode(
        message.googlePlay,
        writer.uint32(114).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OpenGraph_Twitter {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenGraph_Twitter();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.card = reader.string();
          break;
        case 2:
          message.site = reader.string();
          break;
        case 3:
          message.siteId = reader.string();
          break;
        case 4:
          message.creator = reader.string();
          break;
        case 5:
          message.creatorId = reader.string();
          break;
        case 6:
          message.description = reader.string();
          break;
        case 7:
          message.title = reader.string();
          break;
        case 8:
          message.image = reader.string();
          break;
        case 9:
          message.imageAlt = reader.string();
          break;
        case 10:
          message.url = reader.string();
          break;
        case 11:
          message.player = OpenGraph_Twitter_Player.decode(
            reader,
            reader.uint32()
          );
          break;
        case 12:
          message.iphone = OpenGraph_Twitter_IPhone.decode(
            reader,
            reader.uint32()
          );
          break;
        case 13:
          message.ipad = OpenGraph_Twitter_IPad.decode(reader, reader.uint32());
          break;
        case 14:
          message.googlePlay = OpenGraph_Twitter_GooglePlay.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OpenGraph_Twitter {
    return {
      card: isSet(object.card) ? String(object.card) : "",
      site: isSet(object.site) ? String(object.site) : "",
      siteId: isSet(object.siteId) ? String(object.siteId) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      creatorId: isSet(object.creatorId) ? String(object.creatorId) : "",
      description: isSet(object.description) ? String(object.description) : "",
      title: isSet(object.title) ? String(object.title) : "",
      image: isSet(object.image) ? String(object.image) : "",
      imageAlt: isSet(object.imageAlt) ? String(object.imageAlt) : "",
      url: isSet(object.url) ? String(object.url) : "",
      player: isSet(object.player)
        ? OpenGraph_Twitter_Player.fromJSON(object.player)
        : undefined,
      iphone: isSet(object.iphone)
        ? OpenGraph_Twitter_IPhone.fromJSON(object.iphone)
        : undefined,
      ipad: isSet(object.ipad)
        ? OpenGraph_Twitter_IPad.fromJSON(object.ipad)
        : undefined,
      googlePlay: isSet(object.googlePlay)
        ? OpenGraph_Twitter_GooglePlay.fromJSON(object.googlePlay)
        : undefined,
    };
  },

  toJSON(message: OpenGraph_Twitter): unknown {
    const obj: any = {};
    message.card !== undefined && (obj.card = message.card);
    message.site !== undefined && (obj.site = message.site);
    message.siteId !== undefined && (obj.siteId = message.siteId);
    message.creator !== undefined && (obj.creator = message.creator);
    message.creatorId !== undefined && (obj.creatorId = message.creatorId);
    message.description !== undefined &&
      (obj.description = message.description);
    message.title !== undefined && (obj.title = message.title);
    message.image !== undefined && (obj.image = message.image);
    message.imageAlt !== undefined && (obj.imageAlt = message.imageAlt);
    message.url !== undefined && (obj.url = message.url);
    message.player !== undefined &&
      (obj.player = message.player
        ? OpenGraph_Twitter_Player.toJSON(message.player)
        : undefined);
    message.iphone !== undefined &&
      (obj.iphone = message.iphone
        ? OpenGraph_Twitter_IPhone.toJSON(message.iphone)
        : undefined);
    message.ipad !== undefined &&
      (obj.ipad = message.ipad
        ? OpenGraph_Twitter_IPad.toJSON(message.ipad)
        : undefined);
    message.googlePlay !== undefined &&
      (obj.googlePlay = message.googlePlay
        ? OpenGraph_Twitter_GooglePlay.toJSON(message.googlePlay)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OpenGraph_Twitter>, I>>(
    object: I
  ): OpenGraph_Twitter {
    const message = createBaseOpenGraph_Twitter();
    message.card = object.card ?? "";
    message.site = object.site ?? "";
    message.siteId = object.siteId ?? "";
    message.creator = object.creator ?? "";
    message.creatorId = object.creatorId ?? "";
    message.description = object.description ?? "";
    message.title = object.title ?? "";
    message.image = object.image ?? "";
    message.imageAlt = object.imageAlt ?? "";
    message.url = object.url ?? "";
    message.player =
      object.player !== undefined && object.player !== null
        ? OpenGraph_Twitter_Player.fromPartial(object.player)
        : undefined;
    message.iphone =
      object.iphone !== undefined && object.iphone !== null
        ? OpenGraph_Twitter_IPhone.fromPartial(object.iphone)
        : undefined;
    message.ipad =
      object.ipad !== undefined && object.ipad !== null
        ? OpenGraph_Twitter_IPad.fromPartial(object.ipad)
        : undefined;
    message.googlePlay =
      object.googlePlay !== undefined && object.googlePlay !== null
        ? OpenGraph_Twitter_GooglePlay.fromPartial(object.googlePlay)
        : undefined;
    return message;
  },
};

function createBaseOpenGraph_Twitter_Player(): OpenGraph_Twitter_Player {
  return { url: "", width: 0, height: 0, stream: "" };
}

export const OpenGraph_Twitter_Player = {
  encode(
    message: OpenGraph_Twitter_Player,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    if (message.width !== 0) {
      writer.uint32(16).int32(message.width);
    }
    if (message.height !== 0) {
      writer.uint32(24).int32(message.height);
    }
    if (message.stream !== "") {
      writer.uint32(34).string(message.stream);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): OpenGraph_Twitter_Player {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenGraph_Twitter_Player();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.url = reader.string();
          break;
        case 2:
          message.width = reader.int32();
          break;
        case 3:
          message.height = reader.int32();
          break;
        case 4:
          message.stream = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OpenGraph_Twitter_Player {
    return {
      url: isSet(object.url) ? String(object.url) : "",
      width: isSet(object.width) ? Number(object.width) : 0,
      height: isSet(object.height) ? Number(object.height) : 0,
      stream: isSet(object.stream) ? String(object.stream) : "",
    };
  },

  toJSON(message: OpenGraph_Twitter_Player): unknown {
    const obj: any = {};
    message.url !== undefined && (obj.url = message.url);
    message.width !== undefined && (obj.width = Math.round(message.width));
    message.height !== undefined && (obj.height = Math.round(message.height));
    message.stream !== undefined && (obj.stream = message.stream);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OpenGraph_Twitter_Player>, I>>(
    object: I
  ): OpenGraph_Twitter_Player {
    const message = createBaseOpenGraph_Twitter_Player();
    message.url = object.url ?? "";
    message.width = object.width ?? 0;
    message.height = object.height ?? 0;
    message.stream = object.stream ?? "";
    return message;
  },
};

function createBaseOpenGraph_Twitter_IPhone(): OpenGraph_Twitter_IPhone {
  return { name: "", id: "", url: "" };
}

export const OpenGraph_Twitter_IPhone = {
  encode(
    message: OpenGraph_Twitter_IPhone,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.url !== "") {
      writer.uint32(26).string(message.url);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): OpenGraph_Twitter_IPhone {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenGraph_Twitter_IPhone();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.url = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OpenGraph_Twitter_IPhone {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      id: isSet(object.id) ? String(object.id) : "",
      url: isSet(object.url) ? String(object.url) : "",
    };
  },

  toJSON(message: OpenGraph_Twitter_IPhone): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.id !== undefined && (obj.id = message.id);
    message.url !== undefined && (obj.url = message.url);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OpenGraph_Twitter_IPhone>, I>>(
    object: I
  ): OpenGraph_Twitter_IPhone {
    const message = createBaseOpenGraph_Twitter_IPhone();
    message.name = object.name ?? "";
    message.id = object.id ?? "";
    message.url = object.url ?? "";
    return message;
  },
};

function createBaseOpenGraph_Twitter_IPad(): OpenGraph_Twitter_IPad {
  return { name: "", id: "", url: "" };
}

export const OpenGraph_Twitter_IPad = {
  encode(
    message: OpenGraph_Twitter_IPad,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.url !== "") {
      writer.uint32(26).string(message.url);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OpenGraph_Twitter_IPad {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenGraph_Twitter_IPad();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.url = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OpenGraph_Twitter_IPad {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      id: isSet(object.id) ? String(object.id) : "",
      url: isSet(object.url) ? String(object.url) : "",
    };
  },

  toJSON(message: OpenGraph_Twitter_IPad): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.id !== undefined && (obj.id = message.id);
    message.url !== undefined && (obj.url = message.url);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OpenGraph_Twitter_IPad>, I>>(
    object: I
  ): OpenGraph_Twitter_IPad {
    const message = createBaseOpenGraph_Twitter_IPad();
    message.name = object.name ?? "";
    message.id = object.id ?? "";
    message.url = object.url ?? "";
    return message;
  },
};

function createBaseOpenGraph_Twitter_GooglePlay(): OpenGraph_Twitter_GooglePlay {
  return { name: "", id: "", url: "" };
}

export const OpenGraph_Twitter_GooglePlay = {
  encode(
    message: OpenGraph_Twitter_GooglePlay,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.url !== "") {
      writer.uint32(26).string(message.url);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): OpenGraph_Twitter_GooglePlay {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenGraph_Twitter_GooglePlay();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.url = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OpenGraph_Twitter_GooglePlay {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      id: isSet(object.id) ? String(object.id) : "",
      url: isSet(object.url) ? String(object.url) : "",
    };
  },

  toJSON(message: OpenGraph_Twitter_GooglePlay): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.id !== undefined && (obj.id = message.id);
    message.url !== undefined && (obj.url = message.url);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OpenGraph_Twitter_GooglePlay>, I>>(
    object: I
  ): OpenGraph_Twitter_GooglePlay {
    const message = createBaseOpenGraph_Twitter_GooglePlay();
    message.name = object.name ?? "";
    message.id = object.id ?? "";
    message.url = object.url ?? "";
    return message;
  },
};

function createBaseThumbnail(): Thumbnail {
  return { buffer: new Uint8Array(), mime: undefined };
}

export const Thumbnail = {
  encode(message: Thumbnail, writer: Writer = Writer.create()): Writer {
    if (message.buffer.length !== 0) {
      writer.uint32(10).bytes(message.buffer);
    }
    if (message.mime !== undefined) {
      MIME.encode(message.mime, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Thumbnail {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseThumbnail();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.buffer = reader.bytes();
          break;
        case 2:
          message.mime = MIME.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Thumbnail {
    return {
      buffer: isSet(object.buffer)
        ? bytesFromBase64(object.buffer)
        : new Uint8Array(),
      mime: isSet(object.mime) ? MIME.fromJSON(object.mime) : undefined,
    };
  },

  toJSON(message: Thumbnail): unknown {
    const obj: any = {};
    message.buffer !== undefined &&
      (obj.buffer = base64FromBytes(
        message.buffer !== undefined ? message.buffer : new Uint8Array()
      ));
    message.mime !== undefined &&
      (obj.mime = message.mime ? MIME.toJSON(message.mime) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Thumbnail>, I>>(
    object: I
  ): Thumbnail {
    const message = createBaseThumbnail();
    message.buffer = object.buffer ?? new Uint8Array();
    message.mime =
      object.mime !== undefined && object.mime !== null
        ? MIME.fromPartial(object.mime)
        : undefined;
    return message;
  },
};

function createBasePayload(): Payload {
  return { items: [], owner: undefined, size: 0, createdAt: 0 };
}

export const Payload = {
  encode(message: Payload, writer: Writer = Writer.create()): Writer {
    for (const v of message.items) {
      Payload_Item.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.owner !== undefined) {
      Profile.encode(message.owner, writer.uint32(18).fork()).ldelim();
    }
    if (message.size !== 0) {
      writer.uint32(24).int64(message.size);
    }
    if (message.createdAt !== 0) {
      writer.uint32(32).int64(message.createdAt);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Payload {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePayload();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(Payload_Item.decode(reader, reader.uint32()));
          break;
        case 2:
          message.owner = Profile.decode(reader, reader.uint32());
          break;
        case 3:
          message.size = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.createdAt = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Payload {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => Payload_Item.fromJSON(e))
        : [],
      owner: isSet(object.owner) ? Profile.fromJSON(object.owner) : undefined,
      size: isSet(object.size) ? Number(object.size) : 0,
      createdAt: isSet(object.createdAt) ? Number(object.createdAt) : 0,
    };
  },

  toJSON(message: Payload): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? Payload_Item.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.owner !== undefined &&
      (obj.owner = message.owner ? Profile.toJSON(message.owner) : undefined);
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.createdAt !== undefined &&
      (obj.createdAt = Math.round(message.createdAt));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Payload>, I>>(object: I): Payload {
    const message = createBasePayload();
    message.items = object.items?.map((e) => Payload_Item.fromPartial(e)) || [];
    message.owner =
      object.owner !== undefined && object.owner !== null
        ? Profile.fromPartial(object.owner)
        : undefined;
    message.size = object.size ?? 0;
    message.createdAt = object.createdAt ?? 0;
    return message;
  },
};

function createBasePayload_Item(): Payload_Item {
  return {
    mime: undefined,
    size: 0,
    file: undefined,
    url: undefined,
    message: undefined,
    thumbnail: undefined,
    openGraph: undefined,
  };
}

export const Payload_Item = {
  encode(message: Payload_Item, writer: Writer = Writer.create()): Writer {
    if (message.mime !== undefined) {
      MIME.encode(message.mime, writer.uint32(10).fork()).ldelim();
    }
    if (message.size !== 0) {
      writer.uint32(16).int64(message.size);
    }
    if (message.file !== undefined) {
      FileItem.encode(message.file, writer.uint32(26).fork()).ldelim();
    }
    if (message.url !== undefined) {
      UrlItem.encode(message.url, writer.uint32(34).fork()).ldelim();
    }
    if (message.message !== undefined) {
      MessageItem.encode(message.message, writer.uint32(42).fork()).ldelim();
    }
    if (message.thumbnail !== undefined) {
      Thumbnail.encode(message.thumbnail, writer.uint32(50).fork()).ldelim();
    }
    if (message.openGraph !== undefined) {
      OpenGraph_Primary.encode(
        message.openGraph,
        writer.uint32(58).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Payload_Item {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePayload_Item();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mime = MIME.decode(reader, reader.uint32());
          break;
        case 2:
          message.size = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.file = FileItem.decode(reader, reader.uint32());
          break;
        case 4:
          message.url = UrlItem.decode(reader, reader.uint32());
          break;
        case 5:
          message.message = MessageItem.decode(reader, reader.uint32());
          break;
        case 6:
          message.thumbnail = Thumbnail.decode(reader, reader.uint32());
          break;
        case 7:
          message.openGraph = OpenGraph_Primary.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Payload_Item {
    return {
      mime: isSet(object.mime) ? MIME.fromJSON(object.mime) : undefined,
      size: isSet(object.size) ? Number(object.size) : 0,
      file: isSet(object.file) ? FileItem.fromJSON(object.file) : undefined,
      url: isSet(object.url) ? UrlItem.fromJSON(object.url) : undefined,
      message: isSet(object.message)
        ? MessageItem.fromJSON(object.message)
        : undefined,
      thumbnail: isSet(object.thumbnail)
        ? Thumbnail.fromJSON(object.thumbnail)
        : undefined,
      openGraph: isSet(object.openGraph)
        ? OpenGraph_Primary.fromJSON(object.openGraph)
        : undefined,
    };
  },

  toJSON(message: Payload_Item): unknown {
    const obj: any = {};
    message.mime !== undefined &&
      (obj.mime = message.mime ? MIME.toJSON(message.mime) : undefined);
    message.size !== undefined && (obj.size = Math.round(message.size));
    message.file !== undefined &&
      (obj.file = message.file ? FileItem.toJSON(message.file) : undefined);
    message.url !== undefined &&
      (obj.url = message.url ? UrlItem.toJSON(message.url) : undefined);
    message.message !== undefined &&
      (obj.message = message.message
        ? MessageItem.toJSON(message.message)
        : undefined);
    message.thumbnail !== undefined &&
      (obj.thumbnail = message.thumbnail
        ? Thumbnail.toJSON(message.thumbnail)
        : undefined);
    message.openGraph !== undefined &&
      (obj.openGraph = message.openGraph
        ? OpenGraph_Primary.toJSON(message.openGraph)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Payload_Item>, I>>(
    object: I
  ): Payload_Item {
    const message = createBasePayload_Item();
    message.mime =
      object.mime !== undefined && object.mime !== null
        ? MIME.fromPartial(object.mime)
        : undefined;
    message.size = object.size ?? 0;
    message.file =
      object.file !== undefined && object.file !== null
        ? FileItem.fromPartial(object.file)
        : undefined;
    message.url =
      object.url !== undefined && object.url !== null
        ? UrlItem.fromPartial(object.url)
        : undefined;
    message.message =
      object.message !== undefined && object.message !== null
        ? MessageItem.fromPartial(object.message)
        : undefined;
    message.thumbnail =
      object.thumbnail !== undefined && object.thumbnail !== null
        ? Thumbnail.fromPartial(object.thumbnail)
        : undefined;
    message.openGraph =
      object.openGraph !== undefined && object.openGraph !== null
        ? OpenGraph_Primary.fromPartial(object.openGraph)
        : undefined;
    return message;
  },
};

function createBasePayloadList(): PayloadList {
  return { payloads: [], key: "", lastModified: 0 };
}

export const PayloadList = {
  encode(message: PayloadList, writer: Writer = Writer.create()): Writer {
    for (const v of message.payloads) {
      Payload.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    if (message.lastModified !== 0) {
      writer.uint32(24).int64(message.lastModified);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PayloadList {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePayloadList();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.payloads.push(Payload.decode(reader, reader.uint32()));
          break;
        case 2:
          message.key = reader.string();
          break;
        case 3:
          message.lastModified = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PayloadList {
    return {
      payloads: Array.isArray(object?.payloads)
        ? object.payloads.map((e: any) => Payload.fromJSON(e))
        : [],
      key: isSet(object.key) ? String(object.key) : "",
      lastModified: isSet(object.lastModified)
        ? Number(object.lastModified)
        : 0,
    };
  },

  toJSON(message: PayloadList): unknown {
    const obj: any = {};
    if (message.payloads) {
      obj.payloads = message.payloads.map((e) =>
        e ? Payload.toJSON(e) : undefined
      );
    } else {
      obj.payloads = [];
    }
    message.key !== undefined && (obj.key = message.key);
    message.lastModified !== undefined &&
      (obj.lastModified = Math.round(message.lastModified));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PayloadList>, I>>(
    object: I
  ): PayloadList {
    const message = createBasePayloadList();
    message.payloads =
      object.payloads?.map((e) => Payload.fromPartial(e)) || [];
    message.key = object.key ?? "";
    message.lastModified = object.lastModified ?? 0;
    return message;
  },
};

function createBaseSupplyItem(): SupplyItem {
  return { path: "", thumbnail: undefined };
}

export const SupplyItem = {
  encode(message: SupplyItem, writer: Writer = Writer.create()): Writer {
    if (message.path !== "") {
      writer.uint32(10).string(message.path);
    }
    if (message.thumbnail !== undefined) {
      writer.uint32(18).bytes(message.thumbnail);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SupplyItem {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSupplyItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.path = reader.string();
          break;
        case 2:
          message.thumbnail = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SupplyItem {
    return {
      path: isSet(object.path) ? String(object.path) : "",
      thumbnail: isSet(object.thumbnail)
        ? bytesFromBase64(object.thumbnail)
        : undefined,
    };
  },

  toJSON(message: SupplyItem): unknown {
    const obj: any = {};
    message.path !== undefined && (obj.path = message.path);
    message.thumbnail !== undefined &&
      (obj.thumbnail =
        message.thumbnail !== undefined
          ? base64FromBytes(message.thumbnail)
          : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SupplyItem>, I>>(
    object: I
  ): SupplyItem {
    const message = createBaseSupplyItem();
    message.path = object.path ?? "";
    message.thumbnail = object.thumbnail ?? undefined;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
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
  for (const byte of arr) {
    bin.push(String.fromCharCode(byte));
  }
  return btoa(bin.join(""));
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
