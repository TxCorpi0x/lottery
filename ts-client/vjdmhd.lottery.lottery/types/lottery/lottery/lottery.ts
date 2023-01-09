/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "vjdmhd.lottery.lottery";

export interface Lottery {
  id: number;
  startBlock: number;
  endBlock: number;
  betCount: number;
  winnerId: string;
  payout: Coin | undefined;
}

function createBaseLottery(): Lottery {
  return { id: 0, startBlock: 0, endBlock: 0, betCount: 0, winnerId: "", payout: undefined };
}

export const Lottery = {
  encode(message: Lottery, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.startBlock !== 0) {
      writer.uint32(16).int64(message.startBlock);
    }
    if (message.endBlock !== 0) {
      writer.uint32(24).int64(message.endBlock);
    }
    if (message.betCount !== 0) {
      writer.uint32(32).uint64(message.betCount);
    }
    if (message.winnerId !== "") {
      writer.uint32(42).string(message.winnerId);
    }
    if (message.payout !== undefined) {
      Coin.encode(message.payout, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Lottery {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLottery();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.startBlock = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.endBlock = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.betCount = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.winnerId = reader.string();
          break;
        case 6:
          message.payout = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Lottery {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      startBlock: isSet(object.startBlock) ? Number(object.startBlock) : 0,
      endBlock: isSet(object.endBlock) ? Number(object.endBlock) : 0,
      betCount: isSet(object.betCount) ? Number(object.betCount) : 0,
      winnerId: isSet(object.winnerId) ? String(object.winnerId) : "",
      payout: isSet(object.payout) ? Coin.fromJSON(object.payout) : undefined,
    };
  },

  toJSON(message: Lottery): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.startBlock !== undefined && (obj.startBlock = Math.round(message.startBlock));
    message.endBlock !== undefined && (obj.endBlock = Math.round(message.endBlock));
    message.betCount !== undefined && (obj.betCount = Math.round(message.betCount));
    message.winnerId !== undefined && (obj.winnerId = message.winnerId);
    message.payout !== undefined && (obj.payout = message.payout ? Coin.toJSON(message.payout) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Lottery>, I>>(object: I): Lottery {
    const message = createBaseLottery();
    message.id = object.id ?? 0;
    message.startBlock = object.startBlock ?? 0;
    message.endBlock = object.endBlock ?? 0;
    message.betCount = object.betCount ?? 0;
    message.winnerId = object.winnerId ?? "";
    message.payout = (object.payout !== undefined && object.payout !== null)
      ? Coin.fromPartial(object.payout)
      : undefined;
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
