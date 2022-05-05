import { txClient, queryClient, MissingWalletError, registry } from "./module";

import { Params } from "./module/types/blob/params";
import { ThereIs } from "./module/types/blob/there_is";

export { Params, ThereIs };

async function initTxClient(vuexGetters) {
  return await txClient(vuexGetters["common/wallet/signer"], {
    addr: vuexGetters["common/env/apiTendermint"],
  });
}

async function initQueryClient(vuexGetters) {
  return await queryClient({
    addr: vuexGetters["common/env/apiCosmos"],
  });
}

function mergeResults(value, next_values) {
  for (let prop of Object.keys(next_values)) {
    if (Array.isArray(next_values[prop])) {
      value[prop] = [...value[prop], ...next_values[prop]];
    } else {
      value[prop] = next_values[prop];
    }
  }
  return value;
}

function getStructure(template) {
  let structure = { fields: [] };
  for (const [key, value] of Object.entries(template)) {
    let field: any = {};
    field.name = key;
    field.type = typeof value;
    structure.fields.push(field);
  }
  return structure;
}

const getDefaultState = () => {
  return {
    Params: {},
    ThereIs: {},
    ThereIsAll: {},

    _Structure: {
      Params: getStructure(Params.fromPartial({})),
      ThereIs: getStructure(ThereIs.fromPartial({})),
    },
    _Registry: registry,
    _Subscriptions: new Set(),
  };
};

// initial state
const state = getDefaultState();

