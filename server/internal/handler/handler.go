package handler

import (
	"fmt"
	"net/http"

	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

type AWSConfig struct {
	Region				string
	AccessKey			string
	SecretKey			string
	BucketName		string
}

type EmailConfig struct {
	Host string
	Port int
	User string
	Pass string
}

type RecaptchaConfig struct {
	SecretKey string
}

type Config struct {
	PortalUrl			string
	AppName				string
	Version				string
	EmailConf			*EmailConfig
	AWSConf				*AWSConfig
	RecaptchaConf	*RecaptchaConfig
}

type Controller struct {
	Config	*Config
	User		*model.User
}

var Conf *Config
var Ctrl *Controller

func NewAWSConfig(region, accessKey, secretKey, bucketName string) *AWSConfig {
	return &AWSConfig{
		Region: region,
		AccessKey: accessKey,
		SecretKey: secretKey,
		BucketName: bucketName,
	}
}

func NewRecaptchaConfig(secretKey string) *RecaptchaConfig {
	return &RecaptchaConfig{
		SecretKey: secretKey,
	}
}

func NewEmailConfig(host string, port int, user, pass string) *EmailConfig {
	return &EmailConfig{
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
	}
}

func NewConfig(portalUrl, appName, version string, emailConf *EmailConfig, awsConf *AWSConfig, recaptchaConf *RecaptchaConfig) *Config {
	return &Config{
		PortalUrl: portalUrl,
		AppName: appName,
		Version: version,
		EmailConf: emailConf,
		AWSConf: awsConf,
		RecaptchaConf: recaptchaConf,
	}
}

func NewHandler(c *Config) {
	Ctrl = &Controller{
		Config: c,
	}
}

func SetUser(u *model.User) {
	Ctrl.User = u
}

func CheckTokenAlive() bool {
	var t model.TokenInterface = &model.Token{}
	token, err := t.Query(Ctrl.User.Token)

	if err != nil {
		return false
	}

	if !token.IsAlive {
		return false
	}

	return true
}

func CheckToken(w http.ResponseWriter, r *http.Request) bool {
	var t model.TokenInterface = &model.Token{}
	token, err := t.Query(Ctrl.User.Token)

	fmt.Println(err)
	fmt.Println(token.IsAlive)
	if err != nil || !token.IsAlive {
		fmt.Println(1)
		return false
	}

	fmt.Println(2)
	return true
}

func GetUnauthorizedResponse() util.ResponseFormat {
	err := map[string]interface{}{
		"code": http.StatusUnauthorized,
		"message": "Unauthorized",
	}
	res := util.GetResponse(nil, err)

	return res
}

func RenderTemplate(w http.ResponseWriter, tmplPath string, data interface{})  {
	tmpl, err := util.ParseTemplate(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}