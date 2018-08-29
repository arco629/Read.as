package readas

import (
	"fmt"
	"net/http"
)

func handleViewHostMeta(app *app, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Server", serverName)
	w.Header().Set("Content-Type", "application/xrd+xml; charset=utf-8")

	meta := `<?xml version="1.0" encoding="UTF-8"?>
<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0">
  <Link rel="lrdd" type="application/xrd+xml" template="` + app.cfg.Host + `/.well-known/webfinger?resource={uri}"/>
</XRD>`
	fmt.Fprintf(w, meta)

	return nil
}
