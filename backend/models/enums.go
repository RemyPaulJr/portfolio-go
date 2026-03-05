package models

type PublicStatus string  // Custom type for is_public enum
type Role string          // Custom type for role_type enum
type SkillCategory string // Custom type for category enum

const (
	StatusPublic  PublicStatus  = "Yes"
	StatusPrivate PublicStatus  = "No"
	FullTime      Role          = "full-time"
	PartTime      Role          = "part-time"
	Intern        Role          = "internship"
	Contract      Role          = "contract"
	Programming   SkillCategory = "Programming Languages"
	Framework     SkillCategory = "Frameworks"
	Database      SkillCategory = "Databases"
	Cloud         SkillCategory = "Cloud Tools"
	DevOps        SkillCategory = "DevOps"
	Ai            SkillCategory = "AI/ML"
)
