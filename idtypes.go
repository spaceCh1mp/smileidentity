package smileidentity

type countryCode string

const (
	GHANA       countryCode = "GH"
	NIGERIA     countryCode = "NG"
	KENYA       countryCode = "KE"
	SOUTHAFRICA countryCode = "ZA"
)

type smileIDType string

const (
	BVN               smileIDType = "BVN"
	CAC               smileIDType = "CAC"
	NIN               smileIDType = "NIN"
	TIN               smileIDType = "TIN"
	SSNIT             smileIDType = "SSNIT"
	PASSPORT          smileIDType = "PASSPORT"
	VOTERID           smileIDType = "VOTER_ID"
	ALIENCARD         smileIDType = "ALIEN_CARD"
	NATIONALID        smileIDType = "NATIONAL_ID"
	BANKACCOUNT       smileIDType = "BANK_ACCOUNT"
	PHONENUMBER       smileIDType = "PHONE_NUMBER"
	DRIVERSLICENSE    smileIDType = "DRIVERS_LICENSE"
	NATIONALIDNOPHOTO smileIDType = "NATIONAL_ID_NO_PHOTO"
)
