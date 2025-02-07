package extend

import (
	"net/url"

	java "github.com/wnxd/microdbg-java"
)

type Intent interface {
	GetData() *url.URL
	GetBooleanExtra(name string, defValue java.JBoolean) java.JBoolean
	GetByteExtra(name string, defValue java.JByte) java.JByte
	GetCharExtra(name string, defValue java.JChar) java.JChar
	GetShortExtra(name string, defValue java.JShort) java.JShort
	GetIntExtra(name string, defValue java.JInt) java.JInt
	GetLongExtra(name string, defValue java.JLong) java.JLong
	GetFloatExtra(name string, defValue java.JFloat) java.JFloat
	GetDoubleExtra(name string, defValue java.JDouble) java.JDouble
	GetStringExtra(name string) java.IString
}

type DefaultIntent struct {
}

func (DefaultIntent) GetData() *url.URL {
	return nil
}

func (DefaultIntent) GetBooleanExtra(name string, defValue java.JBoolean) java.JBoolean {
	return defValue
}

func (DefaultIntent) GetByteExtra(name string, defValue java.JByte) java.JByte {
	return defValue
}

func (DefaultIntent) GetCharExtra(name string, defValue java.JChar) java.JChar {
	return defValue
}

func (DefaultIntent) GetShortExtra(name string, defValue java.JShort) java.JShort {
	return defValue
}

func (DefaultIntent) GetIntExtra(name string, defValue java.JInt) java.JInt {
	return defValue
}

func (DefaultIntent) GetLongExtra(name string, defValue java.JLong) java.JLong {
	return defValue
}

func (DefaultIntent) GetFloatExtra(name string, defValue java.JFloat) java.JFloat {
	return defValue
}

func (DefaultIntent) GetDoubleExtra(name string, defValue java.JDouble) java.JDouble {
	return defValue
}

func (DefaultIntent) GetStringExtra(name string) java.IString {
	return nil
}
