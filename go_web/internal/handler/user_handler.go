package handler

import (
	"git_truongvudinh/go_web/internal/domain/dto"
	"git_truongvudinh/go_web/internal/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
có 2 loại context trong go:
- context.Context và gin.Context
+ đối với context.Context quản lý vòng đời của tasks, context.Context cho phép
hủy bỏ các tác vụ khi không còn cần thiết
+ truyền dữ liệu giữa các goroutines ? (có thể giống với channel hoặc một thư viện
cải tiến hơn của channel)
+ context.Context còn cho Timeout và Deadline buộc chúng phải hoàn thành
trong một khoảng thời gian nhất định hoặc bị hủy bỏ nếu vượt quá thời gian này
+ context.Context thường dùng ở các lớp để xử lý logic nghiệp vụ vd : usecases layer và repositories layer

+ đối với gin.Context thì dễ hiểu hơn, nó dùng để xử lý các HTTP request
+ truyền dữ liệu giữa middleware và handler ? khá khó hiểu
+ quản lý response vd : trả về json, xml, html, và thiết lập các HTTP status code
tương tự giống ResponseEntity trong spring
+ gin.Context thì dùng ở controller layer (handler layer)
*/

type UserHandler struct {
	userService usecases.IUserService
}

func NewUserHandler(userService usecases.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateNewUser(ctx *gin.Context) {
	var request dto.CreateUserRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// 500
	newUser, err := h.userService.CreateNewUser(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	// trả vể mã 201
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":         newUser.ID,
			"first_name": newUser.FirstName,
			"last_name":  newUser.LastName,
			"email":      newUser.Email,
			"created_at": newUser.CreatedAt.Format("2003-12-23 15:04:05"),
			"updated_at": newUser.UpdatedAt.Format("2003-12-23 15:04:05"),
		},
	})
}

func (h *UserHandler) GetUserById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userService.GetUserByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"password":   user.Password,
		"created_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at": user.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}
