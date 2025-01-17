package apk

type Manifest struct {
	CompileSdkVersion         string `xml:"compileSdkVersion,attr"`
	CompileSdkVersionCodename string `xml:"compileSdkVersionCodename,attr"`
	PlatformBuildVersionName  string `xml:"platformBuildVersionName,attr"`
	PlatformBuildVersionCode  string `xml:"platformBuildVersionCode,attr"`
	Package                   string `xml:"package,attr"`
	VersionName               string `xml:"versionName,attr"`
	VersionCode               string `xml:"versionCode,attr"`

	UsesSdk        UsesSdk          `xml:"uses-sdk"`
	UsesPermission []UsesPermission `xml:"uses-permission"`
	Application    Application      `xml:"application"`
}

type UsesSdk struct {
	MinSdkVersion    string `xml:"minSdkVersion,attr"`
	TargetSdkVersion string `xml:"targetSdkVersion,attr"`
}

type UsesPermission struct {
	Name string `xml:"name,attr"`
}

type Application struct {
	Label string `xml:"label,attr"`
}
