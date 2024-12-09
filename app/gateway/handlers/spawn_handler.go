package handlers

//
// type SpawnVaultRequest struct {
// 	Name string `json:"name"`
// }
//
// func SpawnVault(c echo.Context) error {
// 	ks, err := mpc.NewKeyset()
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	src, err := mpc.NewSource(ks)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	tk, err := src.OriginToken()
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	// Create the vault keyshare auth token
// 	kscid, err := tk.CID()
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	// Create the vault config
// 	dir, err := config.NewFS(config.GetVaultConfig(src.Address(), kscid.String()))
// 	path, err := clients.IPFSAdd(c, dir)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.Redirect(http.StatusFound, path)
// }
