package mocks

import "github.com/Oleg-Skalozub/testtask/src/domain/entity"

var (
	ArrayDataResponse = []entity.DataResponse{
		{
			EventType: "1",
			Result:    255,
		},
		{
			EventType: "2",
			Result:    22,
		},
	}
)
