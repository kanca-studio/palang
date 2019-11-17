package routing

//func checkAuth(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		token := r.Header.Get("Authorization")
//		if err := userManager.ValidateToken(token); err != nil {
//			http.Error(w, "Unauthorized", http.StatusUnauthorized)
//		}
//		context.Set(r, "token", token)
//		next(w, r)
//	}
//}
