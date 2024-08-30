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
	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/require"
)

func TestGetListCategory(t *testing.T) {

	categories := []*db.Category{
		&db.Category{
			CateID:   101,
			CateName: &[]string{"X"}[0],
		},
		&db.Category{
			CateID:   102,
			CateName: &[]string{"y"}[0],
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
					GetAllCategories(gomock.Any()).
					Times(1).
					Return(categories, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.Equal(t, len(categories), rowExpected)
			},
		},
		{
			"InternalError",
			gin.H{},
			func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAllCategories(gomock.Any()).
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

			url := "/api/category/"
			request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)
		})
	}
}

func TestCreateCategory(t *testing.T) {
	//desc := "kuliner"
	cate := &db.Category{
		CateID:   0,
		CateName: &[]string{"RotiBakar"}[0],
	}

	args := cate.CateName

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"cate_name": cate.CateName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCategory(gomock.Any(), args).
					Times(1).
					Return(cate, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				//requireBodyMatchCategory(t, cate, recorder.Body)
			},
		},
		{
			"BadRequestBody",
			gin.H{
				"cate_names": cate.CateName,
			},
			func(store *mockdb.MockStore) {},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
			},
		},
		{
			"DuplicateCateName",
			gin.H{
				"cate_name": cate.CateName,
			},
			func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCategory(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, &pgconn.PgError{ConstraintName: "cate_name_uq"})
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
			},
		},
		{
			"InternalError",
			gin.H{
				"cate_name": cate.CateName,
			},
			func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCategory(gomock.Any(), gomock.Any()).
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

			url := "/api/category/"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)

		})
	}
}

func TestFindCategory(t *testing.T) {
	cate := &db.Category{
		CateID:   100,
		CateName: &[]string{"Course"}[0],
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
					GetCategoryByID(gomock.Any(), cate.CateID).
					Times(1).
					Return(cate, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			"CategoryNotFound",
			gin.H{},
			func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCategoryByID(gomock.Any(), cate.CateID).
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
					GetCategoryByID(gomock.Any(), cate.CateID).
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

			url := "/api/category/" + strconv.Itoa(int(cate.CateID))
			request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	category := &db.Category{
		CateID:   101,
		CateName: &[]string{"Course"}[0],
	}

	args := db.UpdateCategoryParams{
		CateID:   101,
		CateName: category.CateName,
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
				"cate_id":   category.CateID,
				"cate_name": category.CateName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateCategory(gomock.Any(), args).
					Times(1).
					Return(category, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			"CategoryNotFound",
			gin.H{},
			func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateCategory(gomock.Any(), gomock.Any()).
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
					UpdateCategory(gomock.Any(), gomock.Any()).
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

			url := "/api/category/" + strconv.Itoa(int(category.CateID))
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)
		})
	}
}
