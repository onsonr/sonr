import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateService } from "./types/core/identity/tx";
import { MsgDeactivateService } from "./types/core/identity/tx";
import { MsgUpdateDidDocument } from "./types/core/identity/tx";
import { MsgRegisterService } from "./types/core/identity/tx";
import { MsgDeleteDidDocument } from "./types/core/identity/tx";
import { MsgCreateDidDocument } from "./types/core/identity/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/sonrhq.core.identity.MsgUpdateService", MsgUpdateService],
    ["/sonrhq.core.identity.MsgDeactivateService", MsgDeactivateService],
    ["/sonrhq.core.identity.MsgUpdateDidDocument", MsgUpdateDidDocument],
    ["/sonrhq.core.identity.MsgRegisterService", MsgRegisterService],
    ["/sonrhq.core.identity.MsgDeleteDidDocument", MsgDeleteDidDocument],
    ["/sonrhq.core.identity.MsgCreateDidDocument", MsgCreateDidDocument],
    
];

export { msgTypes }