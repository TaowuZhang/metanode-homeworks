package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/repository"
	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/utils"

	"github.com/gin-gonic/gin"
)

type apiResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func newTestRouter(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)
	name := strings.NewReplacer("/", "-", " ", "-").Replace(t.Name())
	db, err := repository.InitDB(fmt.Sprintf("file:%s?mode=memory&cache=shared", name))
	if err != nil {
		t.Fatalf("init db: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("get sql db: %v", err)
	}
	t.Cleanup(func() { _ = sqlDB.Close() })
	return SetupRouter(db, "test-secret", 3600)
}

func request(router http.Handler, method, path, token string, body any) *httptest.ResponseRecorder {
	var buf bytes.Buffer
	if body != nil {
		_ = json.NewEncoder(&buf).Encode(body)
	}
	req := httptest.NewRequest(method, path, &buf)
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func decodeData[T any](t *testing.T, w *httptest.ResponseRecorder) T {
	t.Helper()
	var resp apiResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("decode response: %v body=%s", err, w.Body.String())
	}
	var out T
	if err := json.Unmarshal(resp.Data, &out); err != nil {
		t.Fatalf("decode data: %v body=%s", err, w.Body.String())
	}
	return out
}

func registerAndLogin(t *testing.T, router http.Handler, username string) string {
	t.Helper()
	w := request(router, http.MethodPost, "/api/v1/auth/register", "", map[string]string{
		"username": username, "email": username + "@example.com", "password": "password123",
	})
	if w.Code != http.StatusCreated {
		t.Fatalf("register status=%d body=%s", w.Code, w.Body.String())
	}
	w = request(router, http.MethodPost, "/api/v1/auth/login", "", map[string]string{
		"username": username, "password": "password123",
	})
	if w.Code != http.StatusOK {
		t.Fatalf("login status=%d body=%s", w.Code, w.Body.String())
	}
	data := decodeData[struct {
		Token string `json:"token"`
	}](t, w)
	if data.Token == "" {
		t.Fatal("empty token")
	}
	return data.Token
}

func TestBlogAPIFlow(t *testing.T) {
	router := newTestRouter(t)

	if w := request(router, http.MethodGet, "/health", "", nil); w.Code != http.StatusOK {
		t.Fatalf("health status=%d", w.Code)
	}
	if w := request(router, http.MethodPost, "/api/v1/posts", "", map[string]string{"title": "x", "content": "y"}); w.Code != http.StatusUnauthorized {
		t.Fatalf("unauth create status=%d", w.Code)
	}
	if w := request(router, http.MethodPost, "/api/v1/posts", "not-a-token", map[string]string{"title": "x", "content": "y"}); w.Code != http.StatusUnauthorized {
		t.Fatalf("invalid-token create status=%d", w.Code)
	}

	authorToken := registerAndLogin(t, router, "author")
	otherToken := registerAndLogin(t, router, "other")

	if w := request(router, http.MethodGet, "/api/v1/profile", authorToken, nil); w.Code != http.StatusOK {
		t.Fatalf("profile status=%d body=%s", w.Code, w.Body.String())
	}

	w := request(router, http.MethodPost, "/api/v1/posts", authorToken, map[string]string{"title": "first", "content": "body"})
	if w.Code != http.StatusCreated {
		t.Fatalf("create post status=%d body=%s", w.Code, w.Body.String())
	}
	post := decodeData[struct {
		ID uint `json:"ID"`
	}](t, w)
	if post.ID == 0 {
		t.Fatal("missing post id")
	}
	postPath := fmt.Sprintf("/api/v1/posts/%d", post.ID)
	commentsPath := postPath + "/comments"

	if w := request(router, http.MethodGet, postPath, "", nil); w.Code != http.StatusOK {
		t.Fatalf("get post status=%d body=%s", w.Code, w.Body.String())
	}
	if w := request(router, http.MethodPut, postPath, otherToken, map[string]string{"title": "hack"}); w.Code != http.StatusForbidden {
		t.Fatalf("non-author update status=%d body=%s", w.Code, w.Body.String())
	}
	if w := request(router, http.MethodPut, postPath, authorToken, map[string]string{"title": "updated", "content": "body2"}); w.Code != http.StatusOK {
		t.Fatalf("author update status=%d body=%s", w.Code, w.Body.String())
	}
	if w := request(router, http.MethodGet, "/api/v1/posts", "", nil); w.Code != http.StatusOK {
		t.Fatalf("list status=%d", w.Code)
	}
	if w := request(router, http.MethodPost, commentsPath, otherToken, map[string]string{"content": "nice"}); w.Code != http.StatusCreated {
		t.Fatalf("comment status=%d body=%s", w.Code, w.Body.String())
	}
	if w := request(router, http.MethodGet, commentsPath, "", nil); w.Code != http.StatusOK {
		t.Fatalf("comments status=%d", w.Code)
	}
	if w := request(router, http.MethodDelete, postPath, otherToken, nil); w.Code != http.StatusForbidden {
		t.Fatalf("non-author delete status=%d body=%s", w.Code, w.Body.String())
	}
	if w := request(router, http.MethodDelete, postPath, authorToken, nil); w.Code != http.StatusOK {
		t.Fatalf("author delete status=%d body=%s", w.Code, w.Body.String())
	}
	if w := request(router, http.MethodGet, postPath, "", nil); w.Code != http.StatusNotFound {
		t.Fatalf("deleted post status=%d body=%s", w.Code, w.Body.String())
	}
}

func TestBlogAPIValidationAndAuthEdges(t *testing.T) {
	router := newTestRouter(t)

	if w := request(router, http.MethodGet, "/api/v1/posts/not-a-number", "", nil); w.Code != http.StatusBadRequest {
		t.Fatalf("invalid id status=%d body=%s", w.Code, w.Body.String())
	}
	if w := request(router, http.MethodGet, "/api/v1/posts/999", "", nil); w.Code != http.StatusNotFound {
		t.Fatalf("missing post status=%d body=%s", w.Code, w.Body.String())
	}
	if w := request(router, http.MethodPost, "/api/v1/auth/register", "", map[string]string{"username": "short"}); w.Code != http.StatusBadRequest {
		t.Fatalf("invalid register status=%d body=%s", w.Code, w.Body.String())
	}

	token := registerAndLogin(t, router, "validation-user")
	if w := request(router, http.MethodPost, "/api/v1/posts", token, map[string]string{"title": "", "content": "body"}); w.Code != http.StatusBadRequest {
		t.Fatalf("invalid create status=%d body=%s", w.Code, w.Body.String())
	}
	if w := request(router, http.MethodPost, "/api/v1/posts/bad/comments", token, map[string]string{"content": "x"}); w.Code != http.StatusBadRequest {
		t.Fatalf("invalid comment post id status=%d body=%s", w.Code, w.Body.String())
	}

	expiredToken, err := utils.GenerateToken(1, "test-secret", -time.Second)
	if err != nil {
		t.Fatalf("generate expired token: %v", err)
	}
	if w := request(router, http.MethodGet, "/api/v1/profile", expiredToken, nil); w.Code != http.StatusUnauthorized {
		t.Fatalf("expired token status=%d body=%s", w.Code, w.Body.String())
	}
}
