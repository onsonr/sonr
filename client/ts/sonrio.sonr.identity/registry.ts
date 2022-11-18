import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDeleteDidDocument } from "./types/sonr/identity/tx";
import { MsgCreateDidDocument } from "./types/sonr/identity/tx";
import { MsgUpdateDidDocument } from "./types/sonr/identity/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sonrio.sonr.identity.MsgDeleteDidDocument", MsgDeleteDidDocument],
    ["/sonrio.sonr.identity.MsgCreateDidDocument", MsgCreateDidDocument],
    ["/sonrio.sonr.identity.MsgUpdateDidDocument", MsgUpdateDidDocument],
    
];

export { msgTypes }