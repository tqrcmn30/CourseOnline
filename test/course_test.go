package test

import (
	"bytes"
	db "courseonline/db/sqlc"
	mockdb "courseonline/services/mocks"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
)

func TestGetListCourse(t *testing.T) {

	courses := []*db.Course{
		&db.Course{
			CoursID:   101,
			CoursName: &[]string{"Presiden"}[0],
		},
		&db.Course{
			CoursID:   102,
			CoursName: &[]string{"Wakil Presiden"}[0],
		},
	}

	rowExpected := 2

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAllCourses(gomock.Any()).
					Times(1).
					Return(courses, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.Equal(t, len(courses), rowExpected)
			},
		},
		{
			"InternalError",
			gin.H{},
			func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAllCourses(gomock.Any()).
					Times(1).
					Return(nil, fmt.Errorf("internal error"))
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			testCase.buildStubs(store)

			server := newTestHttpServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(testCase.body)
			require.NoError(t, err)

			url := "/api/course/"
			request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)
		})
	}
}

func TestCreateCourse(t *testing.T) {
	//desc := "sejarah"
	cour := &db.Course{
		CoursID:   0,
		CoursName: &[]string{"Presiden"}[0],
		CoursDesc: &[]string{"Sejarah"}[0],
	}

	args := db.CreateCourseParams{
		CoursName: cour.CoursName,
		CoursDesc: cour.CoursDesc,
	}

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"cours_name": cour.CoursName,
				"cours_desc": cour.CoursDesc,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCourse(gomock.Any(), args).
					Times(1).
					Return(cour, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				//requireBodyMatchcours(t, cour, recorder.Body)
			},
		},
		{
			"BadRequestBody",
			gin.H{
				"cours_names": cour.CoursName,
				"cours_desc":  cour.CoursDesc,
			},
			func(store *mockdb.MockStore) {},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
			},
		},
		{
			"DuplicateCoursename",
			gin.H{
				"cours_name": cour.CoursName,
				"cours_desc": cour.CoursDesc,
			},
			func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCourse(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, &pgconn.PgError{ConstraintName: "cours_name_uq"})
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
			},
		},
		{
			"InternalError",
			gin.H{
				"cours_name": cour.CoursName,
				"cours_desc": cour.CoursDesc,
			},
			func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCourse(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, fmt.Errorf("internal error"))
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			testCase.buildStubs(store)

			server := newTestHttpServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(testCase.body)
			require.NoError(t, err)

			url := "/api/course/"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)

		})
	}
}

func TestUpdateCourse(t *testing.T) {
	cour := &db.Course{
		CoursID:   101,
		CoursName: &[]string{"Presiden"}[0],
		CoursDesc: &[]string{"Sejarah"}[0],
	}

	args := db.UpdateCourseParams{
		CoursID:   101,
		CoursName: cour.CoursName,
	}

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"cours_id":   cour.CoursID,
				"cours_name": cour.CoursName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateCourse(gomock.Any(), args).
					Times(1).
					Return(cour, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			"CourseNotFound",
			gin.H{},
			func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateCourse(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, nil)
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			"InternalError",
			gin.H{},
			func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateCourse(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, fmt.Errorf("internal error"))
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			testCase.buildStubs(store)

			server := newTestHttpServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(testCase.body)
			require.NoError(t, err)

			url := "/api/course/" + strconv.Itoa(int(cour.CoursID))
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)
		})
	}
}

func TestFindCourse(t *testing.T) {
	cour := &db.Course{
		CoursID:   101,
		CoursName: &[]string{"Presiden"}[0],
		CoursDesc: &[]string{"Sejarah"}[0],
	}

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCourseByID(gomock.Any(), cour.CoursID).
					Times(1).
					Return(cour, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			"CourseNotFound",
			gin.H{},
			func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCourseByID(gomock.Any(), cour.CoursID).
					Times(1).
					Return(nil, nil)
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			"InternalError",
			gin.H{},
			func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCourseByID(gomock.Any(), cour.CoursID).
					Times(1).
					Return(nil, fmt.Errorf("internal error"))
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			testCase.buildStubs(store)

			server := newTestHttpServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(testCase.body)
			require.NoError(t, err)

			url := "/api/course/" + strconv.Itoa(int(cour.CoursID))
			request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)
		})
	}
}
