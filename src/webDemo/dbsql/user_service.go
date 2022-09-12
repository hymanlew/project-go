package dbsql

import "fmt"

func (user *User) Add() error {
	sql := "insert into user(id,name,pass,age) values(?,?,?,?)"

	//Prepare 创建一个连接用于查询或其他操作，且可以同时执行多个命令
	instmt, err := DB.Prepare(sql)
	if err != nil {
		fmt.Println("预编译出现异常", err.Error())
		return err
	}

	_, err2 := instmt.Exec(user.Id, user.Name, user.Pass, user.Age)
	if err2 != nil {
		fmt.Println("执行操作出现异常", err2.Error())
		return err2
	}

	return nil
}

func (user *User) Add2() error {
	sql := "insert into user(id,name,pass,age) values(?,?,?,?)"

	//Prepare 创建一个连接用于查询或其他操作，且可以同时执行多个命令
	_, err := DB.Exec(sql, user.Id, user.Name, user.Pass, user.Age)
	if err != nil {
		fmt.Println("执行操作出现异常", err.Error())
		return err
	}
	return nil
}

func (user *User) GetById() (*User, error) {
	sql := "select id, name, age from user where id = ?"
	row := DB.QueryRow(sql, user.Id)

	var id int
	var name string
	var age int
	err := row.Scan(&id, &name, &age)
	if err != nil {
		return nil, err
	}

	usr := &User{
		Id:   id,
		Name: name,
		Age:  age,
	}
	return usr, nil
}

func (user *User) GetAll() ([]*User, error) {
	sql := "select id, name, age from user"
	rows, err := DB.Query(sql)
	if err != nil {
		return nil, err
	}

	var users []*User

	if rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			return nil, err
		}

		usr := &User{
			Id:   id,
			Name: name,
			Age:  age,
		}
		users = append(users, usr)
	}

	for k, v := range users {
		fmt.Printf("第%d个用户是：%v \n", k+1, v)
	}
	return users, nil
}

func GetByName(name string) (*User, error) {
	sql := "select id, name, age from user where name = ?"
	rows, err := DB.Query(sql, name)
	if err != nil {
		return nil, err
	}

	user := &User{}
	rows.Scan(&user.Id, &user.Name, &user.Age)
	return user, nil
}
