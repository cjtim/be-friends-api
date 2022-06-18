package repository

import "go.uber.org/zap"

func PrepareDB() (func(), error) {
	_, err := Connect()
	if err != nil {
		zap.L().Panic("error postgresql", zap.Error(err))
		return func() {}, err
	}

	_, err = ConnectRedis(DEFAULT)
	if err != nil {
		zap.L().Panic("error redis", zap.Error(err))
		return func() {
			zap.L().Info("closed db conn",
				zap.NamedError("postgres", DB.Close()),
			)
		}, err
	}
	_, err = ConnectRedis(JWT)
	if err != nil {
		zap.L().Panic("error redis", zap.Error(err))
		return func() {
			zap.L().Info("closed db conn",
				zap.NamedError("postgres", DB.Close()),
				zap.NamedError("redis 0", RedisDefault.Client.Close()),
			)
		}, err
	}
	_, err = ConnectRedis(CALLBACK)
	if err != nil {
		zap.L().Panic("error redis", zap.Error(err))
		return func() {
			zap.L().Info("closed db conn",
				zap.NamedError("postgres", DB.Close()),
				zap.NamedError("redis 0", RedisDefault.Client.Close()),
				zap.NamedError("redis 1", RedisJwt.Client.Close()),
			)
		}, err
	}
	return func() {
		zap.L().Info("closed db conn",
			zap.NamedError("postgres", DB.Close()),
			zap.NamedError("redis 0", RedisDefault.Client.Close()),
			zap.NamedError("redis 1", RedisJwt.Client.Close()),
			zap.NamedError("redis 2", RedisCallback.Client.Close()),
		)
	}, err
}
