/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "vjdmhd.lottery.lottery";

/** Params defines the parameters for the module. */
export interface Params {
  lotteryFee: number;
  betSize: BetSize | undefined;
}

export interface BetSize {
  minBet: number;
  maxBet: number;
}

function createBaseParams(): Params {
  return { lotteryFee: 0, betSize: undefined };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.lotteryFee !== 0) {
      writer.uint32(8).uint64(message.lotteryFee);
    }
    if (message.betSize !== undefined) {
      BetSize.encode(message.betSize, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.lotteryFee = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.betSize = BetSize.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    return {
      lotteryFee: isSet(object.lotteryFee) ? Number(object.lotteryFee) : 0,
      betSize: isSet(object.betSize) ? BetSize.fromJSON(object.betSize) : undefined,
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.lotteryFee !== undefined && (obj.lotteryFee = Math.round(message.lotteryFee));
    message.betSize !== undefined && (obj.betSize = message.betSize ? BetSize.toJSON(message.betSize) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.lotteryFee = object.lotteryFee ?? 0;
    message.betSize = (object.betSize !== undefined && object.betSize !== null)
      ? BetSize.fromPartial(object.betSize)
      : undefined;
    return message;
  },
};

function createBaseBetSize(): BetSize {
  return { minBet: 0, maxBet: 0 };
}

export const BetSize = {
  encode(message: BetSize, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.minBet !== 0) {
      writer.uint32(8).uint64(message.minBet);
    }
    if (message.maxBet !== 0) {
      writer.uint32(16).uint64(message.maxBet);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BetSize {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBetSize();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.minBet = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.maxBet = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BetSize {
    return {
      minBet: isSet(object.minBet) ? Number(object.minBet) : 0,
      maxBet: isSet(object.maxBet) ? Number(object.maxBet) : 0,
    };
  },

  toJSON(message: BetSize): unknown {
    const obj: any = {};
    message.minBet !== undefined && (obj.minBet = Math.round(message.minBet));
    message.maxBet !== undefined && (obj.maxBet = Math.round(message.maxBet));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BetSize>, I>>(object: I): BetSize {
    const message = createBaseBetSize();
    message.minBet = object.minBet ?? 0;
    message.maxBet = object.maxBet ?? 0;
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
