CREATE TABLE IF NOT EXISTS projects (
    project_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    github_link TEXT,
    image_url TEXT,
    is_public ENUM('Yes', 'No') NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS work_experience (
    experience_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    job_title TEXT NOT NULL,
    company_name TEXT NOT NULL,
    description TEXT,
    role_type ENUM('full-time', 'part-time', 'internship', 'contract') NOT NULL,
    is_public ENUM('Yes', 'No') NOT NULL,
    date_started TIMESTAMP NOT NULL,
    date_ended TIMESTAMP
);

CREATE TABLE IF NOT EXISTS skills (
    skill_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(55) NOT NULL,
    thumbnail TEXT NOT NULL,
    category ENUM('Programming Languages', 'Frameworks', 'Databases', 'Cloud Tools', 'DevOps', 'AI/ML') NOT NULL,
    position INT NOT NULL
);

CREATE TABLE IF NOT EXISTS certifications (
    certification_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(55) NOT NULL,
    company VARCHAR(55) NOT NULL,
    verification_url TEXT,
    date_issued TIMESTAMP NOT NULL,
    date_expiration TIMESTAMP,
    is_public ENUM('Yes', 'No') NOT NULL
);

CREATE TABLE IF NOT EXISTS blog_posts (
    post_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    content_url TEXT NOT NULL,
    thumbnail TEXT,
    slug VARCHAR(255) NOT NULL,
    is_public ENUM('Yes', 'No') NOT NULL,
    date_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS skill_project (
    skill_id INT NOT NULL,
    project_id INT NOT NULL,
    PRIMARY KEY (skill_id, project_id),
    FOREIGN KEY (skill_id) REFERENCES skills(skill_id),
    FOREIGN KEY (project_id) REFERENCES projects(project_id)
);

CREATE TABLE IF NOT EXISTS skill_experience (
    skill_id INT NOT NULL,
    experience_id INT NOT NULL,
    PRIMARY KEY (skill_id, experience_id),
    FOREIGN KEY (skill_id) REFERENCES skills(skill_id),
    FOREIGN KEY (experience_id) REFERENCES work_experience(experience_id)
);

CREATE TABLE IF NOT EXISTS skill_blog (
    skill_id INT NOT NULL,
    post_id INT NOT NULL,
    PRIMARY KEY (skill_id, post_id),
    FOREIGN KEY (skill_id) REFERENCES skills(skill_id),
    FOREIGN KEY (post_id) REFERENCES blog_posts(post_id)
);

CREATE TABLE IF NOT EXISTS project_blog (
    project_id INT NOT NULL,
    post_id INT NOT NULL,
    PRIMARY KEY (project_id, post_id),
    FOREIGN KEY (project_id) REFERENCES projects(project_id),
    FOREIGN KEY (post_id) REFERENCES blog_posts(post_id)
);

DELIMITER //

CREATE TRIGGER before_project_update
BEFORE UPDATE ON projects
FOR EACH ROW
BEGIN
    SET NEW.date_updated = CURRENT_TIMESTAMP;
END //

DELIMITER ;

DELIMITER //

CREATE TRIGGER before_blog_post_update
BEFORE UPDATE ON blog_posts
FOR EACH ROW
BEGIN
    SET NEW.date_updated = CURRENT_TIMESTAMP;
END //

DELIMITER ;