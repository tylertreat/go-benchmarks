package test

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Struct) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Field1":
			z.Field1, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Field2":
			z.Field2, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Field3":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Field3) >= int(xsz) {
				z.Field3 = z.Field3[:xsz]
			} else {
				z.Field3 = make([]string, xsz)
			}
			for xvk := range z.Field3 {
				z.Field3[xvk], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "Field4":
			z.Field4, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "Field5":
			z.Field5, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Field6":
			z.Field6, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Struct) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "Field1"
	err = en.Append(0x86, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Field1)
	if err != nil {
		return
	}
	// write "Field2"
	err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Field2)
	if err != nil {
		return
	}
	// write "Field3"
	err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x33)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Field3)))
	if err != nil {
		return
	}
	for xvk := range z.Field3 {
		err = en.WriteString(z.Field3[xvk])
		if err != nil {
			return
		}
	}
	// write "Field4"
	err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x34)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.Field4)
	if err != nil {
		return
	}
	// write "Field5"
	err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x35)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Field5)
	if err != nil {
		return
	}
	// write "Field6"
	err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x36)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Field6)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Struct) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Field1"
	o = append(o, 0x86, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31)
	o = msgp.AppendString(o, z.Field1)
	// string "Field2"
	o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32)
	o = msgp.AppendInt(o, z.Field2)
	// string "Field3"
	o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x33)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Field3)))
	for xvk := range z.Field3 {
		o = msgp.AppendString(o, z.Field3[xvk])
	}
	// string "Field4"
	o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x34)
	o = msgp.AppendUint64(o, z.Field4)
	// string "Field5"
	o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x35)
	o = msgp.AppendString(o, z.Field5)
	// string "Field6"
	o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x36)
	o = msgp.AppendString(o, z.Field6)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Struct) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Field1":
			z.Field1, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Field2":
			z.Field2, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Field3":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Field3) >= int(xsz) {
				z.Field3 = z.Field3[:xsz]
			} else {
				z.Field3 = make([]string, xsz)
			}
			for xvk := range z.Field3 {
				z.Field3[xvk], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "Field4":
			z.Field4, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "Field5":
			z.Field5, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Field6":
			z.Field6, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Struct) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.Field1) + 7 + msgp.IntSize + 7 + msgp.ArrayHeaderSize
	for xvk := range z.Field3 {
		s += msgp.StringPrefixSize + len(z.Field3[xvk])
	}
	s += 7 + msgp.Uint64Size + 7 + msgp.StringPrefixSize + len(z.Field5) + 7 + msgp.StringPrefixSize + len(z.Field6)
	return
}