export default {
  namespaced: true,
  state,
  mutations: {
    RESET_STATE(state) {
      Object.assign(state, getDefaultState());
    },
    QUERY(state, { query, key, value }) {
      state[query][JSON.stringify(key)] = value;
    },
    SUBSCRIBE(state, subscription) {
      state._Subscriptions.add(JSON.stringify(subscription));
    },
    UNSUBSCRIBE(state, subscription) {
      state._Subscriptions.delete(JSON.stringify(subscription));
    },
  },
  getters: {
    getParams:
      (state) =>
      (params = { params: {} }) => {
        if (!(<any>params).query) {
          (<any>params).query = null;
        }
        return state.Params[JSON.stringify(params)] ?? {};
      },
    getThereIs:
      (state) =>
      (params = { params: {} }) => {
        if (!(<any>params).query) {
          (<any>params).query = null;
        }
        return state.ThereIs[JSON.stringify(params)] ?? {};
      },
    getThereIsAll:
      (state) =>
      (params = { params: {} }) => {
        if (!(<any>params).query) {
          (<any>params).query = null;
        }
        return state.ThereIsAll[JSON.stringify(params)] ?? {};
      },

    getTypeStructure: (state) => (type) => {
      return state._Structure[type].fields;
    },
    getRegistry: (state) => {
      return state._Registry;
    },
  },
  actions: {
    init({ dispatch, rootGetters }) {
      console.log("Vuex module: sonrio.sonr.blob initialized!");
      if (rootGetters["common/env/client"]) {
        rootGetters["common/env/client"].on("newblock", () => {
          dispatch("StoreUpdate");
        });
      }
    },
    resetState({ commit }) {
      commit("RESET_STATE");
    },
    unsubscribe({ commit }, subscription) {
      commit("UNSUBSCRIBE", subscription);
    },
    async StoreUpdate({ state, dispatch }) {
      state._Subscriptions.forEach(async (subscription) => {
        try {
          const sub = JSON.parse(subscription);
          await dispatch(sub.action, sub.payload);
        } catch (e) {
          throw new Error("Subscriptions: " + e.message);
        }
      });
    },

    async QueryParams(
      { commit, rootGetters, getters },
      {
        options: { subscribe, all } = { subscribe: false, all: false },
        params,
        query = null,
      }
    ) {
      try {
        const key = params ?? {};
        const queryClient = await initQueryClient(rootGetters);
        let value = (await queryClient.queryParams()).data;

        commit("QUERY", {
          query: "Params",
          key: { params: { ...key }, query },
          value,
        });
        if (subscribe)
          commit("SUBSCRIBE", {
            action: "QueryParams",
            payload: { options: { all }, params: { ...key }, query },
          });
        return getters["getParams"]({ params: { ...key }, query }) ?? {};
      } catch (e) {
        throw new Error(
          "QueryClient:QueryParams API Node Unavailable. Could not perform query: " +
            e.message
        );
      }
    },

    async QueryThereIs(
      { commit, rootGetters, getters },
      {
        options: { subscribe, all } = { subscribe: false, all: false },
        params,
        query = null,
      }
    ) {
      try {
        const key = params ?? {};
        const queryClient = await initQueryClient(rootGetters);
        let value = (await queryClient.queryThereIs(key.index)).data;

        commit("QUERY", {
          query: "ThereIs",
          key: { params: { ...key }, query },
          value,
        });
        if (subscribe)
          commit("SUBSCRIBE", {
            action: "QueryThereIs",
            payload: { options: { all }, params: { ...key }, query },
          });
        return getters["getThereIs"]({ params: { ...key }, query }) ?? {};
      } catch (e) {
        throw new Error(
          "QueryClient:QueryThereIs API Node Unavailable. Could not perform query: " +
            e.message
        );
      }
    },

    async QueryThereIsAll(
      { commit, rootGetters, getters },
      {
        options: { subscribe, all } = { subscribe: false, all: false },
        params,
        query = null,
      }
    ) {
      try {
        const key = params ?? {};
        const queryClient = await initQueryClient(rootGetters);
        let value = (await queryClient.queryThereIsAll(query)).data;

        while (
          all &&
          (<any>value).pagination &&
          (<any>value).pagination.next_key != null
        ) {
          let next_values = (
            await queryClient.queryThereIsAll({
              ...query,
              "pagination.key": (<any>value).pagination.next_key,
            })
          ).data;
          value = mergeResults(value, next_values);
        }
        commit("QUERY", {
          query: "ThereIsAll",
          key: { params: { ...key }, query },
          value,
        });
        if (subscribe)
          commit("SUBSCRIBE", {
            action: "QueryThereIsAll",
            payload: { options: { all }, params: { ...key }, query },
          });
        return getters["getThereIsAll"]({ params: { ...key }, query }) ?? {};
      } catch (e) {
        throw new Error(
          "QueryClient:QueryThereIsAll API Node Unavailable. Could not perform query: " +
            e.message
        );
      }
    },

    async sendMsgDownloadBlob({ rootGetters }, { value, fee = [], memo = "" }) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgDownloadBlob(value);
        const result = await txClient.signAndBroadcast([msg], {
          fee: { amount: fee, gas: "200000" },
          memo,
        });
        return result;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgDownloadBlob:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgDownloadBlob:Send Could not broadcast Tx: " + e.message
          );
        }
      }
    },
    async sendMsgUploadBlob({ rootGetters }, { value, fee = [], memo = "" }) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgUploadBlob(value);
        const result = await txClient.signAndBroadcast([msg], {
          fee: { amount: fee, gas: "200000" },
          memo,
        });
        return result;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgUploadBlob:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgUploadBlob:Send Could not broadcast Tx: " + e.message
          );
        }
      }
    },
    async sendMsgCreateThereIs(
      { rootGetters },
      { value, fee = [], memo = "" }
    ) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgCreateThereIs(value);
        const result = await txClient.signAndBroadcast([msg], {
          fee: { amount: fee, gas: "200000" },
          memo,
        });
        return result;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgCreateThereIs:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgCreateThereIs:Send Could not broadcast Tx: " +
              e.message
          );
        }
      }
    },
    async sendMsgUpdateThereIs(
      { rootGetters },
      { value, fee = [], memo = "" }
    ) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgUpdateThereIs(value);
        const result = await txClient.signAndBroadcast([msg], {
          fee: { amount: fee, gas: "200000" },
          memo,
        });
        return result;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgUpdateThereIs:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgUpdateThereIs:Send Could not broadcast Tx: " +
              e.message
          );
        }
      }
    },
    async sendMsgDeleteBlob({ rootGetters }, { value, fee = [], memo = "" }) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgDeleteBlob(value);
        const result = await txClient.signAndBroadcast([msg], {
          fee: { amount: fee, gas: "200000" },
          memo,
        });
        return result;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgDeleteBlob:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgDeleteBlob:Send Could not broadcast Tx: " + e.message
          );
        }
      }
    },
    async sendMsgSyncBlob({ rootGetters }, { value, fee = [], memo = "" }) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgSyncBlob(value);
        const result = await txClient.signAndBroadcast([msg], {
          fee: { amount: fee, gas: "200000" },
          memo,
        });
        return result;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgSyncBlob:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgSyncBlob:Send Could not broadcast Tx: " + e.message
          );
        }
      }
    },
    async sendMsgDeleteThereIs(
      { rootGetters },
      { value, fee = [], memo = "" }
    ) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgDeleteThereIs(value);
        const result = await txClient.signAndBroadcast([msg], {
          fee: { amount: fee, gas: "200000" },
          memo,
        });
        return result;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgDeleteThereIs:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgDeleteThereIs:Send Could not broadcast Tx: " +
              e.message
          );
        }
      }
    },

    async MsgDownloadBlob({ rootGetters }, { value }) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgDownloadBlob(value);
        return msg;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgDownloadBlob:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgDownloadBlob:Create Could not create message: " +
              e.message
          );
        }
      }
    },
    async MsgUploadBlob({ rootGetters }, { value }) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgUploadBlob(value);
        return msg;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgUploadBlob:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgUploadBlob:Create Could not create message: " +
              e.message
          );
        }
      }
    },
    async MsgCreateThereIs({ rootGetters }, { value }) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgCreateThereIs(value);
        return msg;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgCreateThereIs:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgCreateThereIs:Create Could not create message: " +
              e.message
          );
        }
      }
    },
    async MsgUpdateThereIs({ rootGetters }, { value }) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgUpdateThereIs(value);
        return msg;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgUpdateThereIs:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgUpdateThereIs:Create Could not create message: " +
              e.message
          );
        }
      }
    },
    async MsgDeleteBlob({ rootGetters }, { value }) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgDeleteBlob(value);
        return msg;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgDeleteBlob:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgDeleteBlob:Create Could not create message: " +
              e.message
          );
        }
      }
    },
    async MsgSyncBlob({ rootGetters }, { value }) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgSyncBlob(value);
        return msg;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgSyncBlob:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgSyncBlob:Create Could not create message: " + e.message
          );
        }
      }
    },
    async MsgDeleteThereIs({ rootGetters }, { value }) {
      try {
        const txClient = await initTxClient(rootGetters);
        const msg = await txClient.msgDeleteThereIs(value);
        return msg;
      } catch (e) {
        if (e == MissingWalletError) {
          throw new Error(
            "TxClient:MsgDeleteThereIs:Init Could not initialize signing client. Wallet is required."
          );
        } else {
          throw new Error(
            "TxClient:MsgDeleteThereIs:Create Could not create message: " +
              e.message
          );
        }
      }
    },
  },
};
