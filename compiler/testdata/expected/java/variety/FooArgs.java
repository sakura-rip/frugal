/**
 * Autogenerated by Frugal Compiler (3.14.2)
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *
 * @generated
 */
package variety.java;

import org.apache.thrift.scheme.IScheme;
import org.apache.thrift.scheme.SchemeFactory;
import org.apache.thrift.scheme.StandardScheme;

import org.apache.thrift.scheme.TupleScheme;
import org.apache.thrift.protocol.TTupleProtocol;
import org.apache.thrift.protocol.TProtocolException;
import org.apache.thrift.EncodingUtils;
import org.apache.thrift.TException;
import org.apache.thrift.async.AsyncMethodCallback;
import org.apache.thrift.server.AbstractNonblockingServer.*;
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
import java.util.Objects;
import java.nio.ByteBuffer;
import java.util.Arrays;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class FooArgs implements org.apache.thrift.TBase<FooArgs, FooArgs._Fields>, java.io.Serializable, Cloneable, Comparable<FooArgs> {
	private static final org.apache.thrift.protocol.TStruct STRUCT_DESC = new org.apache.thrift.protocol.TStruct("FooArgs");

	private static final org.apache.thrift.protocol.TField NEW_MESSAGE_FIELD_DESC = new org.apache.thrift.protocol.TField("newMessage", org.apache.thrift.protocol.TType.STRING, (short)1);
	private static final org.apache.thrift.protocol.TField MESSAGE_ARGS_FIELD_DESC = new org.apache.thrift.protocol.TField("messageArgs", org.apache.thrift.protocol.TType.STRING, (short)2);
	private static final org.apache.thrift.protocol.TField MESSAGE_RESULT_FIELD_DESC = new org.apache.thrift.protocol.TField("messageResult", org.apache.thrift.protocol.TType.STRING, (short)3);

	public String newMessage;
	public String messageArgs;
	public String messageResult;
	/** The set of fields this struct contains, along with convenience methods for finding and manipulating them. */
	public enum _Fields implements org.apache.thrift.TFieldIdEnum {
		NEW_MESSAGE((short)1, "newMessage"),
		MESSAGE_ARGS((short)2, "messageArgs"),
		MESSAGE_RESULT((short)3, "messageResult")
		;

		private static final Map<String, _Fields> byName = new HashMap<String, _Fields>();

		static {
			for (_Fields field : EnumSet.allOf(_Fields.class)) {
				byName.put(field.getFieldName(), field);
			}
		}

		/**
		 * Find the _Fields constant that matches fieldId, or null if its not found.
		 */
		public static _Fields findByThriftId(int fieldId) {
			switch(fieldId) {
				case 1: // NEW_MESSAGE
					return NEW_MESSAGE;
				case 2: // MESSAGE_ARGS
					return MESSAGE_ARGS;
				case 3: // MESSAGE_RESULT
					return MESSAGE_RESULT;
				default:
					return null;
			}
		}

		/**
		 * Find the _Fields constant that matches fieldId, throwing an exception
		 * if it is not found.
		 */
		public static _Fields findByThriftIdOrThrow(int fieldId) {
			_Fields fields = findByThriftId(fieldId);
			if (fields == null) throw new IllegalArgumentException("Field " + fieldId + " doesn't exist!");
			return fields;
		}

		/**
		 * Find the _Fields constant that matches name, or null if its not found.
		 */
		public static _Fields findByName(String name) {
			return byName.get(name);
		}

		private final short _thriftId;
		private final String _fieldName;

		_Fields(short thriftId, String fieldName) {
			_thriftId = thriftId;
			_fieldName = fieldName;
		}

		public short getThriftFieldId() {
			return _thriftId;
		}

		public String getFieldName() {
			return _fieldName;
		}
	}

	// isset id assignments
	public FooArgs() {
	}

	public FooArgs(
		String newMessage,
		String messageArgs,
		String messageResult) {
		this();
		this.newMessage = newMessage;
		this.messageArgs = messageArgs;
		this.messageResult = messageResult;
	}

	/**
	 * Performs a deep copy on <i>other</i>.
	 */
	public FooArgs(FooArgs other) {
		if (other.isSetNewMessage()) {
			this.newMessage = other.newMessage;
		}
		if (other.isSetMessageArgs()) {
			this.messageArgs = other.messageArgs;
		}
		if (other.isSetMessageResult()) {
			this.messageResult = other.messageResult;
		}
	}

	public FooArgs deepCopy() {
		return new FooArgs(this);
	}

	@Override
	public void clear() {
		this.newMessage = null;

		this.messageArgs = null;

		this.messageResult = null;

	}

	public String getNewMessage() {
		return this.newMessage;
	}

	public FooArgs setNewMessage(String newMessage) {
		this.newMessage = newMessage;
		return this;
	}

	public void unsetNewMessage() {
		this.newMessage = null;
	}

	/** Returns true if field newMessage is set (has been assigned a value) and false otherwise */
	public boolean isSetNewMessage() {
		return this.newMessage != null;
	}

	public void setNewMessageIsSet(boolean value) {
		if (!value) {
			this.newMessage = null;
		}
	}

	public String getMessageArgs() {
		return this.messageArgs;
	}

	public FooArgs setMessageArgs(String messageArgs) {
		this.messageArgs = messageArgs;
		return this;
	}

	public void unsetMessageArgs() {
		this.messageArgs = null;
	}

	/** Returns true if field messageArgs is set (has been assigned a value) and false otherwise */
	public boolean isSetMessageArgs() {
		return this.messageArgs != null;
	}

	public void setMessageArgsIsSet(boolean value) {
		if (!value) {
			this.messageArgs = null;
		}
	}

	public String getMessageResult() {
		return this.messageResult;
	}

	public FooArgs setMessageResult(String messageResult) {
		this.messageResult = messageResult;
		return this;
	}

	public void unsetMessageResult() {
		this.messageResult = null;
	}

	/** Returns true if field messageResult is set (has been assigned a value) and false otherwise */
	public boolean isSetMessageResult() {
		return this.messageResult != null;
	}

	public void setMessageResultIsSet(boolean value) {
		if (!value) {
			this.messageResult = null;
		}
	}

	public void setFieldValue(_Fields field, Object value) {
		switch (field) {
		case NEW_MESSAGE:
			if (value == null) {
				unsetNewMessage();
			} else {
				setNewMessage((String)value);
			}
			break;

		case MESSAGE_ARGS:
			if (value == null) {
				unsetMessageArgs();
			} else {
				setMessageArgs((String)value);
			}
			break;

		case MESSAGE_RESULT:
			if (value == null) {
				unsetMessageResult();
			} else {
				setMessageResult((String)value);
			}
			break;

		}
	}

	public Object getFieldValue(_Fields field) {
		switch (field) {
		case NEW_MESSAGE:
			return getNewMessage();

		case MESSAGE_ARGS:
			return getMessageArgs();

		case MESSAGE_RESULT:
			return getMessageResult();

		}
		throw new IllegalStateException();
	}

	/** Returns true if field corresponding to fieldID is set (has been assigned a value) and false otherwise */
	public boolean isSet(_Fields field) {
		if (field == null) {
			throw new IllegalArgumentException();
		}

		switch (field) {
		case NEW_MESSAGE:
			return isSetNewMessage();
		case MESSAGE_ARGS:
			return isSetMessageArgs();
		case MESSAGE_RESULT:
			return isSetMessageResult();
		}
		throw new IllegalStateException();
	}

	@Override
	public boolean equals(Object that) {
		if (that == null)
			return false;
		if (that instanceof FooArgs)
			return this.equals((FooArgs)that);
		return false;
	}

	public boolean equals(FooArgs that) {
		if (that == null)
			return false;
		if (!Objects.equals(this.newMessage, that.newMessage))
			return false;
		if (!Objects.equals(this.messageArgs, that.messageArgs))
			return false;
		if (!Objects.equals(this.messageResult, that.messageResult))
			return false;
		return true;
	}

	@Override
	public int hashCode() {
		List<Object> list = new ArrayList<Object>();

		boolean present_newMessage = true && (isSetNewMessage());
		list.add(present_newMessage);
		if (present_newMessage)
			list.add(newMessage);

		boolean present_messageArgs = true && (isSetMessageArgs());
		list.add(present_messageArgs);
		if (present_messageArgs)
			list.add(messageArgs);

		boolean present_messageResult = true && (isSetMessageResult());
		list.add(present_messageResult);
		if (present_messageResult)
			list.add(messageResult);

		return list.hashCode();
	}

	@Override
	public int compareTo(FooArgs other) {
		if (!getClass().equals(other.getClass())) {
			return getClass().getName().compareTo(other.getClass().getName());
		}

		int lastComparison = 0;

		lastComparison = Boolean.compare(isSetNewMessage(), other.isSetNewMessage());
		if (lastComparison != 0) {
			return lastComparison;
		}
		if (isSetNewMessage()) {
			lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.newMessage, other.newMessage);
			if (lastComparison != 0) {
				return lastComparison;
			}
		}
		lastComparison = Boolean.compare(isSetMessageArgs(), other.isSetMessageArgs());
		if (lastComparison != 0) {
			return lastComparison;
		}
		if (isSetMessageArgs()) {
			lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.messageArgs, other.messageArgs);
			if (lastComparison != 0) {
				return lastComparison;
			}
		}
		lastComparison = Boolean.compare(isSetMessageResult(), other.isSetMessageResult());
		if (lastComparison != 0) {
			return lastComparison;
		}
		if (isSetMessageResult()) {
			lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.messageResult, other.messageResult);
			if (lastComparison != 0) {
				return lastComparison;
			}
		}
		return 0;
	}

	public _Fields fieldForId(int fieldId) {
		return _Fields.findByThriftId(fieldId);
	}

	public void read(org.apache.thrift.protocol.TProtocol iprot) throws org.apache.thrift.TException {
		if (iprot.getScheme() != StandardScheme.class) {
			throw new UnsupportedOperationException();
		}
		new FooArgsStandardScheme().read(iprot, this);
	}

	public void write(org.apache.thrift.protocol.TProtocol oprot) throws org.apache.thrift.TException {
		if (oprot.getScheme() != StandardScheme.class) {
			throw new UnsupportedOperationException();
		}
		new FooArgsStandardScheme().write(oprot, this);
	}

	@Override
	public String toString() {
		StringBuilder sb = new StringBuilder("FooArgs(");
		boolean first = true;

		sb.append("newMessage:");
		sb.append(this.newMessage);
		first = false;
		if (!first) sb.append(", ");
		sb.append("messageArgs:");
		sb.append(this.messageArgs);
		first = false;
		if (!first) sb.append(", ");
		sb.append("messageResult:");
		sb.append(this.messageResult);
		first = false;
		sb.append(")");
		return sb.toString();
	}

	public void validate() throws org.apache.thrift.TException {
		// check for required fields
		// check for sub-struct validity
	}

	private void writeObject(java.io.ObjectOutputStream out) throws java.io.IOException {
		try {
			write(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(out)));
		} catch (org.apache.thrift.TException te) {
			throw new java.io.IOException(te);
		}
	}

	private void readObject(java.io.ObjectInputStream in) throws java.io.IOException, ClassNotFoundException {
		try {
			// it doesn't seem like you should have to do this, but java serialization is wacky, and doesn't call the default constructor.
			read(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(in)));
		} catch (org.apache.thrift.TException te) {
			throw new java.io.IOException(te);
		}
	}

	private static class FooArgsStandardScheme extends StandardScheme<FooArgs> {

		public void read(org.apache.thrift.protocol.TProtocol iprot, FooArgs struct) throws org.apache.thrift.TException {
			org.apache.thrift.protocol.TField schemeField;
			iprot.readStructBegin();
			while (true) {
				schemeField = iprot.readFieldBegin();
				if (schemeField.type == org.apache.thrift.protocol.TType.STOP) {
					break;
				}
				switch (schemeField.id) {
					case 1: // NEW_MESSAGE
						if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
							struct.newMessage = iprot.readString();
							struct.setNewMessageIsSet(true);
						} else {
							org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
						}
						break;
					case 2: // MESSAGE_ARGS
						if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
							struct.messageArgs = iprot.readString();
							struct.setMessageArgsIsSet(true);
						} else {
							org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
						}
						break;
					case 3: // MESSAGE_RESULT
						if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
							struct.messageResult = iprot.readString();
							struct.setMessageResultIsSet(true);
						} else {
							org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
						}
						break;
					default:
						org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
				}
				iprot.readFieldEnd();
			}
			iprot.readStructEnd();

			// check for required fields of primitive type, which can't be checked in the validate method
			struct.validate();
		}

		public void write(org.apache.thrift.protocol.TProtocol oprot, FooArgs struct) throws org.apache.thrift.TException {
			struct.validate();

			oprot.writeStructBegin(STRUCT_DESC);
			if (struct.isSetNewMessage()) {
				oprot.writeFieldBegin(NEW_MESSAGE_FIELD_DESC);
				String elem126 = struct.newMessage;
				oprot.writeString(elem126);
				oprot.writeFieldEnd();
			}
			if (struct.isSetMessageArgs()) {
				oprot.writeFieldBegin(MESSAGE_ARGS_FIELD_DESC);
				String elem127 = struct.messageArgs;
				oprot.writeString(elem127);
				oprot.writeFieldEnd();
			}
			if (struct.isSetMessageResult()) {
				oprot.writeFieldBegin(MESSAGE_RESULT_FIELD_DESC);
				String elem128 = struct.messageResult;
				oprot.writeString(elem128);
				oprot.writeFieldEnd();
			}
			oprot.writeFieldStop();
			oprot.writeStructEnd();
		}

	}

}
