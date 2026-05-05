package errors

/*
Constants Generator
Help to auto-generated correct constants for lists
Martin Prestone (c) 2024
github.com/monopolly

Struct options
#up            Uppercase for golang vars
name=Status    Named var. Ex name=Category
type=string    Type of const list. Ex type=string, int by default
#swift         Create swift enum
#index         Create int index anyway

Fields options
#              Add lists for fields. Ex: id int //#readonly #must...
list{}         Create a.List() for field. Ex: list{colors, text}
iota=100       Add iota counter
name{}         Add new name for fields. Ex name{newjsonname}
go{}           Add new golang var name for fields. Ex go{ArtList}
title{}        Add title for fields. Ex title{News}
desc{}         Add desc for fields. Ex desc{New movie is good}
*/
type EE int

const (
	ErrorChars           = EE(iota) + 400 //#C400 iota=400
	ErrorContact                          //#C400
	ErrorDate                             //#C400
	ErrorDkim                             //#C400
	ErrorDuplicate                        //#C400
	ErrorEmail                            //#C400
	ErrorEmpty                            //#C400
	ErrorExist                            //#C400
	ErrorExpired                          //#C400
	ErrorFacebook                         //#C400
	ErrorFile                             //#C400
	ErrorForbidden                        //#C400
	ErrorFormat                           //#C400
	ErrorId                               //#C400
	ErrorInit                             //#C400
	ErrorInvalid                          //#C400
	ErrorIp                               //#C400
	ErrorLanguage                         //#C400
	ErrorLen                              //#C400
	ErrorLink                             //#C400
	ErrorLocale                           //#C400
	ErrorLong                             //#C400
	ErrorMarshal                          //#C400
	ErrorOpen                             //#C400
	ErrorOutdate                          //#C400
	ErrorParse                            //#C400
	ErrorPassword                         //#C400
	ErrorPayment                          //#C400
	ErrorProvider                         //#C400
	ErrorQuery                            //#C400
	ErrorRange                            //#C400
	ErrorReadonly                         //#C400
	ErrorRequest                          //#C400
	ErrorSave                             //#C400
	ErrorShort                            //#C400
	ErrorSimple                           //#C400
	ErrorStatus                           //#C400
	ErrorUnknown                          //#C400
	ErrorUnmarshal                        //#C400
	ErrorUpdate                           //#C400
	ErrorValid                            //#C400
	ErrorAccess                           //#C401
	ErrorBan                              //#C401
	ErrorCreds                            //#C401
	ErrorCrypt                            //#C401
	ErrorHack                             //#C401
	ErrorLogin                            //#C401
	ErrorPin                              //#C401
	ErrorToken                            //#C401
	ErrorVerify                           //#C401
	ErrorBusy                             //#C402
	ErrorLimit                            //#C403
	ErrorMethod                           //#C404
	ErrorNotexist                         //#C404
	ErrorNotfound                         //#C404
	ErrorToomanyRequests                  //#C429
)

const (
	ErrorBackup      = EE(iota) + 500 //#C500  iota=500
	ErrorCdn                          //#C500
	ErrorConfirm                      //#C500
	ErrorConnection                   //#C500
	ErrorDatabase                     //#C500
	ErrorDownload                     //#C500
	ErrorHostname                     //#C500
	ErrorInternal                     //#C500
	ErrorOverflow                     //#C500
	ErrorServer                       //#C500
	ErrorTimeout                      //#C500
	ErrorTry                          //#C500
	ErrorUnavailable                  //#C500
	ErrorWebsocket                    //#C500
)

