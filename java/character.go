package java

import (
	java "github.com/wnxd/microdbg-java"
)

type FakeCharacter java.JChar

func (c FakeCharacter) GetClass() java.IClass {
	return FakeCharacterClass
}

func (c FakeCharacter) HashCode() java.JInt {
	return java.JInt(c)
}

func (c FakeCharacter) Equals(obj java.IObject) java.JBoolean {
	return c == obj
}

func (c FakeCharacter) ToString() java.IString {
	return FakeString(rune(c))
}

func init() {
	FakeCharacterClass.Set("TYPE", FakeCharTYPE)
}
