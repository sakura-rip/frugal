/**
 * Autogenerated by Frugal Compiler (3.14.2)
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *
 * @generated
 */

package include_vendor.java;

import com.workiva.frugal.FContext;
import com.workiva.frugal.exception.TApplicationExceptionType;
import com.workiva.frugal.middleware.InvocationHandler;
import com.workiva.frugal.middleware.ServiceMiddleware;
import com.workiva.frugal.protocol.*;
import com.workiva.frugal.provider.FScopeProvider;
import com.workiva.frugal.transport.FPublisherTransport;
import com.workiva.frugal.transport.FSubscriberTransport;
import com.workiva.frugal.transport.FSubscription;
import com.workiva.frugal.transport.TMemoryOutputBuffer;
import org.apache.thrift.TException;
import org.apache.thrift.TApplicationException;
import org.apache.thrift.transport.TTransport;
import org.apache.thrift.transport.TTransportException;
import org.apache.thrift.protocol.*;

import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.EnumMap;
import java.util.Set;
import java.util.HashSet;
import java.util.EnumSet;
import java.util.Collections;
import java.util.BitSet;
import java.nio.ByteBuffer;
import java.util.Arrays;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;




public class MyScopeSubscriber {

	public interface Iface {
		public FSubscription subscribenewItem(final newItemHandler handler) throws TException;

	}

	public interface IfaceThrowable {
		public FSubscription subscribenewItemThrowable(final newItemThrowableHandler handler) throws TException;

	}

	public interface newItemHandler {
		void onnewItem(FContext ctx, some.vendored.pkg.Item req) throws TException;
	}

	public interface newItemThrowableHandler {
		void onnewItem(FContext ctx, some.vendored.pkg.Item req) throws TException;
	}

	public static class Client implements Iface, IfaceThrowable {
		private static final String DELIMITER = ".";
		private static final Logger LOGGER = LoggerFactory.getLogger(Client.class);

		private final FScopeProvider provider;
		private final ServiceMiddleware[] middleware;

		public Client(FScopeProvider provider, ServiceMiddleware... middleware) {
			this.provider = provider;
			List<ServiceMiddleware> combined = new ArrayList<ServiceMiddleware>(Arrays.asList(middleware));
			combined.addAll(provider.getMiddleware());
			this.middleware = combined.toArray(new ServiceMiddleware[0]);
		}

		public FSubscription subscribenewItem(final newItemHandler handler) throws TException {
			final String op = "newItem";
			String prefix = "";
			final String topic = String.format("%sMyScope%s%s", prefix, DELIMITER, op);
			final FScopeProvider.Subscriber subscriber = provider.buildSubscriber();
			final FSubscriberTransport transport = subscriber.getTransport();
			final newItemHandler proxiedHandler = InvocationHandler.composeMiddleware(handler, newItemHandler.class, middleware);
			transport.subscribe(topic, recvnewItem(op, subscriber.getProtocolFactory(), proxiedHandler));
			return FSubscription.of(topic, transport);
		}

		private FAsyncCallback recvnewItem(String op, FProtocolFactory pf, newItemHandler handler) {
			return new FAsyncCallback() {
				public void onMessage(TTransport tr) throws TException {
					FProtocol iprot = pf.getProtocol(tr);
					FContext ctx = iprot.readRequestHeader();
					TMessage msg = iprot.readMessageBegin();
					if (!msg.name.equals(op)) {
						TProtocolUtil.skip(iprot, TType.STRUCT);
						iprot.readMessageEnd();
						throw new TApplicationException(TApplicationExceptionType.UNKNOWN_METHOD);
					}
					some.vendored.pkg.Item received = new some.vendored.pkg.Item();
					received.read(iprot);
					iprot.readMessageEnd();
					handler.onnewItem(ctx, received);
				}
			};
		}

		public FSubscription subscribenewItemThrowable(final newItemThrowableHandler handler) throws TException {
			final String op = "newItem";
			String prefix = "";
			final String topic = String.format("%sMyScope%s%s", prefix, DELIMITER, op);
			final FScopeProvider.Subscriber subscriber = provider.buildSubscriber();
			final FSubscriberTransport transport = subscriber.getTransport();
			final newItemThrowableHandler proxiedHandler = InvocationHandler.composeMiddleware(handler, newItemThrowableHandler.class, middleware);
			transport.subscribe(topic, recvnewItem(op, subscriber.getProtocolFactory(), proxiedHandler));
			return FSubscription.of(topic, transport);
		}

		private FAsyncCallback recvnewItem(String op, FProtocolFactory pf, newItemThrowableHandler handler) {
			return new FAsyncCallback() {
				public void onMessage(TTransport tr) throws TException {
					FProtocol iprot = pf.getProtocol(tr);
					FContext ctx = iprot.readRequestHeader();
					TMessage msg = iprot.readMessageBegin();
					if (!msg.name.equals(op)) {
						TProtocolUtil.skip(iprot, TType.STRUCT);
						iprot.readMessageEnd();
						throw new TApplicationException(TApplicationExceptionType.UNKNOWN_METHOD);
					}
					some.vendored.pkg.Item received = new some.vendored.pkg.Item();
					received.read(iprot);
					iprot.readMessageEnd();
					handler.onnewItem(ctx, received);
				}
			};
		}
	}

}
