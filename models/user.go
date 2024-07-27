package models

import (
	"fmt"
	"reflect"
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

func (Defs BasicDefs) AutoIncrementUserId() (int64, error) {

	defer Defs.CtxCancel()

	counter, err := Defs.DbConn.Incr(Defs.Ctx, UserKey).Result()

	if err != nil {

		fmt.Println("err", err)

		return 0, err
	}

	return counter, nil
}

func (Defs BasicDefs) CreateUser(key string, user User) error {

	defer Defs.CtxCancel()

    mapData := make(map[string]interface{})

	v := reflect.ValueOf(user)

	t := reflect.TypeOf(user)

	for i:=0;i<v.NumField();i++{

		field := t.Field(i)

		tagName := field.Tag.Get("json")

		fieldValue := v.Field(i)

		mapData[tagName] = fieldValue.Interface()
	}

	fmt.Println("mapData",mapData)

	if _, err := Defs.DbConn.HSet(Defs.Ctx, key, mapData).Result(); err != nil {

		fmt.Println("err", err)

		return err
	}

	return nil
}
