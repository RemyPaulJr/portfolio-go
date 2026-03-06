package repository

import (
	"context"

	"github.com/RemyPaulJr/portfolio-go/backend/models"
)

/*
	The purpose of this package is to initialize my repository interfaces.
	These are the method signatures that i will define to do specific tasks.
	Like GET, CREATE, UPDATE, DELETE, ADD, and REMOVE.
	They will go hand in hand with my CRUD operations.
	As a Admin for my site, these interfaces allow me to perform operations and
	my code will understand the operation and how to properly manage my database.
	And pass information to my frontend to display to me and visitors of my site.
*/

type ProjectRepository interface {
	GetProject(ctx context.Context, id int) (models.ProjectDetail, error) // For single project we want rich information so we include the projectdetail
	GetProjects(ctx context.Context) ([]models.Project, error)            // For all products we want a lean return so just the project without the skills and blog posts

	CreateProject(ctx context.Context, project models.Project) (models.Project, error)
	UpdateProject(ctx context.Context, project models.Project) (models.Project, error)
	DeleteProject(ctx context.Context, id int) error

	AddSkillToProject(ctx context.Context, skillID int, projectID int) error // I only want to worry about the skill and project id. At this point the project is already loaded so i will update the skill associated with that project.
	RemoveSkillFromProject(ctx context.Context, skillID int, projectID int) error
}

type WorkExperienceRepository interface {
	GetWorkExperience(ctx context.Context, id int) (models.WorkExperienceDetail, error)
	GetWorkExperiences(ctx context.Context) ([]models.WorkExperience, error)

	CreateWorkExperience(ctx context.Context, workExperience models.WorkExperience) (models.WorkExperience, error)
	UpdateWorkExperience(ctx context.Context, workExperience models.WorkExperience) (models.WorkExperience, error)
	DeleteWorkExperience(ctx context.Context, id int) error

	AddSkillToExperience(ctx context.Context, skillID int, experienceID int) error
	RemoveSkillFromExperience(ctx context.Context, skillID int, experienceID int) error
}

type SkillRepository interface {
	GetSkill(ctx context.Context, id int) (models.Skill, error)
	GetSkills(ctx context.Context) ([]models.Skill, error)

	CreateSkill(ctx context.Context, skill models.Skill) (models.Skill, error)
	UpdateSkill(ctx context.Context, skill models.Skill) (models.Skill, error)
	DeleteSkill(ctx context.Context, id int) error
}

type CertificationRepository interface {
	GetCertification(ctx context.Context, id int) (models.Certification, error)
	GetCertifications(ctx context.Context) ([]models.Certification, error)

	CreateCertification(ctx context.Context, certification models.Certification) (models.Certification, error)
	UpdateCertification(ctx context.Context, certification models.Certification) (models.Certification, error)
	DeleteCertification(id int) error
}

type BlogPostRepository interface {
	GetBlogPost(ctx context.Context, id int) (models.BlogPostDetail, error)
	GetBlogPosts(ctx context.Context) ([]models.BlogPost, error)

	CreateBlogPost(ctx context.Context, blogPost models.BlogPost) (models.BlogPost, error)
	UpdateBlogPost(ctx context.Context, blogPost models.BlogPost) (models.BlogPost, error)
	DeleteBlogPost(ctx context.Context, id int) error

	AddProjectToBlog(ctx context.Context, projectID int, blogPostID int) error
	RemoveProjectFromBlog(ctx context.Context, projectID int, blogPostID int) error

	AddSkillToBlog(ctx context.Context, skillID int, blogPostID int) error
	RemoveSkillFromBlog(ctx context.Context, skillID int, blogPostID int) error
}
