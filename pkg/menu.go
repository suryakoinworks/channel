package pkg

import (
	"fmt"
	"sync"
	"time"
)

type menu struct {
	RoleId int
	Id     int
	Name   string
	Path   string
}

func newMenu(roleId int, menuId int, name string, path string) menu {
	return menu{RoleId: roleId, Id: menuId, Name: name, Path: path}
}

func GetMenuByRoleId(roleId int) []menu {
	menus := make([]menu, 0)
	for i := 1; i <= 7; i++ {
		time.Sleep(10 * time.Millisecond)

		menus = append(menus, newMenu(roleId, i, fmt.Sprintf("Menu %d", i), fmt.Sprintf("/menu/%d", i)))
	}

	return menus
}

func FastGetMenuByRoleId(roleId int) <-chan menu {
	cMenu := make(chan menu)
	go func(roleId int) {
		wg := sync.WaitGroup{}
		for i := 1; i <= 7; i++ {
			wg.Add(1)

			go func(roleId int, menudId int, wg *sync.WaitGroup) {
				time.Sleep(10 * time.Millisecond)

				cMenu <- newMenu(roleId, menudId, fmt.Sprintf("Menu %d", menudId), fmt.Sprintf("/menu/%d", menudId))

				wg.Done()
			}(roleId, i, &wg)
		}

		wg.Wait()
		close(cMenu)
	}(roleId)

	return cMenu
}
