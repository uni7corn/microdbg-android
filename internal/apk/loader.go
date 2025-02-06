package apk

import (
	"archive/zip"
	"strconv"
	"strings"

	"github.com/wnxd/microdbg-android/internal"
	"github.com/wnxd/microdbg-android/res"
)

func Load(name string) (internal.Package, error) {
	z, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}
	return loadWithZip(z)
}

func loadWithZip(z *zip.ReadCloser) (internal.Package, error) {
	manifestFile, err := z.Open("AndroidManifest.xml")
	if err != nil {
		z.Close()
		return nil, err
	}
	defer manifestFile.Close()
	decoder, err := res.NewXMLDecoder(manifestFile)
	if err != nil {
		z.Close()
		return nil, err
	}
	var manifest Manifest
	err = decoder.Decode(&manifest)
	if err != nil {
		z.Close()
		return nil, err
	}
	resourcesFile, err := z.Open("resources.arsc")
	if err != nil {
		z.Close()
		return nil, err
	}
	defer resourcesFile.Close()
	resources, err := res.ParseTable(resourcesFile)
	if err != nil {
		z.Close()
		return nil, err
	}
	info := &info{
		fs:         z,
		name:       manifest.Package,
		version:    version{name: manifest.VersionName},
		permission: make([]string, len(manifest.UsesPermission)),
	}
	info.version.code, _ = strconv.Atoi(manifest.VersionCode)
	info.sdk.min, _ = strconv.Atoi(manifest.UsesSdk.MinSdkVersion)
	info.sdk.target, _ = strconv.Atoi(manifest.UsesSdk.TargetSdkVersion)
	label := strings.TrimPrefix(manifest.Application.Label, "0x")
	id, _ := strconv.ParseInt(label, 16, 0)
	if value, ok := resources.Get(int(id)); ok {
		info.label = value.Value.(res.Value).String
	}
	for i := range manifest.UsesPermission {
		info.permission[i] = manifest.UsesPermission[i].Name
	}
	return info.init(), nil
}
