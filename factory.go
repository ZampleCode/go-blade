package blade

func NewBlade(viewPath, cachePath string) *Blade {
	blade := new(Blade)
	blade.loadViewPath = viewPath
	blade.cachePath = cachePath
	blade.bootstrap()

	return blade
}
