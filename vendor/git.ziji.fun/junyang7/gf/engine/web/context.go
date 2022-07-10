package web

import (
	"git.ziji.fun/junyang7/gf/common"
	"git.ziji.fun/junyang7/gf/common/interceptor"
	"git.ziji.fun/junyang7/gf/tool/_file"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type context struct {
	w          http.ResponseWriter
	r          *http.Request
	timeS      time.Time
	conf       *Conf
	filepath   string
	e          []byte
	headerList map[string]string
	method     string
	path       string
}

func newContext(w http.ResponseWriter, r *http.Request, conf *Conf) *context {
	this := &context{timeS: time.Now()}
	this.w = w
	this.r = r
	this.conf = conf
	this.headerList = map[string]string{}
	this.path = r.URL.Path
	this.method = r.Method
	return this
}
func (this *context) do() {
	defer this.handleException()
	this.checkOrigin()
	this.checkResource()
}
func (this *context) handleException() {
	if err := recover(); nil != err {
		switch err.(type) {
		case *common.Response:
			response := err.(*common.Response)
			this.e = []byte(response.Message)
		default:
			this.e = []byte(common.FailureMessage)
		}
	}
}
func (this *context) checkOrigin() {
	origin := this.r.Header.Get("Origin")
	matchedList := regexp.MustCompile("(\\S+)://([^:]+):?(\\d+)?").FindStringSubmatch(strings.Trim(origin, "/"))
	if 0 == len(matchedList) {
		return
	}
	for _, origin := range this.conf.Origin {
		if "*" == origin || matchedList[2] == origin || "." == origin[0:1] && matchedList[2][len(matchedList[2])-len(origin):] == origin {
			headerValue := matchedList[1] + "://" + matchedList[2]
			if 4 == len(matchedList) {
				headerValue += ":" + matchedList[3]
			}
			this.headerList["access-control-allow-origin"] = headerValue
			return
		}
	}
	interceptor.Insure(false).
		Message("跨域阻止").
		Data(map[string]interface{}{"origin": origin}).
		Do()
}
func (this *context) checkResource() {
	root := strings.TrimRight(this.conf.Root, " /") + "/"
	filepath := root + strings.Trim(this.path, " /")
	if _file.Exists(filepath) {
		this.filepath = filepath
		return
	}
	filepath = root + "index.html"
	if _file.Exists(filepath) {
		this.filepath = filepath
		return
	}
	interceptor.Insure(false).
		Message("资源不存在").
		Data(map[string]interface{}{"path": this.r.URL.Path}).
		Do()
}
func (this *context) render() {
	if len(this.e) > 0 {
		this.w.Write(this.e)
		return
	}
	this.w.Header().Set("content-type", _file.ContentType(this.filepath))
	f, err := os.Open(this.filepath)
	if nil == err {
		defer f.Close()
		io.Copy(this.w, f)
	}
}
