// Code generated by goctl. DO NOT EDIT.
// Source: core.proto

package coreclient

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiInfo                  = core.ApiInfo
	ApiListReq               = core.ApiListReq
	ApiListResp              = core.ApiListResp
	BaseIDResp               = core.BaseIDResp
	BaseMsg                  = core.BaseMsg
	BaseResp                 = core.BaseResp
	BaseUUIDResp             = core.BaseUUIDResp
	CallbackReq              = core.CallbackReq
	DepartmentInfo           = core.DepartmentInfo
	DepartmentListReq        = core.DepartmentListReq
	DepartmentListResp       = core.DepartmentListResp
	DictionaryDetailInfo     = core.DictionaryDetailInfo
	DictionaryDetailListReq  = core.DictionaryDetailListReq
	DictionaryDetailListResp = core.DictionaryDetailListResp
	DictionaryInfo           = core.DictionaryInfo
	DictionaryListReq        = core.DictionaryListReq
	DictionaryListResp       = core.DictionaryListResp
	Empty                    = core.Empty
	IDReq                    = core.IDReq
	IDsReq                   = core.IDsReq
	MenuInfo                 = core.MenuInfo
	MenuInfoList             = core.MenuInfoList
	MenuRoleInfo             = core.MenuRoleInfo
	MenuRoleListResp         = core.MenuRoleListResp
	Meta                     = core.Meta
	OauthLoginReq            = core.OauthLoginReq
	OauthProviderInfo        = core.OauthProviderInfo
	OauthProviderListReq     = core.OauthProviderListReq
	OauthProviderListResp    = core.OauthProviderListResp
	OauthRedirectResp        = core.OauthRedirectResp
	PageInfoReq              = core.PageInfoReq
	PositionInfo             = core.PositionInfo
	PositionListReq          = core.PositionListReq
	PositionListResp         = core.PositionListResp
	RoleInfo                 = core.RoleInfo
	RoleListReq              = core.RoleListReq
	RoleListResp             = core.RoleListResp
	RoleMenuAuthorityReq     = core.RoleMenuAuthorityReq
	RoleMenuAuthorityResp    = core.RoleMenuAuthorityResp
	StockInfo                = core.StockInfo
	StockListReq             = core.StockListReq
	StockListResp            = core.StockListResp
	StockUserInfo            = core.StockUserInfo
	StockUserListReq         = core.StockUserListReq
	StockUserListResp        = core.StockUserListResp
	StockUserMobileReq       = core.StockUserMobileReq
	StockUsernameReq         = core.StockUsernameReq
	TokenInfo                = core.TokenInfo
	TokenListReq             = core.TokenListReq
	TokenListResp            = core.TokenListResp
	UUIDReq                  = core.UUIDReq
	UUIDsReq                 = core.UUIDsReq
	UserInfo                 = core.UserInfo
	UserListReq              = core.UserListReq
	UserListResp             = core.UserListResp
	UsernameReq              = core.UsernameReq

	Core interface {
		// API management
		CreateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetApiList(ctx context.Context, in *ApiListReq, opts ...grpc.CallOption) (*ApiListResp, error)
		GetApiById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*ApiInfo, error)
		DeleteApi(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetMenuAuthority(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleMenuAuthorityResp, error)
		CreateOrUpdateMenuAuthority(ctx context.Context, in *RoleMenuAuthorityReq, opts ...grpc.CallOption) (*BaseResp, error)
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// Department management
		CreateDepartment(ctx context.Context, in *DepartmentInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateDepartment(ctx context.Context, in *DepartmentInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetDepartmentList(ctx context.Context, in *DepartmentListReq, opts ...grpc.CallOption) (*DepartmentListResp, error)
		GetDepartmentById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DepartmentInfo, error)
		DeleteDepartment(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Dictionary management
		CreateDictionary(ctx context.Context, in *DictionaryInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateDictionary(ctx context.Context, in *DictionaryInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetDictionaryList(ctx context.Context, in *DictionaryListReq, opts ...grpc.CallOption) (*DictionaryListResp, error)
		GetDictionaryById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DictionaryInfo, error)
		DeleteDictionary(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// DictionaryDetail management
		CreateDictionaryDetail(ctx context.Context, in *DictionaryDetailInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateDictionaryDetail(ctx context.Context, in *DictionaryDetailInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetDictionaryDetailList(ctx context.Context, in *DictionaryDetailListReq, opts ...grpc.CallOption) (*DictionaryDetailListResp, error)
		GetDictionaryDetailById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DictionaryDetailInfo, error)
		DeleteDictionaryDetail(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetDictionaryDetailByDictionaryName(ctx context.Context, in *BaseMsg, opts ...grpc.CallOption) (*DictionaryDetailListResp, error)
		CreateMenu(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateMenu(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteMenu(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetMenuListByRole(ctx context.Context, in *BaseMsg, opts ...grpc.CallOption) (*MenuInfoList, error)
		GetMenuList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*MenuInfoList, error)
		// OauthProvider management
		CreateOauthProvider(ctx context.Context, in *OauthProviderInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateOauthProvider(ctx context.Context, in *OauthProviderInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetOauthProviderList(ctx context.Context, in *OauthProviderListReq, opts ...grpc.CallOption) (*OauthProviderListResp, error)
		GetOauthProviderById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*OauthProviderInfo, error)
		DeleteOauthProvider(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		OauthLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*OauthRedirectResp, error)
		OauthCallback(ctx context.Context, in *CallbackReq, opts ...grpc.CallOption) (*UserInfo, error)
		// Position management
		CreatePosition(ctx context.Context, in *PositionInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdatePosition(ctx context.Context, in *PositionInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetPositionList(ctx context.Context, in *PositionListReq, opts ...grpc.CallOption) (*PositionListResp, error)
		GetPositionById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*PositionInfo, error)
		DeletePosition(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Role management
		CreateRole(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateRole(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetRoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error)
		GetRoleById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleInfo, error)
		DeleteRole(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Stock management
		CreateStock(ctx context.Context, in *StockInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		UpdateStock(ctx context.Context, in *StockInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetStockList(ctx context.Context, in *StockListReq, opts ...grpc.CallOption) (*StockListResp, error)
		GetStockById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*StockInfo, error)
		DeleteStock(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// StockUser management
		CreateStockUser(ctx context.Context, in *StockUserInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		UpdateStockUser(ctx context.Context, in *StockUserInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetStockUserList(ctx context.Context, in *StockUserListReq, opts ...grpc.CallOption) (*StockUserListResp, error)
		GetStockUserById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*StockUserInfo, error)
		GetStockUserByUsername(ctx context.Context, in *StockUsernameReq, opts ...grpc.CallOption) (*StockUserInfo, error)
		GetStockUserByMobile(ctx context.Context, in *StockUserMobileReq, opts ...grpc.CallOption) (*StockUserInfo, error)
		DeleteStockUser(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Token management
		CreateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		DeleteToken(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetTokenList(ctx context.Context, in *TokenListReq, opts ...grpc.CallOption) (*TokenListResp, error)
		GetTokenById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*TokenInfo, error)
		BlockUserAllToken(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*BaseResp, error)
		UpdateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error)
		// User management
		CreateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		UpdateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetUserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		GetUserById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*UserInfo, error)
		GetUserByUsername(ctx context.Context, in *UsernameReq, opts ...grpc.CallOption) (*UserInfo, error)
		DeleteUser(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultCore struct {
		cli zrpc.Client
	}
)

func NewCore(cli zrpc.Client) Core {
	return &defaultCore{
		cli: cli,
	}
}

// API management
func (m *defaultCore) CreateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateApi(ctx, in, opts...)
}

func (m *defaultCore) UpdateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateApi(ctx, in, opts...)
}

func (m *defaultCore) GetApiList(ctx context.Context, in *ApiListReq, opts ...grpc.CallOption) (*ApiListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetApiList(ctx, in, opts...)
}

func (m *defaultCore) GetApiById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*ApiInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetApiById(ctx, in, opts...)
}

func (m *defaultCore) DeleteApi(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteApi(ctx, in, opts...)
}

func (m *defaultCore) GetMenuAuthority(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleMenuAuthorityResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetMenuAuthority(ctx, in, opts...)
}

func (m *defaultCore) CreateOrUpdateMenuAuthority(ctx context.Context, in *RoleMenuAuthorityReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateOrUpdateMenuAuthority(ctx, in, opts...)
}

func (m *defaultCore) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

// Department management
func (m *defaultCore) CreateDepartment(ctx context.Context, in *DepartmentInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateDepartment(ctx, in, opts...)
}

func (m *defaultCore) UpdateDepartment(ctx context.Context, in *DepartmentInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateDepartment(ctx, in, opts...)
}

func (m *defaultCore) GetDepartmentList(ctx context.Context, in *DepartmentListReq, opts ...grpc.CallOption) (*DepartmentListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetDepartmentList(ctx, in, opts...)
}

func (m *defaultCore) GetDepartmentById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DepartmentInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetDepartmentById(ctx, in, opts...)
}

func (m *defaultCore) DeleteDepartment(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteDepartment(ctx, in, opts...)
}

// Dictionary management
func (m *defaultCore) CreateDictionary(ctx context.Context, in *DictionaryInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateDictionary(ctx, in, opts...)
}

func (m *defaultCore) UpdateDictionary(ctx context.Context, in *DictionaryInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateDictionary(ctx, in, opts...)
}

func (m *defaultCore) GetDictionaryList(ctx context.Context, in *DictionaryListReq, opts ...grpc.CallOption) (*DictionaryListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetDictionaryList(ctx, in, opts...)
}

func (m *defaultCore) GetDictionaryById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DictionaryInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetDictionaryById(ctx, in, opts...)
}

func (m *defaultCore) DeleteDictionary(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteDictionary(ctx, in, opts...)
}

// DictionaryDetail management
func (m *defaultCore) CreateDictionaryDetail(ctx context.Context, in *DictionaryDetailInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateDictionaryDetail(ctx, in, opts...)
}

func (m *defaultCore) UpdateDictionaryDetail(ctx context.Context, in *DictionaryDetailInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateDictionaryDetail(ctx, in, opts...)
}

func (m *defaultCore) GetDictionaryDetailList(ctx context.Context, in *DictionaryDetailListReq, opts ...grpc.CallOption) (*DictionaryDetailListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetDictionaryDetailList(ctx, in, opts...)
}

func (m *defaultCore) GetDictionaryDetailById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DictionaryDetailInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetDictionaryDetailById(ctx, in, opts...)
}

func (m *defaultCore) DeleteDictionaryDetail(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteDictionaryDetail(ctx, in, opts...)
}

func (m *defaultCore) GetDictionaryDetailByDictionaryName(ctx context.Context, in *BaseMsg, opts ...grpc.CallOption) (*DictionaryDetailListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetDictionaryDetailByDictionaryName(ctx, in, opts...)
}

func (m *defaultCore) CreateMenu(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateMenu(ctx, in, opts...)
}

func (m *defaultCore) UpdateMenu(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateMenu(ctx, in, opts...)
}

func (m *defaultCore) DeleteMenu(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteMenu(ctx, in, opts...)
}

func (m *defaultCore) GetMenuListByRole(ctx context.Context, in *BaseMsg, opts ...grpc.CallOption) (*MenuInfoList, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetMenuListByRole(ctx, in, opts...)
}

func (m *defaultCore) GetMenuList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*MenuInfoList, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetMenuList(ctx, in, opts...)
}

// OauthProvider management
func (m *defaultCore) CreateOauthProvider(ctx context.Context, in *OauthProviderInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateOauthProvider(ctx, in, opts...)
}

func (m *defaultCore) UpdateOauthProvider(ctx context.Context, in *OauthProviderInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateOauthProvider(ctx, in, opts...)
}

func (m *defaultCore) GetOauthProviderList(ctx context.Context, in *OauthProviderListReq, opts ...grpc.CallOption) (*OauthProviderListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetOauthProviderList(ctx, in, opts...)
}

func (m *defaultCore) GetOauthProviderById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*OauthProviderInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetOauthProviderById(ctx, in, opts...)
}

func (m *defaultCore) DeleteOauthProvider(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteOauthProvider(ctx, in, opts...)
}

func (m *defaultCore) OauthLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*OauthRedirectResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.OauthLogin(ctx, in, opts...)
}

func (m *defaultCore) OauthCallback(ctx context.Context, in *CallbackReq, opts ...grpc.CallOption) (*UserInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.OauthCallback(ctx, in, opts...)
}

// Position management
func (m *defaultCore) CreatePosition(ctx context.Context, in *PositionInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreatePosition(ctx, in, opts...)
}

func (m *defaultCore) UpdatePosition(ctx context.Context, in *PositionInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdatePosition(ctx, in, opts...)
}

func (m *defaultCore) GetPositionList(ctx context.Context, in *PositionListReq, opts ...grpc.CallOption) (*PositionListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetPositionList(ctx, in, opts...)
}

func (m *defaultCore) GetPositionById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*PositionInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetPositionById(ctx, in, opts...)
}

func (m *defaultCore) DeletePosition(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeletePosition(ctx, in, opts...)
}

// Role management
func (m *defaultCore) CreateRole(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateRole(ctx, in, opts...)
}

func (m *defaultCore) UpdateRole(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

func (m *defaultCore) GetRoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetRoleList(ctx, in, opts...)
}

func (m *defaultCore) GetRoleById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetRoleById(ctx, in, opts...)
}

func (m *defaultCore) DeleteRole(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteRole(ctx, in, opts...)
}

// Stock management
func (m *defaultCore) CreateStock(ctx context.Context, in *StockInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateStock(ctx, in, opts...)
}

func (m *defaultCore) UpdateStock(ctx context.Context, in *StockInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateStock(ctx, in, opts...)
}

func (m *defaultCore) GetStockList(ctx context.Context, in *StockListReq, opts ...grpc.CallOption) (*StockListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetStockList(ctx, in, opts...)
}

func (m *defaultCore) GetStockById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*StockInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetStockById(ctx, in, opts...)
}

func (m *defaultCore) DeleteStock(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteStock(ctx, in, opts...)
}

// StockUser management
func (m *defaultCore) CreateStockUser(ctx context.Context, in *StockUserInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateStockUser(ctx, in, opts...)
}

func (m *defaultCore) UpdateStockUser(ctx context.Context, in *StockUserInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateStockUser(ctx, in, opts...)
}

func (m *defaultCore) GetStockUserList(ctx context.Context, in *StockUserListReq, opts ...grpc.CallOption) (*StockUserListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetStockUserList(ctx, in, opts...)
}

func (m *defaultCore) GetStockUserById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*StockUserInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetStockUserById(ctx, in, opts...)
}

func (m *defaultCore) GetStockUserByUsername(ctx context.Context, in *StockUsernameReq, opts ...grpc.CallOption) (*StockUserInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetStockUserByUsername(ctx, in, opts...)
}

func (m *defaultCore) GetStockUserByMobile(ctx context.Context, in *StockUserMobileReq, opts ...grpc.CallOption) (*StockUserInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetStockUserByMobile(ctx, in, opts...)
}

func (m *defaultCore) DeleteStockUser(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteStockUser(ctx, in, opts...)
}

// Token management
func (m *defaultCore) CreateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateToken(ctx, in, opts...)
}

func (m *defaultCore) DeleteToken(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteToken(ctx, in, opts...)
}

func (m *defaultCore) GetTokenList(ctx context.Context, in *TokenListReq, opts ...grpc.CallOption) (*TokenListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetTokenList(ctx, in, opts...)
}

func (m *defaultCore) GetTokenById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*TokenInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetTokenById(ctx, in, opts...)
}

func (m *defaultCore) BlockUserAllToken(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.BlockUserAllToken(ctx, in, opts...)
}

func (m *defaultCore) UpdateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateToken(ctx, in, opts...)
}

// User management
func (m *defaultCore) CreateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultCore) UpdateUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultCore) GetUserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetUserList(ctx, in, opts...)
}

func (m *defaultCore) GetUserById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*UserInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetUserById(ctx, in, opts...)
}

func (m *defaultCore) GetUserByUsername(ctx context.Context, in *UsernameReq, opts ...grpc.CallOption) (*UserInfo, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.GetUserByUsername(ctx, in, opts...)
}

func (m *defaultCore) DeleteUser(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := core.NewCoreClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}
