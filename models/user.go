package models

import (
	"fmt"
)

func (Defs BasicDefs) GetUserDetail(key string) (map[string]string, error) {

	defer Defs.CtxCancel()

	user, err := Defs.DbConn.HGetAll(Defs.Ctx, key).Result()

	if err != nil {

		fmt.Println("err", err)

		return map[string]string{}, err
	}

	return user, nil
}

func (Defs BasicDefs) IncrementUserId() (int64, error) {

	defer Defs.CtxCancel()

	counter, err := Defs.DbConn.Incr(Defs.Ctx, UserKey).Result()

	if err != nil {

		fmt.Println("err", err)

		return 0, err
	}

	return counter, nil
}

func (Defs BasicDefs) SetUser(key string, user User) error {

	defer Defs.CtxCancel()

	mapData, err := ConvertStructToMap(user)

	if err != nil {

		fmt.Println("err", err)

		return err
	}

	if _, err := Defs.DbConn.HSet(Defs.Ctx, key, mapData).Result(); err != nil {

		fmt.Println("err", err)

		return err
	}

	return nil
}

func (Defs BasicDefs) DeleteUser(key string) error {

	defer Defs.CtxCancel()

	if _, err := Defs.DbConn.Del(Defs.Ctx, key).Result(); err != nil {

		fmt.Println("err", err)

		return err
	}

	return nil
}

func (Defs BasicDefs) DecrementUserId() error {

	defer Defs.CtxCancel()

	if _, err := Defs.DbConn.Decr(Defs.Ctx, UserKey).Result(); err != nil {

		return err
	}

	return nil
}

func (Defs BasicDefs) DeleteUserTransaction(key string)error{

	defer Defs.CtxCancel()

	pipeline := Defs.DbConn.TxPipeline()

	pipeline.Del(Defs.Ctx,key)

	pipeline.Decr(Defs.Ctx,UserKey)

	if _, err := pipeline.Exec(Defs.Ctx);err!= nil{

		fmt.Println("err",err)

		return err
	}

	return nil
}
