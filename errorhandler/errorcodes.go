package errorhandler

const (
	ErrorTransactionBuilder             = 1000
	ErrorTransactionSigning             = 1001
	ErrorConvertBase64                  = 1002
	ErrorSubmitTransaction              = 1003
	ErrorInvalidToken                   = 1004
	ErrorInsertTransaction              = 1005
	ErrorParsingJSON                    = 1006
	ErrorReadingHistory                 = 1007
	ErrorUpdatingZagazTokenCharges      = 1008
	ErrorUpdatingZagazServiceCharges    = 1009
	ErrorUpdatingZagazConversionRate    = 1010
	ErrorUpdatingCurrencyConversionRate = 1011
	ErrorGetZagazConversionRate         = 1012
	ErrorGetZagazServiceCharge          = 1013
	ErrorGetTransaction                 = 1014
	ErrorUnknownDatabaseException       = 1015
	ErrorRecordNotFound                 = 1016
	ErrorInvalidTxnType                 = 1017
	ErrorInvalidFilter                  = 1018
	ErrorNoRowsAffected                 = 1019
	ErrorMissingQueryParameters         = 1020
	ErrorRequestHorizon                 = 1021
	ErrorReadingResponse                = 1022
	ErrorInvalidDateFormat              = 1023
	ErrorTimeZoneFormat                 = 1024
	ErrorInvalidUserDetails             = 3001
	ErrorExceedUIDLength                = 3002
	ErrorShortUID                       = 3003
	ErrorExceedPasswordLength           = 3004
	ErrorShortPassword                  = 3005
	ErrorConfirmPassword                = 3006
	ErrorUnknowdDatabaseError           = 3007
	ErrorDuplicateData                  = 3008
	ErrorUserIDExists                   = 3009
	ErrorUnknownDuplicateData           = 3010
	ErrorMobileNumberExists             = 3011
	ErrorRequestData                    = 3012
	ErrorOnRequest                      = 3013
	ErrorRequestingService              = 3014
	ErrorResponseBody                   = 3015
	ErrorEmptyToken                     = 3016
	ErrorUnknownResponseBody            = 3017
	ErrorFromHarmony                    = 3018
	ErrorSendSMS                        = 3019
	ErrorOtpMismatch                    = 3020
	ErrorCountryDoesNotExist            = 3022
	ErrorUploadToS3                     = 3023
	ErrorPasswordDoesNotMatch           = 3024
	ErrorMissingQueryParams             = 3025
	ErrorSigningToken                   = 3026
	ErrorPushToQue                      = 3027
	ErrorSMSRegion                      = 3028
	ErrorUnexpectedToken                = 3029
	ErrorTokenExpired                   = 3031
	ErrorUserKeysResponse               = 3032
	ErrorReadRecord                     = 3036
	ErrorUpdateRecord                   = 3037
	ErrorSetDataIntoQueue               = 3038
	ErrorVerifyApp                      = 3039
	ErrorGetDataFromQueue               = 3040
	ErrorInvalidInputData               = 3041
	ErrorAuthorizeApp                   = 3042
	ErrorDefineRole                     = 3043
	ErrorGetUserAccessControlList       = 3044
	ErrorPermissionDenied               = 3045
	ErrorDecodeString                   = 3046
)

const (
	PaymentSuccess = 0

	SuccessCreateAccount                 = 2000
	SuccessInsertTransaction             = 2001
	SuccessPaymentCompleted              = 2002
	SuccessPaymentStatusUpdated          = 2003
	SuccessReadingHistory                = 2004
	SuccessReadingTransactionDetails     = 2005
	SuccessUpdatingZagazTokenCharges     = 2006
	SuccessUpdatingZagazServiceCharge    = 2007
	SucessUpdatingCurrencyConversionRate = 2008
	SuccessGetZagazCharges               = 2009
	SuccessGetZagazConverionRate         = 2010
	SuccessGetActiveZagazServiceCharge   = 2011
	SuccessGetTransaction                = 2012
	SuccessGetActiveZagazTokenCharge     = 2013
	SuccessUpdateZagazTokenCharge        = 2014
	SuccessInsertZagazTokenCharge        = 2015
	SuccessGetActiveZagazConversionRate  = 2016
	SuccessUpdateZagazConversionRate     = 2017
	SuccessInsertZagazConversionRate     = 2018
	SuccessUpdateZagazServiceCharge      = 2019
	SuccessInsertZagazServiceCharge      = 2020
	SuccessUpdateTransaction             = 2021
	SuccessReadingBlocks                 = 2022
	SuccessSendRequest                   = 2023
	SuccessGetTransactionsBetweenDate    = 2021
	SuccessCreateUser                    = 5000
	SuccessValidUserDetails              = 5001
	SuccessInsertUser                    = 5002
	SuccessGenerateToken                 = 5003
	SuccessRequest                       = 5004
	SuccessParseResponse                 = 5005
	SuccessGenerateOTP                   = 5006
	SuccessSendSMS                       = 5007
	SuccessSendOTP                       = 5008
	SuccessInsertUserExtend              = 5009
	SuccessInsertRecord                  = 5010
	SuccessValidToken                    = 5011
	SuccessValidOtp                      = 5013
	SuccessCreateUserAccount             = 5014
	SuccessValidAccountDetails           = 5015
	SuccessCreateUserKYC                 = 5017
	SuccessUploadToS3                    = 5018
	SuccessLogin                         = 5020
	SuccessGetActiveStatus               = 5021
	SuccessApproveSI                     = 5026
	SuccessDeclareSI                     = 5027
	SuccessGetSIBankList                 = 5028
	SuccessGetUser                       = 5029
	SuccessGetKycHash                    = 5031
	SuccessGetDeviceID                   = 5032
	SuccessGetToken                      = 5033
	SuccessCreateUserKeys                = 5034
	SuccessGetWalletAddress              = 5037
	SuccessParseQueryParams              = 5038
	SuccessPromoteUser                   = 5039
	SuccessReadRecord                    = 5041
	SuccessGetKycDocList                 = 5042
	SuccessGetPendingSI                  = 5043
	SuccessUpdateRecord                  = 5044
	SuccessReadUserKYCNotVerified        = 5045
	SuccessReadUserKYCDetails            = 5046
	SuccessApproveKYCDetails             = 5047
	SuccessGetAccessControlList          = 5048
	SuccessDefineRole                    = 5049
	SuccessCreatePayLoad                 = 5050
	SuccessGenerateAccessControls        = 5051
	SuccessGetAccessControlStructure     = 5052
	SuccessAuthorizeApp                  = 5053
	SuccessVerifyApp                     = 5054
)
