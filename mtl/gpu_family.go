// +build darwin

package mtl

type MTLGPUFamily uint16

const (

	//MTLGPUFamilyCommon1 A GPU supporting the common family 1 characteristics.
	MTLGPUFamilyCommon1 MTLGPUFamily = 3001

	//MTLGPUFamilyCommon2 A GPU supporting the common family 2 characteristics.
	MTLGPUFamilyCommon2 MTLGPUFamily = 3002

	//MTLGPUFamilyCommon3 A GPU supporting the common family 3 characteristics.
	MTLGPUFamilyCommon3 MTLGPUFamily = 3003

	//MTLGPUFamilyApple1 An Apple family 1 GPU.
	MTLGPUFamilyApple1 MTLGPUFamily = 1001

	//MTLGPUFamilyApple1 An Apple family 2 GPU.
	MTLGPUFamilyApple2 MTLGPUFamily = 1002

	//MTLGPUFamilyApple1 An Apple family 3 GPU.
	MTLGPUFamilyApple3 MTLGPUFamily = 1003

	//MTLGPUFamilyApple1 An Apple family 4 GPU.
	MTLGPUFamilyApple4 MTLGPUFamily = 1004

	//MTLGPUFamilyApple1 An Apple family 5 GPU.
	MTLGPUFamilyApple5 MTLGPUFamily = 1005

	//MTLGPUFamilyApple1 An Apple family 6 GPU.
	MTLGPUFamilyApple6 MTLGPUFamily = 1006

	//MTLGPUFamilyApple1 An Apple family 7 GPU.
	MTLGPUFamilyApple7 MTLGPUFamily = 1007

	//MTLGPUFamilyMac1 A family 1 Mac GPU.
	MTLGPUFamilyMac1 MTLGPUFamily = 2001

	//MTLGPUFamilyMac2 A family 2 Mac GPU.
	MTLGPUFamilyMac2 MTLGPUFamily = 2002

	//MTLGPUFamilyMacCatalyst1 A family 1 Mac GPU when running in an iOS app.
	MTLGPUFamilyMacCatalyst1 = 4001

	//MTLGPUFamilyMacCatalyst1 A family 2 Mac GPU when running in an iOS app.
	MTLGPUFamilyMacCatalyst2 = 4002
)
