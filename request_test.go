package utils

import (
	"encoding/binary"
	"log"
	"os"
	"testing"
)

type (
	// BaiDuError 错误提示
	BaiDuError struct {
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
	}
	// BaiDuAccessToken AccessToken信息
	BaiDuAccessToken struct {
		RefreshToken  string `json:"refresh_token"`
		ExpiresIn     int    `json:"expires_in"`
		Scope         string `json:"scope"`
		SessionKey    string `json:"session_key"`
		AccessToken   string `json:"access_token"`
		SessionSecret string `json:"session_secret"`
		BaiDuError
	}
	// BaiDuSpeechQuick SpeechQuick信息
	BaiDuSpeechQuick struct {
		ErrNo  int      `json:"err_no"`
		ErrMsg string   `json:"err_msg"`
		SN     string   `json:"sn"`
		Result []string `json:"result"`
	}
	// BaiDuRobotDialogue 机器人对话
	BaiDuRobotDialogue struct {
		ErrorCode int    `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
		Result    struct {
			Version       string `json:"version"`
			ServiceID     string `json:"service_id"`
			LogID         string `json:"log_id"`
			InteractionID string `json:"interaction_id"`
			Response      struct {
				Status     int    `json:"status"`
				Msg        string `json:"msg"`
				ActionList []struct {
					Confidence  float64 `json:"confidence"`
					ActionID    string  `json:"action_id"`
					Say         string  `json:"say"`
					CustomReply string  `json:"custom_reply"`
					Type        string  `json:"type"`
				} `json:"action_list"`
				Schema struct {
					Confidence       float64 `json:"confidence"`
					Intent           string  `json:"intent"`
					IntentConfidence float64 `json:"intent_confidence"`
					Slots            []struct {
						Confidence     float64 `json:"confidence"`
						Begin          int     `json:"begin"`
						Length         int     `json:"length"`
						OriginalWord   string  `json:"original_word"`
						NormalizedWord string  `json:"normalized_word"`
						WordType       string  `json:"word_type"`
						Name           string  `json:"name"`
						SessionOffset  string  `json:"session_offset"`
						MergeMethod    string  `json:"merge_method"`
					} `json:"slots"`
				} `json:"schema"`
			} `json:"response"`
		} `json:"result"`
	}
)

const (
	baiDuClientID     string = "Sy0tLT7bHWE2RhollVOqelHq"
	baiDuClientSecret string = "jSr0a2Isaivi1yvgk2TXlB7tqg21Gf1m"
	//baiDuClientID     string = "MDNsII2jkUtbF729GQOZt7FS"
	//baiDuClientSecret string = "0vWCVCLsbWHMSH1wjvxaDq4VmvCZM2O9"

	// baiDuRequestURLForAccessToken 获取token地址
	baiDuRequestURLForAccessToken string = "https://aip.baidubce.com/oauth/2.0/token"
	// baiDuRequestURLForVoiceQuick 语音识别极速版
	baiDuRequestURLForVoiceQuick string = "https://vop.baidu.com/pro_api"
	// baiDuRequestURLForRobotDialogue 语音机器人对话
	baiDuRequestURLForRobotDialogue string = "https://aip.baidubce.com/rpc/2.0/unit/bot/chat"
	//baiDuRequestURLForRobotDialogue string = "https://aip.baidubce.com/rpc/2.0/unit/service/chat"
)

func TestNewClient(t *testing.T) {
	//file := "../upload/20210624/16k1.pcm"
	//
	client2 := NewClient(baiDuRequestURLForAccessToken, MethodForPost, map[string]interface{}{
		"grant_type": "client_credentials", "client_id": baiDuClientID, "client_secret": baiDuClientSecret,
	})
	resp2, err := client2.Request(RequestBodyFormatForFormData)

	if err != nil {
		t.Log(err)
		return
	}
	response := new(BaiDuAccessToken)

	_ = FromJSONBytes(resp2, response)

	if response.Error != "" {
		t.Logf("获取百度AccessToken错误：%v - %v", response.Error, response.ErrorDescription)
		return
	}
	//t.Log(response.AccessToken)
	//
	//reader, err := os.OpenFile(file, os.O_RDONLY, 0666)
	//
	//if err != nil {
	//	t.Log(err)
	//	return
	//}
	//defer reader.Close()
	//
	//content, _ := ioutil.ReadAll(reader)
	//
	//cuid := ""
	//
	//netitfs, err := net.Interfaces()
	//
	//if err != nil {
	//	cuid = "anonymous_sqzn"
	//} else {
	//	for _, itf := range netitfs {
	//		if cuid = itf.HardwareAddr.String(); len(cuid) > 0 {
	//			break
	//		}
	//	}
	//}
	//t.Log(cuid)
	//t.Log(base64.StdEncoding.EncodeToString(content))
	//t.Log(fmt.Sprintf("%d", len(content)))
	//
	//_params := map[string]interface{}{
	//	"format":  file[len(file)-3:],
	//	"rate":    16000,
	//	"dev_pid": 1537,
	//	"channel": 1,
	//	"token":   "24.1f876b06d070d7403c90832dddb813cb.2592000.1627110943.282335-24431674",
	//	"cuid":    cuid,
	//	"speech":  base64.StdEncoding.EncodeToString(content),
	//	"len":     len(content),
	//}
	//_json, _ := json.Marshal(_params)
	//
	//req, err := http.NewRequest("GET", "http://vop.baidu.com/server_api", bytes.NewBuffer(_json))
	//
	//if err != nil {
	//	t.Log(err)
	//	return
	//}
	//resp := new(http.Response)
	//
	//client := new(http.Client)
	//
	//if resp, err = client.Do(req); err != nil {
	//	t.Log(err)
	//	return
	//}
	//bytes, err := ioutil.ReadAll(resp.Body)
	//defer resp.Body.Close()
	//
	//response1 := new(BaiDuSpeechQuick)
	//
	//_ = FromJSONBytes(bytes, response1)
	//
	//t.Logf("resp：%v\n", AnyToJSON(response1))

	serviceID := "1101579"

	params2 := map[string]interface{}{
		"version": "2.0",
		//"service_id": serviceID,
		"bot_id":     serviceID,
		"log_id":     Md5String(AnyToJSON(2040256374931197952), serviceID),
		"session_id": Sha256String(AnyToJSON(2040256374931197952) + serviceID),
		"request": map[string]interface{}{
			"query":   "公告",
			"user_id": AnyToJSON(2040256374931197952),
			"query_info": map[string]interface{}{
				"asr_candidates": []string{},
				"source":         "KEYBOARD",
				"type":           "TEXT",
			},
		},
		"bernard_level": 1,
	}
	t.Log(params2)
	client3 := NewClient(baiDuRequestURLForRobotDialogue+"?access_token="+response.AccessToken,
		MethodForPost, params2)

	resp3, err := client3.Request(RequestBodyFormatForRaw, Headers{ContentType: "application/json; charset=UTF-8"})

	if err != nil {
		t.Log(err)
		return
	}
	response3 := new(BaiDuRobotDialogue)

	_ = FromJSONBytes(resp3, response3)

	t.Log(AnyToJSON(response3))

	return
	//
	//client1 := NewClient("http://vop.baidu.com/server_api", MethodForPost, params)
	//
	//resp1 := make([]byte, 0)
	//
	//if resp1, err = client1.Request(RequestBodyFormatForRaw, Headers{
	//	ContentType: "application/json",
	//}); err != nil {
	//	t.Log(err)
	//	return
	//}
	//response1 := new(BaiDuSpeechQuick)
	//
	//_ = FromJSONBytes(resp1, response1)
	//
	//t.Logf("resp：%v\n", AnyToJSON(response1))
}

func TestClient_Request(t *testing.T) {
	request := NewClient("https://image1.ljcdn.com/hdic-resblock/4494aa6e-4165-4f4a-b7ba-4ab095dd1ffa.JPG.710x400.jpg", "GET", nil)

	resp, err := request.Request(RequestBodyFormatForFormData, Headers{
		Others: map[string]string{
			"Referer":    "http://drc.hefei.gov.cn/group4/M00/07/4D/wKgEIWEM9X-AXLhsAAONk965l5o088.png",
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36",
			"Cookie":     "__jsluid_h=9fcece02433a024c638dd5e8d4cf6f92; __jsl_clearance=1628842172.968|0|8rBRZzH5SoW3MMG1%2FWkYpLUeRXA%3D",
		},
	})
	f, err := os.Create("test.jpg")
	if err != nil {
		log.Fatal("Couldn't open file")
	}
	defer f.Close()

	err = binary.Write(f, binary.LittleEndian, resp)

	if err != nil {
		log.Fatal("Write failed")
	}

	//resp, err := request.Request(RequestBodyFormatForFormData, Headers{
	//	//Others: map[string]string{
	//	//	"Referer":    "http://drc.hefei.gov.cn/group4/M00/07/4D/wKgEIWEM9X-AXLhsAAONk965l5o088.png",
	//	//	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36",
	//	//	"Cookie":     "__jsluid_h=2844bb494bad8b1cd372e28c65844abd; __jsl_clearance=1628840623.068|0|vNUfDD1V4muQrHrWy%2BmhoGbOFr0%3D",
	//	//},
	//})
	//f, err := os.Create("test.jpg")
	//if err != nil {
	//	log.Fatal("Couldn't open file")
	//}
	//defer f.Close()
	//
	//err = binary.Write(f, binary.LittleEndian, resp)
	//
	//if err != nil {
	//	log.Fatal("Write failed")
	//}
}
