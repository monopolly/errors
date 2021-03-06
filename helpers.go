package errors

import (
	"github.com/monopolly/errors/e"
)

func Access(c ...interface{}) E     { return New(401, e.Access, c...) }
func Backup(c ...interface{}) E     { return New(500, e.Backup, c...) }
func Ban(c ...interface{}) E        { return New(401, e.Ban, c...) }
func Busy(c ...interface{}) E       { return New(400, e.Busy, c...) }
func CDN(c ...interface{}) E        { return New(500, e.CDN, c...) }
func Chars(c ...interface{}) E      { return New(400, e.Chars, c...) }
func Connection(c ...interface{}) E { return New(500, e.Conn, c...) }
func Creds(c ...interface{}) E      { return New(401, e.Creds, c...) }
func Crypt(c ...interface{}) E      { return New(401, e.Crypt, c...) }
func Database(c ...interface{}) E   { return New(500, e.Database, c...) }
func Dkim(c ...interface{}) E       { return New(400, e.Dkim, c...) }
func Download(c ...interface{}) E   { return New(500, e.Download, c...) }
func Duplicate(c ...interface{}) E  { return New(400, e.Duplicate, c...) }
func ID(c ...interface{}) E         { return New(400, e.ID, c...) }
func Login(c ...interface{}) E      { return New(401, e.Login, c...) }
func Email(c ...interface{}) E      { return New(400, e.Email, c...) }
func Password(c ...interface{}) E   { return New(400, e.Password, c...) }
func Len(c ...interface{}) E        { return New(400, e.Len, c...) }
func Empty(c ...interface{}) E      { return New(301, e.Empty, c...) }
func Exist(c ...interface{}) E      { return New(400, e.Exist, c...) }
func Expired(c ...interface{}) E    { return New(400, e.Expired, c...) }
func Facebook(c ...interface{}) E   { return New(400, e.Facebook, c...) }
func File(c ...interface{}) E       { return New(400, e.File, c...) }
func Format(c ...interface{}) E     { return New(400, e.Format, c...) }
func Hostname(c ...interface{}) E   { return New(500, e.Hostname, c...) }
func Internal(c ...interface{}) E   { return New(500, e.Internal, c...) }
func Invalid(c ...interface{}) E    { return New(400, e.Invalid, c...) }
func IP(c ...interface{}) E         { return New(400, e.Invalid, c...) }
func Language(c ...interface{}) E   { return New(400, e.Lang, c...) }
func Limit(c ...interface{}) E      { return New(403, e.Limit, c...) }
func Link(c ...interface{}) E       { return New(400, e.Link, c...) }
func Locale(c ...interface{}) E     { return New(400, e.Locale, c...) }
func Long(c ...interface{}) E       { return New(400, e.Long, c...) }
func Marshal(c ...interface{}) E    { return New(400, e.Marshal, c...) }
func Method(c ...interface{}) E     { return New(404, e.Method, c...) }
func NotExist(c ...interface{}) E   { return New(404, e.NotExist, c...) }
func NotFound(c ...interface{}) E   { return New(404, e.NotFound, c...) }
func Open(c ...interface{}) E       { return New(400, e.Open, c...) }
func Outdate(c ...interface{}) E    { return New(400, e.Outdate, c...) }
func Range(c ...interface{}) E      { return New(400, e.Range, c...) }
func Parse(c ...interface{}) E      { return New(400, e.Parse, c...) }
func Payment(c ...interface{}) E    { return New(400, e.Payment, c...) }
func Provider(c ...interface{}) E   { return New(400, e.Provider, c...) }
func Query(c ...interface{}) E      { return New(400, e.Query, c...) }
func Save(c ...interface{}) E       { return New(400, e.Save, c...) }
func Server(c ...interface{}) E     { return New(500, e.Server, c...) }
func Short(c ...interface{}) E      { return New(400, e.Short, c...) }
func Simple(c ...interface{}) E     { return New(400, e.Simple, c...) }
func Contact(c ...interface{}) E    { return New(200, e.Contact, c...) }
func Status(c ...interface{}) E     { return New(400, e.Status, c...) }
func Timeout(c ...interface{}) E    { return New(500, e.Timeout, c...) }
func Token(c ...interface{}) E      { return New(401, e.Token, c...) }
func Pin(c ...interface{}) E        { return New(401, e.Pin, c...) }
func Hack(c ...interface{}) E       { return New(401, e.Hack, c...) }
func Request(c ...interface{}) E    { return New(400, e.Request, c...) }
func Unknown(c ...interface{}) E    { return New(400, e.Unknown, c...) }
func Unmarshal(c ...interface{}) E  { return New(400, e.Unmarshal, c...) }
func Update(c ...interface{}) E     { return New(400, e.Update, c...) }
func Verify(c ...interface{}) E     { return New(401, e.Verify, c...) }
func Websocket(c ...interface{}) E  { return New(500, e.Websocket, c...) }
func Overflow(c ...interface{}) E   { return New(500, e.Overflow, c...) }
