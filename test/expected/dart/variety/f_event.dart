// Autogenerated by Frugal Compiler (3.0.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

import 'dart:typed_data' show Uint8List;
import 'package:thrift/thrift.dart' as thrift;
import 'package:variety/variety.dart' as t_variety;
import 'package:actual_base_dart/actual_base_dart.dart' as t_actual_base_dart;
import 'package:intermediate_include/intermediate_include.dart' as t_intermediate_include;
import 'package:validStructs/validStructs.dart' as t_validStructs;
import 'package:ValidTypes/ValidTypes.dart' as t_ValidTypes;
import 'package:subdir_include_ns/subdir_include_ns.dart' as t_subdir_include_ns;

/// This docstring gets added to the generated code because it has
/// the @ sign.
class Event implements thrift.TBase {
  static final thrift.TStruct _STRUCT_DESC = new thrift.TStruct("Event");
  static final thrift.TField _ID_FIELD_DESC = new thrift.TField("ID", thrift.TType.I64, 1);
  static final thrift.TField _MESSAGE_FIELD_DESC = new thrift.TField("Message", thrift.TType.STRING, 2);

  /// ID is a unique identifier for an event.
  int _iD = 0;
  static const int ID = 1;
  /// Message contains the event payload.
  String _message;
  static const int MESSAGE = 2;

  bool __isset_iD = false;

  Event() {
    this.iD = t_variety.VarietyConstants.DEFAULT_ID;
  }

  /// ID is a unique identifier for an event.
  int get iD => this._iD;

  /// ID is a unique identifier for an event.
  set iD(int iD) {
    this._iD = iD;
    this.__isset_iD = true;
  }

  bool isSetID() => this.__isset_iD;

  unsetID() {
    this.__isset_iD = false;
  }

  /// Message contains the event payload.
  String get message => this._message;

  /// Message contains the event payload.
  set message(String message) {
    this._message = message;
  }

  bool isSetMessage() => this.message != null;

  unsetMessage() {
    this.message = null;
  }

  @override
  getFieldValue(int fieldID) {
    switch (fieldID) {
      case ID:
        return this.iD;
      case MESSAGE:
        return this.message;
      default:
        throw new ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  @override
  setFieldValue(int fieldID, Object value) {
    switch (fieldID) {
      case ID:
        if (value == null) {
          unsetID();
        } else {
          this.iD = value as int;
        }
        break;

      case MESSAGE:
        if (value == null) {
          unsetMessage();
        } else {
          this.message = value as String;
        }
        break;

      default:
        throw new ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  // Returns true if the field corresponding to fieldID is set (has been assigned a value) and false otherwise
  @override
  bool isSet(int fieldID) {
    switch (fieldID) {
      case ID:
        return isSetID();
      case MESSAGE:
        return isSetMessage();
      default:
        throw new ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  @override
  read(thrift.TProtocol iprot) {
    iprot.readStructBegin();
    for (thrift.TField field = iprot.readFieldBegin();
        field.type != thrift.TType.STOP;
        field = iprot.readFieldBegin()) {
      switch (field.id) {
        case ID:
          if (field.type == thrift.TType.I64) {
            iD = iprot.readI64();
            this.__isset_iD = true;
          } else {
            thrift.TProtocolUtil.skip(iprot, field.type);
          }
          break;
        case MESSAGE:
          if (field.type == thrift.TType.STRING) {
            message = iprot.readString();
          } else {
            thrift.TProtocolUtil.skip(iprot, field.type);
          }
          break;
        default:
          thrift.TProtocolUtil.skip(iprot, field.type);
          break;
      }
      iprot.readFieldEnd();
    }
    iprot.readStructEnd();

    // check for required fields of primitive type, which can't be checked in the validate method
    validate();
  }

  @override
  write(thrift.TProtocol oprot) {
    validate();

    oprot.writeStructBegin(_STRUCT_DESC);
    oprot.writeFieldBegin(_ID_FIELD_DESC);
    oprot.writeI64(iD);
    oprot.writeFieldEnd();
    if (this.message != null) {
      oprot.writeFieldBegin(_MESSAGE_FIELD_DESC);
      oprot.writeString(message);
      oprot.writeFieldEnd();
    }
    oprot.writeFieldStop();
    oprot.writeStructEnd();
  }

  @override
  String toString() {
    StringBuffer ret = new StringBuffer("Event(");

    ret.write("iD:");
    ret.write(this.iD);

    ret.write(", ");
    ret.write("message:");
    if (this.message == null) {
      ret.write("null");
    } else {
      ret.write(this.message);
    }

    ret.write(")");

    return ret.toString();
  }

  @override
  bool operator ==(Object o) {
    if (o is Event) {
      return this.iD == o.iD &&
        this.message == o.message;
    }
    return false;
  }

  @override
  int get hashCode {
    var value = 17;
    value = (value * 31) ^ iD.hashCode;
    value = (value * 31) ^ message.hashCode;
    return value;
  }

  Event clone({
    int iD: null,
    String message: null,
  }) {
    return new Event()
      ..iD = iD ?? this.iD
      ..message = message ?? this.message;
  }

  validate() {
    // check for required fields
    // check that fields of type enum have valid values
  }
}
