/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "vjdmhd.lottery.bet";

export interface MsgCreateBet {
  creator: string;
  amount: string;
}

export interface MsgCreateBetResponse {
}

function createBaseMsgCreateBet(): MsgCreateBet {
  return { creator: "", amount: "" };
}

export const MsgCreateBet = {
  encode(message: MsgCreateBet, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateBet {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateBet();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateBet {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
    };
  },

  toJSON(message: MsgCreateBet): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateBet>, I>>(object: I): MsgCreateBet {
    const message = createBaseMsgCreateBet();
    message.creator = object.creator ?? "";
    message.amount = object.amount ?? "";
    return message;
  },
};

function createBaseMsgCreateBetResponse(): MsgCreateBetResponse {
  return {};
}

export const MsgCreateBetResponse = {
  encode(_: MsgCreateBetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateBetResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateBetResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateBetResponse {
    return {};
  },

  toJSON(_: MsgCreateBetResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateBetResponse>, I>>(_: I): MsgCreateBetResponse {
    const message = createBaseMsgCreateBetResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateBet(request: MsgCreateBet): Promise<MsgCreateBetResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateBet = this.CreateBet.bind(this);
  }
  CreateBet(request: MsgCreateBet): Promise<MsgCreateBetResponse> {
    const data = MsgCreateBet.encode(request).finish();
    const promise = this.rpc.request("vjdmhd.lottery.bet.Msg", "CreateBet", data);
    return promise.then((data) => MsgCreateBetResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
