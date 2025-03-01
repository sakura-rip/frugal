// Autogenerated by Frugal Compiler (3.14.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING



// ignore_for_file: unused_import
// ignore_for_file: unused_field
import 'dart:async';
import 'dart:typed_data' show Uint8List;

import 'package:thrift/thrift.dart' as thrift;
import 'package:frugal/frugal.dart' as frugal;

import 'package:v1_music/v1_music.dart' as t_v1_music;


const String delimiter = '.';

AlbumWinnersPublisher albumWinnersPublisherFactory(frugal.FScopeProvider provider, {List<frugal.Middleware> middleware}) =>
    AlbumWinnersPublisher(provider, middleware);

/// Scopes are a Frugal extension to the IDL for declaring PubSub
/// semantics. Subscribers to this scope will be notified if they win a contest.
/// Scopes must have a prefix.
class AlbumWinnersPublisher {
  frugal.FPublisherTransport transport;
  frugal.FProtocolFactory protocolFactory;
  Map<String, frugal.FMethod> _methods;
  AlbumWinnersPublisher(frugal.FScopeProvider provider, [List<frugal.Middleware> middleware]) {
    transport = provider.publisherTransportFactory.getTransport();
    protocolFactory = provider.protocolFactory;
    var combined = middleware ?? [];
    combined.addAll(provider.middleware);
    this._methods = {};
    this._methods['ContestStart'] = frugal.FMethod(this._publishContestStart, 'AlbumWinners', 'publishContestStart', combined);
    this._methods['TimeLeft'] = frugal.FMethod(this._publishTimeLeft, 'AlbumWinners', 'publishTimeLeft', combined);
    this._methods['Winner'] = frugal.FMethod(this._publishWinner, 'AlbumWinners', 'publishWinner', combined);
  }

  Future open() {
    return transport.open();
  }

  Future close() {
    return transport.close();
  }

  Future publishContestStart(frugal.FContext ctx, List<t_v1_music.Album> req) {
    return this._methods['ContestStart']([ctx, req]);
  }

  Future _publishContestStart(frugal.FContext ctx, List<t_v1_music.Album> req) async {
    var op = 'ContestStart';
    var prefix = 'v1.music.';
    var topic = '${prefix}AlbumWinners$delimiter$op';
    var memoryBuffer = frugal.TMemoryOutputBuffer(transport.publishSizeLimit);
    var oprot = protocolFactory.getProtocol(memoryBuffer);
    var msg = thrift.TMessage(op, thrift.TMessageType.CALL, 0);
    oprot.writeRequestHeader(ctx);
    oprot.writeMessageBegin(msg);
    oprot.writeListBegin(thrift.TList(thrift.TType.STRUCT, req.length));
    for(var elem4 in req) {
      elem4.write(oprot);
    }
    oprot.writeListEnd();
    oprot.writeMessageEnd();
    // sync in this version but async in v2. Mitigate breaking changes by always awaiting.
    // ignore: await_only_futures
    await transport.publish(topic, memoryBuffer.writeBytes);
  }


  Future publishTimeLeft(frugal.FContext ctx, double req) {
    return this._methods['TimeLeft']([ctx, req]);
  }

  Future _publishTimeLeft(frugal.FContext ctx, double req) async {
    var op = 'TimeLeft';
    var prefix = 'v1.music.';
    var topic = '${prefix}AlbumWinners$delimiter$op';
    var memoryBuffer = frugal.TMemoryOutputBuffer(transport.publishSizeLimit);
    var oprot = protocolFactory.getProtocol(memoryBuffer);
    var msg = thrift.TMessage(op, thrift.TMessageType.CALL, 0);
    oprot.writeRequestHeader(ctx);
    oprot.writeMessageBegin(msg);
    oprot.writeDouble(req);
    oprot.writeMessageEnd();
    // sync in this version but async in v2. Mitigate breaking changes by always awaiting.
    // ignore: await_only_futures
    await transport.publish(topic, memoryBuffer.writeBytes);
  }


  Future publishWinner(frugal.FContext ctx, t_v1_music.Album req) {
    return this._methods['Winner']([ctx, req]);
  }

  Future _publishWinner(frugal.FContext ctx, t_v1_music.Album req) async {
    var op = 'Winner';
    var prefix = 'v1.music.';
    var topic = '${prefix}AlbumWinners$delimiter$op';
    var memoryBuffer = frugal.TMemoryOutputBuffer(transport.publishSizeLimit);
    var oprot = protocolFactory.getProtocol(memoryBuffer);
    var msg = thrift.TMessage(op, thrift.TMessageType.CALL, 0);
    oprot.writeRequestHeader(ctx);
    oprot.writeMessageBegin(msg);
    req.write(oprot);
    oprot.writeMessageEnd();
    // sync in this version but async in v2. Mitigate breaking changes by always awaiting.
    // ignore: await_only_futures
    await transport.publish(topic, memoryBuffer.writeBytes);
  }
}


AlbumWinnersSubscriber albumWinnersSubscriberFactory(frugal.FScopeProvider provider, {List<frugal.Middleware> middleware}) =>
    AlbumWinnersSubscriber(provider, middleware);

