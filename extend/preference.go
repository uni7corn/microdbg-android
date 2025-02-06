package extend

import java "github.com/wnxd/microdbg-java"

type SharedPreference interface {
	Contains(key string) java.JBoolean
	GetBoolean(key string, defValue java.JBoolean) java.JBoolean
	GetFloat(key string, defValue java.JFloat) java.JFloat
	GetInt(key string, defValue java.JInt) java.JInt
	GetLong(key string, defValue java.JLong) java.JLong
	GetString(key string, defValue java.IString) java.IString
}

type SharedPreferenceEditor interface {
	Clear()
	SetBoolean(key string, value java.JBoolean)
	SetFloat(key string, value java.JFloat)
	SetInt(key string, value java.JInt)
	SetLong(key string, value java.JLong)
	SetString(key string, value java.IString)
	Remove(key string)
}

type DefaultPreference struct {
}

func (DefaultPreference) Contains(key string) java.JBoolean {
	return false
}

func (DefaultPreference) GetBoolean(key string, defValue java.JBoolean) java.JBoolean {
	return defValue
}

func (DefaultPreference) GetFloat(key string, defValue java.JFloat) java.JFloat {
	return defValue
}

func (DefaultPreference) GetInt(key string, defValue java.JInt) java.JInt {
	return defValue
}

func (DefaultPreference) GetLong(key string, defValue java.JLong) java.JLong {
	return defValue
}

func (DefaultPreference) GetString(key string, defValue java.IString) java.IString {
	return defValue
}

func (DefaultPreference) Clear() {
}

func (DefaultPreference) SetBoolean(key string, value java.JBoolean) {
}

func (DefaultPreference) SetFloat(key string, value java.JFloat) {
}

func (DefaultPreference) SetInt(key string, value java.JInt) {
}

func (DefaultPreference) SetLong(key string, value java.JLong) {
}

func (DefaultPreference) SetString(key string, value java.IString) {
}

func (DefaultPreference) Remove(key string) {
}
