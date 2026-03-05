package models

import "time"

type Project struct {
	ProjectID   int          `json:"project_id"`
	Title       string       `json:"title"`
	Description *string      `json:"description"`
	GithubLink  *string      `json:"github_link"`
	ImageURL    *string      `json:"image_url"`
	IsPublic    PublicStatus `json:"is_public"`
	DateCreated time.Time    `json:"date_created"`
	DateUpdated time.Time    `json:"date_updated"`
}

type WorkExperience struct {
	WorkExperienceID int          `json:"experience_id"`
	JobTitle         string       `json:"job_title"`
	CompanyName      string       `json:"company_name"`
	Description      *string      `json:"description"`
	RoleType         Role         `json:"role_type"`
	IsPublic         PublicStatus `json:"is_public"`
	DateStarted      time.Time    `json:"date_started"`
	DateEnded        *time.Time   `json:"date_ended"`
}

type Skill struct {
	SkillID   int           `json:"skill_id"`
	Name      string        `json:"name"`
	Thumbnail string        `json:"thumbnail"`
	Category  SkillCategory `json:"category"`
	Position  int           `json:"position"`
}

type Certification struct {
	CertificationID int          `json:"certification_id"`
	Name            string       `json:"name"`
	Company         string       `json:"company"`
	VerificationURL *string      `json:"verification_url"`
	DateIssued      time.Time    `json:"date_issued"`
	DateExpiration  *time.Time   `json:"date_expiration"`
	IsPublic        PublicStatus `json:"is_public"`
}

type BlogPost struct {
	PostID      int          `json:"post_id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	ContentURL  string       `json:"content_url"`
	Thumbnail   *string      `json:"thumbnail"`
	Slug        string       `json:"slug"`
	IsPublic    PublicStatus `json:"is_public"`
	DateCreated time.Time    `json:"date_created"`
	DateUpdated time.Time    `json:"date_updated"`
}

type ProjectDetail struct {
	Project
	Skills    []Skill    `json:"skills"`
	BlogPosts []BlogPost `json:"blog_posts"`
}

type WorkExperienceDetail struct {
	WorkExperience
	Skills []Skill `json:"skills"`
}

type BlogPostDetail struct {
	BlogPost
	Skills   []Skill   `json:"skills"`
	Projects []Project `json:"projects"`
}
