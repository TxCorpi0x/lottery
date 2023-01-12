/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "vjdmhd.lottery.lottery";

/** Params defines the parameters for the module. */
export interface Params {
  lotteryFee: string;
  betSize: BetSize | undefined;
}

export interface BetSize {
  minBet: string;
  maxBet: string;
}

function createBaseParams(): Params {
  return { lotteryFee: "", betSize: undefined };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.lotteryFee !== "") {
      writer.uint32(10).string(message.lotteryFee);
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
          message.lotteryFee = reader.string();
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
      lotteryFee: isSet(object.lotteryFee) ? String(object.lotteryFee) : "",
      betSize: isSet(object.betSize) ? BetSize.fromJSON(object.betSize) : undefined,
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.lotteryFee !== undefined && (obj.lotteryFee = message.lotteryFee);
    message.betSize !== undefined && (obj.betSize = message.betSize ? BetSize.toJSON(message.betSize) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.lotteryFee = object.lotteryFee ?? "";
    message.betSize = (object.betSize !== undefined && object.betSize !== null)
      ? BetSize.fromPartial(object.betSize)
      : undefined;
    return message;
  },
};

function createBaseBetSize(): BetSize {
  return { minBet: "", maxBet: "" };
}

export const BetSize = {
  encode(message: BetSize, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.minBet !== "") {
      writer.uint32(10).string(message.minBet);
    }
    if (message.maxBet !== "") {
      writer.uint32(18).string(message.maxBet);
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
          message.minBet = reader.string();
          break;
        case 2:
          message.maxBet = reader.string();
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
      minBet: isSet(object.minBet) ? String(object.minBet) : "",
      maxBet: isSet(object.maxBet) ? String(object.maxBet) : "",
    };
  },

  toJSON(message: BetSize): unknown {
    const obj: any = {};
    message.minBet !== undefined && (obj.minBet = message.minBet);
    message.maxBet !== undefined && (obj.maxBet = message.maxBet);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BetSize>, I>>(object: I): BetSize {
    const message = createBaseBetSize();
    message.minBet = object.minBet ?? "";
    message.maxBet = object.maxBet ?? "";
    return message;
  },
};

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
