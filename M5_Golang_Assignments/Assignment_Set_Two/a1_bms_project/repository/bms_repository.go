package repository

import (
	"a1_bms_project/model"
	"database/sql"
)

type BlogRepository struct {
	DB *sql.DB
}

func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{DB: db}
}

func (repo *BlogRepository) CreateBlog(blog *model.Blog) (*model.Blog, error) {
	query := "INSERT INTO blogs (title, content, author) VALUES (?, ?, ?)"
	result, err := repo.DB.Exec(query, blog.Title, blog.Content, blog.Author)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	blog.ID = int(id)
	return blog, nil
}

func (repo *BlogRepository) GetBlog(id int) (*model.Blog, error) {
	query := "SELECT * FROM blogs WHERE id = ?"
	row := repo.DB.QueryRow(query, id)

	var blog model.Blog
	err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (repo *BlogRepository) GetAllBlogs() ([]model.Blog, error) {
	query := "SELECT * FROM blogs ORDER BY timestamp DESC"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []model.Blog
	for rows.Next() {
		var blog model.Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (repo *BlogRepository) UpdateBlog(blog *model.Blog) (*model.Blog, error) {
	query := "UPDATE blogs SET title = ?, content = ?, author = ? WHERE id = ?"
	_, err := repo.DB.Exec(query, blog.Title, blog.Content, blog.Author, blog.ID)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (repo *BlogRepository) DeleteBlog(id int) error {
	query := "DELETE FROM blogs WHERE id = ?"
	_, err := repo.DB.Exec(query, id)
	return err
}
