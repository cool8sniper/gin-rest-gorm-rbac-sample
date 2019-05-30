package utils

//
//import (
//	"log"
//	"reflect"
//	"runtime"
//	"strconv"
//	"strings"
//
//	"github.com/gin-gonic/gin"
//)
//
//type HandlerDecoratored func(gin.HandlerFunc, ...string) gin.HandlerFunc
//
//func Decorator(h gin.HandlerFunc, decors ...HandlerDecoratored) gin.HandlerFunc {
//	for i := range decors {
//		d := decors[len(decors)-1-i] // iterate in reverse
//		h = d(h)
//	}
//	return h
//}
//
//func CheckToken(user_id uint, token_in_request string) bool {
//	key := "user:" + strconv.FormatUint(uint64(user_id), 10) + ":token"
//	token, err := redis.GetRedisKey(key)
//
//	if err == false {
//		log.Println("Can not find token key: in redis.", token_in_request)
//		return false
//	}
//
//	if token == token_in_request {
//		log.Println("token verification success.")
//		return true
//	} else {
//		return false
//	}
//}
//
//func CheckParamAndHeader(input gin.HandlerFunc, http_params ...string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		http_params_local := append([]string{"param:user_id", "header:token"}, http_params...)
//		required_params_str := strings.Join(http_params_local, ", ")
//		required_params_str = "Required parameters include: " + required_params_str
//		log.Println(http_params_local, required_params_str, len(http_params_local))
//
//		for _, v := range http_params_local {
//			ret := strings.Split(v, ":")
//
//			switch ret[0] {
//			case "header":
//				header := c.Request.Header.Get(ret[1])
//
//				if header == "" {
//					c.JSON(200, gin.H{
//						"code":   3,
//						"result": "failed",
//						"msg":    required_params_str + ". Missing " + v,
//					})
//					return
//				}
//				if ret[1] == "token" {
//					user_id_str, _ := c.GetQuery("user_id")
//					user_id, _ := strconv.Atoi(user_id_str)
//
//					if CheckToken(uint(user_id), header) == true {
//						log.Println("token verification success.")
//					} else {
//						c.JSON(200, gin.H{
//							"code":   4,
//							"result": "failed",
//							"msg":    "Login timeout. Please re-login.",
//						})
//						return
//					}
//				}
//			case "param":
//				_, err := c.GetQuery(ret[1])
//				if err == false {
//					c.JSON(200, gin.H{
//						"code":   3,
//						"result": "failed",
//						"msg":    required_params_str + ". Missing " + v,
//					})
//					return
//
//				}
//			case "body":
//				body_param := c.PostForm(ret[1])
//
//				if body_param == "" {
//					c.JSON(200, gin.H{
//						"code":   3,
//						"result": "failed",
//						"msg":    required_params_str + ". Missing " + v,
//					})
//					return
//				}
//			default:
//				log.Println("Unsupported checking type: %s", ret[0])
//			}
//		}
//		input(c)
//	}
//}
//
//func CheckPermission(input gin.HandlerFunc, http_params ...string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		user_id_str, _ := c.GetQuery("user_id")
//		user_id, _ := strconv.Atoi(user_id_str)
//		function_name_str := runtime.FuncForPC(reflect.ValueOf(input).Pointer()).Name()
//
//		function_name_array := strings.Split(function_name_str, "/")
//		module_method := strings.Split(function_name_array[len(function_name_array)-1], ".")
//		module := module_method[0]
//		method := module_method[1]
//
//		log.Logger.Debugf("module: %s, method: %s", module, method)
//		permissions := GetUserRolePermission(uint(user_id), 0)
//		method_info := mysql.GetMethodByName(method)
//
//		for i := range permissions {
//			if (permissions[i].Method_id == method_info.ID) && (permissions[i].Module_id == method_info.Module_id) {
//				log.Println("user: %d has the permission to access module: %s, method: %s", user_id, module, method)
//
//				input(c)
//				return
//			}
//		}
//
//		c.JSON(200, gin.H{
//			"code":   6,
//			"result": "failed",
//			"msg":    "User have no permission to access this API.",
//		})
//		return
//	}
//}
