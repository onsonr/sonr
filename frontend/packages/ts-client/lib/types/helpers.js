"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getStructure = exports.MissingWalletError = void 0;
exports.MissingWalletError = new Error("wallet is required");
function getStructure(template) {
    let structure = { fields: [] };
    for (const [key, value] of Object.entries(template)) {
        let field = {};
        field.name = key;
        field.type = typeof value;
        structure.fields.push(field);
    }
    return structure;
}
exports.getStructure = getStructure;
