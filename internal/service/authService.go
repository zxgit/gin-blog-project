package service

type GetTokenOpt struct {
	AppKey    string `json:"appKey" binding:"required"`
	AppSecret string `json:"appSecret" binding:"required"`
}

func (s Service) GetToken(a GetTokenOpt) (token string, err error) {
	return s.Dao.GetToken(a.AppKey, a.AppSecret)
}
