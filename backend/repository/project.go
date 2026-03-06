package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/RemyPaulJr/portfolio-go/backend/models"
)

type ProjectRepo struct {
	db *sql.DB
}

func (p ProjectRepo) GetProject(ctx context.Context, id int) (models.ProjectDetail, error) {

	var project models.Project
	var projectSkills []models.Skill
	var projectBlogs []models.BlogPost

	projectCheck := p.db.QueryRowContext(ctx, "SELECT project_id, title, description, github_link, image_url, is_public, date_created, date_updated FROM projects WHERE project_id = ? ", id)
	if err := projectCheck.Scan(&project.ProjectID, &project.Title, &project.Description, &project.GithubLink, &project.ImageURL, &project.IsPublic, &project.DateCreated, &project.DateUpdated); err != nil {
		log.Print("Error querying project: ", err)
		return models.ProjectDetail{}, err
	} // no error then this returns one row, just one project, and we will proceed to check this project for it's associated skills and blog posts

	projectSkillsCheck, err := p.db.QueryContext(ctx, "SELECT s.skill_id, s.name, s.thumbnail, s.category, s.position FROM skills s INNER JOIN skill_project sp ON s.skill_id = sp.skill_id WHERE sp.project_id = ? ", id)
	if err != nil {
		log.Print("Error querying skill_project: ", err)
		return models.ProjectDetail{}, err
	}

	defer projectSkillsCheck.Close() // defer close connection to avoid connection leaks

	for projectSkillsCheck.Next() { // loop through our skills table because for every project we can have multiple skills associated
		var skill models.Skill

		if err := projectSkillsCheck.Scan(&skill.SkillID, &skill.Name, &skill.Thumbnail, &skill.Category, &skill.Position); err != nil {
			log.Print("Error scanning project skills: ", err)
			return models.ProjectDetail{}, err
		}
		projectSkills = append(projectSkills, skill) // appending our skill {skill_id, name, etc} to our models.Skill slice
	}

	if err := projectSkillsCheck.Err(); err != nil { // This will catch if the context was canceled OR a database error occurred
		log.Print("Error from loop checking skills apart of a project: ", err)
		return models.ProjectDetail{}, err
	}

	projectBlogsCheck, err := p.db.QueryContext(ctx, "SELECT bp.post_id, bp.title, bp.description, bp.content_url, bp.thumbnail, bp.slug, bp.is_public, bp.date_created, bp.date_updated FROM blog_posts bp INNER JOIN project_blog pb ON bp.post_id = pb.post_id WHERE pb.project_id = ? ", id)
	if err != nil {
		log.Print("Error querying project_blog: ", err)
		return models.ProjectDetail{}, err
	}

	defer projectBlogsCheck.Close()

	for projectBlogsCheck.Next() {
		var blog models.BlogPost

		if err := projectBlogsCheck.Scan(&blog.PostID, &blog.Title, &blog.Description, &blog.ContentURL, &blog.Thumbnail, &blog.Slug, &blog.IsPublic, &blog.DateCreated, &blog.DateUpdated); err != nil {
			log.Print("Error scanning blog_posts table: ", err)
			return models.ProjectDetail{}, err
		}
		projectBlogs = append(projectBlogs, blog)
	}

	if err := projectBlogsCheck.Err(); err != nil {
		log.Print("Error from loop checking blog posts apart of a project: ", err)
		return models.ProjectDetail{}, err
	}

	return models.ProjectDetail{Project: project, Skills: projectSkills, BlogPosts: projectBlogs}, nil
}

func (p ProjectRepo) GetProjects(ctx context.Context) ([]models.Project, error) {

	var projects []models.Project

	projectCheck, err := p.db.QueryContext(ctx, "SELECT project_id, title, description, github_link, image_url, is_public, date_created, date_updated FROM projects")
	if err != nil {
		log.Print("Error querying projects table for all rows: ", err)
		return []models.Project{}, err
	}

	defer projectCheck.Close()

	for projectCheck.Next() {
		var project models.Project

		err := projectCheck.Scan(&project.ProjectID, &project.Title, &project.Description, &project.GithubLink, &project.ImageURL, &project.IsPublic, &project.DateCreated, &project.DateUpdated)
		if err != nil {
			log.Print("Error storing project result query in project struct: ", err)
			return []models.Project{}, err
		}
		projects = append(projects, project)

	}

	return projects, nil
}
