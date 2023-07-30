package ecode

var (
	ChainInvalid       = New(2000510)
	ChainNotFound      = New(2000511)
	ChainCodeChange    = New(2000512)
	ChainCodeInvalid   = New(2000513)
	ChainNameInvalid   = New(2000514)
	ChainStatusInvalid = New(2000515)

	AssetInvalid           = New(2010511)
	AssetExist             = New(2010512)
	AssetCodeChange        = New(2010513)
	AssetIdentityChange    = New(2010514)
	AssetSettingChange     = New(2010515)
	AssetCodeInvalid       = New(2010516)
	AssetNameInvalid       = New(2010517)
	AssetIdentityInvalid   = New(2010518)
	AssetPrecessionInvalid = New(2010519)
	AssetStatusInvalid     = New(2010520)
	AssetSettingInvalid    = New(2010521)
)
