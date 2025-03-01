#
# Autogenerated by Frugal Compiler (3.14.2)
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#



import asyncio
from datetime import timedelta
import inspect

from frugal.aio.processor import FBaseProcessor
from frugal.aio.processor import FProcessorFunction
from frugal.exceptions import TApplicationExceptionType
from frugal.exceptions import TTransportExceptionType
from frugal.middleware import Method
from frugal.transport import TMemoryOutputBuffer
from frugal.util.deprecate import deprecated
from thrift.Thrift import TApplicationException
from thrift.Thrift import TMessageType
from thrift.transport.TTransport import TTransportException
from . import f_BasePinger
from .ttypes import *


class Iface(f_BasePinger.Iface):

    async def ping(self, ctx):
        """
        Args:
            ctx: FContext
        """
        pass


class Client(f_BasePinger.Client, Iface):

    def __init__(self, provider, middleware=None):
        """
        Create a new Client with an FServiceProvider containing a transport
        and protocol factory.

        Args:
            provider: FServiceProvider
            middleware: ServiceMiddleware or list of ServiceMiddleware
        """
        middleware = middleware or []
        if middleware and not isinstance(middleware, list):
            middleware = [middleware]
        super(Client, self).__init__(provider, middleware=middleware)
        middleware += provider.get_middleware()
        self._methods.update({
            'ping': Method(self._ping, middleware),
        })

    async def ping(self, ctx):
        """
        Args:
            ctx: FContext
        """
        return await self._methods['ping']([ctx])

    async def _ping(self, ctx):
        memory_buffer = TMemoryOutputBuffer(self._transport.get_request_size_limit())
        oprot = self._protocol_factory.get_protocol(memory_buffer)
        oprot.write_request_headers(ctx)
        oprot.writeMessageBegin('ping', TMessageType.CALL, 0)
        args = ping_args()
        args.write(oprot)
        oprot.writeMessageEnd()
        response_transport = await self._transport.request(ctx, memory_buffer.getvalue())

        iprot = self._protocol_factory.get_protocol(response_transport)
        iprot.read_response_headers(ctx)
        _, mtype, _ = iprot.readMessageBegin()
        if mtype == TMessageType.EXCEPTION:
            x = TApplicationException()
            x.read(iprot)
            iprot.readMessageEnd()
            if x.type == TApplicationExceptionType.RESPONSE_TOO_LARGE:
                raise TTransportException(type=TTransportExceptionType.RESPONSE_TOO_LARGE, message=x.message)
            raise x
        result = ping_result()
        result.read(iprot)
        iprot.readMessageEnd()

class Processor(f_BasePinger.Processor):

    def __init__(self, handler, middleware=None):
        """
        Create a new Processor.

        Args:
            handler: Iface
        """
        if middleware and not isinstance(middleware, list):
            middleware = [middleware]

        super(Processor, self).__init__(handler, middleware=middleware)
        self.add_to_processor_map('ping', _ping(Method(handler.ping, middleware), self.get_write_lock()))


class _ping(FProcessorFunction):

    def __init__(self, handler, lock):
        super(_ping, self).__init__(handler, lock)

    async def process(self, ctx, iprot, oprot):
        args = ping_args()
        args.read(iprot)
        iprot.readMessageEnd()
        result = ping_result()
        try:
            ret = self._handler([ctx])
            if inspect.iscoroutine(ret):
                ret = await ret
        except TApplicationException as ex:
            async with self._lock:
                _write_application_exception(ctx, oprot, "ping", exception=ex)
                return
        except Exception as e:
            async with self._lock:
                _write_application_exception(ctx, oprot, "ping", ex_code=TApplicationExceptionType.INTERNAL_ERROR, message=str(e))
            raise
        async with self._lock:
            try:
                oprot.write_response_headers(ctx)
                oprot.writeMessageBegin('ping', TMessageType.REPLY, 0)
                result.write(oprot)
                oprot.writeMessageEnd()
                oprot.get_transport().flush()
            except TTransportException as e:
                # catch a request too large error because the TMemoryOutputBuffer always throws that if too much data is written
                if e.type == TTransportExceptionType.REQUEST_TOO_LARGE:
                    raise _write_application_exception(ctx, oprot, "ping", ex_code=TApplicationExceptionType.RESPONSE_TOO_LARGE, message=e.message)
                else:
                    raise e


def _write_application_exception(ctx, oprot, method, ex_code=None, message=None, exception=None):
    if exception is not None:
        x = exception
    else:
        x = TApplicationException(type=ex_code, message=message)
    oprot.write_response_headers(ctx)
    oprot.writeMessageBegin(method, TMessageType.EXCEPTION, 0)
    x.write(oprot)
    oprot.writeMessageEnd()
    oprot.get_transport().flush()
    return x

class ping_args(object):
    def read(self, iprot):
        iprot.readStructBegin()
        while True:
            (fname, ftype, fid) = iprot.readFieldBegin()
            if ftype == TType.STOP:
                break
            else:
                iprot.skip(ftype)
            iprot.readFieldEnd()
        iprot.readStructEnd()
        self.validate()

    def write(self, oprot):
        self.validate()
        oprot.writeStructBegin('ping_args')
        oprot.writeFieldStop()
        oprot.writeStructEnd()

    def validate(self):
        return

    def __hash__(self):
        value = 17
        return value

    def __repr__(self):
        L = ['%s=%r' % (key, value)
            for key, value in self.__dict__.items()]
        return '%s(%s)' % (self.__class__.__name__, ', '.join(L))

    def __eq__(self, other):
        return isinstance(other, self.__class__) and self.__dict__ == other.__dict__

    def __ne__(self, other):
        return not (self == other)

class ping_result(object):
    def read(self, iprot):
        iprot.readStructBegin()
        while True:
            (fname, ftype, fid) = iprot.readFieldBegin()
            if ftype == TType.STOP:
                break
            else:
                iprot.skip(ftype)
            iprot.readFieldEnd()
        iprot.readStructEnd()
        self.validate()

    def write(self, oprot):
        self.validate()
        oprot.writeStructBegin('ping_result')
        oprot.writeFieldStop()
        oprot.writeStructEnd()

    def validate(self):
        return

    def __hash__(self):
        value = 17
        return value

    def __repr__(self):
        L = ['%s=%r' % (key, value)
            for key, value in self.__dict__.items()]
        return '%s(%s)' % (self.__class__.__name__, ', '.join(L))

    def __eq__(self, other):
        return isinstance(other, self.__class__) and self.__dict__ == other.__dict__

    def __ne__(self, other):
        return not (self == other)

