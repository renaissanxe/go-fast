// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package serialization

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// Attributes:
//  - Field1
//  - Field2
//  - Field3
//  - Field4
//  - Field5
//  - Field6
//  - Field7
type StructThrift struct {
  Field1 string `thrift:"field1,1" db:"field1" json:"field1"`
  Field2 int64 `thrift:"field2,2" db:"field2" json:"field2"`
  Field3 []string `thrift:"field3,3" db:"field3" json:"field3"`
  Field4 int64 `thrift:"field4,4" db:"field4" json:"field4"`
  Field5 string `thrift:"field5,5" db:"field5" json:"field5"`
  Field6 string `thrift:"field6,6" db:"field6" json:"field6"`
  Field7 []byte `thrift:"field7,7" db:"field7" json:"field7"`
}

func NewStructThrift() *StructThrift {
  return &StructThrift{}
}


func (p *StructThrift) GetField1() string {
  return p.Field1
}

func (p *StructThrift) GetField2() int64 {
  return p.Field2
}

func (p *StructThrift) GetField3() []string {
  return p.Field3
}

func (p *StructThrift) GetField4() int64 {
  return p.Field4
}

func (p *StructThrift) GetField5() string {
  return p.Field5
}

func (p *StructThrift) GetField6() string {
  return p.Field6
}

func (p *StructThrift) GetField7() []byte {
  return p.Field7
}
func (p *StructThrift) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 4:
      if err := p.ReadField4(iprot); err != nil {
        return err
      }
    case 5:
      if err := p.ReadField5(iprot); err != nil {
        return err
      }
    case 6:
      if err := p.ReadField6(iprot); err != nil {
        return err
      }
    case 7:
      if err := p.ReadField7(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *StructThrift)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Field1 = v
}
  return nil
}

func (p *StructThrift)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Field2 = v
}
  return nil
}

func (p *StructThrift)  ReadField3(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]string, 0, size)
  p.Field3 =  tSlice
  for i := 0; i < size; i ++ {
var _elem0 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem0 = v
}
    p.Field3 = append(p.Field3, _elem0)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *StructThrift)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Field4 = v
}
  return nil
}

func (p *StructThrift)  ReadField5(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.Field5 = v
}
  return nil
}

func (p *StructThrift)  ReadField6(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 6: ", err)
} else {
  p.Field6 = v
}
  return nil
}

func (p *StructThrift)  ReadField7(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(); err != nil {
  return thrift.PrependError("error reading field 7: ", err)
} else {
  p.Field7 = v
}
  return nil
}

func (p *StructThrift) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("StructThrift"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
    if err := p.writeField6(oprot); err != nil { return err }
    if err := p.writeField7(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *StructThrift) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("field1", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:field1: ", p), err) }
  if err := oprot.WriteString(string(p.Field1)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.field1 (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:field1: ", p), err) }
  return err
}

func (p *StructThrift) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("field2", thrift.I64, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:field2: ", p), err) }
  if err := oprot.WriteI64(int64(p.Field2)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.field2 (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:field2: ", p), err) }
  return err
}

func (p *StructThrift) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("field3", thrift.LIST, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:field3: ", p), err) }
  if err := oprot.WriteListBegin(thrift.STRING, len(p.Field3)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.Field3 {
    if err := oprot.WriteString(string(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:field3: ", p), err) }
  return err
}

func (p *StructThrift) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("field4", thrift.I64, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:field4: ", p), err) }
  if err := oprot.WriteI64(int64(p.Field4)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.field4 (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:field4: ", p), err) }
  return err
}

func (p *StructThrift) writeField5(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("field5", thrift.STRING, 5); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:field5: ", p), err) }
  if err := oprot.WriteString(string(p.Field5)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.field5 (5) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 5:field5: ", p), err) }
  return err
}

func (p *StructThrift) writeField6(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("field6", thrift.STRING, 6); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:field6: ", p), err) }
  if err := oprot.WriteString(string(p.Field6)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.field6 (6) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 6:field6: ", p), err) }
  return err
}

func (p *StructThrift) writeField7(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("field7", thrift.STRING, 7); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:field7: ", p), err) }
  if err := oprot.WriteBinary(p.Field7); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.field7 (7) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 7:field7: ", p), err) }
  return err
}

func (p *StructThrift) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("StructThrift(%+v)", *p)
}

