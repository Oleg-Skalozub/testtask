package services

import (
	"github.com/Oleg-Skalozub/testtask/src/infrastructure/errorscan"
	"github.com/Oleg-Skalozub/testtask/src/mock"
	"github.com/jinzhu/gorm"
	"testing"

	"github.com/Oleg-Skalozub/testtask/src/domain/entity"
	"github.com/golang/mock/gomock"
)

func TestFetch_FetchData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockDataRepository(ctrl)
	client := mocks.NewMockClientInterface(ctrl)
	fetch := fetch{
		dataRepository: repo,
		client:         client,
	}

	testCases := []struct {
		name    string
		dbError error
		err     error
	}{
		{"test without errors", nil, nil},
		{"test with DB error", errorscan.BigDayValueError, nil},
		{"test with DB error", errorscan.EmptyResultError, nil},
	}

	for _, tc := range testCases {
		if tc.dbError == nil && tc.err == nil {
			repo.EXPECT().GetData(gomock.Any(), gomock.Any()).Return(mocks.ArrayDataResponse, nil)

		} else if tc.dbError == errorscan.EmptyResultError {
			repo.EXPECT().GetData(gomock.Any(), gomock.Any()).Return(nil, tc.dbError)
			client.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(entity.Contain{}, nil)
		} else if tc.dbError != nil {
			repo.EXPECT().GetData(gomock.Any(), gomock.Any()).Return(nil, tc.dbError)
		}

		_, err := fetch.FetchData(3, 4)

		if tc.dbError == nil {
			if err != nil {
				t.Error(err)
			}
		} else if tc.dbError == errorscan.EmptyResultError {
			if err != nil {
				t.Error(err)
			}
		} else if tc.dbError != nil {
			if err == nil {
				t.Error("expected not <nil> error")
			}
		}
	}

}

func TestFetch_SaveData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockDataRepository(ctrl)

	fetch := fetch{
		dataRepository: repo,
	}

	wg.Add(3)
	repo.EXPECT().SaveData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	fetch.SaveData(2, 4, 6, []entity.Event{})
}

func TestFetch_GetData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockDataRepository(ctrl)

	fetch := fetch{
		dataRepository: repo,
	}

	testCases := []struct {
		name    string
		dbError error
		err     error
	}{
		{"test without errors", nil, nil},
		{"test with DB error", errorscan.EmptyResultError, nil},
		{"test with empty result", nil, gorm.ErrInvalidSQL},
	}

	for _, tc := range testCases {

		if tc.dbError == nil && tc.err == nil {
			repo.EXPECT().GetData(gomock.Any(), gomock.Any()).Return(mocks.ArrayDataResponse, nil)
		} else if tc.err != nil && tc.dbError == nil {
			repo.EXPECT().GetData(gomock.Any(), gomock.Any()).Return([]entity.DataResponse{}, nil)
		} else {
			repo.EXPECT().GetData(gomock.Any(), gomock.Any()).Return(nil, tc.dbError)
		}

		_, err := fetch.GetData(3, 4)

		if tc.dbError == nil && tc.err == nil {
			if err != nil {
				t.Error(err)
			}
		} else if tc.dbError != nil || tc.err != nil {
			if err == nil {
				t.Error("expected not <nil> error")
			}
		}
	}

}
