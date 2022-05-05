package main

import (
	"log"
	"net/http"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

var (
	smsApiClient *dysmsapi20170525.Client

	PHONE_NUM       = os.Getenv("PHONE_NUM")
	ACCESS_KEY      = os.Getenv("ACCESS_KEY")
	ACCESS_SECRET   = os.Getenv("ACCESS_SECRET")
	SIGN_NAME       = os.Getenv("SIGN_NAME")
	TEMPLATE_CODE   = os.Getenv("TEMPLATE_CODE")
	TEMPLATE_PARAMS = os.Getenv("TEMPLATE_PARAMS")
)

func init() {
	var err error
	smsApiClient, err = createClient(&ACCESS_KEY, &ACCESS_SECRET)
	if err != nil {
		log.Fatalf("createClient: %s", err)
	}
}

// Handle a serverless request
func Handle(req []byte) string {
	resp, err := sendSMS(PHONE_NUM, SIGN_NAME, TEMPLATE_CODE, TEMPLATE_PARAMS)
	if err != nil {
		log.Printf("sendSMS: %s", err)
		return err.Error()
	} else {
		return resp.String()
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	resp, err := sendSMS(PHONE_NUM, SIGN_NAME, TEMPLATE_CODE, TEMPLATE_PARAMS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp.String()))
}

func createClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func sendSMS(phoneNum, signName, templateCode, templateParam string) (_result *dysmsapi20170525.SendSmsResponse, _err error) {
	return smsApiClient.SendSms(&dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &phoneNum,
		SignName:      &signName,
		TemplateCode:  &templateCode,
		TemplateParam: &templateParam,
	})
}
