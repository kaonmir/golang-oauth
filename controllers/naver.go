package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/kaonmir/OAuth/config"
)

// type callbackResponse struct {
// 	Code             int    `json:"code"`
// 	State            string `json:"state"`
// 	Error            string `json:"error"`
// 	ErrorDescription string `json:"error_description"`
// }

type accessResponse struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int    `json:"expires_in"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func requestAccessToken(target *accessResponse, code string, state string) error {
	q := url.Values{
		"client_id":     {config.Env().NaverClientID},
		"client_secret": {config.Env().NaverClientSecret},
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"state":         {state},
	}

	res, err := http.Get("https://nid.naver.com/oauth2.0/token" + "?" + q.Encode())
	if err != nil {
		return err
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(target)
	return nil
}

func NaverCallBackHandler(ctx *gin.Context) {
	code := ctx.Query("code")   // 네이버 로그인 인증에 성공하면 반환받는 인증 코드, 접근 토큰(access token) 발급에 사용
	state := ctx.Query("state") // 사이트 간 요청 위조 공격을 방지하기 위해 애플리케이션에서 생성한 상태 토큰으로 URL 인코딩을 적용한 값
	error := ctx.Query("error")
	errorDescription := ctx.Query("error_description")

	_, _ = ctx.Request.Response.Location()

	value, exists := ctx.Get("code")
	if exists {
		log.Printf("[DEBUG] Param : %+v\n", value.(string))
	}

	if error != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Naver Login Callback Error", "error": errorDescription})
		ctx.Abort()
		return
	}

	access := new(accessResponse)
	err := requestAccessToken(access, code, state)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error to request access token", "error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Naver Login Callback Success", "access": access})

}
