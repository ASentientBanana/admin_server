package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string
	Stack       string
	Description string
	Live        string
	Github      string
	Image       string
}

// project_id = models.AutoField(primary_key=True)
// name = models.CharField(max_length=100,null=False)
// stack = models.CharField(max_length=150,null=False)
// description = models.CharField(max_length=200,null=False)
// live = models.CharField(max_length=100,null=True)
// github = models.CharField(max_length=100,null=True)
// image = models.FileField(upload_to='static/assets',null=True)
