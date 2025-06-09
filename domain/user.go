package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                    uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName             string    `gorm:"type:varchar(150);not null" json:"first_name"`
	LastName              string    `gorm:"type:varchar(150);not null" json:"last_name"`
	DocumentType          string    `gorm:"type:varchar(50);not null" json:"document_type"`
	DocumentNumber        string    `gorm:"type:varchar(150);" json:"document_number"`
	Birthdate             time.Time `gorm:"type:date;not null" json:"birthdate"`
	Birthplace            string    `gorm:"type:varchar(150)" json:"birthplace"`
	Address               string    `gorm:"type:varchar(150)" json:"address"`
	Landline              string    `gorm:"type:varchar(50)" json:"landline"`
	Mobile                string    `gorm:"type:varchar(50)" json:"mobile"`
	Email                 string    `gorm:"type:text" json:"email"`
	Neighborhood          string    `gorm:"type:varchar(100)" json:"neighborhood"`
	Locality              string    `gorm:"type:varchar(100)" json:"locality"`
	SocioeconomicStratum  int       `json:"socioeconomic_stratum"`
	BloodType             string    `gorm:"type:varchar(15)" json:"blood_type"`
	Profession            string    `gorm:"type:varchar(100)" json:"profession"`
	Company               string    `gorm:"type:varchar(100)" json:"company"`
	Position              string    `gorm:"type:varchar(100)" json:"position"`
	MaritalStatus         string    `gorm:"type:varchar(50)" json:"marital_status"`
	LeaderName            string    `gorm:"type:varchar(150)" json:"leader_name"`
	ConversionDate        time.Time `gorm:"type:date" json:"conversion_date"`
	InvitedBy             string    `gorm:"type:varchar(150)" json:"invited_by"`
	AttendedOtherChurch   bool      `gorm:"default:false" json:"attended_other_church"`
	OtherChurchName       string    `gorm:"type:varchar(150)" json:"other_church_name"`
	EncounterDate         time.Time `gorm:"type:date" json:"encounter_date"`
	BaptismDate           time.Time `gorm:"type:date" json:"baptism_date"`
	MinisterialFunctionID uint      `gorm:"foreignKey:MinisterialFunctionID" json:"ministerial_function_id"`
	MinisterialFunction   string    `gorm:"type:varchar(100)" json:"ministerial_function"`
	LastSchoolLevelID     uint      `gorm:"foreignKey:LastSchoolLevelID" json:"last_school_level_id"`
	LastSchoolLevel       string    `gorm:"type:varchar(100)" json:"last_school_level"`
	PescaTeamID           uint      `gorm:"foreignKey:PescaTeamID" json:"pesca_team_id"`
	PescaTeam             string    `gorm:"type:varchar(100)" json:"pesca_team"` // Assuming this is a direct string for now, could be a separate entity
	SpouseName            string    `gorm:"type:varchar(150)" json:"spouse_name"`
	SpouseIsMember        bool      `gorm:"default:false" json:"spouse_is_member"`
	MarriageDate          time.Time `gorm:"type:date" json:"marriage_date"`
	ParentsNames          string    `gorm:"type:varchar(255)" json:"parents_names"`
	ParentsAreMembers     bool      `gorm:"default:false" json:"parents_are_members"`
	ChildrenNames         string    `gorm:"type:varchar(255)" json:"children_names"`
	HasAllergy            bool      `gorm:"default:false" json:"has_allergy"`
	DataAuthorization     bool      `gorm:"default:false" json:"data_authorization"`
}
