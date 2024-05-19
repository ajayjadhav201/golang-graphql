package database

import (
	"errors"
	"fmt"

	"github.com/ajayjadhav201/golang-graphql/graph/model"
)

type Database interface {
	CreateJobListing(model.CreateJobListingInput) (*model.JobListing, error)
	UpdateJobListing(string, model.UpdateJobListingInput) (*model.JobListing, error)
	DeleteJobListing(string) (*model.DeleteJobResponse, error)
	GetJobs() ([]*model.JobListing, error)
	GetJobById(string) (*model.JobListing, error)
}

type database struct {
	//
	model.JobListing
}

func NewDatabase() *database {
	return &database{}
}

func (d *database) CreateJobListing(input model.CreateJobListingInput) (*model.JobListing, error) {
	fmt.Println("ajaj ", input)
	return nil, errors.New("aj crete job listing not implemented")
}

func (d *database) UpdateJobListing(id string, input model.UpdateJobListingInput) (*model.JobListing, error) {
	return nil, errors.New("aj update job listing not implemented")
}

func (d *database) DeleteJobListing(id string) (*model.DeleteJobResponse, error) {
	return nil, errors.New("aj delete job listing not implemented")
}

func (*database) GetJobs() ([]*model.JobListing, error) {
	return []*model.JobListing{
		&model.JobListing{
			ID: "1",
		},
		&model.JobListing{
			ID: "2",
		},
		&model.JobListing{
			ID: "3",
		},
		&model.JobListing{
			ID: "4",
		},
		&model.JobListing{
			ID: "5",
		},
		&model.JobListing{
			ID: "6",
		},
		&model.JobListing{
			ID: "7",
		},
		&model.JobListing{
			ID: "8",
		},
		&model.JobListing{
			ID: "9",
		},
		&model.JobListing{
			ID: "10",
		},
	}, nil
	// return nil, errors.New("aj get jobs not implemented")
}

func (d *database) GetJobById(id string) (*model.JobListing, error) {
	return nil, errors.New("aj get job by id not implemented")
}
