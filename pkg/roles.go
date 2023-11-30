package pkg

import (
	"fmt"
	"sync"
	"time"
)

type Role struct {
	UserId int
	Id     int
	Name   string
	Menus  []menu
}

func newRole(userId int, roleId int, name string, menus ...menu) Role {
	return Role{UserId: userId, Id: roleId, Name: name, Menus: menus}
}

func GetRoleByUserId(userId int) []Role {
	roles := make([]Role, 0)
	for i := 1; i <= 7; i++ {
		time.Sleep(10 * time.Millisecond)

		roles = append(roles, newRole(userId, i, fmt.Sprintf("Role %d", i), GetMenuByRoleId(i)...))
	}

	return roles
}

func FastGetRoleByUserId(userId int) <-chan Role {
	cRole := make(chan Role)
	go func(userId int) {
		wg := sync.WaitGroup{}
		for i := 1; i <= 7; i++ {
			wg.Add(1)

			go func(userId int, roleId int, wg *sync.WaitGroup) {
				time.Sleep(10 * time.Millisecond)

				menus := []menu{}
				cMenu := FastGetMenuByRoleId(roleId)
				for menu := range cMenu {
					menus = append(menus, menu)
				}

				cRole <- newRole(userId, roleId, fmt.Sprintf("Role %d", roleId), menus...)

				wg.Done()
			}(userId, i, &wg)
		}

		wg.Wait()
		close(cRole)
	}(userId)

	return cRole
}