func (a EE) String() string {
	switch a {
	case ErrorChars:
		return "errorChars"
	case ErrorContact:
		return "errorContact"
	case ErrorDate:
		return "errorDate"
	case ErrorDkim:
		return "errorDkim"
	case ErrorDuplicate:
		return "errorDuplicate"
	case ErrorEmail:
		return "errorEmail"
	case ErrorEmpty:
		return "errorEmpty"
	case ErrorExist:
		return "errorExist"
	case ErrorExpired:
		return "errorExpired"
	case ErrorFacebook:
		return "errorFacebook"
	case ErrorFile:
		return "errorFile"
	case ErrorForbidden:
		return "errorForbidden"
	case ErrorFormat:
		return "errorFormat"
	case ErrorId:
		return "errorId"
	case ErrorInit:
		return "errorInit"
	case ErrorInvalid:
		return "errorInvalid"
	case ErrorIp:
		return "errorIp"
	case ErrorLanguage:
		return "errorLanguage"
	case ErrorLen:
		return "errorLen"
	case ErrorLink:
		return "errorLink"
	case ErrorLocale:
		return "errorLocale"
	case ErrorLong:
		return "errorLong"
	case ErrorMarshal:
		return "errorMarshal"
	case ErrorOpen:
		return "errorOpen"
	case ErrorOutdate:
		return "errorOutdate"
	case ErrorParse:
		return "errorParse"
	case ErrorPassword:
		return "errorPassword"
	case ErrorPayment:
		return "errorPayment"
	case ErrorProvider:
		return "errorProvider"
	case ErrorQuery:
		return "errorQuery"
	case ErrorRange:
		return "errorRange"
	case ErrorReadonly:
		return "errorReadonly"
	case ErrorRequest:
		return "errorRequest"
	case ErrorSave:
		return "errorSave"
	case ErrorShort:
		return "errorShort"
	case ErrorSimple:
		return "errorSimple"
	case ErrorStatus:
		return "errorStatus"
	case ErrorUnknown:
		return "errorUnknown"
	case ErrorUnmarshal:
		return "errorUnmarshal"
	case ErrorUpdate:
		return "errorUpdate"
	case ErrorValid:
		return "errorValid"
	case ErrorAccess:
		return "errorAccess"
	case ErrorBan:
		return "errorBan"
	case ErrorCreds:
		return "errorCreds"
	case ErrorCrypt:
		return "errorCrypt"
	case ErrorHack:
		return "errorHack"
	case ErrorLogin:
		return "errorLogin"
	case ErrorPin:
		return "errorPin"
	case ErrorToken:
		return "errorToken"
	case ErrorVerify:
		return "errorVerify"
	case ErrorBusy:
		return "errorBusy"
	case ErrorLimit:
		return "errorLimit"
	case ErrorMethod:
		return "errorMethod"
	case ErrorNotexist:
		return "errorNotexist"
	case ErrorNotfound:
		return "errorNotfound"
	case ErrorToomanyRequests:
		return "errorToomanyRequests"
	case ErrorBackup:
		return "errorBackup"
	case ErrorCdn:
		return "errorCdn"
	case ErrorConfirm:
		return "errorConfirm"
	case ErrorConnection:
		return "errorConnection"
	case ErrorDatabase:
		return "errorDatabase"
	case ErrorDownload:
		return "errorDownload"
	case ErrorHostname:
		return "errorHostname"
	case ErrorInternal:
		return "errorInternal"
	case ErrorOverflow:
		return "errorOverflow"
	case ErrorServer:
		return "errorServer"
	case ErrorTimeout:
		return "errorTimeout"
	case ErrorTry:
		return "errorTry"
	case ErrorUnavailable:
		return "errorUnavailable"
	case ErrorWebsocket:
		return "errorWebsocket"
	default:
		return ""
	}
}

