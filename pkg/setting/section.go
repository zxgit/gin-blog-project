package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	Timeout         time.Duration
}

type DatabaseSettingS struct {
	DBType         string
	UserName       string
	Password       string
	Host           string
	DBName         string
	TablePrefix    string
	Charset        string
	ParseTime      bool
	MaxIdleConns   int
	MaxOpenleConns int
}

type JwtSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
