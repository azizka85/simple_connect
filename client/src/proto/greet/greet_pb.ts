// @generated by protoc-gen-es v1.0.0 with parameter "target=ts"
// @generated from file proto/greet/greet.proto (package greet, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message greet.GreetRequest
 */
export class GreetRequest extends Message<GreetRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<GreetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "greet.GreetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GreetRequest {
    return new GreetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GreetRequest {
    return new GreetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GreetRequest {
    return new GreetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GreetRequest | PlainMessage<GreetRequest> | undefined, b: GreetRequest | PlainMessage<GreetRequest> | undefined): boolean {
    return proto3.util.equals(GreetRequest, a, b);
  }
}

/**
 * @generated from message greet.GreetResponse
 */
export class GreetResponse extends Message<GreetResponse> {
  /**
   * @generated from field: string greeting = 1;
   */
  greeting = "";

  constructor(data?: PartialMessage<GreetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "greet.GreetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "greeting", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GreetResponse {
    return new GreetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GreetResponse {
    return new GreetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GreetResponse {
    return new GreetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GreetResponse | PlainMessage<GreetResponse> | undefined, b: GreetResponse | PlainMessage<GreetResponse> | undefined): boolean {
    return proto3.util.equals(GreetResponse, a, b);
  }
}