func (a EE) Title() string {
	switch a {
	case ErrorChars:
		return "ErrorChars"
	case ErrorContact:
		return "ErrorContact"
	case ErrorDate:
		return "ErrorDate"
	case ErrorDkim:
		return "ErrorDkim"
	case ErrorDuplicate:
		return "ErrorDuplicate"
	case ErrorEmail:
		return "ErrorEmail"
	case ErrorEmpty:
		return "ErrorEmpty"
	case ErrorExist:
		return "ErrorExist"
	case ErrorExpired:
		return "ErrorExpired"
	case ErrorFacebook:
		return "ErrorFacebook"
	case ErrorFile:
		return "ErrorFile"
	case ErrorForbidden:
		return "ErrorForbidden"
	case ErrorFormat:
		return "ErrorFormat"
	case ErrorId:
		return "ErrorId"
	case ErrorInit:
		return "ErrorInit"
	case ErrorInvalid:
		return "ErrorInvalid"
	case ErrorIp:
		return "ErrorIp"
	case ErrorLanguage:
		return "ErrorLanguage"
	case ErrorLen:
		return "ErrorLen"
	case ErrorLink:
		return "ErrorLink"
	case ErrorLocale:
		return "ErrorLocale"
	case ErrorLong:
		return "ErrorLong"
	case ErrorMarshal:
		return "ErrorMarshal"
	case ErrorOpen:
		return "ErrorOpen"
	case ErrorOutdate:
		return "ErrorOutdate"
	case ErrorParse:
		return "ErrorParse"
	case ErrorPassword:
		return "ErrorPassword"
	case ErrorPayment:
		return "ErrorPayment"
	case ErrorProvider:
		return "ErrorProvider"
	case ErrorQuery:
		return "ErrorQuery"
	case ErrorRange:
		return "ErrorRange"
	case ErrorReadonly:
		return "ErrorReadonly"
	case ErrorRequest:
		return "ErrorRequest"
	case ErrorSave:
		return "ErrorSave"
	case ErrorShort:
		return "ErrorShort"
	case ErrorSimple:
		return "ErrorSimple"
	case ErrorStatus:
		return "ErrorStatus"
	case ErrorUnknown:
		return "ErrorUnknown"
	case ErrorUnmarshal:
		return "ErrorUnmarshal"
	case ErrorUpdate:
		return "ErrorUpdate"
	case ErrorValid:
		return "ErrorValid"
	case ErrorAccess:
		return "ErrorAccess"
	case ErrorBan:
		return "ErrorBan"
	case ErrorCreds:
		return "ErrorCreds"
	case ErrorCrypt:
		return "ErrorCrypt"
	case ErrorHack:
		return "ErrorHack"
	case ErrorLogin:
		return "ErrorLogin"
	case ErrorPin:
		return "ErrorPin"
	case ErrorToken:
		return "ErrorToken"
	case ErrorVerify:
		return "ErrorVerify"
	case ErrorBusy:
		return "ErrorBusy"
	case ErrorLimit:
		return "ErrorLimit"
	case ErrorMethod:
		return "ErrorMethod"
	case ErrorNotexist:
		return "ErrorNotexist"
	case ErrorNotfound:
		return "ErrorNotfound"
	case ErrorToomanyRequests:
		return "ErrorToomanyRequests"
	case ErrorBackup:
		return "ErrorBackup"
	case ErrorCdn:
		return "ErrorCdn"
	case ErrorConfirm:
		return "ErrorConfirm"
	case ErrorConnection:
		return "ErrorConnection"
	case ErrorDatabase:
		return "ErrorDatabase"
	case ErrorDownload:
		return "ErrorDownload"
	case ErrorHostname:
		return "ErrorHostname"
	case ErrorInternal:
		return "ErrorInternal"
	case ErrorOverflow:
		return "ErrorOverflow"
	case ErrorServer:
		return "ErrorServer"
	case ErrorTimeout:
		return "ErrorTimeout"
	case ErrorTry:
		return "ErrorTry"
	case ErrorUnavailable:
		return "ErrorUnavailable"
	case ErrorWebsocket:
		return "ErrorWebsocket"
	default:
		return ""
	}
}

func (a EE) Desc() string {
	switch a {
	default:
		return ""
	}
}

func (a EE) Int() int {
	return int(a)
}

func ValidError(v int) bool {
	switch EE(v) {
	case ErrorChars, ErrorContact, ErrorDate, ErrorDkim, ErrorDuplicate, ErrorEmail, ErrorEmpty, ErrorExist, ErrorExpired, ErrorFacebook, ErrorFile, ErrorForbidden, ErrorFormat, ErrorId, ErrorInit, ErrorInvalid, ErrorIp, ErrorLanguage, ErrorLen, ErrorLink, ErrorLocale, ErrorLong, ErrorMarshal, ErrorOpen, ErrorOutdate, ErrorParse, ErrorPassword, ErrorPayment, ErrorProvider, ErrorQuery, ErrorRange, ErrorReadonly, ErrorRequest, ErrorSave, ErrorShort, ErrorSimple, ErrorStatus, ErrorUnknown, ErrorUnmarshal, ErrorUpdate, ErrorValid, ErrorAccess, ErrorBan, ErrorCreds, ErrorCrypt, ErrorHack, ErrorLogin, ErrorPin, ErrorToken, ErrorVerify, ErrorBusy, ErrorLimit, ErrorMethod, ErrorNotexist, ErrorNotfound, ErrorToomanyRequests, ErrorBackup, ErrorCdn, ErrorConfirm, ErrorConnection, ErrorDatabase, ErrorDownload, ErrorHostname, ErrorInternal, ErrorOverflow, ErrorServer, ErrorTimeout, ErrorTry, ErrorUnavailable, ErrorWebsocket:
		return true
	default:
		return false
	}
}

