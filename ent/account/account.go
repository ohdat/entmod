// Code generated by ent, DO NOT EDIT.

package account

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAudit holds the string denoting the audit field in the database.
	FieldAudit = "audit"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldDisable holds the string denoting the disable field in the database.
	FieldDisable = "disable"
	// FieldPwdErrorNum holds the string denoting the pwd_error_num field in the database.
	FieldPwdErrorNum = "pwd_error_num"
	// FieldPwdErrorExpireTime holds the string denoting the pwd_error_expire_time field in the database.
	FieldPwdErrorExpireTime = "pwd_error_expire_time"
	// FieldPhoneToken holds the string denoting the phone_token field in the database.
	FieldPhoneToken = "phone_token"
	// FieldFacebook holds the string denoting the facebook field in the database.
	FieldFacebook = "facebook"
	// FieldLine holds the string denoting the line field in the database.
	FieldLine = "line"
	// FieldWeibo holds the string denoting the weibo field in the database.
	FieldWeibo = "weibo"
	// FieldGoogle holds the string denoting the google field in the database.
	FieldGoogle = "google"
	// FieldInstagram holds the string denoting the instagram field in the database.
	FieldInstagram = "instagram"
	// FieldLinkedin holds the string denoting the linkedin field in the database.
	FieldLinkedin = "linkedin"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldCountryCode holds the string denoting the country_code field in the database.
	FieldCountryCode = "country_code"
	// FieldQrcode holds the string denoting the qrcode field in the database.
	FieldQrcode = "qrcode"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldBirthDate holds the string denoting the birth_date field in the database.
	FieldBirthDate = "birth_date"
	// FieldSelfIntroduction holds the string denoting the self_introduction field in the database.
	FieldSelfIntroduction = "self_introduction"
	// FieldCover holds the string denoting the cover field in the database.
	FieldCover = "cover"
	// FieldPhoto holds the string denoting the photo field in the database.
	FieldPhoto = "photo"
	// FieldPlatform holds the string denoting the platform field in the database.
	FieldPlatform = "platform"
	// FieldLastLoginTime holds the string denoting the last_login_time field in the database.
	FieldLastLoginTime = "last_login_time"
	// FieldCredit holds the string denoting the credit field in the database.
	FieldCredit = "credit"
	// FieldHeatNum holds the string denoting the heat_num field in the database.
	FieldHeatNum = "heat_num"
	// FieldDesignerID holds the string denoting the designer_id field in the database.
	FieldDesignerID = "designer_id"
	// FieldIdentity holds the string denoting the identity field in the database.
	FieldIdentity = "identity"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldChainWalletID holds the string denoting the chain_wallet_id field in the database.
	FieldChainWalletID = "chain_wallet_id"
	// FieldDiscordAid holds the string denoting the discord_aid field in the database.
	FieldDiscordAid = "discord_aid"
	// FieldTicketBalance holds the string denoting the ticket_balance field in the database.
	FieldTicketBalance = "ticket_balance"
	// FieldOhdatTicketBalance holds the string denoting the ohdat_ticket_balance field in the database.
	FieldOhdatTicketBalance = "ohdat_ticket_balance"
	// FieldSpaceExpeditionChance holds the string denoting the space_expedition_chance field in the database.
	FieldSpaceExpeditionChance = "space_expedition_chance"
	// FieldBambLockedBalance holds the string denoting the bamb_locked_balance field in the database.
	FieldBambLockedBalance = "bamb_locked_balance"
	// FieldBambClaimableBalance holds the string denoting the bamb_claimable_balance field in the database.
	FieldBambClaimableBalance = "bamb_claimable_balance"
	// FieldBambStaking holds the string denoting the bamb_staking field in the database.
	FieldBambStaking = "bamb_staking"
	// FieldBamblpStaking holds the string denoting the bamblp_staking field in the database.
	FieldBamblpStaking = "bamblp_staking"
	// FieldBamblpCollect holds the string denoting the bamblp_collect field in the database.
	FieldBamblpCollect = "bamblp_collect"
	// FieldChipBalance holds the string denoting the chip_balance field in the database.
	FieldChipBalance = "chip_balance"
	// FieldBambooShootBalance holds the string denoting the bamboo_shoot_balance field in the database.
	FieldBambooShootBalance = "bamboo_shoot_balance"
	// FieldPandaAirdropCount holds the string denoting the panda_airdrop_count field in the database.
	FieldPandaAirdropCount = "panda_airdrop_count"
	// FieldFirstLogin holds the string denoting the first_login field in the database.
	FieldFirstLogin = "first_login"
	// FieldPassedSpaceExpeditionCount holds the string denoting the passed_space_expedition_count field in the database.
	FieldPassedSpaceExpeditionCount = "passed_space_expedition_count"
	// FieldPassedSpaceExpeditionReward holds the string denoting the passed_space_expedition_reward field in the database.
	FieldPassedSpaceExpeditionReward = "passed_space_expedition_reward"
	// FieldReserved holds the string denoting the reserved field in the database.
	FieldReserved = "reserved"
	// FieldCentralizedBeerCapsBalance holds the string denoting the centralized_beer_caps_balance field in the database.
	FieldCentralizedBeerCapsBalance = "centralized_beer_caps_balance"
	// Table holds the table name of the account in the database.
	Table = "account"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldAudit,
	FieldEmail,
	FieldNickname,
	FieldAvatar,
	FieldPassword,
	FieldDisable,
	FieldPwdErrorNum,
	FieldPwdErrorExpireTime,
	FieldPhoneToken,
	FieldFacebook,
	FieldLine,
	FieldWeibo,
	FieldGoogle,
	FieldInstagram,
	FieldLinkedin,
	FieldLanguage,
	FieldPhone,
	FieldCountryCode,
	FieldQrcode,
	FieldGender,
	FieldBirthDate,
	FieldSelfIntroduction,
	FieldCover,
	FieldPhoto,
	FieldPlatform,
	FieldLastLoginTime,
	FieldCredit,
	FieldHeatNum,
	FieldDesignerID,
	FieldIdentity,
	FieldLevel,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCreatedAt,
	FieldChainWalletID,
	FieldDiscordAid,
	FieldTicketBalance,
	FieldOhdatTicketBalance,
	FieldSpaceExpeditionChance,
	FieldBambLockedBalance,
	FieldBambClaimableBalance,
	FieldBambStaking,
	FieldBamblpStaking,
	FieldBamblpCollect,
	FieldChipBalance,
	FieldBambooShootBalance,
	FieldPandaAirdropCount,
	FieldFirstLogin,
	FieldPassedSpaceExpeditionCount,
	FieldPassedSpaceExpeditionReward,
	FieldReserved,
	FieldCentralizedBeerCapsBalance,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}