package models

import (
	"fmt"
)

// Method GetUserDetail used to retrieve a particular user detail by using user id key in redis db
func (Defs BasicDefs) GetUserDetail(key string) (map[string]string, error) {

	defer Defs.CtxCancel()

	user, err := Defs.DbConn.HGetAll(Defs.Ctx, key).Result()

	if err != nil {

		return map[string]string{}, err
	}

	return user, nil
}

// Method IncrementUserId used to increment the value  by 1 to the provided redis integer key in redis db
func (Defs BasicDefs) IncrementUserId() (int64, error) {

	defer Defs.CtxCancel()

	counter, err := Defs.DbConn.Incr(Defs.Ctx, UserKey).Result()

	if err != nil {

		return 0, err
	}

	return counter, nil
}

// Method SetUser used to create and update a user to the provided user key in the redis db 
func (Defs BasicDefs) SetUser(key string, user User) error {

	defer Defs.CtxCancel()

	mapData, err := ConvertStructToMap(user)

	if err != nil {

		return err
	}

	if _, err := Defs.DbConn.HSet(Defs.Ctx, key, mapData).Result(); err != nil {

		return err
	}

	return nil
}

// Method DeleteUser used to delete an existing user stored in the redis db
func (Defs BasicDefs) DeleteUser(key string) error {

	defer Defs.CtxCancel()

	if _, err := Defs.DbConn.Del(Defs.Ctx, key).Result(); err != nil {

		fmt.Println("err", err)

		return err
	}

	return nil
}

// Method DecrementUserId used to delete the assigned user key after the user details successfully deleted fron the redis db
func (Defs BasicDefs) DecrementUserId() error {

	defer Defs.CtxCancel()

	if _, err := Defs.DbConn.Decr(Defs.Ctx, UserKey).Result(); err != nil {

		return err
	}

	return nil
}
 
// Method DeleteUserTransaction used as a pipeline to execute both delete a user details followed removing the user key assigned the deleted user in a single db transaction
func (Defs BasicDefs) DeleteUserTransaction(key string)error{

	defer Defs.CtxCancel()

	pipeline := Defs.DbConn.TxPipeline()

	pipeline.Del(Defs.Ctx,key)

	pipeline.Decr(Defs.Ctx,UserKey)

	if _, err := pipeline.Exec(Defs.Ctx);err!= nil{

		return err
	}

	return nil
}
