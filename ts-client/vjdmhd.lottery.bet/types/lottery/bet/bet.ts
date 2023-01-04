/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "vjdmhd.lottery.bet";

export interface Bet {
  id: string;
  lotteryId: string;
  height: number;
  amount: number;
  creator: string;
}

function createBaseBet(): Bet {
  return { id: "", lotteryId: "", height: 0, amount: 0, creator: "" };
}

export const Bet = {
  encode(message: Bet, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.lotteryId !== "") {
      writer.uint32(18).string(message.lotteryId);
    }
    if (message.height !== 0) {
      writer.uint32(24).uint64(message.height);
    }
    if (message.amount !== 0) {
      writer.uint32(32).uint64(message.amount);
    }
    if (message.creator !== "") {
      writer.uint32(42).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Bet {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBet();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.lotteryId = reader.string();
          break;
        case 3:
          message.height = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.amount = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Bet {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      lotteryId: isSet(object.lotteryId) ? String(object.lotteryId) : "",
      height: isSet(object.height) ? Number(object.height) : 0,
      amount: isSet(object.amount) ? Number(object.amount) : 0,
      creator: isSet(object.creator) ? String(object.creator) : "",
    };
  },

  toJSON(message: Bet): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.lotteryId !== undefined && (obj.lotteryId = message.lotteryId);
    message.height !== undefined && (obj.height = Math.round(message.height));
    message.amount !== undefined && (obj.amount = Math.round(message.amount));
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Bet>, I>>(object: I): Bet {
    const message = createBaseBet();
    message.id = object.id ?? "";
    message.lotteryId = object.lotteryId ?? "";
    message.height = object.height ?? 0;
    message.amount = object.amount ?? 0;
    message.creator = object.creator ?? "";
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}