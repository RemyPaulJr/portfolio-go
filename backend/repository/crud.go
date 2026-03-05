package repository

import "github.com/RemyPaulJr/portfolio-go/backend/models"

type ProjectRepository interface {
	GetProject(id int) (models.Project, error)
	GetProjects() ([]models.Project, error)

	CreateProject(project models.Project) (models.Project, error)
	UpdateProject(project models.Project) (models.Project, error)
	DeleteProject(id int) error

	//CreateProjectSkill(projectskill models.ProjectDetail) (models.ProjectDetail, error)
}

type WorkExperienceRepository interface {
	GetWorkExperience(id int) (models.WorkExperience, error)
	GetWorkExperiences() ([]models.WorkExperience, error)

	CreateWorkExperience(workExperience models.WorkExperience) (models.WorkExperience, error)
	UpdateWorkExperience(workExperience models.WorkExperience) (models.WorkExperience, error)
	DeleteWorkExperience(id int) error
}

type SkillRepository interface {
	GetSkill(id int) (models.Skill, error)
	GetSkills() ([]models.Skill, error)

	CreateSkill(skill models.Skill) (models.Skill, error)
	UpdateSkill(skill models.Skill) (models.Skill, error)
	DeleteSkill(id int) error
}

type CertificationRepository interface {
	GetCerfication(id int) (models.Certification, error)
	GetCertification() ([]models.Certification, error)

	CreateCertification(certification models.Certification) (models.Certification, error)
	UpdateCertification(certification models.Certification) (models.Certification, error)
	DeleteCertification(id int) error
}

type BlogPostRepository interface {
	GetBlogPost(id int) (models.BlogPost, error)
	GetBlogPosts() ([]models.BlogPost, error)

	CreateBlogPost(blogPost models.BlogPost) (models.BlogPost, error)
	UpdateBlogPost(blogPost models.BlogPost) (models.BlogPost, error)
	DeleteBlogPost(id int) error
}
