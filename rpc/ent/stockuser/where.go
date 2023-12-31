// Code generated by ent, DO NOT EDIT.

package stockuser

import (
	"time"

	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldStatus, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldDeletedAt, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldUsername, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldPassword, v))
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldNickname, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldDescription, v))
}

// HomePath applies equality check predicate on the "home_path" field. It's identical to HomePathEQ.
func HomePath(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldHomePath, v))
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldMobile, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldEmail, v))
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldAvatar, v))
}

// LastLoginInfo applies equality check predicate on the "last_login_info" field. It's identical to LastLoginInfoEQ.
func LastLoginInfo(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldLastLoginInfo, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldNotNull(FieldStatus))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldNotNull(FieldDeletedAt))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameIsNil applies the IsNil predicate on the "username" field.
func UsernameIsNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldIsNull(FieldUsername))
}

// UsernameNotNil applies the NotNil predicate on the "username" field.
func UsernameNotNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldNotNull(FieldUsername))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContainsFold(FieldPassword, v))
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldNickname, v))
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldNickname, v))
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldNickname, vs...))
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldNickname, vs...))
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldNickname, v))
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldNickname, v))
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldNickname, v))
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldNickname, v))
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContains(FieldNickname, v))
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasPrefix(FieldNickname, v))
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasSuffix(FieldNickname, v))
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEqualFold(FieldNickname, v))
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContainsFold(FieldNickname, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContainsFold(FieldDescription, v))
}

// HomePathEQ applies the EQ predicate on the "home_path" field.
func HomePathEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldHomePath, v))
}

// HomePathNEQ applies the NEQ predicate on the "home_path" field.
func HomePathNEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldHomePath, v))
}

// HomePathIn applies the In predicate on the "home_path" field.
func HomePathIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldHomePath, vs...))
}

// HomePathNotIn applies the NotIn predicate on the "home_path" field.
func HomePathNotIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldHomePath, vs...))
}

// HomePathGT applies the GT predicate on the "home_path" field.
func HomePathGT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldHomePath, v))
}

// HomePathGTE applies the GTE predicate on the "home_path" field.
func HomePathGTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldHomePath, v))
}

// HomePathLT applies the LT predicate on the "home_path" field.
func HomePathLT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldHomePath, v))
}

// HomePathLTE applies the LTE predicate on the "home_path" field.
func HomePathLTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldHomePath, v))
}

// HomePathContains applies the Contains predicate on the "home_path" field.
func HomePathContains(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContains(FieldHomePath, v))
}

// HomePathHasPrefix applies the HasPrefix predicate on the "home_path" field.
func HomePathHasPrefix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasPrefix(FieldHomePath, v))
}

// HomePathHasSuffix applies the HasSuffix predicate on the "home_path" field.
func HomePathHasSuffix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasSuffix(FieldHomePath, v))
}

// HomePathEqualFold applies the EqualFold predicate on the "home_path" field.
func HomePathEqualFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEqualFold(FieldHomePath, v))
}

// HomePathContainsFold applies the ContainsFold predicate on the "home_path" field.
func HomePathContainsFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContainsFold(FieldHomePath, v))
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldMobile, v))
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldMobile, v))
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldMobile, vs...))
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldMobile, vs...))
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldMobile, v))
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldMobile, v))
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldMobile, v))
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldMobile, v))
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContains(FieldMobile, v))
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasPrefix(FieldMobile, v))
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasSuffix(FieldMobile, v))
}

// MobileIsNil applies the IsNil predicate on the "mobile" field.
func MobileIsNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldIsNull(FieldMobile))
}

// MobileNotNil applies the NotNil predicate on the "mobile" field.
func MobileNotNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldNotNull(FieldMobile))
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEqualFold(FieldMobile, v))
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContainsFold(FieldMobile, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContainsFold(FieldEmail, v))
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldAvatar, v))
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldAvatar, v))
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldAvatar, vs...))
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldAvatar, vs...))
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldAvatar, v))
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldAvatar, v))
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldAvatar, v))
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldAvatar, v))
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContains(FieldAvatar, v))
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasPrefix(FieldAvatar, v))
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasSuffix(FieldAvatar, v))
}

// AvatarIsNil applies the IsNil predicate on the "avatar" field.
func AvatarIsNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldIsNull(FieldAvatar))
}

// AvatarNotNil applies the NotNil predicate on the "avatar" field.
func AvatarNotNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldNotNull(FieldAvatar))
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEqualFold(FieldAvatar, v))
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContainsFold(FieldAvatar, v))
}

// LastLoginInfoEQ applies the EQ predicate on the "last_login_info" field.
func LastLoginInfoEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEQ(FieldLastLoginInfo, v))
}

// LastLoginInfoNEQ applies the NEQ predicate on the "last_login_info" field.
func LastLoginInfoNEQ(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNEQ(FieldLastLoginInfo, v))
}

// LastLoginInfoIn applies the In predicate on the "last_login_info" field.
func LastLoginInfoIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldIn(FieldLastLoginInfo, vs...))
}

// LastLoginInfoNotIn applies the NotIn predicate on the "last_login_info" field.
func LastLoginInfoNotIn(vs ...string) predicate.StockUser {
	return predicate.StockUser(sql.FieldNotIn(FieldLastLoginInfo, vs...))
}

// LastLoginInfoGT applies the GT predicate on the "last_login_info" field.
func LastLoginInfoGT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGT(FieldLastLoginInfo, v))
}

// LastLoginInfoGTE applies the GTE predicate on the "last_login_info" field.
func LastLoginInfoGTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldGTE(FieldLastLoginInfo, v))
}

// LastLoginInfoLT applies the LT predicate on the "last_login_info" field.
func LastLoginInfoLT(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLT(FieldLastLoginInfo, v))
}

// LastLoginInfoLTE applies the LTE predicate on the "last_login_info" field.
func LastLoginInfoLTE(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldLTE(FieldLastLoginInfo, v))
}

// LastLoginInfoContains applies the Contains predicate on the "last_login_info" field.
func LastLoginInfoContains(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContains(FieldLastLoginInfo, v))
}

// LastLoginInfoHasPrefix applies the HasPrefix predicate on the "last_login_info" field.
func LastLoginInfoHasPrefix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasPrefix(FieldLastLoginInfo, v))
}

// LastLoginInfoHasSuffix applies the HasSuffix predicate on the "last_login_info" field.
func LastLoginInfoHasSuffix(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldHasSuffix(FieldLastLoginInfo, v))
}

// LastLoginInfoIsNil applies the IsNil predicate on the "last_login_info" field.
func LastLoginInfoIsNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldIsNull(FieldLastLoginInfo))
}

// LastLoginInfoNotNil applies the NotNil predicate on the "last_login_info" field.
func LastLoginInfoNotNil() predicate.StockUser {
	return predicate.StockUser(sql.FieldNotNull(FieldLastLoginInfo))
}

// LastLoginInfoEqualFold applies the EqualFold predicate on the "last_login_info" field.
func LastLoginInfoEqualFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldEqualFold(FieldLastLoginInfo, v))
}

// LastLoginInfoContainsFold applies the ContainsFold predicate on the "last_login_info" field.
func LastLoginInfoContainsFold(v string) predicate.StockUser {
	return predicate.StockUser(sql.FieldContainsFold(FieldLastLoginInfo, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StockUser) predicate.StockUser {
	return predicate.StockUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StockUser) predicate.StockUser {
	return predicate.StockUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.StockUser) predicate.StockUser {
	return predicate.StockUser(sql.NotPredicates(p))
}
