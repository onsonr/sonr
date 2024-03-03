## Commit Message 

**Issue:** #121 Routing for HTMX Authentication Views

**Type:** feature

### Authentication ViewComponents
- `auth.templ`: New HTML template for the authentication page
- `auth_templ.go`: New Go view rendering for the authentication page

### Login ViewComponents
- `form.templ`: Modified HTML template for the login form. Now supports HTMX.
- `form_templ.go`: Modified Go view rendering. Added HTMX libraries and integrated new HTMX routing features.
- `page.templ`: Modified HTML template for login page, mainly in forms calling htmx.js functions
- `page_templ.go`: Modified Go view rendering for the login page to handle new htmx requests

### Register ViewComponents
- `form.templ`: Altered HTML template for register form. Now supports htmx routing.
- `form_templ.go`: Adjusted Go view rendering. Added htmx dependencies and integrated new HTMX routing features.
- `page.templ`: Tweaked HTML template for register page, mainly in forms calling htmx.js functions
- `page_templ.go`: Adapted Go view rendering for register page to work with HTMX.

### Landing Page
- `page_templ.go`: A minor refactoring made to accommodate new auth routing.

*N.B.: Docs and changelog have been updated. Ref: `changes.md`.*

