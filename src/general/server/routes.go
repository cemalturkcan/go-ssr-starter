package server

import (
	"projectesc/src/routes/blog"
	"projectesc/src/routes/detail"
	"projectesc/src/routes/filter"
	"projectesc/src/routes/index"
)

func (s *FiberServer) RegisterFiberRoutes() {
	index.Register(s.App)
	filter.Register(s.App)
	blog.Register(s.App)
	detail.Register(s.App)
}
