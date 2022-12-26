// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountColumns holds the columns for the "account" table.
	AccountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "audit", Type: field.TypeUint8},
		{Name: "email", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "disable", Type: field.TypeBool},
		{Name: "pwd_error_num", Type: field.TypeInt32},
		{Name: "pwd_error_expire_time", Type: field.TypeInt32},
		{Name: "phone_token", Type: field.TypeString},
		{Name: "facebook", Type: field.TypeString},
		{Name: "line", Type: field.TypeString},
		{Name: "weibo", Type: field.TypeString},
		{Name: "google", Type: field.TypeString},
		{Name: "instagram", Type: field.TypeString},
		{Name: "linkedin", Type: field.TypeString},
		{Name: "language", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "country_code", Type: field.TypeInt32},
		{Name: "qrcode", Type: field.TypeString},
		{Name: "gender", Type: field.TypeBool},
		{Name: "birth_date", Type: field.TypeTime},
		{Name: "self_introduction", Type: field.TypeString},
		{Name: "cover", Type: field.TypeString},
		{Name: "photo", Type: field.TypeString, Nullable: true},
		{Name: "platform", Type: field.TypeBool},
		{Name: "last_login_time", Type: field.TypeTime},
		{Name: "credit", Type: field.TypeInt32},
		{Name: "heat_num", Type: field.TypeInt32},
		{Name: "designer_id", Type: field.TypeInt32},
		{Name: "identity", Type: field.TypeBool},
		{Name: "level", Type: field.TypeInt32},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "chain_wallet_id", Type: field.TypeString, Unique: true},
		{Name: "discord_aid", Type: field.TypeInt32},
		{Name: "ticket_balance", Type: field.TypeInt32},
		{Name: "ohdat_ticket_balance", Type: field.TypeInt32},
		{Name: "space_expedition_chance", Type: field.TypeBool},
		{Name: "bamb_locked_balance", Type: field.TypeFloat64},
		{Name: "bamb_claimable_balance", Type: field.TypeFloat64},
		{Name: "bamb_staking", Type: field.TypeFloat64},
		{Name: "bamblp_staking", Type: field.TypeFloat64},
		{Name: "bamblp_collect", Type: field.TypeFloat64},
		{Name: "chip_balance", Type: field.TypeFloat64},
		{Name: "bamboo_shoot_balance", Type: field.TypeInt32},
		{Name: "panda_airdrop_count", Type: field.TypeInt32},
		{Name: "first_login", Type: field.TypeBool},
		{Name: "passed_space_expedition_count", Type: field.TypeInt32},
		{Name: "passed_space_expedition_reward", Type: field.TypeFloat64},
		{Name: "reserved", Type: field.TypeString},
		{Name: "centralized_beer_caps_balance", Type: field.TypeInt32},
	}
	// AccountTable holds the schema information for the "account" table.
	AccountTable = &schema.Table{
		Name:       "account",
		Columns:    AccountColumns,
		PrimaryKey: []*schema.Column{AccountColumns[0]},
	}
	// QuestionnaireIqiyiColumns holds the columns for the "questionnaire_iqiyi" table.
	QuestionnaireIqiyiColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "wallet_address", Type: field.TypeString},
		{Name: "twitter_url", Type: field.TypeString},
		{Name: "invite_code", Type: field.TypeString},
		{Name: "q1", Type: field.TypeString},
		{Name: "q2", Type: field.TypeString},
		{Name: "q3", Type: field.TypeString},
		{Name: "q4", Type: field.TypeString},
		{Name: "q5", Type: field.TypeString},
		{Name: "q6", Type: field.TypeString},
		{Name: "q7", Type: field.TypeString},
		{Name: "balance", Type: field.TypeString, Nullable: true},
		{Name: "amount", Type: field.TypeInt32, Nullable: true},
		{Name: "transaction_count", Type: field.TypeInt32, Nullable: true},
		{Name: "answer", Type: field.TypeInt8, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// QuestionnaireIqiyiTable holds the schema information for the "questionnaire_iqiyi" table.
	QuestionnaireIqiyiTable = &schema.Table{
		Name:       "questionnaire_iqiyi",
		Columns:    QuestionnaireIqiyiColumns,
		PrimaryKey: []*schema.Column{QuestionnaireIqiyiColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountTable,
		QuestionnaireIqiyiTable,
	}
)

func init() {
	AccountTable.Annotation = &entsql.Annotation{
		Table: "account",
	}
	QuestionnaireIqiyiTable.Annotation = &entsql.Annotation{
		Table: "questionnaire_iqiyi",
	}
}
