// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/api"
	"github.com/suyuan32/simple-admin-core/rpc/ent/department"
	"github.com/suyuan32/simple-admin-core/rpc/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/rpc/ent/dictionarydetail"
	"github.com/suyuan32/simple-admin-core/rpc/ent/menu"
	"github.com/suyuan32/simple-admin-core/rpc/ent/oauthprovider"
	"github.com/suyuan32/simple-admin-core/rpc/ent/position"
	"github.com/suyuan32/simple-admin-core/rpc/ent/role"
	"github.com/suyuan32/simple-admin-core/rpc/ent/schema"
	"github.com/suyuan32/simple-admin-core/rpc/ent/stock"
	"github.com/suyuan32/simple-admin-core/rpc/ent/token"
	"github.com/suyuan32/simple-admin-core/rpc/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	apiMixin := schema.API{}.Mixin()
	apiMixinFields0 := apiMixin[0].Fields()
	_ = apiMixinFields0
	apiFields := schema.API{}.Fields()
	_ = apiFields
	// apiDescCreatedAt is the schema descriptor for created_at field.
	apiDescCreatedAt := apiMixinFields0[1].Descriptor()
	// api.DefaultCreatedAt holds the default value on creation for the created_at field.
	api.DefaultCreatedAt = apiDescCreatedAt.Default.(func() time.Time)
	// apiDescUpdatedAt is the schema descriptor for updated_at field.
	apiDescUpdatedAt := apiMixinFields0[2].Descriptor()
	// api.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	api.DefaultUpdatedAt = apiDescUpdatedAt.Default.(func() time.Time)
	// api.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	api.UpdateDefaultUpdatedAt = apiDescUpdatedAt.UpdateDefault.(func() time.Time)
	// apiDescMethod is the schema descriptor for method field.
	apiDescMethod := apiFields[3].Descriptor()
	// api.DefaultMethod holds the default value on creation for the method field.
	api.DefaultMethod = apiDescMethod.Default.(string)
	// apiDescIsRequired is the schema descriptor for is_required field.
	apiDescIsRequired := apiFields[4].Descriptor()
	// api.DefaultIsRequired holds the default value on creation for the is_required field.
	api.DefaultIsRequired = apiDescIsRequired.Default.(bool)
	departmentMixin := schema.Department{}.Mixin()
	departmentMixinFields0 := departmentMixin[0].Fields()
	_ = departmentMixinFields0
	departmentMixinFields1 := departmentMixin[1].Fields()
	_ = departmentMixinFields1
	departmentMixinFields2 := departmentMixin[2].Fields()
	_ = departmentMixinFields2
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescCreatedAt is the schema descriptor for created_at field.
	departmentDescCreatedAt := departmentMixinFields0[1].Descriptor()
	// department.DefaultCreatedAt holds the default value on creation for the created_at field.
	department.DefaultCreatedAt = departmentDescCreatedAt.Default.(func() time.Time)
	// departmentDescUpdatedAt is the schema descriptor for updated_at field.
	departmentDescUpdatedAt := departmentMixinFields0[2].Descriptor()
	// department.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	department.DefaultUpdatedAt = departmentDescUpdatedAt.Default.(func() time.Time)
	// department.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	department.UpdateDefaultUpdatedAt = departmentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// departmentDescStatus is the schema descriptor for status field.
	departmentDescStatus := departmentMixinFields1[0].Descriptor()
	// department.DefaultStatus holds the default value on creation for the status field.
	department.DefaultStatus = departmentDescStatus.Default.(uint8)
	// departmentDescSort is the schema descriptor for sort field.
	departmentDescSort := departmentMixinFields2[0].Descriptor()
	// department.DefaultSort holds the default value on creation for the sort field.
	department.DefaultSort = departmentDescSort.Default.(uint32)
	// departmentDescParentID is the schema descriptor for parent_id field.
	departmentDescParentID := departmentFields[6].Descriptor()
	// department.DefaultParentID holds the default value on creation for the parent_id field.
	department.DefaultParentID = departmentDescParentID.Default.(uint64)
	dictionaryMixin := schema.Dictionary{}.Mixin()
	dictionaryMixinFields0 := dictionaryMixin[0].Fields()
	_ = dictionaryMixinFields0
	dictionaryMixinFields1 := dictionaryMixin[1].Fields()
	_ = dictionaryMixinFields1
	dictionaryFields := schema.Dictionary{}.Fields()
	_ = dictionaryFields
	// dictionaryDescCreatedAt is the schema descriptor for created_at field.
	dictionaryDescCreatedAt := dictionaryMixinFields0[1].Descriptor()
	// dictionary.DefaultCreatedAt holds the default value on creation for the created_at field.
	dictionary.DefaultCreatedAt = dictionaryDescCreatedAt.Default.(func() time.Time)
	// dictionaryDescUpdatedAt is the schema descriptor for updated_at field.
	dictionaryDescUpdatedAt := dictionaryMixinFields0[2].Descriptor()
	// dictionary.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	dictionary.DefaultUpdatedAt = dictionaryDescUpdatedAt.Default.(func() time.Time)
	// dictionary.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	dictionary.UpdateDefaultUpdatedAt = dictionaryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// dictionaryDescStatus is the schema descriptor for status field.
	dictionaryDescStatus := dictionaryMixinFields1[0].Descriptor()
	// dictionary.DefaultStatus holds the default value on creation for the status field.
	dictionary.DefaultStatus = dictionaryDescStatus.Default.(uint8)
	dictionarydetailMixin := schema.DictionaryDetail{}.Mixin()
	dictionarydetailMixinFields0 := dictionarydetailMixin[0].Fields()
	_ = dictionarydetailMixinFields0
	dictionarydetailMixinFields1 := dictionarydetailMixin[1].Fields()
	_ = dictionarydetailMixinFields1
	dictionarydetailMixinFields2 := dictionarydetailMixin[2].Fields()
	_ = dictionarydetailMixinFields2
	dictionarydetailFields := schema.DictionaryDetail{}.Fields()
	_ = dictionarydetailFields
	// dictionarydetailDescCreatedAt is the schema descriptor for created_at field.
	dictionarydetailDescCreatedAt := dictionarydetailMixinFields0[1].Descriptor()
	// dictionarydetail.DefaultCreatedAt holds the default value on creation for the created_at field.
	dictionarydetail.DefaultCreatedAt = dictionarydetailDescCreatedAt.Default.(func() time.Time)
	// dictionarydetailDescUpdatedAt is the schema descriptor for updated_at field.
	dictionarydetailDescUpdatedAt := dictionarydetailMixinFields0[2].Descriptor()
	// dictionarydetail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	dictionarydetail.DefaultUpdatedAt = dictionarydetailDescUpdatedAt.Default.(func() time.Time)
	// dictionarydetail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	dictionarydetail.UpdateDefaultUpdatedAt = dictionarydetailDescUpdatedAt.UpdateDefault.(func() time.Time)
	// dictionarydetailDescStatus is the schema descriptor for status field.
	dictionarydetailDescStatus := dictionarydetailMixinFields1[0].Descriptor()
	// dictionarydetail.DefaultStatus holds the default value on creation for the status field.
	dictionarydetail.DefaultStatus = dictionarydetailDescStatus.Default.(uint8)
	// dictionarydetailDescSort is the schema descriptor for sort field.
	dictionarydetailDescSort := dictionarydetailMixinFields2[0].Descriptor()
	// dictionarydetail.DefaultSort holds the default value on creation for the sort field.
	dictionarydetail.DefaultSort = dictionarydetailDescSort.Default.(uint32)
	menuMixin := schema.Menu{}.Mixin()
	menuMixinFields0 := menuMixin[0].Fields()
	_ = menuMixinFields0
	menuMixinFields1 := menuMixin[1].Fields()
	_ = menuMixinFields1
	menuFields := schema.Menu{}.Fields()
	_ = menuFields
	// menuDescCreatedAt is the schema descriptor for created_at field.
	menuDescCreatedAt := menuMixinFields0[1].Descriptor()
	// menu.DefaultCreatedAt holds the default value on creation for the created_at field.
	menu.DefaultCreatedAt = menuDescCreatedAt.Default.(func() time.Time)
	// menuDescUpdatedAt is the schema descriptor for updated_at field.
	menuDescUpdatedAt := menuMixinFields0[2].Descriptor()
	// menu.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	menu.DefaultUpdatedAt = menuDescUpdatedAt.Default.(func() time.Time)
	// menu.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	menu.UpdateDefaultUpdatedAt = menuDescUpdatedAt.UpdateDefault.(func() time.Time)
	// menuDescSort is the schema descriptor for sort field.
	menuDescSort := menuMixinFields1[0].Descriptor()
	// menu.DefaultSort holds the default value on creation for the sort field.
	menu.DefaultSort = menuDescSort.Default.(uint32)
	// menuDescParentID is the schema descriptor for parent_id field.
	menuDescParentID := menuFields[0].Descriptor()
	// menu.DefaultParentID holds the default value on creation for the parent_id field.
	menu.DefaultParentID = menuDescParentID.Default.(uint64)
	// menuDescPath is the schema descriptor for path field.
	menuDescPath := menuFields[3].Descriptor()
	// menu.DefaultPath holds the default value on creation for the path field.
	menu.DefaultPath = menuDescPath.Default.(string)
	// menuDescRedirect is the schema descriptor for redirect field.
	menuDescRedirect := menuFields[5].Descriptor()
	// menu.DefaultRedirect holds the default value on creation for the redirect field.
	menu.DefaultRedirect = menuDescRedirect.Default.(string)
	// menuDescComponent is the schema descriptor for component field.
	menuDescComponent := menuFields[6].Descriptor()
	// menu.DefaultComponent holds the default value on creation for the component field.
	menu.DefaultComponent = menuDescComponent.Default.(string)
	// menuDescDisabled is the schema descriptor for disabled field.
	menuDescDisabled := menuFields[7].Descriptor()
	// menu.DefaultDisabled holds the default value on creation for the disabled field.
	menu.DefaultDisabled = menuDescDisabled.Default.(bool)
	// menuDescHideMenu is the schema descriptor for hide_menu field.
	menuDescHideMenu := menuFields[10].Descriptor()
	// menu.DefaultHideMenu holds the default value on creation for the hide_menu field.
	menu.DefaultHideMenu = menuDescHideMenu.Default.(bool)
	// menuDescHideBreadcrumb is the schema descriptor for hide_breadcrumb field.
	menuDescHideBreadcrumb := menuFields[11].Descriptor()
	// menu.DefaultHideBreadcrumb holds the default value on creation for the hide_breadcrumb field.
	menu.DefaultHideBreadcrumb = menuDescHideBreadcrumb.Default.(bool)
	// menuDescIgnoreKeepAlive is the schema descriptor for ignore_keep_alive field.
	menuDescIgnoreKeepAlive := menuFields[12].Descriptor()
	// menu.DefaultIgnoreKeepAlive holds the default value on creation for the ignore_keep_alive field.
	menu.DefaultIgnoreKeepAlive = menuDescIgnoreKeepAlive.Default.(bool)
	// menuDescHideTab is the schema descriptor for hide_tab field.
	menuDescHideTab := menuFields[13].Descriptor()
	// menu.DefaultHideTab holds the default value on creation for the hide_tab field.
	menu.DefaultHideTab = menuDescHideTab.Default.(bool)
	// menuDescFrameSrc is the schema descriptor for frame_src field.
	menuDescFrameSrc := menuFields[14].Descriptor()
	// menu.DefaultFrameSrc holds the default value on creation for the frame_src field.
	menu.DefaultFrameSrc = menuDescFrameSrc.Default.(string)
	// menuDescCarryParam is the schema descriptor for carry_param field.
	menuDescCarryParam := menuFields[15].Descriptor()
	// menu.DefaultCarryParam holds the default value on creation for the carry_param field.
	menu.DefaultCarryParam = menuDescCarryParam.Default.(bool)
	// menuDescHideChildrenInMenu is the schema descriptor for hide_children_in_menu field.
	menuDescHideChildrenInMenu := menuFields[16].Descriptor()
	// menu.DefaultHideChildrenInMenu holds the default value on creation for the hide_children_in_menu field.
	menu.DefaultHideChildrenInMenu = menuDescHideChildrenInMenu.Default.(bool)
	// menuDescAffix is the schema descriptor for affix field.
	menuDescAffix := menuFields[17].Descriptor()
	// menu.DefaultAffix holds the default value on creation for the affix field.
	menu.DefaultAffix = menuDescAffix.Default.(bool)
	// menuDescDynamicLevel is the schema descriptor for dynamic_level field.
	menuDescDynamicLevel := menuFields[18].Descriptor()
	// menu.DefaultDynamicLevel holds the default value on creation for the dynamic_level field.
	menu.DefaultDynamicLevel = menuDescDynamicLevel.Default.(uint32)
	// menuDescRealPath is the schema descriptor for real_path field.
	menuDescRealPath := menuFields[19].Descriptor()
	// menu.DefaultRealPath holds the default value on creation for the real_path field.
	menu.DefaultRealPath = menuDescRealPath.Default.(string)
	oauthproviderMixin := schema.OauthProvider{}.Mixin()
	oauthproviderMixinFields0 := oauthproviderMixin[0].Fields()
	_ = oauthproviderMixinFields0
	oauthproviderFields := schema.OauthProvider{}.Fields()
	_ = oauthproviderFields
	// oauthproviderDescCreatedAt is the schema descriptor for created_at field.
	oauthproviderDescCreatedAt := oauthproviderMixinFields0[1].Descriptor()
	// oauthprovider.DefaultCreatedAt holds the default value on creation for the created_at field.
	oauthprovider.DefaultCreatedAt = oauthproviderDescCreatedAt.Default.(func() time.Time)
	// oauthproviderDescUpdatedAt is the schema descriptor for updated_at field.
	oauthproviderDescUpdatedAt := oauthproviderMixinFields0[2].Descriptor()
	// oauthprovider.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	oauthprovider.DefaultUpdatedAt = oauthproviderDescUpdatedAt.Default.(func() time.Time)
	// oauthprovider.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	oauthprovider.UpdateDefaultUpdatedAt = oauthproviderDescUpdatedAt.UpdateDefault.(func() time.Time)
	positionMixin := schema.Position{}.Mixin()
	positionMixinFields0 := positionMixin[0].Fields()
	_ = positionMixinFields0
	positionMixinFields1 := positionMixin[1].Fields()
	_ = positionMixinFields1
	positionMixinFields2 := positionMixin[2].Fields()
	_ = positionMixinFields2
	positionFields := schema.Position{}.Fields()
	_ = positionFields
	// positionDescCreatedAt is the schema descriptor for created_at field.
	positionDescCreatedAt := positionMixinFields0[1].Descriptor()
	// position.DefaultCreatedAt holds the default value on creation for the created_at field.
	position.DefaultCreatedAt = positionDescCreatedAt.Default.(func() time.Time)
	// positionDescUpdatedAt is the schema descriptor for updated_at field.
	positionDescUpdatedAt := positionMixinFields0[2].Descriptor()
	// position.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	position.DefaultUpdatedAt = positionDescUpdatedAt.Default.(func() time.Time)
	// position.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	position.UpdateDefaultUpdatedAt = positionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// positionDescStatus is the schema descriptor for status field.
	positionDescStatus := positionMixinFields1[0].Descriptor()
	// position.DefaultStatus holds the default value on creation for the status field.
	position.DefaultStatus = positionDescStatus.Default.(uint8)
	// positionDescSort is the schema descriptor for sort field.
	positionDescSort := positionMixinFields2[0].Descriptor()
	// position.DefaultSort holds the default value on creation for the sort field.
	position.DefaultSort = positionDescSort.Default.(uint32)
	roleMixin := schema.Role{}.Mixin()
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleMixinFields1 := roleMixin[1].Fields()
	_ = roleMixinFields1
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleMixinFields0[1].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleMixinFields0[2].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// role.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	role.UpdateDefaultUpdatedAt = roleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roleDescStatus is the schema descriptor for status field.
	roleDescStatus := roleMixinFields1[0].Descriptor()
	// role.DefaultStatus holds the default value on creation for the status field.
	role.DefaultStatus = roleDescStatus.Default.(uint8)
	// roleDescDefaultRouter is the schema descriptor for default_router field.
	roleDescDefaultRouter := roleFields[2].Descriptor()
	// role.DefaultDefaultRouter holds the default value on creation for the default_router field.
	role.DefaultDefaultRouter = roleDescDefaultRouter.Default.(string)
	// roleDescRemark is the schema descriptor for remark field.
	roleDescRemark := roleFields[3].Descriptor()
	// role.DefaultRemark holds the default value on creation for the remark field.
	role.DefaultRemark = roleDescRemark.Default.(string)
	// roleDescSort is the schema descriptor for sort field.
	roleDescSort := roleFields[4].Descriptor()
	// role.DefaultSort holds the default value on creation for the sort field.
	role.DefaultSort = roleDescSort.Default.(uint32)
	stockMixin := schema.Stock{}.Mixin()
	stockMixinFields0 := stockMixin[0].Fields()
	_ = stockMixinFields0
	stockMixinFields1 := stockMixin[1].Fields()
	_ = stockMixinFields1
	stockFields := schema.Stock{}.Fields()
	_ = stockFields
	// stockDescCreatedAt is the schema descriptor for created_at field.
	stockDescCreatedAt := stockMixinFields0[1].Descriptor()
	// stock.DefaultCreatedAt holds the default value on creation for the created_at field.
	stock.DefaultCreatedAt = stockDescCreatedAt.Default.(func() time.Time)
	// stockDescUpdatedAt is the schema descriptor for updated_at field.
	stockDescUpdatedAt := stockMixinFields0[2].Descriptor()
	// stock.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	stock.DefaultUpdatedAt = stockDescUpdatedAt.Default.(func() time.Time)
	// stock.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	stock.UpdateDefaultUpdatedAt = stockDescUpdatedAt.UpdateDefault.(func() time.Time)
	// stockDescStatus is the schema descriptor for status field.
	stockDescStatus := stockMixinFields1[0].Descriptor()
	// stock.DefaultStatus holds the default value on creation for the status field.
	stock.DefaultStatus = stockDescStatus.Default.(uint8)
	// stockDescIsRecommend is the schema descriptor for is_recommend field.
	stockDescIsRecommend := stockFields[2].Descriptor()
	// stock.DefaultIsRecommend holds the default value on creation for the is_recommend field.
	stock.DefaultIsRecommend = stockDescIsRecommend.Default.(bool)
	// stockDescID is the schema descriptor for id field.
	stockDescID := stockMixinFields0[0].Descriptor()
	// stock.DefaultID holds the default value on creation for the id field.
	stock.DefaultID = stockDescID.Default.(func() uuid.UUID)
	tokenMixin := schema.Token{}.Mixin()
	tokenMixinFields0 := tokenMixin[0].Fields()
	_ = tokenMixinFields0
	tokenMixinFields1 := tokenMixin[1].Fields()
	_ = tokenMixinFields1
	tokenFields := schema.Token{}.Fields()
	_ = tokenFields
	// tokenDescCreatedAt is the schema descriptor for created_at field.
	tokenDescCreatedAt := tokenMixinFields0[1].Descriptor()
	// token.DefaultCreatedAt holds the default value on creation for the created_at field.
	token.DefaultCreatedAt = tokenDescCreatedAt.Default.(func() time.Time)
	// tokenDescUpdatedAt is the schema descriptor for updated_at field.
	tokenDescUpdatedAt := tokenMixinFields0[2].Descriptor()
	// token.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	token.DefaultUpdatedAt = tokenDescUpdatedAt.Default.(func() time.Time)
	// token.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	token.UpdateDefaultUpdatedAt = tokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tokenDescStatus is the schema descriptor for status field.
	tokenDescStatus := tokenMixinFields1[0].Descriptor()
	// token.DefaultStatus holds the default value on creation for the status field.
	token.DefaultStatus = tokenDescStatus.Default.(uint8)
	// tokenDescID is the schema descriptor for id field.
	tokenDescID := tokenMixinFields0[0].Descriptor()
	// token.DefaultID holds the default value on creation for the id field.
	token.DefaultID = tokenDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinHooks2 := userMixin[2].Hooks()
	user.Hooks[0] = userMixinHooks2[0]
	userMixinInters2 := userMixin[2].Interceptors()
	user.Interceptors[0] = userMixinInters2[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userMixinFields1[0].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(uint8)
	// userDescHomePath is the schema descriptor for home_path field.
	userDescHomePath := userFields[4].Descriptor()
	// user.DefaultHomePath holds the default value on creation for the home_path field.
	user.DefaultHomePath = userDescHomePath.Default.(string)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[7].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
	// userDescDepartmentID is the schema descriptor for department_id field.
	userDescDepartmentID := userFields[8].Descriptor()
	// user.DefaultDepartmentID holds the default value on creation for the department_id field.
	user.DefaultDepartmentID = userDescDepartmentID.Default.(uint64)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.12.4"                                         // Version of ent codegen.
	Sum     = "h1:LddPnAyxls/O7DTXZvUGDj0NZIdGSu317+aoNLJWbD8=" // Sum of ent codegen.
)