func EEIndexes() []EE {
	return []EE{ErrorChars, ErrorContact, ErrorDate, ErrorDkim, ErrorDuplicate, ErrorEmail, ErrorEmpty, ErrorExist, ErrorExpired, ErrorFacebook, ErrorFile, ErrorForbidden, ErrorFormat, ErrorId, ErrorInit, ErrorInvalid, ErrorIp, ErrorLanguage, ErrorLen, ErrorLink, ErrorLocale, ErrorLong, ErrorMarshal, ErrorOpen, ErrorOutdate, ErrorParse, ErrorPassword, ErrorPayment, ErrorProvider, ErrorQuery, ErrorRange, ErrorReadonly, ErrorRequest, ErrorSave, ErrorShort, ErrorSimple, ErrorStatus, ErrorUnknown, ErrorUnmarshal, ErrorUpdate, ErrorValid, ErrorAccess, ErrorBan, ErrorCreds, ErrorCrypt, ErrorHack, ErrorLogin, ErrorPin, ErrorToken, ErrorVerify, ErrorBusy, ErrorLimit, ErrorMethod, ErrorNotexist, ErrorNotfound, ErrorToomanyRequests, ErrorBackup, ErrorCdn, ErrorConfirm, ErrorConnection, ErrorDatabase, ErrorDownload, ErrorHostname, ErrorInternal, ErrorOverflow, ErrorServer, ErrorTimeout, ErrorTry, ErrorUnavailable, ErrorWebsocket}
}

func EEC400List() []EE {
	return []EE{ErrorChars, ErrorContact, ErrorDate, ErrorDkim, ErrorDuplicate, ErrorEmail, ErrorEmpty, ErrorExist, ErrorExpired, ErrorFacebook, ErrorFile, ErrorForbidden, ErrorFormat, ErrorId, ErrorInit, ErrorInvalid, ErrorIp, ErrorLanguage, ErrorLen, ErrorLink, ErrorLocale, ErrorLong, ErrorMarshal, ErrorOpen, ErrorOutdate, ErrorParse, ErrorPassword, ErrorPayment, ErrorProvider, ErrorQuery, ErrorRange, ErrorReadonly, ErrorRequest, ErrorSave, ErrorShort, ErrorSimple, ErrorStatus, ErrorUnknown, ErrorUnmarshal, ErrorUpdate, ErrorValid}
}

func (a EE) C400() bool {
	switch a {
	case ErrorChars, ErrorContact, ErrorDate, ErrorDkim, ErrorDuplicate, ErrorEmail, ErrorEmpty, ErrorExist, ErrorExpired, ErrorFacebook, ErrorFile, ErrorForbidden, ErrorFormat, ErrorId, ErrorInit, ErrorInvalid, ErrorIp, ErrorLanguage, ErrorLen, ErrorLink, ErrorLocale, ErrorLong, ErrorMarshal, ErrorOpen, ErrorOutdate, ErrorParse, ErrorPassword, ErrorPayment, ErrorProvider, ErrorQuery, ErrorRange, ErrorReadonly, ErrorRequest, ErrorSave, ErrorShort, ErrorSimple, ErrorStatus, ErrorUnknown, ErrorUnmarshal, ErrorUpdate, ErrorValid:
		return true
	default:
		return false
	}
}
func EEC401List() []EE {
	return []EE{ErrorAccess, ErrorBan, ErrorCreds, ErrorCrypt, ErrorHack, ErrorLogin, ErrorPin, ErrorToken, ErrorVerify}
}

func (a EE) C401() bool {
	switch a {
	case ErrorAccess, ErrorBan, ErrorCreds, ErrorCrypt, ErrorHack, ErrorLogin, ErrorPin, ErrorToken, ErrorVerify:
		return true
	default:
		return false
	}
}
func EEC402List() []EE {
	return []EE{ErrorBusy}
}

