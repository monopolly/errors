package errors

func Access(c ...interface{}) E          { return New(401, IDAccess, c...) }
func Backup(c ...interface{}) E          { return New(500, IDBackup, c...) }
func Ban(c ...interface{}) E             { return New(401, IDBan, c...) }
func Busy(c ...interface{}) E            { return New(402, IDBusy, c...) }
func Readonly(c ...interface{}) E        { return New(400, IDReadonly, c...) }
func CDN(c ...interface{}) E             { return New(500, IDCDN, c...) }
func Chars(c ...interface{}) E           { return New(400, IDChars, c...) }
func Connection(c ...interface{}) E      { return New(500, IDConn, c...) }
func Confirm(c ...interface{}) E         { return New(500, IDConfirm, c...) }
func Creds(c ...interface{}) E           { return New(401, IDCreds, c...) }
func Crypt(c ...interface{}) E           { return New(401, IDCrypt, c...) }
func Database(c ...interface{}) E        { return New(500, IDDatabase, c...) }
func Dkim(c ...interface{}) E            { return New(400, IDDkim, c...) }
func Download(c ...interface{}) E        { return New(500, IDDownload, c...) }
func Date(c ...interface{}) E            { return New(400, IDDate, c...) }
func Duplicate(c ...interface{}) E       { return New(400, IDDuplicate, c...) }
func ID(c ...interface{}) E              { return New(400, IDID, c...) }
func Login(c ...interface{}) E           { return New(401, IDLogin, c...) }
func Email(c ...interface{}) E           { return New(400, IDEmail, c...) }
func Password(c ...interface{}) E        { return New(400, IDPassword, c...) }
func Len(c ...interface{}) E             { return New(400, IDLen, c...) }
func Empty(c ...interface{}) E           { return New(400, IDEmpty, c...) }
func Exist(c ...interface{}) E           { return New(400, IDExist, c...) }
func Expired(c ...interface{}) E         { return New(400, IDExpired, c...) }
func Facebook(c ...interface{}) E        { return New(400, IDFacebook, c...) }
func Forbidden(c ...interface{}) E       { return New(400, IDForbidden, c...) }
func File(c ...interface{}) E            { return New(400, IDFile, c...) }
func Format(c ...interface{}) E          { return New(400, IDFormat, c...) }
func Hostname(c ...interface{}) E        { return New(500, IDHostname, c...) }
func Internal(c ...interface{}) E        { return New(500, IDInternal, c...) }
func Invalid(c ...interface{}) E         { return New(400, IDInvalid, c...) }
func Init(c ...interface{}) E            { return New(400, IDInit, c...) }
func Valid(c ...interface{}) E           { return New(400, IDValid, c...) }
func IP(c ...interface{}) E              { return New(400, IDIP, c...) }
func Language(c ...interface{}) E        { return New(400, IDLang, c...) }
func Limit(c ...interface{}) E           { return New(403, IDLimit, c...) }
func Link(c ...interface{}) E            { return New(400, IDLink, c...) }
func Locale(c ...interface{}) E          { return New(400, IDLocale, c...) }
func Long(c ...interface{}) E            { return New(400, IDLong, c...) }
func Marshal(c ...interface{}) E         { return New(400, IDMarshal, c...) }
func Method(c ...interface{}) E          { return New(404, IDMethod, c...) }
func NotExist(c ...interface{}) E        { return New(404, IDNotExist, c...) }
func NotFound(c ...interface{}) E        { return New(404, IDNotFound, c...) }
func Open(c ...interface{}) E            { return New(400, IDOpen, c...) }
func Outdate(c ...interface{}) E         { return New(400, IDOutdate, c...) }
func Range(c ...interface{}) E           { return New(400, IDRange, c...) }
func Parse(c ...interface{}) E           { return New(400, IDParse, c...) }
func Payment(c ...interface{}) E         { return New(400, IDPayment, c...) }
func Provider(c ...interface{}) E        { return New(400, IDProvider, c...) }
func Query(c ...interface{}) E           { return New(400, IDQuery, c...) }
func Save(c ...interface{}) E            { return New(400, IDSave, c...) }
func Server(c ...interface{}) E          { return New(500, IDServer, c...) }
func Short(c ...interface{}) E           { return New(400, IDShort, c...) }
func Simple(c ...interface{}) E          { return New(400, IDSimple, c...) }
func Contact(c ...interface{}) E         { return New(400, IDContact, c...) }
func Status(c ...interface{}) E          { return New(400, IDStatus, c...) }
func Timeout(c ...interface{}) E         { return New(500, IDTimeout, c...) }
func TooManyRequests(c ...interface{}) E { return New(500, IDTooManyRequests, c...) }
func Try(c ...interface{}) E             { return New(500, IDTry, c...) }
func Token(c ...interface{}) E           { return New(401, IDToken, c...) }
func Pin(c ...interface{}) E             { return New(401, IDPin, c...) }
func Hack(c ...interface{}) E            { return New(401, IDHack, c...) }
func Request(c ...interface{}) E         { return New(400, IDRequest, c...) }
func Unknown(c ...interface{}) E         { return New(400, IDUnknown, c...) }
func Unmarshal(c ...interface{}) E       { return New(400, IDUnmarshal, c...) }
func Update(c ...interface{}) E          { return New(400, IDUpdate, c...) }
func Verify(c ...interface{}) E          { return New(401, IDVerify, c...) }
func Websocket(c ...interface{}) E       { return New(500, IDWebsocket, c...) }
func Overflow(c ...interface{}) E        { return New(500, IDOverflow, c...) }
func Unavailable(c ...interface{}) E     { return New(500, IDUnavailable, c...) }
