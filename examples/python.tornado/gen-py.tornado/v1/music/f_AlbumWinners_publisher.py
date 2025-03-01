#
# Autogenerated by Frugal Compiler (3.14.2)
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#



import sys
import traceback

from thrift.Thrift import TApplicationException
from thrift.Thrift import TMessageType
from thrift.Thrift import TType
from tornado import gen
from frugal.exceptions import TApplicationExceptionType
from frugal.middleware import Method
from frugal.subscription import FSubscription
from frugal.transport import TMemoryOutputBuffer

from .ttypes import *




class AlbumWinnersPublisher(object):
    """
    Scopes are a Frugal extension to the IDL for declaring PubSub
    semantics. Subscribers to this scope will be notified if they win a contest.
    Scopes must have a prefix.
    """

    _DELIMITER = '.'

    def __init__(self, provider, middleware=None):
        """
        Create a new AlbumWinnersPublisher.

        Args:
            provider: FScopeProvider
            middleware: ServiceMiddleware or list of ServiceMiddleware
        """

        middleware = middleware or []
        if middleware and not isinstance(middleware, list):
            middleware = [middleware]
        middleware += provider.get_middleware()
        self._transport, self._protocol_factory = provider.new_publisher()
        self._methods = {
            'publish_ContestStart': Method(self._publish_ContestStart, middleware),
            'publish_TimeLeft': Method(self._publish_TimeLeft, middleware),
            'publish_Winner': Method(self._publish_Winner, middleware),
        }

    @gen.coroutine
    def open(self):
        yield self._transport.open()

    @gen.coroutine
    def close(self):
        yield self._transport.close()

    @gen.coroutine
    def publish_ContestStart(self, ctx, req):
        """
        Args:
            ctx: FContext
            req: list
        """
        yield self._methods['publish_ContestStart']([ctx, req])

    @gen.coroutine
    def _publish_ContestStart(self, ctx, req):
        op = 'ContestStart'
        prefix = 'v1.music.'
        topic = '{}AlbumWinners{}{}'.format(prefix, self._DELIMITER, op)
        buffer = TMemoryOutputBuffer(self._transport.get_publish_size_limit())
        oprot = self._protocol_factory.get_protocol(buffer)
        oprot.write_request_headers(ctx)
        oprot.writeMessageBegin(op, TMessageType.CALL, 0)
        oprot.writeListBegin(TType.STRUCT, len(req))
        for elem3 in req:
            elem3.write(oprot)
        oprot.writeListEnd()
        oprot.writeMessageEnd()
        yield self._transport.publish(topic, buffer.getvalue())


    @gen.coroutine
    def publish_TimeLeft(self, ctx, req):
        """
        Args:
            ctx: FContext
            req: Minutes
        """
        yield self._methods['publish_TimeLeft']([ctx, req])

    @gen.coroutine
    def _publish_TimeLeft(self, ctx, req):
        op = 'TimeLeft'
        prefix = 'v1.music.'
        topic = '{}AlbumWinners{}{}'.format(prefix, self._DELIMITER, op)
        buffer = TMemoryOutputBuffer(self._transport.get_publish_size_limit())
        oprot = self._protocol_factory.get_protocol(buffer)
        oprot.write_request_headers(ctx)
        oprot.writeMessageBegin(op, TMessageType.CALL, 0)
        oprot.writeDouble(req)
        oprot.writeMessageEnd()
        yield self._transport.publish(topic, buffer.getvalue())


    @gen.coroutine
    def publish_Winner(self, ctx, req):
        """
        Args:
            ctx: FContext
            req: Album
        """
        yield self._methods['publish_Winner']([ctx, req])

    @gen.coroutine
    def _publish_Winner(self, ctx, req):
        op = 'Winner'
        prefix = 'v1.music.'
        topic = '{}AlbumWinners{}{}'.format(prefix, self._DELIMITER, op)
        buffer = TMemoryOutputBuffer(self._transport.get_publish_size_limit())
        oprot = self._protocol_factory.get_protocol(buffer)
        oprot.write_request_headers(ctx)
        oprot.writeMessageBegin(op, TMessageType.CALL, 0)
        req.write(oprot)
        oprot.writeMessageEnd()
        yield self._transport.publish(topic, buffer.getvalue())