func (a EE) C402() bool {
	switch a {
	case ErrorBusy:
		return true
	default:
		return false
	}
}
func EEC403List() []EE {
	return []EE{ErrorLimit}
}

func (a EE) C403() bool {
	switch a {
	case ErrorLimit:
		return true
	default:
		return false
	}
}
func EEC404List() []EE {
	return []EE{ErrorMethod, ErrorNotexist, ErrorNotfound}
}

func (a EE) C404() bool {
	switch a {
	case ErrorMethod, ErrorNotexist, ErrorNotfound:
		return true
	default:
		return false
	}
}
func EEC429List() []EE {
	return []EE{ErrorToomanyRequests}
}

func (a EE) C429() bool {
	switch a {
	case ErrorToomanyRequests:
		return true
	default:
		return false
	}
}
func EEC500List() []EE {
	return []EE{ErrorBackup, ErrorCdn, ErrorConfirm, ErrorConnection, ErrorDatabase, ErrorDownload, ErrorHostname, ErrorInternal, ErrorOverflow, ErrorServer, ErrorTimeout, ErrorTry, ErrorUnavailable, ErrorWebsocket}
}

func (a EE) C500() bool {
	switch a {
	case ErrorBackup, ErrorCdn, ErrorConfirm, ErrorConnection, ErrorDatabase, ErrorDownload, ErrorHostname, ErrorInternal, ErrorOverflow, ErrorServer, ErrorTimeout, ErrorTry, ErrorUnavailable, ErrorWebsocket:
		return true
	default:
		return false
	}
}

func Index(v string) int {
	switch v {
	case "errorChars":
		return 400
	case "errorContact":
		return 401
	case "errorDate":
		return 402
	case "errorDkim":
		return 403
	case "errorDuplicate":
		return 404
	case "errorEmail":
		return 405
	case "errorEmpty":
		return 406
	case "errorExist":
		return 407
	case "errorExpired":
		return 408
	case "errorFacebook":
		return 409
	case "errorFile":
		return 410
	case "errorForbidden":
		return 411
	case "errorFormat":
		return 412
	case "errorId":
		return 413
	case "errorInit":
		return 414
	case "errorInvalid":
		return 415
	case "errorIp":
		return 416
	case "errorLanguage":
		return 417
	case "errorLen":
		return 418
	case "errorLink":
		return 419
	case "errorLocale":
		return 420
	case "errorLong":
		return 421
	case "errorMarshal":
		return 422
	case "errorOpen":
		return 423
	case "errorOutdate":
		return 424
	case "errorParse":
		return 425
	case "errorPassword":
		return 426
	case "errorPayment":
		return 427
	case "errorProvider":
		return 428
	case "errorQuery":
		return 429
	case "errorRange":
		return 430
	case "errorReadonly":
		return 431
	case "errorRequest":
		return 432
	case "errorSave":
		return 433
	case "errorShort":
		return 434
	case "errorSimple":
		return 435
	case "errorStatus":
		return 436
	case "errorUnknown":
		return 437
	case "errorUnmarshal":
		return 438
	case "errorUpdate":
		return 439
	case "errorValid":
		return 440
	case "errorAccess":
		return 441
	case "errorBan":
		return 442
	case "errorCreds":
		return 443
	case "errorCrypt":
		return 444
	case "errorHack":
		return 445
	case "errorLogin":
		return 446
	case "errorPin":
		return 447
	case "errorToken":
		return 448
	case "errorVerify":
		return 449
	case "errorBusy":
		return 450
	case "errorLimit":
		return 451
	case "errorMethod":
		return 452
	case "errorNotexist":
		return 453
	case "errorNotfound":
		return 454
	case "errorToomanyRequests":
		return 455
	case "errorBackup":
		return 500
	case "errorCdn":
		return 501
	case "errorConfirm":
		return 502
	case "errorConnection":
		return 503
	case "errorDatabase":
		return 504
	case "errorDownload":
		return 505
	case "errorHostname":
		return 506
	case "errorInternal":
		return 507
	case "errorOverflow":
		return 508
	case "errorServer":
		return 509
	case "errorTimeout":
		return 510
	case "errorTry":
		return 511
	case "errorUnavailable":
		return 512
	case "errorWebsocket":
		return 513
	default:
		return 0
	}
}
