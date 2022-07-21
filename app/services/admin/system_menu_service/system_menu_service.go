package system_menu_service

type MenuService struct {
	RoleIds []int
}

type IMenuService interface {
	findRoleMenus()
}

func (menuService *MenuService) findRoleMenus() {

}
