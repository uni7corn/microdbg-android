package java

import (
	"maps"

	java "github.com/wnxd/microdbg-java"
)

type Map = map[java.IObject]java.IObject

type entry struct {
	key, value java.IObject
}

func MapOf(mp Map) java.IObject {
	return FakeMapClass.NewObject(mp)
}

func init() {
	definePrimitiveMethod(FakeMapClass, "size", "()I", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, _ ...any) any {
		fake := obj.(FakeObject)
		mp := fake.Value().(Map)
		return java.JInt(len(mp))
	})
	definePrimitiveMethod(FakeMapClass, "isEmpty", "()Z", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, _ ...any) any {
		fake := obj.(FakeObject)
		mp := fake.Value().(Map)
		return len(mp) == 0
	})
	definePrimitiveMethod(FakeMapClass, "containsKey", "(Ljava/lang/Object;)Z", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		mp := fake.Value().(Map)
		_, ok := mp[ToObject[java.IObject](args[0])]
		return ok
	})
	definePrimitiveMethod(FakeMapClass, "containsValue", "(Ljava/lang/Object;)Z", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		mp := fake.Value().(Map)
		val := ToObject[java.IObject](args[0])
		for _, v := range mp {
			if v.Equals(val) {
				return true
			}
		}
		return false
	})
	definePrimitiveMethod(FakeMapClass, "get", "(Ljava/lang/Object;)Ljava/lang/Object;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		mp := fake.Value().(Map)
		return mp[ToObject[java.IObject](args[0])]
	})
	definePrimitiveMethod(FakeMapClass, "put", "(Ljava/lang/Object;Ljava/lang/Object;)Ljava/lang/Object;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		mp := fake.Value().(Map)
		key := ToObject[java.IObject](args[0])
		old := mp[key]
		mp[key] = ToObject[java.IObject](args[1])
		return old
	})
	definePrimitiveMethod(FakeMapClass, "remove", "(Ljava/lang/Object;)Ljava/lang/Object;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		mp := fake.Value().(Map)
		key := ToObject[java.IObject](args[0])
		old := mp[key]
		delete(mp, key)
		return old
	})
	definePrimitiveMethod(FakeMapClass, "clear", "()V", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		mp := fake.Value().(Map)
		clear(mp)
		return nil
	})
	definePrimitiveMethod(FakeMapClass, "keySet", "()Ljava/util/Set;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		mp := fake.Value().(Map)
		keys := make(map[java.IObject]struct{}, len(mp))
		for k := range mp {
			keys[k] = struct{}{}
		}
		return FakeSetClass.NewObject(keys)
	})
	definePrimitiveMethod(FakeMapClass, "values", "()Ljava/util/Collection;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		mp := fake.Value().(Map)
		maps.Values(mp)
		vals := make([]java.IObject, 0, len(mp))
		for _, v := range mp {
			vals = append(vals, v)
		}
		return FakeCollectionClass.NewObject(&vals)
	})
	definePrimitiveMethod(FakeMapClass, "entrySet", "()Ljava/util/Set;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		mp := fake.Value().(Map)
		entries := make(map[java.IObject]struct{}, len(mp))
		for k, v := range mp {
			entries[FakeMapEntryClass.NewObject(&entry{k, v})] = struct{}{}
		}
		return FakeSetClass.NewObject(entries)
	})

	definePrimitiveMethod(FakeMapEntryClass, "getKey", "()Ljava/lang/Object;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		entry := fake.Value().(*entry)
		return entry.key
	})
	definePrimitiveMethod(FakeMapEntryClass, "getValue", "()Ljava/lang/Object;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		entry := fake.Value().(*entry)
		return entry.value
	})
	definePrimitiveMethod(FakeMapEntryClass, "setValue", "(Ljava/lang/Object;)Ljava/lang/Object;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		entry := fake.Value().(*entry)
		old := entry.value
		entry.value = ToObject[java.IObject](args[0])
		return old
	})

	definePrimitiveMethod(FakeHashMapClass, ConstructorMethodName, "(I)V", Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		initialCapacity := args[0].(java.JInt)
		return FakeHashMapClass.NewObject(make(Map, initialCapacity))
	})
}
