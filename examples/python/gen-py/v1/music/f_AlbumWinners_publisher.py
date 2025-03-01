#
# Autogenerated by Frugal Compiler (3.14.2)
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#



from thrift.Thrift import TMessageType
from frugal.middleware import Method
from frugal.transport import TMemoryOutputBuffer




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

    def open(self):
        self._transport.open()

    def close(self):
        self._transport.close()

    def publish_ContestStart(self, ctx, req):
        """
        Args:
            ctx: FContext
            req: list
        """
        self._methods['publish_ContestStart']([ctx, req])

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
        self._transport.publish(topic, buffer.getvalue())


    def publish_TimeLeft(self, ctx, req):
        """
        Args:
            ctx: FContext
            req: Minutes
        """
        self._methods['publish_TimeLeft']([ctx, req])

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
        self._transport.publish(topic, buffer.getvalue())


    def publish_Winner(self, ctx, req):
        """
        Args:
            ctx: FContext
            req: Album
        """
        self._methods['publish_Winner']([ctx, req])

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
        self._transport.publish(topic, buffer.getvalue())

