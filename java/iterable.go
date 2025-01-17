package java

import (
	"slices"

	java "github.com/wnxd/microdbg-java"
)

type iterator struct {
	hasNext func() bool
	next    func() java.IObject
	remove  func()
}

func init() {
	definePrimitiveMethod(FakeIterableClass, "iterator", "()Ljava/util/Iterator;", Modifier_PUBLIC|Modifier_ABSTRACT)

	definePrimitiveMethod(FakeIteratorClass, "hasNext", "()Z", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		it := fake.Value().(*iterator)
		return it.hasNext()
	})
	definePrimitiveMethod(FakeIteratorClass, "next", "()Ljava/lang/Object;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		it := fake.Value().(*iterator)
		return it.next()
	})
	definePrimitiveMethod(FakeIteratorClass, "remove", "()V", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		it := fake.Value().(*iterator)
		it.remove()
		return nil
	})

	definePrimitiveMethod(FakeCollectionClass, "iterator", "()Ljava/util/Iterator;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		list := fake.Value().(*[]java.IObject)
		index := 0
		it := &iterator{
			hasNext: func() bool {
				return index < len(*list)
			},
			next: func() java.IObject {
				val := (*list)[index]
				index++
				return val
			},
			remove: func() {
				*list = slices.Delete(*list, index-1, index)
			},
		}
		return FakeIteratorClass.NewObject(it)
	})
	definePrimitiveMethod(FakeCollectionClass, "size", "()I", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		list := fake.Value().(*[]java.IObject)
		return java.JInt(len(*list))
	})
	definePrimitiveMethod(FakeCollectionClass, "isEmpty", "()Z", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		list := fake.Value().(*[]java.IObject)
		return len(*list) == 0
	})
	definePrimitiveMethod(FakeCollectionClass, "contains", "(Ljava/lang/Object;)Z", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		list := fake.Value().(*[]java.IObject)
		val := ToObject[java.IObject](args[0])
		for _, v := range *list {
			if v.Equals(val) {
				return true
			}
		}
		return false
	})
	definePrimitiveMethod(FakeCollectionClass, "toArray", "()[Ljava/lang/Object;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		list := fake.Value().(*[]java.IObject)
		return ArrayOf(FakeObjectArrayClass, *list)
	})
	definePrimitiveMethod(FakeCollectionClass, "toArray", "(Ljava/lang/Object;)[Ljava/lang/Object;", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		list := fake.Value().(*[]java.IObject)
		a := args[0].(java.IGenericArray[java.IObject])
		copy(a.Elements(), *list)
		return a
	})
	definePrimitiveMethod(FakeCollectionClass, "add", "(Ljava/lang/Object;)Z", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		list := fake.Value().(*[]java.IObject)
		*list = append(*list, ToObject[java.IObject](args[0]))
		return true
	})
	definePrimitiveMethod(FakeCollectionClass, "remove", "(Ljava/lang/Object;)Z", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		list := fake.Value().(*[]java.IObject)
		val := ToObject[java.IObject](args[0])
		for i, v := range *list {
			if v.Equals(val) {
				*list = slices.Delete(*list, i, i+1)
				return true
			}
		}
		return false
	})
	definePrimitiveMethod(FakeCollectionClass, "clear", "()V", Modifier_PUBLIC|Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(FakeObject)
		list := fake.Value().(*[]java.IObject)
		clear(*list)
		return nil
	})
}
