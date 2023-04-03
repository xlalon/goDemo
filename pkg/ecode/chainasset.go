package ecode

var (
	ChainInvalid       = New(1000510)
	ChainNotFound      = New(1000511)
	ChainCodeChange    = New(1000512)
	ChainCodeInvalid   = New(1000513)
	ChainNameInvalid   = New(1000514)
	ChainStatusInvalid = New(1000515)

	AssetInvalid           = New(2000511)
	AssetExist             = New(2000512)
	AssetCodeChange        = New(2000513)
	AssetIdentityChange    = New(2000514)
	AssetSettingChange     = New(2000515)
	AssetCodeInvalid       = New(2000516)
	AssetNameInvalid       = New(2000517)
	AssetIdentityInvalid   = New(2000518)
	AssetPrecessionInvalid = New(2000519)
	AssetStatusInvalid     = New(2000520)

	DomainIdInvalid = New(8000001)
	DomainIdChange  = New(8000002)

	AssetSettingInvalid = New(3000511)
)
