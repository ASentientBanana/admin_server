package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `json:"name" xml:"name"`
	Stack       string `json:"stack" xml:"stack"`
	Description string `json:"description" xml:"description"`
	Live        string `json:"live" xml:"live"`
	Github      string `json:"github" xml:"github"`
	Image       string `json:"image" xml:"image"`
	Position    int    `json:"position" xml:"position"`
}

// project_id = models.AutoField(primary_key=True)
// name = models.CharField(max_length=100,null=False)
// stack = models.CharField(max_length=150,null=False)
// description = models.CharField(max_length=200,null=False)
// live = models.CharField(max_length=100,null=True)
// github = models.CharField(max_length=100,null=True)
// image = models.FileField(upload_to='static/assets',null=True)
