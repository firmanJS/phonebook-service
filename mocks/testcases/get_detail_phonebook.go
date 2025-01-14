package testcases

import (
	"database/sql"
	"errors"
	"time"

	"github.com/sapawarga/phonebook-service/model"
	"github.com/stretchr/testify/mock"
)

// GetDetailResponseRepository ...
type GetDetailResponseRepository struct {
	Result *model.PhoneBookResponse
	Error  error
}

// GetDetailResponseUsecase ...
type GetDetailResponseUsecase struct {
	Result *model.PhonebookDetail
	Error  error
}

// CategoryResponse ...
type CategoryResponse struct {
	Result string
	Error  error
}

// LocationResponse ...
type LocationResponse struct {
	Result string
	Error  error
}

// GetDetailPhonebook ...
type GetDetailPhonebook struct {
	Description         string
	UsecaseParams       int64
	GetDetailRequest    int64
	GetCategoryRequest  int64
	GetLocationRequest  int64
	MockUsecase         GetDetailResponseUsecase
	MockPhonebookDetail GetDetailResponseRepository
	MockCategory        CategoryResponse
	MockLocation        LocationResponse
}

var currentTime = time.Now().UTC()

// GetDetailPhonebookData ...
var GetDetailPhonebookData = []GetDetailPhonebook{
	{
		Description:        "success_get_detail",
		UsecaseParams:      1,
		GetDetailRequest:   1,
		GetCategoryRequest: 1,
		GetLocationRequest: 1,
		MockPhonebookDetail: GetDetailResponseRepository{
			Result: &model.PhoneBookResponse{
				ID:             1,
				PhoneNumbers:   sql.NullString{String: `[{"type":"phone", "phone_number":"+62812312131"]`, Valid: true},
				Description:    sql.NullString{String: "test case", Valid: true},
				Name:           sql.NullString{String: "test kantor", Valid: true},
				Address:        sql.NullString{String: "jalan panjang", Valid: true},
				RegencyID:      sql.NullInt64{Int64: 1, Valid: true},
				DistrictID:     sql.NullInt64{Int64: 10, Valid: true},
				VillageID:      sql.NullInt64{Int64: 100, Valid: true},
				Latitude:       sql.NullString{String: "-6.231928", Valid: true},
				Longitude:      sql.NullString{String: "0.988789", Valid: true},
				CoverImagePath: sql.NullString{String: "http://localhost:9080", Valid: true},
				Status:         sql.NullInt64{Int64: 10, Valid: true},
				CreatedAt:      sql.NullTime{Time: currentTime, Valid: true},
				UpdatedAt:      sql.NullTime{Time: currentTime, Valid: true},
				CategoryID:     sql.NullInt64{Int64: 1, Valid: true},
			},
			Error: nil,
		},
		MockCategory: CategoryResponse{
			Result: "category",
			Error:  nil,
		},
		MockLocation: LocationResponse{
			Result: mock.Anything,
			Error:  nil,
		},
		MockUsecase: GetDetailResponseUsecase{
			Result: &model.PhonebookDetail{
				ID:             1,
				Name:           "test kantor",
				CategoryID:     1,
				CategoryName:   "category",
				Address:        "jalan panjang",
				Description:    "test case",
				PhoneNumbers:   `[{"type":"phone", "phone_number":"+62812312131"]`,
				RegencyID:      1,
				RegencyName:    "mock.Anything",
				DistrictID:     10,
				DistrictName:   "mock.Anything",
				VillageID:      100,
				VillageName:    "mock.Anything",
				Latitude:       "-6.231928",
				Longitude:      "0.988789",
				CoverImagePath: "http://localhost:9080",
				Status:         10,
				CreatedAt:      currentTime,
				UpdatedAt:      currentTime,
			},
			Error: nil,
		},
	}, {
		Description:        "failed get phone book",
		UsecaseParams:      0,
		GetDetailRequest:   0,
		GetCategoryRequest: 0,
		GetLocationRequest: 0,
		MockPhonebookDetail: GetDetailResponseRepository{
			Result: nil,
			Error:  errors.New("id_is_invalid"),
		},
		MockCategory: CategoryResponse{
			Result: "",
			Error:  errors.New("id_is_invalid"),
		},
		MockLocation: LocationResponse{
			Result: "",
			Error:  errors.New("id_is_invalid"),
		},
		MockUsecase: GetDetailResponseUsecase{
			Result: nil,
			Error:  errors.New("id_is_invalid"),
		},
	}, {
		Description:        "failed get category",
		UsecaseParams:      1,
		GetDetailRequest:   1,
		GetCategoryRequest: 1,
		GetLocationRequest: 0,
		MockPhonebookDetail: GetDetailResponseRepository{
			Result: &model.PhoneBookResponse{
				ID:             1,
				PhoneNumbers:   sql.NullString{String: `[{"type":"phone", "phone_number":"+62812312131"]`, Valid: true},
				Description:    sql.NullString{String: "test case", Valid: true},
				Name:           sql.NullString{String: "test kantor", Valid: true},
				Address:        sql.NullString{String: "jalan panjang", Valid: true},
				RegencyID:      sql.NullInt64{Int64: 0, Valid: true},
				DistrictID:     sql.NullInt64{Int64: 0, Valid: true},
				VillageID:      sql.NullInt64{Int64: 0, Valid: true},
				Latitude:       sql.NullString{String: "-6.231928", Valid: true},
				Longitude:      sql.NullString{String: "0.988789", Valid: true},
				CoverImagePath: sql.NullString{String: "http://localhost:9080", Valid: true},
				Status:         sql.NullInt64{Int64: 10, Valid: true},
				CreatedAt:      sql.NullTime{Time: currentTime, Valid: true},
				UpdatedAt:      sql.NullTime{Time: currentTime, Valid: true},
				CategoryID:     sql.NullInt64{Int64: 1, Valid: true},
			},
			Error: nil,
		},
		MockCategory: CategoryResponse{
			Result: "",
			Error:  errors.New("id_is_invalid"),
		},
		MockLocation: LocationResponse{
			Result: "",
			Error:  errors.New("id_is_invalid"),
		},
		MockUsecase: GetDetailResponseUsecase{
			Result: nil,
			Error:  errors.New("id_is_invalid"),
		},
	}, {
		Description:        "failed get locations",
		UsecaseParams:      1,
		GetDetailRequest:   1,
		GetCategoryRequest: 0,
		GetLocationRequest: 0,
		MockPhonebookDetail: GetDetailResponseRepository{
			Result: &model.PhoneBookResponse{
				ID:             1,
				PhoneNumbers:   sql.NullString{String: `[{"type":"phone", "phone_number":"+62812312131"]`, Valid: true},
				Description:    sql.NullString{String: "test case", Valid: true},
				Name:           sql.NullString{String: "test kantor", Valid: true},
				Address:        sql.NullString{String: "jalan panjang", Valid: true},
				RegencyID:      sql.NullInt64{Int64: 0, Valid: true},
				DistrictID:     sql.NullInt64{Int64: 0, Valid: true},
				VillageID:      sql.NullInt64{Int64: 0, Valid: true},
				Latitude:       sql.NullString{String: "-6.231928", Valid: true},
				Longitude:      sql.NullString{String: "0.988789", Valid: true},
				CoverImagePath: sql.NullString{String: "http://localhost:9080", Valid: true},
				Status:         sql.NullInt64{Int64: 10, Valid: true},
				CreatedAt:      sql.NullTime{Time: currentTime, Valid: true},
				UpdatedAt:      sql.NullTime{Time: currentTime, Valid: true},
				CategoryID:     sql.NullInt64{Int64: 1, Valid: true},
			},
			Error: nil,
		},
		MockCategory: CategoryResponse{
			Result: "category",
			Error:  nil,
		},
		MockLocation: LocationResponse{
			Result: "",
			Error:  errors.New("id_is_invalid"),
		},
		MockUsecase: GetDetailResponseUsecase{
			Result: nil,
			Error:  errors.New("id_is_invalid"),
		},
	},
}

// DetailPhonebookDescription :
func DetailPhonebookDescription() []string {
	var arr = []string{}
	for _, data := range GetDetailPhonebookData {
		arr = append(arr, data.Description)
	}
	return arr
}
