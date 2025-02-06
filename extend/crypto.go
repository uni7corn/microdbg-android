package extend

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
)

type secretKey struct {
	data      java.IByteArray
	algorithm java.IString
}

type mac struct {
	hash.Hash
	h func() hash.Hash
}

func (ex *extend) defineCrypto() {
	SecretKey := ex.cf.DefineClass("javax.crypto.SecretKey", nil, ex.cf.GetClass("java.security.Key"))

	SecretKeySpec := ex.cf.DefineClass("javax.crypto.spec.SecretKeySpec", gava.FakeObjectClass, SecretKey)
	SecretKeySpec.DefineMethod(gava.ConstructorMethodName, "([BLjava/lang/String;)V", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		return SecretKeySpec.NewObject(&secretKey{
			data:      args[0].(java.IByteArray),
			algorithm: args[1].(java.IString),
		})
	})
	SecretKeySpec.DefineMethod("getAlgorithm", "()Ljava/lang/String;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		key := obj.(gava.FakeObject).Value().(*secretKey)
		return key.algorithm
	})
	SecretKeySpec.DefineMethod("getEncoded", "()[B", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		key := obj.(gava.FakeObject).Value().(*secretKey)
		return key.data
	})
	SecretKeySpec.DefineMethod("getFormat", "()Ljava/lang/String;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		return gava.FakeString("RAW")
	})

	Mac := ex.cf.DefineClass("javax.crypto.Mac")
	Mac.DefineMethod("getInstance", "(Ljava/lang/String;)Ljavax/crypto/Mac;", gava.Modifier_PUBLIC|gava.Modifier_STATIC|gava.Modifier_FINAL).BindCall(func(obj java.IObject, args ...any) any {
		algorithm := args[0].(java.IString).String()
		switch algorithm {
		case "HmacMD5":
			return Mac.NewObject(&mac{h: md5.New})
		case "HmacSHA1":
			return Mac.NewObject(&mac{h: sha1.New})
		case "HmacSHA224":
			return Mac.NewObject(&mac{h: sha256.New224})
		case "HmacSHA256":
			return Mac.NewObject(&mac{h: sha256.New})
		case "HmacSHA384":
			return Mac.NewObject(&mac{h: sha512.New384})
		case "HmacSHA512":
			return Mac.NewObject(&mac{h: sha512.New})
		}
		return nil
	})
	Mac.DefineMethod("init", "(Ljava/security/Key;)V", gava.Modifier_PUBLIC|gava.Modifier_FINAL).BindCall(func(obj java.IObject, args ...any) any {
		mac := obj.(gava.FakeObject).Value().(*mac)
		key := args[0].(gava.FakeObject)
		mac.Hash = hmac.New(mac.h, gava.GetBytes(key.FindMethod("getEncoded", "()[B").CallPrimitive(key).(java.IByteArray)))
		return nil
	})
	Mac.DefineMethod("reset", "()V", gava.Modifier_PUBLIC|gava.Modifier_FINAL).BindCall(func(obj java.IObject, args ...any) any {
		mac := obj.(gava.FakeObject).Value().(*mac)
		mac.Hash.Reset()
		return nil
	})
	Mac.DefineMethod("update", "([B)V", gava.Modifier_PUBLIC|gava.Modifier_FINAL).BindCall(func(obj java.IObject, args ...any) any {
		mac := obj.(gava.FakeObject).Value().(*mac)
		mac.Hash.Write(gava.GetBytes(args[0].(java.IByteArray)))
		return nil
	})
	Mac.DefineMethod("doFinal", "([B)[B", gava.Modifier_PUBLIC|gava.Modifier_FINAL).BindCall(func(obj java.IObject, args ...any) any {
		mac := obj.(gava.FakeObject).Value().(*mac)
		if args[0] != nil {
			mac.Hash.Write(gava.GetBytes(args[0].(java.IByteArray)))
		}
		return gava.BytesOf(mac.Hash.Sum(nil))
	})
}
