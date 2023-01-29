// @generated by protoc-gen-connect-web v0.6.0 with parameter "target=ts"
// @generated from file proto/greet/timing.proto (package greet, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MessageInfo, TimingInputRequest, TimingInputResponse } from "./timing_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service greet.TimingInputService
 */
export const TimingInputService = {
  typeName: "greet.TimingInputService",
  methods: {
    /**
     * @generated from rpc greet.TimingInputService.Send
     */
    send: {
      name: "Send",
      I: TimingInputRequest,
      O: TimingInputResponse,
      kind: MethodKind.ClientStreaming,
    },
  }
} as const;

/**
 * @generated from service greet.TimingOutputService
 */
export const TimingOutputService = {
  typeName: "greet.TimingOutputService",
  methods: {
    /**
     * @generated from rpc greet.TimingOutputService.Accept
     */
    accept: {
      name: "Accept",
      I: TimingInputRequest,
      O: MessageInfo,
      kind: MethodKind.ServerStreaming,
    },
  }
} as const;
