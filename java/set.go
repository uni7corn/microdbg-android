package java

import (
	java "github.com/wnxd/microdbg-java"
)

func init() {
	definePrimitiveMethod(FakeSetClass, "iterator", "()Ljava/util/Iterator;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		set := fake.Value().(map[java.IObject]struct{})
		keys := make([]java.IObject, 0, len(set))
		for k := range set {
			keys = append(keys, k)
		}
		index := 0
		it := &iterator{
			hasNext: func() bool {
				return index < len(keys)
			},
			next: func() java.IObject {
				val := keys[index]
				index++
				return val
			},
			remove: func() {
				delete(set, keys[index-1])
			},
		}
		return FakeIteratorClass.NewObject(it)
	})
	definePrimitiveMethod(FakeSetClass, "size", "()I", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		set := fake.Value().(map[java.IObject]struct{})
		return java.JInt(len(set))
	})
	definePrimitiveMethod(FakeSetClass, "isEmpty", "()Z", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		set := fake.Value().(map[java.IObject]struct{})
		return len(set) == 0
	})
	definePrimitiveMethod(FakeSetClass, "contains", "(Ljava/lang/Object;)Z", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		set := fake.Value().(map[java.IObject]struct{})
		_, ok := set[ToObject[java.IObject](args[0])]
		return ok
	})
	definePrimitiveMethod(FakeSetClass, "toArray", "()[Ljava/lang/Object;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		set := fake.Value().(map[java.IObject]struct{})
		keys := make([]java.IObject, 0, len(set))
		for k := range set {
			keys = append(keys, k)
		}
		return ArrayOf(FakeObjectArrayClass, keys)
	})
	definePrimitiveMethod(FakeSetClass, "toArray", "(Ljava/lang/Object;)[Ljava/lang/Object;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		set := fake.Value().(map[java.IObject]struct{})
		a := args[0].(java.IGenericArray[java.IObject])
		arr := a.Elements()
		i := 0
		for v := range set {
			if i >= len(arr) {
				break
			}
			arr[i] = v
			i++
		}
		return a
	})
	definePrimitiveMethod(FakeSetClass, "add", "(Ljava/lang/Object;)Z", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		set := fake.Value().(map[java.IObject]struct{})
		set[ToObject[java.IObject](args[0])] = struct{}{}
		return true
	})
	definePrimitiveMethod(FakeSetClass, "remove", "(Ljava/lang/Object;)Z", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		set := fake.Value().(map[java.IObject]struct{})
		delete(set, ToObject[java.IObject](args[0]))
		return true
	})
	definePrimitiveMethod(FakeSetClass, "clear", "()V", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		set := fake.Value().(map[java.IObject]struct{})
		clear(set)
		return nil
	})
}
