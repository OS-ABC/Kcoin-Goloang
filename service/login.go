package service

import (
	"Kcoin-Golang/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// 用于存储各个用户的github access token，是一个实现了互斥访问的map
var GithubAccessToken sync.Map

// 获取github的accessToken
func GetAccessToken(code string) (accessToken string, err error) {
	// TODO 需要修改为从配置文件中读取clientId和clientSecret
	url := "https://github.com/login/oauth/access_token?code=" + code + "&client_id=" + "17a8c3e212ff56ace267" + "&client_secret=" + "6725d566c37d0eeb6e27a4d57c93a16602c5cc05"
	fmt.Println(url)
	client := &http.Client{}
	response, err := client.Get(url)
	if err != nil {
		return "", err
	} else {
		defer response.Body.Close()
	}
	var body []byte
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if strings.Split(string(body), "=")[0] == "error" {
		err = errors.New("get Access Token Failed" + strings.Split(string(body), "&")[0])
		return "", err
	} else {
		accessToken = strings.Split(strings.Split(string(body), "&")[0], "=")[1]
	}
	return accessToken, err
}

// 通过从github获取的AccessToken来向githubapi请求用户的信息
func GetGithubUserInfo(accessToken string) (*models.User, error) {
	url := "https://api.github.com/user?access_token=" + accessToken
	client := &http.Client{}
	res, err := client.Get(url)
	if err != nil {
		return nil, errors.New("get github user info failed")
	}
	defer res.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	json.Unmarshal(body, user)
	return user, nil
}
