// Autogenerated by Frugal Compiler (3.14.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

// ignore_for_file: unused_import
// ignore_for_file: unused_field
import 'dart:typed_data' show Uint8List;

import 'package:collection/collection.dart';
import 'package:thrift/thrift.dart' as thrift;
import 'package:v1_music/v1_music.dart' as t_v1_music;

/// Exceptions are converted to the native format for each compiled
/// language.
class PurchasingError extends Error implements thrift.TBase {
  static final thrift.TStruct _STRUCT_DESC = thrift.TStruct('PurchasingError');
  static final thrift.TField _MESSAGE_FIELD_DESC = thrift.TField('message', thrift.TType.STRING, 1);
  static final thrift.TField _ERROR_CODE_FIELD_DESC = thrift.TField('error_code', thrift.TType.I16, 2);

  String _message;
  static const int MESSAGE = 1;
  int _error_code = 0;
  static const int ERROR_CODE = 2;

  bool __isset_error_code = false;

  String get message => this._message;

  set message(String message) {
    this._message = message;
  }

  bool isSetMessage() => this.message != null;

  unsetMessage() {
    this.message = null;
  }

  int get error_code => this._error_code;

  set error_code(int error_code) {
    this._error_code = error_code;
    this.__isset_error_code = true;
  }

  bool isSetError_code() => this.__isset_error_code;

  unsetError_code() {
    this.__isset_error_code = false;
  }

  @override
  getFieldValue(int fieldID) {
    switch (fieldID) {
      case MESSAGE:
        return this.message;
      case ERROR_CODE:
        return this.error_code;
      default:
        throw ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  @override
  setFieldValue(int fieldID, Object value) {
    switch (fieldID) {
      case MESSAGE:
        if (value == null) {
          unsetMessage();
        } else {
          this.message = value as String;
        }
        break;

      case ERROR_CODE:
        if (value == null) {
          unsetError_code();
        } else {
          this.error_code = value as int;
        }
        break;

      default:
        throw ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  // Returns true if the field corresponding to fieldID is set (has been assigned a value) and false otherwise
  @override
  bool isSet(int fieldID) {
    switch (fieldID) {
      case MESSAGE:
        return isSetMessage();
      case ERROR_CODE:
        return isSetError_code();
      default:
        throw ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  @override
  read(thrift.TProtocol iprot) {
    iprot.readStructBegin();
    for (thrift.TField field = iprot.readFieldBegin();
        field.type != thrift.TType.STOP;
        field = iprot.readFieldBegin()) {
      switch (field.id) {
        case MESSAGE:
          if (field.type == thrift.TType.STRING) {
            this.message = iprot.readString();
          } else {
            thrift.TProtocolUtil.skip(iprot, field.type);
          }
          break;
        case ERROR_CODE:
          if (field.type == thrift.TType.I16) {
            this.error_code = iprot.readI16();
            this.__isset_error_code = true;
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

    validate();
  }

  @override
  write(thrift.TProtocol oprot) {
    validate();

    oprot.writeStructBegin(_STRUCT_DESC);
    if (this.message != null) {
      oprot.writeFieldBegin(_MESSAGE_FIELD_DESC);
      oprot.writeString(this.message);
      oprot.writeFieldEnd();
    }
    oprot.writeFieldBegin(_ERROR_CODE_FIELD_DESC);
    oprot.writeI16(this.error_code);
    oprot.writeFieldEnd();
    oprot.writeFieldStop();
    oprot.writeStructEnd();
  }

  @override
  String toString() {
    StringBuffer ret = StringBuffer('PurchasingError(');

    ret.write('message:');
    if (this.message == null) {
      ret.write('null');
    } else {
      ret.write(this.message);
    }

    ret.write(', ');
    ret.write('error_code:');
    ret.write(this.error_code);

    ret.write(')');

    return ret.toString();
  }

  @override
  bool operator ==(Object o) {
    if (o is PurchasingError) {
      return this.message == o.message &&
        this.error_code == o.error_code;
    }
    return false;
  }

  @override
  int get hashCode {
    var value = 17;
    value = (value * 31) ^ this.message.hashCode;
    value = (value * 31) ^ this.error_code.hashCode;
    return value;
  }

  PurchasingError clone({
    String message,
    int error_code,
  }) {
    return PurchasingError()
      ..message = message ?? this.message
      ..error_code = error_code ?? this.error_code;
  }

  validate() {
  }
}
