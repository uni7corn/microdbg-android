package java

func init() {
	definePrimitiveMethod(FakeNumberClass, "byteValue", "()B", Modifier_PUBLIC|Modifier_ABSTRACT)
	definePrimitiveMethod(FakeNumberClass, "shortValue", "()S", Modifier_PUBLIC|Modifier_ABSTRACT)
	definePrimitiveMethod(FakeNumberClass, "intValue", "()I", Modifier_PUBLIC|Modifier_ABSTRACT)
	definePrimitiveMethod(FakeNumberClass, "longValue", "()J", Modifier_PUBLIC|Modifier_ABSTRACT)
	definePrimitiveMethod(FakeNumberClass, "floatValue", "()F", Modifier_PUBLIC|Modifier_ABSTRACT)
	definePrimitiveMethod(FakeNumberClass, "doubleValue", "()D", Modifier_PUBLIC|Modifier_ABSTRACT)
}
