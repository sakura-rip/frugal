/**
 * Autogenerated by Frugal Compiler (3.14.2)
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *
 * @generated
 */
package variety.java;

import java.util.Map;
import java.util.HashMap;
import org.apache.thrift.TEnum;

public enum ItsAnEnum implements org.apache.thrift.TEnum {
	FIRST(2),
	SECOND(3),
	THIRD(4),
	fourth(5),
	Fifth(6),
	sIxItH(7);

	private final int value;

	private ItsAnEnum(int value) {
		this.value = value;
	}

	public int getValue() {
		return value;
	}

	public static ItsAnEnum findByValue(int value) {
		switch (value) {
			case 2:
				return FIRST;
			case 3:
				return SECOND;
			case 4:
				return THIRD;
			case 5:
				return fourth;
			case 6:
				return Fifth;
			case 7:
				return sIxItH;
			default:
				return null;
		}
	}
}
