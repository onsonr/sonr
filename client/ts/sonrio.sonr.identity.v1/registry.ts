import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateDidDocument } from "./types/sonr/identity/v1/tx";
import { MsgDeleteDidDocument } from "./types/sonr/identity/v1/tx";
import { MsgUpdateDidDocument } from "./types/sonr/identity/v1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sonrio.sonr.identity.v1.MsgCreateDidDocument", MsgCreateDidDocument],
    ["/sonrio.sonr.identity.v1.MsgDeleteDidDocument", MsgDeleteDidDocument],
    ["/sonrio.sonr.identity.v1.MsgUpdateDidDocument", MsgUpdateDidDocument],
    
];

export { msgTypes }