/// Scopes are a Frugal extension to the IDL for declaring PubSub
/// semantics. Subscribers to this scope will be notified if they win a contest.
/// Scopes must have a prefix.
class AlbumWinnersSubscriber {
  final frugal.FScopeProvider provider;
  final List<frugal.Middleware> _middleware;

  AlbumWinnersSubscriber(this.provider, [List<frugal.Middleware> middleware])
      : this._middleware = middleware ?? [] {
    this._middleware.addAll(provider.middleware);
}

  Future<frugal.FSubscription> subscribeContestStart(dynamic onlist(frugal.FContext ctx, List<t_v1_music.Album> req)) async {
    var op = 'ContestStart';
    var prefix = 'v1.music.';
    var topic = '${prefix}AlbumWinners$delimiter$op';
    var transport = provider.subscriberTransportFactory.getTransport();
    await transport.subscribe(topic, _recvContestStart(op, provider.protocolFactory, onlist));
    return frugal.FSubscription(topic, transport);
  }

  frugal.FAsyncCallback _recvContestStart(String op, frugal.FProtocolFactory protocolFactory, dynamic onlist(frugal.FContext ctx, List<t_v1_music.Album> req)) {
    frugal.FMethod method = frugal.FMethod(onlist, 'AlbumWinners', 'subscribelist', this._middleware);
    callbackContestStart(thrift.TTransport transport) {
      var iprot = protocolFactory.getProtocol(transport);
      var ctx = iprot.readRequestHeader();
      var tMsg = iprot.readMessageBegin();
      if (tMsg.name != op) {
        thrift.TProtocolUtil.skip(iprot, thrift.TType.STRUCT);
        iprot.readMessageEnd();
        throw thrift.TApplicationError(
        frugal.FrugalTApplicationErrorType.UNKNOWN_METHOD, tMsg.name);
      }
      thrift.TList elem5 = iprot.readListBegin();
      List<t_v1_music.Album> req = List<t_v1_music.Album>();
      for(int elem7 = 0; elem7 < elem5.length; ++elem7) {
        t_v1_music.Album elem6 = t_v1_music.Album();
        elem6.read(iprot);
        req.add(elem6);
      }
      iprot.readListEnd();
      iprot.readMessageEnd();
      method([ctx, req]);
    }
    return callbackContestStart;
  }


  Future<frugal.FSubscription> subscribeTimeLeft(dynamic onMinutes(frugal.FContext ctx, double req)) async {
    var op = 'TimeLeft';
    var prefix = 'v1.music.';
    var topic = '${prefix}AlbumWinners$delimiter$op';
    var transport = provider.subscriberTransportFactory.getTransport();
    await transport.subscribe(topic, _recvTimeLeft(op, provider.protocolFactory, onMinutes));
    return frugal.FSubscription(topic, transport);
  }

  frugal.FAsyncCallback _recvTimeLeft(String op, frugal.FProtocolFactory protocolFactory, dynamic onMinutes(frugal.FContext ctx, double req)) {
    frugal.FMethod method = frugal.FMethod(onMinutes, 'AlbumWinners', 'subscribeMinutes', this._middleware);
    callbackTimeLeft(thrift.TTransport transport) {
      var iprot = protocolFactory.getProtocol(transport);
      var ctx = iprot.readRequestHeader();
      var tMsg = iprot.readMessageBegin();
      if (tMsg.name != op) {
        thrift.TProtocolUtil.skip(iprot, thrift.TType.STRUCT);
        iprot.readMessageEnd();
        throw thrift.TApplicationError(
        frugal.FrugalTApplicationErrorType.UNKNOWN_METHOD, tMsg.name);
      }
      double req = iprot.readDouble();
      iprot.readMessageEnd();
      method([ctx, req]);
    }
    return callbackTimeLeft;
  }


  Future<frugal.FSubscription> subscribeWinner(dynamic onAlbum(frugal.FContext ctx, t_v1_music.Album req)) async {
    var op = 'Winner';
    var prefix = 'v1.music.';
    var topic = '${prefix}AlbumWinners$delimiter$op';
    var transport = provider.subscriberTransportFactory.getTransport();
    await transport.subscribe(topic, _recvWinner(op, provider.protocolFactory, onAlbum));
    return frugal.FSubscription(topic, transport);
  }

  frugal.FAsyncCallback _recvWinner(String op, frugal.FProtocolFactory protocolFactory, dynamic onAlbum(frugal.FContext ctx, t_v1_music.Album req)) {
    frugal.FMethod method = frugal.FMethod(onAlbum, 'AlbumWinners', 'subscribeAlbum', this._middleware);
    callbackWinner(thrift.TTransport transport) {
      var iprot = protocolFactory.getProtocol(transport);
      var ctx = iprot.readRequestHeader();
      var tMsg = iprot.readMessageBegin();
      if (tMsg.name != op) {
        thrift.TProtocolUtil.skip(iprot, thrift.TType.STRUCT);
        iprot.readMessageEnd();
        throw thrift.TApplicationError(
        frugal.FrugalTApplicationErrorType.UNKNOWN_METHOD, tMsg.name);
      }
      t_v1_music.Album req = t_v1_music.Album();
      req.read(iprot);
      iprot.readMessageEnd();
      method([ctx, req]);
    }
    return callbackWinner;
  }
}

